### Testing Driven Development w GIN Framework
-   https://pkg.go.dev/github.com/gin-gonic/gin?tab=doc
-   https://godoc.org/github.com/gin-gonic/gin
-   https://github.com/quii/learn-go-with-tests/blob/master/README.md

### Folder Structure

    ├── alochym.db
    ├── api
    │   ├── controllers
    │   │   └── ping.go
    │   ├── models
    │   │   └── ping.go
    │   ├── routes.go
    │   └── templates
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── README.md
    └── test
        └── ping_test.go

-   api folder is main folder
    -   controllers
        -   ping.go
            -   To define all functions that are related to `/ping` url
    -   models
        -   ping.go
            -   To define all interaction with database
    -   templates
        -   to store all application html templates
    -   routes.go
        -   To define all routes of application
-   test folder
    -   To define all application's testing function
    -   ping_test.go is an example for testing all ping's controller function
-   main.go is an entry point to run application
### How to use
-   `cd test && go test -v`
