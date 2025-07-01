#!/bin/sh

DOCS_FILE="./docs/docs.go"

# Detect OS and set sed in-place flag
if sed --version >/dev/null 2>&1; then
    # GNU sed (Linux)
    SED_INPLACE=(-i)
else
    # BSD sed (macOS)
    SED_INPLACE=(-i '')
fi

# Add os import
sed "${SED_INPLACE[@]}" 's#import "github.com/swaggo/swag"#import (\n\t"github.com/swaggo/swag"\n\t"os"\n)#g' "$DOCS_FILE"

# Set Host dynamically
sed "${SED_INPLACE[@]}" 's#Host:             ""#Host:             os.Getenv("HOST")#g' "$DOCS_FILE"

# Replace ApiKeyAuth with BearerAuth in security definitions
sed "${SED_INPLACE[@]}" 's#"ApiKeyAuth": \[\\\]#"BearerAuth": \[\\\]#g' "$DOCS_FILE"