"""A Python Pulumi program"""

import pulumi
import pulumi_genesiscloud as gc

region = "ARC-IS-HAF-1"

# const sshKey = new SSHKey("philip", {
#   name: "philip",
#   publicKey:
#     "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIG0EyNBL6fu1IPhWLvO6njP6/cEWYCMCv/uTBQgdSM7Q barista@roastery",
# });

ssh_key = gc.SSHKey(
    "philip",
    name="philip",
    public_key="ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIG0EyNBL6fu1IPhWLvO6njP6/cEWYCMCv/uTBQgdSM7Q barista@roastery",
)

floating_ip = gc.FloatingIp(
    resource_name="my-pulumi-floating-ip",
    name="my-pulumi-floating-ip",
    description="My pulumi floating IP",
    region=region,
    version="ipv4",
)


allow_ssh = gc.SecurityGroup(
    "allow-ssh",
    name="allow-ssh",
    description="Allow SSH",
    region=region,
    rules=[
        gc.SecurityGroupRuleArgs(
            direction="ingress",
            protocol="tcp",
            port_range_min=22,
            port_range_max=22,
        )
    ],
)

allow_http = gc.SecurityGroup(
    "allow-http",
    name="allow-http",
    description="Allow HTTP",
    region=region,
    rules=[
        gc.SecurityGroupRuleArgs(
            direction="ingress",
            protocol="tcp",
            port_range_min=80,
            port_range_max=80,
        )
    ],
)

startup_script = """
<<EOF
#!/bin/bash
set -eo pipefail

# Add startup script

EOF
"""

image_id = "2cd0e25f-a39e-4bc6-aa78-b4c40b87072a"

instance = gc.Instance(
    "my-pulumi-instance",
    name="my-pulumi-instance",
    region=region,
    image=image_id,
    placement_option="AUTO",
    type="vcpu-4_memory-12g_disk-80g_nvidia3080-1",
    ssh_key_ids=[ssh_key.id],
    security_group_ids=[allow_ssh.id, allow_http.id],
    metadata={"startup-script": startup_script},
    floating_ip_id=floating_ip.id,
)

# pulumi.export("ssh_key", ssh_key.public_key)
# pulumi.export("floating_ip", floating_ip.ip_address)
