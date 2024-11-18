## Development

> **Warning**: When using the locally-built provider, you do not need to run `terraform init` again after making changes to the provider code. Simply rebuild the provider binary and Terraform will pick up the changes automatically.

### Configuring Terraform to use your locally-built provider

In order to run your locally-built provider, you'll need to configure Terraform to look for the provider binary in a specific location. You can do this by adding a `provider_installation` block to your `~/.terraformrc` file:

```hcl
# @see https://developer.hashicorp.com/terraform/cli/config/config-file
provider_installation {
  dev_overrides {
   "adduc/stateful" = "/path/to/terraform-provider-stateful/bin"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}

```