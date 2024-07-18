package cliparser

import (
	"errors"
	"flag"
	"regexp"
	"strings"

	"github.com/mrumyantsev/mkobj/internal/config"
	"github.com/mrumyantsev/mkobj/internal/models"
	"github.com/mrumyantsev/mkobj/pkg/errno"
)

var (
	ErrNoObjectPathSpecified = errors.New("no object path specified")
)

type CliParser struct {
	config *config.Config
}

func New(cfg *config.Config) *CliParser {
	return &CliParser{config: cfg}
}

func (p *CliParser) Parse() (*models.ObjectInfo, error) {
	flag.Parse()

	objPath := flag.Arg(0)

	if len(objPath) == 0 {
		return nil, errno.Attach(ErrNoObjectPathSpecified, "2.1.1")
	}

	idx := strings.LastIndex(objPath, p.config.PathSeparator)

	var containDir string

	if idx != -1 {
		containDir = strings.ToLower(objPath[:idx])
	}

	dirtyObjDir := objPath[idx+1:]

	re := regexp.MustCompile("[^0-9A-Za-z._-]")

	leapyObjDir := re.ReplaceAllString(dirtyObjDir, "")

	objDir := strings.ToLower(leapyObjDir)

	re = regexp.MustCompile("[._-]")

	pkgName := re.ReplaceAllString(objDir, "")

	objName := re.ReplaceAllString(leapyObjDir, "")

	return &models.ObjectInfo{
		ContainDir:  containDir,
		ObjectDir:   objDir,
		PackageName: pkgName,
		ObjectName:  objName,
	}, nil
}
