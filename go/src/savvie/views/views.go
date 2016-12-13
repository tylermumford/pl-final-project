// Package views handles template loading and rendering to the browser.
//
// Use NewViewData to create a struct that can be passed to templates. Title is a simple way to create such a struct.
// Use RenderView to send HTML to the browser.
package views

import (
	"html/template"
	"net/http"
	"savvie/users"
)

var allTemplates *template.Template

func loadAllTemplates() {
	allTemplates = template.Must(template.ParseGlob("/vagrant/templates/*"))
}

// RenderView takes the filename of a template and passes it the given data argument.
// It then sends the resulting HTML to the browser via the given ResponseWriter.
func RenderView(w http.ResponseWriter, templateName string, data ViewData) {
	loadAllTemplates()
	err := allTemplates.ExecuteTemplate(w, templateName, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ViewData represents the data that templates can render. All templates expect a
// PageTitle and (the currently logged in) User, but they can also be passed ad-hoc data
// via the Key map.
//
// This allows all templates to depend on the structured fields, yet developers can pass any
// kind of data to specific templates freely.
type ViewData struct {
	PageTitle string
	User      users.User
	Key       map[string]interface{}
}

// NewViewData uses the given title and user to initialize a ViewData struct.
// This is the recommended way to create ViewData structs, because it sets system-wide defaults
// for uninitialized fields.
func NewViewData(title string, usr users.User) ViewData {
	t := ViewData{}
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

// Title is a helper function that returns a data struct that can be
// passed to `renderTemplate` to set the page title.
func Title(title string) ViewData {
	t := NewViewData("", users.User{})
	t.PageTitle = title
	return t
}

// Htmlstr simply wraps a string in a way that marks it as "safe." Safe strings
// will not be escaped (that is, they can contain raw HTML) when passed to templates.
func Htmlstr(s string) template.HTML {
	return template.HTML(s)
}
