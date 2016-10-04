package server

import (
        "fmt"
)

type TESTServer struct {
        ConfCh          chan bool
}

func NewTESTServer() *TESTServer {
        testServer := &TESTServer{}
        testServer.ConfCh = make(chan bool)
        return testServer
}

func (server *TESTServer) StartServer() {
        fmt.Println("Inside Start Server...")
        for {
                conf := <-server.ConfCh
                fmt.Println("Received a ping...", conf)
        }
}
