#

.PHONY:	all clean

all:	client server

client: client.go protoc_helloworld
	go build -o client_proto client.go

server: server.go protoc_helloworld
	go build -o server_proto server.go

protoc_helloworld: helloworld.proto
	mkdir protoMsg
	protoc --go_out=protoMsg helloworld.proto
	@touch protoc_helloworld

clean:
	rm -rf client_proto
	rm -rf server_proto
	rm -rf protoMsg
	rm -rf protoc_helloworld
