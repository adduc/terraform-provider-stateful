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

## Modules

module "state" {
  source = "./module"
}

## Outputs

output "state" {
  value = module.state
}
