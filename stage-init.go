package workflow

import "github.com/goyek/goyek/v2"

var StageInit = goyek.Define(goyek.Task{
	Name:  "init",
	Usage: "init stage",
	Deps: goyek.Deps{
		TaskGoMod,
	},
})
