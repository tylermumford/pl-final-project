package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"regexp"
	"time"
	"users"
)

func main() {

	// General initialization
	func() {
		seconds := time.Now().Second()
		rand.Seed(int64(seconds))
	}()

	// Setting up the controller

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		renderTemplate(w, "index.html", pTitle("Home"))
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "about.html", pTitle("About the Site"))
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
		renderTemplate(w, "create.html", pTitle("Create an argument"))
	})

	http.HandleFunc("/create-submit", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Require user to be logged in.
		descr := r.PostFormValue("descr")
		newArg := saveNewArgument(descr)
		http.Redirect(w, r, "/args/"+newArg, http.StatusFound)
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
			renderTemplate(w, "error.html", pTitle("Error"))
		}
	})

	http.HandleFunc("/signup/submit/", func(w http.ResponseWriter, r *http.Request) {
		fname := r.PostFormValue("fname")
		lname := r.PostFormValue("lname")
		email := r.PostFormValue("email")
		pwd := r.PostFormValue("pwd")
		confpwd := r.PostFormValue("confpwd")
		if confpwd == pwd {
			//something to Give an error and return them to the signup page
		}

		u := users.GetUser(email)
		e := users.User{}

		if u.Name == e.Name {
			users.NewUser(fname+lname, email, pwd)
		} else {
			//something to Give an error and return them to the signup page
		}
	})

	http.HandleFunc("/signup/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "signup.html", pTitle("Sign up"))
	})

	http.HandleFunc("/login/submit/", func(w http.ResponseWriter, r *http.Request) {
		email := r.PostFormValue("email")
		pwd := r.PostFormValue("pwd")

		correct := users.Auth(email, pwd)
		if !correct {
			http.Redirect(w, r, "/login", http.StatusNotFound)
		}

		// TODO: Set the user as logged in. Store with a cookie?
	})

	http.HandleFunc("/login/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "login.html", pTitle("Log in"))
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

// pTitle is a helper function that returns a data struct that can be
// passed to `renderTemplate` to set the page title.
func pTitle(title string) interface{} {
	return struct {
		PageTitle string
	}{title}
}
