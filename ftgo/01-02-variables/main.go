package main

import "fmt"

// Global variable declared outside function scope
//
//	In global namespace, don't need to declare the type, directly assign
//	variable. Although go is statically type, type inference occurs.
//	Idomatic way is to only use `var` in the global scope
var globalVar = "Global variable"
var globalVarString string = "Global string variable"
var globalVarNoValue int

// Group variable declarations
var (
	firstName        = "Name"
	surname   string = "Surname"
	age              = 42
)

// Constants are declared in the global scope
//
//	Declared at top of package
//	Always presented in lowercase, camelCase or PascalCase
//	Exportable with a Captial Letter as first character
const newConstant = "Constant string"
const ExportableNewConst = "Exportable constant string"

func main() {
	// Local variables inside a function scope
	// `:=` automatically infers the type at declaration
	version := 1            // infer type int
	buildName := "Software" // infer type string
	buildNumber := 10.1     // infer type float32

	fmt.Println(version)
	fmt.Println(buildName)
	fmt.Println(buildNumber)

	// Don't use `var` to initialise values, prefer `:=` in function scope
	// Idomatic var declaration only for initalisation for non-assignment
	var deployment int

	// Every type is initialised with a default value
	fmt.Println(deployment)
	// => 0

	fmt.Println(newConstant)

	// go run main.go
}
