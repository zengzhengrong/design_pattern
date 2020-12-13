package factory

import "fmt"

// Person is factory interface
type Person interface {
	Body()
}

// Zeron is instace for Person
type Zeron struct{}

// Zengzhengrong is instace for Person
type Zengzhengrong struct{}

// Body is for Zeron
func (p *Zeron) Body() {
	fmt.Println("Zeron body")
}

// Body is for Zengzhengrong
func (p *Zengzhengrong) Body() {
	fmt.Println("Zengzhengrong body")
}

// NewPerson is new instance
func NewPerson(name string) Person {
	switch name {
	case "zeron":
		return &Zeron{}
	case "zengzhengrong":
		return &Zengzhengrong{}
	default:
		return nil
	}
}
