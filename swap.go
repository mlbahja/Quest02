package main

import "fmt"

func Swap(a *int, b *int) {
	var tmp int
	tmp = *a
	*a = *b
	*b = tmp

}
func main() {
	a := 0
	b := 1
	Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
