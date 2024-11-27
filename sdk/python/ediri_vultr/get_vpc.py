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
    'GetVpcResult',
    'AwaitableGetVpcResult',
    'get_vpc',
    'get_vpc_output',
]

@pulumi.output_type
class GetVpcResult:
    """
    A collection of values returned by getVpc.
    """
    def __init__(__self__, date_created=None, description=None, filters=None, id=None, region=None, v4_subnet=None, v4_subnet_mask=None):
        if date_created and not isinstance(date_created, str):
            raise TypeError("Expected argument 'date_created' to be a str")
        pulumi.set(__self__, "date_created", date_created)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if v4_subnet and not isinstance(v4_subnet, str):
            raise TypeError("Expected argument 'v4_subnet' to be a str")
        pulumi.set(__self__, "v4_subnet", v4_subnet)
        if v4_subnet_mask and not isinstance(v4_subnet_mask, int):
            raise TypeError("Expected argument 'v4_subnet_mask' to be a int")
        pulumi.set(__self__, "v4_subnet_mask", v4_subnet_mask)

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> str:
        """
        The date the VPC was added to your Vultr account.
        """
        return pulumi.get(self, "date_created")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The VPC's description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetVpcFilterResult']]:
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
    def region(self) -> str:
        """
        The ID of the region that the VPC is in.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="v4Subnet")
    def v4_subnet(self) -> str:
        """
        The IPv4 network address. For example: 10.1.1.0.
        """
        return pulumi.get(self, "v4_subnet")

    @property
    @pulumi.getter(name="v4SubnetMask")
    def v4_subnet_mask(self) -> int:
        """
        The number of bits for the netmask in CIDR notation. Example: 20
        """
        return pulumi.get(self, "v4_subnet_mask")


class AwaitableGetVpcResult(GetVpcResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcResult(
            date_created=self.date_created,
            description=self.description,
            filters=self.filters,
            id=self.id,
            region=self.region,
            v4_subnet=self.v4_subnet,
            v4_subnet_mask=self.v4_subnet_mask)


def get_vpc(filters: Optional[Sequence[Union['GetVpcFilterArgs', 'GetVpcFilterArgsDict']]] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcResult:
    """
    Get information about a Vultr VPC.

    ## Example Usage

    Get the information for a VPC by `description`:

    ```python
    import pulumi
    import pulumi_vultr as vultr

    my_vpc = vultr.get_vpc(filters=[{
        "name": "description",
        "values": ["my-vpc-description"],
    }])
    ```


    :param Sequence[Union['GetVpcFilterArgs', 'GetVpcFilterArgsDict']] filters: Query parameters for finding VPCs.
    """
    __args__ = dict()
    __args__['filters'] = filters
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vultr:index/getVpc:getVpc', __args__, opts=opts, typ=GetVpcResult).value

    return AwaitableGetVpcResult(
        date_created=pulumi.get(__ret__, 'date_created'),
        description=pulumi.get(__ret__, 'description'),
        filters=pulumi.get(__ret__, 'filters'),
        id=pulumi.get(__ret__, 'id'),
        region=pulumi.get(__ret__, 'region'),
        v4_subnet=pulumi.get(__ret__, 'v4_subnet'),
        v4_subnet_mask=pulumi.get(__ret__, 'v4_subnet_mask'))
def get_vpc_output(filters: Optional[pulumi.Input[Optional[Sequence[Union['GetVpcFilterArgs', 'GetVpcFilterArgsDict']]]]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVpcResult]:
    """
    Get information about a Vultr VPC.

    ## Example Usage

    Get the information for a VPC by `description`:

    ```python
    import pulumi
    import pulumi_vultr as vultr

    my_vpc = vultr.get_vpc(filters=[{
        "name": "description",
        "values": ["my-vpc-description"],
    }])
    ```


    :param Sequence[Union['GetVpcFilterArgs', 'GetVpcFilterArgsDict']] filters: Query parameters for finding VPCs.
    """
    __args__ = dict()
    __args__['filters'] = filters
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vultr:index/getVpc:getVpc', __args__, opts=opts, typ=GetVpcResult)
    return __ret__.apply(lambda __response__: GetVpcResult(
        date_created=pulumi.get(__response__, 'date_created'),
        description=pulumi.get(__response__, 'description'),
        filters=pulumi.get(__response__, 'filters'),
        id=pulumi.get(__response__, 'id'),
        region=pulumi.get(__response__, 'region'),
        v4_subnet=pulumi.get(__response__, 'v4_subnet'),
        v4_subnet_mask=pulumi.get(__response__, 'v4_subnet_mask')))
