package conf

import (
	"fmt"
	"os"
)

// 判断其中部分参数是否符合规则
func ParamWringOrErr() {
	// 如果启用钉钉，配置信息不能为空
	if CFG.Dingtalk.FLAG {
		if CFG.Dingtalk.API == "" || CFG.Dingtalk.KEY == "" {
			panic("钉钉信息配置错误")
		}
	}
	for _, monitorContent := range CFG.Monitor {
		if !DirExists(monitorContent.PATH) {
			panic(fmt.Sprintf("监控目录[%s]不存在", monitorContent.PATH))
		} else {
			if monitorContent.IGNOREFLAG {
				num := len(monitorContent.PATH)
				for _, path := range monitorContent.IGNOREPATH {
					if monitorContent.PATH != path[:num] {
						panic(fmt.Sprintf("[%s]不是[%s]的子目录", monitorContent.IGNOREPATH, monitorContent.PATH))
					}
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
