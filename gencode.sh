protoc -I . --go_out=plugins=grpc:. --grpc-gateway_out=. protos/v1/health/*.proto
protoc -I . --go_out=plugins=grpc:. --grpc-gateway_out=. protos/v1/oauth2/*.proto