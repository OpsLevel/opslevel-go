package opslevel

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hasura/go-graphql-client"
)

type ErrorCode int

const (
	ErrorUnknown ErrorCode = iota
	ErrorRequestError
	ErrorAPIError
	ErrorResourceNotFound
)

type ClientError struct {
	error
	ErrorCode ErrorCode
}

func NewClientError(code ErrorCode, message string, args ...any) error {
	return &ClientError{
		error:     fmt.Errorf(message, args...),
		ErrorCode: code,
	}
}

func HandleErrors(opts ...any) error {
	output := (error)(nil)
	for _, opt := range opts {
		switch v := opt.(type) {
		case error:
			if !IsOpsLevelApiError(v) {
				output = errors.Join(output, NewClientError(ErrorRequestError, v.Error()))
			} else {
				output = errors.Join(output, v)
			}
		case []Error:
			output = errors.Join(output, HasAPIErrors(v))
		}
	}
	return output
}

func HasAPIErrors(errs []Error) error {
	if len(errs) == 0 {
		return nil
	}

	message := "OpsLevel API Errors:"
	for _, e := range errs {
		if len(e.Path) == 1 && e.Path[0] == "base" {
			e.Path[0] = ""
		}
		message += fmt.Sprintf("\n\t- '%s' %s", strings.Join(e.Path, "."), e.Message)
	}

	return NewClientError(ErrorAPIError, message)
}

func IsResourceFound(resource any) error {
	// TODO: Also Check if ID is valid somehow `.Id == ""`
	if resource == nil {
		return NewClientError(ErrorResourceNotFound, "resource '%T' not found", resource)
	}
	if casted, ok := resource.(Identifiable); ok && casted.GetID() == "" {
		return NewClientError(ErrorResourceNotFound, "resource '%T' not found", resource)
	}
	return nil
}

func ErrIs(err error, code ErrorCode) bool {
	var clientErr *ClientError
	if errors.As(err, &clientErr) {
		if clientErr.ErrorCode == code {
			return true
		}
	}
	return false
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
