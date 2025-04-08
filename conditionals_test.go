package condutil

import (
	"testing"
)

func TestIsZeroValue(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		want  bool
	}{
		// Boolean tests
		{"Zero boolean", false, true},
		{"Non-zero boolean", true, false},

		// Integer tests
		{"Zero int", 0, true},
		{"Non-zero int", 1, false},
		{"Zero int64", int64(0), true},
		{"Non-zero int64", int64(1), false},

		// Unsigned integer tests
		{"Zero uint", uint(0), true},
		{"Non-zero uint", uint(1), false},

		// Float tests
		{"Zero float32", float32(0), true},
		{"Non-zero float32", float32(1.1), false},
		{"Zero float64", float64(0), true},
		{"Non-zero float64", float64(1.1), false},

		// Complex number tests
		{"Zero complex64", complex64(0), true},
		{"Non-zero complex64", complex64(1 + 2i), false},
		{"Zero complex128", complex128(0), true},
		{"Non-zero complex128", complex128(1 + 2i), false},

		// String tests
		{"Zero string", "", true},
		{"Non-zero string", "hello", false},

		// Pointer tests
		{"Nil pointer", (*int)(nil), true},
		{"Non-nil pointer", new(int), false},

		// Slice tests
		{"Nil slice", []int(nil), true},
		{"Empty slice", []int{}, false},
		{"Non-empty slice", []int{1, 2, 3}, false},

		// Map tests
		{"Nil map", map[string]int(nil), true},
		{"Empty map", map[string]int{}, false},
		{"Non-empty map", map[string]int{"a": 1}, false},

		// Array tests
		{"Zero array", [3]int{}, true},
		{"Non-zero array", [3]int{1, 2, 3}, false},

		// Struct tests
		{"Zero struct", struct{}{}, true},
		{"Non-zero struct", struct{ A int }{A: 1}, false},

		// Channel tests
		{"Nil channel", (chan int)(nil), true},
		{"Non-nil channel", make(chan int), false},

		// Function tests
		{"Nil function", (func())(nil), true},
		{"Non-nil function", func() {}, false},

		// Interface tests
		{"Nil interface", nil, true},
		{"Non-nil interface", interface{}("hello"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsZeroValue(tt.value); got != tt.want {
				t.Errorf("IsZeroValue(%v) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}

func TestIsEmptySlice(t *testing.T) {
	var nilSlice []int
	tests := []struct {
		name  string
		slice interface{}
		want  bool
	}{
		{"Empty slice", []int{}, true},
		{"Non-empty slice", []int{1, 2, 3}, false},
		{"Nil slice", nilSlice, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmptySlice(tt.slice); got != tt.want {
				t.Errorf("IsEmptySlice(%v) = %v, want %v", tt.slice, got, tt.want)
			}
		})
	}
}

func TestIsEmptyMap(t *testing.T) {
	var nilMap map[string]int
	tests := []struct {
		name string
		m    interface{}
		want bool
	}{
		{"Empty map", map[string]int{}, true},
		{"Non-empty map", map[string]int{"a": 1}, false},
		{"Nil map", nilMap, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmptyMap(tt.m); got != tt.want {
				t.Errorf("IsEmptyMap(%v) = %v, want %v", tt.m, got, tt.want)
			}
		})
	}
}
