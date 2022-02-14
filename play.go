package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main () {
	x := "boo ****"
	var newString string = x
    newString = strings.ReplaceAll(newString,"*", " ")
    // strings.TrimSpace(newString)
    fmt.Println("" + newString)

	
// if (v :=2) {
//     fmt.Println(v)
// } else {
//     fmt.Println(v)
// }

choice1 :=  "Ford Focus"
choice2 :=  "Ford Pinto"

fmt.Println(choice1>choice2)

withData := []int{0,1,2,3,4,5}
// len(withData)
fmt.Println()
fmt.Println(withData)
withData = append(withData, 6)
var value int = 4
var length int = 10
var slice = make([]int,length)
fmt.Printf("Length of slice= %d \n", len(slice))
for e := 0 ; e< length; e++{
	fmt.Println("accessing " + strconv.Itoa(e) + "index")
	slice[e] = value

}
fmt.Println(slice)



}

