package assert

import (
	"testing"
)

var t = testing.T{}

type Assert struct {
	T *testing.T
}

func NewAssert(t *testing.T) *Assert {
	return &Assert{
		T: t,
	}
}

// Equal(1,2,"1 should equal to 2")
func (this *Assert) Equal(vals ...interface{}) {
	if len(vals) == 3 {
		this.T.Log(vals[2])
	}
	if len(vals) >= 2 {
		if vals[0] != vals[1] {
			this.T.Errorf("Expected %v, got %v.", vals[1], vals[0])
		}
	}

}
