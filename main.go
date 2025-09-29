package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/fvhockney/terraform-provider-sops/sops"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/fvhockney/sops",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), sops.New, opts)
	if err != nil {
		log.Fatal(err.Error())
	}
}
