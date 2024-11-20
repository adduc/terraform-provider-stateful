terraform {
  required_providers {
    stateful = {
      source = "adduc/stateful"
    }
  }
}

variable "input" {
  type = object({
    key1 = string
    key2 = number
    key3 = list(string)
  })
}

output "state" {
  value = var.input
}