package workflow

import "github.com/goyek/goyek/v2"

var TaskGoGenerate = goyek.Define(goyek.Task{
	Name:  "go-generate",
	Usage: "go generate",
	Action: func(tf *goyek.TF) {
		Exec(tf, "go generate ./...")
	},
})
