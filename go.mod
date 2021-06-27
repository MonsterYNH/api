module github.com/MonsterYNH/api

go 1.16

require (
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.14.5
	google.golang.org/genproto v0.0.0-20200423170343-7949de9c1215
	google.golang.org/grpc v1.38.0
)

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5
	github.com/golang/protobuf => github.com/golang/protobuf v1.3.2
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)
