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
)

const (
	SALT_BYTES  = 32
	HASH_BYTES  = 64
	data_folder = "/vagrant/data/users/"
)

type User struct {
	Name, Email  string
	pwd          []byte
	Date_Started time.Time
}

func NewUser(name, email, pword string) {
	f, err := os.Create(data_folder + email + ".txt")
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

func GetUser(email string) (user User) {
	user = makeStruct(email)

	return
}

func pHash(pword string) (hash []byte) {
	// salt := []byte{118, 168, 47, 97, 146, 191, 30, 94, 167, 60, 50, 8, 191, 83, 179, 255, 216, 56, 220, 235, 139, 162, 140, 200, 91, 241, 88, 9, 98, 231, 9, 81}
	//_, err := io.ReadFull(rand.Reader, salt)
	//check(err)

	// hash, err := scrypt.Key([]byte(pword), salt, 1<<14, 8, 1, HASH_BYTES)
	// check(err)

	sum := sha256.Sum256([]byte(pword))
	hash = sum[:]

	return
}

func Auth(email, pword string) bool {
	user := makeStruct(email)

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
	f, err := os.Open(data_folder + email + ".txt")
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
