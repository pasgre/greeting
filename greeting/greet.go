// Greeting allows you to greet people
package greeting

// Greet greets users with their name provided
func Greet(name string) string {
	return "Hello " + name + "!"
}

// Person represents a real world human
type Person struct {
	// Name desribes how the person is called
	Name string
}

// Greet greets with the persons name
func (p *Person) Greet() string {
	return Greet(p.Name)
}
