package opslevel

import (
	"fmt"
	"strings"

	"github.com/shurcooL/graphql"
)

type IdentifierInput struct {
	Id    graphql.ID     `graphql:"id,omitempty" json:"id,omitempty"`
	Alias graphql.String `graphql:"alias,omitempty" json:"alias,omitempty"`
}

type PageInfo struct {
	HasNextPage     graphql.Boolean `graphql:"hasNextPage"`
	HasPreviousPage graphql.Boolean `graphql:"hasPreviousPage"`
	Start           graphql.String  `graphql:"startCursor"`
	End             graphql.String  `graphql:"endCursor"`
}

type PayloadVariables map[string]interface{}

type OpsLevelErrors struct {
	Message string
	Path    []string
}

type IdResponsePayload struct {
	Id     graphql.ID `graphql:"deletedCheckId"`
	Errors []OpsLevelErrors
}

func (p *IdResponsePayload) Mutate(client *Client, m interface{}, v PayloadVariables) error {
	if err := client.Mutate(m, v); err != nil {
		return err
	}
	return FormatErrors(p.Errors)
}

func NewId(id string) *IdentifierInput {
	return &IdentifierInput{
		Id: graphql.ID(id),
	}
}

func NewIdFromAlias(alias string) *IdentifierInput {
	return &IdentifierInput{
		Alias: graphql.String(alias),
	}
}

func FormatErrors(errs []OpsLevelErrors) error {
	if len(errs) == 0 {
		return nil
	}

	var errstrings []string
	errstrings = append(errstrings, "OpsLevel API Errors:")
	for _, err := range errs {
		errstrings = append(errstrings, fmt.Sprintf("\t* %s", string(err.Message)))
	}

	return fmt.Errorf(strings.Join(errstrings, "\n"))
}

func removeDuplicates(data []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range data {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
