package opslevel

// OnCall A user that is currently on call for a service.
type OnCall struct {
	ExternalEmail string  `graphql:"externalEmail"`
	Id            ID      `graphql:"id"`
	Name          string  `graphql:"name"`
	User          *UserId `graphql:"user"`
}

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
