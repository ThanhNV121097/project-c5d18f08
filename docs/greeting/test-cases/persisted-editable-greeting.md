# Test Cases — Persisted editable greeting

Module: `greeting`
Function: Persisted editable greeting
Risk level: Medium. Story is small, but it writes shared persisted data through browser, API, and PostgreSQL, so save, validation, persistence, and contract failures need explicit checks.

## Page rendering and saved value

**Scenario**: First visit shows initial greeting when no row exists
**Given**: Greeting storage has no current greeting row.
**When**: Visitor opens the page.
**Then**: The visible greeting heading text is exactly `Hello, World!`.
Traces: SC-1 (GREETING-001 AC-1)
Check: render_url

**Scenario**: Existing stored greeting appears as heading
**Given**: Stored greeting is `Pipeline accepted`.
**When**: Visitor opens the page.
**Then**: The visible greeting heading text is exactly `Pipeline accepted`.
Traces: SC-2 (GREETING-001 AC-2)
Check: render_url

**Scenario**: Existing stored greeting prefills input
**Given**: Stored greeting is `Pipeline accepted`.
**When**: Visitor opens the page.
**Then**: Greeting text input value is exactly `Pipeline accepted`.
Traces: SC-3 (GREETING-001 AC-3)
Check: render_url

**Scenario**: Page has one visible greeting heading
**Given**: Visitor opens the page.
**When**: Page renders.
**Then**: Exactly one visible `h1` is present, and its text is current greeting.
Traces: SC-4 (GREETING-001 AC-4)
Check: render_url

**Scenario**: Input has accessible name from associated label
**Given**: Visitor opens the page.
**When**: Page renders.
**Then**: Text input has accessible name `Greeting` from associated label.
Traces: SC-5 (GREETING-001 AC-5)
Check: render_url

**Scenario**: Save submit button is present
**Given**: Visitor opens the page.
**When**: Page renders.
**Then**: One submit button named `Save` is present.
Traces: SC-6 (GREETING-001 AC-6)
Check: render_url

**Scenario**: Page has no navigation or extra section
**Given**: Visitor opens the page.
**When**: Page renders.
**Then**: No navigation landmark is present, and exactly one visible content section is present.
Traces: SC-7 (GREETING-001 AC-7)
Check: render_url

**Scenario**: Greeting section has no animation or transition
**Given**: Visitor opens the page.
**When**: Page renders.
**Then**: Greeting section and its visible descendants have `animation-name: none` and `transition-duration: 0s`.
Traces: SC-8 (GREETING-001 AC-8)
Check: measure_styles

## Saving and persistence

**Scenario**: Save button updates heading with valid greeting
**Given**: Page shows stored greeting `Hello, World!`.
**When**: Visitor enters `Hello, Pipeline!` and selects `Save`.
**Then**: Heading text becomes exactly `Hello, Pipeline!`.
Traces: SC-9 (GREETING-002 AC-1)
Check: interact_page

**Scenario**: Enter key updates heading with valid greeting
**Given**: Page shows stored greeting `Hello, World!`.
**When**: Visitor enters `Hello, Enter!` and presses Enter in the input.
**Then**: Heading text becomes exactly `Hello, Enter!`.
Traces: SC-10 (GREETING-002 AC-2)
Check: interact_page

**Scenario**: Reload keeps saved greeting in heading
**Given**: Visitor saved `Hello, Pipeline!` successfully.
**When**: Visitor reloads the page.
**Then**: Heading text is exactly `Hello, Pipeline!`.
Traces: SC-11 (GREETING-002 AC-3)
Check: interact_page

**Scenario**: Reload keeps saved greeting in input
**Given**: Visitor saved `Hello, Pipeline!` successfully.
**When**: Visitor reloads the page.
**Then**: Greeting text input value is exactly `Hello, Pipeline!`.
Traces: SC-12 (GREETING-002 AC-4)
Check: interact_page

**Scenario**: Save trims surrounding whitespace
**Given**: Stored greeting is `Hello, World!`.
**When**: Visitor enters `   Trimmed greeting   ` and selects `Save`.
**Then**: Heading text becomes exactly `Trimmed greeting`.
Traces: SC-13 (GREETING-002 AC-5)
Check: interact_page

**Scenario**: Empty submit keeps existing greeting unchanged
**Given**: Stored greeting is `Hello, World!`.
**When**: Visitor clears input and submits empty value.
**Then**: Heading text remains exactly `Hello, World!`.
Traces: SC-14 (GREETING-002 AC-6)
Check: interact_page

