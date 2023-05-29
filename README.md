# gRPC Practice with Golang

## How to generate code

`protobuf` code is generated with this sentence

```bash
protoc --go-grpc_out=./<gen_directory_name> --go_out=./<gen_directory_name> <proto_file_name>
```

For example,

```bash
protoc --go-grpc_out=./chat --go_out=./chat chat.proto
```