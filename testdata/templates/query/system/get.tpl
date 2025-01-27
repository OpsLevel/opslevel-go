{{- define "system_get" -}}
{
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
  }
}
{{- end -}}