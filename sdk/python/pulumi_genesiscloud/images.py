# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'ImagesResult',
    'AwaitableImagesResult',
    'images',
    'images_output',
]

@pulumi.output_type
class ImagesResult:
    """
    A collection of values returned by Images.
    """
    def __init__(__self__, filter=None, id=None, images=None, timeouts=None):
        if filter and not isinstance(filter, dict):
            raise TypeError("Expected argument 'filter' to be a dict")
        pulumi.set(__self__, "filter", filter)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if images and not isinstance(images, list):
            raise TypeError("Expected argument 'images' to be a list")
        pulumi.set(__self__, "images", images)
        if timeouts and not isinstance(timeouts, dict):
            raise TypeError("Expected argument 'timeouts' to be a dict")
        pulumi.set(__self__, "timeouts", timeouts)

    @property
    @pulumi.getter
    def filter(self) -> 'outputs.ImagesFilterResult':
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def images(self) -> Sequence['outputs.ImagesImageResult']:
        return pulumi.get(self, "images")

    @property
    @pulumi.getter
    def timeouts(self) -> Optional['outputs.ImagesTimeoutsResult']:
        return pulumi.get(self, "timeouts")


class AwaitableImagesResult(ImagesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ImagesResult(
            filter=self.filter,
            id=self.id,
            images=self.images,
            timeouts=self.timeouts)


def images(filter: Optional[pulumi.InputType['ImagesFilterArgs']] = None,
           timeouts: Optional[pulumi.InputType['ImagesTimeoutsArgs']] = None,
           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableImagesResult:
    """
    Images data source

    ## Example Usage

    ```python
    import pulumi
    import pulumi_genesiscloud as genesiscloud

    cloud_images = genesiscloud.images(filter=genesiscloud.ImagesFilterArgs(
        type="cloud-image",
    ))
    snapshots = genesiscloud.images(filter=genesiscloud.ImagesFilterArgs(
        region="ARC-IS-HAF-1",
        type="snapshot",
    ))
    preconfigured_images = genesiscloud.images(filter=genesiscloud.ImagesFilterArgs(
        type="preconfigured",
    ))
    ```
    """
    __args__ = dict()
    __args__['filter'] = filter
    __args__['timeouts'] = timeouts
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('genesiscloud:index/images:Images', __args__, opts=opts, typ=ImagesResult).value

    return AwaitableImagesResult(
        filter=pulumi.get(__ret__, 'filter'),
        id=pulumi.get(__ret__, 'id'),
        images=pulumi.get(__ret__, 'images'),
        timeouts=pulumi.get(__ret__, 'timeouts'))


@_utilities.lift_output_func(images)
def images_output(filter: Optional[pulumi.Input[pulumi.InputType['ImagesFilterArgs']]] = None,
                  timeouts: Optional[pulumi.Input[Optional[pulumi.InputType['ImagesTimeoutsArgs']]]] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ImagesResult]:
    """
    Images data source

    ## Example Usage

    ```python
    import pulumi
    import pulumi_genesiscloud as genesiscloud

    cloud_images = genesiscloud.images(filter=genesiscloud.ImagesFilterArgs(
        type="cloud-image",
    ))
    snapshots = genesiscloud.images(filter=genesiscloud.ImagesFilterArgs(
        region="ARC-IS-HAF-1",
        type="snapshot",
    ))
    preconfigured_images = genesiscloud.images(filter=genesiscloud.ImagesFilterArgs(
        type="preconfigured",
    ))
    ```
    """
    ...
