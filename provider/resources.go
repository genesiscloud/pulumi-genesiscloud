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

//go:embed cmd/pulumi-resource-genesiscloud/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	Version := version.Version
	prov := tfbridge.ProviderInfo{
		P:                 tfpf.ShimProvider(genesiscloudshim.NewProvider(Version)),
		DisplayName:       "Genesis Cloud",
		Name:              "genesiscloud",
		Description:       "A Pulumi package for creating and managing genesiscloud cloud resources.",
		Keywords:          []string{"pulumi", "genesiscloud", "category/cloud"},
		License:           "Apache-2.0",
		Homepage:          "https://www.pulumi.io",
		Repository:        "https://github.com/genesiscloud/pulumi-genesiscloud",
		PluginDownloadURL: "github://api.github.com/genesiscloud",
		LogoURL:           "https://avatars.githubusercontent.com/u/38134186?s=200&v=4",
		Version:           Version,
		GitHubOrg:         "genesiscloud",
		MetadataInfo:      tfbridge.NewProviderMetadata(metadata),
		Config: map[string]*tfbridge.SchemaInfo{
			"token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GENESISCLOUD_TOKEN"},
				},
			},
			"endpoint": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GENESISCLOUD_ENDPOINT"},
				},
			},
			// Add any rquired configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("regionu", "Region"),
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
			PackageName: "@genesiscloud/pulumi-genesiscloud",
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
			},
		},
		Python: &tfbridge.PythonInfo{
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
				"genesiscloud": "genesiscloud",
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
