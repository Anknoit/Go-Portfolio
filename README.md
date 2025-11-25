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
3. Literals 
    - Integer literal
    - Float literal
    - Rune Literal - 'single quote for rune'
    - String literal - "double quote for strings"
