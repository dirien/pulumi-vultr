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

__all__ = [
    'GetBlockStorageResult',
    'AwaitableGetBlockStorageResult',
    'get_block_storage',
    'get_block_storage_output',
]

@pulumi.output_type
class GetBlockStorageResult:
    """
    A collection of values returned by getBlockStorage.
    """
    def __init__(__self__, attached_to_instance=None, block_type=None, cost=None, date_created=None, filters=None, id=None, label=None, mount_id=None, region=None, size_gb=None, status=None):
        if attached_to_instance and not isinstance(attached_to_instance, str):
            raise TypeError("Expected argument 'attached_to_instance' to be a str")
        pulumi.set(__self__, "attached_to_instance", attached_to_instance)
        if block_type and not isinstance(block_type, str):
            raise TypeError("Expected argument 'block_type' to be a str")
        pulumi.set(__self__, "block_type", block_type)
        if cost and not isinstance(cost, int):
            raise TypeError("Expected argument 'cost' to be a int")
        pulumi.set(__self__, "cost", cost)
        if date_created and not isinstance(date_created, str):
            raise TypeError("Expected argument 'date_created' to be a str")
        pulumi.set(__self__, "date_created", date_created)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if label and not isinstance(label, str):
            raise TypeError("Expected argument 'label' to be a str")
        pulumi.set(__self__, "label", label)
        if mount_id and not isinstance(mount_id, str):
            raise TypeError("Expected argument 'mount_id' to be a str")
        pulumi.set(__self__, "mount_id", mount_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if size_gb and not isinstance(size_gb, int):
            raise TypeError("Expected argument 'size_gb' to be a int")
        pulumi.set(__self__, "size_gb", size_gb)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="attachedToInstance")
    def attached_to_instance(self) -> str:
        """
        The ID of the VPS the block storage subscription is attached to.
        """
        return pulumi.get(self, "attached_to_instance")

    @property
    @pulumi.getter(name="blockType")
    def block_type(self) -> str:
        """
        The type of block storage volume.
        """
        return pulumi.get(self, "block_type")

    @property
    @pulumi.getter
    def cost(self) -> int:
        """
        The cost per month of the block storage subscription in USD.
        """
        return pulumi.get(self, "cost")

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> str:
        """
        The date the block storage subscription was added to your Vultr account.
        """
        return pulumi.get(self, "date_created")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetBlockStorageFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def label(self) -> str:
        """
        The label of the block storage subscription.
        """
        return pulumi.get(self, "label")

    @property
    @pulumi.getter(name="mountId")
    def mount_id(self) -> str:
        """
        An ID associated with the instance, when mounted the ID can be found in /dev/disk/by-id prefixed with virtio.
        """
        return pulumi.get(self, "mount_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        The region ID of the block storage subscription.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="sizeGb")
    def size_gb(self) -> int:
        """
        The size of the block storage subscription in GB.
        """
        return pulumi.get(self, "size_gb")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the block storage subscription.
        """
        return pulumi.get(self, "status")


class AwaitableGetBlockStorageResult(GetBlockStorageResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBlockStorageResult(
            attached_to_instance=self.attached_to_instance,
            block_type=self.block_type,
            cost=self.cost,
            date_created=self.date_created,
            filters=self.filters,
            id=self.id,
            label=self.label,
            mount_id=self.mount_id,
            region=self.region,
            size_gb=self.size_gb,
            status=self.status)


def get_block_storage(filters: Optional[Sequence[Union['GetBlockStorageFilterArgs', 'GetBlockStorageFilterArgsDict']]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBlockStorageResult:
    """
    Get information about a Vultr block storage subscription.

    ## Example Usage

    Get the information for a block storage subscription by `label`:

    ```python
    import pulumi
    import pulumi_vultr as vultr

    my_block_storage = vultr.get_block_storage(filters=[{
        "name": "label",
        "values": ["my-block-storage-label"],
    }])
    ```


    :param Sequence[Union['GetBlockStorageFilterArgs', 'GetBlockStorageFilterArgsDict']] filters: Query parameters for finding block storage subscriptions.
    """
    __args__ = dict()
    __args__['filters'] = filters
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vultr:index/getBlockStorage:getBlockStorage', __args__, opts=opts, typ=GetBlockStorageResult).value

    return AwaitableGetBlockStorageResult(
        attached_to_instance=pulumi.get(__ret__, 'attached_to_instance'),
        block_type=pulumi.get(__ret__, 'block_type'),
        cost=pulumi.get(__ret__, 'cost'),
        date_created=pulumi.get(__ret__, 'date_created'),
        filters=pulumi.get(__ret__, 'filters'),
        id=pulumi.get(__ret__, 'id'),
        label=pulumi.get(__ret__, 'label'),
        mount_id=pulumi.get(__ret__, 'mount_id'),
        region=pulumi.get(__ret__, 'region'),
        size_gb=pulumi.get(__ret__, 'size_gb'),
        status=pulumi.get(__ret__, 'status'))
def get_block_storage_output(filters: Optional[pulumi.Input[Optional[Sequence[Union['GetBlockStorageFilterArgs', 'GetBlockStorageFilterArgsDict']]]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBlockStorageResult]:
    """
    Get information about a Vultr block storage subscription.

    ## Example Usage

    Get the information for a block storage subscription by `label`:

    ```python
    import pulumi
    import pulumi_vultr as vultr

    my_block_storage = vultr.get_block_storage(filters=[{
        "name": "label",
        "values": ["my-block-storage-label"],
    }])
    ```


    :param Sequence[Union['GetBlockStorageFilterArgs', 'GetBlockStorageFilterArgsDict']] filters: Query parameters for finding block storage subscriptions.
    """
    __args__ = dict()
    __args__['filters'] = filters
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vultr:index/getBlockStorage:getBlockStorage', __args__, opts=opts, typ=GetBlockStorageResult)
    return __ret__.apply(lambda __response__: GetBlockStorageResult(
        attached_to_instance=pulumi.get(__response__, 'attached_to_instance'),
        block_type=pulumi.get(__response__, 'block_type'),
        cost=pulumi.get(__response__, 'cost'),
        date_created=pulumi.get(__response__, 'date_created'),
        filters=pulumi.get(__response__, 'filters'),
        id=pulumi.get(__response__, 'id'),
        label=pulumi.get(__response__, 'label'),
        mount_id=pulumi.get(__response__, 'mount_id'),
        region=pulumi.get(__response__, 'region'),
        size_gb=pulumi.get(__response__, 'size_gb'),
        status=pulumi.get(__response__, 'status')))