**Scenario**: Empty submit returns focus to input
**Given**: Stored greeting is `Hello, World!`.
**When**: Visitor clears input and submits empty value.
**Then**: Browser focus is on greeting input.
Traces: SC-15 (GREETING-002 AC-7)
Check: interact_page

**Scenario**: Whitespace-only submit keeps existing greeting and no visible error
**Given**: Stored greeting is `Hello, World!`.
**When**: Visitor enters `   ` and selects `Save`.
**Then**: Heading text remains exactly `Hello, World!`, focus is on greeting input, and no visible error message appears.
Traces: SC-14 (GREETING-002 AC-6), SC-15 (GREETING-002 AC-7)
Check: interact_page

**Scenario**: Later completed save wins
**Given**: Stored greeting is `First`.
**When**: Visitor A saves `Second`, then Visitor B saves `Third`, and page reloads.
**Then**: Heading text is exactly `Third`.
Traces: SC-16 (GREETING-002 AC-8)
Check: interact_page

## Minimal visual design

**Scenario**: Body background is white
**Given**: Visitor opens the page.
**When**: Page renders.
**Then**: Body computed background color is `#FFFFFF`.
Traces: SC-17 (GREETING-003 AC-1)
Check: measure_styles

**Scenario**: Heading text is black
**Given**: Visitor opens the page.
**When**: Page renders.
**Then**: Greeting heading computed text color is `#000000`.
Traces: SC-18 (GREETING-003 AC-2)
Check: measure_styles

**Scenario**: Heading is centered above form
**Given**: Visitor opens the page.
**When**: Page renders.
**Then**: Greeting heading center x matches form center x within 1px, and heading bottom is above form top.
Traces: SC-19 (GREETING-003 AC-3)
Check: measure_styles

**Scenario**: Input height is 48px
**Given**: Visitor opens the page.
**When**: Page renders.
**Then**: Greeting input computed height is `48px`.
Traces: SC-20 (GREETING-003 AC-4)
Check: measure_styles

**Scenario**: Input border is black, 1px, square-cornered
**Given**: Visitor opens the page.
**When**: Page renders.
**Then**: Greeting input computed border color is `#000000`, border width is `1px`, and border radius is `0px`.
Traces: SC-21 (GREETING-003 AC-5)
Check: measure_styles

**Scenario**: Save button background is blue
**Given**: Visitor opens the page.
**When**: Page renders.
**Then**: `Save` button computed background color is `#2563EB`.
Traces: SC-22 (GREETING-003 AC-6)
Check: measure_styles

**Scenario**: Save button text style matches design
**Given**: Visitor opens the page.
**When**: Page renders.
**Then**: `Save` button computed text color is `#FFFFFF`, font weight is bold, and font size is `18px`.
Traces: SC-23 (GREETING-003 AC-7)
Check: measure_styles

**Scenario**: Input keyboard focus outline matches design
**Given**: Visitor opens the page.
**When**: Visitor tabs to the input.
**Then**: Input computed focus outline color is `#2563EB`, outline width is `3px`, and outline offset is `3px`.
Traces: SC-24 (GREETING-003 AC-8)
Check: interact_page

**Scenario**: Save button keyboard focus outline matches design
**Given**: Visitor opens the page.
**When**: Visitor tabs to the `Save` button.
**Then**: Button computed focus outline color is `#2563EB`, outline width is `3px`, and outline offset is `3px`.
Traces: SC-25 (GREETING-003 AC-9)
Check: interact_page

**Scenario**: Desktop-width controls sit side by side with 12px gap
**Given**: Viewport width is 560px.
**When**: Page renders.
**Then**: Input and button top edges align, input right edge is 12px from button left edge, and both controls are in one row.
Traces: SC-26 (GREETING-003 AC-10)
Check: measure_styles

**Scenario**: Compact-width controls stack full width with 48px height
**Given**: Viewport width is 390px.
**When**: Page renders.
**Then**: Input and button are stacked vertically, each spans available section width, and each computed height is `48px`.
Traces: SC-27 (GREETING-003 AC-11)
Check: measure_styles

## API contract

**Scenario**: GET greeting returns current greeting success shape
**Given**: Stored greeting is `Hello, World!`.
**When**: Client requests `GET /v1/greeting`.
**Then**: Response status is `200 OK`, content type is `application/json`, and body is exactly `{"greeting":"Hello, World!"}`.
Traces: contract (GET /v1/greeting)
Check: fetch_url

