package tutorial

import (
	"math"
	"testing"
)

func TestSlowAdd(t *testing.T) {
	if result := SlowAdd(1, 1); result != 2 {
		t.Errorf("wrong result: %d", result)
	}
}

func TestSlowAdd_Overflow(t *testing.T) {
	if result := SlowAdd(math.MaxInt64, 1); result != math.MinInt64 {
		t.Errorf("wrong result: %d", result)
	}
}

func TestSlowSub(t *testing.T) {
	if result := SlowSub(1, 1); result != 0 {
		t.Errorf("wrong result: %d", result)
	}
}
