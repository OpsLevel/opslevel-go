package opslevel

const teamQuery = `
query($teamAlias: String) {
  account {
    team(alias: $teamAlias){
      id
      name
      responsibilities
      manager {
        name
        email
      }
      contacts {
        displayName
        address
      }
    }
  }
}
`

const tagCreateMutation = `
mutation create($input: TagCreateInput!){
  tagCreate(input: $input){
    tag{
      key
      value
    }
    errors {
      path
      message
    }
  }
}
`
