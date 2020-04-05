package cli

import (
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	initCmdPkgName string
)

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.PersistentFlags().StringVar(&initCmdPkgName, "name", "", "name for service (default is the base of path given)")
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize skelton for mskit",
	Long:  "Initialize skelton for mskit",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var pkgName string

		directoryPath := args[0]

		if initCmdPkgName != "" {
			pkgName = initCmdPkgName
		} else {
			pkgDestPath := directoryPath

			if pkgDestPath == "." {
				pkgDestPath, err = filepath.Abs(pkgDestPath)
				if err != nil {
					panic(err)
				}
			}

			pkgName = filepath.Base(pkgDestPath)
		}

		err = project.initializeService(directoryPath, pkgName)
		if err != nil {
			panic(err)
		}
	},
}
