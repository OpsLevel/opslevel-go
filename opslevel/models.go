package opslevel

import "github.com/machinebox/graphql"

type Client struct {
	url string
	bearerToken string
	graphqlClient *graphql.Client
}

type Team struct {
	Id               string
	Name             string
	Responsibilities string
	Manager          User
	Contacts         []Contact
}

type User struct {
	Name string
	Email string
}

type Contact struct {
	DisplayName string
	Address string
}

type Tag struct {
	Id string
	Owner string
	Key string
	Value string
}

type graphqlError struct {
	Path []string
	Message string
}

type createTagResponse struct {
	TagCreate struct {
		Tag *Tag
		Errors []graphqlError
	}
}

type teamResponse struct {
	Account struct {
		Team *Team
	}
}

