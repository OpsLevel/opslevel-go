package opslevel

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type Identifiable interface {
	GetID() ID
}

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
	Id      ID       `graphql:"id" json:"id"`
	Aliases []string `graphql:"aliases" json:"aliases"`
}

func (s Identifier) GetID() ID {
	return s.Id
}

func (identifierInput IdentifierInput) MarshalJSON() ([]byte, error) {
	if identifierInput.Id == nil && identifierInput.Alias == nil {
		return []byte("null"), nil
	}
	var out string
	if identifierInput.Id != nil {
		out = fmt.Sprintf(`{"id":"%s"}`, *identifierInput.Id)
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
			Alias: &value[0],
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

// NullableConstraint defines what types can be nullable - keep separated using the union operator (pipe)
type NullableConstraint interface {
	any
}

// Nullable can be used to unset a value using an OpsLevel input struct type, should always be instantiated using a constructor.
type Nullable[T NullableConstraint] struct {
	Value   T
	SetNull bool
}

func (nullable Nullable[T]) MarshalJSON() ([]byte, error) {
	if nullable.SetNull {
		return []byte("null"), nil
	}
	return json.Marshal(nullable.Value)
}

func (nullable *Nullable[T]) UnmarshalJSON(data []byte) error {
	stuff := json.Unmarshal(data, &nullable.Value)
	return stuff
}

func (nullable *Nullable[T]) MarshalYAML() (interface{}, error) {
	if nullable == nil || nullable.SetNull {
		return nil, nil
	}
	return nullable.Value, nil
}

// UnmarshalYAML implements the yaml.Unmarshaler interface
func (nullable *Nullable[T]) UnmarshalYAML(value *yaml.Node) error {
	// Handle null values
	if value.Kind == yaml.ScalarNode && value.Tag == "!!null" {
		*nullable = Nullable[T]{
			SetNull: true,
		}
		return nil
	}

	// Handle direct value unmarshaling
	if value.Kind == yaml.ScalarNode {
		var v T
		if err := value.Decode(&v); err != nil {
			return err
		}
		*nullable = Nullable[T]{
			Value:   v,
			SetNull: false,
		}
		return nil
	}

	// Handle structured format
	if value.Kind == yaml.MappingNode {
		for i := 0; i < len(value.Content); i += 2 {
			key := value.Content[i]
			val := value.Content[i+1]

			switch key.Value {
			case "value":
				var v T
				if err := val.Decode(&v); err != nil {
					return err
				}
				nullable.Value = v
				nullable.SetNull = false
			case "setNull":
				var setNull bool
				if err := val.Decode(&setNull); err != nil {
					return err
				}
				nullable.SetNull = setNull
			}
		}
		return nil
	}

	return fmt.Errorf("invalid YAML node kind: %v", value.Kind)
}

// NewNull returns a Nullable string that will always marshal into `null`, can be used to unset fields
func NewNull[T string]() *Nullable[T] {
	return NewNullOf[T]()
}

// NewNullOf returns a Nullable of any type that fits NullableConstraint that will always marshal into `null`, can be used to unset fields
func NewNullOf[T NullableConstraint]() *Nullable[T] {
	return &Nullable[T]{
		SetNull: true,
	}
}

// NewNullableFrom returns a Nullable that will never marshal into `null`, can be used to change fields or even set them to an empty value (like "")
func NewNullableFrom[T NullableConstraint](value T) *Nullable[T] {
	return &Nullable[T]{
		Value:   value,
		SetNull: false,
	}
}
