# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .filesystem import *
from .floating_ip import *
from .images import *
from .instance import *
from .instance_status import *
from .provider import *
from .security_group import *
from .snapshot import *
from .ssh_key import *
from .volume import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_genesiscloud.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_genesiscloud.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "genesiscloud",
  "mod": "index/filesystem",
  "fqn": "pulumi_genesiscloud",
  "classes": {
   "genesiscloud:index/filesystem:Filesystem": "Filesystem"
  }
 },
 {
  "pkg": "genesiscloud",
  "mod": "index/floatingIp",
  "fqn": "pulumi_genesiscloud",
  "classes": {
   "genesiscloud:index/floatingIp:FloatingIp": "FloatingIp"
  }
 },
 {
  "pkg": "genesiscloud",
  "mod": "index/instance",
  "fqn": "pulumi_genesiscloud",
  "classes": {
   "genesiscloud:index/instance:Instance": "Instance"
  }
 },
 {
  "pkg": "genesiscloud",
  "mod": "index/instanceStatus",
  "fqn": "pulumi_genesiscloud",
  "classes": {
   "genesiscloud:index/instanceStatus:InstanceStatus": "InstanceStatus"
  }
 },
 {
  "pkg": "genesiscloud",
  "mod": "index/sSHKey",
  "fqn": "pulumi_genesiscloud",
  "classes": {
   "genesiscloud:index/sSHKey:SSHKey": "SSHKey"
  }
 },
 {
  "pkg": "genesiscloud",
  "mod": "index/securityGroup",
  "fqn": "pulumi_genesiscloud",
  "classes": {
   "genesiscloud:index/securityGroup:SecurityGroup": "SecurityGroup"
  }
 },
 {
  "pkg": "genesiscloud",
  "mod": "index/snapshot",
  "fqn": "pulumi_genesiscloud",
  "classes": {
   "genesiscloud:index/snapshot:Snapshot": "Snapshot"
  }
 },
 {
  "pkg": "genesiscloud",
  "mod": "index/volume",
  "fqn": "pulumi_genesiscloud",
  "classes": {
   "genesiscloud:index/volume:Volume": "Volume"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "genesiscloud",
  "token": "pulumi:providers:genesiscloud",
  "fqn": "pulumi_genesiscloud",
  "class": "Provider"
 }
]
"""
)
