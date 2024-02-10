package main

import "fmt"

func UltimateDivMod(a *int, b *int) {
	tmp := *a
	*a = tmp / *b
	*b = tmp % *b

}
func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
