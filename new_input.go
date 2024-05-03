package opslevel

// ServiceUpdater exists for backwards compatability between ServiceUpdateInput and ServiceUpdateInputV2
type ServiceUpdater interface {
	updatesService() bool // exists only to restrict which types qualify as a ServiceUpdater
}

func (inputType ServiceUpdateInput) updatesService() bool {
	return true
}

func (inputType ServiceUpdateInputV2) updatesService() bool {
	return true
}

// ServiceUpdateInputV2 enables setting fields like Framework and foreign keys like TierAlias to `null`
type ServiceUpdateInputV2 struct {
	Parent                *IdentifierInput `json:"parent,omitempty" yaml:"parent,omitempty"`                                               // The parent system for the service. (Optional.)
	Id                    *ID              `json:"id,omitempty" yaml:"id,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`             // The id of the service to be updated. (Optional.)
	Alias                 *OptionalString  `json:"alias,omitempty" yaml:"alias,omitempty" example:"example_alias"`                         // The alias of the service to be updated. (Optional.)
	Name                  *OptionalString  `json:"name,omitempty" yaml:"name,omitempty" example:"example_name"`                            // The display name of the service. (Optional.)
	Product               *OptionalString  `json:"product,omitempty" yaml:"product,omitempty" example:"example_product"`                   // A product is an application that your end user interacts with. Multiple services can work together to power a single product. (Optional.)
	Description           *OptionalString  `json:"description,omitempty" yaml:"description,omitempty" example:"example_description"`       // A brief description of the service. (Optional.)
	Language              *OptionalString  `json:"language,omitempty" yaml:"language,omitempty" example:"example_language"`                // The primary programming language that the service is written in. (Optional.)
	Framework             *OptionalString  `json:"framework,omitempty" yaml:"framework,omitempty" example:"example_framework"`             // The primary software development framework that the service uses. (Optional.)
	TierAlias             *OptionalString  `json:"tierAlias,omitempty" yaml:"tierAlias,omitempty" example:"example_alias"`                 // The software tier that the service belongs to. (Optional.)
	OwnerInput            *IdentifierInput `json:"ownerInput,omitempty" yaml:"ownerInput,omitempty"`                                       // The owner for the service. (Optional.)
	LifecycleAlias        *OptionalString  `json:"lifecycleAlias,omitempty" yaml:"lifecycleAlias,omitempty" example:"example_alias"`       // The lifecycle stage of the service. (Optional.)
	SkipAliasesValidation *bool            `json:"skipAliasesValidation,omitempty" yaml:"skipAliasesValidation,omitempty" example:"false"` // Allows updating a service with invalid aliases. (Optional.)
}

func (inputType ServiceUpdateInputV2) GetGraphQLType() string {
	return "ServiceUpdateInput"
}
