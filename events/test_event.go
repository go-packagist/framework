package events

type testEvent struct {
	Name string
}

var _ Event = (*testEvent)(nil)

func NewTestEvent(name string) Event {
	return &testEvent{
		Name: name,
	}
}
