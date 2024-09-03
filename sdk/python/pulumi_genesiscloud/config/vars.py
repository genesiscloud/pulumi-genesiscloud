# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

import types

__config__ = pulumi.Config('genesiscloud')


class _ExportableConfig(types.ModuleType):
    @property
    def endpoint(self) -> Optional[str]:
        """
        Genesis Cloud API endpoint. May also be provided via `GENESISCLOUD_ENDPOINT` environment variable. If neither is
        provided, defaults to `https://api.genesiscloud.com/compute/v1`.
        """
        return __config__.get('endpoint') or _utilities.get_env('GENESISCLOUD_ENDPOINT')

    @property
    def polling_interval(self) -> Optional[str]:
        """
        The polling interval. - The string must be a positive [time duration](https://pkg.go.dev/time#ParseDuration), for
        example "10s".
        """
        return __config__.get('pollingInterval')

    @property
    def token(self) -> Optional[str]:
        """
        Genesis Cloud API token. May also be provided via `GENESISCLOUD_TOKEN` environment variable.
        """
        return __config__.get('token') or _utilities.get_env('GENESISCLOUD_TOKEN')

