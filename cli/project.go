package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/eiji03aero/mskit/cli/tpl"
	"github.com/iancoleman/strcase"
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

// -------------------- initialize --------------------
func (p *Project) initializeService(
	directoryPath string,
	pkgName string,
) (err error) {
	p.PkgName = pkgName

	if !filepath.IsAbs(directoryPath) {
		directoryPath = filepath.Join(p.WorkingDir, directoryPath)
		directoryPath, err = filepath.Abs(directoryPath)
		if err != nil {
			return
		}
	}

	if _, err = os.Stat(directoryPath); os.IsExist(err) {
		return fmt.Errorf("directory exists: %s", directoryPath)
	}

	err = createDir(directoryPath)
	if err != nil {
		return
	}

	err = os.Chdir(directoryPath)
	if err != nil {
		return
	}

	err = exec.Command("go", "mod", "init", p.PkgName).Run()
	if err != nil {
		return
	}

	err = createFileWithTemplate(
		directoryPath,
		"service.go",
		tpl.RootService(),
		p,
	)
	if err != nil {
		return
	}

	appDirectoryPath := fmt.Sprintf("%s/cmd/app", directoryPath)
	err = createDir(appDirectoryPath)
	if err != nil {
		return
	}

	err = createFileWithTemplate(
		appDirectoryPath,
		"main.go",
		tpl.CmdAppTemplate(),
		p,
	)
	if err != nil {
		return
	}

	err = createFileWithTemplate(
		appDirectoryPath,
		"env.go",
		tpl.CmdEnvTemplate(),
		p,
	)
	if err != nil {
		return
	}

	domainDirectoryPath := fmt.Sprintf("%s/domain", directoryPath)
	err = createDir(domainDirectoryPath)
	if err != nil {
		return
	}

	serviceDirectoryPath := fmt.Sprintf("%s/service", directoryPath)

	err = createDir(serviceDirectoryPath)
	if err != nil {
		return
	}

	err = createFileWithTemplate(
		serviceDirectoryPath,
		"service.go",
		tpl.ServiceTemplate(),
		p,
	)
	if err != nil {
		return
	}

	return
}

// -------------------- domain --------------------
func (p *Project) generateDomainAggregate(name string) (err error) {
	data := struct {
		Name          string
		LowerName     string
		AggregateName string
		SnakeName     string
		NameInitial   string
	}{}
	data.Name = name
	data.LowerName = strings.ToLower(name)
	data.AggregateName = name + "Aggregate"
	data.SnakeName = strcase.ToSnake(name)
	data.NameInitial = string([]rune(data.LowerName)[0])

	dir := fmt.Sprintf("%s/domain/entity/%s", p.WorkingDir, data.LowerName)

	err = createDir(dir)
	if err != nil {
		return
	}

	err = createFileWithTemplate(
		dir,
		data.LowerName+".go",
		tpl.DomainEntityTemplate(),
		data,
	)
	if err != nil {
		return
	}

	err = createFileWithTemplate(
		dir,
		data.LowerName+"_aggregate.go",
		tpl.DomainAggregateTemplate(),
		data,
	)
	if err != nil {
		return
	}

	err = createFileWithTemplate(
		dir,
		data.LowerName+"_commands.go",
		tpl.DomainCommandsTemplate(),
		data,
	)
	if err != nil {
		return
	}

	err = createFileWithTemplate(
		dir,
		data.LowerName+"_events.go",
		tpl.DomainEventsTemplate(),
		data,
	)
	if err != nil {
		return
	}

	return
}

