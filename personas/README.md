# AI Review Personas

AI personas provide automated, multi-perspective review of decision proposals.
Each persona evaluates an RFC through a specialized lens, surfacing
considerations that human reviewers might miss — or that would require assembling
a large review committee to cover.

## How Personas Work

When an RFC pull request is marked as ready for review (status: `proposed`), the
CI workflow runs all personas in parallel. Each persona:

1. Reads its prompt template from this directory.
2. Reads the RFC content from the PR.
3. Evaluates the RFC against its domain-specific criteria and the engineering
   principles.
4. Posts a PR comment with findings classified by severity.

## Severity Levels

| Level | Meaning |
|-------|---------|
| **CRITICAL** | Must be addressed before the RFC can be accepted |
| **CONCERN** | Should be addressed; may be accepted with documented rationale |
| **SUGGESTION** | Nice to have; improves quality but not blocking |

## Advisory Only

Persona reviews are **advisory only** and do not block merge. However, RFC
sponsors must explicitly address any CRITICAL findings before accepting a
proposal (see [decision-framework.md](../governance/decision-framework.md)).

## Why Parallel Review Matters

Personas run independently with no visibility into each other's output. This
eliminates anchoring bias — each perspective is formed from first principles
rather than being influenced by another reviewer's framing.

## Personas

| Persona | Lens |
|---------|------|
| [Security & Privacy](security-and-privacy.md) | Threat models, auth, secrets, compliance |
| [Reliability & Operations](reliability-and-operations.md) | Failure modes, observability, blast radius |
| [Developer Experience](developer-experience.md) | API design, cognitive load, migration burden |
| [Simplicity Advocate](simplicity-advocate.md) | Scope creep, over-engineering, YAGNI |
| [Performance & Scalability](performance-and-scalability.md) | Throughput, latency, resource utilization |
| [Cost Efficiency](cost-efficiency.md) | Infrastructure costs, resource optimization |
| [Accessibility & Inclusivity](accessibility-and-inclusivity.md) | Universal design, inclusive language |
| [Ecosystem & Interoperability](ecosystem-and-interoperability.md) | Platform consistency, API contracts |
| [Testing & Maintainability](testing-and-maintainability.md) | Test strategy, technical debt, coupling |

## Versioning

Persona definitions are versioned alongside the rest of the repository. Changes
to persona prompts should be tested against existing RFCs to verify review
quality.
