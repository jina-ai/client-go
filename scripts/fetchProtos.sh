#!/bin/bash

# This script fetches the latest protobuf files from the Jina/Docarray repository
# and copies them to the `protos` directory.

# The script is meant to be run from the root of the repository.

JINA_VERSION=$1
DOCARRAY_VERSION=$2

if [ -z "$JINA_VERSION" ]; then
    echo "Please provide a Jina version as the first argument."
    exit 1
fi

if [ -z "$DOCARRAY_VERSION" ]; then
    echo "Please provide a Docarray version as the second argument."
    exit 1
fi

echo "Fetching protos for Jina version $JINA_VERSION"
wget https://raw.githubusercontent.com/jina-ai/jina/$JINA_VERSION/jina/proto/jina.proto -O protos/jina.proto

echo "Fetching protos for Docarray version $DOCARRAY_VERSION"
wget https://raw.githubusercontent.com/jina-ai/docarray/$DOCARRAY_VERSION/docarray/proto/docarray.proto -O protos/docarray.proto
