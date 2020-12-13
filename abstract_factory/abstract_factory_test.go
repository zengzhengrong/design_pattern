package abstractfactory

import (
	"testing"
)

func TestNewHome(t *testing.T) {
	home := NewHome("3")
	p := home.GetMember("zzr")
	ps := home.GetMembers()
	p.Body()
	for _, item := range ps {
		item.Body()
	}

}
