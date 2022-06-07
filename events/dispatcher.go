package events

type Dispatcher struct {
	listeners map[Event][]Lister
}

func NewEventDispatcher() *Dispatcher {
	return &Dispatcher{
		listeners: make(map[Event][]Lister),
	}
}

func (d *Dispatcher) Listen(event Event, listener Lister) {
	d.listeners[event] = append(d.listeners[event], listener)
}

func (d *Dispatcher) dispatch(event Event, args ...interface{}) []interface{} {
	var results []interface{}

	for _, listener := range d.listeners[event] {
		result, err := listener.Handle(args...)
		if err != nil {
			break
		}

		results = append(results, result)
	}

	return results
}
