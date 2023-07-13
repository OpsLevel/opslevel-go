package opslevel

type EntityOwnerGroup struct {
	Alias string `json:"alias,omitempty" graphql:"groupAlias:alias"`
	Id    ID     `json:"id"`
}

type EntityOwnerTeam struct {
	Alias string `json:"alias,omitempty" graphql:"teamAlias:alias"`
	Id    ID     `json:"id"`
}

type EntityOwner struct {
	OnGroup EntityOwnerGroup `graphql:"... on Group"`
	OnTeam  EntityOwnerTeam  `graphql:"... on Team"`
}

func (s *EntityOwner) Alias() string {
	if s.OnGroup.Id == "" {
		return s.OnTeam.Alias
	} else {
		return s.OnGroup.Alias
	}
}

func (s *EntityOwner) Id() ID {
	if s.OnGroup.Id == "" {
		return s.OnTeam.Id
	} else {
		return s.OnGroup.Id
	}
}

func (s *EntityOwnerGroup) AsGroup() GroupId {
	return GroupId{
		Alias: s.Alias,
		Id:    s.Id,
	}
}

func (s *EntityOwnerTeam) AsTeam() TeamId {
	return TeamId{
		Alias: s.Alias,
		Id:    s.Id,
	}
}
