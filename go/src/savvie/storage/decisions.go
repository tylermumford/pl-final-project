package storage

// Decision structs are the top-level things on the site
type Decision struct {
	ID          string
	Description string
	Options     []Option
}

// ListDecisions returns a list of all the decisions on the site.
func ListDecisions() []Decision {
	// TODO
	return nil
}

// GetDecision returns a specific decision.
func GetDecision(decisionID string) Decision {
	// TODO
	return Decision{}
}
