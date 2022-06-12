package cache

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestMemory(t *testing.T) {
	m := NewMemory()

	// put/get
	assert.Equal(t, nil, m.Put("key", "value", time.Second*1))
	assert.Equal(t, "value", m.Get("key").Value)

	// gc
	time.Sleep(time.Second * 2)
	assert.Error(t, m.Get("key").Error)

	// forever
	assert.Equal(t, nil, m.Forever("key2", "value2"))
	assert.Equal(t, "value2", m.Get("key2").Value)

	// incr
	assert.Equal(t, 1, m.Increment("key3").Value)
	assert.Equal(t, 1, m.Increment("key4", 1).Value)
	assert.Equal(t, 2, m.Increment("key4", 1).Value)

}

func BenchmarkMemory(b *testing.B) {
	m := NewMemory()
	wg := &sync.WaitGroup{}

	for i := 1; i <= b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Increment("key")
		}(i)
	}

	wg.Wait()

	// assert.Equal(b, b.N, m.Get("key").Value)
}
