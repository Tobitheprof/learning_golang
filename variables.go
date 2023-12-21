package master

import (
	"fmt"
	"strconv"
)

// Variable declaration in Go

// Mrthod One(Global Scope)

var i int = 123 //  Global Scope Declaration

// Var Block Example
var (
	actorName string = "Johnny Cage"
)

// Method Two(Function Scope)
func main() {

	var j int = 23 // Manually Var Type
	k := "food"    // Automatically Var Type

	fmt.Println(i)

	fmt.Println(j)

	fmt.Printf(k)

	fmt.Printf("\n%v, %T", actorName, actorName) // Print Formatting (Initial args in the first quotes represent what exactly you want to print out to the console, type, value etc...)

	// Var conversion
	var z float32
	z = float32(i)

	fmt.Printf("\n%v, %T\n", z, z)

	var s string
	s = strconv.Itoa(i)

	fmt.Printf("%v,%T", s, s)

}

// NB vars can not be redeclared in the same scope, they can only have their values reassigned.
// Deeper var scopes are higher priority
// Vars in Go always have to be used
// Lower case var names have scopes withing the package level
// Uppercase var names have global scope
// Vars in code blocks are limited to the code block
// The name of a variable should reflect the life of a variable
// Acronyms should be in uppercase (primarily for readability)
