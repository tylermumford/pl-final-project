package comments

import (
	"encoding/gob"
	"os"
	"time"
)

// Comment stores everything about an argument's comment. User contains a username
// (email address). Argument contains an argument.ID. Body will be escaped, so HTML
// inside will not be rendered.
type Comment struct {
	User     string
	Argument string
	Date     time.Time
	Body     string
}

const dataFolder = "/vagrant/data/comments/"

// Load returns all of the comments on the given argument, sorted somehow.
func Load(argID string) []Comment {
	filename := dataFolder + argID + "-comments.txt"
	file, err := os.Open(filename)
	if err != nil {
		return []Comment{}
	}
	defer file.Close()

	result := []Comment{}
	dec := gob.NewDecoder(file)
	dec.Decode(&result)
	return result
}

// Save stores a new comment with the given information.
func Save(user, argID, body string) error {
	if user == "" || argID == "" || body == "" {
		return Error{"Could not create comment with given information."}
	}

	all := Load(argID)
	filename := dataFolder + argID + "-comments.txt"
	file, err := os.Create(filename)
	if err != nil {
		return Error{"Could not open file: " + filename}
	}
	defer file.Close()

	c := Comment{
		User:     user,
		Argument: argID,
		Date:     time.Now(),
		Body:     body,
	}
	all = append(all, c)

	enc := gob.NewEncoder(file)
	enc.Encode(all)
	return nil
}

// Error provides information about what went wrong.
type Error struct {
	e string
}

func (e Error) Error() string {
	return "comments: " + e.e
}
