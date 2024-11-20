##
# Exercise: Types Inputs Example
#
# This exercise demonstrates how the configuration passed to the
# `stateful` provider keeps the types of the values passed in the
# `state` argument.
##

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
    key2 = 2
    key3 = ["value3", "value4"]
  }
}

## Data Sources

data "stateful_state" "state" {}

## Modules

module "state" {
  source = "./module"
  input  = data.stateful_state.state.state
}

## Outputs

output "state" {
  value = module.state
}
