# Annotations from Codely.tv course about Go
* People
    - @adrianpgl
    - @joanjan14
* Example source code: https://github.com/CodelyTV/golang-introduction


## General
* Go is compiled
* Multiplatform
* Single thread

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
    - `"fmt` >> `fmt.Println()`
    - `. "fmt`: dot imports >> `Println()`
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
    - Now the dependencies come natively, we do not need `$GOPATH` to create the project
    - Go modules does not come activated by default in `$GOPATH`: you need to set `GO111MODULE=on`
* `go mod init`: inside an empty folder, to initialize the module (a file `go.mod` gets generated)
    - `go.mod` should never be modified manually
* `go mod tidy`: download dependencies
    - it generates `go.sum` if it does not exist with the dependencies versions and integrity hashes
    - no more need to run `go get ./...`
    - downloaded libraries use semantic version
    - it updated `go.mod` with ALL the possible combinations of OS, architecture and build tags.
* `go build` and `go test` update the `go.mod` file with the required packages
* Optional: run `go mod vendor` if you want to create a backwards-compatible `vendor` folder.
* **Go flags**
    - https://golang.org/pkg/flag/
    - https://blog.friendsofgo.tech/posts/crear-tu-primer-cli-en-go/
    - `flag.NewFlagSet()`, `flag.String()`, `flag.NArg()`, `flag.Arg(0)`
* To access a variable which is a pointer, use an asterisk: `*variableNAme`
* Declare a pointer to a variable/struct:  `&variableName`
* **Cobra** library
    - https://github.com/spf13/cobra
    - Library for CLI (the Docker client is developed with Cobra)
    - Two parts: the library and the command line
    - Install the library: `go get -u github.com/spf13/cobra/cobra`
    - To create the project: `cobra init`
    - To add new commands: `cobra add <command_name>`

     
## Data structures in Go
* The array size defines its type
* A slice is a data structure that represents an underlying array: it contains
    - a pointer to the underlying array
    - length of the underlying array (what it is actually used)
    - capacity
* Creation of a slice: `slices := make([]string, 3)`
* Usually you work with slices, unless you know the size.
* Reading files
    - Simple: `bufio`, `os`...
    - Rich: `encoding/json`, `encoding/xml`...


## Types
* struct:  son estructuras de datos formadas por listados de atributos caracterizados por un nombre y un tipo
    - `var page struct { ... }`
* Defined alias = type:  you can create types from other types (struct, string, interface, func, channel, etc.)
* There is no `Enum` type, but a convention like https://play.golang.org/p/fO9qEoBCdWg:
```
type Weekday int

const (
	Monday    Weekday = 0
	Tuesday   Weekday = 1
	Wednesday Weekday = 2
	Thursday  Weekday = 3
	Friday    Weekday = 4
	Saturday  Weekday = 5
	Sunday    Weekday = 6
)

func (d Weekday) String() string {
	names := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	return names[d]
}
```

