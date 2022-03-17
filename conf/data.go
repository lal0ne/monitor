package conf

// 配置文件对应结构体
type Cfg struct {
	LOGFLAG  bool
	LOG      string
	Dingtalk Dingtalk
	Monitor  []Monitor
}

type Dingtalk struct {
	FLAG bool
	API  string
	KEY  string
}

type Monitor struct {
	PATH       string
	IGNOREFLAG bool
	IGNOREPATH []string
	TYPEFLAG   bool
	TYPE       []string
}
