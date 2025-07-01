#!/bin/bash

DOCS_FILE="/Users/syamil/Projects/gahmen_api/docs/docs.go"

# Add os import
sed -i '' 's#import "github.com/swaggo/swag"#import (\n\t"github.com/swaggo/swag"\n\t"os"\n)#g' $DOCS_FILE

# Set Host dynamically
sed -i '' 's#Host:             ""#Host:             os.Getenv("HOST")#g' $DOCS_FILE

# Replace ApiKeyAuth with BearerAuth in security definitions
sed -i '' 's#"ApiKeyAuth": \[\\\]#"BearerAuth": \[\\\]#g' $DOCS_FILE
