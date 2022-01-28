package main

import (
	"os"

	"github.com/pylons-tech/juno/cmd/parse"

	"github.com/pylons-tech/juno/modules/messages"
	"github.com/pylons-tech/juno/modules/registrar"

	"github.com/pylons-tech/juno/cmd"
)

func main() {
	// JunoConfig the runner
	config := cmd.NewConfig("juno").
		WithParseConfig(parse.NewConfig().
			WithRegistrar(registrar.NewDefaultRegistrar(
				messages.CosmosMessageAddressesParser,
			)),
		)

	// Run the commands and panic on any error
	exec := cmd.BuildDefaultExecutor(config)
	err := exec.Execute()
	if err != nil {
		os.Exit(1)
	}
}
