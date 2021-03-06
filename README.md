# terraform-provider-kong
[![CircleCI](https://circleci.com/gh/alexashley/terraform-provider-kong/tree/master.svg?style=svg)](https://circleci.com/gh/alexashley/terraform-provider-kong/tree/master)

A Terraform provider for the api gateway [Kong](https://github.com/Kong/kong).

- [example](/examples/simple/main.tf)
- [documentation](https://alexashley.github.io/terraform-provider-kong/)

## Features

- Supports Kong's Enterprise authentication
- Resources for individual plugins, including EE-only plugins 
- Import individual consumers, services, routes, or plugins
- [Bulk import tool](./importer/README.md) to ease migrating existing infrastructure into Terraform

## Unsupported Resources

No plans to support the legacy [API resource](https://docs.konghq.com/0.12.x/admin-api/#api-object).

I also don't intend to add support for the following resources, since I don't currently use them:

- [Certificate](https://docs.konghq.com/0.14.x/admin-api/#certificate-object)
- [SNI](https://docs.konghq.com/0.14.x/admin-api/#sni-objects)
- [Upstream](https://docs.konghq.com/0.14.x/admin-api/#upstream-objects)
- [Target](https://docs.konghq.com/0.14.x/admin-api/#target-object)

That said, PRs are welcome and may I try to implement them once the other resources are mature.

## Development

- install Go 1.11 (this project uses [Go modules](https://github.com/golang/go/wiki/Modules#installing-and-activating-module-support))
- install Terraform.
- `docker-compose up --build -d` to stand up an instance of Kong 0.14 CE w/ Postgres.
- `make build` to create the provider.
- `make example` to run the provider against Kong.

### Adding a new resource
- If you're adding a new plugin, use the `generic_plugin_resource` to abstract the Terraform operations for the plugin.
    - Supply one function that maps the resource to the plugin config (`provider.ConfigMapper`)
    - And another function to map from the plugin's API representation to Terraform state (`provider.ResourceMapper`)
- For all resources, add a new YAML file under `docsgen/resource-metadata` with information about the resource
    - Documentation is generated by reading the provider, so you must supply a description for each field
