// Package storage handles data persistence.
//
// It works with our C# code to store options (formerly arguments) and uses Go
// code to store comments.
//
// Options
//
// makeCmd is the backbone. All other functions use it to communicate with C#. Generally, you'll want to use GetOpt to get a specified option struct.
//
// Comments
//
// See the documentation of individual functions. Generally, you'll want to load a slice of all the comments for a specified option (LoadComments).
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
	optionsFolder = "/vagrant/data/storage/"
)

func init() {
	// Seed the random number generator.
	seconds := time.Now().Second()
	rand.Seed(int64(seconds))
}

// Option :
// contains information describing an option to a decision. IDs are 5-digit
// decimal numbers.
type Option struct {
	ID          string
	Description string
	Upvotes     int
	Downvotes   int
}

// ListOpts returns a list of all the options on the site.
func ListOpts() (optIDs []Option) {
	// sort by modification time. // parse out name w/o ".txt"
	data, _ := os.Open(optionsFolder)
	file, _ := data.Readdirnames(0)
	//info, _ := data.Readdir(-1) // ignoring this error
	//sort.Sort(info)

	for _, value := range file {
		// value contains "55555.txt", so we slice it to get just the number.
		optIDs = append(optIDs, GetOpt(value[:5]))
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

// SaveNewOption takes an option description and saves it as a new option.
// It returns the ID of the new option.
func SaveNewOption(descr string) string {
	fn := fmt.Sprintf("%0d", rand.Intn(99999))
	for GetOpt(fn).ID != "" {
		// Loop until we get a non-existing option ID
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

// GetOpt takes an option's 5-digit ID and returns the corresponding
// option struct.
func GetOpt(id string) Option {
	c := makeCmd(id, "export", "")
	str, _ := c.Output()
	parts := strings.Split(string(str), "@@@")

	// Should consist of description, upvotes, and downvotes.
	if len(parts) != 3 {
		return Option{}
	}

	u, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
	d, _ := strconv.Atoi(strings.TrimSpace(parts[2]))
	return Option{
		ID:          id,
		Description: parts[0],
		Upvotes:     u,
		Downvotes:   d,
	}
}

// Upvote simply upvotes the specified option.
func Upvote(id string) {
	c := makeCmd(id, "upvote", "")
	c.Run()
}

// Downvote simply downvotes the specified option.
func Downvote(id string) {
	c := makeCmd(id, "downvote", "")
	c.Run()
}

// Score returns upvotes minus downvotes.
func (a Option) Score() int {
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
