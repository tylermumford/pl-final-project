/*
* Written by Katrina Mehring as part of Code Camp and final project
*   for a Programming Languages class.
*
*	This package is used by the main package. It is used to
*		hash user passwords, verify user login information,
*		and store user information. Because of hashing issues
*		at Code Camp, the old hash has been commented out, and
*		a new hash is being used. The code is left here so that
*		it can hopefully be used in future implementations.
 */

package users

import (
	"bufio"
	"os"
	"time"

	// "github.com/golang/crypto/scrypt"
	//"crypto/rand"
	//"io"
	//"fmt"
	"crypto/sha256"
	"log"
)

const (
	saltBytes  = 32
	hashBytes  = 64
	dataFolder = "/vagrant/data/users/"
)

// User a Struct made only for the purpose
// of being returned to main.go to
// verify new users don't already exist
// and to verify passwords and whatnot.
type User struct {
	Name, Email string
	pwd         []byte
	DateStarted time.Time
}

//NewUser :
// A function that allows the main package to
// create a new user after verifying that the
// username doesn't already exist (usernames
// are emails).
func NewUser(name, email, pword string) {
	f, err := os.Create(dataFolder + email + ".txt")
	check(err)
	defer f.Close()

	_, err = f.WriteString(name + "\n")
	check(err)

	date := time.Now().Format(time.RFC850)
	_, err = f.WriteString(date + "\n")

	_, err = f.Write(pHash(pword))
	check(err)

	f.Sync()
}

// GetUser :
// A function to return a user struct to the main package.
// This is used to connect a user to an argument, verify
// that a user exists or not, etc.
func GetUser(email string) (user User) {
	user = makeStruct(email)

	return
}

func pHash(pword string) (hash []byte) {
	// salt := []byte{118, 168, 47, 97, 146, 191, 30, 94, 167, 60, 50, 8, 191, 83, 179, 255, 216, 56, 220, 235, 139, 162, 140, 200, 91, 241, 88, 9, 98, 231, 9, 81}
	//_, err := io.ReadFull(rand.Reader, salt)
	//check(err)

	// hash, err := scrypt.Key([]byte(pword), salt, 1<<14, 8, 1, hashBytes)
	// check(err)

	sum := sha256.Sum256([]byte(pword))
	hash = sum[:]

	return
}

// Auth :
// A function that authenticates
// a user's password based on the username
// and password input.
func Auth(email, pword string) bool {
	user := makeStruct(email)
	if user.Email == "" {
		return false
	}

	password := pHash(pword)

	return compareSlice(user.pwd, password)
}

func compareSlice(a, b []byte) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func makeStruct(email string) (user User) {
	f, err := os.Open(dataFolder + email + ".txt")
	if err != nil {
		log.Println("Error making User struct:", err)
		return User{}
	}
	check(err)

	b := bufio.NewReader(f)

	name, err := b.ReadString('\n')
	check(err)

	dateStr, err := b.ReadString('\n')
	check(err)

	dateStr = dateStr[:len(dateStr)-1]

	date, err := time.Parse(time.RFC850, dateStr)
	check(err)

	pwd, _, err := b.ReadLine()
	check(err)

	user = User{email, name, pwd, date}

	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
