package workflow

import "github.com/goyek/goyek/v2"

var TaskGoMod = goyek.Define(goyek.Task{
	Name:  "go-mod",
	Usage: "go mod tidy",
	Action: func(tf *goyek.TF) {
		Exec(tf, "go mod tidy")
	},
})
