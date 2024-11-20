terraform {
  required_providers {
    stateful = {
      source = "adduc/stateful"
    }
  }
}

data "stateful_state" "state" {}

output "state" {
  value = data.stateful_state.state
}