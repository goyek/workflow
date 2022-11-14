// Package workflow is a customizable goyek flow.
package workflow

import (
	"os"

	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/boot"
)

func init() {
	goyek.SetDefault(PipelineAll)
}

// Library undefines all tasks that are related to an application build pipeline.
func Library() {
	goyek.Undefine(StageBuild)
	goyek.Undefine(TaskGoBuild)
}

// Main parses the args and runs the workflow.
func Main() {
	if err := os.Chdir(".."); err != nil {
		panic(err)
	}
	boot.Main()
}
