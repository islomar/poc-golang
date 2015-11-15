# Playground for Go language
The goal of this repository is to concentrate my learnings about the Go language.

## General information
* Go is an open source project developed by a team at Google and many contributors from the open source community.

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

## Setup
* You have to configure a $GOPATH variable pointing to your workspace.
* Append to your PATH: $GOPATH/bin
* If you keep your code in a source repository somewhere, then you should use the root of that source repository as your base path. For instance, if you have a GitHub account at github.com/islomar, that should be your base path.
* `go install`: to compile and install (under /bin if the functions are not reusable, under /pkg if they are)
 * After compiling the go file, it creates an executable file under /bin folder
* `go build`: compile the Go file


## Packages
* Use capital letters in functions, so that they can be reused.
* Use `go install` to compile and install a package (something that can be reused): it will create a /pkg subfolder.

## Interesting links

### Videos
* Writing, building and testing Go code: [https://www.youtube.com/watch?v=XCsL89YtqCs]()
* Andrew Gerrand - Go: a simple programming environment - [https://vimeo.com/69237265](https://vimeo.com/69237265) >> slides: [https://speakerdeck.com/railsberry/go-a-simple-programming-environment-by-andrew-gerrand](https://speakerdeck.com/railsberry/go-a-simple-programming-environment-by-andrew-gerrand)


### Books and tutorials
* [https://golang.org/doc/](https://golang.org/doc/)
* An introduction to programming in Go: [https://www.golang-book.com/books/intro](https://www.golang-book.com/books/intro)
* The way to Go: [http://www.amazon.com/The-Way-To-Introduction-Programming/dp/1469769166/ref=sr_1_1?ie=UTF8&qid=1377988316&sr=8-1&keywords=the+way+to+Go](http://www.amazon.com/The-Way-To-Introduction-Programming/dp/1469769166/ref=sr_1_1?ie=UTF8&qid=1377988316&sr=8-1&keywords=the+way+to+Go)
* The Go Programming Language: [http://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440/ref=pd_sim_14_4?ie=UTF8&dpID=41%2BNlBKiHXL&dpSrc=sims&preST=_AC_UL160_SR111%2C160_&refRID=0897WKSK8BQ75AVPXKJ3](http://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440/ref=pd_sim_14_4?ie=UTF8&dpID=41%2BNlBKiHXL&dpSrc=sims&preST=_AC_UL160_SR111%2C160_&refRID=0897WKSK8BQ75AVPXKJ3)


### General
* https://github.com/go-lang-plugin-org/go-lang-idea-plugin/wiki/v1.0.0-Setup-initial-project