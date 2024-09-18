{{ define "integration_1" }}
{
    {{ template "id1" }},
    "name": "{{ template "name1" }}",
    "type": "deploy"
}
{{ end }}

{{ define "integration_2" }}
{
    {{ template "id2" }},
    "name": "{{ template "name2" }}",
    "type": "payload"
}
{{ end }}

{{ define "integration_3" }}
{
    {{ template "id3" }},
    "name": "{{ template "name3" }}",
    "type": "generic"
}
{{ end }}

{{- define "deploy_integration_response" }}
  "id": "Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkRlcGxveUludGVncmF0aW9uLzMyNw",
  "name": "Deploy",
  "type": "deploy"
{{ end }}

{{- define "payload_integration_response" }}
  "id": "Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OlBheWxvYWRJbnRlZ3JhdGlvbi8zNDg",
  "name": "Payload",
  "type": "payload"
{{ end }}
{{ define "kubernetes_integration_response" }}
  "id": "Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkdlbmVyaWNJbnRlZ3JhdGlvbi81ODQ",
  "name": "Kubernetes",
  "type": "generic"
{{ end }}

{{- define "integration_request" -}}
{id,name,type,createdAt,installedAt,... on ApiDocIntegration{},... on ArgocdDeployIntegration{},... on AwsIntegration{iamRole,externalId,awsTagsOverrideOwnership,ownershipTagKeys},... on AzureResourcesIntegration{aliases,ownershipTagKeys,subscriptionId,tagsOverrideOwnership,tenantId},... on CircleciDeployIntegration{},... on DeployIntegration{},... on FluxIntegration{},... on GenericIntegration{},... on GitLabCIntegration{},... on GithubActionsIntegration{},... on GoogleCloudIntegration{aliases,clientEmail,ownershipTagKeys,projects{id,name,url},tagsOverrideOwnership},... on NewRelicIntegration{baseUrl,accountKey},... on OctopusDeployIntegration{}}
{{- end }}