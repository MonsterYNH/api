protoc -I . --go_out=plugins=grpc:. --grpc-gateway_out=. v1/health/*.proto
protoc -I . --go_out=plugins=grpc:. --grpc-gateway_out=. v1/oauth2/*.proto