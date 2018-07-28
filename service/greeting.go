package service

import (
	"fmt"
)

// Greeting is the greeting service
type Greeting struct {
	greeter string
}

// NewGreeting returns a new greeting service
func NewGreeting(msg string) (*Greeting, error) {
	return &Greeting{
		greeter: msg,
	}, nil
}

// GreetDefault greets person with default msg
func (s *Greeting) GreetDefault(name string) (string, error) {
	if name == "" {
		return "Hello! Who are you?", nil
	}
	return fmt.Sprintf("Hello %s! I'm %s", name, s.greeter), nil
}
