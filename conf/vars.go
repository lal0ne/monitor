package conf

import "sync"

var (
	CFG = Cfg{}
	// 读写锁
	RWfile = sync.RWMutex{}
)
