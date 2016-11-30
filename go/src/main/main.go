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
		addHeader(w)
		fmt.Fprint(w, "<p>Hello from <strong>Go</strong></p>")
	})

	http.HandleFunc("/args/", func(w http.ResponseWriter, r *http.Request) {
		addHeader(w)

		if argID, found := findArgIDInPath(r.URL.Path); found {
			fmt.Fprintf(w, "<h1>Argument %v:</h1>", argID)
			openBody(w)
			displayArg(w, argID)
			fmt.Fprintf(w, "<p><a href='/upvote/%v' class='button' style='float:right'>Upvote</a></p>", argID)
			fmt.Fprintf(w, "<p><a href='/downvote/%v' class='button alert' style='float:right'>Downvote</a></p>", argID)
			closeBody(w)
			return
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

	http.HandleFunc("/upvote/", func(w http.ResponseWriter, r *http.Request) {
		if argID, found := findArgIDInPath(r.URL.Path); found {
			go upvote(argID)
			http.Redirect(w, r, "/args/"+argID, http.StatusSeeOther)
		} else {
			addHeader(w)
			fmt.Fprintf(w, "Not found...")
		}

	})

	http.HandleFunc("/downvote/", func(w http.ResponseWriter, r *http.Request) {
		if argID, found := findArgIDInPath(r.URL.Path); found {
			go downvote(argID)
			http.Redirect(w, r, "/args/"+argID, http.StatusSeeOther)
		} else {
			addHeader(w)
			fmt.Fprintf(w, "Does not exist")
		}
	})

	http.HandleFunc("/create/", func(w http.ResponseWriter, r *http.Request) {
		descr := r.PostFormValue("description")
		saveNewArgument(descr)
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

func displayArg(w http.ResponseWriter, argID string) {
	arg := getArg(argID)

	fmt.Fprintf(w, "<div class='callout'><h2>%v</h2><ul><li>Score: %v</li><li>Upvotes: %v</li><li>Downvotes: %v</li></div>", arg.description, arg.upvotes-arg.downvotes, arg.upvotes, arg.downvotes)
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