func (p *Project) generateDomainService(name string) (err error) {
	data := struct {
		PkgName       string
		Name          string
		LowerName     string
		SnakeName     string
		InterfaceName string
	}{
		PkgName:       p.PkgName,
		Name:          name,
		LowerName:     strings.ToLower(name),
		SnakeName:     strcase.ToSnake(name),
		InterfaceName: name + "Service",
	}

	rootFileName := "domain.go"
	rootFilePath := fmt.Sprintf("%s/%s", p.WorkingDir, rootFileName)

	if _, err = os.Stat(rootFilePath); os.IsNotExist(err) {
		err = createFileWithTemplate(
			p.WorkingDir,
			rootFileName,
			tpl.RootProxy(),
			data,
		)
		if err != nil {
			panic(err)
		}
	}

	err = appendToFileWithTemplate(
		rootFilePath,
		tpl.Interface(),
		data,
	)
	if err != nil {
		panic(err)
	}

	dmnDir := fmt.Sprintf("%s/domain/service/%s", p.WorkingDir, data.LowerName)
	err = createDir(dmnDir)
	if err != nil {
		return
	}

	err = createFileWithTemplate(
		dmnDir,
		data.LowerName+".go",
		tpl.DomainServiceTemplate(),
		data,
	)
	if err != nil {
		return
	}

	return
}

// -------------------- adapters --------------------
func (p *Project) generatePublisher() (err error) {
	dir := fmt.Sprintf("%s/adapter/publisher", p.WorkingDir)

	err = createDir(dir)
	if err != nil {
		return
	}

	err = createFileWithTemplate(
		dir,
		"publisher.go",
		tpl.PublisherTemplate(),
		p,
	)
	if err != nil {
		return
	}

	return
}

func (p *Project) generateConsumer() (err error) {
	dir := fmt.Sprintf("%s/adapter/consumer", p.WorkingDir)

	err = createDir(dir)
	if err != nil {
		return
	}

	err = createFileWithTemplate(
		dir,
		"consumer.go",
		tpl.ConsumerTemplate(),
		p,
	)
	if err != nil {
		return
	}

	return
}

func (p *Project) generateRPCEndpoint() (err error) {
	dir := fmt.Sprintf("%s/adapter/rpcendpoint", p.WorkingDir)

	err = createDir(dir)
	if err != nil {
		return
	}

	err = createFileWithTemplate(
		dir,
		"rpcendpoint.go",
		tpl.RPCEndpointTemplate(),
		p,
	)
	if err != nil {
		return
	}

	return
}

func (p *Project) generateProxy(name string) (err error) {
	dir := p.WorkingDir
	data := struct {
		PkgName       string
		Name          string
		LowerName     string
		InterfaceName string
	}{
		PkgName:       p.PkgName,
		Name:          name,
		LowerName:     strings.ToLower(name),
		InterfaceName: fmt.Sprintf("%sProxy", name),
	}

	fileName := "proxy.go"
	rootFilePath := fmt.Sprintf("%s/%s", dir, fileName)
	if _, err = os.Stat(rootFilePath); os.IsNotExist(err) {
		err = createFileWithTemplate(
			dir,
			fileName,
			tpl.RootProxy(),
			data,
		)
		if err != nil {
			panic(err)
		}
	}

	err = appendToFileWithTemplate(
		rootFilePath,
		tpl.Interface(),
		data,
	)
	if err != nil {
		panic(err)
	}

	proxyDir := fmt.Sprintf("%s/adapter/proxy/%s", dir, data.LowerName)

	err = createDir(proxyDir)
	if err != nil {
		return
	}

	err = createFileWithTemplate(
		proxyDir,
		"proxy.go",
		tpl.ProxyTemplate(),
		data,
	)
	if err != nil {
		panic(err)
	}

	return
}

// -------------------- saga --------------------
func (p *Project) generateSaga(name string) (err error) {
	dir := p.WorkingDir
	data := struct {
		Name      string
		PkgName   string
		LowerName string
	}{
		Name:      name,
		PkgName:   p.PkgName,
		LowerName: strings.ToLower(name),
	}

	sagaDir := fmt.Sprintf("%s/saga", dir)
	sagaImplDir := fmt.Sprintf("%s/%s", sagaDir, data.LowerName)

	err = createDir(sagaImplDir)
	if err != nil {
		return
	}

	err = createFileWithTemplate(
		sagaImplDir,
		"manager.go",
		tpl.SagaManagerTemplate(),
		data,
	)
	if err != nil {
		return
	}

	err = createFileWithTemplate(
		sagaImplDir,
		"state.go",
		tpl.SagaStateTemplate(),
		data,
	)
	if err != nil {
		return
	}

	return
}
