package main

import (
	"fmt"
	"sync"
	"time"
)

type Property struct {
	Key   string
	Value string
}

type Event struct {
	Name       string
	Properties []Property
	Timestamp  time.Time
}

type Events struct {
	sync.Mutex
	start  time.Time
	events []Event
}

func NewEvents() *Events {
	return &Events{
		events: make([]Event, 0),
		start:  time.Now(),
	}
}

func (events *Events) Add(name string, props ...Property) {
	events.Lock()
	defer events.Unlock()

	events.events = append(events.events, Event{
		Name:       name,
		Properties: props,
		Timestamp:  time.Now(),
	})
}

func (events *Events) Print() {
	for _, event := range events.events {
		fmt.Printf("%s: %v ", event.Name, event.Timestamp.Sub(events.start))
		for _, prop := range event.Properties {
			fmt.Printf("[%s:%s] ", prop.Key, prop.Value)
		}
		fmt.Println()
	}
}