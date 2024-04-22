package main

import (
	"github.com/genesiscloud/pulumi-genesiscloud/sdk/go/genesiscloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		region := pulumi.String("ARC-IS-HAF-1")

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

		allowSSH, allowSSHErr := genesiscloud.NewSecurityGroup(ctx, "allow-ssh",
			&genesiscloud.SecurityGroupArgs{
				Name:        pulumi.String("allow-ssh"),
				Description: pulumi.String("Allow SSH"),
				Region:      region,
				Rules: genesiscloud.SecurityGroupRuleArray{
					genesiscloud.SecurityGroupRuleArgs{
						Direction:    pulumi.String("ingress"),
						Protocol:     pulumi.String("tcp"),
						PortRangeMin: pulumi.Int(22),
						PortRangeMax: pulumi.Int(22),
					},
				},
			})
		if allowSSHErr != nil {
			return allowSSHErr
		}

		allowHTTP, allowHTTPErr := genesiscloud.NewSecurityGroup(ctx, "allow-http",
			&genesiscloud.SecurityGroupArgs{
				Name:        pulumi.String("allow-http"),
				Description: pulumi.String("Allow HTTP"),
				Region:      region,
				Rules: genesiscloud.SecurityGroupRuleArray{
					genesiscloud.SecurityGroupRuleArgs{
						Direction:    pulumi.String("ingress"),
						Protocol:     pulumi.String("tcp"),
						PortRangeMin: pulumi.Int(80),
						PortRangeMax: pulumi.Int(80),
					},
				},
			})
		if allowHTTPErr != nil {
			return allowHTTPErr
		}

		_, instanceErr := genesiscloud.NewInstance(ctx, "my-pulumi-instance", &genesiscloud.InstanceArgs{
			Region:           region,
			Image:            pulumi.String("ubuntu-ml-nvidia-pytorch"),
			Name:             pulumi.String("my-pulumi-instance"),
			DiskSize:         pulumi.Int(128),
			Type:             pulumi.String("vcpu-4_memory-12g_disk-80g_nvidia3080-1"),
			SshKeyIds:        pulumi.StringArray{sshKey.ID()},
			SecurityGroupIds: pulumi.StringArray{allowSSH.ID(), allowHTTP.ID()},
			FloatingIpId:     floatingIP.ID(),
		})
		if instanceErr != nil {
			return instanceErr
		}

		return nil
	})
}
