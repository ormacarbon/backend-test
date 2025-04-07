package internal

import "context"

const EventTypeIncreasePoint = "increase_point"

type IncreasePointEventData struct {
	Referal string
}

func RegisterIncreasePointHandler(eventBus *EventBus, usecase *IncreasePoint) {
	ch := make(chan Event)

	_ = eventBus.Subscribe(EventTypeIncreasePoint, ch)

	go func() {
		for evt := range ch {
			data, ok := evt.Data.(IncreasePointEventData)
			if !ok {
				continue
			}

			_ = usecase.Execute(context.Background(), InputIncreasePoint{
				Referal: data.Referal,
			})
		}
	}()
}
