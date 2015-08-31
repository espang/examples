package fib

import "testing"

var result int

var fibTests = []struct {
	n   int
	exp int
}{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{8, 21},
	{9, 34},
}

func TestFib1(t *testing.T) {
	for _, tt := range fibTests {
		cur := fib1(tt.n)
		if cur != tt.exp {
			t.Errorf("Fib1(%d): expected %d, is %d", tt.n, tt.exp, cur)
		}
	}
}

func TestFib2(t *testing.T) {
	for _, tt := range fibTests {
		cur := fib2(tt.n)
		if cur != tt.exp {
			t.Errorf("Fib2(%d): expected %d, is %d", tt.n, tt.exp, cur)
		}
	}
}

func TestFib3(t *testing.T) {
	for _, tt := range fibTests {
		cur := fib3(tt.n)
		if cur != tt.exp {
			t.Errorf("Fib3(%d): expected %d, is %d", tt.n, tt.exp, cur)
		}
	}
}

func TestFib4(t *testing.T) {
	for _, tt := range fibTests {
		cur := fib4(tt.n)
		if cur != tt.exp {
			t.Errorf("Fib4(%d): expected %d, is %d", tt.n, tt.exp, cur)
		}
	}
}

func benchmarkFib1(f func(int) int, i int, b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = f(i)
	}
	result = r
}

func BenchmarkFib1_1(b *testing.B)  { benchmarkFib1(fib1, 1, b) }
func BenchmarkFib1_5(b *testing.B)  { benchmarkFib1(fib1, 5, b) }
func BenchmarkFib1_10(b *testing.B) { benchmarkFib1(fib1, 10, b) }
func BenchmarkFib1_15(b *testing.B) { benchmarkFib1(fib1, 15, b) }
func BenchmarkFib1_20(b *testing.B) { benchmarkFib1(fib1, 20, b) }
func BenchmarkFib1_30(b *testing.B) { benchmarkFib1(fib1, 30, b) }

func BenchmarkFib3_1(b *testing.B)  { benchmarkFib1(fib3, 1, b) }
func BenchmarkFib3_5(b *testing.B)  { benchmarkFib1(fib3, 5, b) }
func BenchmarkFib3_10(b *testing.B) { benchmarkFib1(fib3, 10, b) }
func BenchmarkFib3_15(b *testing.B) { benchmarkFib1(fib3, 15, b) }
func BenchmarkFib3_20(b *testing.B) { benchmarkFib1(fib3, 20, b) }
func BenchmarkFib3_30(b *testing.B) { benchmarkFib1(fib3, 30, b) }

func BenchmarkFib4_1(b *testing.B)  { benchmarkFib1(fib4, 1, b) }
func BenchmarkFib4_5(b *testing.B)  { benchmarkFib1(fib4, 5, b) }
func BenchmarkFib4_10(b *testing.B) { benchmarkFib1(fib4, 10, b) }
func BenchmarkFib4_15(b *testing.B) { benchmarkFib1(fib4, 15, b) }
func BenchmarkFib4_20(b *testing.B) { benchmarkFib1(fib4, 20, b) }
func BenchmarkFib4_30(b *testing.B) { benchmarkFib1(fib4, 30, b) }
