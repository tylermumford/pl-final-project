package main

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
)

func main() {

	// Setting up the controller

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		renderTemplate(w, "index.html", nil)
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "about.html", nil)
	})

	http.HandleFunc("/args/", func(w http.ResponseWriter, r *http.Request) {

		if argID, found := findArgIDInPath(r.URL.Path); found {
			a := getArg(argID)
			renderTemplate(w, "args.html", a)
		}

		// fmt.Fprint(w, "<h1>All Arguments</h1>")

		// //	TODO: Load arguments from Nick
		// arguments := loadArguments(0)

		// if len(arguments) == 0 {
		// 	fmt.Fprint(w, "<p>No arguments</p>")
		// } else {
		// 	fmt.Fprint(w, "<ul>")
		// 	for _, a := range arguments {
		// 		fmt.Fprintf(w, "<li>%v</li>", a.title)
		// 	}
		// 	fmt.Fprintf(w, "</ul>")
		// }

	})

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Require user to be logged in.
		renderTemplate(w, "create.html", nil)
	})

	http.HandleFunc("/create-submit", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Require user to be logged in.
		descr := r.PostFormValue("descr")
		saveNewArgument(descr)
	})

	http.HandleFunc("/upvote/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Require user to be logged in.
		if argID, found := findArgIDInPath(r.URL.Path); found {
			upvote(argID)
			http.Redirect(w, r, "/args/"+argID, http.StatusSeeOther)
		} else {
			addHeader(w)
			fmt.Fprintf(w, "Not found...")
		}

	})

	http.HandleFunc("/downvote/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Require user to be logged in.
		if argID, found := findArgIDInPath(r.URL.Path); found {
			downvote(argID)
			http.Redirect(w, r, "/args/"+argID, http.StatusSeeOther)
		} else {
			addHeader(w)
			fmt.Fprintf(w, "Does not exist")
		}
	})

	http.HandleFunc("/signup/submit/", func(w http.ResponseWriter, r *http.Request) {
		//if(user == User{})
	})

	http.HandleFunc("/signup/", func(w http.ResponseWriter, r *http.Request) {
		//renderTemplate(w, "signup.html", nil)
	})

	http.ListenAndServe("localhost:8000", nil)
}

func findArgIDInPath(p string) (string, bool) {
	re := regexp.MustCompile("/([0-9]{5})$")
	sub := re.FindStringSubmatch(p)

	if sub == nil {
		return "", false
	}

	return sub[1], true
}

func addHeader(w http.ResponseWriter) {
	fmt.Fprint(w, "<head><link href='https://cdnjs.cloudflare.com/ajax/libs/foundation/6.2.4/foundation.css' rel='stylesheet'><style>.body{max-width:600px;margin: 20 auto;}</style></head>")
}

func openBody(w http.ResponseWriter) {
	fmt.Fprint(w, "<div class='body'>")
}

func closeBody(w http.ResponseWriter) {
	fmt.Fprint(w, "</div>")

}

var allTemplates *template.Template

func loadAllTemplates() {
	allTemplates = template.Must(template.ParseGlob("/vagrant/templates/*"))
}

func renderTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	if allTemplates == nil {
		loadAllTemplates()
	}
	err := allTemplates.ExecuteTemplate(w, templateName, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
