package test

import (
	"fmt"
)

//平均数
func Average(sl []float64) float64 {
	var sum float64
	for _, i := range sl {
		sum += i
	}
	return sum / float64(len(sl))
}

//排序
func SortNum(a int, b int) (int, int) {
	if a > b {
		return b, a
	} else {
		return a, b
	}
}

// 栈
var stack []string

func Push(k string) {
	stack = append(stack, k)
	fmt.Printf("My stack %v\n", stack)
}
func Pop() (l string) {

	index := len(stack) - 1
	l = stack[index]
	stack = append(stack[:index], stack[index+1:]...)
	fmt.Printf("My stack %v\n", stack)
	return
}

//变参
func Varargs(arg ...int) {
	for _, n := range arg {
		fmt.Printf("%d\n", n)
	}
}

// 斐波那契数列
func Fibonacci(index int) (sl []int) {
	num1 := index
	num2 := index
	sl = append(sl, num1)
	sl = append(sl, num2)
	for i := 0; i < 10; i++ {
		count := num1 + num2
		num1 = num2
		num2 = count
		sl = append(sl, count)
	}
	return
}

// map函数
func Map(f func(int) int, arr []int) (result []int) {
	for _, i := range arr {
		result = append(result, f(i))
	}
	return
}

// 最大值
func Max(arr []int) (max int) {
	max = arr[0]
	for _, i := range arr {
		if i > max {
			max = i
		}
	}
	return
}

// 最小值
func Min(arr []int) (min int) {
	min = arr[0]
	for _, i := range arr {
		if i < min {
			min = i
		}
	}
	return
}

// 冒泡排序
func BubbleSort(n []int) {
	for i := 0; i < len(n)-1; i++ {
		for j := i + 1; j < len(n); j++ {
			if n[j] < n[i] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}
}

// 函数返回一个函数
func PlusX(x int) func(int) int {
	return func(y int) int { return x + y }
}
