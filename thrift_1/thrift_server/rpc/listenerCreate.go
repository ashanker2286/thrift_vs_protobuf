package rpc

import (
        //"fmt"
        //"testServer"
        //"thrift_vs_protobuf/thrift/thrift_server/server"
)

func (h *TESTHandler) SendPing() {
        h.server.ConfCh <-true
        return
}

func (h *TESTHandler) Ping(str string) (string, error) {
        //fmt.Println("Received Msg")
        //h.SendPing()
        return "Echo Reply", nil
}
