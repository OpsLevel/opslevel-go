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
	id    *ID     `graphql:"id" json:"id,omitempty" yaml:"id,omitempty"`
	alias *string `graphql:"alias" json:"alias,omitempty" yaml:"alias,omitempty"`
}

func (i *IdentifierInput) ID() *ID {
	return i.id
}

func (i *IdentifierInput) Alias() *string {
	return i.alias
}

func NewIdentifier(value ...string) *IdentifierInput {
	var output *IdentifierInput
	if len(value) == 1 {
		if IsID(value[0]) {
			return &IdentifierInput{
				id: NewID(value[0]),
			}
		}
		return &IdentifierInput{
			alias: NewString(value[0]),
		}
	}
	return output
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
