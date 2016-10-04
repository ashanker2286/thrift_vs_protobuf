package rpc

import (
	"thrift_vs_protobuf/thrift/thrift_server/server"
)

type TESTHandler struct {
	server *server.TESTServer
}

func NewTESTHandler(server *server.TESTServer) *TESTHandler {
	h := new(TESTHandler)
	h.server = server
	return h
}
