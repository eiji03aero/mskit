package main

import (
	"fmt"
	"os"

	"github.com/eiji03aero/mskit/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