**Scenario**: GET greeting returns stored changed greeting
**Given**: Stored greeting is `Pipeline accepted`.
**When**: Client requests `GET /v1/greeting`.
**Then**: Response status is `200 OK` and JSON body has `greeting` equal to `Pipeline accepted`.
Traces: contract (GET /v1/greeting)
Check: fetch_url

**Scenario**: GET greeting seeds initial value when row is absent
**Given**: Greeting storage has no current greeting row.
**When**: Client requests `GET /v1/greeting`.
**Then**: Response status is `200 OK` and JSON body has `greeting` equal to `Hello, World!`.
Traces: contract (GET /v1/greeting)
Check: fetch_url

**Scenario**: GET greeting database failure returns internal error envelope
**Given**: Database query for current greeting fails after connection is established.
**When**: Client requests `GET /v1/greeting`.
**Then**: Response status is `500 Internal Server Error`, content type is `application/json`, and body has `error.code` equal to `INTERNAL_ERROR` with no database detail.
Traces: contract (GET /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting saves valid trimmed text
**Given**: Stored greeting is `Hello, World!`.
**When**: Client requests `PUT /v1/greeting` with JSON body `{"greeting":"  Hello, API!  "}`.
**Then**: Response status is `200 OK` and body is exactly `{"greeting":"Hello, API!"}`.
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting rejects missing greeting field
**Given**: Stored greeting is `Hello, World!`.
**When**: Client requests `PUT /v1/greeting` with JSON body `{}`.
**Then**: Response status is `400 Bad Request`, body has `error.code` equal to `VALIDATION_ERROR`, and stored greeting remains `Hello, World!`.
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting rejects non-string greeting
**Given**: Stored greeting is `Hello, World!`.
**When**: Client requests `PUT /v1/greeting` with JSON body `{"greeting":123}`.
**Then**: Response status is `400 Bad Request`, body has `error.code` equal to `VALIDATION_ERROR`, and stored greeting remains `Hello, World!`.
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting rejects empty greeting
**Given**: Stored greeting is `Hello, World!`.
**When**: Client requests `PUT /v1/greeting` with JSON body `{"greeting":""}`.
**Then**: Response status is `400 Bad Request`, body has `error.code` equal to `VALIDATION_ERROR`, and stored greeting remains `Hello, World!`.
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting rejects whitespace-only greeting
**Given**: Stored greeting is `Hello, World!`.
**When**: Client requests `PUT /v1/greeting` with JSON body `{"greeting":"   "}`.
**Then**: Response status is `400 Bad Request`, body has `error.code` equal to `VALIDATION_ERROR`, and stored greeting remains `Hello, World!`.
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting rejects unknown request field
**Given**: Stored greeting is `Hello, World!`.
**When**: Client requests `PUT /v1/greeting` with JSON body `{"greeting":"Hello, API!","extra":"ignored?"}`.
**Then**: Response status is `400 Bad Request`, body has `error.code` equal to `VALIDATION_ERROR`, and stored greeting remains `Hello, World!`.
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting database failure returns internal error envelope
**Given**: Database update for current greeting fails after connection is established.
**When**: Client requests `PUT /v1/greeting` with JSON body `{"greeting":"Hello, API!"}`.
**Then**: Response status is `500 Internal Server Error`, content type is `application/json`, and body has `error.code` equal to `INTERNAL_ERROR` with no database detail.
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: Unknown API path returns not found envelope
**Given**: Backend is running.
**When**: Client requests `GET /v1/unknown`.
**Then**: Response status is `404 Not Found`, content type is `application/json`, and body has `error.code` equal to `NOT_FOUND`.
Traces: contract (unknown paths)
Check: fetch_url

**Scenario**: Unsupported API method returns method not allowed envelope
**Given**: Backend is running.
**When**: Client requests `POST /v1/greeting`.
**Then**: Response status is `405 Method Not Allowed`, content type is `application/json`, and body has `error.code` equal to `METHOD_NOT_ALLOWED`.
Traces: contract (unsupported methods)
Check: fetch_url

**Scenario**: Health check passes after migration and database connectivity
**Given**: Backend migration completed and database connectivity succeeds.
**When**: Client requests `GET /healthz`.
**Then**: Response status is `200 OK` and body is exactly `{"status":"ok"}`.
Traces: contract (GET /healthz)
Check: fetch_url

**Scenario**: Health check reports unavailable when database is unreachable
**Given**: Backend cannot connect to database.
**When**: Client requests `GET /healthz`.
**Then**: Response status is `503 Service Unavailable`.
Traces: contract (GET /healthz)
Check: fetch_url
