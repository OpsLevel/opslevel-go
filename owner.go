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