## Entities and Repositories
* https://github.com/CodelyTV/golang-introduction/blob/master/04-modeling_data/internal/beer.go
* In Go usually no dependency injection libraries are used, though it exists some like [Wire](https://github.com/google/wire)
    - https://blog.friendsofgo.tech/posts/gestion-de-dependencias-en-golang/
* NewRepository 
```
type repository struct {
}

func NewRepository() beerscli.BeerRepo {
   	return &repository{}
}
```
    - `&repository{}` is the way to initialize an empty repository struct. The & means that we are returning a pointer to the struct 
* `Unknown BeerType = iota`, `iota` means that we want to assign incremental values
    - https://github.com/CodelyTV/golang-introduction/blob/master/04-modeling_data/internal/beer.go#L17
    - `Unknown` means that it would start with 0
* The "complex" types are defined as pointers, to be able to pass them by reference and modified them, e.g. BeerType here:
```
type Beer struct {
	ProductID int
	Type      *BeerType
```
* Because Go packages must be self-contained, the repository interface declaration goes in the same file than the type . Only the interface implementation would go in a different file.


## HTTP requests
* `net/http` package
* Public Beers API for playing around: http://ontariobeerapi.ca/beers/
* Other interesting public APIs for trying:
    - https://pokeapi.co 
    - https://swapi.co 
    - https://api.github.com 
    - https://punkapi.com
* `UnmarshalJSON()` is the name which should have the struct function for unmarshalling in a specific way: https://github.com/CodelyTV/golang-introduction/blob/master/05-parsing_http_response/internal/beer.go#L65
* Declare **struct tags** for matching the entity and repository data: https://github.com/CodelyTV/golang-introduction/blob/master/05-parsing_http_response/internal/beer.go#L9
* **There is no inheritance** in Go.
* There is both composition for structs and for interfaces.


## Error handling
* https://github.com/CodelyTV/golang-introduction/tree/master/06-error_handling
* There are no exceptions in Go, but **errors**.
* `log.Fatal(err)`: it prints the error and a System exit with status != 0 and stops the execution (but it does not propagate an error up).
* `panic(err)`: it propagates the error up. It can be captured with the clause `defer func()`
    - Do NOT use `panic()` usually.
* To create an error: `errors.New("This is the error description")`
* When you want to return an error, it must be the last argument.
* Go proverbs by Rob Pike:
    - https://go-proverbs.github.io/
    - https://www.youtube.com/watch?v=PAAkCSZUG1c
* Three strategies to handle errors:
    - *Sentinel errors*: 
        - when using specific errors to decide what to do
        - `errors.New("Ooooppss!!")`
        - discouraged in general because from the outside you can couple to it
    - *Custom errors*
        - it's about adding more context to the error
        - create a struct with a function `Error()` which returns a string
        - and then check the error type...
        ```
            if err, ok := err.(*BeerNotFoundErr); ok {
                log.Fatal(err)
            }	
        ```
    - *Opaque errors*
        - The less you know about the error, the better.
        - `"github.com/pkg/errors"`
        - `errors.Wrapf(err, "Something bad happened: %s", file)`: we wrap it when it is an error from an external library of from the Go core
        - if it is our own error, we return it just like this `errors.Errorf("my error")`
* The "recommended way":
    - https://github.com/CodelyTV/golang-introduction/blob/master/07-behaviour_error_handling/internal/errors/errortypes.go
    - Handle errors based on their behaviour
    - private struct with wrapper methods and `func IsXxxErr(err error) bool`
    - to "hide" the error to the outside. You get decoupled from the error...


## Automated testing
* https://github.com/CodelyTV/golang-introduction/tree/master/08-automated_tests
* define under `package xx_test`, so that you don't have access to private methods
* `func TestXxxx()`
* test files like `xxxx_test.go` under same folder (though it could be all together under a `tests` folder)
* Go does not compile the tests, it is something independent when running `go test xx`
* Instead of asserting, we fail if something did not go as expected. Use `testify` for asserting
* Run: `$ go test ./...`
* Very common to use mocks from `testify`
* Data providers
    - Table driven tests
    - https://github.com/CodelyTV/golang-introduction/blob/master/08-automated_tests/internal/fetching/service_test.go#L15
    - `go test ./... -run=TestFetchByID/"valid beer"`, to run only one test
* Mocks:
    - Usually you create it on your own.
    - The trainer usually uses the mocking from `testify`
    - Another option: https://github.com/matryer/moq, you need to manually generate it from an interface (it generates a file)
    - A very interesting one: https://github.com/vektra/mockery


## Debugging errors
### Profiling
* https://github.com/CodelyTV/golang-introduction/blob/master/10-profiling/cmd/beers-cli/main.go
    - `go tool pprof <app_to_be_run> <profiling_filename>`, e.g. beers.mem.prof
    - Execute a command, e.g. `top 5` to check the top 5 functions which have whatever resource consumption we are measuring (memory, CPU time, etc.)
    - Execute `web` to show the critical path and visual execution.
* Amdahl's law: https://en.wikipedia.org/wiki/Amdahl%27s_law
* We can profile as HTTP as well: https://golang.org/pkg/net/http/pprof/
* https://blog.golang.org/profiling-go-programs
* To read and understand the files generated:
    - go tool pprof
    - go-torch

### Benchmarking
* https://github.com/CodelyTV/golang-introduction/tree/master/09-benchmarking
* https://github.com/CodelyTV/golang-introduction/blob/master/09-benchmarking/internal/storage/ontario/repository_test.go
* Once we have updated our code for adding the benchmarking (e.g. in a test file), we can run `go test -run=GetBeers -bench=. > bench.old`
* `go test -bench=.`
* To measure which implementation is better:
    - `go get -u golang.org/x/tools/cmd/benchcmp`
    - `go test -bench=. > bench.old`
    - `go test -bench=. > bench.new`
    - `benchcmp bench.old bench.new`
* The library `jsoniter` is much more optimal for un/marshalling json.


## Concurrency and parallelism
* https://github.com/CodelyTV/golang-introduction/tree/master/11-sharing_memory_concurrency
* Shared memory:
    - Maybe needed when dealing with big data
    - https://golang.org/pkg/sync/#Mutex
    - Example: https://github.com/CodelyTV/golang-introduction/blob/master/11-sharing_memory_concurrency/internal/fetching/service.go#L45
* Message sending:
    - Use [copy](https://golang.org/pkg/builtin/#copy) when dealing with slices, in order to not share the array.
* **Don’t communicate by sharing memory, share memory by communicating.**

## General questions
### Go project structure
* Having an `internal` folder is quite common: everything inside there will not be visible from the outside
* https://blog.friendsofgo.tech/posts/como_estructurar_tus_aplicaciones_go/
* sadfasf
    - `cmd`: entrypoints, client actions.
    - `pkg`: everything that could be reused from the outside
    - `internal`: it is protected, only for our own stuff, not visible.

### Go mod + init
* Vendoring: having all the dependencies inside our project, under `/vendor`
    - Useful for example in a CI tool, for performance reason: instead of having to download each time all the dependencies, they are already in the project.
* With Go modules, you don't need the vendoring: 
    - https://blog.friendsofgo.tech/posts/go-modules-en-tres-pasos/
* `init()` will be executed each time the package where it's contained it is imported
    - not recommended 

### Next steps
* REST API: https://blog.friendsofgo.tech/posts/como_crear_una_api_rest_en_golang/


## Exercises
* Error handling: https://pro.codely.tv/library/introduccion-a-go-tu-primera-app/89042/path/step/59271607/
* Testing: https://pro.codely.tv/library/introduccion-a-go-tu-primera-app/89042/path/step/59282342/


## WTFs
* Visibility public/private according to upper/lower case
* Looks like not differentiating "A has same type of B" and "A is compound of B"
* single letters as variable names: bad readability
* repeated block for handling err...
* Implicit interfaces
* The array size defines its type
* creating an interface for the service... just because, looks like a convention, even if you don't have another implementation
* No Enum
* test 
    - files like `xxxx_test.go` under same folder (though it could be all together under a `tests` folder)
    - Instead of asserting, we fail if something did not go as expected. Use `testify` for asserting
* Dependency management: vendoring or Go module, very inmature...
* If with a short statement: bad readability
* Naked returns: bad readability
* constants with camel case
* switch/case not indented: bad readability


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

### People to follow
* Dave Cheney
* [Francesc Campoy] (https://twitter.com/francesc)
* [Bill Kennedy] (https://twitter.com/goinggodotnet)
* [Jon Calhoun] (https://twitter.com/joncalhoun)
* [Mat Ryer] (https://twitter.com/matryer)


### Useful resources
* [Gophercises](https://gophercises.com/) - Muy útil para hacer ejercicios variados, próximamente tendremos el recurso en español por nosotros en Friends of Go
* [Go by example](https://gobyexample.com/)
* [Go Tour] (https://tour.golang.org/welcome/1)
* [Go 101] (https://go101.org/article/101.html)
* [Golang Programs] (http://www.golangprograms.com/)
* [Practical Go] (https://dave.cheney.net/practical-go) - Próximamente traducido al español por nosotros en Friends of Go


## To ask or research
* Equivalent to Hamcrest?