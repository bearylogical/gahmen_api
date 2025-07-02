#!/bin/sh

DOCS_FILE="./docs/docs.go"

# Cross-platform sed in-place function
sedi() {
  # $1 = sed script, $2 = file
  if sed --version >/dev/null 2>&1; then
    # GNU sed
    sed -i "$1" "$2"
  else
    # BSD sed (macOS)
    sed -i '' "$1" "$2"
  fi
}

# Add os import (use printf to insert literal newlines)
sedi "/import \"github.com\/swaggo\/swag\"/{
    s#import \"github.com/swaggo/swag\"#import (\\
    \t\"github.com/swaggo/swag\"\\
    \t\"os\"\\
    )#
}" "$DOCS_FILE"

# Set Host dynamically
sedi "s#Host:             \"\"#Host:             os.Getenv(\"HOST\")#g" "$DOCS_FILE"

sedi "s#Version:          \"1.0\"#Version:          os.Getenv(\"API_VERSION\")#g" "$DOCS_FILE"

# Replace ApiKeyAuth with BearerAuth in security definitions
sedi "s#\"ApiKeyAuth\": [\\]#\"BearerAuth\": [\\]#g" "$DOCS_FILE"