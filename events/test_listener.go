package events

type testListener struct {
	event Event
}

var _ Lister = (*testListener)(nil)

func newTestListener(event Event) Lister {
	return &testListener{
		event: event,
	}
}

func (t *testListener) Handle(i ...interface{}) (interface{}, error) {
	return i.event.Name, nil
	return t.event + i.(string), nil
}
