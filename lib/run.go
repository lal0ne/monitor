package lib

import (
	"github.com/fsnotify/fsnotify"
)

func Run() {
	watch, _ := fsnotify.NewWatcher()
	w := Watch{
		watch: watch,
	}
	w.watchDir(CFG.Monitor.PATH)
	// 阻塞循环
	select {}
}
