package opslevel

type EntityOwnerTeam struct {
	Alias string `json:"alias,omitempty" graphql:"teamAlias:alias"`
	Id    ID     `json:"id"`
}

func (entityOwner *EntityOwner) Alias() string {
	return entityOwner.OnTeam.Alias
}

func (entityOwner *EntityOwner) Id() ID {
	return entityOwner.OnTeam.Id
}

func (entityOwnerTeam *EntityOwnerTeam) AsTeam() TeamId {
	return TeamId(*entityOwnerTeam)
}

type EntityOwnerService struct {
	OnService ServiceId `graphql:"... on Service"`
}

func (entityOwnerService *EntityOwnerService) Aliases() []string {
	return entityOwnerService.OnService.Aliases
}

func (entityOwnerService *EntityOwnerService) Id() ID {
	return entityOwnerService.OnService.Id
}

// HasPropertiesOwner represents the owner of a Property (Service or Team).
// GraphQL type: HasProperties (interface implemented by Service and Team).
type HasPropertiesOwner struct {
	OnService ServiceId       `graphql:"... on Service"`
	OnTeam    EntityOwnerTeam `graphql:"... on Team"`
}

func (o HasPropertiesOwner) Id() ID {
	if o.OnService.Id != "" {
		return o.OnService.Id
	}
	return o.OnTeam.Id
}

func (o HasPropertiesOwner) Alias() string {
	if o.OnService.Id != "" {
		if len(o.OnService.Aliases) > 0 {
			return o.OnService.Aliases[0]
		}
		return ""
	}
	return o.OnTeam.Alias
}
