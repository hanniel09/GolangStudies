package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	x := 10.0
	y := "olá bom dia"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T", y, y)

}
