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
  alias = "state1"
  state = {
    key1 = "value1"
    key2 = "value2"
  }
}

provider "stateful" {
  alias = "state2"
  state = {
    key1 = "value1"
    key2 = "value2"
  }
}

## Data Sources

data "stateful_state" "state1" {
  provider = stateful.state1
}

data "stateful_state" "state2" {
  provider = stateful.state2
}

## Outputs

output "state1" {
  value = data.stateful_state.state1
}

output "state2" {
  value = data.stateful_state.state2
}