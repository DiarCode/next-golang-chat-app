protos:
	protoc --go_out=. --go-grpc_out=. protos/auth/*.proto
	protoc --go_out=. --go-grpc_out=. protos/event/*.proto
	protoc --go_out=. --go-grpc_out=. protos/user/*.proto
	protoc --go_out=. --go-grpc_out=. protos/queue/*.proto