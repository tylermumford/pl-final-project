package main

import (
	"comments"
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
		user := getLoggedIn(r)
		data := newTemplateData("Home", user)
		renderTemplate(w, "index.html", data)
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		user := getLoggedIn(r)
		data := newTemplateData("About the Site", user)
		renderTemplate(w, "about.html", data)
	})

	http.HandleFunc("/args/", func(w http.ResponseWriter, r *http.Request) {
		argID, found := findArgIDInPath(r.URL.Path)
		a := getArg(argID)
		if !found || a.ID == "" {
			w.WriteHeader(http.StatusNotFound)
			renderTemplate(w, "error.html", pTitle("Not found"))
			return
		}

		data := newTemplateData(a.Description, getLoggedIn(r))
		data.Key["argument"] = a
		data.Key["comments"] = comments.Load(a.ID)
		renderTemplate(w, "args.html", data)
	})

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		if !requireLoggedIn(w, r) {
			return
		}
		data := newTemplateData("Create an argument", getLoggedIn(r))
		renderTemplate(w, "create.html", data)
	})

	http.HandleFunc("/create-submit", func(w http.ResponseWriter, r *http.Request) {
		if !requireLoggedIn(w, r) {
			return
		}
		descr := r.PostFormValue("descr")
		newArg := saveNewArgument(descr)
		http.Redirect(w, r, "/args/"+newArg, http.StatusFound)
	})

	http.HandleFunc("/upvote/", func(w http.ResponseWriter, r *http.Request) {
		if !requireLoggedIn(w, r) {
			return
		}
		if argID, found := findArgIDInPath(r.URL.Path); found {
			upvote(argID)
			http.Redirect(w, r, "/args/"+argID, http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/error", http.StatusNotFound)
		}

	})

	http.HandleFunc("/downvote/", func(w http.ResponseWriter, r *http.Request) {
		if !requireLoggedIn(w, r) {
			return
		}
		if argID, found := findArgIDInPath(r.URL.Path); found {
			downvote(argID)
			http.Redirect(w, r, "/args/"+argID, http.StatusSeeOther)
		} else {
			w.WriteHeader(http.StatusNotFound)
			renderTemplate(w, "error.html", pTitle("Error"))
		}
	})

	http.HandleFunc("/comment/", func(w http.ResponseWriter, r *http.Request) {
		if !requireLoggedIn(w, r) {
			return
		}
		data := newTemplateData("Error", getLoggedIn(r))

		argID, found := findArgIDInPath(r.URL.Path)
		if !found {
			renderTemplate(w, "error.html", data)
			return
		}

		err := comments.Save(getLoggedIn(r).Email, argID, r.FormValue("commentBody"))
		if err != nil {
			data := newTemplateData("Error", getLoggedIn(r))
			data.Key["errorMessage"] = err.Error()
			renderTemplate(w, "error.html", data)
			return
		}

		http.Redirect(w, r, "/args/"+argID, http.StatusSeeOther)
	})

	http.HandleFunc("/signup-submit", func(w http.ResponseWriter, r *http.Request) {
		fname := r.FormValue("fname")
		lname := r.FormValue("lname")
		email := r.FormValue("email")
		pwd := r.FormValue("pwd")
		confpwd := r.FormValue("confpwd")
		var problems = make([]string, 0)
		if fname == "" || lname == "" {
			problems = append(problems, "Please provide a name.")
		}
		if email == "" {
			problems = append(problems, "Please provide an email address.")
		}
		if pwd != confpwd {
			problems = append(problems, "Passwords do not match.")
		}
		if len(problems) == 0 && users.GetUser(email).Email != "" {
			problems = append(problems, "That user already exists.")
		}

		if len(problems) != 0 {
			fmt.Fprint(w, problems)
			return
		}

		users.NewUser(fname+" "+lname, email, pwd)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		setLoggedIn(w, users.GetUser(email))
	})

	http.HandleFunc("/signup/", func(w http.ResponseWriter, r *http.Request) {
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
			data := newTemplateData("Log in", users.User{})
			data.Key["lastAttempt"] = email
			renderTemplate(w, "login.html", data)
			return
		}

		setLoggedIn(w, users.GetUser(email))
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.HandleFunc("/logout/", func(w http.ResponseWriter, r *http.Request) {
		setLoggedOut(w)
		http.Redirect(w, r, "/", http.StatusSeeOther)
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

// setLoggedIn marks the given user as "logged in" by storing their
// email address in a cookie.
func setLoggedIn(w http.ResponseWriter, u users.User) {
	c := http.Cookie{
		Name:   userCookieName,
		Value:  u.Email,
		Path:   "/",
		Secure: false,
	}
	http.SetCookie(w, &c)
}

// Logs out by setting the cookie to "".
func setLoggedOut(w http.ResponseWriter) {
	c := http.Cookie{
		Name:   userCookieName,
		Value:  "",
		Path:   "/",
		Secure: false,
	}
	http.SetCookie(w, &c)
}

// getLoggedIn returns the User associated with the given request, based
// on a cookie. If the given request did not come from a logged-in user
// (or if the claimed user does not exist), it returns an empty User.
func getLoggedIn(r *http.Request) users.User {
	c, err := r.Cookie(userCookieName)
	if err != nil {
		return users.User{}
	}
	return users.GetUser(c.Value)
}

func requireLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	if getLoggedIn(r).Email == "" {
		data := newTemplateData("Error", users.User{})
		data.Key["errorMessage"] = htmlstr(`You must <a href="/login">Log in</a> to access this page.`)
		renderTemplate(w, "error.html", data)
		return false
	}
	return true
}
