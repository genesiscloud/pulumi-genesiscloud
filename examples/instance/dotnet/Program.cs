using Pulumi;
using GenesisCloud.PulumiGenesisCloud;

class GenesisCloudInstance : Stack
{
    public GenesisCloudInstance()
    {
        var region = "ARC-IS-HAF-1";

        var ssh_key = new SSHKey("ssh-key", new SSHKeyArgs
        {
            Name = "ssh-key",
            PublicKey = "<YOUR_SSH_PUBLIC>"
        });

        var floating_ip = new FloatingIp("floating-ip", new FloatingIpArgs
        {
            Name = "my-pulumi-floating-ip",
            Description = "My pulumi floating IP",
            Region = region,
            Version = "ipv4"
        });

        var allow_ssh = new SecurityGroup("allow-ssh", new SecurityGroupArgs
        {
            Name = "allow-ssh",
            Description = "Allow SSH",
            Region = region,
            Rules = new SecurityGroupRuleArgs
            {
              Direction = "ingress",
              Protocol = "tcp",
              PortRangeMin = 22,
              PortRangeMax = 22
            }
        });

        var allow_http = new SecurityGroup("allow-http", new SecurityGroupArgs
        {
            Name = "allow-http",
            Description = "Allow HTTP",
            Region = region,
            Rules = new SecurityGroupRuleArgs
            {
              Direction = "ingress",
              Protocol = "tcp",
              PortRangeMin = 80,
              PortRangeMax = 80
            }
        });

        var startup_script = @"
<<EOF
#!/bin/bash
set -eo pipefail

# Add startup script

EOF"

        var image_id = "2cd0e25f-a39e-4bc6-aa78-b4c40b87072a"

        var instance = new Instance("my-pulumi-instance", new InstanceArgs
        {
            Name = "my-pulumi-instance",
            Region = region,
            Image = image_id,
            PlacementOption = "AUTO",
            DiskSize = 128,
            Type = "vcpu-4_memory-12g_disk-80g_nvidia3080-1",
            SshKeyIds = { ssh_key.Id },
            SecurityGroupIds = { allow_ssh.Id, allow_http.Id },
            Metadata = new InputMap<string>
            {
              { "startup-script", startup_script }
            },
            FloatingIpId = floating_ip.Id
        });
    }
}

class Program
{
    static Task<int> Main(string[] args) => Deployment.RunAsync<GenesisCloudInstance>();
}
