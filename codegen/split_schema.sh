#!/usr/bin/env bash
if [ ! -f schema.graphql ]; then
  echo "No schema.graphql found. Getting new OpsLevel GraphQL schema."
  curl https://app.opslevel.com/public/schemas/schema.graphql > schema.graphql
fi

# Get number of "records"
CODE_BLOCK_COUNT=$(awk -v RS=\} 'END{print FNR}' schema.graphql)

# Make sure these files are empty
truncate -s0 \
  graph/enums.graphqls \
  graph/interfaces.graphqls \
  graph/inputs.graphqls \
  graph/objects.graphqls

echo "Splitting OpsLevel GraphQL API schema into subschemas..."
for BLOCK_NUM in $(seq 1 "$CODE_BLOCK_COUNT"); do
  CODE_BLOCK=$(awk -v RS=\} NR=="$BLOCK_NUM" schema.graphql)
  if echo "$CODE_BLOCK" | grep \
    -e "type Account {" \
    -e "type Query {" \
    -e "type Mutation {" > /dev/null; then
    continue
  fi

  if echo "$CODE_BLOCK" | grep -e "type [A-Za-z0-9]* {" -e "type [A-Za-z0-9]* implements [A-Za-z]* {" > /dev/null; then
    echo "$CODE_BLOCK" >> graph/objects.graphqls;
    echo -n "}" >> graph/objects.graphqls;
  elif echo "$CODE_BLOCK" | grep -e "interface [A-Za-z0-9]* {" > /dev/null; then
    echo "$CODE_BLOCK" >> graph/interfaces.graphqls;
    echo -n "}" >> graph/interfaces.graphqls;
  elif echo "$CODE_BLOCK" | grep -e "input [A-Za-z0-9]* {" > /dev/null; then
    echo "$CODE_BLOCK" >> graph/inputs.graphqls;
    echo -n "}" >> graph/inputs.graphqls;
  elif echo "$CODE_BLOCK" | grep -e "enum [A-Za-z0-9]* {" > /dev/null; then
    echo "$CODE_BLOCK" >> graph/enums.graphqls;
    echo -n "}" >> graph/enums.graphqls;
  fi
done

echo "Success! Listing subschemas for gqlgen code generation:"
ls -1 ./graph/*.graphqls
