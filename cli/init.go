package cli

import (
	"path/filepath"

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
		directoryPath := args[0]
		pkgName := filepath.Base(directoryPath)
		err := project.initializeService(directoryPath, pkgName)
		if err != nil {
			panic(err)
		}
	},
}
