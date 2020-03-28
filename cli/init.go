package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/eiji03aero/mskit/cli/tpl"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize skelton for mskit",
	Long:  "Initialize skelton for mskit",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		directoryPath := args[0]
		project.PkgName = filepath.Base(directoryPath)

		if !filepath.IsAbs(directoryPath) {
			directoryPath = filepath.Join(project.WorkingDir, directoryPath)
			directoryPath, err = filepath.Abs(directoryPath)
			if err != nil {
				panic(err)
			}
		}

		if _, err = os.Stat(directoryPath); os.IsExist(err) {
			panic(fmt.Errorf("directory exists: ", directoryPath))
		}

		err = os.MkdirAll(directoryPath, 0777)
		if err != nil {
			panic(err)
		}

		err = os.Chdir(directoryPath)
		if err != nil {
			panic(err)
		}

		err = exec.Command("go", "mod", "init", project.PkgName).Run()
		if err != nil {
			panic(err)
		}

		// -------------------- root --------------------
		err = createFileWithTemplate(
			directoryPath,
			"service.go",
			tpl.RootService(),
			project,
		)
		if err != nil {
			panic(err)
		}

		// -------------------- domain --------------------
		domainDirectoryPath := fmt.Sprintf("%s/domain", directoryPath)
		err = os.MkdirAll(domainDirectoryPath, 0777)
		if err != nil {
			panic(err)
		}

		// -------------------- service --------------------
		serviceDirectoryPath := fmt.Sprintf("%s/service", directoryPath)

		err = os.MkdirAll(serviceDirectoryPath, 0777)
		if err != nil {
			panic(err)
		}

		err = createFileWithTemplate(
			serviceDirectoryPath,
			"service.go",
			tpl.ServiceTemplate(),
			project,
		)
		if err != nil {
			panic(err)
		}
	},
}
