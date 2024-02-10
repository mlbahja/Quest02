package main

import "fmt"

func StrLen(s string) int {
	len := 0
	for range s {
		len++
	}
	return len
}

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}
