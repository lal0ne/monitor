package lib

func HandleIgnoreMsg(NowTime, msg, evName string) {
	if CFG.LOGFLAG {
		_msg := "time： " + NowTime + " event： " + msg + "\n"
		WriteLog(_msg)
	}

	if CFG.Monitor.IGNOREFLAG {
		if !StrContainOrInList(evName, CFG.Monitor.IGNOREPATH) {
			HandleTypeMsg(NowTime, msg, evName)
		}
	} else {
		HandleTypeMsg(NowTime, msg, evName)
	}
}

func HandleTypeMsg(NowTime, msg, evName string) {
	if CFG.Monitor.TYPEFLAG {
		if FileExtInList(evName, CFG.Monitor.TYPE) {
			SendMsgOrWriteLog(NowTime, msg)
		}
	} else {
		SendMsgOrWriteLog(NowTime, msg)
	}
}

func SendMsgOrWriteLog(NowTime, msg string) {
	if CFG.Dingtalk.FLAG {
		SendMsg(NowTime, msg)
	}
}
