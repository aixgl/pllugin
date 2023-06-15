package login

import (
	"github.com/slclub/go-tips/stringbyte"
	"strings"
)

// 数据中心
func DataSrvRoute(req_path string) string {
	buf := stringbyte.StringToBytes(req_path)
	if len(buf) <= 1 {
		return DataSrvDomain()
	}
	if buf[0] == '/' {
		buf = buf[1:]
	}
	return DataSrvDomain() + "/" + stringbyte.BytesToString(buf)
}

func DataSrvDomain() string {
	return strings.Join([]string{Database.Scheme, Database.Addr}, "")
}
