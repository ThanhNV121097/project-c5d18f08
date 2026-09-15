# SRS — Greeting

Module: `greeting`
Design: [View the approved design](http://localhost:8080/design/c5d18f08-7743-4ce3-8eeb-86ec592eb6be)
Design system: `design/design-system.md`

> One file per module, at `docs/{module}/SRS.md`. It covers only the functions
> that belong to this module. Never write `docs/SRS.md`.

## 1. Purpose

The greeting module lets any visitor view and update one persisted greeting for "Hello World Acceptance 5". It proves the full pipeline end to end: PostgreSQL keeps the greeting, the Go API serves and updates it, and the Next.js page renders and edits it. Without this module, the product becomes a static page and cannot validate persistence across reloads.

## 2. Actors

| Actor | Who they are | What they may do in this module |
|---|---|---|
| Visitor | Anyone who opens the page; no sign-in exists | View the current greeting, edit the greeting text, and save it |

## 3. Scope

**In scope** — the functions specified below, by their plan titles:

- Persisted editable greeting

**Out of scope** — name what a reader would reasonably expect here and say
where it lives instead. This section prevents the same argument twice.

- Sign-in and user-specific greetings — deliberately not built; the stakeholder specified no sign-in.
- Navigation and extra sections — deliberately not built; the approved design has one centered section only.
- External services — deliberately not built; the stakeholder specified none.
- Loading, error, hover, active, and disabled visual states — deliberately not built; the approved design does not draw them.

## 4. Functional requirements

### 4.1 Persisted editable greeting

**Requirement GREETING-001 — Show saved greeting**

*As a* Visitor, *I want to* see the saved greeting when I open the page, *so that* the page reflects persisted data instead of fixed copy.

Behaviour:

1. When the Visitor opens the page for the first time and no greeting has been changed, the page shows `Hello, World!` as the large greeting heading.
2. When the stored greeting has another value, the page shows that stored value as the large greeting heading.
3. The greeting text field is prefilled with the same value shown in the heading.
4. The page displays one centered section only: visible greeting heading, visually hidden `Greeting` label associated to the input, text input, and `Save` button.
5. The page displays no navigation, no extra section, no secondary text, and no animation.

**Acceptance criteria** — each is proved by at least one test case in
`docs/greeting/test-cases/persisted-editable-greeting.md`, through the story plan that cites it
(`SC-1 [GREETING-001 AC-1]`). Given/When/Then, no compound
conditions: one behaviour per criterion.

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | No greeting has previously been saved | Visitor opens the page | Heading text is `Hello, World!` |
| AC-2 | Stored greeting is `Pipeline accepted` | Visitor opens the page | Heading text is `Pipeline accepted` |
| AC-3 | Stored greeting is `Pipeline accepted` | Visitor opens the page | Text input value is `Pipeline accepted` |
| AC-4 | Visitor opens the page | Page renders | Exactly one visible `h1` greeting heading is present |
| AC-5 | Visitor opens the page | Page renders | Input has accessible name `Greeting` from an associated label |
| AC-6 | Visitor opens the page | Page renders | `Save` submit button is present |
| AC-7 | Visitor opens the page | Page renders | No navigation landmark or extra content section is present |
| AC-8 | Visitor opens the page | Page renders | No CSS animation or transition changes the greeting section |

**Requirement GREETING-002 — Save changed greeting**

*As a* Visitor, *I want to* save a new greeting, *so that* future visits and reloads show my latest text.

Behaviour:

1. The Visitor edits the greeting in the text field and submits the form with the `Save` button or by pressing Enter in the input.
2. When the text field contains at least one non-whitespace character after trimming, the saved greeting becomes the trimmed submitted text.
3. After a successful save, the heading updates to the saved greeting and the text field value matches it.
4. After the page reloads, the heading and text field still show the latest saved greeting.
5. If the Visitor submits an empty or whitespace-only value, the existing saved greeting is unchanged and focus returns to the input; no visible error message is shown in the approved design.
6. If two Visitors save different valid greetings, the later completed save is the greeting shown after reload.

**Acceptance criteria** — each is proved by at least one test case in
`docs/greeting/test-cases/persisted-editable-greeting.md`, through the story plan that cites it
(`SC-1 [GREETING-002 AC-1]`). Given/When/Then, no compound
conditions: one behaviour per criterion.

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | Page shows stored greeting `Hello, World!` | Visitor enters `Hello, Pipeline!` and selects `Save` | Heading text becomes `Hello, Pipeline!` |
| AC-2 | Page shows stored greeting `Hello, World!` | Visitor enters `Hello, Enter!` and presses Enter in the input | Heading text becomes `Hello, Enter!` |
| AC-3 | Visitor saved `Hello, Pipeline!` successfully | Visitor reloads the page | Heading text is `Hello, Pipeline!` |
| AC-4 | Visitor saved `Hello, Pipeline!` successfully | Visitor reloads the page | Text input value is `Hello, Pipeline!` |
| AC-5 | Stored greeting is `Hello, World!` | Visitor enters `   Trimmed greeting   ` and selects `Save` | Heading text becomes `Trimmed greeting` |
| AC-6 | Stored greeting is `Hello, World!` | Visitor submits an empty value | Heading text remains `Hello, World!` |
| AC-7 | Stored greeting is `Hello, World!` | Visitor submits an empty value | Focus is on the greeting input |
| AC-8 | Stored greeting is `First` | Visitor A saves `Second`, then Visitor B saves `Third` | Reload shows heading text `Third` |

**Requirement GREETING-003 — Match approved minimal design**

*As a* Visitor, *I want to* use the greeting form in the approved minimal layout, *so that* the page stays readable and usable at desktop and compact widths.

Behaviour:

1. The page background is white `#FFFFFF`; text is black `#000000`.
2. The visible greeting appears as a bold, large heading centered above the form.
3. The form is centered in the same section as the heading.
4. The input uses white background, black text, black 1px border, square corners, 48px height, and 18px text.
5. The `Save` button uses blue `#2563EB` background and border, white bold text, square corners, 48px height, and 18px text.
6. Keyboard focus on the input or button shows a blue `#2563EB` 3px outline with 3px offset.
7. At widths above 520px, the input and button sit in one row with a 12px gap.
8. At widths 520px and below, the input and button stack vertically, each spans the available section width, and each keeps 48px height.

**Acceptance criteria** — each is proved by at least one test case in
`docs/greeting/test-cases/persisted-editable-greeting.md`, through the story plan that cites it
(`SC-1 [GREETING-003 AC-1]`). Given/When/Then, no compound
conditions: one behaviour per criterion.

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | Visitor opens the page | Page renders | Body background color is `#FFFFFF` |
| AC-2 | Visitor opens the page | Page renders | Greeting heading text color is `#000000` |
| AC-3 | Visitor opens the page | Page renders | Greeting heading is centered above the form |
| AC-4 | Visitor opens the page | Page renders | Input height is 48px |
| AC-5 | Visitor opens the page | Page renders | Input border is black, 1px, and square-cornered |
| AC-6 | Visitor opens the page | Page renders | `Save` button background is `#2563EB` |
| AC-7 | Visitor opens the page | Page renders | `Save` button text is white, bold, and 18px |
| AC-8 | Visitor tabs to the input | Input receives keyboard focus | Input shows blue 3px focus outline with 3px offset |
| AC-9 | Visitor tabs to the `Save` button | Button receives keyboard focus | Button shows blue 3px focus outline with 3px offset |
| AC-10 | Viewport width is 560px | Page renders | Input and button are side by side with 12px gap |
| AC-11 | Viewport width is 390px | Page renders | Input and button are stacked, full width, and both 48px tall |

**Failure, boundary and permission behaviour** — the part most often skipped
and most often the source of bugs. Every case this function actually has needs a
defined outcome; "should not happen" is not an outcome.

| Case | Condition | Expected behaviour |
|---|---|---|
| Invalid input | Submitted greeting is empty or whitespace-only | Existing saved greeting remains unchanged; focus returns to input; no visible error message is shown because the approved design has no error state |
| Boundary | Submitted greeting contains leading or trailing whitespace around non-whitespace text | Saved greeting is trimmed before display and persistence |
| Boundary | Submitted greeting is long enough to exceed the heading width | Heading wraps within the centered section; no horizontal page scroll appears |
| Not found | Persisted greeting row is absent on first use | Visitor sees and can edit initial greeting `Hello, World!` |
| Not permitted | Any Visitor opens the page | Not applicable: no sign-in or roles exist; every Visitor has the same view and save permission |
| Conflict | Two Visitors save valid greetings in sequence | Later completed save wins and is shown after reload |
| Upstream failure | API or database cannot complete a read or write | No error or empty state is part of the approved design; the API contract's error envelope is specified in the service contract |

**Data touched** — the fields this function reads and writes, in product terms.
The physical schema is TL's job in `docs/architecture/erd.md`; this is the list
that document has to satisfy.

| Field | Type | Required | Rule |
|---|---|---|---|
| Greeting text | text | yes | Initial value is `Hello, World!`; submitted value is trimmed; saved value must contain at least one non-whitespace character |
| Last saved value | text | yes | One current greeting exists for the product; reload returns the latest saved value |

## 5. Screens

The design is the source of truth for appearance; this section maps functions
onto it so nothing in the design is unaccounted for and nothing specified here
is missing from the design.

List only the states the approved design actually shows. A screen the design
draws once, with no variant for waiting, for no data, or for a failure, has
exactly **one** state and its name is `default`. That is not a placeholder and
not an invented state: it is what "this screen has one appearance" is called,
and it is the correct and complete answer for a static screen. Writing
`loading`, `empty` or `error` for a screen whose design has no such variant
invents work, and the reviewer will reject it.

| Screen | Section in the design | Functions it serves | States that must exist |
|---|---|---|---|
| Greeting page | `<main><section aria-labelledby="greeting-heading">` | GREETING-001, GREETING-002, GREETING-003 | default |

## 6. Non-functional requirements

Only what is real for this module. Delete rows that do not apply rather than
inventing a number nobody will check.

| Area | Requirement |
|---|---|
| Accessibility | Input and button are keyboard reachable; input has accessible name `Greeting`; visible focus outline exists for input and button; text and controls have contrast ratio ≥ 4.5:1 where WCAG AA requires text contrast |
| Responsive | Page works at viewport widths from 320px upward with no horizontal page scroll; at 390px the input and button are stacked, full width, and both 48px tall |
| Privacy | Greeting text is public shared content; no personal data, account data, or external-service data is stored |

## 7. Dependencies and assumptions

- **Depends on:** Go API, for reading and updating the current greeting.
- **Depends on:** PostgreSQL, for persistence across reloads and process restarts.
- **Depends on:** Next.js frontend, for rendering the approved page and submitting edits.
- **Assumption:** One shared greeting exists for the whole product. If the stakeholder later wants per-user greetings, sign-in and user-specific storage become new scope.

Anything genuinely undecided goes here as an open question with a proposed
default — never leave a blank for someone else to discover mid-build:

| Open question | Proposed default | Who decides |
|---|---|---|
| — | No open questions | — |

## 8. Traceability

Every plan item in this module appears exactly once, and every requirement id
traces to a test case. A gap in this table is a gap in the build.

| Plan item | Requirement ids | Test cases |
|---|---|---|
| Persisted editable greeting | GREETING-001, GREETING-002, GREETING-003 | `test-cases/persisted-editable-greeting.md` |
