package templates

import (
	"html/template"
	"math"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
)

// Setup memuat semua template HTML dan static assets
func Setup(r *gin.Engine) {
	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"round": func(f float64) float64 {
			return math.Round(f)
		},
		"firstChar": func(s string) string {
			if s == "" {
				return "?"
			}
			firstRune, _ := utf8.DecodeRuneInString(s)
			return string(unicode.ToUpper(firstRune))
		},
		"split": func(s, sep string) []string {
			return strings.Split(s, sep)
		},
	}

	// 1. Parse sections (navbar, footer, sidebar, topbar, produk_section, faq_section, testimoni_section, etc.)
	all := template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/sections/*.html"))

	// 2. Parse all layouts
	all = template.Must(template.Must(all.Clone()).ParseGlob("templates/layout-*.html"))

	// 3. Parse index.html — uses {{define "content-index"}}
	all = template.Must(template.Must(all.Clone()).ParseGlob("templates/index.html"))

	// 4. Parse each page file individually — each uses {{define "content-*"}}
	pageFiles := []string{
		"templates/_pages/quiz.html",
		"templates/_pages/paywall.html",
		"templates/_pages/hasil.html",
		"templates/_pages/tentang.html",
		"templates/_pages/login.html",
		"templates/_pages/dashboard.html",
		"templates/_pages/user_detail.html",
		"templates/_pages/error.html",
	}
	for _, f := range pageFiles {
		parsed := template.Must(template.Must(all.Clone()).ParseFiles(f))
		for _, t := range parsed.Templates() {
			if t.Name() != "" && strings.HasPrefix(t.Name(), "content-") {
				all = template.Must(all.AddParseTree(t.Name(), t.Tree))
			}
		}
	}

	// 5. Create wrapper templates named after each page file so Gin's c.HTML() can find them.
	//    Each wrapper directly renders the full HTML by calling the layout's components inline.
	//
	//    CRITICAL: Wrapper templates must NOT use {{define}} blocks! Go's template system stores
	//    {{define}} blocks globally in the template set. If multiple wrappers define the same name
	//    (e.g. {{define "content"}}), the last one processed overwrites all previous ones, causing
	//    every page to render the same (wrong) content.
	//
	//    Layout component mapping:
	//      Public pages  → navbar + footer (index, tentang, paywall, hasil, error)
	//      Quiz page     → focused, no navbar/footer
	//      Auth page     → no navbar/footer, centered (login)
	//      Dashboard     → sidebar + topbar (dashboard, user_detail)
	wrapperTemplates := map[string]string{
		// Public pages — navbar + footer
		"index.html":   `<!DOCTYPE html><html lang="id">{{template "head" .}}<body class="{{if .BodyClass}}{{.BodyClass}}{{end}}">{{template "navbar" .}}<main>{{template "content-index" .}}</main>{{template "footer" .}}</body></html>`,
		"tentang.html": `<!DOCTYPE html><html lang="id">{{template "head" .}}<body class="{{if .BodyClass}}{{.BodyClass}}{{end}}">{{template "navbar" .}}<main>{{template "content-tentang" .}}</main>{{template "footer" .}}</body></html>`,
		"paywall.html": `<!DOCTYPE html><html lang="id">{{template "head" .}}<body class="{{if .BodyClass}}{{.BodyClass}}{{end}}">{{template "navbar" .}}<main>{{template "content-paywall" .}}</main>{{template "footer" .}}</body></html>`,
		"hasil.html":   `<!DOCTYPE html><html lang="id">{{template "head" .}}<body class="{{if .BodyClass}}{{.BodyClass}}{{end}}">{{template "navbar" .}}<main>{{template "content-hasil" .}}</main>{{template "footer" .}}</body></html>`,
		"error.html":   `<!DOCTYPE html><html lang="id">{{template "head" .}}<body class="{{if .BodyClass}}{{.BodyClass}}{{end}}">{{template "navbar" .}}<main>{{template "content-error" .}}</main>{{template "footer" .}}</body></html>`,

		// Quiz page — focused, no public navbar/footer
		"quiz.html": `<!DOCTYPE html><html lang="id">{{template "head" .}}<body class="{{if .BodyClass}}{{.BodyClass}}{{end}}"><main>{{template "content-quiz" .}}</main></body></html>`,

		// Auth page — no navbar/footer, centered
		"login.html": `<!DOCTYPE html><html lang="id">{{template "head" .}}<body class="{{if .BodyClass}}{{.BodyClass}}{{end}}"><main>{{template "content-login" .}}</main></body></html>`,

		// Dashboard pages — sidebar + topbar
		"dashboard.html":   `<!DOCTYPE html><html lang="id">{{template "head" .}}<body class="{{if .BodyClass}}{{.BodyClass}}{{end}}">{{template "sidebar" .}}<div class="dashboard-main">{{template "topbar" .}}<main>{{template "content-dashboard" .}}</main></div></body></html>`,
		"user_detail.html": `<!DOCTYPE html><html lang="id">{{template "head" .}}<body class="{{if .BodyClass}}{{.BodyClass}}{{end}}">{{template "sidebar" .}}<div class="dashboard-main">{{template "topbar" .}}<main>{{template "content-user_detail" .}}</main></div></body></html>`,
	}

	for name, content := range wrapperTemplates {
		template.Must(all.New(name).Parse(content))
	}

	r.SetHTMLTemplate(all)
	r.Static("/assets", "./assets")
}
