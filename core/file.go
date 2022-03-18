package core

import (
	"bufio"
	"monitor/conf"
	"os"
)

// 记录文件操作日志
func WriteLog(msg string) {
	// 加入文件锁
	conf.RWfile.Lock()

	// 打开日志文件
	file, err := os.OpenFile(conf.CFG.LOG, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 写入日志文件
	write := bufio.NewWriter(file)
	write.WriteString(msg + "\n")
	write.Flush()

	// 关闭文件锁
	conf.RWfile.Unlock()
}
