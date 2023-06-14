package ants

import (
	"errors"
	ants "github.com/panjf2000/ants/v2"
	"runtime"
)

var ants_pool *ants.Pool
var ants_pool_task *ants.Pool

//  please remenber use release method to free your memory
func AntsPool() *ants.Pool {

	if ants_pool == nil {
		var err error = errors.New("")
		ants_pool, err = ants.NewPool(runtime.NumCPU()*8, ants.WithLogger(new(logger)))
		if err != nil {
			panic(any(err))
		}
	}
	return ants_pool
}

func AntsTaskPool() *ants.Pool {

	if ants_pool_task == nil {
		var err error = errors.New("")
		ants_pool_task, err = ants.NewPool(runtime.NumCPU()*8, ants.WithLogger(new(logger)))
		if err != nil {
			panic(any(err))
		}
	}
	return ants_pool_task
}

// 设置ants 日志
type PrintF func(format string, args ...interface{})

var _handle PrintF

type logger struct {
}

func (this *logger) Printf(format string, args ...interface{}) {
	_handle(format, args...)
}

func SetPrintf(fn PrintF) {
	_handle = fn
}

/**
example:
	ants.Init(func(){
		ants.SetPrintf(ants.PrintF(f))
	})
*/
func Init(f func()) {
	f()
}
