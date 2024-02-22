#!/usr/bin/env bash
if [ ! -f ./graph/model/models_gen.go ]; then
  echo "Error: no generated models_gen.go found."
  exit 1
fi
cd graph/model || exit 1

# Get number of "records"
CODE_BLOCK_COUNT=$(awk -v RS=\} 'END{print FNR}' models_gen.go)

# Make sure objects.go is empty
truncate -s0 objects.go

echo 'package opslevel

import "github.com/relvacode/iso8601"' >> objects.go

echo "Extracting generated objects from models_gen.go to objects.go ..."
for BLOCK_NUM in $(seq 1 "$CODE_BLOCK_COUNT"); do
  CODE_BLOCK=$(awk -v RS=\} NR=="$BLOCK_NUM" models_gen.go)

  if echo "$CODE_BLOCK" | grep -e "type [A-Za-z0-9]* struct {" > /dev/null; then
    echo "$CODE_BLOCK" >> objects.go;
    echo -n "}" >> objects.go;
  fi
done

echo "Success!"
