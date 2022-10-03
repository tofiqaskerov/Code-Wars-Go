//  1) Make a function that will return a greeting statement that uses an input; your program should return, "Hello, <name> how are you doing today?".

package main

import (
	"fmt"
	"strconv"
	"strings"
)

var name = "Tofiq"

func main() {
	Greet(name)
	GetGrade(50, 80, 90)
	AbbrevName("Askerov Tofiq")
	NumberToString(50)
}

// 1) Make a simple function called greet that returns the most-famous "hello world!".
func Greet(name string) {
	fmt.Printf("Hello, %s how are you doing today?", name)
}

// 2)  Complete the function so that it finds the average of the three scores passed to it and returns the letter value associated with that grade.

func GetGrade(a, b, c int) {
	middleValue := float32((a + b + c)) / 3
	switch {

	case middleValue >= 60 && middleValue < 70:
		fmt.Print(rune(68))
	case middleValue >= 70 && middleValue < 80:
		fmt.Print(rune(67))
	case middleValue >= 80 && middleValue < 90:
		fmt.Print(rune(66))
	case middleValue >= 90 && middleValue < 100:
		fmt.Print(rune(65))
	default:
		fmt.Print(rune(70))
	}

	fmt.Printf("%v %T \n", middleValue, middleValue)
	fmt.Print(rune(65))
}

/* 3) Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.

The output should be two capital letters with a dot separating them.

It should look like this:

Sam Harris => S.H

patrick feeney => P.F */

func AbbrevName(name string) string {

	nameSplit := strings.Split(name, " ")
	var firstIndex, lastIndex = nameSplit[0][0], nameSplit[1][0]
	return fmt.Sprintf("%s.%s", strings.ToUpper(string(firstIndex)), strings.ToUpper(string(lastIndex)))
}

/* 4) We need a function that can transform a number (integer) into a string.

What ways of achieving this do you know? */

func NumberToString(n int) string {
	str := strconv.Itoa(n)
	fmt.Printf("%T", str)
	return str
}
