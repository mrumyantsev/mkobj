package errno

import "fmt"

const (
	tmplAttach = "[Errno %s] %w"
)

func Attach(err error, errno string) error {
	return fmt.Errorf(tmplAttach, errno, err)
}
