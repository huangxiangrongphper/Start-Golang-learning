package main

import "fmt"

func main() {
	a := 20
	c := 200
	c = a
	fmt.Println("赋值操作，把a赋值给c，所以c的值为：", c)
	c += a
	fmt.Println("相加和赋值运算符，实际为 c = c + a ，所以c的值为：", c)
	c -= a
	fmt.Println("相减和赋值运算符，实际为 c = c - a ，所以c的值为：", c)
	c *= a
	fmt.Println("相乘和赋值运算符，实际为 c = c * a ，所以c的值为：", c)
	c /= a
	fmt.Println("相除和赋值运算符，实际为 c = c / a ，所以c的值为：", c)
	c <<= 2
	fmt.Println("左移和赋值运算符，所以c的值为：", c)
	c >>= 2
	fmt.Println("右移和赋值运算符，所以c的值为：", c)
	c &= 2
	fmt.Println("按位与和赋值运算符，所以c的值为：", c)
	c ^= 2
	fmt.Println("按位异或和赋值运算符，所以c的值为：", c)
	c |= 2
	fmt.Println("按位或和赋值运算符，所以c的值为：", c)

	// 并行赋值的写法：
	const 人数, 费用, 班级 = 40, 200.01, "一班"

	const (
		Monday, Tuesday, Wednesday = 1, 2, 3
		Thursday, Friday, Saturday = 4, 5, 6
	)
}
