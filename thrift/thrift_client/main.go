package main

import (
        "os"
        "time"
        "utils/ipcutils"
        "fmt"
        "../../../../../generated/src/testServer"
        "strconv"
        "git.apache.org/thrift.git/lib/go/thrift"
)

type TestClientBase struct {
        Address         string
        Transport       thrift.TTransport
        PtrProtocolFactory      *thrift.TBinaryProtocolFactory
}

type TestClient struct {
        TestClientBase
        ClientHdl *testServer.TESTServicesClient
}

var testClient TestClient

func connect_to_testServer(port string) {
        //testClient.Address = "localhost:" + strconv.Itoa(9000)
        testClient.Address = "localhost:" + port
        testClient.Transport, testClient.PtrProtocolFactory, _ = ipcutils.CreateIPCHandles(testClient.Address)
        testClient.ClientHdl = testServer.NewTESTServicesClientFactory(testClient.Transport, testClient.PtrProtocolFactory)
}

func sendPing(msg string) {
        //fmt.Println("Sending Msg")
        testClient.ClientHdl.Ping(msg)
        //fmt.Println("Received Msg")
}

func main() {
        argsCnt := len(os.Args[1:])
        if argsCnt != 3 {
                fmt.Println("Usage: ./client_thrift <server port number> <num of iterations> <num of bytes>")
                return
        }
        port := os.Args[1]
        num, _ := strconv.Atoi(os.Args[2])
        numOfBytes, _ := strconv.Atoi(os.Args[3])
        var msg string
        for i := 0; i < numOfBytes; i++ {
                msg = msg + "A"
        }
        //fmt.Println(msg)
        connect_to_testServer(port)
        t := time.Now()
        for i := 0; i < num; i++ {
                sendPing(msg)
        }
        fmt.Println(time.Since(t))
}
