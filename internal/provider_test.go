package internal

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func protoV6ProviderFactory() map[string]func() (tfprotov6.ProviderServer, error) {
	return map[string]func() (tfprotov6.ProviderServer, error){
		"stateful":  providerserver.NewProtocol6WithError(New("test")()),
		"stateful2": providerserver.NewProtocol6WithError(New("test2")()),
	}
}

func TestStatefulProviderMissingConfig(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactory(),
		Steps: []resource.TestStep{
			{
				Config: `
					# check that an error is returned when the state
					# argument is not provided
					provider "stateful" {}

					# data source provided to allow the provider to be
					# initialized
					data "stateful_state" "state" {}
				`,
				ExpectError: regexp.MustCompile(`The argument "state" is required, but no definition was found.`),
			},
		},
	})
}

func TestStatefulProviderConfigured(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactory(),
		Steps: []resource.TestStep{
			{
				Config: `
					# check that the provider accepts the state argument
					provider "stateful" {
						state = {
							"key" = "value"
						}
					}

					data "stateful_state" "state" {}

					output "state" { value = data.stateful_state.state }
				`,
			},
		},
	})
}
