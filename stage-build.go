package workflow

import "github.com/goyek/goyek/v2"

var StageBuild = goyek.Define(goyek.Task{
	Name:  "build",
	Usage: "build stage",
	Deps: goyek.Deps{
		TaskGoBuild,
	},
})
