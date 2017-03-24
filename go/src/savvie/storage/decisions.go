package storage

// Decision structs are the top-level things on the site
type Decision struct {
	ID          string
	Description string
	Options     []Option
}

var decisionList = []Decision{
	Decision{
		"692",
		"When to hold the meeting",
		[]Option{
			Option{
				"78912",
				"Hold the meeting on Saturday at 5PM",
				4,
				1,
			},
			Option{
				"89123",
				"Hold it Wednesday at 8AM to get it out of the way, while we're all ready for work",
				1,
				9,
			},
		},
	},
	Decision{
		"186",
		"What refreshments are wanted at the meeting",
		[]Option{
			Option{
				"67891",
				"Popcorn - Something that is still easy to talk while eating",
				2,
				6,
			},
			Option{
				"56789",
				"Apples and fruit dip - Something healthy",
				4,
				3,
			},
			Option{
				"45678",
				"Cookies - Not greasy, but easy to eat while talking",
				4,
				3,
			},
		},
	},
	Decision{
		"925",
		"Where the meeting should take place",
		[]Option{
			Option{
				"23456",
				"Official Conference Room - Will have a more formal atmosphere, safe from weather",
				3,
				2,
			},
			Option{
				"12345",
				"Park across the street - Weather is getting better and it will help us focus more when we come back",
				9,
				3,
			},
			Option{
				"34567",
				"Standing up, on the workfloor - Give us an opportunity to stretch our limbs, yet safe from bad weather",
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
	for _, d := range decisionList {
		if d.ID == decisionID {
			return d
		}
	}
	return Decision{}
}

// DecisionForOption returns the Decision associated with the given Option.
func DecisionForOption(opt Option) Decision {
	for decision := range decisionList {
		for option := range decisionList[decision].Options {
			if decisionList[decision].Options[option].ID == opt.ID {
				return decisionList[decision]
			}
		}
	}
	return Decision{}
}
