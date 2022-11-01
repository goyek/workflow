package workflow

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var TaskGoMod = goyek.Define(goyek.Task{
	Name:  "go-mod",
	Usage: "go mod tidy",
	Action: func(tf *goyek.TF) {
		cmd.Exec(tf, "go mod tidy")
	},
})
