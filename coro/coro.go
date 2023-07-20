package coro

import (
	"errors"
	"fmt"
)

type msg[T any] struct {
	panic any
	val   T
}

var ErrCanceled = errors.New("coroutine canceled")

func New[In, Out any](f func(in In, yield func(Out) In) Out) (resume func(In) (Out, bool), cancel func()) {
	cin := make(chan msg[In])
	cout := make(chan msg[Out])
	running := true
	resume = func(in In) (out Out, ok bool) {
		if !running {
			return
		}
		cin <- msg[In]{val: in}
		m := <-cout
		if m.panic != nil {
			panic(m.panic)
		}
		return m.val, running
	}
	cancel = func() {
		e := fmt.Errorf("%w", ErrCanceled) // unique wrapper
		cin <- msg[In]{panic: e}
		m := <-cout
		if m.panic != nil && m.panic != e {
			panic(m.panic)
		}
	}
	yield := func(out Out) In {
		cout <- msg[Out]{val: out}
		m := <-cin
		if m.panic != nil {
			panic(m.panic)
		}
		return m.val
	}
	go func() {
		defer func() {
			if running {
				running = false
				cout <- msg[Out]{panic: recover()}
			}
		}()
		var out Out
		m := <-cin
		if m.panic == nil {
			out = f(m.val, yield)
		}
		running = false
		cout <- msg[Out]{val: out}
	}()
	return resume, cancel
}

type msg0 struct {
	panic any
}

func New0(f func(yield func())) (resume func() bool, cancel func()) {
	cin := make(chan msg0)
	cout := make(chan msg0)
	running := true
	resume = func() bool {
		if !running {
			return false
		}
		cin <- msg0{}
		m := <-cout
		if m.panic != nil {
			panic(m.panic)
		}
		return running
	}
	cancel = func() {
		e := fmt.Errorf("%w", ErrCanceled) // unique wrapper
		cin <- msg0{panic: e}
		m := <-cout
		if m.panic != nil && m.panic != e {
			panic(m.panic)
		}
	}
	yield := func() {
		cout <- msg0{}
		m := <-cin
		if m.panic != nil {
			panic(m.panic)
		}
		return
	}
	go func() {
		defer func() {
			if running {
				running = false
				cout <- msg0{panic: recover()}
			}
		}()
		m := <-cin
		if m.panic == nil {
			f(yield)
		}
		running = false
		cout <- msg0{}
	}()
	return resume, cancel
}
