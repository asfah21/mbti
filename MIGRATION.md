# Migration Plan: `html/template` + Gin → `templ`

## Overview

Migrate this ShadowSelf/MBTI assessment platform from Go's `html/template` (via Gin) to [`templ`](https://github.com/a-h/templ) — a type-safe HTML templating language that compiles to Go code.

### Current Stack

| Layer | Technology |
|---|---|
| HTTP Framework | Gin (`github.com/gin-gonic/gin`) |
| Template Engine | `html/template` (stdlib) |
| Template Setup | `templates/setup.go` — 101 lines, FuncMap + ParseGlob + wrapper templates |
| Frontend Interactivity | Alpine.js (CDN), Vanilla JS |
| CSS | Custom (`assets/css/style.css`) |
| Icons | Material Symbols (Google Fonts) |

### Target Stack

| Layer | Technology |
|---|---|
| HTTP Framework | Gin (unchanged) |
| Template Engine | `templ` (`github.com/a-h/templ`) |
| Template Setup | `//go:generate templ generate` |
| Frontend Interactivity | Alpine.js (unchanged) |
| CSS | Custom (unchanged) |
| Icons | Material Symbols (unchanged) |

---

## Current Architecture (Deep Dive)

### Template Loading Flow

```
main.go → templates.Setup(r)
  └── templates/setup.go
        ├── FuncMap {add, round, firstChar, split}
        ├── ParseGlob("templates/sections/*.html")     → 8 section templates
        ├── ParseGlob("templates/layout-*.html")        → 4 layout templates
        ├── ParseFiles("templates/index.html")           → content-index
        ├── ParseFiles("templates/_pages/*.html")        → 8 content-* blocks
        └── Wrapper templates (inline strings mapping filename → full HTML)
              └── r.SetHTMLTemplate(all)
              └── r.Static("/assets", "./assets")
```

### Template File Map

| File | Type | `{{define}}` Name | Data Dependencies |
|---|---|---|---|
| `sections/head.html` | Component | `head` | `.Title` |
| `sections/navbar.html` | Component | `navbar` | `.BodyClass` |
| `sections/footer.html` | Component | `footer` | none |
| `sections/sidebar.html` | Component | `sidebar` | none |
| `sections/topbar.html` | Component | `topbar` | none |
| `sections/faq_section.html` | Component | `faq_section` | none |
| `sections/produk_section.html` | Component | `produk_section` | none |
| `sections/testimoni_section.html` | Component | `testimoni_section` | none |
| `layout-public.html` | Layout | — | `.Title`, `.BodyClass` |
| `layout-dashboard.html` | Layout | — | `.Title`, `.BodyClass` |
| `layout-quiz.html` | Layout | — | `.Title`, `.BodyClass` |
| `layout-auth.html` | Layout | — | `.Title`, `.BodyClass` |
| `index.html` | Page | `content-index` | none |
| `_pages/quiz.html` | Page | `content-quiz` | none (Alpine.js handles state) |
| `_pages/paywall.html` | Page | `content-paywall` | `.ID`, `.Nama` |
| `_pages/hasil.html` | Page | `content-hasil` | `.Nama`, `.Narsisme`, `.Machiavellian`, etc. |
| `_pages/tentang.html` | Page | `content-tentang` | none |
| `_pages/login.html` | Page | `content-login` | `.Error` |
| `_pages/dashboard.html` | Page | `content-dashboard` | `.Users`, `.TotalUser`, `.SudahBayar`, `.BelumBayar`, `.TotalPendapatan` |
| `_pages/user_detail.html` | Page | `content-user_detail` | `.User` |
| `_pages/error.html` | Page | `content-error` | `.Message` |

### Route → Handler → Template Map

| Route | Handler | Template | Data |
|---|---|---|---|
| `GET /` | `ShowHome` | `index.html` | — |
| `GET /quiz` | `ShowQuiz` | `quiz.html` | — |
| `POST /submit-tes` | `SubmitTest` | JSON response | — |
| `GET /paywall/:id` | `ShowPaywall` | `paywall.html` | `{ID, Nama}` |
| `POST /konfirmasi-bayar/:id` | `KonfirmasiBayar` | JSON response | — |
| `GET /hasil/:id` | `ShowResult` | `hasil.html` | `{Nama, MBTI, SkorEI, ..., ExecutiveSummary, ...}` |
| `GET /tentang` | `ShowTentang` | `tentang.html` | — |
| `GET /admin/login` | `ShowLogin` | `login.html` | `.Error` (on failure) |
| `POST /admin/login` | `LoginProcess` | `login.html` or redirect | `.Error` |
| `GET /admin/dashboard` | `ShowDashboard` | `dashboard.html` | `{Users, TotalUser, SudahBayar, ...}` |
| `GET /admin/user/:id` | `ShowUserDetail` | `user_detail.html` | `{User}` |
| `GET /admin/logout` | `LogoutProcess` | redirect | — |
| 404 | `Show404` | `error.html` | `{Message}` |

### FuncMap Usage in Templates

