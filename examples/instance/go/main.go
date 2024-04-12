package main

import (
	"github.com/genesiscloud/pulumi-genesiscloud/sdk/go/genesiscloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		region := "ARC-IS-HAF-1"

		sshKey, sshKeyErr := genesiscloud.NewSSHKey(ctx, "ssh-key", &genesiscloud.SSHKeyArgs{
			PublicKey: pulumi.String("<YOUR_SSH_PUBLIC_KEY>"),
		})
		if sshKeyErr != nil {
			return sshKeyErr
		}

		floatingIP, floatingIPErr := genesiscloud.NewFloatingIp(ctx, "my-pulumi-floating-ip", &genesiscloud.FloatingIpArgs{
			Description: pulumi.String("My pulumi floating IP"),
			Region:      pulumi.String(region),
			Version:     pulumi.String("ipv4"),
		})
		if floatingIPErr != nil {
			return floatingIPErr
		}

		_, instanceErr := genesiscloud.NewInstance(ctx, "my-pulumi-instance", &genesiscloud.InstanceArgs{
			Region:           pulumi.String(region),
			Image:            pulumi.String("ubuntu-ml-nvidia-pytorch"),
			Name:             pulumi.String("my-pulumi-instance"),
			DiskSize:         pulumi.Int(128),
			Type:             pulumi.String("vcpu-4_memory-12g_disk-80g_nvidia3080-1"),
			SshKeyIds:        pulumi.StringArray{sshKey.ID()},
			SecurityGroupIds: pulumi.StringArray(nil),
			FloatingIpId:     floatingIP.ID(),
		})
		if instanceErr != nil {
			return instanceErr
		}

		return nil
	})
}
