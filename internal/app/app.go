package app

import (
	"fmt"

	cliparser "github.com/mrumyantsev/mkobj/internal/cli-parser"
	"github.com/mrumyantsev/mkobj/internal/config"
	fsops "github.com/mrumyantsev/mkobj/internal/fs-ops"
	"github.com/mrumyantsev/mkobj/internal/models"
	"github.com/mrumyantsev/mkobj/internal/templates"
)

type CliParser interface {
	Parse() (*models.ObjectInfo, error)
}

type FsOps interface {
	MakeDir(path string) error
	MakeWriteFile(path string, content []byte) error
}

type App struct {
	config *config.Config
	parser CliParser
	fsops  FsOps
}

func New() (*App, error) {
	cfg := config.New()

	return &App{
		config: cfg,
		parser: cliparser.New(cfg),
		fsops:  fsops.New(),
	}, nil
}

func (a *App) Run() error {
	info, err := a.parser.Parse()
	if err != nil {
		return err
	}

	var objDirPath string

	if info.ContainDir != "" {
		objDirPath = info.ContainDir + a.config.PathSeparator + info.ObjectDir
	} else {
		objDirPath = info.ObjectDir
	}

	if err = a.fsops.MakeDir(objDirPath); err != nil {
		return err
	}

	tmpl := fmt.Sprintf(
		templates.TemplateObject,
		info.PackageName,
		info.ObjectName,
		info.ObjectName,
		info.ObjectName,
	)

	objFilePath := objDirPath + a.config.PathSeparator + info.ObjectDir + ".go"

	if err = a.fsops.MakeWriteFile(objFilePath, []byte(tmpl)); err != nil {
		return err
	}

	return nil
}
