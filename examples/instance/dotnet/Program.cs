using Pulumi;
using GenesisCloud.PulumiPackage.Genesiscloud;
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

        var instance = new Instance("my-pulumi-instance", new InstanceArgs
        {
            Name = "my-pulumi-instance",
            Region = region,
            Image = "ubuntu-ml-nvidia-pytorch",
            PlacementOption = "AUTO",
            DiskSize = 128,
            Type = "vcpu-4_memory-12g_disk-80g_nvidia3080-1",
            SshKeyIds = { ssh_key.Id },
            FloatingIpId = floating_ip.Id
        });
    }
}

class Program
{
    static Task<int> Main(string[] args) => Deployment.RunAsync<GenesisCloudInstance>();
}
