package main

import "fmt"

func main() {
	var a1, a2, a3 int
	var b1, b2, b3 int
	fmt.Scanf("%d %d %d", &a1, &a2, &a3)
	fmt.Scanf("%d %d %d", &b1, &b2, &b3)

	ac1, bc1 := compare(a1, b1)
	ac2, bc2 := compare(a2, b2)
	ac3, bc3 := compare(a3, b3)

	alice := ac1 + ac2 + ac3
	bob := bc1 + bc2 + bc3

	fmt.Printf("%d %d\n", alice, bob)
}

func compare(a, b int) (int, int) {
	if a > b {
		return 1, 0
	} else if a < b {
		return 0, 1
	}
	return 0, 0
}
