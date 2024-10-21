{{- define "system_get" -}}
{
  id,
  aliases,
  managedAliases,
  name,
  description,
  htmlUrl,
  owner{
    ... on Team{
      teamAlias:alias,
      id
    }
  },
  parent{
    id,
    aliases,
    description,
    htmlUrl,
    managedAliases,
    name,
    note,
    owner{
      ... on Team{
        teamAlias:alias,
        id
      }
    }
  },
  note
}
{{- end -}}