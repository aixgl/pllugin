package login

import (
	"github.com/slclub/utils/bytesconv"
	"server/conf"
	"strings"
	//github.com/slclub/utils/bytesconv
)

/**
 * 	resp_data := make(map[string]interface{}) // resp_data 也可以是结构体
 *	err := Get(LoginRoute("info"), &resp_data)
 */

// 根据路由生成login 登陆地址 可以直接用于http.Get http.Post
// 也可以使用 我们封装的Get Post等函数获取
func LoginRoute(req_path string) string {
	//req_path = url.QueryEscape(req_path)
	buf := bytesconv.StringToBytes(req_path)
	if len(buf) <= 1 {
		return LoginDomain()
	}
	if buf[0] == '/' {
		buf = buf[1:]
	}
	return LoginDomain() + "/" + bytesconv.BytesToString(buf)
}

func LoginDomain() string {
	return strings.Join([]string{conf.TomlConf.Login.Scheme, conf.TomlConf.Login.LoginAddr}, "")
}
