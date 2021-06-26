package mediator

import (
	"testing"
)

func TestRenting(t *testing.T) {
	m := NewMediator()
	m.tenant.Send(100)

}
