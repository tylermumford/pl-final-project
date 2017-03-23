package storage

// Decision structs are the top-level things on the site
type Decision struct {
	ID          string
	Description string
	Options     []Option
}

var decisionList = []Decision{
	Decision{
		"When to hold the meeting",
		"Ideas and thoughts of when the meeting should be held",
		[]Option{
			Option{
				"Saturday at 5PM",
				"Hold the meeting on Saturday at 5PM",
				4,
				1,
			},
			Option{
				"Wednesday at 8AM",
				"Hold it early to get it out of the way, while we're all ready for work",
				1,
				9,
			},
		},
	},
	Decision{
		"Meeting Refreshments",
		"What refreshments are wanted at the meeting",
		[]Option{
			Option{
				"Popcorn",
				"Something that is still easy to talk while eating",
				2,
				6,
			},
			Option{
				"Apples and fruit dip",
				"Something healthy",
				4,
				3,
			},
			Option{
				"Cookies",
				"Not greasy, but easy to eat while talking",
				4,
				3,
			},
		},
	},
	Decision{
		"Meeting Location",
		"Where the meeting should take place",
		[]Option{
			Option{
				"Official Conference Room",
				"Will have a more formal atmosphere, safe from weather",
				3,
				2,
			},
			Option{
				"Park across the street",
				"Weather is getting better and it will help us focus more when we come back",
				9,
				3,
			},
			Option{
				"Standing up, on the workfloor",
				"Give us an opportunity to stretch our limbs, yet safe from bad weather",
				1,
				4,
			},
		},
	},
}

// ListDecisions returns a list of all the decisions on the site.
func ListDecisions() []Decision {
	return decisionList
}

// GetDecision returns a specific decision.
func GetDecision(decisionID string) Decision {
	// TODO
	return Decision{}
}
