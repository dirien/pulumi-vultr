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
    'GetLoadBalancerResult',
    'AwaitableGetLoadBalancerResult',
    'get_load_balancer',
    'get_load_balancer_output',
]

@pulumi.output_type
class GetLoadBalancerResult:
    """
    A collection of values returned by getLoadBalancer.
    """
    def __init__(__self__, attached_instances=None, balancing_algorithm=None, cookie_name=None, date_created=None, filters=None, firewall_rules=None, forwarding_rules=None, has_ssl=None, health_check=None, id=None, ipv4=None, ipv6=None, label=None, private_network=None, proxy_protocol=None, region=None, ssl=None, ssl_redirect=None, status=None):
        if attached_instances and not isinstance(attached_instances, list):
            raise TypeError("Expected argument 'attached_instances' to be a list")
        pulumi.set(__self__, "attached_instances", attached_instances)
        if balancing_algorithm and not isinstance(balancing_algorithm, str):
            raise TypeError("Expected argument 'balancing_algorithm' to be a str")
        pulumi.set(__self__, "balancing_algorithm", balancing_algorithm)
        if cookie_name and not isinstance(cookie_name, str):
            raise TypeError("Expected argument 'cookie_name' to be a str")
        pulumi.set(__self__, "cookie_name", cookie_name)
        if date_created and not isinstance(date_created, str):
            raise TypeError("Expected argument 'date_created' to be a str")
        pulumi.set(__self__, "date_created", date_created)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if firewall_rules and not isinstance(firewall_rules, list):
            raise TypeError("Expected argument 'firewall_rules' to be a list")
        pulumi.set(__self__, "firewall_rules", firewall_rules)
        if forwarding_rules and not isinstance(forwarding_rules, list):
            raise TypeError("Expected argument 'forwarding_rules' to be a list")
        pulumi.set(__self__, "forwarding_rules", forwarding_rules)
        if has_ssl and not isinstance(has_ssl, bool):
            raise TypeError("Expected argument 'has_ssl' to be a bool")
        pulumi.set(__self__, "has_ssl", has_ssl)
        if health_check and not isinstance(health_check, dict):
            raise TypeError("Expected argument 'health_check' to be a dict")
        pulumi.set(__self__, "health_check", health_check)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ipv4 and not isinstance(ipv4, str):
            raise TypeError("Expected argument 'ipv4' to be a str")
        pulumi.set(__self__, "ipv4", ipv4)
        if ipv6 and not isinstance(ipv6, str):
            raise TypeError("Expected argument 'ipv6' to be a str")
        pulumi.set(__self__, "ipv6", ipv6)
        if label and not isinstance(label, str):
            raise TypeError("Expected argument 'label' to be a str")
        pulumi.set(__self__, "label", label)
        if private_network and not isinstance(private_network, str):
            raise TypeError("Expected argument 'private_network' to be a str")
        pulumi.set(__self__, "private_network", private_network)
        if proxy_protocol and not isinstance(proxy_protocol, bool):
            raise TypeError("Expected argument 'proxy_protocol' to be a bool")
        pulumi.set(__self__, "proxy_protocol", proxy_protocol)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if ssl and not isinstance(ssl, dict):
            raise TypeError("Expected argument 'ssl' to be a dict")
        pulumi.set(__self__, "ssl", ssl)
        if ssl_redirect and not isinstance(ssl_redirect, bool):
            raise TypeError("Expected argument 'ssl_redirect' to be a bool")
        pulumi.set(__self__, "ssl_redirect", ssl_redirect)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="attachedInstances")
    def attached_instances(self) -> Sequence[str]:
        """
        Array of instances that are currently attached to the load balancer.
        """
        return pulumi.get(self, "attached_instances")

    @property
    @pulumi.getter(name="balancingAlgorithm")
    def balancing_algorithm(self) -> str:
        """
        The balancing algorithm for your load balancer.
        """
        return pulumi.get(self, "balancing_algorithm")

    @property
    @pulumi.getter(name="cookieName")
    def cookie_name(self) -> str:
        """
        Name for your given sticky session.
        """
        return pulumi.get(self, "cookie_name")

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> str:
        return pulumi.get(self, "date_created")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetLoadBalancerFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter(name="firewallRules")
    def firewall_rules(self) -> Sequence[Mapping[str, Any]]:
        return pulumi.get(self, "firewall_rules")

    @property
    @pulumi.getter(name="forwardingRules")
    def forwarding_rules(self) -> Sequence[Mapping[str, Any]]:
        """
        Defines the forwarding rules for a load balancer. The configuration of a `forwarding_rules` is listened below.
        """
        return pulumi.get(self, "forwarding_rules")

    @property
    @pulumi.getter(name="hasSsl")
    def has_ssl(self) -> bool:
        """
        Boolean value that indicates if SSL is enabled.
        """
        return pulumi.get(self, "has_ssl")

    @property
    @pulumi.getter(name="healthCheck")
    def health_check(self) -> Mapping[str, Any]:
        """
        Defines the way load balancers should check for health. The configuration of a `health_check` is listed below.
        """
        return pulumi.get(self, "health_check")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ipv4(self) -> str:
        """
        IPv4 address for your load balancer.
        """
        return pulumi.get(self, "ipv4")

    @property
    @pulumi.getter
    def ipv6(self) -> str:
        """
        IPv6 address for your load balancer.
        """
        return pulumi.get(self, "ipv6")

    @property
    @pulumi.getter
    def label(self) -> str:
        """
        The load balancers label.
        """
        return pulumi.get(self, "label")

    @property
    @pulumi.getter(name="privateNetwork")
    def private_network(self) -> str:
        """
        (Deprecated: use `vpc` instead) Defines the private network the load balancer is attached to.
        """
        return pulumi.get(self, "private_network")

    @property
    @pulumi.getter(name="proxyProtocol")
    def proxy_protocol(self) -> Optional[bool]:
        """
        Boolean value that indicates if Proxy Protocol is enabled.
        """
        return pulumi.get(self, "proxy_protocol")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        The region your load balancer is deployed in.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def ssl(self) -> Mapping[str, Any]:
        return pulumi.get(self, "ssl")

    @property
    @pulumi.getter(name="sslRedirect")
    def ssl_redirect(self) -> bool:
        """
        Boolean value that indicates if HTTP calls will be redirected to HTTPS.
        """
        return pulumi.get(self, "ssl_redirect")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Current status for the load balancer
        """
        return pulumi.get(self, "status")


class AwaitableGetLoadBalancerResult(GetLoadBalancerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLoadBalancerResult(
            attached_instances=self.attached_instances,
            balancing_algorithm=self.balancing_algorithm,
            cookie_name=self.cookie_name,
            date_created=self.date_created,
            filters=self.filters,
            firewall_rules=self.firewall_rules,
            forwarding_rules=self.forwarding_rules,
            has_ssl=self.has_ssl,
            health_check=self.health_check,
            id=self.id,
            ipv4=self.ipv4,
            ipv6=self.ipv6,
            label=self.label,
            private_network=self.private_network,
            proxy_protocol=self.proxy_protocol,
            region=self.region,
            ssl=self.ssl,
            ssl_redirect=self.ssl_redirect,
            status=self.status)


def get_load_balancer(filters: Optional[Sequence[pulumi.InputType['GetLoadBalancerFilterArgs']]] = None,
                      proxy_protocol: Optional[bool] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLoadBalancerResult:
    """
    Get information about a Vultr load balancer.

    ## Example Usage

    Get the information for a load balancer by `label`:

    ```python
    import pulumi
    import pulumi_vultr as vultr

    my_lb = vultr.get_load_balancer(filters=[vultr.GetLoadBalancerFilterArgs(
        name="label",
        values=["my-lb-label"],
    )])
    ```


    :param Sequence[pulumi.InputType['GetLoadBalancerFilterArgs']] filters: Query parameters for finding load balancers.
    :param bool proxy_protocol: Boolean value that indicates if Proxy Protocol is enabled.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['proxyProtocol'] = proxy_protocol
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vultr:index/getLoadBalancer:getLoadBalancer', __args__, opts=opts, typ=GetLoadBalancerResult).value

    return AwaitableGetLoadBalancerResult(
        attached_instances=pulumi.get(__ret__, 'attached_instances'),
        balancing_algorithm=pulumi.get(__ret__, 'balancing_algorithm'),
        cookie_name=pulumi.get(__ret__, 'cookie_name'),
        date_created=pulumi.get(__ret__, 'date_created'),
        filters=pulumi.get(__ret__, 'filters'),
        firewall_rules=pulumi.get(__ret__, 'firewall_rules'),
        forwarding_rules=pulumi.get(__ret__, 'forwarding_rules'),
        has_ssl=pulumi.get(__ret__, 'has_ssl'),
        health_check=pulumi.get(__ret__, 'health_check'),
        id=pulumi.get(__ret__, 'id'),
        ipv4=pulumi.get(__ret__, 'ipv4'),
        ipv6=pulumi.get(__ret__, 'ipv6'),
        label=pulumi.get(__ret__, 'label'),
        private_network=pulumi.get(__ret__, 'private_network'),
        proxy_protocol=pulumi.get(__ret__, 'proxy_protocol'),
        region=pulumi.get(__ret__, 'region'),
        ssl=pulumi.get(__ret__, 'ssl'),
        ssl_redirect=pulumi.get(__ret__, 'ssl_redirect'),
        status=pulumi.get(__ret__, 'status'))


@_utilities.lift_output_func(get_load_balancer)
def get_load_balancer_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetLoadBalancerFilterArgs']]]]] = None,
                             proxy_protocol: Optional[pulumi.Input[Optional[bool]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLoadBalancerResult]:
    """
    Get information about a Vultr load balancer.

    ## Example Usage

    Get the information for a load balancer by `label`:

    ```python
    import pulumi
    import pulumi_vultr as vultr

    my_lb = vultr.get_load_balancer(filters=[vultr.GetLoadBalancerFilterArgs(
        name="label",
        values=["my-lb-label"],
    )])
    ```


    :param Sequence[pulumi.InputType['GetLoadBalancerFilterArgs']] filters: Query parameters for finding load balancers.
    :param bool proxy_protocol: Boolean value that indicates if Proxy Protocol is enabled.
    """
    ...
