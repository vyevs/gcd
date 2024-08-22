package gcd

import (
	"fmt"
	"testing"
)

var tests = []struct {
	a, b, want int
} {
	{
		a: 0, b: 0, want: 0,
	},
	{
		a: 1, b: 0, want: 1,
	},
	{
		a: 0, b: 1, want: 1,
	},
	{
		a: 48, b: 18, want: 6,
	},
	{
		a: 18, b: 48, want: 6,
	},
	{
		a: 461952, b: 116298, want: 18,
	},
	{
		a: 7966496, b: 314080416, want: 32,
	},
	{
		a: 24826148, b: 45296490, want: 526,
	},
}

func TestGCD(t *testing.T) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("gcd(%d, %d)", test.a, test.b), func(t *testing.T) {
			got := GCD(test.a, test.b)
			
			if got != test.want {
				t.Fatalf("got %d want %d", got, test.want)
			}
		})
	}
}

func BenchmarkGCD(b *testing.B) {
	funcs := []struct {
		name string
		f func(a, b int) int
	} {
		{
			name: "euclidean",
			f: euclidean[int],
		},
		{
			name: "binary",
			f: binary[int],
		},
	}
	
	for _, testFunc := range funcs {
		for _, test := range tests {
			b.Run(fmt.Sprintf("%s(%d, %d)", testFunc.name, test.a, test.b), func(b *testing.B) {
				for range b.N {
					_ = testFunc.f(test.a, test.b)
				}
			})
		}
	}
}