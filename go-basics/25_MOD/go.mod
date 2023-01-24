module mod

// module is used to define a module and its dependencies

go 1.19

// version of the golang which we are working on

require github.com/gorilla/mux v1.8.0

// go mod verify is used to verify the dependencies
// go mod tidy is used to tidy the dependencies
// go mod init is used to initialize the module
// go list -m all is used to list all the dependencies

// go list -m --versions github.com/gorilla/mux is used to list
//  all the versions of the module
