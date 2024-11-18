terraform {
  required_providers {
    stateful = {
      source  = "adduc/stateful"
      version = "0-dev"
    }
  }
}

data "stateful_state" "state" {}

output "state" {
  value = data.stateful_state.state
}