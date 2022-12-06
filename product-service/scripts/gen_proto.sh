# #!/bin/bash
# CURRENT_DIR=$(pwd)

# rm -rf ./genproto/*

# for module in $(find $CURRENT_DIR/protos/* -type d); do
#     protoc -I=${module} -I $CURRENT_DIR/protos/ \
#            --gofast_out=plugins=grpc:$CURRENT_DIR/ \
#             $module/*.proto;
# done;

#!/bin/bash
CURRENT_DIR=$1
for x in $(find ${CURRENT_DIR}/protos/* -type d); do
  sudo protoc --plugin="protoc-gen-go=${GOPATH}/bin/protoc-gen-go" --plugin="protoc-gen-go-grpc=${GOPATH}/bin/protoc-gen-go-grpc" -I=${x} -I=${CURRENT_DIR}/protos -I /usr/local/include --go_out=${CURRENT_DIR} \
   --go-grpc_out=${CURRENT_DIR} ${x}/*.proto
done

