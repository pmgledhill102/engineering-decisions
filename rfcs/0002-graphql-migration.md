---
rfc: 2
title: "Migrate API Gateway from REST to GraphQL"
status: proposed
authors:
  - Paul Gledhill (@pmgledhill102)
created: 2026-03-18
updated: 2026-03-18
decision-date:
supersedes:
superseded-by:
---

# Migrate API Gateway from REST to GraphQL

## Summary

This RFC proposes migrating the public API gateway from REST to GraphQL. The
current REST API requires clients to make multiple round-trips to assemble views,
leading to over-fetching, under-fetching, and poor mobile performance. A GraphQL
gateway would allow clients to request exactly the data they need in a single
query, reducing payload sizes and network round-trips.

## Background

The platform currently exposes a REST API gateway that aggregates data from 12
backend microservices. The gateway was built three years ago when the platform
had 3 services and a single web client.

Since then, the platform has grown significantly:

- **12 microservices** behind the gateway (up from 3)
- **4 client platforms**: web SPA, iOS, Android, and a partner API
- **47 REST endpoints**, many of which return overlapping data
- **Average of 3.2 API calls per screen** on mobile clients to assemble a single
  view

Pain points with the current REST API:

- **Over-fetching**: The `/users/{id}` endpoint returns 34 fields. The mobile
  app uses 8 of them on the profile screen. The remaining 26 fields are
  transferred and discarded.
- **Under-fetching**: To display a dashboard, the web client calls
  `/users/{id}`, `/users/{id}/orders`, and `/users/{id}/notifications` in
  sequence. Each call waits for the previous one.
- **Endpoint proliferation**: Teams have created view-specific endpoints
  (`/mobile/dashboard`, `/web/dashboard`) to work around the above, creating
  maintenance burden and inconsistency.
- **Documentation drift**: REST endpoint documentation is manually maintained
  and frequently out of date. There is no machine-readable schema enforced at
  the gateway level.

## Business Justification

Mobile performance directly impacts revenue. Industry benchmarks show that each
additional second of load time reduces conversion by 7%. Our current mobile
dashboard requires 3 sequential API calls averaging 340ms total. A single
GraphQL query could reduce this to one round-trip.

Additional business benefits:

- **Reduced bandwidth costs**: Eliminating over-fetching reduces payload sizes
  by an estimated 60% for mobile clients.
- **Faster feature development**: Frontend teams can build new views without
  requesting backend changes to create view-specific endpoints.
- **Self-documenting API**: GraphQL's introspection provides always-accurate
  schema documentation, eliminating documentation drift.
- **Partner API simplification**: Partners currently receive the same bloated
  responses as internal clients. GraphQL would let partners request only the
  fields they need.

## Proposal

### Architecture

Replace the existing REST gateway with a GraphQL gateway that sits in front of
the same 12 backend microservices. Backend services remain unchanged — the
GraphQL gateway handles query resolution and orchestration.

```text
                    ┌─────────────────┐
Clients ──────────► │  GraphQL Gateway │
                    └────────┬────────┘
                             │
           ┌─────────────────┼─────────────────┐
           ▼                 ▼                  ▼
      ┌─────────┐     ┌──────────┐      ┌───────────┐
      │ Users   │     │  Orders  │      │  Payments  │  ... (12 services)
      │ Service │     │  Service │      │  Service   │
      └─────────┘     └──────────┘      └───────────┘
```

### Technology Choice

We propose using **Apollo Server** (Node.js) with **Apollo Federation** for the
gateway layer. Apollo Federation allows each backend team to define their own
GraphQL subgraph, which the gateway composes into a unified schema.

Key reasons for Apollo:

- Mature ecosystem with strong TypeScript support
- Federation model maps well to our existing microservice ownership boundaries
- Built-in performance tracing and metrics
- Large community and extensive documentation

### Schema Design

Each microservice team defines a subgraph schema for their domain. The gateway
composes these into a single federated schema.

Example subgraph (Users service):

```graphql
type User @key(fields: "id") {
  id: ID!
  email: String!
  displayName: String!
  createdAt: DateTime!
  orders: [Order!]!
  notifications: [Notification!]!
}

type Query {
  user(id: ID!): User
  users(filter: UserFilter, limit: Int = 20, offset: Int = 0): UserConnection!
}
```

### Query Complexity Limits

To prevent abusive or accidental expensive queries, the gateway will enforce:

