package handlers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// loadGlobalTemplate parse all templates are used as layout
// Example: templates/layout/base.tmpl, templates/layout/header.tmpl,...
func loadGlobalTemplate() (*template.Template, error) {
	return template.ParseGlob("templates/layout/*")
}

// renderPage render single template use base layout from global template
func renderPage(c *gin.Context, pageTemplate string, data interface{}) {
	tmpl, err := loadGlobalTemplate()
	w := c.Writer
	tmpl, err = tmpl.ParseFiles(pageTemplate)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
