# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

endpoint: Optional[str]
"""
Genesis Cloud API endpoint. May also be provided via `GENESISCLOUD_ENDPOINT` environment variable. If neither is
provided, defaults to `https://api.genesiscloud.com/compute/v1`.
"""

pollingInterval: Optional[str]
"""
The polling interval. - The string must be a positive [time duration](https://pkg.go.dev/time#ParseDuration), for
example "10s".
"""

token: Optional[str]
"""
Genesis Cloud API token. May also be provided via `GENESISCLOUD_TOKEN` environment variable.
"""
