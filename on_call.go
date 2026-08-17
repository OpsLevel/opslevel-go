package opslevel

// OnCallConnection The connection type for OnCall.
type OnCallConnection struct {
	Edges      []OnCallEdge `graphql:"edges"`
	Nodes      []OnCall     `graphql:"nodes"`
	PageInfo   PageInfo     `graphql:"pageInfo"`
	TotalCount int          `graphql:"totalCount"`
}

// OnCallEdge An edge in an on-call connection.
type OnCallEdge struct {
	Cursor string  `graphql:"cursor"`
	Node   *OnCall `graphql:"node"`
}
