# Story — Persisted editable greeting

Module: `greeting`
Plan item: Persisted editable greeting

## User story

As a Visitor, I want to view and save one shared greeting, so that the page always shows the latest greeting persisted in PostgreSQL through the Go API and Next.js UI.

## In scope

- Render one centered greeting page section from the approved design.
- Read the current shared greeting through the Go API backed by PostgreSQL.
- Seed or return `Hello, World!` when no saved greeting exists.
- Prefill the text input with the shown greeting.
- Save a valid edited greeting with the `Save` button or Enter key.
- Trim leading and trailing whitespace before saving.
- Reject empty or whitespace-only submissions without changing stored greeting.
- Keep latest valid saved greeting after page reload.
- Apply approved minimal visual styling, responsive layout, and keyboard focus behaviour.

## Out of scope

- Sign-in, permissions, and user-specific greetings; SRS states every Visitor shares same greeting and no sign-in exists.
- Navigation, extra sections, secondary copy, loading screens, empty screens, visible error messages, hover states, active states, disabled states, transitions, and animations; approved design does not include them.
- External services; stakeholder specified none.
- Multiple greetings, greeting history, audit trail, delete/reset action, or per-visitor storage; plan item owns one current shared greeting only.
- Admin moderation, profanity filtering, length limit beyond layout safety, or rich text formatting; not specified by SRS.
- Changing global scaffold, design tokens, architecture shape, or provider setup; this story consumes existing scaffold and contracts.

## UI scope

- Screen: Greeting page, default state only.
- Section: `<main><section aria-labelledby="greeting-heading">` from approved design.
- Visible elements: one `h1` greeting heading, visually hidden `Greeting` label associated with one text input, and one `Save` submit button.
- Layout: white page, black centered heading above centered form; form controls are side by side above 520px and stacked at 520px and below.
- States: default plus native keyboard focus-visible outline for input and button. No loading, empty, visible error, hover, active, or disabled visual state is introduced.

## Acceptance criteria

### Show saved greeting

- SC-1 [GREETING-001 AC-1]: Given no greeting has previously been saved, when Visitor opens page, heading text is `Hello, World!`.
- SC-2 [GREETING-001 AC-2]: Given stored greeting is `Pipeline accepted`, when Visitor opens page, heading text is `Pipeline accepted`.
- SC-3 [GREETING-001 AC-3]: Given stored greeting is `Pipeline accepted`, when Visitor opens page, text input value is `Pipeline accepted`.
- SC-4 [GREETING-001 AC-4]: Given Visitor opens page, when page renders, exactly one visible `h1` greeting heading is present.
- SC-5 [GREETING-001 AC-5]: Given Visitor opens page, when page renders, input has accessible name `Greeting` from associated label.
- SC-6 [GREETING-001 AC-6]: Given Visitor opens page, when page renders, `Save` submit button is present.
- SC-7 [GREETING-001 AC-7]: Given Visitor opens page, when page renders, no navigation landmark or extra content section is present.
- SC-8 [GREETING-001 AC-8]: Given Visitor opens page, when page renders, no CSS animation or transition changes greeting section.

### Save changed greeting

- SC-9 [GREETING-002 AC-1]: Given page shows stored greeting `Hello, World!`, when Visitor enters `Hello, Pipeline!` and selects `Save`, heading text becomes `Hello, Pipeline!`.
- SC-10 [GREETING-002 AC-2]: Given page shows stored greeting `Hello, World!`, when Visitor enters `Hello, Enter!` and presses Enter in input, heading text becomes `Hello, Enter!`.
- SC-11 [GREETING-002 AC-3]: Given Visitor saved `Hello, Pipeline!` successfully, when Visitor reloads page, heading text is `Hello, Pipeline!`.
- SC-12 [GREETING-002 AC-4]: Given Visitor saved `Hello, Pipeline!` successfully, when Visitor reloads page, text input value is `Hello, Pipeline!`.
- SC-13 [GREETING-002 AC-5]: Given stored greeting is `Hello, World!`, when Visitor enters `   Trimmed greeting   ` and selects `Save`, heading text becomes `Trimmed greeting`.
- SC-14 [GREETING-002 AC-6]: Given stored greeting is `Hello, World!`, when Visitor submits empty value, heading text remains `Hello, World!`.
- SC-15 [GREETING-002 AC-7]: Given stored greeting is `Hello, World!`, when Visitor submits empty value, focus is on greeting input.
- SC-16 [GREETING-002 AC-8]: Given stored greeting is `First`, when Visitor A saves `Second` and then Visitor B saves `Third`, reload shows heading text `Third`.

### Match approved minimal design

- SC-17 [GREETING-003 AC-1]: Given Visitor opens page, when page renders, body background color is `#FFFFFF`.
- SC-18 [GREETING-003 AC-2]: Given Visitor opens page, when page renders, greeting heading text color is `#000000`.
- SC-19 [GREETING-003 AC-3]: Given Visitor opens page, when page renders, greeting heading is centered above form.
- SC-20 [GREETING-003 AC-4]: Given Visitor opens page, when page renders, input height is 48px.
- SC-21 [GREETING-003 AC-5]: Given Visitor opens page, when page renders, input border is black, 1px, and square-cornered.
- SC-22 [GREETING-003 AC-6]: Given Visitor opens page, when page renders, `Save` button background is `#2563EB`.
- SC-23 [GREETING-003 AC-7]: Given Visitor opens page, when page renders, `Save` button text is white, bold, and 18px.
- SC-24 [GREETING-003 AC-8]: Given Visitor tabs to input, when input receives keyboard focus, input shows blue 3px focus outline with 3px offset.
- SC-25 [GREETING-003 AC-9]: Given Visitor tabs to `Save` button, when button receives keyboard focus, button shows blue 3px focus outline with 3px offset.
- SC-26 [GREETING-003 AC-10]: Given viewport width is 560px, when page renders, input and button are side by side with 12px gap.
- SC-27 [GREETING-003 AC-11]: Given viewport width is 390px, when page renders, input and button are stacked, full width, and both 48px tall.

## Dependencies

- Requires existing Next.js frontend scaffold and frozen global design tokens.
- Requires Go API serving greeting read and update endpoints under `/v1/...`.
- Requires PostgreSQL persistence with one shared greeting row and initial value `Hello, World!`.
- Requires backend migrations to run before API health passes.
- Requires frontend configured with `NEXT_PUBLIC_API_URL` or proxy path per architecture.
- No external accounts, secrets beyond existing environment values, or stakeholder answers required.

## Notes for implementation and test

- Empty or whitespace-only submit keeps existing greeting and returns focus to input without visible error message.
- Long greeting must wrap inside centered section with no horizontal page scroll from 320px upward.
- Later completed valid save wins when two Visitors save in sequence.
- API or database failure behaviour is not visible in approved design; backend should use architecture error envelope, while UI must not invent visible error state.
