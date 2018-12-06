package days

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

type Event struct {
	id        int
	timestamp time.Time
	awake     bool
}
type Events []Event

func (es Events) Len() int {
	return len(es)
}
func (es Events) Swap(i, j int) {
	es[i], es[j] = es[j], es[i]
}
func (es Events) Less(i, j int) bool {
	return es[i].timestamp.Before(es[j].timestamp)
}
func (e *Event) DurationDifference(evt Event) time.Duration {
	return e.timestamp.Sub(evt.timestamp)
}

func NewEvent(ts time.Time) Event {
	return Event{timestamp: ts}
}

func (e *Event) SetId(id int) {
	e.id = id
}

func (e *Event) Id() int {
	return e.id
}

func (e *Event) SetAwake(b bool) {
	e.awake = b
}

func (e *Event) isAwake() bool {
	return e.awake
}

func ParseInput(s string) (time.Time, string) {
	s = strings.TrimPrefix(s, "[")
	ss := strings.Split(s, "] ")
	tss := ss[0]
	action := ss[1]
	timestamp, err := time.Parse("2006-01-02 15:04", tss)
	if err != nil {
		log.Fatal(err)
	}
	return timestamp, action
}

func ParseEvents(input []string) []Event {
	events := make([]Event, 0, len(input))
	for _, line := range input {
		timestamp, action := ParseInput(line)
		evt := NewEvent(timestamp)
		evt.SetAwake(action != "falls asleep")
		if strings.ContainsAny(action, "#") {
			var id int
			if _, err := fmt.Sscanf(action, "Guard #%d begins shift", &id); err != nil {
				log.Fatal(err)
			}
			evt.SetId(id)
		}
		events = append(events, evt)
	}
	sort.Sort(Events(events))
	var id int
	for i := 0; i < len(events); i++ {
		if events[i].Id() > 0 {
			id = events[i].Id()
		} else {
			events[i].SetId(id)
		}
	}
	return events
}

func MostAsleep(events []Event) int {
	tallys := make(map[int]float64)
	var evv Event
	for _, evt := range events {
		fmt.Println("id", evt.id, "Timestamp:", evt.timestamp.String(), "Awake:", evt.isAwake())
		if !evt.awake {
			if evv.Id() != 0 {
				tallys[evt.Id()] += evt.DurationDifference(evv).Minutes()
			}
		}
		evv = evt
	}
	var id int
	var max float64
	for k, v := range tallys {
		if max < v {
			id = k
		}
	}
	return id
}
