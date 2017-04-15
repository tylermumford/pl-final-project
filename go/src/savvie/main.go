// SAVvie is a website for social argument voting and discussion.
//
// main.go implements the HTTP server functionality of SAVvie. It handles requests by fetching
// appropriate data from subpackages and presenting it with the view package.
//
// Written as part of Code Camp and as the final project for Programming Languages class.

package main

import (
	"fmt"
	"net/http"
	"regexp"
	"savvie/storage"
	"savvie/users"
	"savvie/views"
)

func main() {

	// Setting up the controller

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		user := getLoggedIn(r)
		data := views.NewViewData("Home", user)
		views.RenderView(w, "index.html", data)
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		user := getLoggedIn(r)
		data := views.NewViewData("About the Site", user)
		views.RenderView(w, "about.html", data)
	})

	http.HandleFunc("/choices/", func(w http.ResponseWriter, r *http.Request) {
		// Show all arguments
		if r.URL.Path == "/choices/" {
			views.RenderView(w, "error.html", views.Title("This page is deprecated; it shouldn't be reachable."))
			return
		}

		// Show a specific choice
		argID, found := findArgIDInPath(r.URL.Path)
		a := storage.GetChoice(argID)
		if !found || a.ID == "" {
			w.WriteHeader(http.StatusNotFound)
			views.RenderView(w, "error.html", views.Title("Not found"))
			return
		}

		data := views.NewViewData(a.Description, getLoggedIn(r))
		data.Key["argument"] = a
		data.Key["decision"] = storage.DecisionForChoice(a)
		data.Key["comments"] = storage.LoadComments(a.ID)
		views.RenderView(w, "args.html", data)
	})

	http.HandleFunc("/decisions/", func(w http.ResponseWriter, r *http.Request) {
		// Show all decisions
		if r.URL.Path == "/decisions/" {
			data := views.NewViewData("All decisions", getLoggedIn(r))
			data.Key["decisions"] = storage.ListDecisions()
			views.RenderView(w, "all-decisions.html", data)
			return
		}

		// Show a specific decision
		decisionID, found := findDecisionIDInPath(r.URL.Path)
		d := storage.GetDecision(decisionID)
		if !found || d.ID == "" {
			w.WriteHeader(http.StatusNotFound)
			views.RenderView(w, "error.html", views.Title("Not found"))
			return
		}

		data := views.NewViewData(d.Description, getLoggedIn(r))
		data.Key["decision"] = d
		data.Key["choices"] = d.Choices
		views.RenderView(w, "decision.html", data)
	})

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		if !requireLoggedIn(w, r) {
			return
		}
		data := views.NewViewData("Create an argument", getLoggedIn(r))
		views.RenderView(w, "create.html", data)
	})

	http.HandleFunc("/create-submit", func(w http.ResponseWriter, r *http.Request) {
		if !requireLoggedIn(w, r) {
			return
		}
		descr := r.PostFormValue("descr")
		newArg := storage.SaveNewChoice(descr)
		http.Redirect(w, r, "/choices/"+newArg, http.StatusFound)
	})

	http.HandleFunc("/upvote/", func(w http.ResponseWriter, r *http.Request) {
		if !requireLoggedIn(w, r) {
			return
		}
		if argID, found := findArgIDInPath(r.URL.Path); found {
			storage.Upvote(argID)
			http.Redirect(w, r, "/choices/"+argID, http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/error", http.StatusNotFound)
		}

	})

	http.HandleFunc("/downvote/", func(w http.ResponseWriter, r *http.Request) {
		if !requireLoggedIn(w, r) {
			return
		}
		if argID, found := findArgIDInPath(r.URL.Path); found {
			storage.Downvote(argID)
			http.Redirect(w, r, "/choices/"+argID, http.StatusSeeOther)
		} else {
			w.WriteHeader(http.StatusNotFound)
			views.RenderView(w, "error.html", views.Title("Error"))
		}
	})

	http.HandleFunc("/comment/", func(w http.ResponseWriter, r *http.Request) {
		if !requireLoggedIn(w, r) {
			return
		}
		data := views.NewViewData("Error", getLoggedIn(r))

		argID, found := findArgIDInPath(r.URL.Path)
		if !found {
			views.RenderView(w, "error.html", data)
			return
		}

		err := storage.SaveNewComment(getLoggedIn(r).Email, argID, r.FormValue("commentBody"), r.FormValue("type"))
		if err != nil {
			data := views.NewViewData("Error", getLoggedIn(r))
			data.Key["errorMessage"] = err.Error()
			views.RenderView(w, "error.html", data)
			return
		}

		http.Redirect(w, r, "/choices/"+argID, http.StatusSeeOther)
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
		views.RenderView(w, "signup.html", views.Title("Sign up"))
	})

	http.HandleFunc("/login/", func(w http.ResponseWriter, r *http.Request) {
		email := r.PostFormValue("email")
		pwd := r.PostFormValue("pwd")

		if email == "" {
			views.RenderView(w, "login.html", views.Title("Log in"))
			return
		}

		correct := users.Auth(email, pwd)
		if !correct {
			data := views.NewViewData("Log in", users.User{})
			data.Key["lastAttempt"] = email
			views.RenderView(w, "login.html", data)
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
		views.RenderView(w, "error.html", views.ViewData{})
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

func findDecisionIDInPath(p string) (string, bool) {
	re := regexp.MustCompile("/([0-9]+)$")
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
		data := views.NewViewData("Error", users.User{})
		data.Key["errorMessage"] = views.Htmlstr(`You must <a href="/login">Log in</a> to access this page.`)
		views.RenderView(w, "error.html", data)
		return false
	}
	return true
}
