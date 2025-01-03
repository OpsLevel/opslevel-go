package opslevel

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/hasura/go-graphql-client"
	"github.com/relvacode/iso8601"
)

type PageInfo struct {
	HasNextPage     bool   `graphql:"hasNextPage"`
	HasPreviousPage bool   `graphql:"hasPreviousPage"`
	Start           string `graphql:"startCursor"`
	End             string `graphql:"endCursor"`
}

type PayloadVariables map[string]interface{}

// WithoutDeactivedUsers filters out deactivated users on ListUsers query
func (pv *PayloadVariables) WithoutDeactivedUsers() *PayloadVariables {
	enumEquals := BasicTypeEnumEquals
	omitDeactivedUsersFilter := UsersFilterInput{
		Key:  UsersFilterEnumDeactivatedAt,
		Type: &enumEquals,
	}
	(*pv)["filter"] = &[]UsersFilterInput{omitDeactivedUsersFilter}
	return pv
}

type OpsLevelWarnings struct {
	Message string
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

func RefOf[T NullableConstraint](value T) *Nullable[T] {
	return NewNullableFrom(value)
}

func RefTo[T NullableConstraint](value T) *Nullable[T] {
	return NewNullableFrom(value)
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

	allErrors := fmt.Errorf("OpsLevel API Errors:")
	for _, err := range errs {
		if len(err.Path) == 1 && err.Path[0] == "base" {
			err.Path[0] = ""
		}
		newErr := fmt.Errorf("\t- '%s' %s", strings.Join(err.Path, "."), err.Message)
		allErrors = errors.Join(allErrors, newErr)
	}

	return allErrors
}

// IsOpsLevelApiError checks if the error is returned by OpsLevel's API
func IsOpsLevelApiError(err error) bool {
	if _, ok := err.(graphql.Errors); !ok {
		return false
	}
	for _, hasuraErr := range err.(graphql.Errors) {
		if len(hasuraErr.Path) > 0 {
			return true
		}
	}
	return false
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

// Given actual aliases and wanted aliases, returns aliasesToCreate and aliasesToDelete lists
func extractAliases(existingAliases, aliasesWanted []string) ([]string, []string) {
	var aliasesToCreate, aliasesToDelete []string

	// collect aliasesToDelete - existing aliases that are no longer wanted
	for _, alias := range existingAliases {
		if !slices.Contains(aliasesWanted, alias) {
			aliasesToDelete = append(aliasesToDelete, alias)
		}
	}

	// collect aliasesToCreate - wanted aliases that do not yet exist
	for _, aliasWanted := range aliasesWanted {
		if !slices.Contains(existingAliases, aliasWanted) {
			aliasesToCreate = append(aliasesToCreate, aliasWanted)
		}
	}
	return aliasesToCreate, aliasesToDelete
}
