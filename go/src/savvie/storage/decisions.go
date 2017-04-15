package storage

// Decision structs are the top-level things on the site
type Decision struct {
	ID          string
	Description string
	Choices     []Choice
}

var decisionList = []Decision{
	Decision{
		"692",
		"When to hold the meeting",
		[]Choice{
			Choice{
				"78912",
				"Hold the meeting on Saturday at 5PM",
				4,
				1,
			},
			Choice{
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
		[]Choice{
			Choice{
				"67891",
				"Popcorn - Something that is still easy to talk while eating",
				2,
				6,
			},
			Choice{
				"56789",
				"Apples and fruit dip - Something healthy",
				4,
				3,
			},
			Choice{
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
		[]Choice{
			Choice{
				"23456",
				"Official Conference Room - Will have a more formal atmosphere, safe from weather",
				3,
				2,
			},
			Choice{
				"12345",
				"Park across the street - Weather is getting better and it will help us focus more when we come back",
				9,
				3,
			},
			Choice{
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

// DecisionForChoice returns the Decision associated with the given Choice.
func DecisionForChoice(opt Choice) Decision {
	for decision := range decisionList {
		for choice := range decisionList[decision].Choices {
			if decisionList[decision].Choices[choice].ID == opt.ID {
				return decisionList[decision]
			}
		}
	}
	return Decision{}
}
