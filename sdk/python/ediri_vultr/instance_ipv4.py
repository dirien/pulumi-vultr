# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['InstanceIpv4Args', 'InstanceIpv4']

@pulumi.input_type
class InstanceIpv4Args:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 reboot: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a InstanceIpv4 resource.
        :param pulumi.Input[str] instance_id: The ID of the instance to be assigned the IPv4 address.
        :param pulumi.Input[bool] reboot: Default true. Determines whether or not the server is rebooted after adding the IPv4 address.
        """
        InstanceIpv4Args._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            instance_id=instance_id,
            reboot=reboot,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             instance_id: pulumi.Input[str],
             reboot: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("instance_id", instance_id)
        if reboot is not None:
            _setter("reboot", reboot)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The ID of the instance to be assigned the IPv4 address.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def reboot(self) -> Optional[pulumi.Input[bool]]:
        """
        Default true. Determines whether or not the server is rebooted after adding the IPv4 address.
        """
        return pulumi.get(self, "reboot")

    @reboot.setter
    def reboot(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "reboot", value)


@pulumi.input_type
class _InstanceIpv4State:
    def __init__(__self__, *,
                 gateway: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 netmask: Optional[pulumi.Input[str]] = None,
                 reboot: Optional[pulumi.Input[bool]] = None,
                 reverse: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InstanceIpv4 resources.
        :param pulumi.Input[str] gateway: The gateway IP address.
        :param pulumi.Input[str] instance_id: The ID of the instance to be assigned the IPv4 address.
        :param pulumi.Input[str] ip: The IPv4 address in canonical format.
        :param pulumi.Input[str] netmask: The IPv4 netmask in dot-decimal notation.
        :param pulumi.Input[bool] reboot: Default true. Determines whether or not the server is rebooted after adding the IPv4 address.
        :param pulumi.Input[str] reverse: The reverse DNS information for this IP address.
        """
        _InstanceIpv4State._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            gateway=gateway,
            instance_id=instance_id,
            ip=ip,
            netmask=netmask,
            reboot=reboot,
            reverse=reverse,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             gateway: Optional[pulumi.Input[str]] = None,
             instance_id: Optional[pulumi.Input[str]] = None,
             ip: Optional[pulumi.Input[str]] = None,
             netmask: Optional[pulumi.Input[str]] = None,
             reboot: Optional[pulumi.Input[bool]] = None,
             reverse: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if gateway is not None:
            _setter("gateway", gateway)
        if instance_id is not None:
            _setter("instance_id", instance_id)
        if ip is not None:
            _setter("ip", ip)
        if netmask is not None:
            _setter("netmask", netmask)
        if reboot is not None:
            _setter("reboot", reboot)
        if reverse is not None:
            _setter("reverse", reverse)

    @property
    @pulumi.getter
    def gateway(self) -> Optional[pulumi.Input[str]]:
        """
        The gateway IP address.
        """
        return pulumi.get(self, "gateway")

    @gateway.setter
    def gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the instance to be assigned the IPv4 address.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv4 address in canonical format.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def netmask(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv4 netmask in dot-decimal notation.
        """
        return pulumi.get(self, "netmask")

    @netmask.setter
    def netmask(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "netmask", value)

    @property
    @pulumi.getter
    def reboot(self) -> Optional[pulumi.Input[bool]]:
        """
        Default true. Determines whether or not the server is rebooted after adding the IPv4 address.
        """
        return pulumi.get(self, "reboot")

    @reboot.setter
    def reboot(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "reboot", value)

    @property
    @pulumi.getter
    def reverse(self) -> Optional[pulumi.Input[str]]:
        """
        The reverse DNS information for this IP address.
        """
        return pulumi.get(self, "reverse")

    @reverse.setter
    def reverse(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reverse", value)


class InstanceIpv4(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 reboot: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Provides a Vultr instance IPv4 resource. This can be used to create, read, and
        modify a IPv4 address. instance is rebooted by default.

        ## Example Usage

        Create a new IPv4 address for a instance:

        ```python
        import pulumi
        import ediri_vultr as vultr

        my_instance = vultr.Instance("myInstance",
            enable_ipv6=True,
            os_id=167,
            plan="vc2-1c-1gb",
            region="ewr")
        my_instance_ipv4 = vultr.InstanceIpv4("myInstanceIpv4",
            instance_id=my_instance.id,
            reboot=False)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: The ID of the instance to be assigned the IPv4 address.
        :param pulumi.Input[bool] reboot: Default true. Determines whether or not the server is rebooted after adding the IPv4 address.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceIpv4Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Vultr instance IPv4 resource. This can be used to create, read, and
        modify a IPv4 address. instance is rebooted by default.

        ## Example Usage

        Create a new IPv4 address for a instance:

        ```python
        import pulumi
        import ediri_vultr as vultr

        my_instance = vultr.Instance("myInstance",
            enable_ipv6=True,
            os_id=167,
            plan="vc2-1c-1gb",
            region="ewr")
        my_instance_ipv4 = vultr.InstanceIpv4("myInstanceIpv4",
            instance_id=my_instance.id,
            reboot=False)
        ```

        :param str resource_name: The name of the resource.
        :param InstanceIpv4Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceIpv4Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            InstanceIpv4Args._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 reboot: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceIpv4Args.__new__(InstanceIpv4Args)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["reboot"] = reboot
            __props__.__dict__["gateway"] = None
            __props__.__dict__["ip"] = None
            __props__.__dict__["netmask"] = None
            __props__.__dict__["reverse"] = None
        super(InstanceIpv4, __self__).__init__(
            'vultr:index/instanceIpv4:InstanceIpv4',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            gateway: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            ip: Optional[pulumi.Input[str]] = None,
            netmask: Optional[pulumi.Input[str]] = None,
            reboot: Optional[pulumi.Input[bool]] = None,
            reverse: Optional[pulumi.Input[str]] = None) -> 'InstanceIpv4':
        """
        Get an existing InstanceIpv4 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] gateway: The gateway IP address.
        :param pulumi.Input[str] instance_id: The ID of the instance to be assigned the IPv4 address.
        :param pulumi.Input[str] ip: The IPv4 address in canonical format.
        :param pulumi.Input[str] netmask: The IPv4 netmask in dot-decimal notation.
        :param pulumi.Input[bool] reboot: Default true. Determines whether or not the server is rebooted after adding the IPv4 address.
        :param pulumi.Input[str] reverse: The reverse DNS information for this IP address.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceIpv4State.__new__(_InstanceIpv4State)

        __props__.__dict__["gateway"] = gateway
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["ip"] = ip
        __props__.__dict__["netmask"] = netmask
        __props__.__dict__["reboot"] = reboot
        __props__.__dict__["reverse"] = reverse
        return InstanceIpv4(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def gateway(self) -> pulumi.Output[str]:
        """
        The gateway IP address.
        """
        return pulumi.get(self, "gateway")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The ID of the instance to be assigned the IPv4 address.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Output[str]:
        """
        The IPv4 address in canonical format.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def netmask(self) -> pulumi.Output[str]:
        """
        The IPv4 netmask in dot-decimal notation.
        """
        return pulumi.get(self, "netmask")

    @property
    @pulumi.getter
    def reboot(self) -> pulumi.Output[Optional[bool]]:
        """
        Default true. Determines whether or not the server is rebooted after adding the IPv4 address.
        """
        return pulumi.get(self, "reboot")

    @property
    @pulumi.getter
    def reverse(self) -> pulumi.Output[str]:
        """
        The reverse DNS information for this IP address.
        """
        return pulumi.get(self, "reverse")

