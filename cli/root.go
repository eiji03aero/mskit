package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	project *Project

	rootCmd = &cobra.Command{
		Use:   "mskit",
		Short: "Provides utility cli tools for mskit",
		Long:  "Provides utility cli tools for mskit",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("root called")
		},
	}
)

func init() {
	project = newProject()
	err := project.initialize()
	if err != nil {
		panic(err)
	}
}

func Execute() error {
	return rootCmd.Execute()
}
