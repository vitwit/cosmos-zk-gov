package main

import (
	"fmt"
	"os"

	simapp "github.com/vishal-kanna/zk/zk-gov/app"
	"github.com/vishal-kanna/zk/zk-gov/app/simd/cmd"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", simapp.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
