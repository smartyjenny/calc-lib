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

func TestSubtraction_Calculate(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{a: 0, b: 0, want: 0},
		{a: -1, b: 0, want: -1},
		{a: 2, b: 3, want: -1},
		{a: 4, b: 3, want: 1},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d - %d = %d", tt.a, tt.b, tt.want), func(t *testing.T) {
			this := &Subtraction{}
			if got := this.Calculate(tt.a, tt.b); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplication_Calculate(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{a: 0, b: 0, want: 0},
		{a: -1, b: 0, want: 0},
		{a: 3, b: 3, want: 9},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d * %d = %d", tt.a, tt.b, tt.want), func(t *testing.T) {
			this := &Multiplication{}
			if got := this.Calculate(tt.a, tt.b); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivision_Calculate(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{a: 4, b: 2, want: 2},
		{a: 1, b: 1, want: 1},
		{a: 9, b: 3, want: 3},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d / %d = %d", tt.a, tt.b, tt.want), func(t *testing.T) {
			this := &Division{}
			if got := this.Calculate(tt.a, tt.b); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
