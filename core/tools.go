package core

import (
	"os"
	"path"
	"regexp"
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
		r, _ := regexp.Compile(checkExt)
		if rawExt != "" && r.MatchString(rawExt[1:]) {
			return true
		}
	}

	return false
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
