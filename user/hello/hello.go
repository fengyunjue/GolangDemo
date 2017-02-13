package main

import (
	"fmt"
	"user/stringutil"
	"user/test"
)

func main() {
	fmt.Printf(stringutil.Reverse("!oG ,olleH"))

	sl := []float64{1, 2, 3, 4, 5, 6}
	aver := test.Average(sl)
	fmt.Printf("%f", aver)

	a, b := test.SortNum(12, 14)
	fmt.Printf("%d,%d", a, b)

	// test.Push("m")
	// test.Push("l")
	// test.Pop()

	// test.Varargs(1, 2)

	sl1 := test.Fibonacci(1)
	fmt.Printf("%d", sl1)

	fc := func(a int) (b int) {
		b = a + 1
		return
	}
	fmt.Printf("%d", test.Map(fc, []int{1, 2, 3, 4, 5, 6}))

	fmt.Printf("%d", test.Max([]int{3, 5, 635, 666, 234, 62, 563}))

	fmt.Printf("  %d ", test.Min([]int{3, 5, 635, 666, 234, 62, 563}))

	p := test.PlusX(3)
	fmt.Printf("%d ", p(5))

}
