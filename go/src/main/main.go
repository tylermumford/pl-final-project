package main

import (
	"users"
	"fmt"
)

func main()  {
    fmt.Println("Main running.")

	a := User{"uid", "fname", "lname", "pword"}

	fmt.Println(a.fname)
}