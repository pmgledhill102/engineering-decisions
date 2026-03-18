# Decision Process

This document defines the lifecycle of an RFC from inception to resolution.

## When to Write an RFC

Write an RFC when a decision:

- Affects multiple teams or shared infrastructure
- Introduces a new system, service, or significant component
- Deprecates or removes existing functionality
- Changes an engineering principle or standard

An RFC is **not** needed for:

- Bug fixes
- Single-team feature work
- Minor tooling changes
- Dependency version bumps

## Who Can Submit

Any engineer, regardless of seniority. Good ideas are not correlated with title.

## RFC Numbering

RFCs are numbered sequentially with zero-padded 4-digit identifiers: `0001`,
`0002`, etc. The filename format is `NNNN-short-slug.md`.

## Lifecycle

```text
draft --> proposed --> accepted
                  \-> rejected
                  \-> withdrawn

accepted --> superseded
```

### draft

The author opens a **draft pull request**. CI validation runs to check format
and structure. AI persona reviews do **not** run on drafts.

Use this phase to iterate on the proposal with trusted reviewers before broader
exposure.

### proposed

The author marks the PR as **ready for review** and sets the frontmatter status
to `proposed`. This starts the formal review period:

- AI persona reviews run automatically.
- The 10 business day review window begins.
- The author identifies affected teams and requests reviews.
- A sponsor (senior engineer or architect, not the author) is assigned.

### accepted

The sponsor determines that consensus has been reached. The RFC is merged to
main. Implementation can begin.

Acceptance may be unconditional or **accepted with conditions** ("Yes, if...")
where specific conditions are documented and must be met during implementation.

### rejected

The sponsor determines the proposal should not proceed. The reason is documented
in the PR. The RFC is still merged to main — rejected RFCs serve as prior art
and prevent rehashing the same discussions.

### withdrawn

The author decides not to proceed. The RFC is merged with `withdrawn` status.

### superseded

A newer RFC replaces this one. The original RFC's frontmatter is updated with a
`superseded-by` field pointing to the replacement.

## Review Period

The formal review period is **10 business days** from the date the RFC enters
`proposed` status.

### Required Reviewers

At minimum, one engineer from each team affected by the proposal. The author
is responsible for identifying affected teams in the PR description.

### Stale RFCs

- **30 days** of inactivity on a `proposed` RFC triggers an automated reminder.
- **60 days** of inactivity results in automatic withdrawal.

## Roles

### Author

The engineer who writes and champions the RFC. Responsible for:

- Filling in all required sections of the template
- Identifying affected teams
- Responding to feedback
- Iterating on the proposal

### Sponsor

A senior engineer or architect who shepherds the RFC through the process. The
sponsor is **not** the author. Responsible for:

- Ensuring adequate review from affected teams
- Facilitating discussion and resolving disagreements
- Making the final accept/reject decision
- Documenting the rationale for the decision
