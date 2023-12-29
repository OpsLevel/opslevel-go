package opslevel

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

type ID string

func NewID(id ...string) *ID {
	var output ID
	if len(id) == 1 {
		output = ID(id[0])
	}
	return &output
}

func (s ID) GetGraphQLType() string { return "ID" }

func (s *ID) MarshalJSON() ([]byte, error) {
	if *s == "" {
		return []byte("null"), nil
	}
	return []byte(strconv.Quote(string(*s))), nil
}

type Identifier struct {
	Id      ID       `graphql:"id"`
	Aliases []string `graphql:"aliases"`
}

type IdentifierInput struct {
	Id    *ID     `json:"id,omitempty" yaml:"id,omitempty"`
	Alias *string `json:"alias,omitempty" yaml:"alias,omitempty"`
}

func (i IdentifierInput) MarshalJSON() ([]byte, error) {
	// this case can only occur by using an EmptyIdentifier()
	if i.Id == nil && i.Alias == nil {
		return []byte("null"), nil
	}

	var out string
	if i.Id != nil {
		out = fmt.Sprintf(`{"id":"%s"}`, string(*i.Id))
	} else {
		out = fmt.Sprintf(`{"alias":"%s"}`, string(*i.Alias))
	}
	return []byte(out), nil
}

func NewIdentifier(value string) *IdentifierInput {
	if IsID(value) {
		return &IdentifierInput{
			Id: NewID(value),
		}
	}
	return &IdentifierInput{
		Alias: NewString(value),
	}
}

func EmptyIdentifier() *IdentifierInput {
	var output IdentifierInput
	return &output
}

func NewIdentifierArray(values []string) []IdentifierInput {
	output := []IdentifierInput{}
	for _, value := range values {
		output = append(output, *NewIdentifier(value))
	}
	return output
}

func IsID(value string) bool {
	decoded, err := base64.RawURLEncoding.DecodeString(value)
	if err != nil {
		return false
	}
	return strings.HasPrefix(string(decoded), "gid://")
}
