package rpcx

import (
	"context"
)

var _serv_desc *serverDesc

func init() {
	_serv_desc = &serverDesc{}
}

func ServerDesc() *serverDesc {
	return _serv_desc
}

type serverDesc struct {
	ServerID   int
	ServerName string
	ListenAddr string
}

// ServerDesc Param
type RequestDesc struct {
	ID int
}

type ResponseDesc struct {
	MsgCode int
}

// rpcx method

type ServerDescSync struct {
}

func (ServerDescSync) SyncID(ctx context.Context, req *RequestDesc, res *ResponseDesc) error {
	//ServerDesc().ServerID = req.ID
	//res.MsgCode = code.SUCCESS
	return nil
}
