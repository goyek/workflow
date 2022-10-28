package workflow

import (
	"os/exec"

	"github.com/goyek/goyek/v2"
)

var goBuild = goyek.Define(goyek.Task{
	Name:  "go-build",
	Usage: "go build",
	Action: func(tf *goyek.TF) {
		ext, err := exec.CommandContext(tf.Context(), "go", "env", "GOEXE").CombinedOutput()
		if err != nil {
			tf.Fatal(err)
		}
		Exec(tf, `go build -ldflags="-s -w" -o dist/app`+string(ext))
	},
})
