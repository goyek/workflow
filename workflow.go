// Package workflow is a customizable goyek flow.
package workflow

import (
	"flag"
	"fmt"
	"strings"

	"github.com/goyek/goyek/v2"
	"github.com/goyek/goyek/v2/middleware"
)

var (
	dryRun  = flag.Bool("dry-run", false, "print all tasks that would be run without running them")
	longRun = flag.Duration("long-run", 0, "print when a task takes longer")
	noDeps  = flag.Bool("no-deps", false, "do not process dependencies")
	skip    = flag.String("skip", "", "skip processing the `comma-separated tasks`")
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
	flag.CommandLine.SetOutput(goyek.Output())
	flag.Usage = usage
	flag.Parse()

	if *dryRun {
		goyek.Use(middleware.DryRun)
	}
	goyek.Use(reportStatus)
	if *longRun > 0 {
		goyek.Use(middleware.ReportLongRun(*longRun))
	}

	var opts []goyek.Option
	if *noDeps {
		opts = append(opts, goyek.NoDeps())
	}
	if *skip != "" {
		skippedTasks := strings.Split(*skip, ",")
		opts = append(opts, goyek.Skip(skippedTasks...))
	}

	goyek.SetUsage(usage)
	goyek.SetLogger(goyek.FmtLogger{})
	goyek.Main(flag.Args(), opts...)
}

func usage() {
	fmt.Println("Usage of build: [flags] [--] [tasks]")
	goyek.Print()
	fmt.Println("Flags:")
	flag.PrintDefaults()
}
