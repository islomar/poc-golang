# Annotations from Codely.tv course about Go
@adrianpgl
@joanjan14

## Go in 15 minutes
* https://www.youtube.com/watch?time_continue=8&v=5QokUwp99oY
* Implicit interfaces
* No inheritance: only composition
* Errors instead of Exceptions
* Safe pointers, no arithmetic operations on pointers
* Type alias: to define new types, reserved word `type`

## Introduction to Go
* Frameworks
    - Buffalo
    - Martini
    - Gin
    - Revel
    - Beego

* Most used libraries
    - Testify (testing)
    - Logrus (log)
    - pkg/errors (errors)
    - cobra (cli)
    - godog (testing)
    
* APIs
    - Gorilla mux
    - Negroni
    - google/jsonapi
    - grpc
    - echo
    
* Web-scraping frameworks
    - Pholcus
    - go_spider
    - ants-go
    - colly
    - Dataflow Kit
    
* To follow
    - @golang
    - @golang_news
    - @golangweekly
    - @francesc
    - @bketelsen
    - @goingodotnet
    - @matryer
    - @golangbcn
    - @FriendsOfGo

* Cool features
    - `godoc -http=:6060`
    - `gofmt -d -w`
    - Cross-compile: with the same command and saying the architecture, you got it. There are libraries to be used in your CI/CD.
        - The binary generated is different for each architecture

## Development environment
* https://play.golang.org/
* Visibility
    - Upper case: public
    - Lower case: private
* It works with packages (like Python). Each folder is a package. Each package should be completely independent.
* Package import, three ways:
    - `"fmt`
    - `. "fmt`: dot imports >> `Println
    - `format fmg`: alias import, you can do `format.Println`
    - `_ "fmt`: to tell the compiler that we are not using directly in our code (transitive dependencies). E.g. when doing tests or using an ORM.
* `const`
* Entrypoint: 
    - `package main`
    - `func main()`
* How to declare a variable:
    - `var variableName variableType = variableValue` (usually when we want to asign a default initial value or no value at all)
    - `variableName := variableValue`  >> we let the compiler to infer the type (usually when we want to initialize the variable)
* Common types: `int`, `int32`, `int64`, `float64`, etc.
* `For` loops:
    - `for j:=0; j < 5; j++`
    - `for idx, n := range nums {`
* Functions
    - `func functionName() returntype { return "Hello World"`    
    - ```
        func functionName() (variableName returntype) { 
            msg= Hello World
            return 
       }
       ```   
       Implicit return, the variableName will be created in the scope of the function and then you can use it
* There are no classes but structs
* You can not have unused imports or variables (the compiler fails)
* Switch-case: use `fallthrough` to concatenate two cases
* Semicolons are optional (the compiler does it before compiling)
* `nil` is not a valid value for all the types: https://blog.friendsofgo.tech/posts/los_nil_seran_nil/

## Making it work
* `go build main.go`
* `go run main.go`: compile and run
* `go build -o <binaryFilename> main.go`: it generates a binary file for the architecture on the machine where it is run


## First steps: creating a project from scratch
* `$GOPATH`: bin (binaries), pkg (dependencies), src (source code)
* Go modules: https://blog.friendsofgo.tech/posts/go-modules-en-tres-pasos/


## WTFs
* Visibility public/private according to upper/lower case
* Implicit interfaces


## Links
* Online playground: https://play.golang.org/
* Blogs, slacks, meetups:
    * https://blog.friendsofgo.tech/
    * https://github.com/golang/go/wiki
    * https://gophers.slack.com/
    * https://go-meetups.appspot.com/
* https://gopherize.me/
* https://golang.org/doc/effective_go.html
* https://blog.golang.org/
* https://tour.golang.org/welcome/1
* Friends of Go:
    - Twitter: https://twitter.com/friendsofgotech
    - Blog: https://friendsofgo.tech
