package workflow

import "github.com/goyek/goyek/v2"

var stageBuild = goyek.Define(goyek.Task{
	Name:  "build",
	Usage: "build stage",
	Deps: goyek.Deps{
		goBuild,
	},
})
