package main

import (
	"users"
	"fmt"
)

func main()  {
    fmt.Println("Main running.")

	users.NewUser("myname", "email1", "password")

	u := users.GetUser("email1")

	fmt.Println(u.Date_Started.String())

	if users.Auth("email1", "password") {
		fmt.Println("true")
	} else {
		fmt.Println("There's a problem")
	}
}