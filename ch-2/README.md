# Chapter 2: Functions and If-Else Statements

This chapter covers the fundamentals of **functions** and **if-else statements** in Go, along with some important theoretical concepts.

---

## Functions in Go

Functions are blocks of reusable code that perform specific tasks and help in modularizing your program.

### Key Points

1. **Definition:** Functions are declared using the `func` keyword.
2. **Parameters and Return Types:** Functions can accept multiple parameters and return multiple values.
3. **Modularization:** Functions make code cleaner and reusable.
4. **Exported Functions:** Functions starting with an uppercase letter (e.g., `Add()`) are accessible from other packages.

### Syntax

```go
func functionName(parameters) returnType {
    // Function body
    return value
}


--------------------------
package main

import "fmt"

// Function to multiply two numbers
func multiply(a int, b int) int {
    return a * b
}

func main() {
    result := multiply(4, 5)
    fmt.Println("Product:", result)  // Output: Product: 20
}


#If-Else Statements
If-else statements are conditional structures that help in decision-making based on conditions.

Key Points
>>>Single Condition: Use if to execute code when a condition is true.
>>>Alternative Execution: Use else for an alternative block of code if the condition is false.
>>>Multiple Conditions: Use else if to test additional conditions.
>>>Short Variable Declaration in If: Declare variables within an if block for concise conditions.

if condition {
    // Code if condition is true
} else if anotherCondition {
    // Code if anotherCondition is true
} else {
    // Code if all conditions are false
}


//shoet if else
if num := 10; num > 5 {
    fmt.Println("Number is greater than 5")
}

//nested if else

if a > b {
    if a > c {
        fmt.Println("a is the largest")
    }
} else {
    fmt.Println("b or c is larger")
}


//Return Multiple Values from Functions:

func swap(a, b int) (int, int) {
    return b, a
}

func main() {
    x, y := swap(10, 20)
    fmt.Println(x, y)  // Output: 20, 10
}


Defer Statement in Functions:
Use defer to delay execution of a statement until the function exits.


func main() {
    defer fmt.Println("This is printed last")
    fmt.Println("This is printed first")
}



#Concepts Covered

1. Defining and calling functions.
2. Parameters and return types in functions.
3. Using if, else if, and else statements.
4. Short and nested if-else structures.
5. Returning multiple values and using defer.
```