- **Maximum query depth**: 7 levels
- **Maximum query complexity score**: 1000 (weighted by field resolver cost)
- **Timeout**: 10 seconds per query

### Caching Strategy

- **CDN-level caching**: Use `Cache-Control` headers with `@cacheControl`
  directive for public, read-heavy queries.
- **DataLoader pattern**: Batch and deduplicate backend service calls within a
  single query resolution using DataLoader.
- **Redis query cache**: Cache resolved query results for frequently-accessed,
  slowly-changing data (e.g., product catalog).

### Client Migration

Clients migrate incrementally. Both REST and GraphQL endpoints run in parallel
during the transition period:

1. **Phase 1**: Deploy GraphQL gateway alongside existing REST gateway. No
   client changes.
2. **Phase 2**: Migrate the mobile dashboard (highest impact) to GraphQL.
   Measure performance improvement.
3. **Phase 3**: Migrate remaining mobile screens.
4. **Phase 4**: Migrate web SPA.
5. **Phase 5**: Migrate partner API with versioned GraphQL schema.
6. **Phase 6**: Deprecate REST endpoints with 6-month sunset timeline.

## Alternatives Considered

### BFF (Backend for Frontend) pattern

Create dedicated backend-for-frontend services for each client platform (web
BFF, mobile BFF, partner BFF). This solves the over-fetching problem but creates
3 separate aggregation layers to maintain. Each new feature requires changes in
both the backend service and the relevant BFFs. Rejected because the maintenance
burden scales linearly with the number of client platforms.

### REST with sparse fieldsets (JSON:API)

Adopt JSON:API specification with `?fields[user]=id,email,displayName` query
parameters. This addresses over-fetching but does not solve the N+1 call problem
for assembling composite views. Clients would still need multiple round-trips.
Considered as a lighter-weight option but rejected because it does not address
the primary pain point.

### gRPC gateway

Replace REST with a gRPC gateway. gRPC offers strong typing, code generation,
and efficient binary serialization. However, gRPC has poor browser support
(requires gRPC-Web proxy), limited tooling for frontend developers, and would
require significant changes to all client applications. The developer experience
gap for frontend teams was the deciding factor against this option.

### Keep REST, add GraphQL for mobile only

Deploy GraphQL as an additional layer for mobile clients only, keeping REST for
web and partners. Rejected because this creates two API paradigms to maintain and
does not address the underlying problem — it just defers it.

## Migration and Rollout

### Phase 1: Gateway deployment (Weeks 1-4)

Deploy Apollo Gateway in the staging environment. Implement subgraphs for the
3 most-used services (Users, Orders, Notifications). Run synthetic traffic to
validate correctness.

### Phase 2: Mobile pilot (Weeks 5-8)

Migrate the mobile dashboard screen to GraphQL. This is the highest-impact
screen with 3 sequential REST calls. Run A/B test comparing REST vs GraphQL
performance with 10% of mobile traffic.

### Phase 3: Full mobile migration (Weeks 9-16)

Migrate remaining mobile screens based on Phase 2 learnings. Remaining backend
teams implement their subgraphs.

### Phase 4: Web and partner migration (Weeks 17-24)

Migrate web SPA and partner API. Begin REST deprecation timeline.

### Monitoring

- Track query latency p50, p95, p99 at the gateway
- Monitor backend service call volume (should decrease as DataLoader batches)
- Alert on query complexity violations
- Dashboard for schema composition health

## Risks and Mitigations

| Risk | Likelihood | Impact | Mitigation |
|------|-----------|--------|------------|
| N+1 resolver problem causes backend overload | High | High | DataLoader pattern for batching; per-service rate limiting |
| Complex queries cause gateway timeouts | Medium | Medium | Query complexity limits and depth restrictions |
| Team unfamiliarity with GraphQL slows development | Medium | Medium | Internal workshop series; pair programming during initial subgraph development |
| Schema composition conflicts between teams | Low | High | CI validation of schema composition; schema review in PR process |
| Apollo vendor lock-in | Low | Medium | Federation spec is open; subgraphs are portable to other federation-compatible gateways |

## References

- [Apollo Federation documentation](https://www.apollographql.com/docs/federation/)
- [GraphQL best practices](https://graphql.org/learn/best-practices/)
- [DataLoader pattern](https://github.com/graphql/dataloader)
- [Netflix: Our learnings from adopting GraphQL](https://netflixtechblog.com/beyond-rest-1b76f7c20ef6)
