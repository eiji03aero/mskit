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

// -------------------- commands --------------------
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
		return fmt.Errorf("directory exists: ", directoryPath)
	}

	err = os.MkdirAll(directoryPath, 0777)
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

	// -------------------- root --------------------
	err = createFileWithTemplate(
		directoryPath,
		"service.go",
		tpl.RootService(),
		p,
	)
	if err != nil {
		return
	}

	// -------------------- cmd --------------------
	appDirectoryPath := fmt.Sprintf("%s/cmd/app", directoryPath)
	err = os.MkdirAll(appDirectoryPath, 0777)
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

	// -------------------- domain --------------------
	domainDirectoryPath := fmt.Sprintf("%s/domain", directoryPath)
	err = os.MkdirAll(domainDirectoryPath, 0777)
	if err != nil {
		return
	}

	// -------------------- service --------------------
	serviceDirectoryPath := fmt.Sprintf("%s/service", directoryPath)

	err = os.MkdirAll(serviceDirectoryPath, 0777)
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

func (p *Project) generateAggregate(name string) (err error) {
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

	dir := fmt.Sprintf("%s/domain/%s", p.WorkingDir, data.LowerName)

	if _, err = os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0777); err != nil {
			return
		}
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

func (p *Project) generatePublisher() (err error) {
	dir := fmt.Sprintf("%s/transport/publisher", p.WorkingDir)

	if _, err = os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0777); err != nil {
			return
		}
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
	dir := fmt.Sprintf("%s/transport/consumer", p.WorkingDir)

	if _, err = os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0777); err != nil {
			return
		}
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
	dir := fmt.Sprintf("%s/transport/rpcendpoint", p.WorkingDir)

	if _, err = os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0777); err != nil {
			return
		}
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

	// -------------------- root --------------------
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

	// -------------------- transport/proxy --------------------
	proxyDir := fmt.Sprintf("%s/transport/proxy/%s", dir, data.LowerName)
	if _, err = os.Stat(proxyDir); os.IsNotExist(err) {
		if err = os.MkdirAll(proxyDir, 0777); err != nil {
			panic(err)
		}
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