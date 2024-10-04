package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"discord-terraform/discord"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website
//go:generate terraform fmt -recursive ./examples/
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate -provider-name discord

var (
	version string = "dev"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: discord.Provider(version),
		Debug:        debugMode,
	})
}
