package test

import (
	"fmt"
)

func Other() {
	// a, b := 20, 16
	// // var a int
	// // var b bool
	// // 15 = 15
	// b = 20
	// a = 30
	// fmt.Printf("%p,%p", a, b)

	// var a int
	// var b int32
	// a = 15
	// b = a + a
	// b = b + 5

	// const (
	// 	a = iota
	// 	b = iota
	// )

	// var s string = "hello"
	// c := []rune(s)
	// c[0] = 'c'
	// s2 := string(c)
	// fmt.Printf("%s\n", s2)

	// var ss string = "123567" +
	// 	"000000000"

	// var ss2 string = `hehehehe   sdfsdf
	// werwerwfdsfer     `

	// fmt.Printf("   %s---%s", ss, ss2)

	// var e error = 123

	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("%d", i)
	// }

	// 	i := -1
	// add:
	// 	i += 1
	// 	fmt.Printf("%d", i)
	// 	if i < 9 {
	// 		goto add
	// 	}

	//---------------------------------------
	//                For-loop
	//---------------------------------------
	// array := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// for _, v := range array {
	// 	fmt.Printf("%d", v)
	// }

	// var a int = 15
	// fmt.Printf("%d", a)
	//---------------------------------------
	//                FizzBuzz
	//---------------------------------------
	// for i := 1; i < 101; i++ {
	// 	str := strconv.Itoa(i)
	// 	if i%3 == 0 {
	// 		str = "Fizz"
	// 		if i%5 == 0 {
	// 			str = str + "Buzz"
	// 		}
	// 	} else if i%5 == 0 {
	// 		str = "Buzz"
	// 	}
	// 	fmt.Printf(str + "\n")
	// }
	//---------------------------------------
	//                字符串
	//---------------------------------------

	// str := "A"
	// for i := 0; i < 100; i++ {
	// 	fmt.Printf(str + "\n")
	// 	str += "A"
	// }

	str := "asSASA ddd dsjkdsjs dk"
	a := []rune(str)
	copy(a[4:7], []rune("abc"))
	// str2 := string(a)
	// fmt.Printf("%d---%s\n", len(str),, str2)
	//---------------------------------------
	//                平均值
	//---------------------------------------
	sl := []int{1, 2, 3, 4, 5, 6}

	count := 0
	for _, i := range sl {
		count += i
	}
	ss := count / len(sl)
	fmt.Printf("%d", ss)
}
