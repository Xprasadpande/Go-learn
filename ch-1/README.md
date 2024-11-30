# Go Language Learning Repository

Welcome to my **Go Language** learning repository! This repository serves as a storage for my notes, examples, and exercises as I explore the fundamentals of Go.

## Chapter 1: Why Write Go?

Go, also known as Golang, was designed by Google engineers to address issues in software scalability, maintainability, and speed. Here are some of its standout features:

- **Simplicity:** The language emphasizes simplicity and readability.
- **Performance:** Compiled to machine code, Go provides the speed of C/C++.
- **Concurrency:** Goroutines and channels make concurrent programming intuitive and efficient.
- **Cross-platform:** Offers cross-platform portability.
- **Modern Tooling:** Provides excellent tools for testing, benchmarking, and documentation.

### Why Choose Go?

- Ideal for backend systems, CLI tools, and microservices.
- Large companies like Google, Uber, and Dropbox use it extensively.
- Growing community support and a strong standard library.

---

## Chapter 1: Variables and Basic Syntax

### Declaring Variables

Go supports explicit and implicit variable declarations.

```go
// Explicit Declaration
var a int = 10

// Implicit Declaration
b := 20

var x, y, z int = 1, 2, 3
// OR
i, j, k := 4, 5, 6


Zero Values
Variables in Go are initialized with zero values if not explicitly assigned.

int → 0
float64 → 0.0
string → "" (empty string)
bool → false


Multiple Declarations:

var x, y, z = 1, 2, 3
var (
    a int = 10
    b string = "Go"
)


----Concepts Covered

1. Why Write Go?
2. Variables and Declarations
```
