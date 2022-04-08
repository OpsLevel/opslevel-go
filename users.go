package opslevel

import "github.com/shurcooL/graphql"

type MemberInput struct {
	Email string `json:"email"`
}

type User struct {
	Email   string
	HTMLUrl string
	Id      graphql.ID
	Name    string
	Role    UserRole
}

type UserConnection struct {
	Nodes    []User
	PageInfo PageInfo
}
