package workflow

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var TaskGoGenerate = goyek.Define(goyek.Task{
	Name:  "go-generate",
	Usage: "go generate",
	Action: func(a *goyek.A) {
		cmd.Exec(a, "go generate ./...")
	},
})
