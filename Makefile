bank:
	protoc --go_opt=module=grpcbank --go_out=. ./proto/bank/*.proto

bank-grpc:
	protoc --go-grpc_opt=module=grpcbank --go-grpc_out=. ./proto/bank/*.proto