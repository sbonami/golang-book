An Introduction to Programming in Go
====================================

Notes and code sampesl for the first meeting of the Squaremouth book club.
Covering Chapters 1 - 5 of the go lang book.

Chapter Notes
-------------

### Chapter 1

Nothing of note in this chapter.

### Chapter 2 - Your First Program

Finally get to a ["Hello World"](chapter2/hello.go).
Briefly explains packages ("main"), importing, and main.
Some of the simplest types are broken down.

### Chapter 3 - Types

Starts of by explaining types in terms of philosophers and
mathematics (man after my own heart).

#### Integers

* `uint8`, `uint16`, `uint32`, `uint64`,`int8`,`int16`,`int32`,`int64`
* Alias types: `byte` == `uint8` and `rune` == `int32`
* Machine specific: `uint`, `int`, `uintpr`

#### Non-integer Numbers

* `float32` and `float64` (increase in precision with bits).
* `NaN`
* `∞` - (OPTION + 5)
* `complex64` and `complex128`

#### Strings

* Can be created with "" - Single line strings that allow escape characters
* Can be created with `` - Multi-line strings but no escape characters
* `len("String")` - String length
* `"Sample"[0]` - String indexing, does not return what you'd expect

#### Booleans

* `true`, `false` - literals
* `&&`, `||`, `!` - operations

### Chapter 4 - Variables

Now we start looking at how we can declare variables in go. There are a few
different ways to do the same thing.

* `var x int = 8` - declaring and assigning in one line
* `var x int` - only declaring followed by `x = 18` later on
* `x := 8` - assigning and inferring the type
* `const stuff string` - Declaring a constant variable, can't be re-assigned

### Chapter 5 - Control Structures

Can almost write a useful program now.

#### Loops

All loops are written in terms of `for`

* `for stuff == true { ... }` - While loop equivalent
* `for i := 1; i <= 5; i += 2 { ... }` - Traditional for-loop
* Yet to be seen for-each style loop

#### Conditionals

Standard conditionals using real english words.

* `if`, `else if`, `else` - Main conditional statements
* `switch` statement also exists

Discussion Points
-----------------

* Capital lettered function/method names. ex)

  ```go
  fmt.Println("Hello World")
  ```

* The `go` tool and the `godoc` command. Run `go help` to see what it can do.
* Go has a metric ton of types. Is this worth it?
* You can use UTF-8 characters, think about that. Show [gomega](http://onsi.github.io/gomega/).
* Go will not allow you to declare variables that aren't used (_ values)
