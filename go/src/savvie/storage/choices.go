// Package storage handles data persistence.
//
// It works with our C# code to store choices (formerly arguments) and uses Go
// code to store comments.
//
// Choices
//
// makeCmd is the backbone. All other functions use it to communicate with C#. Generally, you'll want to use GetChoice to get a specified choice struct.
//
// Comments
//
// See the documentation of individual functions. Generally, you'll want to load a slice of all the comments for a specified choice (LoadComments).
package storage

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const (
	choicesFolder = "/vagrant/data/storage/"
)

func init() {
	// Seed the random number generator.
	seconds := time.Now().Second()
	rand.Seed(int64(seconds))
}

// Choice :
// contains information describing an choice to a decision. IDs are 5-digit
// decimal numbers.
type Choice struct {
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

// SaveNewChoice takes a choice description and saves it as a new choice.
// It returns the ID of the new choice.
func SaveNewChoice(descr string) string {
	fn := fmt.Sprintf("%0d", rand.Intn(99999))
	for GetChoice(fn).ID != "" {
		// Loop until we get a non-existing choice ID
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

// GetChoice takes a choice's 5-digit ID and returns the corresponding
// choice struct.
func GetChoice(id string) Choice {
	for decision := range decisionList {
		for choice := range decisionList[decision].Choices {
			if decisionList[decision].Choices[choice].ID == id {
				return decisionList[decision].Choices[choice]
			}
		}
	}
	return Choice{}
}

// Upvote simply upvotes the specified choice.
func Upvote(id string) {
	for decision := range decisionList {
		for choice := range decisionList[decision].Choices {
			if decisionList[decision].Choices[choice].ID == id {
				decisionList[decision].Choices[choice].Upvotes++
			}
		}
	}
}

// Downvote simply downvotes the specified choice.
func Downvote(id string) {
	for decision := range decisionList {
		for choice := range decisionList[decision].Choices {
			if decisionList[decision].Choices[choice].ID == id {
				decisionList[decision].Choices[choice].Downvotes++
			}
		}
	}
}

// Score returns upvotes minus downvotes.
func (a Choice) Score() int {
	return a.Upvotes - a.Downvotes
}
