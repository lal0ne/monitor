package cmd

import (
	"github.com/fsnotify/fsnotify"

	"monitor/conf"
	"monitor/lib/notify"
)

// 调用运行监控程序
func Run() {
	watch, _ := fsnotify.NewWatcher()
	w := notify.Watch{
		Watch: watch,
	}
	w.WatchDir(conf.CFG.Monitor.PATH)
	// 阻塞循环
	select {}
}
