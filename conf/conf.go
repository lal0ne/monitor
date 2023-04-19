package conf

import (
	"fmt"
	"github.com/spf13/viper"
	//"gopkg.in/yaml.v2"
	//"io/ioutil"
	"os"
)

// 读取配置文件数据
func init() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	/*
		yaml版本
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
	*/

	// 获取项目路径
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// 读取配置
	vip := viper.New()
	vip.AddConfigPath(path)
	vip.SetConfigName("conf")
	vip.SetConfigType("yaml")
	// 读取配置
	if err := vip.ReadInConfig(); err != nil {
		panic(err)
	}
	err = vip.Unmarshal(&CFG)
	if err != nil {
		panic(err)
	}

	// 处理参数是否符合
	ParamWringOrErr()
}
