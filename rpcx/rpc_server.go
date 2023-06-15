package rpcx

import (
	"github.com/slclub/go-tips"
	rpcxServer "github.com/smallnest/rpcx/server"
)

var rpcxServerObj = rpcxServer.NewServer()

func init() {
	Register("ServerDescSync", new(ServerDescSync))
}

func Start(RpcxAddr string) {
	go rpcxServerObj.Serve("tcp", ":"+tips.StrEnd(RpcxAddr, ":"))
}

/**
 * 注册自定义的 对外类，类下挂载着方法，从rpcx的client端旧可以调用类
 */
func Register(controller string, obj interface{}) {
	if tips.IsNil(obj) {
		panic(any("[RPCX][REGISTER][PANIC]" + controller))
	}
	rpcxServerObj.RegisterName(controller, obj, "")
}
