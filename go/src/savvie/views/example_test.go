package views_test

import (
	"net/http"
	"savvie/users"
	"savvie/views"
)

func ExampleRenderView() {
	// Typically, this would appear inside a handler function,
	// so w would be provided as a parameter.
	var w http.ResponseWriter
	data := views.NewViewData("Contact Us", users.User{})
	views.RenderView(w, "contact.html", data)
}
