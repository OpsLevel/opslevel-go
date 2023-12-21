package opslevel

import (
	"encoding/json"
	"strconv"

	"github.com/rs/zerolog/log"
)

// JSON is a specialized map[string]string to support proper graphql serialization
type JSON map[string]any

func (s JSON) GetGraphQLType() string { return "JSON" }

func NewJSON(data string) JSON {
	result := make(JSON)
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		panic(err)
	}
	return result
}

func (s JSON) ToJSON() string {
	dto := map[string]any{}
	for k, v := range s {
		dto[k] = v
	}
	b, _ := json.Marshal(dto)
	return string(b)
}

func (s JSON) MarshalJSON() ([]byte, error) {
	dto := map[string]any{}
	for k, v := range s {
		dto[k] = v
	}
	b, err := json.Marshal(dto)
	return []byte(strconv.Quote(string(b))), err
}

// JSONString is a specialized input type to support serialization to JSON for input to graphql
type JSONString string

func (s JSONString) GetGraphQLType() string { return "JSONSchema" }

func NewJSONInput(data any) JSONString {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return JSONString(bytes)
}

func JsonStringAs[T any](data JSONString) (T, error) {
	var result T
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		log.Warn().Err(err).Msgf("unable to marshal json as %T", result)
		return result, err
	}
	return result, nil
}

func (s JSONString) AsBool() bool {
	value, _ := JsonStringAs[bool](s)
	return value
}

func (s JSONString) AsInt() int {
	value, _ := JsonStringAs[int](s)
	return value
}

func (s JSONString) AsFloat64() float64 {
	value, _ := JsonStringAs[float64](s)
	return value
}

func (s JSONString) AsString() string {
	value, _ := JsonStringAs[string](s)
	return value
}

func (s JSONString) AsArray() []any {
	value, _ := JsonStringAs[[]any](s)
	return value
}

func (s JSONString) AsMap() map[string]any {
	value, _ := JsonStringAs[map[string]any](s)
	return value
}

//
//func (s *JSON) UnmarshalJSON(data []byte) error {
//	escaped, err := strconv.Unquote(string(data))
//	if err != nil {
//		return err
//	}
//	dto := map[string]string{}
//	if err := json.Unmarshal([]byte(escaped), &dto); err != nil {
//		return err
//	}
//	if (*s) == nil {
//		(*s) = JSON{}
//	}
//	for k, v := range dto {
//		(*s)[k] = v
//	}
//	return nil
//}
