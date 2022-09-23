package greeting

import "testing"

func TestGreet(t *testing.T) {
	greeting := Greet("John Doe")
	if greeting != "Hello John Doe!" {
		t.Errorf("'%s' was not the expected string 'Hello John Doe!'", greeting)
	}
}

func TestPersonGreet(t *testing.T) {
	p := Person{Name: "John Doe"}
	greeting := p.Greet()
	if greeting != "Hello John Doe!" {
		t.Errorf("'%s' was not the expected string 'Hello John Doe!'", greeting)
	}
}
