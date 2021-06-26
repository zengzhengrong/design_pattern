package mediator

import "fmt"

// 中介
type Mediator interface {
	Communicate(c Customer, num int, t int)
}

// 顾客
type Customer interface {
	SetMediator(mediaotr Mediator)
}

// 租客
type Tenant struct {
	mediator Mediator
}

// 找中介
func (t *Tenant) SetMediator(mediator Mediator) {

	t.mediator = mediator
}

func (t *Tenant) Send(num int) {
	// Before Send logic
	t.mediator.Communicate(t, num, 1)
}

// 从中介获取房东信息
func (t *Tenant) Receive(num int) {

	t.mediator.Communicate(t, num, 2)
}

// 房东
type Landlord struct {
	mediator Mediator
}

// 找中介
func (l *Landlord) SetMediator(mediator Mediator) {

	l.mediator = mediator
}

func (l *Landlord) Send(num int) {
	// Before Send logic
	l.mediator.Communicate(l, num, 1)
}

// 从中介获取租客信息
func (l *Landlord) Receive(num int) {

	l.mediator.Communicate(l, num, 2)
}

// 经理
type Manager struct {
	tenant   *Tenant
	landlord *Landlord
}

func (m *Manager) Communicate(c Customer, num int, t int) {
	if c == m.tenant && t == 1 {
		num = num - 50
		content := fmt.Sprintf("%v ,ok?", num)
		fmt.Println(content)
		m.landlord.mediator.Communicate(m.landlord, num, 1)
	}
	if c == m.landlord && t == 1 {
		if 0 < num && num < 100 {
			fmt.Println("no!")
		} else {
			fmt.Println("ok!")
		}

	}
}

func NewMediator() *Manager {
	tenant := &Tenant{}
	landlord := &Landlord{}
	mediator := &Manager{
		tenant:   tenant,
		landlord: landlord,
	}
	mediator.tenant.SetMediator(mediator)
	mediator.landlord.SetMediator(mediator)
	return mediator
}
