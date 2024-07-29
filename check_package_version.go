package opslevel

type PackageVersionCheckFragment struct {
	MissingPackageResult       *CheckResultStatusEnum `graphql:"missingPackageResult"`       // The check result if the package isn't being used by a service.
	PackageConstraint          PackageConstraintEnum  `graphql:"packageConstraint"`          // The package constraint the service is to be checked for.
	PackageManager             PackageManagerEnum     `graphql:"packageManager"`             // The package manager (ecosystem) this package relates to.
	PackageName                string                 `graphql:"packageName"`                // The name of the package to be checked.
	PackageNameIsRegex         bool                   `graphql:"packageNameIsRegex"`         // Whether or not the value in the package name field is a regular expression.
	VersionConstraintPredicate *Predicate             `graphql:"versionConstraintPredicate"` // The predicate that describes the version constraint the package must satisfy.
}

// CreateCheckPackageVersion Creates a package version check.
func (client *Client) CreateCheckPackageVersion(input CheckPackageVersionCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkPackageVersionCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckPackageVersionCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

// UpdateCheckPackageVersion Updates a package version check.
func (client *Client) UpdateCheckPackageVersion(input CheckPackageVersionUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkPackageVersionUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckPackageVersionUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}
