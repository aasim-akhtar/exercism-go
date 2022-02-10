package main

import (
	"fmt"
)

func main () {
	var expression bool = false && !(true && false)
	fmt.Println(expression)
}
