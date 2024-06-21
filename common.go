package opslevel

import (
	"fmt"
	"slices"
	"strings"
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

func RefOf[T any](v T) *T {
	return &v
}

func RefTo[T any](v T) *T {
	return &v
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

	var sb strings.Builder
	sb.WriteString("OpsLevel API Errors:\n")
	for _, err := range errs {
		if len(err.Path) == 1 && err.Path[0] == "base" {
			err.Path[0] = ""
		}
		sb.WriteString(fmt.Sprintf("\t- '%s' %s\n", strings.Join(err.Path, "."), err.Message))
	}

	return fmt.Errorf(sb.String())
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

// Returns copy of stringsToKeep after removing stringsToExclude
func getSliceWithStringsRemoved(stringsToKeep, stringsToExclude []string) []string {
	return slices.DeleteFunc(stringsToKeep, func(value string) bool {
		return slices.Contains(stringsToExclude, value)
	})
}
