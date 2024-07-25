#!/bin/sh

VERSION=$(grep '^VERSION=' .version | cut -d '=' -f 2-)
echo VERSION=$VERSION

# if commit is not set, use the latest commit
if [ -z "$COMMIT" ]; then
  COMMIT=$(git rev-parse HEAD)
fi
echo COMMIT=$COMMIT

go build \
  -tags docker \
  -ldflags="-s -w \
    -X 'github.com/Lumerin-protocol/Morpheus-Lumerin-Node/proxy-router/internal/config.BuildVersion=$VERSION' \
    -X 'github.com/Lumerin-protocol/Morpheus-Lumerin-Node/proxy-router/internal/config.Commit=$COMMIT' \
  " \
  -o bin/proxy-router cmd/main.go
