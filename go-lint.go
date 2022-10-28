package workflow

import "github.com/goyek/goyek/v2"

var goLint = goyek.Define(goyek.Task{
	Name:  "go-lint",
	Usage: "go vet",
	Action: func(tf *goyek.TF) {
		Exec(tf, "go vet ./...")
	},
})
