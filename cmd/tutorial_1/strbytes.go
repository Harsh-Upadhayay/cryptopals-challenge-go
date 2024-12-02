package main

import (
	"fmt"
	"strings"
)

// Strings are immutable in go.
func main() {

	myStr := "résumé"
	ch2 := myStr[1] //rune(integer) first half (8 bits of 16 bits)
	fmt.Printf("%v, %T\n", ch2, ch2)

	// looping handles multibyte chars indexing automatically.
	for idx, val := range myStr {
		fmt.Println(idx, val) // (0,1,((3)),4,5,6) skips 2 (é is stored at 1 & 2 idx)
	}

	// cast manually to int arr, now each index contain a integer.
	// (rune is just an alias for int32)
	myStr2 := []rune("résumé")
	ch := myStr2[1]
	fmt.Printf("%v, %T\n", ch, ch)

	for idx, val := range myStr2 {
		fmt.Println(idx, val) // (0,1,2,3,4,5) each integer at a idx.
	}

	// array of strings (not letters/runes/ints)
	strSlice := []string{"a", "b", "c"}
	catStr := ""

	// Inefficient, coz each time a concatination happens a new string is made.
	for i := range strSlice {
		catStr += strSlice[i]
	}

	runSlice := []rune{97, 98, 99}
	val := ""

	// Same as above with additional overhead of casting rune to string.
	for i := range runSlice {
		val += string(runSlice[i])
	}

	fmt.Println(val)

	// internally allocates an array.
	var strBuilder strings.Builder
	for i := range strSlice {
		// appends data to array in O(1)
		strBuilder.WriteString(strSlice[i])
	}

	// builds string only after calling String() fn.
	finalStr := strBuilder.String()
	fmt.Printf("%v, %T\n", strBuilder, strBuilder)
	fmt.Printf("%v, %T\n", finalStr, finalStr)

}
