---
rfc: 1
title: "AI-Augmented Decision Process for Engineering"
status: proposed
authors:
  - Paul Gledhill (@pmgledhill102)
created: 2026-02-14
updated: 2026-02-16
decision-date:
supersedes:
superseded-by:
---

# AI-Augmented Decision Process for Engineering

## Summary

This RFC proposes a structured decision process for cross-team engineering
decisions. Engineers submit proposals as pull requests, automated checks validate
format and quality, and AI personas review proposals against shared engineering
principles from multiple specialized perspectives. The process produces durable,
machine-readable decision records optimized for both human readers and AI agents.

## Background

Most engineering organizations lack a structured process for cross-team technical
decisions. Decisions about standards, shared infrastructure, and architectural
patterns are made through ad hoc discussions in meetings, chat threads, and email
chains. This leads to several problems:

- **Lost context**: Decision rationale is not recorded, leading to repeated
  debates about the same topics.
- **Inconsistent review**: Some proposals receive deep scrutiny while others are
  approved without adequate review from affected teams.
- **No prior art**: Rejected ideas are not documented, so engineers waste time
  proposing approaches that were previously considered and dismissed.
- **Uneven participation**: Decisions are often made by whoever is in the room
  rather than by systematically involving affected stakeholders.

RFC processes are well-established at companies like HashiCorp, Squarespace,
Uber, and Oxide Computer. This proposal adapts those patterns and adds
AI-augmented review as a force multiplier for thoroughness — reducing the human
effort required while improving the quality of review coverage.

## Business Justification

The cost of poor architectural decisions compounds over time. A single
undocumented decision to adopt an inappropriate technology can create years of
technical debt. The cost of this process is minimal (a few hours per proposal to
write and review) compared to the cost of reversing bad decisions or maintaining
inconsistent systems.

Specific business benefits:

- **Reduced rework**: Systematic multi-perspective review catches problems before
  implementation begins.
- **Faster onboarding**: New engineers can read accepted RFCs to understand why
  systems are built the way they are.
- **Better alignment**: Cross-team decisions are visible and reviewable by all
  stakeholders.
- **Preserved knowledge**: Both accepted and rejected proposals remain as
  searchable, machine-readable prior art.
- **Lower review burden**: AI personas handle the broad analysis, freeing human
  reviewers to focus on domain-specific judgment calls.

## Proposal

### RFC Lifecycle

RFCs progress through a defined lifecycle:

```text
draft --> proposed --> accepted
                  \-> rejected
                  \-> withdrawn

accepted --> superseded
```

- **draft**: Author iterates on the proposal in a draft PR. CI validates
  format.
- **proposed**: PR marked ready. AI personas review. 10 business day review
  period starts.
- **accepted/rejected**: Sponsor decides. Both outcomes are merged to preserve
  the decision record.
- **superseded**: A newer RFC replaces this one.

### Template

Each RFC uses a structured template with YAML frontmatter for machine-readable
metadata and required sections: Summary, Background, Business Justification,
Proposal, Alternatives Considered, Migration and Rollout, and Risks and
Mitigations.

The YAML frontmatter ensures every decision record is parseable by agents and
automation. The structured sections ensure consistency across proposals,
making it straightforward for both humans and AI to locate specific information.

### Validation

A Go tool (`rfclint`) validates RFC format automatically:

- Frontmatter fields are present and valid
- RFC number matches the filename
- Title is consistent between frontmatter and H1 heading
- Status is a valid lifecycle value
- All required sections are present and non-empty

### AI Review Personas

Nine AI personas review proposals from specialized perspectives:

1. **Security & Privacy** — threat models, auth, secrets, compliance
2. **Reliability & Operations** — failure modes, observability, blast radius
3. **Developer Experience** — API design, cognitive load, migration burden
4. **Simplicity Advocate** — scope creep, over-engineering, YAGNI
5. **Performance & Scalability** — throughput, latency, resource utilization
6. **Cost Efficiency** — infrastructure costs, resource optimization
7. **Accessibility & Inclusivity** — universal design, inclusive language
8. **Ecosystem & Interoperability** — platform consistency, API contracts
9. **Testing & Maintainability** — test strategy, technical debt, coupling

