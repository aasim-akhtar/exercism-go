/* CONCEPTS LEARNED:
	SLICES
	ITERATION OF SLICE
	MAKE() TO INITIALIZE SLICE WITH USER/VARIABLE INPUT
	APPEND()
	strconv.Itoa(e)
	... parameter  slice = append(slice[:index], slice[index+1:]...)
	_ The blank identifier, not implemented
	for loop
*/
package cards
import (
    "fmt"
    "strconv"
)
// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
    if index<len(slice) && index >=0 {
        return slice[index],true
    } else {
    	return 0,false
    }
    
	panic("Please implement the GetItem function")
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
    if index >=0 && index < len(slice) {
        slice[index] = value
    }else {
    	slice = append(slice, value) 
    }
    return slice
	panic("Please implement the SetItem function")
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
    if length <0 {
        var slice []int
        return slice
    }else {
    	var slice = make([]int,length)
		for e := 0 ; e < length; e++{
        fmt.Println("accessing " + strconv.Itoa(e) + " index")
        slice[e] = value
    }
	return slice
    }

	panic("Please implement the PrefilledSlice function")
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if(index < len(slice) && index >=0){
            slice = append(slice[:index], slice[index+1:]...)
    }
    return slice
	panic("Please implement the RemoveItem function")
}
