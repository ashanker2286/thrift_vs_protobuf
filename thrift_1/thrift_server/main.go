package main

import (
	"fmt"
        "os"
	"thrift_vs_protobuf/thrift/thrift_server/rpc"
	"thrift_vs_protobuf/thrift/thrift_server/server"
)

func main() {
        argsCnt := len(os.Args[1:])
        if argsCnt != 1 {
                fmt.Println("Usage: ./server_thrift <server port number>")
                return
        }
        port := os.Args[1]
	fmt.Println("Starting test thrift server")
	fmt.Println("Start logger")

	fmt.Println("Starting test thrift server...")
	testServer := server.NewTESTServer()
	go testServer.StartServer()
        fmt.Println("Starting listener...")
        serverIface := rpc.NewTESTHandler(testServer)
        rpc.StartServer(serverIface, port)
}
