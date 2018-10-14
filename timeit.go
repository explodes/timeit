package timeit

import (
	"fmt"
	"log"
	"time"
)

var _noopTimeit = noopTimeit{}

type Timeit interface {
	Log() Timeit
	Reset(msg string)
	Resetf(msg string, args ...interface{})
}

func New(msg string) Timeit {
	t := &defaultTimeit{}
	t.Reset(msg)
	return t
}

func Newf(msg string, args ...interface{}) Timeit {
	t := &defaultTimeit{}
	t.Resetf(msg, args...)
	return t
}

func NewEnabled(enabled bool, msg string) Timeit {
	if !enabled {
		return _noopTimeit
	}
	return New(msg)
}

func NewEnabledf(enabled bool, msg string, args ...interface{}) Timeit {
	if !enabled {
		return _noopTimeit
	}
	return Newf(msg, args...)
}

type noopTimeit struct{}

func (t noopTimeit) Log() Timeit                            { return t }
func (t noopTimeit) Reset(msg string)                       {}
func (t noopTimeit) Resetf(msg string, args ...interface{}) {}

type defaultTimeit struct {
	start time.Time
	msg   string
}

func (t *defaultTimeit) Log() Timeit {
	duration := time.Now().Sub(t.start)
	log.Printf("TIMEIT: %s: %v\n", t.msg, duration)
	return t
}

func (t *defaultTimeit) Reset(msg string) {
	t.start = time.Now()
	t.msg = msg
}

func (t *defaultTimeit) Resetf(msg string, args ...interface{}) {
	t.Reset(fmt.Sprintf(msg, args...))
}
