package stream

type EventBus struct {
	subscribers []chan string
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make([]chan string, 0),
	}
}

func (b *EventBus) Subscribe(ch chan string) {
	b.subscribers = append(b.subscribers, ch)
}

func (b *EventBus) Publish(event string) {
	for _, sub := range b.subscribers {
		select {
		case sub <- event:
		default:
		}
	}
}
