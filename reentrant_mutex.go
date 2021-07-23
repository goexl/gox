package gox

import (
	`fmt`
	`sync`
	`sync/atomic`
)

type reentrantMutex struct {
	sync.Mutex

	owner     uint64
	recursion int32
}

// NewReentrantMutex 创建新的可重入锁
func NewReentrantMutex() (locker sync.Locker) {
	rm := &reentrantMutex{
		Mutex:     sync.Mutex{},
		owner:     0,
		recursion: 0,
	}
	locker = rm

	return
}

func (rm *reentrantMutex) Lock() {
	gid := Gid()
	if gid == atomic.LoadUint64(&rm.owner) {
		rm.recursion++
	} else {
		rm.Mutex.Lock()
		atomic.StoreUint64(&rm.owner, gid)
		rm.recursion = 1
	}
}

func (rm *reentrantMutex) Unlock() {
	gid := Gid()
	if gid != atomic.LoadUint64(&rm.owner) {
		panic(fmt.Sprintf("错误的协程持有者（%d）：%d！", rm.owner, gid))
	}

	rm.recursion--
	if 0 != rm.recursion {
		return
	}

	atomic.StoreUint64(&rm.owner, -1)
	rm.Mutex.Unlock()
}
