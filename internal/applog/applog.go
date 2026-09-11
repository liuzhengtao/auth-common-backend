package applog

import (
	"sync"

	"github.com/gogf/gf/v2/os/glog"
)

var (
	mu     sync.RWMutex
	logger glog.ILogger
)

// Set 由宿主通过 install.WithLogger 注入；传 nil 忽略。
func Set(l glog.ILogger) {
	if l == nil {
		return
	}
	mu.Lock()
	logger = l
	mu.Unlock()
}

// Get 返回已注入的 Logger；未注入时懒加载 glog.New().Line(true)。
func Get() glog.ILogger {
	mu.RLock()
	l := logger
	mu.RUnlock()
	if l != nil {
		return l
	}
	mu.Lock()
	defer mu.Unlock()
	if logger == nil {
		logger = glog.New().Line(true)
	}
	return logger
}
