package factory

import "testing"

func TestNewPerson(t *testing.T) {
	NewPerson("zeron").Body()
	NewPerson("zengzhengrong").Body()
}
