package main

import "fmt"

func main() {

	var (
		name   string
		fixed  float64
		amount float64
	)

	fmt.Scan(&name)
	fmt.Scan(&fixed)
	fmt.Scan(&amount)

	fmt.Printf("TOTAL = R$ %.2f\n", (15.00/100.00*amount)+fixed)
}
