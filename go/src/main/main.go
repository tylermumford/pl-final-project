package main

import (
	"fmt"
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

		data := newTemplateData()
		data.PageTitle = a.Description
		data.Key["argument"] = a
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

	http.HandleFunc("/signup-submit", func(w http.ResponseWriter, r *http.Request) {
		fname := r.FormValue("fname")
		lname := r.FormValue("lname")
		email := r.FormValue("email")
		pwd := r.FormValue("pwd")
		confpwd := r.FormValue("confpwd")
		if confpwd != pwd {
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

	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "signup.html", pTitle("Sign up"))
	})

	http.HandleFunc("/login/", func(w http.ResponseWriter, r *http.Request) {
		email := r.PostFormValue("email")
		pwd := r.PostFormValue("pwd")

		if email == "" {
			renderTemplate(w, "login.html", pTitle("Log in"))
			return
		}

		correct := users.Auth(email, pwd)
		if !correct {
			data := newTemplateData()
			data.PageTitle = "Log in"
			data.Key["lastAttempt"] = email
			renderTemplate(w, "login.html", data)
			return
		}

		setLoggedIn(w, users.GetUser(email))
		fmt.Fprintf(w, "Logged in as %v", email)
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

const userCookieName = "email"

// setLoggedIn marks the current user as "logged in" by storing their
// email address in a cookie.
func setLoggedIn(w http.ResponseWriter, u users.User) {
	c := http.Cookie{
		Name:   userCookieName,
		Value:  u.Email,
		Secure: false,
	}
	http.SetCookie(w, &c)
}

// getLoggedIn returns the User associated with the given request, based
// on a cookie. If the given request did not come from a logged-in user
// (or if the claimed user does not exist), it returns an empty User.
func getLoggedIn(r http.Request) users.User {
	c, err := r.Cookie(userCookieName)
	if err != nil {
		return users.User{}
	}
	return users.GetUser(c.Value)
}
