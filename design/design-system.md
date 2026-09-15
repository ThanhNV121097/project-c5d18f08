# Design System — Hello World Acceptance 5

> Source of truth: the approved `index.html`.
> Every value below is extracted from it. Changing a value here without changing the approved design is a defect.

Last updated: 2026-09-15

## 1. Foundations

### 1.1 Color

Semantic tokens. Name by job, never by hue.

| Token | Value | Used for |
|---|---|---|
| `--color-bg` | `#FFFFFF` | Page background, input background, button text |
| `--color-text` | `#000000` | Heading, input text, input border |
| `--color-primary-action` | `#2563EB` | Save button background, Save button border, focus outline |

#### Contrast audit

Every text-on-background pair actually used. Body text ≥ 4.5:1, large text (≥ 18.66px bold or ≥ 24px) ≥ 3:1, UI borders ≥ 3:1.

| Foreground | Background | Ratio | Passes |
|---|---|---|---|
| `--color-text` | `--color-bg` | `21:1` | AA |
| `--color-bg` | `--color-primary-action` | `5.17:1` | AA |
| `--color-primary-action` | `--color-bg` | `4.06:1` | UI border/focus pass |

### 1.2 Spacing

Base unit: `1px`. Every margin, padding, and gap in the approved design uses one of these values.

| Token | Value |
|---|---|
| `--space-screen` | `24px` |
| `--space-screen-compact` | `20px` |
| `--space-form-gap` | `12px` |
| `--space-heading-gap` | `32px` |
| `--space-input-inline` | `14px` |
| `--space-button-inline` | `22px` |
| `--space-focus-offset` | `3px` |
| `--space-sr-margin` | `-1px` |

### 1.3 Typography

Font families:

- Body: `Arial, Helvetica, sans-serif`, loaded from system fonts.
- Headings: `Arial, Helvetica, sans-serif`, inherited from body.

| Token | Size | Line height | Weight | Used for |
|---|---|---|---|---|
| `--text-control` | `18px` | Browser normal via `font: inherit` | `400` input, `700` button | Input and Save button |
| `--text-heading` | `clamp(40px, 8vw, 72px)` | `1.05` | `700` | Greeting h1 |

Heading levels are used in order: one `h1`, no skipped heading level.

| Token | Value | Used for |
|---|---|---|
| `--font-weight-body` | `400` | Input text and inherited page text |
| `--font-weight-action` | `700` | Save button |
| `--font-weight-heading` | `700` | Greeting h1 |

### 1.4 Radius, border, shadow, motion

| Token | Value | Used for |
|---|---|---|
| `--radius-control` | `0` | Input and Save button |
| `--border-width-control` | `1px` | Input and Save button border |
| `--focus-ring-width` | `3px` | Input and Save button focus outline |

No shadows are used. No transitions or animation are used.

### 1.5 Layout and breakpoints

| Name | Width | Container | Columns | Gutter |
|---|---|---|---|---|
| `base` | all widths | `min(100%, 560px)` | One centered section | `12px` form gap |
| `compact` | `max-width: 520px` | `min(100%, 560px)` | Form controls stack vertically | `12px` form gap |

Z-index scale:

| Layer | Value |
|---|---|
| Base | `0` |

## 2. Components

### 2.1 Greeting section

**Purpose** — Show current persisted greeting and provide one edit form. Do not use for navigation, lists, or extra content.

**Anatomy** — `[heading] [form]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Default | `--color-bg`, `--color-text`, `--text-heading`, `--space-heading-gap` | Only page section |

**Sizes**

| Size | Width | Padding | Text token |
|---|---|---|---|
| Default | `min(100%, 560px)` | Page padding `24px`; `20px` at compact width | `--text-heading` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | Centered heading above form | `--color-bg`, `--color-text` |

**Accessibility** — Section uses `aria-labelledby="greeting-heading"`; heading is visible and is page's first heading.

### 2.2 Greeting text field

**Purpose** — Let visitor edit greeting text before saving. Do not use for multi-line input.

**Anatomy** — `[visually hidden label] [text input]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Default | `--color-bg`, `--color-text`, `--border-width-control`, `--radius-control` | Greeting edit input |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Default | `48px` | `0 14px` | `--text-control` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | White background, black text, black 1px border | `--color-bg`, `--color-text` |
| Focus visible | Blue 3px outline, 3px offset | `--color-primary-action`, `--focus-ring-width`, `--space-focus-offset` |

**Accessibility** — Native text input with associated `label`; label is visually hidden. Keyboard focus must show `:focus-visible` outline. Minimum control height is `48px`. Required field rejects empty submission by focusing input.

### 2.3 Save button

**Purpose** — Submit greeting edit. Do not use for secondary actions.

**Anatomy** — `[label]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Primary action | `--color-primary-action`, `--color-bg`, `--font-weight-action`, `--radius-control` | Save greeting |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Default | `48px` | `0 22px` | `--text-control` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | Blue background, blue border, white bold text | `--color-primary-action`, `--color-bg`, `--font-weight-action` |
| Focus visible | Blue 3px outline, 3px offset | `--color-primary-action`, `--focus-ring-width`, `--space-focus-offset` |

**Accessibility** — Native `button type="submit"`; reachable by keyboard; Enter from input submits form; visible focus ring. Minimum control height is `48px`.

## 3. Content and formatting

- Voice and tone: plain, direct, functional.
- Date, time, number, and currency formats: none shown in approved design.
- Capitalization: heading uses supplied greeting text exactly; button label uses title case: `Save`; label uses title case: `Greeting`.
- Empty-state and error-message wording pattern: no visible empty or error message in approved design; empty input returns focus to field.

## 4. Known deviations

| Where | Deviation | Why it stands | Follow-up |
|---|---|---|---|
| Spacing scale | Uses non-4px values: `3px`, `14px`, `22px`, and `-1px` for focus offset, input padding, button padding, and visually hidden label margin. | Approved design values are source of truth; compact one-page layout intentionally plain. | Revisit only if broader UI adds more screens. |
| Component states | Hover, active, disabled, loading, and error visuals are not drawn. | Approved design shows success path plus focus-visible state only. | Add states only when future approved design draws them. |

AI default check: approved design avoids purple/indigo palette, decorative gradients, maximum rounding, oversized nested padding, heavy shadows, generic feature-grid layout, emoji iconography, filler copy, removed focus states, blank empty states, text over images, and hover-only affordances.

## 5. Change log

| Date | Change | Design PR |
|---|---|---|
| 2026-09-15 | Initial design system extracted from approved `index.html`. | Current PR |
