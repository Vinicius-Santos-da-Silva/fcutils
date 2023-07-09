package events

import "errors"

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

var ErrorAlreadyHandlerRegistered = errors.New("handler already registered")

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (e *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if _, ok := e.handlers[eventName]; ok {
		for _, h := range e.handlers[eventName] {
			if h == handler {
				return ErrorAlreadyHandlerRegistered
			}
		}
	}
	e.handlers[eventName] = append(e.handlers[eventName], handler)
	return nil
}
