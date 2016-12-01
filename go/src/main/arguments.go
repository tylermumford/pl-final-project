/*
*	The code in this file is meant to work with
*		our C# files to store arguments and comments.
*
*	makeCmd is the backbone. All other functions use
*		it to communicate with C#.
 */

package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type argument struct {
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

func saveNewArgument(descr string) string {
	// Don't include a file extension.
	fn := fmt.Sprintf("%0d", rand.Intn(99999))
	cmd := makeCmd(fn, "create", descr)

	_, err := cmd.Output()
	if err != nil {
		panic(err.Error())
	} else {
		return fn
	}
}

func getArg(id string) argument {
	c := makeCmd(id, "export", "")
	str, _ := c.Output()
	parts := strings.Split(string(str), "@@@")

	// Should consist of description, upvotes, and downvotes.
	if len(parts) != 3 {
		return argument{}
	}

	u, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
	d, _ := strconv.Atoi(strings.TrimSpace(parts[2]))
	return argument{
		ID:          id,
		Description: parts[0],
		Upvotes:     u,
		Downvotes:   d,
	}
}

func upvote(id string) {
	c := makeCmd(id, "upvote", "")
	c.Run()
}

func downvote(id string) {
	c := makeCmd(id, "downvote", "")
	c.Run()
}

func (a argument) Score() int {
	return a.Upvotes - a.Downvotes
}
