# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['BlockStorageArgs', 'BlockStorage']

@pulumi.input_type
class BlockStorageArgs:
    def __init__(__self__, *,
                 region: pulumi.Input[str],
                 size_gb: pulumi.Input[int],
                 attached_to_instance: Optional[pulumi.Input[str]] = None,
                 block_type: Optional[pulumi.Input[str]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 live: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a BlockStorage resource.
        :param pulumi.Input[str] region: Region in which this block storage will reside in. (Currently only NJ/NY supported region "ewr")
        :param pulumi.Input[int] size_gb: The size of the given block storage.
        :param pulumi.Input[str] attached_to_instance: VPS ID that you want to have this block storage attached to.
        :param pulumi.Input[str] block_type: Determines on the type of block storage volume that will be created. Soon to become a required parameter. Options are `high_perf` or `storage_opt`.
        :param pulumi.Input[str] label: Label that is given to your block storage.
        :param pulumi.Input[bool] live: Boolean value that will allow attachment of the volume to an instance without a restart. Default is false.
        """
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "size_gb", size_gb)
        if attached_to_instance is not None:
            pulumi.set(__self__, "attached_to_instance", attached_to_instance)
        if block_type is not None:
            pulumi.set(__self__, "block_type", block_type)
        if label is not None:
            pulumi.set(__self__, "label", label)
        if live is not None:
            pulumi.set(__self__, "live", live)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        Region in which this block storage will reside in. (Currently only NJ/NY supported region "ewr")
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="sizeGb")
    def size_gb(self) -> pulumi.Input[int]:
        """
        The size of the given block storage.
        """
        return pulumi.get(self, "size_gb")

    @size_gb.setter
    def size_gb(self, value: pulumi.Input[int]):
        pulumi.set(self, "size_gb", value)

    @property
    @pulumi.getter(name="attachedToInstance")
    def attached_to_instance(self) -> Optional[pulumi.Input[str]]:
        """
        VPS ID that you want to have this block storage attached to.
        """
        return pulumi.get(self, "attached_to_instance")

    @attached_to_instance.setter
    def attached_to_instance(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "attached_to_instance", value)

    @property
    @pulumi.getter(name="blockType")
    def block_type(self) -> Optional[pulumi.Input[str]]:
        """
        Determines on the type of block storage volume that will be created. Soon to become a required parameter. Options are `high_perf` or `storage_opt`.
        """
        return pulumi.get(self, "block_type")

    @block_type.setter
    def block_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "block_type", value)

    @property
    @pulumi.getter
    def label(self) -> Optional[pulumi.Input[str]]:
        """
        Label that is given to your block storage.
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter
    def live(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean value that will allow attachment of the volume to an instance without a restart. Default is false.
        """
        return pulumi.get(self, "live")

    @live.setter
    def live(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "live", value)


@pulumi.input_type
class _BlockStorageState:
    def __init__(__self__, *,
                 attached_to_instance: Optional[pulumi.Input[str]] = None,
                 block_type: Optional[pulumi.Input[str]] = None,
                 cost: Optional[pulumi.Input[float]] = None,
                 date_created: Optional[pulumi.Input[str]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 live: Optional[pulumi.Input[bool]] = None,
                 mount_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 size_gb: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BlockStorage resources.
        :param pulumi.Input[str] attached_to_instance: VPS ID that you want to have this block storage attached to.
        :param pulumi.Input[str] block_type: Determines on the type of block storage volume that will be created. Soon to become a required parameter. Options are `high_perf` or `storage_opt`.
        :param pulumi.Input[float] cost: The monthly cost of this block storage.
        :param pulumi.Input[str] date_created: The date this block storage was created.
        :param pulumi.Input[str] label: Label that is given to your block storage.
        :param pulumi.Input[bool] live: Boolean value that will allow attachment of the volume to an instance without a restart. Default is false.
        :param pulumi.Input[str] mount_id: An ID associated with the instance, when mounted the ID can be found in /dev/disk/by-id prefixed with virtio.
        :param pulumi.Input[str] region: Region in which this block storage will reside in. (Currently only NJ/NY supported region "ewr")
        :param pulumi.Input[int] size_gb: The size of the given block storage.
        :param pulumi.Input[str] status: Current status of your block storage.
        """
        if attached_to_instance is not None:
            pulumi.set(__self__, "attached_to_instance", attached_to_instance)
        if block_type is not None:
            pulumi.set(__self__, "block_type", block_type)
        if cost is not None:
            pulumi.set(__self__, "cost", cost)
        if date_created is not None:
            pulumi.set(__self__, "date_created", date_created)
        if label is not None:
            pulumi.set(__self__, "label", label)
        if live is not None:
            pulumi.set(__self__, "live", live)
        if mount_id is not None:
            pulumi.set(__self__, "mount_id", mount_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if size_gb is not None:
            pulumi.set(__self__, "size_gb", size_gb)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="attachedToInstance")
    def attached_to_instance(self) -> Optional[pulumi.Input[str]]:
        """
        VPS ID that you want to have this block storage attached to.
        """
        return pulumi.get(self, "attached_to_instance")

    @attached_to_instance.setter
    def attached_to_instance(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "attached_to_instance", value)

    @property
    @pulumi.getter(name="blockType")
    def block_type(self) -> Optional[pulumi.Input[str]]:
        """
        Determines on the type of block storage volume that will be created. Soon to become a required parameter. Options are `high_perf` or `storage_opt`.
        """
        return pulumi.get(self, "block_type")

    @block_type.setter
    def block_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "block_type", value)

    @property
    @pulumi.getter
    def cost(self) -> Optional[pulumi.Input[float]]:
        """
        The monthly cost of this block storage.
        """
        return pulumi.get(self, "cost")

    @cost.setter
    def cost(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "cost", value)

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> Optional[pulumi.Input[str]]:
        """
        The date this block storage was created.
        """
        return pulumi.get(self, "date_created")

    @date_created.setter
    def date_created(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "date_created", value)

    @property
    @pulumi.getter
    def label(self) -> Optional[pulumi.Input[str]]:
        """
        Label that is given to your block storage.
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter
    def live(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean value that will allow attachment of the volume to an instance without a restart. Default is false.
        """
        return pulumi.get(self, "live")

    @live.setter
    def live(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "live", value)

    @property
    @pulumi.getter(name="mountId")
    def mount_id(self) -> Optional[pulumi.Input[str]]:
        """
        An ID associated with the instance, when mounted the ID can be found in /dev/disk/by-id prefixed with virtio.
        """
        return pulumi.get(self, "mount_id")

    @mount_id.setter
    def mount_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mount_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Region in which this block storage will reside in. (Currently only NJ/NY supported region "ewr")
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="sizeGb")
    def size_gb(self) -> Optional[pulumi.Input[int]]:
        """
        The size of the given block storage.
        """
        return pulumi.get(self, "size_gb")

    @size_gb.setter
    def size_gb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size_gb", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Current status of your block storage.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class BlockStorage(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attached_to_instance: Optional[pulumi.Input[str]] = None,
                 block_type: Optional[pulumi.Input[str]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 live: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 size_gb: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a Vultr Block Storage resource. This can be used to create, read, modify, and delete Block Storage.

        ## Example Usage

        Create a new Block Storage

        ```python
        import pulumi
        import dirien_vultr as vultr

        my_blockstorage = vultr.BlockStorage("myBlockstorage",
            region="ewr",
            size_gb=10)
        ```

        ## Import

        Block Storage can be imported using the Block Storage `ID`, e.g.

        ```sh
         $ pulumi import vultr:index/blockStorage:BlockStorage my_blockstorage e315835e-d466-4e89-9b4c-dfd8788d7685
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attached_to_instance: VPS ID that you want to have this block storage attached to.
        :param pulumi.Input[str] block_type: Determines on the type of block storage volume that will be created. Soon to become a required parameter. Options are `high_perf` or `storage_opt`.
        :param pulumi.Input[str] label: Label that is given to your block storage.
        :param pulumi.Input[bool] live: Boolean value that will allow attachment of the volume to an instance without a restart. Default is false.
        :param pulumi.Input[str] region: Region in which this block storage will reside in. (Currently only NJ/NY supported region "ewr")
        :param pulumi.Input[int] size_gb: The size of the given block storage.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BlockStorageArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Vultr Block Storage resource. This can be used to create, read, modify, and delete Block Storage.

        ## Example Usage

        Create a new Block Storage

        ```python
        import pulumi
        import dirien_vultr as vultr

        my_blockstorage = vultr.BlockStorage("myBlockstorage",
            region="ewr",
            size_gb=10)
        ```

        ## Import

        Block Storage can be imported using the Block Storage `ID`, e.g.

        ```sh
         $ pulumi import vultr:index/blockStorage:BlockStorage my_blockstorage e315835e-d466-4e89-9b4c-dfd8788d7685
        ```

        :param str resource_name: The name of the resource.
        :param BlockStorageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BlockStorageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attached_to_instance: Optional[pulumi.Input[str]] = None,
                 block_type: Optional[pulumi.Input[str]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 live: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 size_gb: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BlockStorageArgs.__new__(BlockStorageArgs)

            __props__.__dict__["attached_to_instance"] = attached_to_instance
            __props__.__dict__["block_type"] = block_type
            __props__.__dict__["label"] = label
            __props__.__dict__["live"] = live
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            if size_gb is None and not opts.urn:
                raise TypeError("Missing required property 'size_gb'")
            __props__.__dict__["size_gb"] = size_gb
            __props__.__dict__["cost"] = None
            __props__.__dict__["date_created"] = None
            __props__.__dict__["mount_id"] = None
            __props__.__dict__["status"] = None
        super(BlockStorage, __self__).__init__(
            'vultr:index/blockStorage:BlockStorage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attached_to_instance: Optional[pulumi.Input[str]] = None,
            block_type: Optional[pulumi.Input[str]] = None,
            cost: Optional[pulumi.Input[float]] = None,
            date_created: Optional[pulumi.Input[str]] = None,
            label: Optional[pulumi.Input[str]] = None,
            live: Optional[pulumi.Input[bool]] = None,
            mount_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            size_gb: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'BlockStorage':
        """
        Get an existing BlockStorage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attached_to_instance: VPS ID that you want to have this block storage attached to.
        :param pulumi.Input[str] block_type: Determines on the type of block storage volume that will be created. Soon to become a required parameter. Options are `high_perf` or `storage_opt`.
        :param pulumi.Input[float] cost: The monthly cost of this block storage.
        :param pulumi.Input[str] date_created: The date this block storage was created.
        :param pulumi.Input[str] label: Label that is given to your block storage.
        :param pulumi.Input[bool] live: Boolean value that will allow attachment of the volume to an instance without a restart. Default is false.
        :param pulumi.Input[str] mount_id: An ID associated with the instance, when mounted the ID can be found in /dev/disk/by-id prefixed with virtio.
        :param pulumi.Input[str] region: Region in which this block storage will reside in. (Currently only NJ/NY supported region "ewr")
        :param pulumi.Input[int] size_gb: The size of the given block storage.
        :param pulumi.Input[str] status: Current status of your block storage.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BlockStorageState.__new__(_BlockStorageState)

        __props__.__dict__["attached_to_instance"] = attached_to_instance
        __props__.__dict__["block_type"] = block_type
        __props__.__dict__["cost"] = cost
        __props__.__dict__["date_created"] = date_created
        __props__.__dict__["label"] = label
        __props__.__dict__["live"] = live
        __props__.__dict__["mount_id"] = mount_id
        __props__.__dict__["region"] = region
        __props__.__dict__["size_gb"] = size_gb
        __props__.__dict__["status"] = status
        return BlockStorage(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="attachedToInstance")
    def attached_to_instance(self) -> pulumi.Output[Optional[str]]:
        """
        VPS ID that you want to have this block storage attached to.
        """
        return pulumi.get(self, "attached_to_instance")

    @property
    @pulumi.getter(name="blockType")
    def block_type(self) -> pulumi.Output[str]:
        """
        Determines on the type of block storage volume that will be created. Soon to become a required parameter. Options are `high_perf` or `storage_opt`.
        """
        return pulumi.get(self, "block_type")

    @property
    @pulumi.getter
    def cost(self) -> pulumi.Output[float]:
        """
        The monthly cost of this block storage.
        """
        return pulumi.get(self, "cost")

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> pulumi.Output[str]:
        """
        The date this block storage was created.
        """
        return pulumi.get(self, "date_created")

    @property
    @pulumi.getter
    def label(self) -> pulumi.Output[Optional[str]]:
        """
        Label that is given to your block storage.
        """
        return pulumi.get(self, "label")

    @property
    @pulumi.getter
    def live(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean value that will allow attachment of the volume to an instance without a restart. Default is false.
        """
        return pulumi.get(self, "live")

    @property
    @pulumi.getter(name="mountId")
    def mount_id(self) -> pulumi.Output[str]:
        """
        An ID associated with the instance, when mounted the ID can be found in /dev/disk/by-id prefixed with virtio.
        """
        return pulumi.get(self, "mount_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Region in which this block storage will reside in. (Currently only NJ/NY supported region "ewr")
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="sizeGb")
    def size_gb(self) -> pulumi.Output[int]:
        """
        The size of the given block storage.
        """
        return pulumi.get(self, "size_gb")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Current status of your block storage.
        """
        return pulumi.get(self, "status")

