{{- define "infra_get" -}}
{
  id,
  aliases,
  name,
  type @include(if: $all),
  providerResourceType @include(if: $all),
  providerData @include(if: $all){
    accountName,
    externalUrl,
    providerName
  },
  owner @include(if: $all){
    ... on Team{
      teamAlias:alias,
      id
    }
  },
  ownerLocked @include(if: $all),
  data @include(if: $all),
  rawData @include(if: $all)
}
{{- end }}