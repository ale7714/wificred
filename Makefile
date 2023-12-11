gen:
	protoc --proto_path=proto proto/*.proto --go_out=grpc-server --go-grpc_out=grpc-server
	protoc --proto_path=proto proto/*.proto --js_out=import_style=commonjs:client/pb \
      --grpc-web_out=import_style=commonjs,mode=grpcwebtext:client/pb

clean:
	rm -rf server/pb/
