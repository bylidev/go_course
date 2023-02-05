### Description

This proyect gives you an idea about go modules.

#### Steps
- create greetings folder and hello folder
- In the grettings folder create a module with the next command
` go mod init byli.dev/greetings`
- We added a greetings.go file with has the function to be imported in another module. in this case in the hello module.
- In the hello folder we have to create the hello module with the next cli command
`go mod init byli.dev/hello`
- We created a hello.go file wich imports 'byli.dev/greetings' module.
- As byli.dev is not a go registry we have to replace the url reference to the local directory with this command inside the hello folder
`go mod edit -replace byli.dev/greetings=../greetings`
- refreshing dependences in hello folder
`go mod tidy`
- init the hello application
` go run .` 

## Notes

To reference a published module, a go.mod file would typically omit the replace directive and use a require directive with a tagged version number at the end.

`require example.com/greetings v1.1.0`