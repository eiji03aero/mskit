package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.AddCommand(generateAggregateCmd)
	generateCmd.AddCommand(generatePublisherCmd)
	generateCmd.AddCommand(generateConsumerCmd)
	generateCmd.AddCommand(generateRPCEndpointCmd)
	generateCmd.AddCommand(generateProxyCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Provides utility to generate assets",
	Long:  "Provides utility to generate assets",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
	},
}

var generateAggregateCmd = &cobra.Command{
	Use:   "aggregate",
	Short: "Generates aggregate template",
	Long:  "Generates aggregate template",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		err := project.generateAggregate(name)
		if err != nil {
			panic(err)
		}
	},
}

var generatePublisherCmd = &cobra.Command{
	Use:   "publisher",
	Short: "Generates publisher template",
	Long:  "Generates publisher template",
	Run: func(cmd *cobra.Command, args []string) {
		err := project.generatePublisher()
		if err != nil {
			panic(err)
		}
	},
}

var generateConsumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Generates consumer template",
	Long:  "Generates consumer template",
	Run: func(cmd *cobra.Command, args []string) {
		err := project.generateConsumer()
		if err != nil {
			panic(err)
		}
	},
}

var generateRPCEndpointCmd = &cobra.Command{
	Use:   "rpcendpoint",
	Short: "Generates rpcendpoint template",
	Long:  "Generates rpcendpoint template",
	Run: func(cmd *cobra.Command, args []string) {
		err := project.generateRPCEndpoint()
		if err != nil {
			panic(err)
		}
	},
}

var generateProxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Generates proxy template",
	Long:  "Generates proxy template",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		err := project.generateProxy(name)
		if err != nil {
			panic(err)
		}
	},
}
