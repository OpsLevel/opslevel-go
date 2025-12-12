{{- define "component_type_graphql" }}
{id,aliases,description,href,icon{color,name},isDefault,name,ownerRelationship{managementRules{operator,sourceProperty,sourcePropertyBuiltin,targetCategory,targetProperty,targetPropertyBuiltin,targetType}},timestamps{createdAt,updatedAt}}
{{end}}
{{- define "component_type_1_response" }}
{
    {{ template "id1" }},
    "aliases": [
      "example1"
    ],
    "name": "Example1",
    "description": "Description",
    "href": "https://app.opslevel-staging.com/catalog/domains/platformdomain",
    "icon": {
      "color": "#FFFFFF",
      "name": "PhBird"
    },
    "ownerRelationship": {
      "managementRules": [
        {
          "operator": "EQUALS",
          "sourceProperty": "tag_key_eq:owner",
          "sourcePropertyBuiltin": true,
          "targetCategory": null,
          "targetProperty": "name",
          "targetPropertyBuiltin": true,
          "targetType": "team"
        }
      ]
    }
}
{{end}}
{{- define "component_type_2_response" }}
{
    {{ template "id2" }},
    "aliases": [
      "example2"
    ],
    "name": "Example2",
    "description": "Description",
    "href": "https://app.opslevel-staging.com/catalog/domains/platformdomain",
    "icon": {
      "color": "#FFFFFF",
      "name": "PhBird"
    }
}
{{end}}
{{- define "component_type_3_response" }}
{
    {{ template "id3" }},
    "aliases": [
      "example3"
    ],
    "name": "Example3",
    "description": "Description",
    "href": "https://app.opslevel-staging.com/catalog/domains/platformdomain",
    "icon": {
      "color": "#FFFFFF",
      "name": "PhBird"
    }
}
{{end}}
