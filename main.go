package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/sumup/terraform-provider-rundeck/rundeck"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: rundeck.Provider})
}
