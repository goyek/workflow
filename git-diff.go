package workflow

import (
	"io"
	"strings"

	"github.com/goyek/goyek/v2"
)

var TaskGitDiff = goyek.Define(goyek.Task{
	Name:  "git-diff",
	Usage: "git diff",
	Action: func(tf *goyek.TF) {
		Exec(tf, "git diff --exit-code")

		tf.Log("Exec: git status --porcelain")
		cmd := tf.Cmd("git", "status", "--porcelain")
		sb := &strings.Builder{}
		cmd.Stdout = io.MultiWriter(tf.Output(), sb)
		if err := cmd.Run(); err != nil {
			tf.Error(err)
		}
		if sb.Len() > 0 {
			tf.Error("git status --porcelain returned output")
		}
	},
})
