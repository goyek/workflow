package workflow

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var TaskGoVet = goyek.Define(goyek.Task{
	Name:  "go-vet",
	Usage: "go vet",
	Action: func(a *goyek.A) {
		cmd.Exec(a, "go vet ./...")
	},
})
