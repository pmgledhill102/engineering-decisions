# Engineering Decisions

A framework for making better engineering decisions through structured proposals
reviewed by AI personas — reducing human effort while improving decision quality.

## The Problem

Engineering decisions are expensive when they go wrong. Across most organizations,
cross-team technical decisions happen through ad hoc discussions in meetings, chat
threads, and email chains. This creates compounding problems:

- **Lost context** — decision rationale is not recorded, leading to repeated debates
- **Inconsistent review** — some proposals get deep scrutiny, others get rubber-stamped
- **No prior art** — rejected ideas are undocumented, so engineers waste time
  re-proposing previously dismissed approaches
- **Uneven participation** — decisions are made by whoever is in the room, not by
  systematically involving affected stakeholders

## The Solution

This framework combines a structured RFC (Request for Comments) process with
**AI-augmented multi-perspective review**. Nine specialized AI personas
systematically stress-test every proposal before humans spend time on it.

```text
Author writes RFC
        │
        ▼
   CI validates format (rfclint)
        │
        ▼
   9 AI personas review in parallel ◄── The force multiplier
        │
        ▼
   Humans review with AI insights already surfaced
        │
        ▼
   Decision recorded as structured, machine-readable artifact
```

### What makes this different

**AI personas replace the need to assemble large review committees.** Each persona
evaluates the proposal from a specialized lens — security, reliability, cost,
simplicity, and more — surfacing concerns that would otherwise require expertise
from across the organization. Reviews run in parallel with no anchoring bias.

**Decisions are stored as structured artifacts.** Every RFC has YAML frontmatter
(machine-readable metadata) and required sections that follow a consistent schema.
This format is optimized for both human readers and AI agents — any agent can
parse the decision history, understand the rationale, and build on prior art.

**Both accepted and rejected proposals are preserved.** The decision record is the
product, not just the outcome. Rejected RFCs are as valuable as accepted ones —
they document why an approach was considered and dismissed.

## How It Works

1. **Author** copies the [RFC template](rfcs/_template.md) and writes a proposal
2. **CI** validates the RFC format and structure automatically
   ([rfclint](tools/rfclint/))
3. **AI personas** review the proposal from 9 specialized perspectives — advisory
   only, findings classified as CRITICAL / CONCERN / SUGGESTION
4. **Human reviewers** discuss and refine, with AI insights already on the table
5. **Sponsor** makes the final accept/reject decision with full context

## AI Review Personas

| Persona | Lens |
|---------|------|
| [Security & Privacy](personas/security-and-privacy.md) | Threat models, auth, secrets, compliance |
| [Reliability & Operations](personas/reliability-and-operations.md) | Failure modes, observability, blast radius |
| [Developer Experience](personas/developer-experience.md) | API design, cognitive load, migration burden |
| [Simplicity Advocate](personas/simplicity-advocate.md) | Scope creep, over-engineering, YAGNI |
| [Performance & Scalability](personas/performance-and-scalability.md) | Throughput, latency, resource utilization |
| [Cost Efficiency](personas/cost-efficiency.md) | Infrastructure costs, resource optimization |
| [Accessibility & Inclusivity](personas/accessibility-and-inclusivity.md) | Universal design, inclusive language |
| [Ecosystem & Interoperability](personas/ecosystem-and-interoperability.md) | Platform consistency, API contracts |
| [Testing & Maintainability](personas/testing-and-maintainability.md) | Test strategy, technical debt, coupling |

Each persona is a structured prompt with defined evaluation criteria, severity
classification rules, and a focus on specific
[engineering principles](principles/principles.md). Persona definitions are
version-controlled and iterate based on review quality.

## Repository Structure

| Directory | Description |
|-----------|-------------|
| [`rfcs/`](rfcs/) | Decision proposals and index |
| [`principles/`](principles/) | Engineering principles — the evaluation framework |
| [`governance/`](governance/) | Process rules and decision framework |
| [`personas/`](personas/) | AI reviewer persona definitions |
| [`tools/`](tools/) | Validation tooling (rfclint) |

## Quick Start

- [Submit a decision proposal](rfcs/README.md)
- [Engineering Principles](principles/principles.md)
- [Decision Process](governance/process.md)
- [Decision Framework](governance/decision-framework.md)

## Design Principles

This framework is designed with three goals:

1. **Reduce human effort** — AI personas do the broad analysis; humans focus on
   judgment calls and domain-specific context that AI cannot provide
2. **Improve decision quality** — systematic multi-perspective review catches
   blind spots that ad hoc processes miss
3. **Create durable, agent-readable records** — every decision is stored in a
   structured format that both humans and AI agents can parse, query, and reason
   about
