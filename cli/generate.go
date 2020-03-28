package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/eiji03aero/mskit/cli/tpl"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)

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

var generateRPCEndpointCmd = &cobra.Command{
	Use:   "rpcendpoint",
	Short: "Generates rpcendpoint template",
	Long:  "Generates rpcendpoint template",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		dir := fmt.Sprintf("%s/transport/rpcendpoint", project.WorkingDir)

		if _, err = os.Stat(dir); os.IsNotExist(err) {
			if err = os.MkdirAll(dir, 0777); err != nil {
				panic(err)
			}
		}

		err = createFileWithTemplate(
			dir,
			"rpcendpoint.go",
			tpl.RPCEndpointTemplate(),
			project,
		)
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
		var err error
		name := args[0]
		rootDir := project.WorkingDir
		data := struct {
			PkgName       string
			Name          string
			LowerName     string
			InterfaceName string
		}{
			PkgName:       project.PkgName,
			Name:          name,
			LowerName:     strings.ToLower(name),
			InterfaceName: fmt.Sprintf("%sProxy", name),
		}

		// -------------------- root --------------------
		fileName := "proxy.go"
		rootFilePath := fmt.Sprintf("%s/%s", rootDir, fileName)
		if _, err = os.Stat(rootFilePath); os.IsNotExist(err) {
			err = createFileWithTemplate(
				rootDir,
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
		dir := fmt.Sprintf("%s/transport/proxy/%s", rootDir, data.LowerName)
		if _, err = os.Stat(dir); os.IsNotExist(err) {
			if err = os.MkdirAll(dir, 0777); err != nil {
				panic(err)
			}
		}

		err = createFileWithTemplate(
			dir,
			"proxy.go",
			tpl.ProxyTemplate(),
			data,
		)
		if err != nil {
			panic(err)
		}

		err = createFileWithTemplate(
			dir,
			fmt.Sprintf("%s.go", data.LowerName),
			tpl.ProxyImplTemplate(),
			data,
		)
		if err != nil {
			panic(err)
		}
	},
}
