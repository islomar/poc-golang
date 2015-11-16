# Playground for Go language
The goal of this repository is to concentrate my learnings about the Go language.

## General information
* Go is an open source project developed by a team at Google and many contributors from the open source community:
 * [https://en.wikipedia.org/wiki/Go_(programming_language)](https://en.wikipedia.org/wiki/Go_(programming_language))
* It was born on 10.11.2009
* Go v1 was released in March 2012
* Current version: 1.5 (August 2015)

### Go 1.5 [https://golang.org/doc/go1.5](https://golang.org/doc/go1.5)
* The compiler and runtime are now written entirely in Go (with a little assembler). C is no longer involved in the implementation, and so the C compiler that was once necessary for building the distribution is gone. The only C source left in the tree is related to testing or to cgo.
* The garbage collector is now concurrent and provides dramatically lower pause times by running, when possible, in parallel with other goroutines.
* By default, Go programs run with GOMAXPROCS set to the number of cores available; in prior releases it defaulted to 1.
* Support for internal packages is now provided for all repositories, not just the Go core.
* The go command now provides experimental support for "vendoring" external dependencies.
* A new go tool trace command supports fine-grained tracing of program execution.
* A new go doc command (distinct from godoc) is customized for command-line use.


## Characteristics
* Compiles
* Statically typed
* Clean syntax
* Symple type system
* Concurrent primitives
* Rich standard library
* Great tools
* Open source

## The language
* Goroutines are lightweight threads that are managed by the Go runtime.
* To run a function in a new goroutine, just put "go" before the function call.
 * **Packages**: functions which start with capital letter, can be exported and reused.

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

** Goroutines
* light-weight processes

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


## Interesting links

### Videos
* Writing, building and testing Go code: [https://www.youtube.com/watch?v=XCsL89YtqCs]()
* Andrew Gerrand - Go: a simple programming environment - [https://vimeo.com/69237265](https://vimeo.com/69237265) >> slides: [https://speakerdeck.com/railsberry/go-a-simple-programming-environment-by-andrew-gerrand](https://speakerdeck.com/railsberry/go-a-simple-programming-environment-by-andrew-gerrand)


### Books and tutorials
* A Tour of Go: [https://tour.golang.org/list](https://tour.golang.org/list)
* [https://golang.org/doc/](https://golang.org/doc/)
* An introduction to programming in Go: [https://www.golang-book.com/books/intro](https://www.golang-book.com/books/intro)
* The way to Go: [http://www.amazon.com/The-Way-To-Introduction-Programming/dp/1469769166/ref=sr_1_1?ie=UTF8&qid=1377988316&sr=8-1&keywords=the+way+to+Go](http://www.amazon.com/The-Way-To-Introduction-Programming/dp/1469769166/ref=sr_1_1?ie=UTF8&qid=1377988316&sr=8-1&keywords=the+way+to+Go)
* The Go Programming Language: [http://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440/ref=pd_sim_14_4?ie=UTF8&dpID=41%2BNlBKiHXL&dpSrc=sims&preST=_AC_UL160_SR111%2C160_&refRID=0897WKSK8BQ75AVPXKJ3](http://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440/ref=pd_sim_14_4?ie=UTF8&dpID=41%2BNlBKiHXL&dpSrc=sims&preST=_AC_UL160_SR111%2C160_&refRID=0897WKSK8BQ75AVPXKJ3)


### General
* https://github.com/go-lang-plugin-org/go-lang-idea-plugin/wiki/v1.0.0-Setup-initial-project