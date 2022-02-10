package techpalace

import (
    "strings"
    "fmt"
)
// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	var welcome string = "Welcome to the Tech Palace, " + strings.ToUpper(customer)
    return welcome
	panic("Please implement the WelcomeMessage() function")
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	msg := strings.Repeat("*",numStarsPerLine)+"\n" + welcomeMsg +"\n" + strings.Repeat("*",numStarsPerLine)
    
    return msg
    panic("Please implement the AddBorder() function")
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    var newString string = strings.ReplaceAll(oldMsg,"*", " ")
    // newString = strings.ReplaceAll(newString,"*", " ")
    newString = strings.TrimSpace(newString)
    fmt.Println("DEBUGINFO: " + newString)
    return newString 
	panic("Please implement the CleanupMessage() function")
}
