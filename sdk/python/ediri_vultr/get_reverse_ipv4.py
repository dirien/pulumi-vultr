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
    'GetReverseIpv4Result',
    'AwaitableGetReverseIpv4Result',
    'get_reverse_ipv4',
    'get_reverse_ipv4_output',
]

@pulumi.output_type
class GetReverseIpv4Result:
    """
    A collection of values returned by getReverseIpv4.
    """
    def __init__(__self__, filters=None, gateway=None, id=None, instance_id=None, ip=None, netmask=None, reverse=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if gateway and not isinstance(gateway, str):
            raise TypeError("Expected argument 'gateway' to be a str")
        pulumi.set(__self__, "gateway", gateway)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if ip and not isinstance(ip, str):
            raise TypeError("Expected argument 'ip' to be a str")
        pulumi.set(__self__, "ip", ip)
        if netmask and not isinstance(netmask, str):
            raise TypeError("Expected argument 'netmask' to be a str")
        pulumi.set(__self__, "netmask", netmask)
        if reverse and not isinstance(reverse, str):
            raise TypeError("Expected argument 'reverse' to be a str")
        pulumi.set(__self__, "reverse", reverse)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetReverseIpv4FilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def gateway(self) -> str:
        """
        The gateway IP address.
        """
        return pulumi.get(self, "gateway")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        The ID of the instance the IPv4 reverse DNS record was set for.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def ip(self) -> str:
        """
        The IPv4 address in canonical format used in the reverse DNS record.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def netmask(self) -> str:
        """
        The IPv4 netmask in dot-decimal notation.
        """
        return pulumi.get(self, "netmask")

    @property
    @pulumi.getter
    def reverse(self) -> str:
        """
        The hostname used in the IPv4 reverse DNS record.
        """
        return pulumi.get(self, "reverse")


class AwaitableGetReverseIpv4Result(GetReverseIpv4Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetReverseIpv4Result(
            filters=self.filters,
            gateway=self.gateway,
            id=self.id,
            instance_id=self.instance_id,
            ip=self.ip,
            netmask=self.netmask,
            reverse=self.reverse)


def get_reverse_ipv4(filters: Optional[Sequence[pulumi.InputType['GetReverseIpv4FilterArgs']]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetReverseIpv4Result:
    """
    Get information about a Vultr Reverse IPv4.

    ## Example Usage

    Get the information for an IPv4 reverse DNS record by `reverse`:

    ```python
    import pulumi
    import pulumi_vultr as vultr

    my_reverse_ipv4 = vultr.get_reverse_ipv4(filters=[vultr.GetReverseIpv4FilterArgs(
        name="reverse",
        values=["host.example.com"],
    )])
    ```


    :param Sequence[pulumi.InputType['GetReverseIpv4FilterArgs']] filters: Query parameters for finding IPv4 reverse DNS records.
    """
    __args__ = dict()
    __args__['filters'] = filters
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vultr:index/getReverseIpv4:getReverseIpv4', __args__, opts=opts, typ=GetReverseIpv4Result).value

    return AwaitableGetReverseIpv4Result(
        filters=pulumi.get(__ret__, 'filters'),
        gateway=pulumi.get(__ret__, 'gateway'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        ip=pulumi.get(__ret__, 'ip'),
        netmask=pulumi.get(__ret__, 'netmask'),
        reverse=pulumi.get(__ret__, 'reverse'))


@_utilities.lift_output_func(get_reverse_ipv4)
def get_reverse_ipv4_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetReverseIpv4FilterArgs']]]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetReverseIpv4Result]:
    """
    Get information about a Vultr Reverse IPv4.

    ## Example Usage

    Get the information for an IPv4 reverse DNS record by `reverse`:

    ```python
    import pulumi
    import pulumi_vultr as vultr

    my_reverse_ipv4 = vultr.get_reverse_ipv4(filters=[vultr.GetReverseIpv4FilterArgs(
        name="reverse",
        values=["host.example.com"],
    )])
    ```


    :param Sequence[pulumi.InputType['GetReverseIpv4FilterArgs']] filters: Query parameters for finding IPv4 reverse DNS records.
    """
    ...
