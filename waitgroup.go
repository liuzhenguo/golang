package utils

import (
	"fmt"
	"sync"
)

type Monitor struct {
	routines chan int
	wg       *sync.WaitGroup
}

func NewMoni(size int) *Monitor {
	if size <= 0 {
		size = 0
	}
	return &Monitor{
		routines: make(chan int, size),
		wg:       &sync.WaitGroup{},
	}
}

func (t *Monitor) Add() {
	t.wg.Add(1)
	t.routines <- 1
}
func (t *Monitor) Done() {
	<-t.routines
	t.wg.Done()
}
func (t *Monitor) Wait() {
	t.wg.Wait()
}

type MonitorWrapper struct {
	moni *Monitor
}

func NewMoniWrap(size int) *MonitorWrapper {
	return &MonitorWrapper{
		moni: NewMoni(size),
	}
}
func (t *MonitorWrapper) Wrap(cb func()) {

	t.moni.Add()
	go func() {
		defer t.moni.Done()
		cb()

	}()
	t.Wait()
}

func (t *MonitorWrapper) Wait() {
	fmt.Println("wait func")
	t.moni.Wait()
}
