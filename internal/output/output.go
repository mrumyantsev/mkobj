package output

import (
	"os"

	"github.com/mrumyantsev/mkobj/internal/core"
)

const (
	colSpc = ": "
	eol    = "\n"
)

func Fatal(err error) {
	os.Stderr.WriteString(core.AppMark + colSpc + err.Error() + eol)

	os.Exit(core.FatalExitCode)
}
