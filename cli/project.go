package cli

import (
	"os"
)

type Project struct {
	WorkingDir string
	PkgName    string
}

func newProject() (p *Project) {
	return &Project{}
}

func (p *Project) initialize() (err error) {
	wd, err := os.Getwd()
	if err != nil {
		return
	}

	p.WorkingDir = wd

	_ = p.acquirePkgName()

	return
}

func (p *Project) acquirePkgName() (err error) {
	name, err := getPakcageName(p.WorkingDir)
	if err != nil {
		return
	}

	// Acquire and sets the name only when it succeeds
	p.PkgName = name

	return
}
