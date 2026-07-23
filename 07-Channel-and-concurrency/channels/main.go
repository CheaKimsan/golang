package main

import (
	"fmt"
	"sync"
)

type EventProcessor struct {
	events    chan string
	processed map[string]bool
	wg        sync.WaitGroup
	mu        sync.Mutex
}

func NewEventProcessor(bufferSize int) *EventProcessor {
	return &EventProcessor{
		events:    make(chan string, bufferSize),
		processed: make(map[string]bool),
	}
}

func (ep *EventProcessor) Start() {
	go func() {
		for event := range ep.events {
			go ep.handleEvent(event)
		}
	}()
}

func (ep *EventProcessor) Process(event string) {
	ep.events <- event
}

func (ep *EventProcessor) Stop() {
	close(ep.events)
}
func (ep *EventProcessor) Wait() {
	ep.wg.Wait()
}

func (ep *EventProcessor) handleEvent(event string) {
	ep.mu.Lock()
	ep.processed[event] = true
	ep.mu.Unlock()
}

func (ep *EventProcessor) GetProcessedEvent() map[string]bool {
	ep.mu.Lock()
	defer ep.mu.Unlock()

	result := make(map[string]bool)
	for k, v := range ep.processed {
		result[k] = v
	}
	return result

}

func main() {
	processer := NewEventProcessor(10)
	processer.Start()

	processer.Process("event1")
	processer.Process("event2")
	processer.Process("event3")

	// processer.Stop()
	processer.Wait()

	processed := processer.GetProcessedEvent()

	if !processed["event1"] {
		fmt.Println("event 1 was not processed")
	}
	if !processed["event2"] {
		fmt.Println("event 2 was not processed")
	}
	if !processed["event3"] {
		fmt.Println("event 3 was not processed")
	}

	if len(processed) != 3 {
		fmt.Println("expected 3 processed events, got", len(processed))
	}
}
