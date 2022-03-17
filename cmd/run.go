package cmd

import (
	"github.com/fsnotify/fsnotify"

	"monitor/conf"
	"monitor/lib/notify"
)

// 调用运行监控程序
func Run() {
	for _, monitorContent := range conf.CFG.Monitor {
		go RunOneWatch(monitorContent)
	}
}

// 文件监控执行
func RunOneWatch(monitor conf.Monitor) {
	watch, _ := fsnotify.NewWatcher()
	w := notify.Watch{
		Watch: watch,
	}
	w.WatchDir(monitor)
	// 阻塞循环
	select {}
}
