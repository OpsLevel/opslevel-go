package opslevel

import (
	"fmt"
	"strings"
)

type PageInfo struct {
	HasNextPage     bool   `graphql:"hasNextPage"`
	HasPreviousPage bool   `graphql:"hasPreviousPage"`
	Start           string `graphql:"startCursor"`
	End             string `graphql:"endCursor"`
}

type PayloadVariables map[string]interface{}

type OpsLevelErrors struct {
	Message string
	Path    []string
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
