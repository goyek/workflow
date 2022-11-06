package workflow

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var TaskGoMod = goyek.Define(goyek.Task{
	Name:  "go-mod",
	Usage: "go mod tidy",
	Action: func(a *goyek.A) {
		cmd.Exec(a, "go mod tidy")
	},
})
