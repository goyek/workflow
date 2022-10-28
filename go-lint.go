package workflow

import "github.com/goyek/goyek/v2"

var TaskGoVet = goyek.Define(goyek.Task{
	Name:  "go-vet",
	Usage: "go vet",
	Action: func(tf *goyek.TF) {
		Exec(tf, "go vet ./...")
	},
})
