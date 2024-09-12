package genesiscloud

import (
	"bytes"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
)

var replaceTerraformWithPulumi = tfbridge.DocsEdit{
	Path: "*",
	Edit: func(_ string, content []byte) ([]byte, error) {
		content = bytes.ReplaceAll(content, []byte("Terraform"), []byte("Pulumi"))

		return content, nil
	},
}

func editRules(defaults []tfbridge.DocsEdit) []tfbridge.DocsEdit {
	return append(defaults,
		replaceTerraformWithPulumi,
	)
}
