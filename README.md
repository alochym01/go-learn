### VS Code setting

-   Setting GOPATH for terminal `/opt/hadn/go`
-   Using `go mod init github.com/alochym01/go-learn` in current directory
-   go-tool for vs code
    -   user setting
        ```json
        {
            "editor.fontFamily": "'Fira Code'",
            "editor.fontSize": 18,
            "editor.fontLigatures": true,
            "editor.fontWeight": "bold",
            "terminal.integrated.fontFamily": "'Fira Code'",
            "files.trimTrailingWhitespace": true,
            "go.toolsGopath": "/opt/hadn/go",
        }
        ```

### Learning Go

-   Write test first
-   Write code second
-   A simple calculator code
    -   [ ] Not validated a return value in a range of int type
-   View go document `godoc -http:=8080`
    -   Navigate to `http://localhost:8080/pkg/calc`
### Example

-   folder structure
    ```bash
    ├── bin
    ├── pkg
    ├── README.md
    └── src
        ├── calc
        │   ├── calc.go
        │   └── calc_test.go
        └── main.go

    4 directories, 4 files
    ```
-   sample of calc_test.go
    ```go
    package calc

    import (
        "fmt"
        "testing"
    )

    func TestCalc(t *testing.T) {
        t.Run("Running test add function", func(t *testing.T) {
            got := Add(2, 3)
            want := 5
            if got != want {
                t.Errorf("got %q want %q", got, want)
            }
        })
        t.Run("Running test substract function", func(t *testing.T) {
            got := Substract(3, 2)
            want := 1
            if got != want {
                t.Errorf("got %q want %q", got, want)
            }
        })
        t.Run("Running test multiply function", func(t *testing.T) {
            got := Multiply(2, 3)
            want := 6
            if got != want {
                t.Errorf("got %q want %q", got, want)
            }
        })
        t.Run("Running test divide function", func(t *testing.T) {
            got := Divide(4, 2)
            want := 2
            if got != want {
                t.Errorf("got %q want %q", got, want)
            }
        })
    }

    func ExampleAdd() {
        sum := Add(1, 5)
        fmt.Println(sum)
        // Output: 6
    }

    func ExampleSubstract() {
        sum := Substract(1, 5)
        fmt.Println(sum)
        // Output: -4
    }

    func ExampleMultiply() {
        sum := Multiply(1, 5)
        fmt.Println(sum)
        // Output: 5
    }

    func ExampleDivide() {
        sum := Divide(1, 5)
        fmt.Println(sum)
        // Output: 0
    }
    ```
-   sample of calc.go
    ```go
    // Package calc is implemented Add/Substract/Nultiply/Divide function.
    package calc

    // Add takes two integers and returns the sum of them.
    func Add(num1, num2 int) int {
        return num1 + num2
    }

    // Substract takes two integers and returns the substract of them.
    func Substract(num1, num2 int) int {
        return num1 - num2
    }

    // Multiply takes two integers and returns the multiply of them.
    func Multiply(num1, num2 int) int {
        return num1 * num2
    }

    // Divide takes two integers and returns the divide of them.
    func Divide(num1, num2 int) int {
        return num1 / num2
    }
    ```
-   sample test run
    ```bash
    [hadn@hadn src]$ go test -v calc
    === RUN   TestCalc
    === RUN   TestCalc/Running_test_add_function
    === RUN   TestCalc/Running_test_substract_function
    === RUN   TestCalc/Running_test_multiply_function
    === RUN   TestCalc/Running_test_divide_function
    --- PASS: TestCalc (0.00s)
        --- PASS: TestCalc/Running_test_add_function (0.00s)
        --- PASS: TestCalc/Running_test_substract_function (0.00s)
        --- PASS: TestCalc/Running_test_multiply_function (0.00s)
        --- PASS: TestCalc/Running_test_divide_function (0.00s)
    PASS
    ok      calc    (cached)
    ```
-   create `main.go` file in `src` folder, content `main.go`
    ```go
    package main

    import (
        "fmt"

        "github.com/alochym01/go-learn/src/calc"
    )

    func main() {
        fmt.Println(calc.Add(2, 3))
        fmt.Println(calc.Substract(2, 3))
        fmt.Println(calc.Multiply(2, 3))
        fmt.Println(calc.Divide(2, 3))
    }
    ```
-   run `go run main.go`
