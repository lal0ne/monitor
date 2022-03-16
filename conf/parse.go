package conf

import (
	// "fmt"
	"os"
)

// 判断其中部分参数是否符合规则
func ParamWringOrErr() {
	if CFG.Dingtalk.FLAG {
		if CFG.Dingtalk.API == "" || CFG.Dingtalk.KEY == "" {
			panic("钉钉信息配置错误")
		}
	}
	if !DirExists(CFG.Monitor.PATH) {
		panic("监控目录不存在")
	} else {
		if CFG.Monitor.IGNOREFLAG {
			num := len(CFG.Monitor.PATH)
			for _, path := range CFG.Monitor.IGNOREPATH {
				if CFG.Monitor.PATH != path[:num] {
					panic("ignorepath不是监控目录的子目录")
				}
			}
		}
	}
}

// 判断路径是否存在
func DirExists(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}
	return f.IsDir()
}
