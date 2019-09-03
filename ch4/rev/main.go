package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	reverse(a[:])
	fmt.Println(a)
	fmt.Println("")
	b := []string{"Domingo", "Segunda", "Terça"}
	c := []string{"Domingo", "Segunda", "Quinta"}
	fmt.Println(equal(b, c))
}
