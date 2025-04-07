package internal

type EventBusInterface interface {
	Subscribe(eventName string, subscriber chan<- Event) error
	Publish(event Event)
}

type Event struct {
	Type string
	Data interface{}
}

type EventBus struct {
	subscribes map[string][]chan<- Event
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribes: make(map[string][]chan<- Event),
	}
}

func (e *EventBus) Subscribe(eventName string, subscriber chan<- Event) error {
	e.subscribes[eventName] = append(e.subscribes[eventName], subscriber)
	return nil
}

func (e *EventBus) Publish(event Event) {
	subscribes := e.subscribes[event.Type]
	for _, subscriber := range subscribes {
		subscriber <- event
	}
}
