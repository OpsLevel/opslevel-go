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

type PropertyOwner struct {
	Typename   string `graphql:"__typename"`
	*TeamId    `graphql:"... on Team"`
	*ServiceId `graphql:"... on Service"`
}

func (o PropertyOwner) Id() ID {
	if o.ServiceId != nil {
		return o.ServiceId.Id
	}
	if o.TeamId != nil {
		return o.TeamId.Id
	}
	return ""
}
