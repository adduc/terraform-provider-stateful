## Required Providers

terraform {
  required_providers {
    stateful = {
      source  = "adduc/stateful"
      version = "0-dev"
    }
  }
}

## Provider Configuration

provider "stateful" {
  state = {
    key1 = "value1"
    key2 = "value2"
  }
}

## Data Sources

data "stateful_state" "state" {}

## Outputs

output "state" {
  value = data.stateful_state.state
}
