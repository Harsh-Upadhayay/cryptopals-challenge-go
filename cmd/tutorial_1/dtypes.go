package main // Has to be same across all files of this folder.
import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main1() {
	fmt.Println("Hello World!")

	// Staticly typed: Requries type to be known at compile time.
	var intNum int16 = 32761
	fmt.Println(intNum)

	var floatNum float32 = 23.2859121211212607948
	fmt.Println(floatNum)

	// Strongly typed: Need to cast manually for cross datatype operations.
	var result float32 = 2*floatNum + float32(intNum)
	fmt.Println((result))

	var myString string = "Hello" + `world!
	in next line, after a tab. γ`
	fmt.Println((myString))

	// len(): returns the num of bytes used, not utf8 chars use more than one byte (γ: 2)
	fmt.Println(len(myString), "!=", utf8.RuneCountInString(myString))

	// Shorthand.
	newVar := "text"
	tisVar := 23.521
	fmt.Println(newVar, tisVar)

	var1, var2 := 24, "54a"
	fmt.Println(var1, var2)

	// Const: similar to var, just unmutable.
	const x = "hg"
	fmt.Println(x)
	var myErr error = errors.New("test error")
	fmt.Println(myErr.Error())

}
