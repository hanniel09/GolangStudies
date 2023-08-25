package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	var a [2]string
	x := "Hello"
	y := "World"
	a[0] = x
	a[1] = y
	fmt.Println(a)

	s := "Hi, everyone"
	sb := []byte(s)

	for _, v := range sb {
		fmt.Printf("%v - %T - %#U - %#X", v, v, v, v)
	}

	for _, v := range s {
		fmt.Printf("%v - %T - %#U - %#X", v, v, v, v)
	}

	fmt.Println(" ")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v - %T - %#U - %#X", s[i], s[i], s[i], s[i])
	}

}
