# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SnapshotArgs', 'Snapshot']

@pulumi.input_type
class SnapshotArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Snapshot resource.
        :param pulumi.Input[str] instance_id: ID of a given instance that you want to create a snapshot from.
        :param pulumi.Input[str] description: The description for the given snapshot.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        ID of a given instance that you want to create a snapshot from.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description for the given snapshot.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class _SnapshotState:
    def __init__(__self__, *,
                 app_id: Optional[pulumi.Input[int]] = None,
                 date_created: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 os_id: Optional[pulumi.Input[int]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Snapshot resources.
        :param pulumi.Input[int] app_id: The app id which the snapshot is associated with.
        :param pulumi.Input[str] date_created: The date the snapshot was created.
        :param pulumi.Input[str] description: The description for the given snapshot.
        :param pulumi.Input[str] instance_id: ID of a given instance that you want to create a snapshot from.
        :param pulumi.Input[int] os_id: The os id which the snapshot is associated with.
        :param pulumi.Input[int] size: The size of the snapshot in Bytes.
        :param pulumi.Input[str] status: The status for the given snapshot.
        """
        if app_id is not None:
            pulumi.set(__self__, "app_id", app_id)
        if date_created is not None:
            pulumi.set(__self__, "date_created", date_created)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if os_id is not None:
            pulumi.set(__self__, "os_id", os_id)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> Optional[pulumi.Input[int]]:
        """
        The app id which the snapshot is associated with.
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> Optional[pulumi.Input[str]]:
        """
        The date the snapshot was created.
        """
        return pulumi.get(self, "date_created")

    @date_created.setter
    def date_created(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "date_created", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description for the given snapshot.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of a given instance that you want to create a snapshot from.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="osId")
    def os_id(self) -> Optional[pulumi.Input[int]]:
        """
        The os id which the snapshot is associated with.
        """
        return pulumi.get(self, "os_id")

    @os_id.setter
    def os_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "os_id", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[int]]:
        """
        The size of the snapshot in Bytes.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status for the given snapshot.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class Snapshot(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Vultr Snapshot resource. This can be used to create, read, modify, and delete Snapshot.

        ## Example Usage

        Create a new Snapshot

        ```python
        import pulumi
        import ediri_vultr as vultr

        my_instance = vultr.Instance("myInstance",
            label="my_instance",
            os_id=167,
            plan="201",
            region="ewr")
        my_snapshot = vultr.Snapshot("mySnapshot",
            description="my instances snapshot",
            instance_id=my_instance.id)
        ```

        ## Import

        Snapshots can be imported using the Snapshot `ID`, e.g.

        ```sh
        $ pulumi import vultr:index/snapshot:Snapshot my_snapshot 283941e8-0783-410e-9540-71c86b833992
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description for the given snapshot.
        :param pulumi.Input[str] instance_id: ID of a given instance that you want to create a snapshot from.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SnapshotArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Vultr Snapshot resource. This can be used to create, read, modify, and delete Snapshot.

        ## Example Usage

        Create a new Snapshot

        ```python
        import pulumi
        import ediri_vultr as vultr

        my_instance = vultr.Instance("myInstance",
            label="my_instance",
            os_id=167,
            plan="201",
            region="ewr")
        my_snapshot = vultr.Snapshot("mySnapshot",
            description="my instances snapshot",
            instance_id=my_instance.id)
        ```

        ## Import

        Snapshots can be imported using the Snapshot `ID`, e.g.

        ```sh
        $ pulumi import vultr:index/snapshot:Snapshot my_snapshot 283941e8-0783-410e-9540-71c86b833992
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
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SnapshotArgs.__new__(SnapshotArgs)

            __props__.__dict__["description"] = description
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["app_id"] = None
            __props__.__dict__["date_created"] = None
            __props__.__dict__["os_id"] = None
            __props__.__dict__["size"] = None
            __props__.__dict__["status"] = None
        super(Snapshot, __self__).__init__(
            'vultr:index/snapshot:Snapshot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_id: Optional[pulumi.Input[int]] = None,
            date_created: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            os_id: Optional[pulumi.Input[int]] = None,
            size: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'Snapshot':
        """
        Get an existing Snapshot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] app_id: The app id which the snapshot is associated with.
        :param pulumi.Input[str] date_created: The date the snapshot was created.
        :param pulumi.Input[str] description: The description for the given snapshot.
        :param pulumi.Input[str] instance_id: ID of a given instance that you want to create a snapshot from.
        :param pulumi.Input[int] os_id: The os id which the snapshot is associated with.
        :param pulumi.Input[int] size: The size of the snapshot in Bytes.
        :param pulumi.Input[str] status: The status for the given snapshot.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SnapshotState.__new__(_SnapshotState)

        __props__.__dict__["app_id"] = app_id
        __props__.__dict__["date_created"] = date_created
        __props__.__dict__["description"] = description
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["os_id"] = os_id
        __props__.__dict__["size"] = size
        __props__.__dict__["status"] = status
        return Snapshot(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Output[int]:
        """
        The app id which the snapshot is associated with.
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> pulumi.Output[str]:
        """
        The date the snapshot was created.
        """
        return pulumi.get(self, "date_created")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description for the given snapshot.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        ID of a given instance that you want to create a snapshot from.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="osId")
    def os_id(self) -> pulumi.Output[int]:
        """
        The os id which the snapshot is associated with.
        """
        return pulumi.get(self, "os_id")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[int]:
        """
        The size of the snapshot in Bytes.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status for the given snapshot.
        """
        return pulumi.get(self, "status")

