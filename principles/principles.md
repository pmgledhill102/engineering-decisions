# Engineering Principles

These 10 principles guide technical decision-making. They are the foundation
against which RFCs are evaluated and trade-offs are reasoned about.

## 1. Simplicity Over Cleverness

Choose the simplest solution that meets the requirements. Clever code is hard to
debug, hard to maintain, and hard to hand off. When in doubt, prefer the boring
approach.

## 2. Design for Failure

Assume every component will fail and design accordingly. Systems should degrade
gracefully, blast radius should be bounded, and recovery should be automated
where possible. Availability is a feature.

## 3. Optimize for the Reader

Code is read far more often than it is written. Prioritize comprehension over
writing convenience. Names should be descriptive, flows should be
straightforward, and surprising behavior should be documented.

## 4. Secure by Default

Security belongs in the default configuration, not bolted on after the fact.
Secrets must never appear in code or logs. Authentication and authorization
models should be explicit and reviewable.

## 5. Observability as a First-Class Concern

Every system ships with metrics, structured logging, and distributed tracing.
If you cannot observe it, you cannot operate it. Instrument before you ship,
not after the first incident.

## 6. Incremental Delivery

Deliver in independently valuable milestones rather than big-bang releases.
Each increment should be deployable, reversible, and measurable. Smaller
batches reduce risk and accelerate feedback.

## 7. Own Your Dependencies

Evaluate dependencies before adopting them. Understand their maintenance status,
security posture, license, and transitive dependency tree. A dependency you do
not understand is a liability, not a shortcut.

## 8. Reduce Cognitive Load

Prefer explicit over implicit. Minimize the number of concepts a reader must
hold in their head to understand a system. Consistent patterns across services
reduce the cost of context-switching between teams.

## 9. Backward Compatibility by Default

Breaking changes require explicit justification and a documented migration path.
Consumers should have time and tooling to migrate. When in doubt, keep the old
interface working.

## 10. Data-Informed Decisions

Support technical proposals with benchmarks, metrics, or evidence. Opinions are
welcome; data is better. When data is unavailable, be explicit about assumptions
and plan to validate them.
