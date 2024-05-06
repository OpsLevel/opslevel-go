package opslevel

import (
	"encoding/base64"
	"encoding/json"
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

func (id ID) GetGraphQLType() string { return "ID" }

func (id *ID) MarshalJSON() ([]byte, error) {
	if *id == "" {
		return []byte("null"), nil
	}
	return []byte(strconv.Quote(string(*id))), nil
}

type Identifier struct {
	Id      ID       `graphql:"id"`
	Aliases []string `graphql:"aliases"`
}

func (identifierInput IdentifierInput) MarshalJSON() ([]byte, error) {
	if identifierInput.Id == nil && identifierInput.Alias == nil {
		return []byte("null"), nil
	}
	var out string
	if identifierInput.Id != nil {
		out = fmt.Sprintf(`{"id":"%s"}`, string(*identifierInput.Id))
	} else {
		out = fmt.Sprintf(`{"alias":"%s"}`, *identifierInput.Alias)
	}
	return []byte(out), nil
}

func NewIdentifier(value ...string) *IdentifierInput {
	if len(value) == 1 {
		if IsID(value[0]) {
			return &IdentifierInput{
				Id: NewID(value[0]),
			}
		}
		return &IdentifierInput{
			Alias: RefOf(value[0]),
		}
	}
	var output IdentifierInput
	return &output
}

func NewIdentifierArray(values []string) []IdentifierInput {
	output := make([]IdentifierInput, 0)
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

// NullableValue can be used to unset a value using an OpsLevel input struct type, should always be instantiated using a constructor.
type NullableValue[T any] struct {
	Value   T
	SetNull bool
}

func (nullableValue NullableValue[T]) MarshalJSON() ([]byte, error) {
	if nullableValue.SetNull {
		return []byte("null"), nil
	}
	return json.Marshal(nullableValue.Value)
}

func NewNullValue[T any]() *NullableValue[T] {
	return &NullableValue[T]{
		SetNull: true,
	}
}

func NewNullableValue[T any](value T, setNull bool) *NullableValue[T] {
	return &NullableValue[T]{
		Value:   value,
		SetNull: setNull,
	}
}
