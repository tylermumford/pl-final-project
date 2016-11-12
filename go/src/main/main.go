package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Testing "users" package

	saveNewArgument("'Waffles are better than pancakes.'")

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

	http.ListenAndServe("localhost:8000", nil)
}

func findArgIDInPath(p string) (string, bool) {
	return "", false
	// if p[0:len(p)-3]
}
