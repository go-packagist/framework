package cache

import (
	"errors"
	"sync"
	"time"
)

// _data is a struct for data
type _data struct {
	val    interface{}
	expire time.Time
}

// memory is a cache store that stores data in memory.
type memory struct {
	data    map[string]*_data
	locker  *sync.RWMutex
	locker2 *sync.RWMutex
}

var _ Cacher = (*memory)(nil)

// NewMemory New memory store returns a new memory store.
func NewMemory() Cacher {
	m := &memory{
		data:    make(map[string]*_data),
		locker:  &sync.RWMutex{},
		locker2: &sync.RWMutex{},
	}

	go m.GC()

	return m
}

// Get returns the value by key.
func (m *memory) Get(key string) *Result {
	data, ok := m.data[key]

	if ok {
		return &Result{data.val, nil}
	}

	return &Result{nil, errors.New("key not found")}
}

// Put puts the value by key.
func (m *memory) Put(key string, value interface{}, expire time.Duration) error {
	m.locker.Lock()
	defer m.locker.Unlock()

	m.data[key] = &_data{
		val:    value,
		expire: time.Now().Add(expire),
	}

	return nil
}

// Has returns whether the key exists.
func (m *memory) Has(key string) bool {
	_, ok := m.data[key]

	return ok
}

// Remember returns the value by key.
func (m *memory) Remember(key string, fc func() interface{}, expire time.Duration) *Result {
	if !m.Has(key) {
		if err := m.Put(key, fc(), expire); err != nil {
			return &Result{nil, err}
		}
	}

	return m.Get(key)
}

// Forget deletes the key.
func (m *memory) Forget(key string) error {
	if !m.Has(key) {
		return errors.New("key not found")
	}

	m.locker.Lock()
	defer m.locker.Unlock()

	delete(m.data, key)

	return nil
}

// Forever puts the value by key forever.
func (m *memory) Forever(key string, value interface{}) error {
	return m.Put(key, value, time.Hour*24*365*100) // set long long ... long time
}

// Increment increments the value by key.
func (m *memory) Increment(key string, num ...int) *Result {
	// rw := &sync.RWMutex{}
	// m.locker2.Lock()
	// defer m.locker2.Unlock()

	var n = 1

	if len(num) > 0 {
		n = num[0]
	}

	if r := m.Get(key); r.Error == nil {
		switch r.Value.(type) {
		case int:
			r.Value = n + r.Value.(int)
			return &Result{r.Value, nil}
		default:
			return &Result{nil, errors.New("num is not int")}
		}
	}

	if err := m.Forever(key, n); err != nil {
		return &Result{nil, err}
	}

	return &Result{n, nil}
}

// Decrement decrements the value by key.
func (m *memory) Decrement(key string, num ...int) *Result {
	var n = 1

	if len(num) > 0 {
		n = num[0]
	}

	return m.Increment(key, -1*n)
}

// GC garbage collection.
func (m *memory) GC() {
	for {
		time.Sleep(time.Second * 1)

		for k, v := range m.data {
			if v.expire.Before(time.Now()) {
				m.locker.Lock()
				delete(m.data, k)
				m.locker.Unlock()
			}
		}
	}
}
