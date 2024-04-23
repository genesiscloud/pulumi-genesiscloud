using Pulumi;
using GenesisCloud.PulumiPackage.Genesiscloud;
using GenesisCloud.PulumiPackage.Genesiscloud.Inputs;
using System.Threading.Tasks;

class GenesisCloudInstance : Stack
{
    public GenesisCloudInstance()
    {
        var region = "ARC-IS-HAF-1";

        var ssh_key = new SSHKey("ssh-key", new SSHKeyArgs
        {
            Name = "ssh-key",
            PublicKey = "<YOUR_SSH_PUBLIC_KEY>"
        });

        var floating_ip = new FloatingIp("floating-ip", new FloatingIpArgs
        {
            Name = "my-pulumi-floating-ip",
            Description = "My pulumi floating IP",
            Region = region,
            Version = "ipv4"
        });

        var allow_ssh = new SecurityGroup("allow_ssh", new SecurityGroupArgs
        {
            Name = "allow-ssh",
            Description = "Allow SSH",
            Region = region,
            Rules = {
              new SecurityGroupRuleArgs
              {
                Direction = "ingress",
                Protocol = "tcp",
                PortRangeMin = 22,
                PortRangeMax = 22,
              }
            }
        });

        var allow_http = new SecurityGroup("allow_http", new SecurityGroupArgs
        {
            Name = "allow-http",
            Description = "Allow HTTP",
            Region = region,
            Rules = {
              new SecurityGroupRuleArgs
              {
                Direction = "ingress",
                Protocol = "tcp",
                PortRangeMin = 80,
                PortRangeMax = 80,
              }
            }
        });

        var instance = new Instance("my-pulumi-instance", new InstanceArgs
        {
            Name = "my-pulumi-instance",
            Region = region,
            Image = "ubuntu-ml-nvidia-pytorch",
            DiskSize = 128,
            Type = "vcpu-4_memory-12g_disk-80g_nvidia3080-1",
            SshKeyIds = { ssh_key.Id },
            FloatingIpId = floating_ip.Id
            SecurityGroupIds = { allow_ssh.Id, allow_http.Id },
        });
    }
}

class Program
{
    static Task<int> Main(string[] args) => Deployment.RunAsync<GenesisCloudInstance>();
}