Personas run in parallel (no anchoring bias), post findings as PR comments, and
classify issues as CRITICAL, CONCERN, or SUGGESTION. Reviews are advisory only
and do not block merge.

### Governance

- **Who can submit**: Any engineer, regardless of seniority.
- **Sponsor**: A senior engineer or architect (not the author) shepherds the
  decision.
- **Review period**: 10 business days from proposed status.
- **Required reviewers**: At least one engineer from each affected team.
- **Decision**: Sponsor accepts, accepts with conditions, or rejects.
- **Stale RFCs**: 30-day inactivity reminder, 60-day auto-withdrawal.
- **Tie-breaking**: Sponsor escalates to engineering leadership, who must decide
  within 5 business days.

### Engineering Principles

Ten engineering principles serve as the evaluation framework:

1. Simplicity Over Cleverness
2. Design for Failure
3. Optimize for the Reader
4. Secure by Default
5. Observability as a First-Class Concern
6. Incremental Delivery
7. Own Your Dependencies
8. Reduce Cognitive Load
9. Backward Compatibility by Default
10. Data-Informed Decisions

## Alternatives Considered

### No formal process (status quo)

Continue making decisions through ad hoc discussions. Rejected because this
fails to produce durable records, does not ensure cross-team review, and leads
to repeated debates about previously decided topics.

### Lightweight ADR (Architecture Decision Records)

ADRs are shorter and less structured than RFCs. Considered but rejected as
the primary mechanism because ADRs lack the formal review period and
multi-stakeholder input that cross-team decisions require. ADRs may be
appropriate for team-level decisions that do not meet the RFC threshold.

### RFC without AI review

A pure human-review RFC process. Considered viable but augmenting with AI
review provides consistent, thorough coverage across all perspectives at
near-zero marginal cost. AI review is positioned as advisory to avoid
over-reliance.

### Commercial decision-tracking tools

Tools like Notion, Confluence, or dedicated RFC platforms. Rejected because
keeping RFCs in the same Git repository as the code ensures they are versioned,
reviewable via PR workflow, and discoverable alongside the codebase. No
additional tool licenses or onboarding required.

## Migration and Rollout

### Phase 1: Foundation (immediate)

Deploy the RFC template, engineering principles, governance documentation, and
validation tooling. No process changes yet.

### Phase 2: Tooling (immediate)

Deploy the rfclint validation tool with pre-commit hook and CI integration.
Engineers can begin familiarizing themselves with the template.

### Phase 3: AI Review (immediate)

Deploy AI persona definitions and the review workflow. This RFC serves as the
first test.

### Phase 4: Adoption (ongoing)

Begin requiring RFCs for cross-team decisions. Iterate on persona prompts based
on review quality. Adjust review periods and governance rules based on
experience.

### Rollback

If the process proves too burdensome, it can be simplified by:

- Reducing the number of required sections
- Shortening the review period
- Disabling specific AI personas
- Reverting to a lighter-weight ADR format

All changes are reversible because the process is defined entirely in
version-controlled configuration.

## Risks and Mitigations

| Risk | Likelihood | Impact | Mitigation |
|------|-----------|--------|------------|
| Process perceived as bureaucratic | Medium | High | Start with clear guidance on when RFCs are and are not needed. Keep the bar for "not needed" low. |
| AI reviews generate false positives | Medium | Low | Reviews are advisory only. Iterate on prompts. Sponsors can dismiss irrelevant findings. |
| Low adoption due to writing burden | Medium | Medium | Provide a clear template with examples. Celebrate early RFCs. Offer writing support. |
| Stale RFCs accumulate | Low | Low | Automated 30-day reminder and 60-day auto-withdrawal. |
| AI review costs escalate | Low | Low | Monitor API usage. Personas only run on non-draft, proposed RFCs. |

## References

- [HashiCorp RFC Process](https://works.hashicorp.com/articles/writing-practices-and-culture)
- [Squarespace Engineering Blog on RFCs](https://engineering.squarespace.com/blog/2019/the-power-of-yes-if)
- [Oxide Computer RFD Process](https://oxide.computer/blog/rfd-1-requests-for-discussion)
