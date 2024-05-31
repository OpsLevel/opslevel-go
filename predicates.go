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
	predicatesWithNoValue := []PredicateTypeEnum{
		PredicateTypeEnumDoesNotExist,
		PredicateTypeEnumExists,
	}
	if slices.Contains(predicatesWithNoValue, p.Type) && p.Value != "" {
		return fmt.Errorf("Predicate type '%s' cannot have a value. Given value '%s'", p.Type, p.Value)
	} else if p.Value == "" {
		return fmt.Errorf("Predicate type '%s' requires a value", p.Type)
	}
	return nil
}
