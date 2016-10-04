#

.PHONY:	all clean

all:	client server

client: client.go protoc_ping
	go build -o client_proto client.go

server: server.go protoc_ping
	go build -o server_proto server.go

protoc_ping: ping.proto
	mkdir protoMsg
	protoc --go_out=protoMsg ping.proto
	@touch protoc_ping

clean:
	rm -rf client_proto
	rm -rf server_proto
	rm -rf protoMsg
	rm -rf protoc_ping
