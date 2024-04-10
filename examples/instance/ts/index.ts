import {
  SecurityGroup,
  SSHKey,
  Instance,
  FloatingIp,
} from "@genesiscloud/pulumi-genesiscloud";

const region = "ARC-IS-HAF-1";

const sshKey = new SSHKey("ssh-key", {
  name: "ssh-key",
  publicKey: "<YOUR_SSH_PUBLIC_KEY>",
});

const floatingIP = new FloatingIp("my-pulumi-floating-ip", {
  name: "my-pulumi-floating-ip",
  description: "my-pulumi-floating-ip",
  region,
  version: "ipv4",
});

const allowSSH = new SecurityGroup("allow-ssh", {
  name: "allow-ssh",
  region,
  description: "Allow SSH access",
  rules: [
    {
      direction: "ingress",
      protocol: "tcp",
      portRangeMin: 22,
      portRangeMax: 22,
    },
  ],
});

const allowHTTP = new SecurityGroup("allow-http", {
  name: "allow-http",
  region,
  description: "Allow HTTP access",
  rules: [
    {
      direction: "ingress",
      protocol: "tcp",
      portRangeMin: 80,
      portRangeMax: 80,
    },
  ],
});

const startupScript = `
<<EOF
#!/bin/bash
set -eo pipefail

# Add startup script

EOF
`;

const firstPulumiInstance = new Instance("first-pulumi-instance", {
  name: "first-pulumi-instance",
  region,
  image: "ubuntu-ml-nvidia-pytorch",
  type: "vcpu-4_memory-12g_disk-80g_nvidia3080-1",
  diskSize: 128,
  sshKeyIds: [sshKey.id],
  securityGroupIds: [allowSSH.id, allowHTTP.id],
  metadata: {
    startupScript,
  },
  floatingIpId: floatingIP.id,
});
