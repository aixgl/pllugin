package login

import (
	"fmt"
	"os"
	"server/conf"
	"testing"
)

var root_path string

func testinit() {
	apath, _ := os.Getwd()
	root_path = apath + "/../../../../"
	fmt.Println(root_path)

	// 初始化配置
	conf.ReadTomlFrom(root_path+"conf/conf.toml", &conf.TomlConf)
}
func TestLoginInfo(t *testing.T) {
	testinit()
	t.Log("test request url:", LoginRoute("info"))
	resp_data := make(map[string]interface{})
	err := Get(LoginRoute("info"), &resp_data)
	t.Log(err, resp_data)
}
