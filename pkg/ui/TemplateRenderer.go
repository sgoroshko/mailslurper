// Copyright 2013-2018 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package ui

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/mailslurper/mailslurper/static"
)

var templates map[string]*template.Template

// TemplateRenderer
type TemplateRenderer struct {
	templates *template.Template
}

// NewTemplateRenderer
func NewTemplateRenderer() *TemplateRenderer {
	renderer := &TemplateRenderer{}
	renderer.LoadTemplates()
	return renderer
}

// Render
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
	var tmpl *template.Template
	var ok bool

	if tmpl, ok = templates[name]; !ok {
		return fmt.Errorf("cannot find template %s", name)
	}

	return tmpl.ExecuteTemplate(w, "layout", data)
}

// LoadTemplates
func (t *TemplateRenderer) LoadTemplates() {
	templates = make(map[string]*template.Template)

	templates["mainLayout:admin"] = template.Must(
		template.New("layout").
			ParseFS(static.Files,
				"www/mailslurper/layouts/mainLayout.gohtml",
				"www/mailslurper/pages/admin.gohtml"))

	templates["mainLayout:index"] = template.Must(template.New("layout").
		ParseFS(static.Files,
			"www/mailslurper/layouts/mainLayout.gohtml",
			"www/mailslurper/pages/index.gohtml"))

	templates["mainLayout:manageSavedSearches"] = template.Must(template.New("layout").
		ParseFS(static.Files,
			"www/mailslurper/layouts/mainLayout.gohtml",
			"www/mailslurper/pages/manageSavedSearches.gohtml"))

	templates["loginLayout:login"] = template.Must(template.New("layout").
		ParseFS(static.Files,
			"www/mailslurper/layouts/loginLayout.gohtml",
			"www/mailslurper/pages/login.gohtml"))
}
