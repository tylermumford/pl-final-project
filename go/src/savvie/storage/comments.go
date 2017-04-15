package storage

import (
	"savvie/users"
	"time"
)

// Comment stores everything about an choice's comment. User contains a username
// (email address). Choice contains an Choice.ID. Body will be escaped, so HTML
// inside will not be rendered.
type Comment struct {
	User   string
	Choice string
	Date   time.Time
	Body   string
	Type   string // "Comment", "Pro", or "Con".
}

var commentList = []Comment{
	Comment{
		"johnsmith",
		"78912",
		time.Now(),
		"Come into work on our day off",
		"Con",
	},
	Comment{
		"tess-terson",
		"78912",
		time.Now(),
		"No reason, work or other, to miss the meeting",
		"Pro",
	},
	Comment{
		"mitchmasters",
		"89123",
		time.Now(),
		"less likely for us to be exhausted from the week.",
		"Pro",
	},
	Comment{
		"janitha",
		"89123",
		time.Now(),
		"Not everyone is at work by then!",
		"Con",
	},
	Comment{
		"anonymous",
		"89123",
		time.Now(),
		"Even if I'm at work, I may not be mentally acute yet",
		"Con",
	},
	Comment{
		"mysharona",
		"89123",
		time.Now(),
		"This is going to get in the way of productive work",
		"Con",
	},
	Comment{
		"kylejenkins",
		"67891",
		time.Now(),
		"Not everyone likes popcorn, or the same flavor",
		"Con",
	},
	Comment{
		"tylermumford",
		"67891",
		time.Now(),
		"popcorn + rootbeer is the tastiest!!",
		"Pro",
	},
	Comment{
		"katrinamehring",
		"67891",
		time.Now(),
		"But what about my figure????",
		"Con",
	},
	Comment{
		"leeeerooooyyy-jeeeeeeenkiiiins",
		"67891",
		time.Now(),
		"It'll make us feel parched the rest of the day",
		"Con",
	},
	Comment{
		"banananana",
		"45678",
		time.Now(),
		"What flavor of cookies makes a difference. Not oatmeal raisin!",
		"Comment",
	},
	Comment{
		"jessicajones",
		"45678",
		time.Now(),
		"Baked goods are healthier",
		"Pro",
	},
	Comment{
		"jonesy",
		"23456",
		time.Now(),
		"We're sitting all day, why would we want to sit for longer",
		"Con",
	},
	Comment{
		"shannon",
		"23456",
		time.Now(),
		"Safe from weather",
		"Pro",
	},
	Comment{
		"bonnie",
		"12345",
		time.Now(),
		"fresh air",
		"Pro",
	},
	Comment{
		"clyde",
		"12345",
		time.Now(),
		"It is springtime, might as well",
		"Comment",
	},
	Comment{
		"anderson",
		"12345",
		time.Now(),
		"Weather is hard to predict",
		"Con",
	},
	Comment{
		"safetydance",
		"12345",
		time.Now(),
		"It would be good to change it up",
		"Comment",
	},
	Comment{
		"stacysmom",
		"34567",
		time.Now(),
		"Wouldnt' be great for people that can't stand for long periods of time",
		"Con",
	},
	Comment{
		"santa",
		"34567",
		time.Now(),
		"Very distracting to see everyone moving and fidgetting",
		"Con",
	},
	Comment{
		"lucky",
		"34567",
		time.Now(),
		"It's similar to the choice to sit in a meeting room",
		"Comment",
	},
}

// NiceDate returns an output-ready, human-recognizable date string.
// The format of the date is not guaranteed to remain constant. It may change without notice.
func (c *Comment) NiceDate() string {
	return c.Date.Local().Format("Jan 2, 2006 at 15:04 MST")
}

// NiceName returns the human name of the comment's creator.
func (c *Comment) NiceName() string {
	return users.GetUser(c.User).Name
}

const commentsFolder = "/vagrant/data/comments/"

// LoadComments returns all of the comments on the given choice, but they're not guaranteed to be sorted.
// They will be sorted in a future version.
func LoadComments(optID string) []Comment {
	result := []Comment{}
	for c := range commentList {
		if commentList[c].Choice == optID {
			result = append(result, commentList[c])
		}
	}
	return result
}

// SaveNewComment persists a new comment with the given information.
func SaveNewComment(user, optID, body, whichType string) error {
	if user == "" || optID == "" || body == "" {
		return Error{"Could not create comment with given information."}
	}

	c := Comment{
		User:   user,
		Choice: optID,
		// TODO: set location when we set the time
		Date: time.Now(),
		Body: body,
		Type: whichType,
	}
	commentList = append(commentList, c)
	return nil
}

// Error provides information about what went wrong.
type Error struct {
	e string
}

func (e Error) Error() string {
	return "comments: " + e.e
}
