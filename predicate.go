package opslevel

import (
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
	var predicateAsJson string
	if p.Value == nil {
		predicateAsJson = fmt.Sprintf(`{"type":"%s"}`, *p.Type)
	} else {
		predicateAsJson = fmt.Sprintf(`{"type":"%s","value":"%s"}`, *p.Type, *p.Value)
	}
	return []byte(predicateAsJson), nil
}