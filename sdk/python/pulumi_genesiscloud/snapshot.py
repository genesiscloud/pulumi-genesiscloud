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
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['SnapshotArgs', 'Snapshot']

@pulumi.input_type
class SnapshotArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 retain_on_delete: Optional[pulumi.Input[bool]] = None,
                 timeouts: Optional[pulumi.Input['SnapshotTimeoutsArgs']] = None):
        """
        The set of arguments for constructing a Snapshot resource.
        :param pulumi.Input[str] instance_id: The id of the instance to snapshot. - If the value of this attribute changes, the resource will be replaced.
        :param pulumi.Input[str] name: The human-readable name for the snapshot.
        :param pulumi.Input[bool] retain_on_delete: Flag to retain the snapshot when the resource is deleted. - Sets the default value "false" if the attribute is not set.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if retain_on_delete is not None:
            pulumi.set(__self__, "retain_on_delete", retain_on_delete)
        if timeouts is not None:
            pulumi.set(__self__, "timeouts", timeouts)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The id of the instance to snapshot. - If the value of this attribute changes, the resource will be replaced.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The human-readable name for the snapshot.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="retainOnDelete")
    def retain_on_delete(self) -> Optional[pulumi.Input[bool]]:
        """
        Flag to retain the snapshot when the resource is deleted. - Sets the default value "false" if the attribute is not set.
        """
        return pulumi.get(self, "retain_on_delete")

    @retain_on_delete.setter
    def retain_on_delete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "retain_on_delete", value)

    @property
    @pulumi.getter
    def timeouts(self) -> Optional[pulumi.Input['SnapshotTimeoutsArgs']]:
        return pulumi.get(self, "timeouts")

    @timeouts.setter
    def timeouts(self, value: Optional[pulumi.Input['SnapshotTimeoutsArgs']]):
        pulumi.set(self, "timeouts", value)


@pulumi.input_type
class _SnapshotState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 retain_on_delete: Optional[pulumi.Input[bool]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 timeouts: Optional[pulumi.Input['SnapshotTimeoutsArgs']] = None):
        """
        Input properties used for looking up and filtering Snapshot resources.
        :param pulumi.Input[str] created_at: The timestamp when this snapshot was created in RFC 3339.
        :param pulumi.Input[str] instance_id: The id of the instance to snapshot. - If the value of this attribute changes, the resource will be replaced.
        :param pulumi.Input[str] name: The human-readable name for the snapshot.
        :param pulumi.Input[str] region: The region identifier.
        :param pulumi.Input[bool] retain_on_delete: Flag to retain the snapshot when the resource is deleted. - Sets the default value "false" if the attribute is not set.
        :param pulumi.Input[int] size: The storage size of this snapshot given in GiB.
        :param pulumi.Input[str] status: The snapshot status.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if retain_on_delete is not None:
            pulumi.set(__self__, "retain_on_delete", retain_on_delete)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if timeouts is not None:
            pulumi.set(__self__, "timeouts", timeouts)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The timestamp when this snapshot was created in RFC 3339.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the instance to snapshot. - If the value of this attribute changes, the resource will be replaced.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The human-readable name for the snapshot.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region identifier.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="retainOnDelete")
    def retain_on_delete(self) -> Optional[pulumi.Input[bool]]:
        """
        Flag to retain the snapshot when the resource is deleted. - Sets the default value "false" if the attribute is not set.
        """
        return pulumi.get(self, "retain_on_delete")

    @retain_on_delete.setter
    def retain_on_delete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "retain_on_delete", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[int]]:
        """
        The storage size of this snapshot given in GiB.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The snapshot status.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def timeouts(self) -> Optional[pulumi.Input['SnapshotTimeoutsArgs']]:
        return pulumi.get(self, "timeouts")

    @timeouts.setter
    def timeouts(self, value: Optional[pulumi.Input['SnapshotTimeoutsArgs']]):
        pulumi.set(self, "timeouts", value)


class Snapshot(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retain_on_delete: Optional[pulumi.Input[bool]] = None,
                 timeouts: Optional[pulumi.Input[Union['SnapshotTimeoutsArgs', 'SnapshotTimeoutsArgsDict']]] = None,
                 __props__=None):
        """
        Snapshot resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_genesiscloud as genesiscloud

        target = genesiscloud.Instance("target")
        # ...
        example = genesiscloud.Snapshot("example",
            instance_id=target.id,
            retain_on_delete=True)
        # optional
        ```

        ## Import

        ```sh
        $ pulumi import genesiscloud:index/snapshot:Snapshot example 18efeec8-94f0-4776-8ff2-5e9b49c74608
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: The id of the instance to snapshot. - If the value of this attribute changes, the resource will be replaced.
        :param pulumi.Input[str] name: The human-readable name for the snapshot.
        :param pulumi.Input[bool] retain_on_delete: Flag to retain the snapshot when the resource is deleted. - Sets the default value "false" if the attribute is not set.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SnapshotArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Snapshot resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_genesiscloud as genesiscloud

        target = genesiscloud.Instance("target")
        # ...
        example = genesiscloud.Snapshot("example",
            instance_id=target.id,
            retain_on_delete=True)
        # optional
        ```

        ## Import

        ```sh
        $ pulumi import genesiscloud:index/snapshot:Snapshot example 18efeec8-94f0-4776-8ff2-5e9b49c74608
        ```

        :param str resource_name: The name of the resource.
        :param SnapshotArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SnapshotArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retain_on_delete: Optional[pulumi.Input[bool]] = None,
                 timeouts: Optional[pulumi.Input[Union['SnapshotTimeoutsArgs', 'SnapshotTimeoutsArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SnapshotArgs.__new__(SnapshotArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["name"] = name
            __props__.__dict__["retain_on_delete"] = retain_on_delete
            __props__.__dict__["timeouts"] = timeouts
            __props__.__dict__["created_at"] = None
            __props__.__dict__["region"] = None
            __props__.__dict__["size"] = None
            __props__.__dict__["status"] = None
        super(Snapshot, __self__).__init__(
            'genesiscloud:index/snapshot:Snapshot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            retain_on_delete: Optional[pulumi.Input[bool]] = None,
            size: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None,
            timeouts: Optional[pulumi.Input[Union['SnapshotTimeoutsArgs', 'SnapshotTimeoutsArgsDict']]] = None) -> 'Snapshot':
        """
        Get an existing Snapshot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The timestamp when this snapshot was created in RFC 3339.
        :param pulumi.Input[str] instance_id: The id of the instance to snapshot. - If the value of this attribute changes, the resource will be replaced.
        :param pulumi.Input[str] name: The human-readable name for the snapshot.
        :param pulumi.Input[str] region: The region identifier.
        :param pulumi.Input[bool] retain_on_delete: Flag to retain the snapshot when the resource is deleted. - Sets the default value "false" if the attribute is not set.
        :param pulumi.Input[int] size: The storage size of this snapshot given in GiB.
        :param pulumi.Input[str] status: The snapshot status.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SnapshotState.__new__(_SnapshotState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        __props__.__dict__["retain_on_delete"] = retain_on_delete
        __props__.__dict__["size"] = size
        __props__.__dict__["status"] = status
        __props__.__dict__["timeouts"] = timeouts
        return Snapshot(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The timestamp when this snapshot was created in RFC 3339.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The id of the instance to snapshot. - If the value of this attribute changes, the resource will be replaced.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The human-readable name for the snapshot.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region identifier.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="retainOnDelete")
    def retain_on_delete(self) -> pulumi.Output[bool]:
        """
        Flag to retain the snapshot when the resource is deleted. - Sets the default value "false" if the attribute is not set.
        """
        return pulumi.get(self, "retain_on_delete")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[int]:
        """
        The storage size of this snapshot given in GiB.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The snapshot status.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def timeouts(self) -> pulumi.Output[Optional['outputs.SnapshotTimeouts']]:
        return pulumi.get(self, "timeouts")

