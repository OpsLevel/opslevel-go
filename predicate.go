package opslevel

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
)

type Predicate struct {
	Type  PredicateTypeEnum `graphql:"type"`
	Value string            `graphql:"value"`
}

var existsTypes = []PredicateTypeEnum{
	PredicateTypeEnumDoesNotExist,
	PredicateTypeEnumExists,
}

func (p *Predicate) Validate() error {
	if slices.Contains(existsTypes, p.Type) && p.Value != "" {
		return fmt.Errorf("Predicate type '%s' cannot have a value. Given value '%s'", p.Type, p.Value)
	} else if !slices.Contains(existsTypes, p.Type) && p.Value == "" {
		return fmt.Errorf("Predicate type '%s' requires a value", p.Type)
	}

	numericTypes := []PredicateTypeEnum{
		PredicateTypeEnumGreaterThanOrEqualTo,
		PredicateTypeEnumLessThanOrEqualTo,
	}
	if slices.Contains(numericTypes, p.Type) {
		if _, err := strconv.Atoi(p.Value); err != nil {
			return fmt.Errorf("FilterPredicate type '%s' requires a numeric value. Given '%s'", p.Type, p.Value)
		}
	}

	return nil
}

func (p *PredicateUpdateInput) MarshalJSON() ([]byte, error) {
	if p == nil || p.Type == nil || *p.Type == "" {
		return []byte("null"), nil
	}
	m := map[string]string{
		"type": string(*p.Type),
	}

	if p.Value != nil {
		m["value"] = *p.Value
	}

	return json.Marshal(m)
}
