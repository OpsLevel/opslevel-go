package opslevel

import (
	"strconv"
)

type ID string

func (s ID) GetGraphQLType() string { return "ID" }

func (s *ID) MarshalJSON() ([]byte, error) {
	if s == nil {
		return []byte("\"\""), nil
	}
	if *s == "" {
		return []byte("null"), nil
	}
	return []byte(strconv.Quote(string(*s))), nil
}
