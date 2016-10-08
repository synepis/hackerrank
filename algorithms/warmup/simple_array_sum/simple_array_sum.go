package main

import "fmt"

func main() {
	var n, sum, tmp int = 0, 0, 0
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &tmp)
		sum += tmp
	}
	fmt.Printf("%d\n", sum)
}
