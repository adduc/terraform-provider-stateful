package internal

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
)

// test ideas for stateful_state data source
// - test that a various types of state can be read

// Check that the state can be read from the provider
func TestStateDataSourceString(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactory(),
		Steps: []resource.TestStep{
			{
				Config: `
					provider "stateful" { state = "hello" }

					data "stateful_state" "state" {}

					output "state" { value = data.stateful_state.state.state }
				`,

				// check that the state is "hello"
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"state",
						knownvalue.StringExact("hello"),
					),
				},
			},
		},
	})
}

// Check that state is kept separate between different instances of the
// provider
func TestStateDataSourceMultipleProviders(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactory(),
		Steps: []resource.TestStep{
			{
				Config: `
					provider "stateful" { 
					  state = "hello" 
					}

					# We would usually use a provider alias here, but
					# the test framework doesn't support that yet. To
					# work around this, we use a different provider name
					provider "stateful2" { 
					  state = "world"
					}

					data "stateful_state" "hello" {
					  provider = "stateful"
					}

					data "stateful_state" "world" {
					  provider = "stateful2"
					}

					output "hello" { value = data.stateful_state.hello.state }
					output "world" { value = data.stateful_state.world.state }
				`,

				// check that the state is "world"
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"hello",
						knownvalue.StringExact("hello"),
					),
					statecheck.ExpectKnownOutputValue(
						"world",
						knownvalue.StringExact("world"),
					),
				},
			},
		},
	})
}

// Check that the same state can be fetched multiple times
func TestStateDataSourceMultipleFetches(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactory(),
		Steps: []resource.TestStep{
			{
				Config: `
					provider "stateful" { 
					  state = "hello" 
					}

					data "stateful_state" "state" {}

					output "state1" { value = data.stateful_state.state.state }
					output "state2" { value = data.stateful_state.state.state }
				`,

				// check that the state is "hello" in both outputs
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"state1",
						knownvalue.StringExact("hello"),
					),
					statecheck.ExpectKnownOutputValue(
						"state2",
						knownvalue.StringExact("hello"),
					),
				},
			},
		},
	})
}
