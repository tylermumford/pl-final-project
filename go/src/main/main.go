package main

import (
	"users"
	"fmt"
)

func main()  {
    fmt.Println("Main running.")

	users.NewUser("myname", "email1", "password")

	fmt.Println(users.GetName("email1"))

	if users.Auth("email1", "password") {
		fmt.Println("true")
	} else {
		fmt.Println("There's a problem")
	}
}