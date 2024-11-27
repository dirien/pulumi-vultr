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
    'GetPlanResult',
    'AwaitableGetPlanResult',
    'get_plan',
    'get_plan_output',
]

@pulumi.output_type
class GetPlanResult:
    """
    A collection of values returned by getPlan.
    """
    def __init__(__self__, bandwidth=None, disk=None, disk_count=None, filters=None, gpu_type=None, gpu_vram=None, id=None, locations=None, monthly_cost=None, ram=None, type=None, vcpu_count=None):
        if bandwidth and not isinstance(bandwidth, int):
            raise TypeError("Expected argument 'bandwidth' to be a int")
        pulumi.set(__self__, "bandwidth", bandwidth)
        if disk and not isinstance(disk, int):
            raise TypeError("Expected argument 'disk' to be a int")
        pulumi.set(__self__, "disk", disk)
        if disk_count and not isinstance(disk_count, int):
            raise TypeError("Expected argument 'disk_count' to be a int")
        pulumi.set(__self__, "disk_count", disk_count)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if gpu_type and not isinstance(gpu_type, str):
            raise TypeError("Expected argument 'gpu_type' to be a str")
        pulumi.set(__self__, "gpu_type", gpu_type)
        if gpu_vram and not isinstance(gpu_vram, int):
            raise TypeError("Expected argument 'gpu_vram' to be a int")
        pulumi.set(__self__, "gpu_vram", gpu_vram)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if locations and not isinstance(locations, list):
            raise TypeError("Expected argument 'locations' to be a list")
        pulumi.set(__self__, "locations", locations)
        if monthly_cost and not isinstance(monthly_cost, float):
            raise TypeError("Expected argument 'monthly_cost' to be a float")
        pulumi.set(__self__, "monthly_cost", monthly_cost)
        if ram and not isinstance(ram, int):
            raise TypeError("Expected argument 'ram' to be a int")
        pulumi.set(__self__, "ram", ram)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if vcpu_count and not isinstance(vcpu_count, int):
            raise TypeError("Expected argument 'vcpu_count' to be a int")
        pulumi.set(__self__, "vcpu_count", vcpu_count)

    @property
    @pulumi.getter
    def bandwidth(self) -> int:
        """
        The bandwidth available on the plan in GB.
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter
    def disk(self) -> int:
        """
        The amount of disk space in GB available on the plan.
        """
        return pulumi.get(self, "disk")

    @property
    @pulumi.getter(name="diskCount")
    def disk_count(self) -> int:
        """
        The number of disks that this plan offers.
        """
        return pulumi.get(self, "disk_count")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetPlanFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter(name="gpuType")
    def gpu_type(self) -> str:
        """
        For GPU plans, the GPU card type.
        """
        return pulumi.get(self, "gpu_type")

    @property
    @pulumi.getter(name="gpuVram")
    def gpu_vram(self) -> int:
        """
        For GPU plans, the VRAM available in the plan.
        """
        return pulumi.get(self, "gpu_vram")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def locations(self) -> Sequence[str]:
        return pulumi.get(self, "locations")

    @property
    @pulumi.getter(name="monthlyCost")
    def monthly_cost(self) -> float:
        """
        The price per month of the plan in USD.
        """
        return pulumi.get(self, "monthly_cost")

    @property
    @pulumi.getter
    def ram(self) -> int:
        """
        The amount of memory available on the plan in MB.
        """
        return pulumi.get(self, "ram")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of plan it is.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vcpuCount")
    def vcpu_count(self) -> int:
        """
        The number of virtual CPUs available on the plan.
        """
        return pulumi.get(self, "vcpu_count")


class AwaitableGetPlanResult(GetPlanResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPlanResult(
            bandwidth=self.bandwidth,
            disk=self.disk,
            disk_count=self.disk_count,
            filters=self.filters,
            gpu_type=self.gpu_type,
            gpu_vram=self.gpu_vram,
            id=self.id,
            locations=self.locations,
            monthly_cost=self.monthly_cost,
            ram=self.ram,
            type=self.type,
            vcpu_count=self.vcpu_count)


def get_plan(filters: Optional[Sequence[Union['GetPlanFilterArgs', 'GetPlanFilterArgsDict']]] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPlanResult:
    """
    Get information about a Vultr plan.

    ## Example Usage

    Get the information for a plan by `id`:

    ```python
    import pulumi
    import pulumi_vultr as vultr

    my_plan = vultr.get_plan(filters=[{
        "name": "id",
        "values": ["vc2-1c-2gb"],
    }])
    ```


    :param Sequence[Union['GetPlanFilterArgs', 'GetPlanFilterArgsDict']] filters: Query parameters for finding plans.
    """
    __args__ = dict()
    __args__['filters'] = filters
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vultr:index/getPlan:getPlan', __args__, opts=opts, typ=GetPlanResult).value

    return AwaitableGetPlanResult(
        bandwidth=pulumi.get(__ret__, 'bandwidth'),
        disk=pulumi.get(__ret__, 'disk'),
        disk_count=pulumi.get(__ret__, 'disk_count'),
        filters=pulumi.get(__ret__, 'filters'),
        gpu_type=pulumi.get(__ret__, 'gpu_type'),
        gpu_vram=pulumi.get(__ret__, 'gpu_vram'),
        id=pulumi.get(__ret__, 'id'),
        locations=pulumi.get(__ret__, 'locations'),
        monthly_cost=pulumi.get(__ret__, 'monthly_cost'),
        ram=pulumi.get(__ret__, 'ram'),
        type=pulumi.get(__ret__, 'type'),
        vcpu_count=pulumi.get(__ret__, 'vcpu_count'))
def get_plan_output(filters: Optional[pulumi.Input[Optional[Sequence[Union['GetPlanFilterArgs', 'GetPlanFilterArgsDict']]]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPlanResult]:
    """
    Get information about a Vultr plan.

    ## Example Usage

    Get the information for a plan by `id`:

    ```python
    import pulumi
    import pulumi_vultr as vultr

    my_plan = vultr.get_plan(filters=[{
        "name": "id",
        "values": ["vc2-1c-2gb"],
    }])
    ```


    :param Sequence[Union['GetPlanFilterArgs', 'GetPlanFilterArgsDict']] filters: Query parameters for finding plans.
    """
    __args__ = dict()
    __args__['filters'] = filters
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vultr:index/getPlan:getPlan', __args__, opts=opts, typ=GetPlanResult)
    return __ret__.apply(lambda __response__: GetPlanResult(
        bandwidth=pulumi.get(__response__, 'bandwidth'),
        disk=pulumi.get(__response__, 'disk'),
        disk_count=pulumi.get(__response__, 'disk_count'),
        filters=pulumi.get(__response__, 'filters'),
        gpu_type=pulumi.get(__response__, 'gpu_type'),
        gpu_vram=pulumi.get(__response__, 'gpu_vram'),
        id=pulumi.get(__response__, 'id'),
        locations=pulumi.get(__response__, 'locations'),
        monthly_cost=pulumi.get(__response__, 'monthly_cost'),
        ram=pulumi.get(__response__, 'ram'),
        type=pulumi.get(__response__, 'type'),
        vcpu_count=pulumi.get(__response__, 'vcpu_count')))
