package workflow

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var TaskGoTest = goyek.Define(goyek.Task{
	Name:  "go-test",
	Usage: "go test",
	Action: func(tf *goyek.TF) {
		if !cmd.Exec(tf, "go test -race -covermode=atomic -coverprofile=coverage.out -coverpkg=./... ./...") {
			return
		}
		cmd.Exec(tf, "go tool cover -html=coverage.out -o coverage.html")
	},
})
