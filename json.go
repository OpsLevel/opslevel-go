package opslevel

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

type JSONString string

func NewJsonString(text string) (JSONString, error) {
	jsonString := JSONString(text)
	typeFound, err := jsonString.GetType()
	if err != nil {
		return "", err
	}
	if !json.Valid([]byte(text)) {
		return "", fmt.Errorf(
			"JSONString with type '%s' is not valid JSON. String provided: %s",
			typeFound,
			string(text),
		)
	}
	return jsonString, nil
}

func (s JSONString) GetType() (string, error) {
	typeFound := string(s)
	switch typeFound {
	case "true":
		return "bool", nil
	case "false":
		return "bool", nil
	case "null":
		return "null", nil
	}
	if strings.HasPrefix(typeFound, "{") || strings.HasSuffix(typeFound, "}") {
		return "object", nil
	}
	if strings.HasPrefix(typeFound, "[") || strings.HasSuffix(typeFound, "]") {
		return "array", nil
	}
	if strings.HasPrefix(typeFound, "\"") || strings.HasSuffix(typeFound, "\"") {
		return "string", nil
	}
	// REGEX:  starts with zero or one '-' for negative numbers, followed by
	// either: one or more digits, no other characters
	//     or: one or more digits, one dot, one or more digits, no other characters
	if ok, _ := regexp.MatchString(`^-?(\d+$|\d+\.?\d+)`, string(s)); ok {
		return "number", nil
	}
	return "", fmt.Errorf("unknown JSONString type: '%s'", string(s))
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
