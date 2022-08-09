package main

import "fmt"

func main() {

	var (
		code, quantit [2]int64
		unitaryVulue  [2]float64
	)

	for l := 0; l < 2; l++ {
		fmt.Scan(&code[l])
		fmt.Scan(&quantit[l])
		fmt.Scan(&unitaryVulue[l])
	}

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", (float64(quantit[0])*unitaryVulue[0])+(float64(quantit[1])*unitaryVulue[1]))
}
