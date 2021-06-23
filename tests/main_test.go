package tests

import (
	"fmt"
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d, wanted: -2", ans)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(name, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, wang %d", ans, tt.want)
			}
		})
	}
}
