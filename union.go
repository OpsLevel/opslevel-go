// Code generated by gen.go; DO NOT EDIT.

package opslevel

// Approver The actor responsible for the approval/denial of an approvable resource.
type Approver struct {
	User UserId `graphql:"... on User"`
}

// CheckOwner represents the owner a check can belong to.
type CheckOwner struct {
	Team TeamId `graphql:"... on Team"`
	// User UserId `graphql:"... on User"` // TODO: will this be public?
}

// CodeIssueProjectResource represents resource linked to the CodeIssueProject. Can be either Service or Repository.
type CodeIssueProjectResource struct {
	Repository Repository `graphql:"... on Repository"`
	Service    Service    `graphql:"... on Service"`
}

// ContactOwner represents the owner of this contact.
type ContactOwner struct {
	Team TeamId `graphql:"... on Team"`
	User UserId `graphql:"... on User"`
}

// CustomActionsAssociatedObject represents the object that an event was triggered on.
type CustomActionsAssociatedObject struct {
	Service Service `graphql:"... on Service"`
}

// EntityOwner represents the Group or Team owning the entity.
type EntityOwner struct {
	OnTeam EntityOwnerTeam `graphql:"... on Team"`
}

// RelationshipResource represents a resource that can have relationships to other resources.
type RelationshipResource struct {
	Domain                 DomainId                 `graphql:"... on Domain"`
	InfrastructureResource InfrastructureResourceId `graphql:"... on InfrastructureResource"`
	Service                ServiceId                `graphql:"... on Service"`
	System                 SystemId                 `graphql:"... on System"`
	Team                   TeamId                   `graphql:"... on Team"`
}

// ServiceDocumentSource represents the source of a document.
type ServiceDocumentSource struct {
	IntegrationId     `graphql:"... on ApiDocIntegration"`
	ServiceRepository `graphql:"... on ServiceRepository"`
}

// TagOwner represents a resource that a tag can be applied to.
type TagOwner struct {
	Domain                 DomainId                 `graphql:"... on Domain"`
	InfrastructureResource InfrastructureResourceId `graphql:"... on InfrastructureResource"`
	Repository             Repository               `graphql:"... on Repository"`
	Service                Service                  `graphql:"... on Service"`
	System                 SystemId                 `graphql:"... on System"`
	Team                   TeamId                   `graphql:"... on Team"`
	User                   UserId                   `graphql:"... on User"`
}
