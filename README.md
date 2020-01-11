### Learning Go

1.   Write test first
2.   Write code second
3.   A simple calculator code
    -   [ ] Validation a return value in a range of int type
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

    import "testing"

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
    ```
-   sample of calc.go
    ```go
    package calc

    func Add(num1, num2 int) int {
        return num1 + num2
    }

    func Substract(num1, num2 int) int {
        return num1 - num2
    }
    func Multiply(num1, num2 int) int {
        return num1 * num2
    }
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
