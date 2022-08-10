scopes:
  - ./rpc
protoc: v3.20.1
protocWorkDir: ""
plugins:
  protoc-gen-go: google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.0
  protoc-gen-go-grpc: google.golang.org/grpc/cmd/protoc-gen-go-grpc@ad51f572fd270f2323e3aa2c1d2775cab9087af2
  protoc-gen-grpc-gateway: github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.11.0
  protoc-gen-openapiv2: github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.11.0
#  protoc-gen-oaago: github.com/oaago/protoc-gen-oaago@v0.0.5
  protoc-gen-govalidators: github.com/mwitkow/go-proto-validators/protoc-gen-govalidators@v0.3.2
  protoc-gen-doc: github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@v1.5.1
  protoc-go-inject-tag: github.com/favadi/protoc-go-inject-tag@v1.3.0
repositories:
  GOGO_PROTOBUF: https://github.com/gogo/protobuf@226206f39bd7276e88ec684ea0028c18ec2c91ae
  GOOGLE_APIS: https://github.com/googleapis/googleapis@75e9812478607db997376ccea247dd6928f70f45
options:
  - --go_out=./internal/api
  - --go-grpc_out=./internal/api
  - --go-grpc_opt=paths=source_relative
  - --oaago_opt=paths=import
  - --oaago_out=./internal/api
  - --go_opt=paths=source_relative
  - --grpc-gateway_out=./internal/api
  - --grpc-gateway_opt=paths=source_relative
  - --grpc-gateway_opt=generate_unbound_methods=true
  - --grpc-gateway_opt=logtostderr=true
  - --grpc-gateway_opt=register_func_suffix=GW
  - --grpc-gateway_opt=allow_delete_body=true
  - --govalidators_out=paths=source_relative:./internal/api
  - --doc_out=./docs
  - --openapiv2_out=./docs
importPaths:
  - .
  - $GOPATH
  - $GOPATH/src
  - $POWERPROTO_INCLUDE
  - $SOURCE_RELATIVE
  - $GOOGLE_APIS/github.com/googleapis/googleapis
  - $GOGO_PROTOBUF
  - $GOPATH/pkg/mod
  - ./rpc
postActions:
  - name: replace
    args:
      - ./apis/**/*.go
      - ",omitempty"
      - ""
#             protoc -I ./ -I ./rpc \
#             --proto_path=$GOPATH/src \
#             --proto_path=${GOPATH}/pkg/mod \
#             --govalidators_out=paths=source_relative:./internal/api \
#             --go_out=paths=source_relative:./internal/api \
#             --go-grpc_out=./internal/api --go-grpc_opt=paths=import \
#             --oaago_out=./internal/api \
#             --oaago_opt=paths=source_relative \
#             --grpc-gateway_out ./internal/api --grpc-gateway_opt paths=source_relative \
#             --grpc-gateway_opt logtostderr=true \
#             --grpc-gateway_opt generate_unbound_methods=true \
#             --grpc-gateway_opt register_func_suffix=gateway \
#             --grpc-gateway_opt allow_delete_body=true \
#             --doc_out=./docs \
#             --doc_opt=html,index.html \
#             --openapiv2_out ./docs --openapiv2_opt logtostderr=true \
#             ./rpc/ccc/ddd/ccc_ddd.proto

