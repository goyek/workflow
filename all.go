package workflow

import "github.com/goyek/goyek/v2"

var PipelineAll = goyek.Define(goyek.Task{
	Name:  "all",
	Usage: "build pipeline",
	Deps: goyek.Deps{
		StageInit,
		StageBuild,
		StageTest,
	},
})

func init() {
	goyek.SetDefault(PipelineAll)
}
