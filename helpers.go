package workflow

import (
	"os"

	"github.com/goyek/goyek/v2"
	"github.com/mattn/go-shellwords"
)

// Exec runs the command in given directory.
// It calls tf.Error[f] and returns false in case of any problems.
func Exec(tf *goyek.TF, cmdLine string) bool {
	tf.Helper()
	tf.Logf("Exec: %s", cmdLine)
	args, err := shellwords.Parse(cmdLine)
	if err != nil {
		tf.Errorf("parse command line: %v", err)
		return false
	}
	cmd := tf.Cmd(args[0], args[1:]...)
	if err := cmd.Run(); err != nil {
		tf.Error(err)
		return false
	}
	return true
}

// Remove removes path and any children it contains.
// It calls tf.Error in case there is an error during removal.
func Remove(tf *goyek.TF, path string) {
	tf.Helper()
	if _, err := os.Stat(path); err != nil {
		return
	}
	tf.Log("Remove: " + path)
	if err := os.RemoveAll(path); err != nil {
		tf.Error(err)
	}
}
