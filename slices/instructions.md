Introduction
Slices
Slices in Go are similar to lists or arrays in other languages. They hold a number of elements of a specific type (or interface).

Slices in Go are based on arrays. Arrays have a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.

A slice is written like []T with T being the type of the elements in the slice:

var empty []int                 // an empty slice
withData := []int{0,1,2,3,4,5}  // a slice pre-filled with some data
You can get/set an element at a given zero-based index using square-bracket notation:

withData[1] = 5
x := withData[1] // x is now 5
You can create a new slice from an existing slice by getting a range of elements, once again using square-bracket notation, but specifying both a starting (inclusive) and ending (exclusive) index. If you don't specify a starting index, it defaults to 0. If you don't specify an ending index, it defaults to the length of the slice.

newSlice := withData[2:4] // newSlice == []int{2,3}
Multiple Return Values
Go functions and methods can return multiple values. Very often, a second return value is used to return an error. For example:

func GetCard() (Card, error) { ... }
The assignment for multiple return values just uses a comma to separate the variables:

card, err := GetCard()
If statements can use an initializer before the condition separated by a semicolon. This is a common idiom seen for error handling:

if card, err := GetCard(); err != nil {
    // handle the error
}
Instructions
As a magician-to-be, Elyse needs to practice some basics. She has a stack of cards that she wants to manipulate.

To make things a bit easier she only uses the cards 1 to 10.

Task 1
Retrieve a card from a stack

Return the card at position index from the given stack.

card, ok := GetItem([]int{1, 2, 4, 1}, 2)
fmt.Println(card)
// Output: 4
fmt.Println(ok)
// Output: true
If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to return 0 and false:

card, ok := GetItem([]int{1, 2, 4, 1}, 10)
fmt.Println(card)
// Output: 0
fmt.Println(ok)
// Output: false

Stuck? Reveal Hints
Opens in a modal
Task 2
Exchange a card in the stack

Exchange the card at position index with the new card provided and return the adjusted stack. Note that this will also change the input slice which is ok.

index := 2
newCard := 6
SetItem([]int{1, 2, 4, 1}, index, newCard)
// Output: []int{1, 2, 6, 1}
If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to append the new card to the end of the stack:

index := -1
newCard := 6
SetItem([]int{1, 2, 4, 1}, index, newCard)
// Output: []int{1, 2, 4, 1, 6}

Stuck? Reveal Hints
Opens in a modal
Task 3
Create a stack of cards

Create a stack of given length and fill it with cards of the given value. If the given length is negative or zero, return an empty stack.

PrefilledSlice(8, 3)
// Output: []int{8, 8, 8}

Stuck? Reveal Hints
Opens in a modal
Task 4
Remove a card from the stack

Remove the card at position index from the stack and return the stack.

RemoveItem([]int{3, 2, 6, 4, 8}, 2)
// Output: []int{3, 2, 4, 8}
If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to leave the stack unchanged:

RemoveItem([]int{3, 2, 6, 4, 8}, 11)
// Output: []int{3, 2, 6, 4, 8}

Stuck? Reveal Hints
Opens in a modal
How to debug
When a test fails, a message is displayed describing what went wrong and for which input. You can also use the fact that console output will be shown too. You can write to the console using:

import "fmt"
fmt.Println("Debug message")
Note: When debugging with the in-browser editor, using e.g. fmt.Print will not add a newline after the output which can provoke a bug in go test --json and result in messed up test output.