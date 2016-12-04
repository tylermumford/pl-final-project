package main

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"users"
)

func main() {

	// Setting up the controller

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		renderTemplate(w, "index.html", struct{ PageTitle string }{"Home"})
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "about.html", struct{ PageTitle string }{"About the Site"})
	})

	http.HandleFunc("/args/", func(w http.ResponseWriter, r *http.Request) {
		argID, found := findArgIDInPath(r.URL.Path)
		a := getArg(argID)
		if !found || a.ID == "" {
			w.WriteHeader(http.StatusNotFound)
			renderTemplate(w, "error.html", struct{ PageTitle string }{"Not Found"})
			return
		}

		data := struct {
			PageTitle string
			argument
		}{a.Description, a}
		renderTemplate(w, "args.html", data)
	})

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Require user to be logged in.
		renderTemplate(w, "create.html", struct{ PageTitle string }{"Create"})
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
			http.Redirect(w, r, "/error", http.StatusNotFound)
		}

	})

	http.HandleFunc("/downvote/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Require user to be logged in.
		if argID, found := findArgIDInPath(r.URL.Path); found {
			downvote(argID)
			http.Redirect(w, r, "/args/"+argID, http.StatusSeeOther)
		} else {
			w.WriteHeader(http.StatusNotFound)
			renderTemplate(w, "error.html", struct{ PageTitle string }{"Error"})
		}
	})

	http.HandleFunc("/signup/submit/", func(w http.ResponseWriter, r *http.Request) {
		fname := r.PostFormValue("fname")
		lname := r.PostFormValue("lname")
		email := r.PostFormValue("email")
		pwd := r.PostFormValue("pwd")
		confpwd := r.PostFormValue("confpwd")
		if confpwd != pwd {
			//something to Give an error and return them to the signup page
			fmt.Fprintln(w, "Those 2 passwords weren't the same. Please try again.")
		}

		u := users.GetUser(email)
		e := users.User{}

		if u.Email != e.Email {
			// User already exists. Log in.
			// print ("That user already exist. Please use the login link to login")
			fmt.Fprintln(w, "We already have a user with that address. Please login.")
			return
		}
		users.NewUser(fname+lname, email, pwd)
	})

	http.HandleFunc("/signup/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "signup.html", nil)
	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "error.html", nil)
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
