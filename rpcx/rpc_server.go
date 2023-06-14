package rpcx

import (
	"github.com/slclub/go-tips"
	rpcxServer "github.com/smallnest/rpcx/server"
	"server/conf"
	"server/utils"
)

var rpcxServerObj = rpcxServer.NewServer()

func init() {
	Register("ServerDescSync", new(ServerDescSync))
}

func Start() {
	go rpcxServerObj.Serve("tcp", ":"+tips.StrEnd(conf.TomlConf.Server.RpcxAddr, ":"))
}

/**
 * 注册自定义的 对外类，类下挂载着方法，从rpcx的client端旧可以调用类
 */
func Register(controller string, obj interface{}) {
	if utils.IsNil(obj) {
		panic(any("[RPCX][REGISTER][PANIC]" + controller))
	}
	rpcxServerObj.RegisterName(controller, obj, "")
}
