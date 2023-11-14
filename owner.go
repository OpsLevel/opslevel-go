package opslevel

type EntityOwnerTeam struct {
	Alias string `json:"alias,omitempty" graphql:"teamAlias:alias"`
	Id    ID     `json:"id"`
}

type EntityOwner struct {
	OnTeam EntityOwnerTeam `graphql:"... on Team"`
}

func (s *EntityOwner) Alias() string {
	return s.OnTeam.Alias
}

func (s *EntityOwner) Id() ID {
	return s.OnTeam.Id
}

func (s *EntityOwnerTeam) AsTeam() TeamId {
	return TeamId{
		Alias: s.Alias,
		Id:    s.Id,
	}
}
