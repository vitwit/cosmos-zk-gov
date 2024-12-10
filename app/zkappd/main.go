package main

import (
	"fmt"
	"os"

	simapp "github.com/vitwit/cosmos-zk-gov/app"
	"github.com/vitwit/cosmos-zk-gov/app/zkappd/cmd"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", simapp.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
