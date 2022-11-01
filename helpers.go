package workflow

import (
	"os"

	"github.com/goyek/goyek/v2"
)

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
