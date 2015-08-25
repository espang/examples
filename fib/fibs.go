//
// 1 1 2 3 5 8
//
package fib

func fib1(n int) int {
	if n < 1 {
		panic("fib not defined numbers < 1")
	}
	switch n {
	case 1:
		return 1
	case 2:
		return 1
	default:
		return fib1(n-1) + fib1(n-2)
	}
}

func fib2(n int) int {
	if n < 1 {
		panic("fib not defined for numbers < 1")
	}
	return fib2i(n)
}
func fib2i(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 1
	default:
		return fib2i(n-1) + fib2i(n-2)
	}
}

type recfunc func(recfunc, int) int

func fib3(n int) int {
	if n < 1 {
		panic("fib not defined for numbers < 1")
	}
	fibs := map[int]int{
		1: 1,
		2: 1,
	}

	ifib := func(f recfunc, n int) int {
		if val, ok := fibs[n]; ok {
			return val
		}
		v := f(f, n-1) + f(f, n-2)
		fibs[n] = v
		return v
	}

	return ifib(ifib, n)
}

func fib4(n int) int {
	if n < 1 {
		panic("fib not defined for numbers < 1")
	}
	if n == 1 || n == 2 {
		return 1
	}
	a := 1
	b := 1
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
