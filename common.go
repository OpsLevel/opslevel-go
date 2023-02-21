package opslevel

import (
	"fmt"
	"strings"
	"time"

	"github.com/hasura/go-graphql-client"
	"github.com/relvacode/iso8601"
)

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

type Timestamps struct {
	CreatedAt iso8601.Time `json:"createdAt"`
	UpdatedAt iso8601.Time `json:"updatedAt"`
}

func NullString() *string {
	var output *string
	return output
}

// TODO: this can be replaced with NewString when its no longer a graphql.String type
func EmptyString() *string {
	output := ""
	return &output
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
