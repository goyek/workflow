package workflow

import "github.com/goyek/goyek/v2"

var TaskGoTest = goyek.Define(goyek.Task{
	Name:  "go-test",
	Usage: "go test",
	Action: func(tf *goyek.TF) {
		if !Exec(tf, "go test -race -covermode=atomic -coverprofile=coverage.out -coverpkg=./... ./...") {
			return
		}
		Exec(tf, "go tool cover -html=coverage.out -o coverage.html")
	},
})
