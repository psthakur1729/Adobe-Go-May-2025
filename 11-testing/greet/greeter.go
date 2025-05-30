package greet

import (
	"fmt"
	"time"
)

type TimeProvider interface {
	GetCurrent() time.Time
}

type Greeter struct {
	timeProvider TimeProvider
}

func NewGreeter(tp TimeProvider) *Greeter {
	return &Greeter{
		timeProvider: tp,
	}
}

func (g Greeter) Greet(userName string) string {
	if g.timeProvider.GetCurrent().Hour() < 12 {
		return fmt.Sprintf("Hi %s, Good Morning!", userName)
	}
	return fmt.Sprintf("Hi %s, Good Day!", userName)
}
