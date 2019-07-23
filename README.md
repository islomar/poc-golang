# Playground for Go language
The goal of this repository is to concentrate my learnings about the Go language.

## General information
* Go is an open source project developed by a team at Google and many contributors from the open source community:
 * [https://en.wikipedia.org/wiki/Go_(programming_language)](https://en.wikipedia.org/wiki/Go_(programming_language))
* It was officially born on 10.11.2009
* Go v1 was released in March 2012
* Current version: 1.12 (March 2019), https://golang.org/doc/go1.12
* Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.


## Characteristics
* Compiles
* Statically typed
* Clean syntax
* Simple type system
* Concurrent primitives
* Rich standard library
* Great tools
* Open source

## The language
* Goroutines are lightweight threads that are managed by the Go runtime.
* To run a function in a new goroutine, just put "go" before the function call.
 * **Packages**: functions which start with capital letter, can be exported and reused.
* By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.
* In Go, a name is exported if it begins with a capital letter.
* In a function, the type comes after the variable name.
* Go's declaration syntax: [https://blog.golang.org/gos-declaration-syntax](https://blog.golang.org/gos-declaration-syntax)
* When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
* A function can return any number of results: `func swap(x, y string) (string, string)`
* Go's return values may be named. If so, they are treated as variables defined at the top of the function.
* A return statement without arguments returns the named return values. This is known as a "naked" return.
 * Naked return statements should be used only in short functions. They can harm readability in longer functions.


## Setup and CLI
* You have to configure a $GOPATH variable pointing to your workspace.
* Append to your PATH: $GOPATH/bin
* If you keep your code in a source repository somewhere, then you should use the root of that source repository as your base path. For instance, if you have a GitHub account at github.com/islomar, that should be your base path.
* `go install`: to compile and install (under /bin if the functions are not reusable, under /pkg if they are)
 * After compiling the go file, it creates an executable file under /bin folder
* `go build`: compile the Go file
* `go test`: execute method **Test** inside **_test.go** file.
* `go get`: remote package management

## Testing
* File with **_test** suffix.

* Goroutines: light-weight processes

## Packages
* Use capital letters in functions, so that they can be reused.
* Use `go install` to compile and install a package (something that can be reused): it will create a /pkg subfolder.

## Channels
* Channels are typed conduits for synchronization and communication between goroutines.
* They're a versatile and expressive means of modelling concurrent processes.
* It's a kind of pipeline.

## Sync
* Channels are great, but sometimes other concurrency mechanisms are a better fit.
* The sync package provides mutexes, condition variables and more useful primitives.

## net/http
* The net/http package provides an HTTP client.
* The client handles HTTP Keep-Alive using a pool of connections, by default.

## html/template
* It doesn't do automatic escaping

## flag
The flag package provides a simple API for parsing command-line flags.

## expvar
* You can monitor your app using expvar.
* Export variables via an HTTP handler registered at /debug/vars (http://localhost:8080/debug/vars)

## Static analysis
* gofmt for formatting


## How to write Go code
* https://golang.org/doc/code.html
* `go get` will fetch, build, and install it automatically under `$GOPATH/src`
* `go install`: This command builds the hello command, producing an executable binary. It then installs that binary to the workspace's bin directory.


## A Tour of Go
* Go's return values may be named. If so, they are treated as variables defined at the top of the function.
    - A return statement without arguments returns the named return values. This is known as a "naked" return.
    - Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
* A `var` statement can be at package or function level.   
* Outside a function, every statement begins with a keyword (var, func, and so on) and so the `:=` construct is not available.   
* Basic types: 
    - https://tour.golang.org/basics/11
    - The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. 
* The expression `T(v)` converts the value v to the type T.

### Flow control
* Go has only one looping construct, the `for` loop.
* The `while` is done with a for continued: `for sum < 1000 {`
* Forever: `for {`
* Like for, the if statement can start with a short statement to execute before the condition.
* Variables declared inside an if short statement are also available inside any of the else blocks.
* Switch:
    - Go only runs the selected case, not all the cases that follow. 
    - In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. 
    - Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.
    - Switch without a condition is the same as switch true. This construct can be a clean way to write long if-then-else chains.
* A `defer` statement defers the execution of a function until the surrounding function returns.
    - The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
    - Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.


### Pointers
* Go has pointers. A pointer holds the memory address of a value.
* The type `*T` is a pointer to a T value. Its zero value is nil
* The `&` operator generates a pointer to its operand.
* The * operator denotes the pointer's underlying value.
* This is known as "dereferencing" or "indirecting".
* Unlike C, Go has no pointer arithmetic.

### Structs
A struct is a collection of fields.



## Interesting links

### Videos
* Writing, building and testing Go code: [https://www.youtube.com/watch?v=XCsL89YtqCs]()
* Andrew Gerrand - Go: a simple programming environment - [https://vimeo.com/69237265](https://vimeo.com/69237265) >> slides: [https://speakerdeck.com/railsberry/go-a-simple-programming-environment-by-andrew-gerrand](https://speakerdeck.com/railsberry/go-a-simple-programming-environment-by-andrew-gerrand)


### Books and tutorials
* A Tour of Go: [https://tour.golang.org/list](https://tour.golang.org/list)
* [Best practices](https://peter.bourgon.org/go-best-practices-2016/)
* [https://golang.org/doc/](https://golang.org/doc/)
* An introduction to programming in Go: [https://www.golang-book.com/books/intro](https://www.golang-book.com/books/intro)
* The way to Go: [http://www.amazon.com/The-Way-To-Introduction-Programming/dp/1469769166/ref=sr_1_1?ie=UTF8&qid=1377988316&sr=8-1&keywords=the+way+to+Go](http://www.amazon.com/The-Way-To-Introduction-Programming/dp/1469769166/ref=sr_1_1?ie=UTF8&qid=1377988316&sr=8-1&keywords=the+way+to+Go)
* The Go Programming Language: [http://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440/ref=pd_sim_14_4?ie=UTF8&dpID=41%2BNlBKiHXL&dpSrc=sims&preST=_AC_UL160_SR111%2C160_&refRID=0897WKSK8BQ75AVPXKJ3](http://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440/ref=pd_sim_14_4?ie=UTF8&dpID=41%2BNlBKiHXL&dpSrc=sims&preST=_AC_UL160_SR111%2C160_&refRID=0897WKSK8BQ75AVPXKJ3)
* The little book of Go: http://raulexposito.com/documentos/go/

### General
* https://github.com/go-lang-plugin-org/go-lang-idea-plugin/wiki/v1.0.0-Setup-initial-project

### To read
* How to create a REST API in Golang: https://blog.friendsofgo.tech/posts/como_crear_una_api_rest_en_golang/
* https://github.com/slok/goresilience

### TDD and Goland
* https://quii.gitbook.io/learn-go-with-tests/