package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/simapp"

	"github.com/pylons-tech/juno/cmd"
	"github.com/pylons-tech/juno/config"
	stddb "github.com/pylons-tech/juno/db/builder"
)

func main() {
	// Register modules
	// registrar.RegisterModules(staking.Module{}, consensus.Module{}, ...)

	// Build the exec
	exec := cmd.BuildDefaultExecutor("juno", config.DefaultSetup, simapp.MakeCodec, stddb.Builder)

	// Run the commands and panic on any error
	err := exec.Execute()
	if err != nil {
		os.Exit(1)
	}
}
