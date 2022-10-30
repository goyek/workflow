package workflow

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/goyek/goyek/v2"
)

func init() {
	color.NoColor = os.Getenv("NO_COLOR") != ""
}

//nolint:errcheck,gosec // not checking errors when writing to output
func reportStatus(next goyek.Runner) goyek.Runner {
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

type codeLineLogger struct {
	mu          sync.Mutex
	helperPCs   map[uintptr]struct{} // functions to be skipped when writing file/line info
	helperNames map[string]struct{}  // helperPCs converted to function names
}

// Log is used internally in order to provide proper prefix.
func (l *codeLineLogger) Log(w io.Writer, args ...interface{}) {
	txt := fmt.Sprint(args...)
	txt = l.decorate(txt)
	io.WriteString(w, txt) //nolint:errcheck,gosec // not checking errors when writing to output
}

// Logf is used internally in order to provide proper prefix.
func (l *codeLineLogger) Logf(w io.Writer, format string, args ...interface{}) {
	txt := fmt.Sprintf(format, args...)
	txt = l.decorate(txt)
	io.WriteString(w, txt) //nolint:errcheck,gosec // not checking errors when writing to output
}

func (l *codeLineLogger) Error(w io.Writer, args ...interface{}) {
	txt := fmt.Sprint(args...)
	txt = l.decorate(txt)
	color.New(color.FgRed).Fprint(w, txt) //nolint:errcheck,gosec // not checking errors when writing to output
}

func (l *codeLineLogger) Errorf(w io.Writer, format string, args ...interface{}) {
	txt := fmt.Sprintf(format, args...)
	txt = l.decorate(txt)
	color.New(color.FgRed).Fprint(w, txt) //nolint:errcheck,gosec // not checking errors when writing to output
}

func (l *codeLineLogger) Fatal(w io.Writer, args ...interface{}) {
	txt := fmt.Sprint(args...)
	txt = l.decorate(txt)
	color.New(color.FgRed).Fprint(w, txt) //nolint:errcheck,gosec // not checking errors when writing to output
}

func (l *codeLineLogger) Fatalf(w io.Writer, format string, args ...interface{}) {
	txt := fmt.Sprintf(format, args...)
	txt = l.decorate(txt)
	color.New(color.FgRed).Fprint(w, txt) //nolint:errcheck,gosec // not checking errors when writing to output
}

func (l *codeLineLogger) Skip(w io.Writer, args ...interface{}) {
	txt := fmt.Sprint(args...)
	txt = l.decorate(txt)
	color.New(color.FgYellow).Fprint(w, txt) //nolint:errcheck,gosec // not checking errors when writing to output
}

func (l *codeLineLogger) Skipf(w io.Writer, format string, args ...interface{}) {
	txt := fmt.Sprintf(format, args...)
	txt = l.decorate(txt)
	color.New(color.FgYellow).Fprint(w, txt) //nolint:errcheck,gosec // not checking errors when writing to output
}

// Helper marks the calling function as a helper function.
// When printing file and line information, that function will be skipped.
// Helper may be called simultaneously from multiple goroutines.
func (l *codeLineLogger) Helper() {
	var pc [1]uintptr
	const skip = 3 // skip: runtime.Callers + codeLineLogger.Helper + TF.Helper
	n := runtime.Callers(skip, pc[:])
	if n == 0 {
		panic("zero callers found")
	}

	l.mu.Lock()
	defer l.mu.Unlock()
	if l.helperPCs == nil {
		l.helperPCs = make(map[uintptr]struct{})
	}
	if _, found := l.helperPCs[pc[0]]; !found {
		l.helperPCs[pc[0]] = struct{}{}
		l.helperNames = nil // map will be recreated next time it is needed
	}
}

// decorate prefixes the string with the file and line of the call site
// and inserts the final newline and indentation spaces for formatting.
func (l *codeLineLogger) decorate(s string) string {
	const skip = 3
	frame := l.frameSkip(skip)
	file := frame.File
	line := frame.Line
	if file != "" {
		// Truncate file name at last file name separator.
		if index := strings.LastIndex(file, "/"); index >= 0 {
			file = file[index+1:]
		} else if index = strings.LastIndex(file, "\\"); index >= 0 {
			file = file[index+1:]
		}
	} else {
		file = "???"
	}
	if line == 0 {
		line = 1
	}
	buf := &strings.Builder{}
	// Every line is indented at least 6 spaces.
	buf.WriteString("      ")
	fmt.Fprintf(buf, "%s:%d: ", file, line)
	lines := strings.Split(s, "\n")
	if l := len(lines); l > 1 && lines[l-1] == "" {
		lines = lines[:l-1]
	}
	for i, line := range lines {
		if i > 0 {
			// Second and subsequent lines are indented an additional 4 spaces.
			buf.WriteString("\n          ")
		}
		buf.WriteString(line)
	}
	buf.WriteByte('\n')
	return buf.String()
}

// frameSkip searches, starting after skip frames, for the first caller frame
// in a function not marked as a helper and returns that frame.
// The search stops if it finds a tRunner function that
// was the entry point into the test and the test is not a subtest.
// This function must be called with l.mu held.
func (l *codeLineLogger) frameSkip(skip int) runtime.Frame {
	// The maximum number of stack frames to go through when skipping helper functions for
	// the purpose of decorating log messages.
	const maxStackLen = 50
	var pc [maxStackLen]uintptr

	const skipMore = 2 // skip: runtime.Callers + codeLineLogger.frameSkip
	n := runtime.Callers(skip+skipMore, pc[:])
	if n == 0 {
		panic("zero callers found")
	}

	frames := runtime.CallersFrames(pc[:n])
	l.mu.Lock()
	defer l.mu.Unlock()
	var firstFrame, prevFrame, frame runtime.Frame
	for more := true; more; prevFrame = frame {
		frame, more = frames.Next()
		if frame.Function == "runtime.gopanic" {
			continue
		}
		if firstFrame.PC == 0 {
			firstFrame = frame
		}
		if frame.Function == "github.com/goyek/goyek/v2.taskRunner.run.func1" {
			// We've gone up all the way to the runner calling
			// the action (so the user must have
			// called tf.Helper from inside that action).
			return prevFrame
		}
		// If more helper PCs have been added since we last did the conversion
		if l.helperNames == nil {
			l.helperNames = make(map[string]struct{})
			for pc := range l.helperPCs {
				l.helperNames[pcToName(pc)] = struct{}{}
			}
		}
		if _, ok := l.helperNames[frame.Function]; !ok {
			// Found a frame that wasn't inside a helper function.
			return frame
		}
	}
	return firstFrame
}

func pcToName(pc uintptr) string {
	pcs := []uintptr{pc}
	frames := runtime.CallersFrames(pcs)
	frame, _ := frames.Next()
	return frame.Function
}
