package lock

import (
	"TranslateUtil/constent"
	"sync"
	"time"
)

var SyncCounts = 0
var mux sync.Mutex

/*
上锁 注意这是限时锁 所以无需解锁
 */
func LockSeconds() int{
	if SyncCounts < 0 {
		panic("mutex counts less than zero!")
	}
	tar := SyncCounts
	mux.Lock()
	SyncCounts++
	mux.Unlock()
	//限定时间后-1
	go func() {
		time.Sleep(constent.LOCKTIME)
		mux.Lock()
		SyncCounts--
		mux.Unlock()
	}()
	return tar
}