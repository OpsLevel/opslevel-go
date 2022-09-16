mutation ($input: LevelUpdateInput!) {
  levelUpdate(input: $input) {
    level {
      alias
      description
      id
      index
      name
    }
    errors {
      message
      path
    }
  }
}