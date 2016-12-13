// Package storage handles data persistence.
//
// It works with our C# code to store arguments and uses Go code to store comments.
//
// Arguments
//
// makeCmd is the backbone. All other functions use it to communicate with C#. Generally, you'll want to use GetArg to get a specified Argument struct.
//
// Comments
//
// See the documentation of individual functions. Generally, you'll want to load a slice of all the comments for a specified argument (LoadComments).
package storage

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	argumentsFolder = "/vagrant/data/storage/"
)

func init() {
	// Seed the random number generator.
	seconds := time.Now().Second()
	rand.Seed(int64(seconds))
}

// Argument :
// contains information describing an argument. IDs are 5-digit
// decimal numbers.
type Argument struct {
	ID          string
	Description string
	Upvotes     int
	Downvotes   int
}

// ListArgs returns a list of all the arguments on the site.
func ListArgs() (argIDs []Argument) {
	// sort by modification time. // parse out name w/o ".txt"
	data, _ := os.Open(argumentsFolder)
	file, _ := data.Readdirnames(0)
	//info, _ := data.Readdir(-1) // ignoring this error
	//sort.Sort(info)

	for _, value := range file {
		// value contains "55555.txt", so we slice it to get just the number.
		argIDs = append(argIDs, GetArg(value[:5]))
	}

	return
}

func makeCmd(filename string, sCmd string, descr string) exec.Cmd {
	rgs := []string{"blank", filename, sCmd, descr}

	result := exec.Cmd{
		Path:   "/vagrant/bin/storage",
		Args:   rgs,
		Stderr: os.Stderr,
	}
	return result
}

// SaveNewArgument takes an argument description and saves it as a new argument.
// It returns the ID of the new argument.
func SaveNewArgument(descr string) string {
	fn := fmt.Sprintf("%0d", rand.Intn(99999))
	for GetArg(fn).ID != "" {
		// Loop until we get a non-existing argument ID
		fn = fmt.Sprintf("%0d", rand.Intn(99999))
	}
	cmd := makeCmd(fn, "create", descr)

	_, err := cmd.Output()
	if err != nil {
		panic(err.Error())
	} else {
		return fn
	}
}

// GetArg takes an argument's 5-digit ID and returns the corresponding
// argument struct.
func GetArg(id string) Argument {
	c := makeCmd(id, "export", "")
	str, _ := c.Output()
	parts := strings.Split(string(str), "@@@")

	// Should consist of description, upvotes, and downvotes.
	if len(parts) != 3 {
		return Argument{}
	}

	u, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
	d, _ := strconv.Atoi(strings.TrimSpace(parts[2]))
	return Argument{
		ID:          id,
		Description: parts[0],
		Upvotes:     u,
		Downvotes:   d,
	}
}

// Upvote simply upvotes the specified argument.
func Upvote(id string) {
	c := makeCmd(id, "upvote", "")
	c.Run()
}

// Downvote simply downvotes the specified argument.
func Downvote(id string) {
	c := makeCmd(id, "downvote", "")
	c.Run()
}

// Score returns upvotes minus downvotes.
func (a Argument) Score() int {
	return a.Upvotes - a.Downvotes
}

/*
func (list *[]os.FileInfo) Len() int {
	return len(list)
}
func (list *[]FileInfo) Less(i, j int) bool {
	return list[i].ModTime < list[j].ModTime
}
func (list *[]FileInfo) Swap(i, j int) {
	tmp := list[i]
	list[i] = list[j]
	list[j] = tmp
}
*/
