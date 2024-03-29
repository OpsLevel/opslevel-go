package opslevel

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type (
	// JSON represents a json object with keys and values for use with the OpsLevel API.
	// Instantiate using NewJSON.
	// Has a different graphql type compared to JSONSchema.
	JSON map[string]any

	// JSONSchema represents a json object with keys and values for use with the OpsLevel API.
	// Instantiate using NewJSONSchema.
	// Has a different graphql type compared to JSON.
	JSONSchema map[string]any

	// JsonString is a specialized input type to support serialization of any json compatible type
	// (bool, string, int, map, slice, etc.) for use with the OpsLevel API.
	// Instantiate using NewJSONInput.
	JsonString string
)

func (jsonSchema JSONSchema) GetGraphQLType() string { return "JSONSchema" }

func NewJSONSchema(data string) (*JSONSchema, error) {
	result := make(JSONSchema)
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// AsString returns a string containing its key value pairs marshalled as a json object.
func (jsonSchema JSONSchema) AsString() string {
	dto := map[string]any{}
	for k, v := range jsonSchema {
		dto[k] = v
	}
	b, _ := json.Marshal(dto)
	return string(b)
}

func (jsonSchema JSONSchema) MarshalJSON() ([]byte, error) {
	dto := map[string]any{}
	for k, v := range jsonSchema {
		dto[k] = v
	}
	b, err := json.Marshal(dto)
	return []byte(strconv.Quote(string(b))), err
}

func (jsonObject JSON) GetGraphQLType() string { return "JSON" }

func NewJSON(data string) (*JSON, error) {
	result := make(JSON)
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ToJSON returns a string containing its key value pairs marshalled as a json object.
func (jsonObject JSON) ToJSON() string {
	dto := map[string]any{}
	for k, v := range jsonObject {
		dto[k] = v
	}
	b, _ := json.Marshal(dto)
	return string(b)
}

func (jsonObject JSON) MarshalJSON() ([]byte, error) {
	dto := map[string]any{}
	for k, v := range jsonObject {
		dto[k] = v
	}
	b, err := json.Marshal(dto)
	return []byte(strconv.Quote(string(b))), err
}

func (jsonString JsonString) GetGraphQLType() string { return "JsonString" }

// NewJSONInput converts any json compatible type (bool, string, int, map, slice, etc.) into a valid JsonString.
// If passed a json object or array wrapped in a string, it will not use json.Marshal(data) and instead simply return
// the value of of JsonString(data) to prevent adding unnecessary escape characters.
func NewJSONInput(data any) (*JsonString, error) {
	if s, ok := data.(string); ok && wrappedObjectOrArray(s) {
		result := JsonString(s)
		return &result, nil
	}
	var result JsonString
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	result = JsonString(bytes)
	return &result, nil
}

func JsonStringAs[T any](data JsonString) (T, error) {
	var result T
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		return result, errors.Wrap(err, fmt.Sprintf("unable to marshal json as %T", result))
	}
	return result, nil
}

func (jsonString JsonString) AsBool() bool {
	value, _ := JsonStringAs[bool](jsonString)
	return value
}

func (jsonString JsonString) AsInt() int {
	value, _ := JsonStringAs[int](jsonString)
	return value
}

func (jsonString JsonString) AsFloat64() float64 {
	value, _ := JsonStringAs[float64](jsonString)
	return value
}

func (jsonString JsonString) AsString() string {
	value, _ := JsonStringAs[string](jsonString)
	return value
}

func (jsonString JsonString) AsArray() []any {
	value, _ := JsonStringAs[[]any](jsonString)
	return value
}

func (jsonString JsonString) AsMap() map[string]any {
	value, _ := JsonStringAs[map[string]any](jsonString)
	return value
}

func wrappedObjectOrArray(s string) bool {
	if (strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}")) || (strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]")) {
		return true
	}
	return false
}
