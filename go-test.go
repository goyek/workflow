package workflow

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var TaskGoTest = goyek.Define(goyek.Task{
	Name:  "go-test",
	Usage: "go test",
	Action: func(a *goyek.A) {
		if !cmd.Exec(a, "go test -race -covermode=atomic -coverprofile=coverage.out -coverpkg=./... ./...") {
			return
		}
		cmd.Exec(a, "go tool cover -html=coverage.out -o coverage.html")
	},
})
