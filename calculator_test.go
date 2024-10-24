package calc_lib

import (
	"fmt"
	"testing"
)

func TestAddition_Calculate(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{a: 0, b: 0, want: 0},
		{a: -1, b: 0, want: -1},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d + %d = %d", tt.a, tt.b, tt.want), func(t *testing.T) {
			this := &Addition{}
			if got := this.Calculate(tt.a, tt.b); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