| FuncMap Entry | Usage (template) | templ Replacement |
|---|---|---|
| `add(a, b int)` | `{{add $i 1}}` | `i + 1` in Go |
| `round(f float64)` | `{{round .Narsisme}}` | `math.Round(float64(...))` in Go |
| `firstChar(s string)` | `{{firstChar .Nama}}` | Go helper or inline |
| `split(s, sep string)` | `{{split $profile "\n\n"}}` | `strings.Split(profile, "\n\n")` in Go |

---

## Known Data Model Mismatch (Must Fix During Migration)

**Bug identified:** The `ShowResult` handler in `handlers/quiz.go` passes **MBTI raw scores** (`SkorEI`, `SkorSN`, etc.) but the `hasil.html` template renders **Dark Triad percentiles** (`Narsisme`, `Machiavellian`, `Psikopati`). The narrative generator `services/narasi.go` (`GenerateAllNarratives`) is never actually called — placeholder text is used instead.

**Fix:** Create a proper `HasilData` struct that includes both the fully generated narratives from `GenerateAllNarratives` and any MBTI data. Call `GenerateAllNarratives` in the handler.

---

## Target Architecture (After Migration)

```
templ/
├── components/
│   ├── head.templ          # <head> block
│   ├── navbar.templ        # Alpine.js responsive nav
│   ├── footer.templ        # Site footer
│   ├── sidebar.templ       # Dashboard sidebar
│   ├── topbar.templ        # Dashboard topbar
│   ├── faq_section.templ   # FAQ accordion
│   ├── produk_section.templ
│   └── testimoni_section.templ
├── layouts/
│   ├── public_layout.templ     # navbar + footer
│   ├── quiz_layout.templ       # minimal, no nav
│   ├── auth_layout.templ       # centered, no nav
│   └── dashboard_layout.templ  # sidebar + topbar
└── pages/
    ├── index_page.templ        # Landing page
    ├── quiz_page.templ         # Alpine.js quiz app
    ├── paywall_page.templ      # Payment gate
    ├── hasil_page.templ        # Assessment results
    ├── tentang_page.templ      # About page
    ├── login_page.templ        # Admin login
    ├── dashboard_page.templ    # Admin dashboard
    ├── user_detail_page.templ  # User detail (admin)
    └── error_page.templ        # 404/500 errors
```

### Template Loading (After)

```
main.go → handlers.SetupRoutes(r)
  └── No more templates.Setup(r) call
  └── r.Static("/assets", "./assets") // unchanged
  └── Handler calls: pages.ComponentName().Render(ctx, c.Writer)
```

---

## Migration Phases

### Phase 1: Project Setup & Scaffolding
- [ ] Add `templ` dependency: `go get github.com/a-h/templ`
- [ ] Install `templ` CLI: `go install github.com/a-h/templ/cmd/templ@latest`
- [ ] Create `templ/` directory structure
- [ ] Create `generate.go` with `//go:generate templ generate`
- [ ] Verify `templ generate` runs cleanly on empty scaffold

### Phase 2: Component Migration — Sections
- [ ] `templ/components/head.templ`
- [ ] `templ/components/navbar.templ`
- [ ] `templ/components/footer.templ`
- [ ] `templ/components/sidebar.templ`
- [ ] `templ/components/topbar.templ`
- [ ] `templ/components/faq_section.templ`
- [ ] `templ/components/produk_section.templ`
- [ ] `templ/components/testimoni_section.templ`

### Phase 3: Layout Migration
- [ ] `templ/layouts/public_layout.templ`
- [ ] `templ/layouts/quiz_layout.templ`
- [ ] `templ/layouts/auth_layout.templ`
- [ ] `templ/layouts/dashboard_layout.templ`

### Phase 4: Page Migration — Static Pages (no data dependencies)
- [ ] `templ/pages/index_page.templ`
- [ ] `templ/pages/tentang_page.templ`
- [ ] `templ/pages/error_page.templ`
- [ ] `templ/pages/login_page.templ`

### Phase 5: Page Migration — Dynamic Pages (with data dependencies)
- [ ] `templ/pages/quiz_page.templ` (Alpine.js app, no Go data)
- [ ] `templ/pages/paywall_page.templ` (needs `PaywallData`)
- [ ] `templ/pages/hasil_page.templ` (needs unified `HasilData`)
- [ ] `templ/pages/dashboard_page.templ` (needs user list + stats)
- [ ] `templ/pages/user_detail_page.templ` (needs `User`)

### Phase 6: Create Data Models for templ Pages
- [ ] Create `templ/types/hasil_data.go` — unified result data struct
- [ ] Create `templ/types/dashboard_data.go` — dashboard data structs
- [ ] Ensure all narrative fields are populated via `GenerateAllNarratives`

### Phase 7: Remove `templates/setup.go`
- [ ] Delete `templates/setup.go`
- [ ] Remove `html/template` import
- [ ] Ensure no remaining references to `r.SetHTMLTemplate()`

