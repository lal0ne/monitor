package dingtalk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"monitor/conf"
)

// 发送钉钉消息
func SendMsg(NowTime, msg string) {
	text := `{
		"msgtype": "text",
		"text": {
			"content": "%s"
		}
	}`
	msg = strings.ReplaceAll(msg, "\\", "/")
	_msg := conf.CFG.Dingtalk.KEY + "： 异常\n事件： " + msg + "\n时间： " + NowTime
	content := fmt.Sprintf(text, _msg)
	// fmt.Println(content)
	client := &http.Client{}
	req, err := http.NewRequest("POST", conf.CFG.Dingtalk.API, strings.NewReader(content))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resq, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resq.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
