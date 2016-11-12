package main

import (
	"fmt"
	"os"
	"os/exec"
)

type argument struct {
	id          string
	description string
	upvotes     int
	downvotes   int
}

func makeCmd(filename string, sCmd string, descr ...string) exec.Cmd {
	rgs := []string{filename, sCmd}

	for _, data := range descr {
		rgs = append(rgs, data)
	}

	result := exec.Cmd{
		Path:   "/vagrant/bin/storage",
		Args:   rgs,
		Stderr: os.Stderr,
	}
	return result
}

func saveNewArgument(descr string) {
	// Don't include a file extension.
	// fn := fmt.Sprintf("%v", rand.Int())
	cmd := makeCmd("file", "create", descr)

	result1, err := cmd.Output()
	if err != nil {
		fmt.Printf("There's an error: %v\n\n", err.Error())
		// return ""
	} else {
		fmt.Printf("We have some results: %v\n\n", string(result1))
		// return fn
	}
	// TODO: Return filename so someone can redirect to it.
}

func setDescription() {

}
