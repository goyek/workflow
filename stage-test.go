package workflow

import "github.com/goyek/goyek/v2"

var StageTest = goyek.Define(goyek.Task{
	Name:  "test",
	Usage: "test stage",
	Deps: goyek.Deps{
		TaskGoLint,
		TaskGoTest,
	},
})
