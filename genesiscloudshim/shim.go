package genesiscloudshim

import (
	"github.com/genesiscloud/terraform-provider-genesiscloud/internal/provider"
	tfpf "github.com/hashicorp/terraform-plugin-framework/provider"
)

func NewProvider(version string) tfpf.Provider {
	return provider.New(version)()
}
