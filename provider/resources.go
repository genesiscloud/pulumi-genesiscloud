package genesiscloud

import (
	"fmt"
	"path"

	_ "embed"

	tfpf "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"

	"github.com/genesiscloud/pulumi-genesiscloud/provider/pkg/version"
	"github.com/genesiscloud/terraform-provider-genesiscloud/genesiscloudshim"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "genesiscloud"
	// modules:
	mainMod = "index" // the genesiscloud module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
// func preConfigureCallback(resource.PropertyMap, shim.ResourceConfig) error {
// 	return nil
// }

//go:embed cmd/pulumi-resource-genesiscloud/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Version := "0.0.1"
	Version := version.Version
	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		// Instantiate the Terraform provider
		P:       tfpf.ShimProvider(genesiscloudshim.NewProvider(Version)),
		Name:    "genesiscloud",
		Version: Version,
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		// DisplayName: "",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "Genesis Cloud",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "",
		Description:       "A Pulumi package for creating and managing genesiscloud cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "genesiscloud", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.genesiscloud.com",
		Repository: "https://github.com/genesiscloud/pulumi-genesiscloud",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg:    "genesiscloud",
		MetadataInfo: tfbridge.NewProviderMetadata(metadata),
		Config:       map[string]*tfbridge.SchemaInfo{
			// Add any rquired configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"genesiscloud_instance":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Instance")},
			"genesiscloud_security_group": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SecurityGroup")},
			"genesiscloud_snapshot":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Snapshot")},
			"genesiscloud_ssh_key":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SSHKey")},
			"genesiscloud_volume":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Volume")},
			"genesiscloud_floating_ip":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FloatingIp")},
			"genesiscloud_filesystem":     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Filesystem")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"genesiscloud_images": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "Images")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			// Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				fmt.Sprintf("github.com/genesiscloud/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: map[string]string{
				"genesiscloud": "GenesisCloud",
			},
		},
	}

	// These are new API's that you may opt to use to automatically compute resource
	// tokens, and apply auto aliasing for full backwards compatibility.  For more
	// information, please reference:
	// https://pkg.go.dev/github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge#ProviderInfo.ComputeTokens
	prov.MustComputeTokens(tokens.SingleModule("genesiscloud_", mainMod,
		tokens.MakeStandard(mainPkg)))
	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
