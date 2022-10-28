package workflow

import (
	"io"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/goyek/goyek/v2"
)

//nolint:errcheck,gosec // not checking errors when writing to output
func reportStatus(next goyek.Runner) goyek.Runner {
	color.NoColor = os.Getenv("NO_COLOR") != ""

	return func(in goyek.Input) goyek.Result {
		c := color.New(color.FgBlue)

		// report start task
		c.Fprintf(in.Output, "===== TASK  %s\n", in.TaskName)
		start := time.Now()

		// run
		res := next(in)

		// report task end
		c = color.New(color.FgGreen)
		status := "PASS"
		switch res.Status {
		case goyek.StatusFailed:
			c = color.New(color.FgRed, color.Bold)
			status = "FAIL"
		case goyek.StatusSkipped:
			c = color.New(color.FgYellow)
			status = "SKIP"
		case goyek.StatusNotRun:
			status = "NOOP"
		}
		c.Fprintf(in.Output, "----- %s: %s (%.2fs)\n", status, in.TaskName, time.Since(start).Seconds())

		// report panic if happened
		if res.PanicStack != nil {
			if res.PanicValue != nil {
				c.Fprint(in.Output, "panic:", res.PanicValue)
			} else {
				c.Fprint(in.Output, "panic(nil) or runtime.Goexit() called")
			}
			io.WriteString(in.Output, "\n\n")
			c.Fprintf(in.Output, "%s", res.PanicStack)
		}

		return res
	}
}
