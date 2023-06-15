package opslevel

import (
	"encoding/json"
	"strconv"
)

// JSON is a specialized map[string]string to support proper graphql serialization
type JSON map[string]any

func (s JSON) GetGraphQLType() string { return "JSON" }

func NewJSON(data string) JSON {
	result := make(JSON)
	json.Unmarshal([]byte(data), &result)
	return result
}

func (s JSON) MarshalJSON() ([]byte, error) {
	dto := map[string]any{}
	for k, v := range s {
		dto[k] = v
	}
	b, err := json.Marshal(dto)
	return []byte(strconv.Quote(string(b))), err
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
