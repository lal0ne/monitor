package core

import (
	"bufio"
	"os"
	"path"
	"strings"

	"monitor/conf"
)

// 判读字符串是否包含列表中某个值
// @params:
//    rawStr: string
//    checkStrList: []string
// @return: bool
func StrContainOrInList(rawStr string, checkStrList []string) bool {
	for _, checkStr := range checkStrList {
		if strings.Contains(rawStr, checkStr) {
			return true
		}
	}

	return false
}

// 记录文件操作日志
func WriteLog(msg string) {
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
}

// 文件后缀名是否在列表中
func FileExtInList(rawName string, checkExtList []string) bool {
	rawExt := path.Ext(rawName)
	for _, checkExt := range checkExtList {
		if rawExt != "" && (rawExt == checkExt || rawExt[1:] == checkExt) {
			return true
		}
	}

	return false
}
