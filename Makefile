proto:
	@protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. proto/s1.proto
build:
	@docker build -t shippy-service-consignment .
run:
	@docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 shippy-service-consignment

.PHONY: proto build run
