package events

import (
	"fmt"
	"testing"
)

func TestDispatcher_Listen(t *testing.T) {
	d := NewEventDispatcher()
	e := NewTestEvent()

	d.Listen(e, newTestListener(e))
	d.Listen(e, newTestListener(e))

	fmt.Println(d.dispatch(e))
}
