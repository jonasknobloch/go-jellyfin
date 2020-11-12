#!/bin/bash

rm -rf client
openapi-generator generate \
    -i ./spec/openapi.json \
    -g go --package-name "client" \
    --git-repo-id go-jellyfin/client \
    --git-user-id jonasknobloch \
    --git-host github.com \
    -o client \
    -p enumClassPrefix=true
