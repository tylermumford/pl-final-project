package main

import (
	"fmt"
	"net/http"
	"users"
	"regexp"
)

func main() {

	// Testing "users" package

	fmt.Println("Main running.")

	users.NewUser("myname", "email1", "password")

	u := users.GetUser("email1")

	fmt.Println(u.Date_Started.String())

	if users.Auth("email1", "password") {
		fmt.Println("true")
	} else {
		fmt.Println("There's a problem")
	}

	// Setting up the controller

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		fmt.Fprint(w, "<p>Hello from <strong>Go</strong></p>")
	})

	http.HandleFunc("/args", func(w http.ResponseWriter, r *http.Request) {
		if argID, found := findArgIDInPath(r.URL.Path); found {
			_ = argID
			// displayArg(argId)
			return
		}

		fmt.Fprint(w, "<h1>All Arguments</h1>")

		// TODO: Load arguments from Nick
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

	http.HandleFunc("/upvote/", func(w http.ResponseWriter, r *http.Request){
		if argID, found := findArgIDInPath(r.URL.Path); found {
			_ = argID
			fmt.Fprintf(w,argID)
			return
		} else {
			fmt.Fprintf(w, "Not found...")
		}
	})

	http.HandleFunc("/downvote/", func(w http.ResponseWriter, r *http.Request) {
		if argID, found := findArgIDInPath(r.URL.Path); found {
			fmt.Fprintf(w, argID)
			return
		} else {
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
	fmt.Println(sub[1])

	return sub[1], true
}
