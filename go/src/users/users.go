package users

import (
	"os"
	"bufio"
	"github.com/golang/crypto/scrypt"
	//"crypto/rand"
	//"io"
	//"fmt"
)

const(
	SALT_BYTES = 32
	HASH_BYTES = 64
)

//date signed up

const data_folder = "/vagrant/go/src/users/data/"

func NewUser(name, email, pword string) {
	//user{uname, uemail, upword}
	f, err := os.Create(data_folder+email+".txt")
	check(err)
	defer f.Close()

	_,err = f.WriteString(name+"\n")
	check(err)

	_,err = f.Write(pHash(pword))
	check(err)

	f.Sync()
}

func GetName(email string) (name string) {
	f, err := os.Open(data_folder+email+".txt")
	check(err)

	b := bufio.NewReader(f)
	name, err = b.ReadString('\n')
	check(err)

	name = name[:len(name)-1]

	return
}

func pHash(pword string) (hash []byte) {
	salt := []byte{118, 168, 47, 97, 146, 191, 30, 94, 167, 60, 50, 8, 191, 83, 179, 255, 216, 56, 220, 235, 139, 162, 140, 200, 91, 241, 88, 9, 98, 231, 9, 81}
	//_, err := io.ReadFull(rand.Reader, salt)
	//check(err)

	hash, err := scrypt.Key([]byte(pword), salt, 1<<14, 8, 1, HASH_BYTES)
	check(err)

	return
}

func Auth(email, pword string) bool {
	f, err := os.Open(data_folder+email+".txt")
	check(err)

	b := bufio.NewReader(f)

	_, err = b.ReadString('\n')
	check(err)

	pwd, _, err := b.ReadLine()
	check(err)
	password := pHash(pword)

	return  compareSlice(pwd, password)
}

func compareSlice(a, b [] byte) bool {
	if(a == nil && b == nil){
		return true
	}

	if(a == nil || b == nil){
		return false
	}

	if(len(a) != len(b)) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true;
}

func check(e error){
	if e != nil {
		panic(e)
	}
}