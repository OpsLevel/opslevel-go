package opslevel

type EntityOwner struct {
	OnGroup GroupId `graphql:"... on Group"`
	OnTeam  TeamId  `graphql:"... on Team"`
}

func (s *EntityOwner) Id() ID {
	if s.OnGroup.Id == "" {
		return s.OnTeam.Id
	} else {
		return s.OnGroup.Id
	}
}
