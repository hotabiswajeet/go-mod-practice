package main

import (
	"fmt"

	"github.com/hotabiswajeet/go-mod-practice/math"
	"github.com/hotabiswajeet/go-mod-practice/stringutil"
)

func main() {
	a := 10
	b := 20
	fmt.Println(math.CalculateSum(a, b))
	c := "Hello"
	fmt.Println(stringutil.StringReverse(c))
}
