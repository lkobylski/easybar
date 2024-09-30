package main

import (
	"fmt"
	"strings"
	"sync"
	"unicode/utf8"
)

type Option func(bar *EasyBar)

const maxName = 20

type EasyBar struct {
	max     int
	current int
	name    string
	done    bool
	order   int
	maxName int
	lock    sync.RWMutex
}

func NewEasyBar(max int, name string, opts ...Option) *EasyBar {

	if utf8.RuneCountInString(name) > maxName {
		name = name[:maxName-3] + "..."
	}

	eb := &EasyBar{
		name:    name,
		max:     max,
		maxName: maxName,
	}

	for _, opt := range opts {
		opt(eb)
	}

	return eb
}

func WithOrder(n int) Option {
	return func(eb *EasyBar) {
		eb.order = n
	}
}

func UseMultiBars() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	fmt.Print("\033[?25l")
}

func ClearBars() {
	fmt.Printf("\n\033[?25h")
}

func (e *EasyBar) GetMax() int {
	return e.max
}

func (e *EasyBar) GetCurrent() int {
	return e.current
}

func (e *EasyBar) Add(val int) {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.current += val

	if e.current > e.max {
		e.current = e.max
	}
	e.render()
}

func (e *EasyBar) finish() {
	e.done = true
	fmt.Print(" ✅ \r\n")
}

func (e *EasyBar) render() {
	if !e.done {
		percent := (float32(e.current) / float32(e.max)) * 100
		nameWithPadding := fmt.Sprintf("%-*s", maxName, e.name)
		fmt.Printf("\033[%d;0H\033[K%s [%-20s] %.1f%%", e.order, nameWithPadding, strings.Repeat("#", int(percent/5)), percent)
	}
	if e.current == e.max {
		e.finish()
	}
}
