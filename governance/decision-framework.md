# Decision Framework

This document describes how RFC decisions are made, how disagreements are
resolved, and what "done" looks like.

## Decision Authority

The RFC sponsor makes the final decision. The sponsor is a senior engineer or
architect who is **not** the author of the RFC.

## Outcomes

### Accepted

Consensus exists that the proposal should proceed. The sponsor documents any
conditions and the RFC is merged with `accepted` status.

### Accepted with Conditions

The proposal is approved contingent on specific, documented conditions being met
during implementation. This follows the Squarespace pattern of "Yes, if..." —
it avoids blocking good proposals on minor issues while ensuring those issues
are tracked.

Conditions are recorded in the PR merge comment and in the RFC frontmatter.

### Rejected

The sponsor determines the proposal should not proceed. The rejection rationale
is documented in full. The RFC is still merged — rejected RFCs are valuable
prior art that prevents the same ground from being relitigated.

## AI Persona Reviews

AI persona reviews are **advisory only**. They do not block merge. Their purpose
is to surface considerations that human reviewers might miss.

However, sponsors must explicitly address any **CRITICAL** findings from persona
reviews before accepting an RFC. "Address" means one of:

- Modifying the proposal to resolve the concern
- Documenting why the concern does not apply
- Acknowledging the risk and documenting the accepted trade-off

## Resolving Disagreements

### Consensus-Seeking

The default process is consensus-seeking. The sponsor facilitates discussion and
works toward a decision that all affected parties can support, even if it is not
everyone's first choice.

### Tie-Breaking

When consensus cannot be reached within the review period:

1. The sponsor escalates to engineering leadership.
2. Engineering leadership must decide within **5 business days**.
3. The decision and rationale are documented in the RFC.

## What "Done" Looks Like

An RFC decision is final when:

1. The review period has elapsed (or been explicitly shortened by the sponsor).
2. All affected teams have had the opportunity to review.
3. Any CRITICAL AI persona findings have been addressed.
4. The sponsor has recorded the decision and rationale.
5. The PR is merged to main.

## Revisiting Decisions

Accepted RFCs can be revisited by submitting a new RFC that supersedes the
original. The new RFC must explain what has changed since the original decision
(new information, changed requirements, lessons learned from implementation).