### Phase 8: Rewrite Handlers to Use templ
- [ ] Create `helpers/render.go` with `Render(c, status, component)` helper
- [ ] Rewrite `handlers/page.go` — all page handlers
- [ ] Rewrite `handlers/quiz.go` — ShowPaywall, ShowResult
- [ ] Rewrite `handlers/admin.go` — ShowDashboard, ShowUserDetail

### Phase 9: Fix Data Model Mismatch
- [ ] Update `GetQuizResult` service to call `GenerateAllNarratives`
- [ ] Pass proper `HasilData` to hasil_page

### Phase 10: Cleanup & Build Integration
- [ ] Delete `templates/` directory entirely
- [ ] Update `Dockerfile` to run `templ generate` before `go build`
- [ ] Update `task_progress.md` / docs
- [ ] Run `go mod tidy`
- [ ] Full application test

---

## Render Helper (to be created)

```go
// helpers/render.go
package helpers

import (
    "net/http"
    "github.com/a-h/templ"
    "github.com/gin-gonic/gin"
)

func Render(c *gin.Context, status int, component templ.Component) {
    c.Status(status)
    if err := component.Render(c.Request.Context(), c.Writer); err != nil {
        // Fallback: log error, send plain text
        http.Error(c.Writer, "Internal Server Error", http.StatusInternalServerError)
    }
}
```

---

## Handler Migration Examples

### Before (Current)
```go
// handlers/page.go
func ShowHome(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", nil)
}

// handlers/quiz.go
func ShowPaywall(c *gin.Context) {
    data, err := services.GetPaywallData(id)
    // ... error handling ...
    c.HTML(http.StatusOK, "paywall.html", gin.H{
        "ID":   data.ID,
        "Nama": data.Nama,
    })
}
```

### After (templ)
```go
// handlers/page.go
func ShowHome(c *gin.Context) {
    helpers.Render(c, http.StatusOK, pages.IndexPage())
}

// handlers/quiz.go
func ShowPaywall(c *gin.Context) {
    data, err := services.GetPaywallData(id)
    // ... error handling ...
    helpers.Render(c, http.StatusOK, pages.PaywallPage(*data))
}
```

---

## Alpine.js Compatibility Notes

templ passes all HTML attributes through unchanged. This means Alpine.js works seamlessly:

| Alpine Feature | Current Usage | templ Compatibility |
|---|---|---|
| `x-data` | `x-data="{ mobileOpen: false }"` | ✅ Passes through |
| `x-init` | `x-init="init()"` | ✅ Passes through |
| `x-show` | `x-show="currentStep === idx"` | ✅ Passes through |
| `x-cloak` | `x-cloak` | ✅ Passes through |
| `x-transition` | `x-transition:enter="..."` | ✅ Passes through |
| `x-model` | `x-model="nama"` | ✅ Passes through |
| `x-text` | `x-text="q.text"` | ✅ Passes through |
| `x-html` | Not used | ✅ Passes through |
| `@click` | `@click="answers[idx] = opt.value"` | ✅ Passes through |
| `@submit.prevent` | `@submit.prevent="submitQuiz()"` | ✅ Passes through |
| `:style` | `:style="'width:' + ... + '%'"` | ✅ Passes through |
| `:class` | `:class="{ 'is-selected': ... }"` | ✅ Passes through |

---

## Inline `<script>` Migration

Pages with inline JS that uses Go template interpolation (`{{.ID}}`, `{{.Nama}}`) need special handling:

### Paywall Page (`paywall.html`)
```javascript
// BEFORE (inside Go template)
fetch('/konfirmasi-bayar/{{.ID}}', {...})

// AFTER (templ — string interpolation in script)
<script>
  const id = "{ data.ID }";
  fetch('/konfirmasi-bayar/' + id, {...})
</script>
```

### Quiz Page (`quiz.html`)
The entire `quizApp()` function (500 lines) remains as an inline `<script>` with zero changes — it contains no Go template interpolation.

---

## Build Process Changes

### Before
```bash
go build -o mbti .
```

### After
```bash
templ generate
go build -o mbti .
```

### Dockerfile Update
```dockerfile
# Before
RUN go build -o mbti .

# After
RUN go install github.com/a-h/templ/cmd/templ@latest && \
    templ generate && \
    go build -o mbti .
```

---

## Risk Matrix

| Risk | Likelihood | Impact | Mitigation |
|---|---|---|---|
| Alpine.js breakage | Low | High | All Alpine.js attributes pass through unchanged |
| Data model mismatch | Medium | High | Fix in Phase 9; compiler catches field errors |
| Inline JS template interpolation | Low | Medium | Migrate to JS string concat per paywall pattern |
| Build process change | Low | Medium | Update Dockerfile + CI scripts |
| Missing `split`/`add` in template | Low | Medium | Replace with Go functions in page components |
| `firstChar` used for avatars | Low | Low | Implement as Go package-level helper |

---

## Rollback Plan

If critical issues arise during or after migration:

1. Revert to the last known-good commit (`0d54c6dbed439befc78df69e0390389d84c0f28d`)
2. Keep `templates/setup.go` intact until Phase 7 is fully validated
3. Each phase is independently revertible