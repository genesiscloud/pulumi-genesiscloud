module github.com/genesiscloud/terraform-provider-genesiscloud/genesiscloudshim

go 1.23.0

toolchain go1.23.1

replace github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20230912190043-e6d96b3b8f7e

require (
	github.com/genesiscloud/terraform-provider-genesiscloud v1.1.13
	github.com/hashicorp/terraform-plugin-framework v1.14.1
)

require (
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/genesiscloud/genesiscloud-go v1.0.16 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.7 // indirect
	github.com/hashicorp/terraform-plugin-framework-timeouts v0.5.0 // indirect
	github.com/hashicorp/terraform-plugin-framework-validators v0.18.0 // indirect
	github.com/hashicorp/terraform-plugin-go v0.26.0 // indirect
	github.com/hashicorp/terraform-plugin-log v0.9.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/oapi-codegen/oapi-codegen/v2 v2.4.1 // indirect
	github.com/oapi-codegen/runtime v1.1.1 // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
)
