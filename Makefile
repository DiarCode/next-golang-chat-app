protos:
	protoc --go_out=. --go-grpc_out=. protos/auth/*.proto
	protoc --go_out=. --go-grpc_out=. protos/posts/*.proto
	protoc --go_out=. --go-grpc_out=. protos/user/*.proto