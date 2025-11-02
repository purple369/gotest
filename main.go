package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(decimal.New(123456789, 0))
	price1 := 0.1
	price2 := 0.21111
	//total := price1 + price2
	p1 := decimal.NewFromFloat(price1)
	p2 := decimal.NewFromFloat(price2)
	total, exp := p1.Add(p2).Round(3).Float64()
	fmt.Println(total, exp)
}
