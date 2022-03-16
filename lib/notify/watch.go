package notify

import (
	// "bufio"
	// "fmt"
	"github.com/fsnotify/fsnotify"
	"os"
	"path/filepath"
	// "strings"
	"time"

	"monitor/lib"
)

type Watch struct {
	Watch *fsnotify.Watcher
}

//监控目录
func (w *Watch) WatchDir(dir string) {
	//通过Walk来遍历目录下的所有子目录
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		//这里判断是否为目录，只需监控目录即可
		//目录下的文件也在监控范围内，不需要我们一个一个加
		if info.IsDir() {
			path, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			err = w.Watch.Add(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	// 检测协程
	go func() {
		for {
			select {
			case ev, ok := <-w.Watch.Events:
				if !ok {
					return
				}
				var msg string
				nowTime := time.Now().Format("2006-01-02 15:04:05")
				if ev.Op&fsnotify.Create == fsnotify.Create {
					msg = "Create file: " + ev.Name
					//这里获取新创建文件的信息，如果是目录，则加入监控中
					fi, err := os.Stat(ev.Name)
					if err == nil && fi.IsDir() {
						w.Watch.Add(ev.Name)
						//fmt.Println("添加监控 : ", ev.Name)
					}
				}
				if ev.Op&fsnotify.Write == fsnotify.Write {
					msg = "Write file: " + ev.Name
				}
				if ev.Op&fsnotify.Remove == fsnotify.Remove {
					msg = "Delete file: " + ev.Name
					//如果删除文件是目录，则移除监控
					fi, err := os.Stat(ev.Name)
					if err == nil && fi.IsDir() {
						w.Watch.Remove(ev.Name)
						//fmt.Println("删除监控 : ", ev.Name)
					}
				}
				if ev.Op&fsnotify.Rename == fsnotify.Rename {
					msg = "Rename file: " + ev.Name
					//如果重命名文件是目录，则移除监控
					//注意这里无法使用os.Stat来判断是否是目录了
					//因为重命名后，go已经无法找到原文件来获取信息了
					//所以这里就简单粗爆的直接remove好了
					w.Watch.Remove(ev.Name)
				}
				if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
					msg = "Attrib file: " + ev.Name
				}

				if msg != "" {
					go lib.HandleIgnoreMsg(nowTime, msg, ev.Name)
					//fmt.Println(nowTime + " Event: " + msg)
				}

			case _, ok := <-w.Watch.Errors:
				if !ok {
					return
				}

				return

			}
		}
	}()
}
