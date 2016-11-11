package users

import (
	//"io/ioutil"
)

const data_folder = "/vagrant/go/src/users/data"

type User struct{
	Uid string
	Name string
	Pword string
}