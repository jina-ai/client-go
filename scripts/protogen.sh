#/bin/sh

set -e

# This script is used to generate gRPC client stubs from the proto files.
JINA_VERSION=$1
DOCARRAY_VERSION=$2

GO_MODULE="github.com/jina-ai/client-go"
DOCARRAY_PROTO="docarray.proto"
DOCARRAY_DIR="../docarray"
DOCARRAY_PACKAGE="$GO_MODULE/docarray"

JINA_PROTO="jina.proto"
JINA_DIR="../jina"
JINA_PACKAGE="$GO_MODULE/jina"

cd protos

[-d $DOCARRAY_DIR/v$DOCARRAY_VERSION ] || mkdir $DOCARRAY_DIR/v$DOCARRAY_VERSION
mkdir $JINA_DIR/v$JINA_VERSION

grep -q '^option go_package = ' docarray.proto || sed -i '/package docarray;/aoption go_package = "'${DOCARRAY_PACKAGE}'";' docarray.proto
protoc --go_out=${DOCARRAY_DIR} \
       --go_opt=paths=source_relative \
       --go_opt=M${DOCARRAY_PROTO}=${DOCARRAY_PACKAGE} \
       --go-grpc_out=${DOCARRAY_DIR}/v${DOCARRAY_VERSION} \
       --go-grpc_opt=paths=source_relative \
       ${DOCARRAY_PROTO} 

cp $DOCARRAY_DIR/*.go $DOCARRAY_DIR/v$DOCARRAY_VERSION

grep -q '^option go_package = ' jina.proto || sed -i '/package jina;/aoption go_package = "'${JINA_PACKAGE}'";' jina.proto
protoc --go_out=${JINA_DIR} \
       --go_opt=paths=source_relative \
       --go_opt=M${JINA_PROTO}=${JINA_PACKAGE} \
       --go-grpc_out=${JINA_DIR} \
       --go-grpc_opt=paths=source_relative \
       ${JINA_PROTO} 

cp $JINA_DIR/*.go $JINA_DIR/v$JINA_VERSION

cd -
