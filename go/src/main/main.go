package main

import (
	"users"
	"fmt"
)

func main()  {
    fmt.Println("Main running.")

	a := users.User{"uid", "name", "pword"}

	fmt.Println(a.Name)
}