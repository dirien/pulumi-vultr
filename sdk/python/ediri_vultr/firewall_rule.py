# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FirewallRuleArgs', 'FirewallRule']

@pulumi.input_type
class FirewallRuleArgs:
    def __init__(__self__, *,
                 firewall_group_id: pulumi.Input[str],
                 ip_type: pulumi.Input[str],
                 protocol: pulumi.Input[str],
                 subnet: pulumi.Input[str],
                 subnet_size: pulumi.Input[int],
                 notes: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallRule resource.
        :param pulumi.Input[str] firewall_group_id: The firewall group that the firewall rule will belong to.
        :param pulumi.Input[str] ip_type: The type of ip for this firewall rule. Possible values (v4, v6) **Note** they must be lowercase
        :param pulumi.Input[str] protocol: The type of protocol for this firewall rule. Possible values (icmp, tcp, udp, gre, esp, ah) **Note** they must be lowercase
        :param pulumi.Input[str] subnet: IP address that you want to define for this firewall rule.
        :param pulumi.Input[int] subnet_size: The number of bits for the subnet in CIDR notation. Example: 32.
        :param pulumi.Input[str] notes: A simple note for a given firewall rule
        :param pulumi.Input[str] port: TCP/UDP only. This field can be a specific port or a colon separated port range.
        :param pulumi.Input[str] source: Possible values ("", cloudflare)
        """
        pulumi.set(__self__, "firewall_group_id", firewall_group_id)
        pulumi.set(__self__, "ip_type", ip_type)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "subnet", subnet)
        pulumi.set(__self__, "subnet_size", subnet_size)
        if notes is not None:
            pulumi.set(__self__, "notes", notes)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if source is not None:
            pulumi.set(__self__, "source", source)

    @property
    @pulumi.getter(name="firewallGroupId")
    def firewall_group_id(self) -> pulumi.Input[str]:
        """
        The firewall group that the firewall rule will belong to.
        """
        return pulumi.get(self, "firewall_group_id")

    @firewall_group_id.setter
    def firewall_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "firewall_group_id", value)

    @property
    @pulumi.getter(name="ipType")
    def ip_type(self) -> pulumi.Input[str]:
        """
        The type of ip for this firewall rule. Possible values (v4, v6) **Note** they must be lowercase
        """
        return pulumi.get(self, "ip_type")

    @ip_type.setter
    def ip_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip_type", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input[str]:
        """
        The type of protocol for this firewall rule. Possible values (icmp, tcp, udp, gre, esp, ah) **Note** they must be lowercase
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def subnet(self) -> pulumi.Input[str]:
        """
        IP address that you want to define for this firewall rule.
        """
        return pulumi.get(self, "subnet")

    @subnet.setter
    def subnet(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet", value)

    @property
    @pulumi.getter(name="subnetSize")
    def subnet_size(self) -> pulumi.Input[int]:
        """
        The number of bits for the subnet in CIDR notation. Example: 32.
        """
        return pulumi.get(self, "subnet_size")

    @subnet_size.setter
    def subnet_size(self, value: pulumi.Input[int]):
        pulumi.set(self, "subnet_size", value)

    @property
    @pulumi.getter
    def notes(self) -> Optional[pulumi.Input[str]]:
        """
        A simple note for a given firewall rule
        """
        return pulumi.get(self, "notes")

    @notes.setter
    def notes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notes", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[str]]:
        """
        TCP/UDP only. This field can be a specific port or a colon separated port range.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        """
        Possible values ("", cloudflare)
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)


@pulumi.input_type
class _FirewallRuleState:
    def __init__(__self__, *,
                 firewall_group_id: Optional[pulumi.Input[str]] = None,
                 ip_type: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 subnet: Optional[pulumi.Input[str]] = None,
                 subnet_size: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering FirewallRule resources.
        :param pulumi.Input[str] firewall_group_id: The firewall group that the firewall rule will belong to.
        :param pulumi.Input[str] ip_type: The type of ip for this firewall rule. Possible values (v4, v6) **Note** they must be lowercase
        :param pulumi.Input[str] notes: A simple note for a given firewall rule
        :param pulumi.Input[str] port: TCP/UDP only. This field can be a specific port or a colon separated port range.
        :param pulumi.Input[str] protocol: The type of protocol for this firewall rule. Possible values (icmp, tcp, udp, gre, esp, ah) **Note** they must be lowercase
        :param pulumi.Input[str] source: Possible values ("", cloudflare)
        :param pulumi.Input[str] subnet: IP address that you want to define for this firewall rule.
        :param pulumi.Input[int] subnet_size: The number of bits for the subnet in CIDR notation. Example: 32.
        """
        if firewall_group_id is not None:
            pulumi.set(__self__, "firewall_group_id", firewall_group_id)
        if ip_type is not None:
            pulumi.set(__self__, "ip_type", ip_type)
        if notes is not None:
            pulumi.set(__self__, "notes", notes)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if subnet is not None:
            pulumi.set(__self__, "subnet", subnet)
        if subnet_size is not None:
            pulumi.set(__self__, "subnet_size", subnet_size)

    @property
    @pulumi.getter(name="firewallGroupId")
    def firewall_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The firewall group that the firewall rule will belong to.
        """
        return pulumi.get(self, "firewall_group_id")

    @firewall_group_id.setter
    def firewall_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "firewall_group_id", value)

    @property
    @pulumi.getter(name="ipType")
    def ip_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of ip for this firewall rule. Possible values (v4, v6) **Note** they must be lowercase
        """
        return pulumi.get(self, "ip_type")

    @ip_type.setter
    def ip_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_type", value)

    @property
    @pulumi.getter
    def notes(self) -> Optional[pulumi.Input[str]]:
        """
        A simple note for a given firewall rule
        """
        return pulumi.get(self, "notes")

    @notes.setter
    def notes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notes", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[str]]:
        """
        TCP/UDP only. This field can be a specific port or a colon separated port range.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        The type of protocol for this firewall rule. Possible values (icmp, tcp, udp, gre, esp, ah) **Note** they must be lowercase
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        """
        Possible values ("", cloudflare)
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def subnet(self) -> Optional[pulumi.Input[str]]:
        """
        IP address that you want to define for this firewall rule.
        """
        return pulumi.get(self, "subnet")

    @subnet.setter
    def subnet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet", value)

    @property
    @pulumi.getter(name="subnetSize")
    def subnet_size(self) -> Optional[pulumi.Input[int]]:
        """
        The number of bits for the subnet in CIDR notation. Example: 32.
        """
        return pulumi.get(self, "subnet_size")

    @subnet_size.setter
    def subnet_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "subnet_size", value)


class FirewallRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 firewall_group_id: Optional[pulumi.Input[str]] = None,
                 ip_type: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 subnet: Optional[pulumi.Input[str]] = None,
                 subnet_size: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a Vultr Firewall Rule resource. This can be used to create, read, modify, and delete Firewall rules.

        ## Example Usage

        Create a Firewall Rule

        ```python
        import pulumi
        import ediri_vultr as vultr

        my_firewallgroup = vultr.FirewallGroup("myFirewallgroup", description="base firewall")
        my_firewallrule = vultr.FirewallRule("myFirewallrule",
            firewall_group_id=my_firewallgroup.id,
            protocol="tcp",
            ip_type="v4",
            subnet="0.0.0.0",
            subnet_size=0,
            port="8090",
            notes="my firewall rule")
        ```

        ## Import

        Firewall Rules can be imported using the Firewall Group `ID` and Firewall Rule `ID`, e.g.

        ```sh
         $ pulumi import vultr:index/firewallRule:FirewallRule my_rule b6a859c5-b299-49dd-8888-b1abbc517d08,1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] firewall_group_id: The firewall group that the firewall rule will belong to.
        :param pulumi.Input[str] ip_type: The type of ip for this firewall rule. Possible values (v4, v6) **Note** they must be lowercase
        :param pulumi.Input[str] notes: A simple note for a given firewall rule
        :param pulumi.Input[str] port: TCP/UDP only. This field can be a specific port or a colon separated port range.
        :param pulumi.Input[str] protocol: The type of protocol for this firewall rule. Possible values (icmp, tcp, udp, gre, esp, ah) **Note** they must be lowercase
        :param pulumi.Input[str] source: Possible values ("", cloudflare)
        :param pulumi.Input[str] subnet: IP address that you want to define for this firewall rule.
        :param pulumi.Input[int] subnet_size: The number of bits for the subnet in CIDR notation. Example: 32.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Vultr Firewall Rule resource. This can be used to create, read, modify, and delete Firewall rules.

        ## Example Usage

        Create a Firewall Rule

        ```python
        import pulumi
        import ediri_vultr as vultr

        my_firewallgroup = vultr.FirewallGroup("myFirewallgroup", description="base firewall")
        my_firewallrule = vultr.FirewallRule("myFirewallrule",
            firewall_group_id=my_firewallgroup.id,
            protocol="tcp",
            ip_type="v4",
            subnet="0.0.0.0",
            subnet_size=0,
            port="8090",
            notes="my firewall rule")
        ```

        ## Import

        Firewall Rules can be imported using the Firewall Group `ID` and Firewall Rule `ID`, e.g.

        ```sh
         $ pulumi import vultr:index/firewallRule:FirewallRule my_rule b6a859c5-b299-49dd-8888-b1abbc517d08,1
        ```

        :param str resource_name: The name of the resource.
        :param FirewallRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 firewall_group_id: Optional[pulumi.Input[str]] = None,
                 ip_type: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 subnet: Optional[pulumi.Input[str]] = None,
                 subnet_size: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallRuleArgs.__new__(FirewallRuleArgs)

            if firewall_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'firewall_group_id'")
            __props__.__dict__["firewall_group_id"] = firewall_group_id
            if ip_type is None and not opts.urn:
                raise TypeError("Missing required property 'ip_type'")
            __props__.__dict__["ip_type"] = ip_type
            __props__.__dict__["notes"] = notes
            __props__.__dict__["port"] = port
            if protocol is None and not opts.urn:
                raise TypeError("Missing required property 'protocol'")
            __props__.__dict__["protocol"] = protocol
            __props__.__dict__["source"] = source
            if subnet is None and not opts.urn:
                raise TypeError("Missing required property 'subnet'")
            __props__.__dict__["subnet"] = subnet
            if subnet_size is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_size'")
            __props__.__dict__["subnet_size"] = subnet_size
        super(FirewallRule, __self__).__init__(
            'vultr:index/firewallRule:FirewallRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            firewall_group_id: Optional[pulumi.Input[str]] = None,
            ip_type: Optional[pulumi.Input[str]] = None,
            notes: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[str]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            source: Optional[pulumi.Input[str]] = None,
            subnet: Optional[pulumi.Input[str]] = None,
            subnet_size: Optional[pulumi.Input[int]] = None) -> 'FirewallRule':
        """
        Get an existing FirewallRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] firewall_group_id: The firewall group that the firewall rule will belong to.
        :param pulumi.Input[str] ip_type: The type of ip for this firewall rule. Possible values (v4, v6) **Note** they must be lowercase
        :param pulumi.Input[str] notes: A simple note for a given firewall rule
        :param pulumi.Input[str] port: TCP/UDP only. This field can be a specific port or a colon separated port range.
        :param pulumi.Input[str] protocol: The type of protocol for this firewall rule. Possible values (icmp, tcp, udp, gre, esp, ah) **Note** they must be lowercase
        :param pulumi.Input[str] source: Possible values ("", cloudflare)
        :param pulumi.Input[str] subnet: IP address that you want to define for this firewall rule.
        :param pulumi.Input[int] subnet_size: The number of bits for the subnet in CIDR notation. Example: 32.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallRuleState.__new__(_FirewallRuleState)

        __props__.__dict__["firewall_group_id"] = firewall_group_id
        __props__.__dict__["ip_type"] = ip_type
        __props__.__dict__["notes"] = notes
        __props__.__dict__["port"] = port
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["source"] = source
        __props__.__dict__["subnet"] = subnet
        __props__.__dict__["subnet_size"] = subnet_size
        return FirewallRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="firewallGroupId")
    def firewall_group_id(self) -> pulumi.Output[str]:
        """
        The firewall group that the firewall rule will belong to.
        """
        return pulumi.get(self, "firewall_group_id")

    @property
    @pulumi.getter(name="ipType")
    def ip_type(self) -> pulumi.Output[str]:
        """
        The type of ip for this firewall rule. Possible values (v4, v6) **Note** they must be lowercase
        """
        return pulumi.get(self, "ip_type")

    @property
    @pulumi.getter
    def notes(self) -> pulumi.Output[Optional[str]]:
        """
        A simple note for a given firewall rule
        """
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[Optional[str]]:
        """
        TCP/UDP only. This field can be a specific port or a colon separated port range.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        """
        The type of protocol for this firewall rule. Possible values (icmp, tcp, udp, gre, esp, ah) **Note** they must be lowercase
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[Optional[str]]:
        """
        Possible values ("", cloudflare)
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter
    def subnet(self) -> pulumi.Output[str]:
        """
        IP address that you want to define for this firewall rule.
        """
        return pulumi.get(self, "subnet")

    @property
    @pulumi.getter(name="subnetSize")
    def subnet_size(self) -> pulumi.Output[int]:
        """
        The number of bits for the subnet in CIDR notation. Example: 32.
        """
        return pulumi.get(self, "subnet_size")

