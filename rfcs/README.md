# Decision Proposals (RFCs)

RFCs (Requests for Comments) are the mechanism by which cross-team engineering
decisions are proposed, discussed, and recorded.

## Index

| RFC | Title | Status | Authors | Date |
|-----|-------|--------|---------|------|
| [0001](0001-ai-augmented-decision-process.md) | AI-Augmented Decision Process for Engineering | proposed | Paul Gledhill (@pmgledhill102) | 2026-02-14 |
| [0002](0002-graphql-migration.md) | Migrate API Gateway from REST to GraphQL | proposed | Paul Gledhill (@pmgledhill102) | 2026-03-18 |

## Submitting an RFC

1. Copy [`_template.md`](_template.md) to `NNNN-short-slug.md` where `NNNN` is
   the next available number (zero-padded to 4 digits).
2. Fill in all required sections.
3. Open a **draft** pull request. CI will validate the format.
4. When ready for review, mark the PR as **ready** and set the frontmatter
   status to `proposed`. AI persona reviews will run automatically.
5. A sponsor shepherds the RFC through the review period.

See [governance/process.md](../governance/process.md) for full lifecycle details.
