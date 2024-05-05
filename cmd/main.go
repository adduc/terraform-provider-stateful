package main

import (
	"context"
	"flag"
	"log"

	"github.com/adduc/terraform-provider-stateful/internal"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

//go:generate terraform fmt -recursive ./examples/
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate -provider-name scaffolding

var (
	// this will be set by goreleaser
	// @see https://goreleaser.com/cookbooks/using-main.version/
	version string = "dev"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set true to run with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/adduc/stateful",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), internal.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
