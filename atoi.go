package main

import "fmt"

func Atoi(s string) int {
	res := 0
	sign := 1
	for index, value := range s {
		if value == '-' && index == 0 {
			sign = -1
		} else if value == '+' && index == 0 {
			sign = 1
		} else if value >= 'a' && value <= 'z' {
			return 0
		} else if value >= '0' && value <= '9' {
			res *= 10
			res += int(value - '0')
		} else {
			return 0
		}
	}
	return (res * sign)
}
func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
