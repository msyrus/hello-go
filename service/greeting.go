package service

import (
	"fmt"
)

// Greeting is the greeting service
type Greeting struct {
	dflt string
}

// NewGreeting returns a new greeting service
func NewGreeting(msg string) (*Greeting, error) {
	return &Greeting{
		dflt: msg,
	}, nil
}

// GreetUnknown greets unknown
func (s *Greeting) GreetUnknown() (string, error) {
	return "Hello! Who are you?", nil
}

// GreetDefault greets person with default msg
func (s *Greeting) GreetDefault(name string) (string, error) {
	return fmt.Sprintf("Hello %s! Its %s", name, s.dflt), nil
}
