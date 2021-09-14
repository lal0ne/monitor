package lib

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// 读取配置文件数据
func init() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	// 打开配置文件
	file, err := os.Open("conf.yaml")
	if err != nil {
		panic(err)
	}

	// 读取数据
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// 配置与结构体做关系映射
	err = yaml.Unmarshal(bytes, &CFG)
	if err != nil {
		panic(err)
	}

	ParamWringOrErr()
}
