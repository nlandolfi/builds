# needs: protoc see:https://github.com/google/protobuf/releases
# needs: go get -u github.com/golang/protobuf/protoc-gen-go
protoc --proto_path=$GOPATH/src/ \
    $GOPATH/src/github.com/nlandolfi/builds/infra/gub/gub.proto --go_out=plugins=grpc:$GOPATH/src/
