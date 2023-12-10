PHONY: run-server

gen:
	protoc --proto_path=proto proto/*.proto --go_out=server --go-grpc_out=server

clean:
	rm -rf server/pb/


install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
	brew install protobuf
	brew install clang-format
	brew install grpcurl
	export PATH=$PATH:$(go env GOPATH)/bin
