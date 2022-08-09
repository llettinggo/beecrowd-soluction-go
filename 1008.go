package main

import (
	"fmt"
)

func main() {

	var (
		a, b int64
		c    float64
	)

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	fmt.Printf("NUMBER = %d\n", a)
	fmt.Printf("SALARY = U$ %.2f\n", float64(b)*c)
}
