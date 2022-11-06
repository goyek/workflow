package workflow

import (
	"github.com/goyek/goyek/v2"
)

var TaskClean = goyek.Define(goyek.Task{
	Name:  "clean",
	Usage: "remove files created during build pipeline",
	Action: func(a *goyek.A) {
		Remove(a, "dist")
		Remove(a, "coverage.out")
		Remove(a, "coverage.html")
	},
})
