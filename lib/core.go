package lib

import (
	"monitor/conf"
	"monitor/core"
	"monitor/lib/dingtalk"
)

// 信息处理中心
func HandleIgnoreMsg(NowTime, msg, evName string) {
	// 判读是否写入记录
	if conf.CFG.LOGFLAG {
		_msg := "time： " + NowTime + " event： " + msg + "\n"
		core.WriteLog(_msg)
	}

	// 判断是否忽略该文件变动
	if conf.CFG.Monitor.IGNOREFLAG {
		if !core.StrContainOrInList(evName, conf.CFG.Monitor.IGNOREPATH) {
			HandleTypeMsg(NowTime, msg, evName)
		}
	} else {
		HandleTypeMsg(NowTime, msg, evName)
	}
}

// 判断是否处理该后缀文件
func HandleTypeMsg(NowTime, msg, evName string) {
	if conf.CFG.Monitor.TYPEFLAG {
		if core.FileExtInList(evName, conf.CFG.Monitor.TYPE) {
			SendMsgOrWriteLog(NowTime, msg)
		}
	} else {
		SendMsgOrWriteLog(NowTime, msg)
	}
}

// 判断是否发送钉钉消息
func SendMsgOrWriteLog(NowTime, msg string) {
	if conf.CFG.Dingtalk.FLAG {
		dingtalk.SendMsg(NowTime, msg)
	}
}
