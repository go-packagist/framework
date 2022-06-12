package cache

import "time"

type Cacher interface {
	Get(string) *Result
	Put(string, interface{}, time.Duration) error
	Has(string) bool
	Remember(string, func() interface{}, time.Duration) *Result
	Forget(string) error
	Forever(string, interface{}) error
	Increment(string, ...int) *Result
	Decrement(string, ...int) *Result
	GC()
}

type Result struct {
	Value interface{}
	Error error
}
