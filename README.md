# Go-Portfolio
Learning Go and building project along the way, This README will be my note taking place


## NOTES


## CHAPTER 1
1. Go compiles to a single native binary i.e. it does not require aother applications to run and is easy to distribute 
2. go vet - scans for coding mistakes
### 3. <b>Steps to create a Go project </b>
- mkdir project1 --- Make a project dir
- go mod init *project name if in local / repo name if initializing a project from git*
    - A Go project is called a module
    - Each module has a go.mod file, that contains module name, minimum supported go version to run the project and all the dependencies/libraries required in the project to run it. Just like req.txt in python.
    - DONT edit it manually, use go get *package name* & go mod tidy 
    - INSIDE A GO CODE
         ``` go
        package main

        import(
            "fmt"
        )
        func main(){
            Println("Hi my name is Ankit!")
        }
        ```
    - CANNOT import specific functions of a package, ONLY whole pkg imports allowed
    - main package is resposible to start/run the go program using the code under func main(), other packages can be present with different pkg name finally imported under main pkg
- go build - to get a executable single file for starting the program
    - go build -o *custom name for the executable file*

- go fmt - can be run to validate the code formatting, but not needed in VSCODE(as it does so autoatically)
    ``` go
    package main
    // SEMICOLON INSERTION RULE
    import(
        "fmt"
    )
    func main()
    {
        This is the wrong way to declare funciton, as the braces are below the func declaration, 
        So when semicolons are added by lexer during compilation (yes the developers arent required to do that)
        it will put the ; in front of func main();

        In vscode it will show error - unexpected semicolon or newline before { syntax

    }
    ```
- Makefile
    ``` makefile
    .DEFAULT_GOAL:= build
    .PHONY:fmt vet build

    fmt: 
        @echo "Formatting..."
        go fmt ./..
    vet: fmt
        go vet ./..
    build: vet
        go build
    ```
    - Each operation is called a target
    - "./.." means this dir and all sub dirs
    - stuff before : is name, stuff after : is the dependency, i.e. for vet:fmt when run it will ensure to run fmt command before it runs vet
    - .PHONY - ensures whatever written in it is treated as command, for if there s a dir with same name as command can be confusing

    - each target can be run separately using command make fmt, make vet, make build
    - if only make is provided without any target parameters, it will run the whole makefile in order

- Updating golang
```bash
Linux and BSD users need to download the latest version, move the old version to a
backup directory, unpack the new version, and then delete the old version:
$ mv /usr/local/go /usr/local/old-go
$ tar -C /usr/local -xzf go1.20.6.linux-amd64.tar.gz
$ rm -rf /usr/local/old-go
```

## CHAPTER 2 - Predeclared Types and Declarations

1. Predeclared Types - boolean, float, int, strings
2. 0 default value, any var declared but not assigned any value, makes it clearer
3. Literals - Hard coded values - 
    - Integer literal
    - Float literal
    - Rune Literal - 'single quote for rune' - used for unicode values such as newline \n etc....
    Strings
        - Interpreted String Literal - "double quote for strings, interprets unicode values such as new lines etc.."
        - Raw String - `Written in inverted commas, take whole string as it is, usually to write full SQL Queries`

    Literal Type Example Use

    Integer 10, 0b1010, 0xFF Counting, IDs, bit masks, file perms

    Float 3.14, 1e5 Math calculations, scientific values

    Rune 'a', '\\n', \\u03A9 Character processing, unicode

    String "hi", \`multi\\nline\` Logging, JSON, config, SQL, regex

    Bool true, false Conditions

    Nil nil Empty pointer/map/slice

    ``` go
    package main

        import "fmt"

        func main() {
            port := 8080               // integer literal
            pi := 3.1415               // float literal
            newline := '\n'            // rune literal
            path := "C:\\logs\\file"   // interpreted string
            sql := `SELECT * FROM users` // raw string
            isActive := true           // boolean literal

            fmt.Println(port, pi, newline, path, sql, isActive)
        }
    ```
4. Go unlike other langages doesnt have Automatic type conversion, manual conversion needs to be done that follows conversion rules 

5. var vs :=
    - The := operator can do one trick that you cannot do with var: it allows you to assign
    values to existing variables too. As long as at least one new variable is on the lefthand
    side of the :=, any of the other variables can already exist:
    ``` go
    x := 10
    x, y := 30, "hello"
    ```
    - CANNOT use := at package level, only use inside a function. (At package level use var)

6. <b>Avoid declaring variables outside of functions because they compli‐
cate data flow analysis.</b>
    - You should rarely declare variables outside of functions, in what’s called the package
block. Package-level variables whose values change are a
bad idea. When you have a variable outside of a function, it can be difficult to track
the changes made to it, which makes it hard to understand how data is flowing
through your program. This can lead to subtle bugs. As a general rule, you should
only declare variables in the package block that are effectively immutable (constant, not changing).
## General Discoveries

1. Placeholders
    - %d for integers
    - %s for strings
    - %v universal holder with default values set for each type
2. Exported name always begins with capital letter
    - math.Pi   
    - rand.Int
    - - any imported package will have function/methods in capital
3. Function can return any number of results via return
4. ``` go
    type User struct {
        Name string `json:"name"`
        Age int `json:"age"
        Email `
    }

## Interview Questions I encountered for Golang Developer

1. What is Mutex its uses?
2. How do you secure the websocket endpoints?
3. How do you secure connected clients?
4. Explain goroutines and channels
5. Difference between websocket and HTTP, advantages and disadv of websocket.
6. Can we send any data type into channels or is it strict for a the given data type while pushing? - NO Channels are typed
    - make(chan int) - can only receive integer types
