package opslevel

import (
	"encoding/base64"
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
	Id    *ID     `graphql:"id" json:"id,omitempty" yaml:"id,omitempty"`
	Alias *string `graphql:"alias" json:"alias,omitempty" yaml:"alias,omitempty"`
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
