package opslevel

import (
	"slices"
	"time"

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
	omitDeactivedUsersFilter := UsersFilterInput{
		Key:  UsersFilterEnumDeactivatedAt,
		Type: &BasicTypeEnumEquals,
	}
	(*pv)["filter"] = &[]UsersFilterInput{omitDeactivedUsersFilter}
	return pv
}

func NewString(value string) *string {
	return &value
}

func NullString() *string {
	var output *string
	return output
}

func NullOf[T NullableConstraint]() *Nullable[T] {
	output := Nullable[T]{SetNull: true}
	return &output
}

func RefOf[T NullableConstraint](value T) *Nullable[T] {
	return NewNullableFrom(value)
}

func RefTo[T NullableConstraint](value T) *Nullable[T] {
	return NewNullableFrom(value)
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

type Connection interface {
	GetNodes() any
}

func MergeMaps(map1, map2 map[string]any) *map[string]any {
	merged := make(map[string]any)

	for key, value := range map1 {
		merged[key] = value
	}

	for key, value := range map2 {
		if _, present := merged[key]; !present {
			merged[key] = value
		}
	}

	return &merged
}
