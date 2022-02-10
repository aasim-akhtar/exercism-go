package main

import (
	"fmt"
	"strings"
)

func main () {
	x := "boo ****"
	var newString string = x
    newString = strings.ReplaceAll(newString,"*", " ")
    // strings.TrimSpace(newString)
    fmt.Println("DEBUGINFO: " + newString)
}
