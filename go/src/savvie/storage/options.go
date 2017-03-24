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
	for decision := range decisionList {
		for option := range decisionList[decision].Options {
			if decisionList[decision].Options[option].ID == id {
				return decisionList[decision].Options[option]
			}
		}
	}
	return Option{}
}

// Upvote simply upvotes the specified option.
func Upvote(id string) {
	for decision := range decisionList {
		for option := range decisionList[decision].Options {
			if decisionList[decision].Options[option].ID == id {
				decisionList[decision].Options[option].Upvotes++
			}
		}
	}
}

// Downvote simply downvotes the specified option.
func Downvote(id string) {
	for decision := range decisionList {
		for option := range decisionList[decision].Options {
			if decisionList[decision].Options[option].ID == id {
				decisionList[decision].Options[option].Downvotes++
			}
		}
	}
}

// Score returns upvotes minus downvotes.
func (a Option) Score() int {
	return a.Upvotes - a.Downvotes
}
