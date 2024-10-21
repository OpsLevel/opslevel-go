{{- define "service_get" -}}
{
    apiDocumentPath,
    description,
    framework,
    htmlUrl,
    id,
    aliases,
    language,
    lifecycle{
        alias,
        description,
        id,
        index,
        name
    },
    locked,
    managedAliases,
    name,
    note,
    owner{
        alias,
        id
    },
    parent{
        id,
        aliases
    },
    preferredApiDocument{
        id,
        htmlUrl,
        source{
            ... on ApiDocIntegration{
                id,
                name,
                type
            },
            ... on ServiceRepository{
                baseDirectory,
                displayName,
                id,
                repository{
                    id,
                    defaultAlias
                },
                service{
                    id,
                    aliases
                }
            }
        },
        timestamps{
            createdAt,
            updatedAt
        }
    },
    preferredApiDocumentSource,
    product,
    repos{
        edges{
            node{
                id,
                defaultAlias
            },
            serviceRepositories{
                baseDirectory,
                displayName,
                id,
                repository{
                    id,
                    defaultAlias
                },
                service{
                    id,
                    aliases
                }
            }
        },
        {{ template "pagination_request" }},
        totalCount
    },
    defaultServiceRepository{
        baseDirectory,
        displayName,
        id,
        repository{
            id,
            defaultAlias
        },
        service{
            id,
            aliases
        }
    },
    tags{
        nodes{
            id,
            key,
            value
        },
        {{ template "pagination_request" }},
        totalCount
    },
    tier{
        alias,
        description,
        id,
        index,
        name
    },
    timestamps{
        createdAt,
        updatedAt
    },
    tools{
        nodes{
            category,
            categoryAlias,
            displayName,
            environment,
            id,
            url,
            service{
                id,
                aliases
            }
        },
        {{ template "pagination_request" }},
        totalCount
    }
{{- end }}