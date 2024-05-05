terraform {
  required_providers {
    stateful = {
      source  = "adduc/stateful"
      version = "0-dev"
    }
  }
}

provider "stateful" {
  state = {
    key1 = "value1"
    key2 = "value2"
  }
}

data "stateful_state" "state" {}

module "state" {
  source = "./module"
}

output "state" {
  value = data.stateful_state.state
}

output "module_state" {
  value = module.state.state
}