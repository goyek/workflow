package workflow

import (
	"github.com/goyek/goyek/v2"
)

var TaskClean = goyek.Define(goyek.Task{
	Name:  "clean",
	Usage: "remove files created during build pipeline",
	Action: func(tf *goyek.TF) {
		Remove(tf, "dist")
		Remove(tf, "coverage.out")
		Remove(tf, "coverage.html")
	},
})
