package workflow

import (
	"os"

	"github.com/goyek/goyek/v2"
)

// Remove removes path and any children it contains.
// It calls a.Error in case there is an error during removal.
func Remove(a *goyek.A, path string) {
	a.Helper()
	if _, err := os.Stat(path); err != nil {
		return
	}
	a.Log("Remove: " + path)
	if err := os.RemoveAll(path); err != nil {
		a.Error(err)
	}
}
