package opslevel

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/hasura/go-graphql-client"
	"github.com/relvacode/iso8601"
)

type IdentifierInput struct {
	Id    ID             `graphql:"id" json:"id,omitempty"`
	Alias graphql.String `graphql:"alias" json:"alias,omitempty"`
}

type PageInfo struct {
	HasNextPage     graphql.Boolean `graphql:"hasNextPage"`
	HasPreviousPage graphql.Boolean `graphql:"hasPreviousPage"`
	Start           graphql.String  `graphql:"startCursor"`
	End             graphql.String  `graphql:"endCursor"`
}

type PayloadVariables map[string]interface{}

type DeleteInput struct {
	Id ID `json:"id"`
}

type OpsLevelErrors struct {
	Message string
	Path    []string
}

type IdResponsePayload struct {
	Id     ID `graphql:"deletedCheckId"`
	Errors []OpsLevelErrors
}

type ResourceDeletePayload struct {
	Alias  string           `graphql:"deletedAlias" json:"alias,omitempty"`
	Id     ID               `graphql:"deletedId" json:"id,omitempty"`
	Errors []OpsLevelErrors `graphql:"errors" json:"errors,omitempty"`
}

type Connection struct {
	PageInfo PageInfo `graphql:"pageInfo"`
}

type Timestamps struct {
	CreatedAt iso8601.Time `json:"createdAt"`
	UpdatedAt iso8601.Time `json:"updatedAt"`
}

func (p *IdResponsePayload) Mutate(client *Client, m interface{}, v PayloadVariables) error {
	if err := client.Mutate(m, v); err != nil {
		return err
	}
	return FormatErrors(p.Errors)
}

func IsID(value string) bool {
	decoded, err := base64.RawURLEncoding.DecodeString(value)
	if err != nil {
		return false
	}
	return strings.HasPrefix(string(decoded), "gid://")
}

func NewIdentifier(value string) *IdentifierInput {
	if IsID(value) {
		return &IdentifierInput{
			Id: ID(value),
		}
	}
	return &IdentifierInput{
		Alias: graphql.String(value),
	}
}

func NullString() *graphql.String {
	var output *graphql.String
	return output
}

func NewString(value string) *graphql.String {
	output := graphql.String(value)
	return &output
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool {
	p := new(bool)
	*p = v
	return p
}

func HandleErrors(err error, errs []OpsLevelErrors) error {
	if err != nil {
		return err
	}
	return FormatErrors(errs)
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

// NewId use "" to set "null" for ID input fields that can be nullified
func NewID(id string) *ID {
	var output ID
	if id != "" {
		output = ID(id)
	}
	return &output
}

func NewInt(i int) *int {
	output := i
	return &output
}

func NewISO8601Date(datetime string) iso8601.Time {
	date, _ := iso8601.ParseString(datetime)
	return iso8601.Time{Time: date}
}

func NewISO8601DateNow() iso8601.Time {
	return iso8601.Time{Time: time.Now()}
}

func removeDuplicates(data []string) []string {
	keys := make(map[string]bool)
	var list []string

	for _, entry := range data {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
