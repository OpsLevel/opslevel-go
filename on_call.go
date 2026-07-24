package opslevel

// OnCallEdge An edge in an on-call connection.
type OnCallEdge struct {
	Cursor string  `graphql:"cursor"`
	Node   *OnCall `graphql:"node"`
}
