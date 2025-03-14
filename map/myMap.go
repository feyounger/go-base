package main

import (
	"context"
	"errors"
	"sync"
	"time"
)

type MyChan struct {
	ch chan struct{}
	sync.Once
}

func NewMyChan() *MyChan {
	return &MyChan{
		ch: make(chan struct{}),
	}
}

func (handle *MyChan) Close() {
	handle.Once.Do(func() {
		close(handle.ch)
	})
}

type MyCurrentMap struct {
	m map[int]int
	sync.Mutex
	chM map[int]*MyChan
}

func NewMyCurrentMap() *MyCurrentMap {
	return &MyCurrentMap{
		m:   make(map[int]int),
		chM: make(map[int]*MyChan),
	}
}

func (handle *MyCurrentMap) Put(key, value int) {
	handle.Lock()
	defer handle.Unlock()
	handle.m[key] = value
	ch, ok := handle.chM[key]
	if !ok {
		return
	}
	//防止多次关闭同样的ch
	ch.Close()
}

func (handle *MyCurrentMap) Get(key int, duration time.Duration) (int, error) {
	handle.Lock()
	value, ok := handle.m[key]
	if ok {
		handle.Unlock()
		return value, nil
	}
	_, ok = handle.chM[key]
	if !ok {
		handle.chM[key] = NewMyChan()
	}

	timeout, cancelFunc := context.WithTimeout(context.Background(), duration)
	defer cancelFunc()
	handle.Unlock()
	select {
	case <-timeout.Done():
		return 0, errors.New("timeout")
	default:
		<-handle.chM[key].ch
	}
	handle.Lock()
	value = handle.m[key]
	handle.Unlock()
	return value, nil
}

func main() {
	currentMap := NewMyCurrentMap()
	go func() {
		get, err := currentMap.Get(1, time.Minute)
		if err != nil {
			println(err)
			return
		}
		println(get)
	}()

	time.Sleep(time.Second * 3)
	currentMap.Put(1, 1)

	time.Sleep(time.Second * 3)
}
