package events

// type Listener struct {
// 	event *Event
// }

type Lister interface {
	Handle(...interface{}) (interface{}, error)
}
