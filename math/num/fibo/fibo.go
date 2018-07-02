package fibo

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-07-02

import (
	"context"
)

type Fibo struct {
	next   chan int64
	cancel context.CancelFunc
}

func New() *Fibo {
	obj := &Fibo{
		next: make(chan int64, 10),
	}

	ctx, fn := context.WithCancel(context.Background())

	obj.cancel = fn

	go maker(ctx, obj.next)

	return obj
}

func maker(cxt context.Context, stream chan int64) {
	cur := int64(1)
	prev := int64(0)

	for {

		select {
		case stream <- cur:
			cur, prev = cur+prev, cur
		case <-cxt.Done():
			close(stream)
			return
		}

	}
}

func (f *Fibo) Get() int64 {
	return <-f.next
}

func (f *Fibo) Close() {
	f.cancel()
}
