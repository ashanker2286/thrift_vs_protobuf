package rpc

import (
	"testServer"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
        //"strconv"
)

func StartServer(handler *TESTHandler, port string) {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	//serverTransport, err := thrift.NewTServerSocket("localhost:" + strconv.Itoa(9000))
	serverTransport, err := thrift.NewTServerSocket("localhost:" + port)
	if err != nil {
		fmt.Println("StartServer: NewTServerSocket failed with error:", err)
		return
	}
	processor := testServer.NewTESTServicesProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	err = server.Serve()
	if err != nil {
		fmt.Println("Failed to start the listener, err:", err)
	}
	fmt.Println("Start the listener successfully")
	return
}
