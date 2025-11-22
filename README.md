# Go-Portfolio
Learning Go and building project along the way, This README will be my note taking place


## NOTES

1. Go compiles to a single native binary i.e. it does not require aother applications to run and is easy to distribute 
2. go vet - scans for coding mistakes
3. <b>Steps to create a Go project </b>
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