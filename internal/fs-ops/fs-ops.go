package fsops

import (
	"os"

	"github.com/mrumyantsev/mkobj/pkg/errno"
)

const (
	permDirCreate  = 0755
	permFileCreate = 0644
)

type FsOps struct {
}

func New() *FsOps {
	return &FsOps{}
}

func (o *FsOps) MakeDir(path string) error {
	if err := os.MkdirAll(path, permDirCreate); err != nil {
		return errno.Attach(err, "3.1.1")
	}

	return nil
}

func (o *FsOps) MakeWriteFile(path string, content []byte) error {
	if err := os.WriteFile(path, content, permFileCreate); err != nil {
		return errno.Attach(err, "3.2.1")
	}

	return nil
}
