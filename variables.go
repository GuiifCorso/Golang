//Every Go code should start with the "package", which will "store" your content in a pack that can be acessed between files.
package main

//Here we are import the fmt package, that includes useful print and input functions used in the code.
import (
	"fmt"
)

//As the package name in main, the principal function must be "func main()" as a reference to start the code. All your principal code goes inside main().
func main() {
  //There are two main ways to declare a variable inside a function. The first is with :=, that set the implicit type of a variable value.
	intNum := 2 //Here, the implicit value for 2, is int, which means "integer". Default integers value is 0
	floatNum := 1.5 //For 1.5, it's float64. Default float64 value is 0.0
	boolValue := true //For true, it's Boolean (True/False). Default bool is false
	stringText := "Hello world" //For "hello world" it's string. Default string is "", or nothing.

  //Above we are using the fmt package, with the function .Printf, which means "print formatted". It's useful when you need to pass variables values inside a print statement.
  //Here, when we use %d, %s... we are telling the Printf that in those locations is going to have variables, which we will pass after the main text
	fmt.Printf("Int: %d, Float: %.2f, Boolean: %t, String: %s", intNum, floatNum, boolValue, stringText) //As you can see, we have to pass the variables in the order we want it to be printed, so intNum goes first, then floatNum and ...
  //In the Printf, we have specific letters that goes after the % to say to the function what the type of the variable we are sending is. In the code we wrote, we can see that %d goes for int, %.2f goes for the float64 with 2 decimal after the ., %t for the bool and %s for strings.

  //The other way to declare a variable, is to set the type by yourself. This is an explicit declaration
	var name string //Here we set the var "name" as a string
	var age int //var "age" as int
	var favNum int //var "favNum" as int

  //As age and favNum are the same type, we could turn this into one line: (var age, favNum int)

  
	fmt.Print("\nHey, what's your name? ") //Function in fmt that just prints as the other languages, nothing special in this one. The "\n" in the beggining is used to go to the next line.
	fmt.Scanln(&name) //Scan function works like input in other languges like python. It reads the value the user writes and store it in somewhere. The store location will be explained next.

  //That "&" in front of the variable name is what we call a pointer. Simply talking, what & does with "name", is getting the store location of the variable. You can see that by printing &name. It will return the location where "name" is stored
  //Once we (the scan function) get the location and the user writes a response to the input, Scan will get the input value, and pass it to the location of "name", changing the variable value.
  //If we just wrote "fmt.Scan(name)" it returns an error, because the Scan function won't have an "adress" to store the user's input.

	fmt.Print("\nAnd, what's you age? ")
	fmt.Scanln(&age)

	fmt.Print("\nTell me something, what's your favorite number? ")
	fmt.Scanln(&favNum)

	fmt.Printf("Oh, now I see! Your name is %s, you are %d years old, and you favorite number is %d", name, age, favNum)
}
