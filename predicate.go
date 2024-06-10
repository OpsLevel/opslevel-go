package opslevel

import (
	"encoding/json"
	"fmt"
	"slices"
)

type Predicate struct {
	Type  PredicateTypeEnum `graphql:"type"`
	Value string            `graphql:"value"`
}

func (p *Predicate) Validate() error {
	if !slices.Contains(AllPredicateTypeEnum, string(p.Type)) {
		return fmt.Errorf("Invalidate Predicate type '%s'. Expected one of '%v'", p.Type, AllPredicateTypeEnum)
	}

	predicatesWithNoValue := []PredicateTypeEnum{
		PredicateTypeEnumDoesNotExist,
		PredicateTypeEnumExists,
	}
	if slices.Contains(predicatesWithNoValue, p.Type) && p.Value != "" {
		return fmt.Errorf("Predicate type '%s' cannot have a value. Given value '%s'", p.Type, p.Value)
	} else if !slices.Contains(predicatesWithNoValue, p.Type) && p.Value == "" {
		return fmt.Errorf("Predicate type '%s' requires a value", p.Type)
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
