package lib

import (
	"monitor/conf"
	"monitor/core"
	"monitor/lib/dingtalk"
)

// 信息处理中心
func HandleIgnoreMsg(NowTime, msg, evName string, monitor conf.Monitor) {
	// 判读是否写入记录
	if conf.CFG.LOGFLAG && (len(evName) < len(conf.CFG.LOG) || conf.CFG.LOG != evName[len(evName)-len(conf.CFG.LOG):]) {
		_msg := "time： " + NowTime + " event： " + msg + "\n"
		core.WriteLog(_msg)
	}

	// 判断是否忽略该文件变动
	if monitor.IGNOREFLAG {
		if !core.StrContainOrInList(evName, monitor.IGNOREPATH) {
			HandleTypeMsg(NowTime, msg, evName, monitor)
		}
	} else {
		HandleTypeMsg(NowTime, msg, evName, monitor)
	}
}

// 判断是否处理该后缀文件
func HandleTypeMsg(NowTime, msg, evName string, monitor conf.Monitor) {
	if monitor.TYPEFLAG {
		if core.FileExtInList(evName, monitor.TYPE) {
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
