package internal

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// test that the terraform provider StatefulProvider can be initialized
// with no configuration
func TestStatefulProvider(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"stateful": providerserver.NewProtocol6WithError(New("test")()),
		},
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
