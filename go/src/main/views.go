package main

import (
	"html/template"
	"net/http"
	"users"
)

// TODO: Isolate this file into its own package.

var allTemplates *template.Template

func loadAllTemplates() {
	allTemplates = template.Must(template.ParseGlob("/vagrant/templates/*"))
}

func renderTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	loadAllTemplates()
	err := allTemplates.ExecuteTemplate(w, templateName, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type templateData struct {
	PageTitle string
	User      users.User
	Key       map[string]interface{}
}

func newTemplateData(title string, usr users.User) templateData {
	t := templateData{}
	if title == "" {
		title = "Social Argument Voting"
	}
	if usr.Email == "" {
		usr = users.User{}
	}
	t.PageTitle = title
	t.User = usr
	t.Key = make(map[string]interface{}, 1)
	return t
	// TODO: We should probably be passing around pointers to templateData,
	// instead of copying them around.
}

// pTitle is a helper function that returns a data struct that can be
// passed to `renderTemplate` to set the page title.
func pTitle(title string) interface{} {
	t := newTemplateData("", users.User{})
	t.PageTitle = title
	return t
}

func htmlstr(s string) template.HTML {
	return template.HTML(s)
}
