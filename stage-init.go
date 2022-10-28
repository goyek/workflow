package workflow

import "github.com/goyek/goyek/v2"

var stageInit = goyek.Define(goyek.Task{
	Name:  "init",
	Usage: "init stage",
	Deps: goyek.Deps{
		goMod,
	},
})
