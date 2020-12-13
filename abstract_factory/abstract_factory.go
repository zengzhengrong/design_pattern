package abstractfactory

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

// Home is interface
type Home interface {
	GetMember(name string) Person
	GetMembers() []Person
}

// Home1 is
type Home1 struct{}

// Home2 is
type Home2 struct{}

// Home3 is
type Home3 struct{}

//GetMember is
func (h *Home1) GetMember(name string) Person {
	return &Zeron{}
}

// GetMembers is
func (h *Home1) GetMembers() []Person {
	return []Person{&Zeron{}}
}

//GetMember is
func (h *Home2) GetMember(name string) Person {
	return &Zengzhengrong{}

}

// GetMembers is
func (h *Home2) GetMembers() []Person {
	return []Person{&Zengzhengrong{}}
}

//GetMember is
func (h *Home3) GetMember(name string) Person {
	switch name {
	case "zeron":
		return new(Zeron)
	case "zengzhengrong":
		return new(Zengzhengrong)
	}
	return nil

}

//GetMembers is
func (h *Home3) GetMembers() []Person {
	var ms []Person
	z := new(Zeron)
	zz := new(Zengzhengrong)
	ms = append(ms, z)
	ms = append(ms, zz)
	return ms
}

// NewHome is new instance
func NewHome(name string) Home {
	switch name {
	case "1":
		return &Home1{}
	case "2":
		return &Home2{}
	case "3":
		return &Home3{}
	default:
		return nil
	}
}
