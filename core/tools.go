package core

import (
	"path"
	"strings"
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
