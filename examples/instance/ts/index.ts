import * as pulumi from "@pulumi/pulumi";

import {
  SecurityGroup,
  SSHKey,
  Instance,
  FloatingIp,
} from "@pulumi/genesiscloud";

const region = "ARC-IS-HAF-1";

const sshKey = new SSHKey("philip", {
  name: "philip",
  publicKey:
    "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIG0EyNBL6fu1IPhWLvO6njP6/cEWYCMCv/uTBQgdSM7Q barista@roastery",
});

const floatingIP = new FloatingIp("floating-ip", {
  name: "floating-ip",
  description: "floating-ip",
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

const imageId = "2cd0e25f-a39e-4bc6-aa78-b4c40b87072a";

const firstPulumiInstance = new Instance("first-pulumi-instance", {
  name: "first-pulumi-instance",
  region,
  image: imageId,
  placementOption: "AUTO",
  type: "vcpu-4_memory-12g_disk-80g_nvidia3080-1",
  sshKeyIds: [sshKey.id],
  securityGroupIds: [allowSSH.id, allowHTTP.id],
  metadata: {
    startupScript,
  },
  floatingIpId: floatingIP.id,
});

// TODO: create output using the ssh as output to connect to the instance
