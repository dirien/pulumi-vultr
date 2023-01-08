# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ReservedIpArgs', 'ReservedIp']

@pulumi.input_type
class ReservedIpArgs:
    def __init__(__self__, *,
                 ip_type: pulumi.Input[str],
                 region: pulumi.Input[str],
                 instance_id: Optional[pulumi.Input[str]] = None,
                 label: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ReservedIp resource.
        :param pulumi.Input[str] ip_type: The type of reserved IP that you want. Either "v4" or "v6".
        :param pulumi.Input[str] region: The region ID that you want the reserved IP to be created in.
        :param pulumi.Input[str] instance_id: The VPS ID you want this reserved IP to be attached to.
        :param pulumi.Input[str] label: The label you want to give your reserved IP.
        """
        pulumi.set(__self__, "ip_type", ip_type)
        pulumi.set(__self__, "region", region)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if label is not None:
            pulumi.set(__self__, "label", label)

    @property
    @pulumi.getter(name="ipType")
    def ip_type(self) -> pulumi.Input[str]:
        """
        The type of reserved IP that you want. Either "v4" or "v6".
        """
        return pulumi.get(self, "ip_type")

    @ip_type.setter
    def ip_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip_type", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        The region ID that you want the reserved IP to be created in.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The VPS ID you want this reserved IP to be attached to.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def label(self) -> Optional[pulumi.Input[str]]:
        """
        The label you want to give your reserved IP.
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "label", value)


@pulumi.input_type
class _ReservedIpState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 ip_type: Optional[pulumi.Input[str]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 subnet: Optional[pulumi.Input[str]] = None,
                 subnet_size: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering ReservedIp resources.
        :param pulumi.Input[str] instance_id: The VPS ID you want this reserved IP to be attached to.
        :param pulumi.Input[str] ip_type: The type of reserved IP that you want. Either "v4" or "v6".
        :param pulumi.Input[str] label: The label you want to give your reserved IP.
        :param pulumi.Input[str] region: The region ID that you want the reserved IP to be created in.
        :param pulumi.Input[str] subnet: The reserved IP's subnet.
        :param pulumi.Input[int] subnet_size: The reserved IP's subnet size.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if ip_type is not None:
            pulumi.set(__self__, "ip_type", ip_type)
        if label is not None:
            pulumi.set(__self__, "label", label)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if subnet is not None:
            pulumi.set(__self__, "subnet", subnet)
        if subnet_size is not None:
            pulumi.set(__self__, "subnet_size", subnet_size)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The VPS ID you want this reserved IP to be attached to.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="ipType")
    def ip_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of reserved IP that you want. Either "v4" or "v6".
        """
        return pulumi.get(self, "ip_type")

    @ip_type.setter
    def ip_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_type", value)

    @property
    @pulumi.getter
    def label(self) -> Optional[pulumi.Input[str]]:
        """
        The label you want to give your reserved IP.
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region ID that you want the reserved IP to be created in.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def subnet(self) -> Optional[pulumi.Input[str]]:
        """
        The reserved IP's subnet.
        """
        return pulumi.get(self, "subnet")

    @subnet.setter
    def subnet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet", value)

    @property
    @pulumi.getter(name="subnetSize")
    def subnet_size(self) -> Optional[pulumi.Input[int]]:
        """
        The reserved IP's subnet size.
        """
        return pulumi.get(self, "subnet_size")

    @subnet_size.setter
    def subnet_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "subnet_size", value)


class ReservedIp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 ip_type: Optional[pulumi.Input[str]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Vultr reserved IP resource. This can be used to create, read, modify, and delete reserved IP addresses on your Vultr account.

        ## Example Usage

        Create a new reserved IP:

        ```python
        import pulumi
        import ediri_vultr as vultr

        my_reserved_ip = vultr.ReservedIp("myReservedIp",
            ip_type="v4",
            label="my-reserved-ip",
            region="sea")
        ```

        Attach a reserved IP to a instance:

        ```python
        import pulumi
        import ediri_vultr as vultr

        my_reserved_ip = vultr.ReservedIp("myReservedIp",
            instance_id="b9cc6fad-70b1-40ee-ab6a-4d622858962f",
            ip_type="v4",
            label="my-reserved-ip",
            region="sea")
        ```

        ## Import

        Reserved IPs can be imported using the reserved IP `ID`, e.g.

        ```sh
         $ pulumi import vultr:index/reservedIp:ReservedIp my_reserved_ip b9cc6fad-70b1-40ee-ab6a-4d622858962f
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: The VPS ID you want this reserved IP to be attached to.
        :param pulumi.Input[str] ip_type: The type of reserved IP that you want. Either "v4" or "v6".
        :param pulumi.Input[str] label: The label you want to give your reserved IP.
        :param pulumi.Input[str] region: The region ID that you want the reserved IP to be created in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReservedIpArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Vultr reserved IP resource. This can be used to create, read, modify, and delete reserved IP addresses on your Vultr account.

        ## Example Usage

        Create a new reserved IP:

        ```python
        import pulumi
        import ediri_vultr as vultr

        my_reserved_ip = vultr.ReservedIp("myReservedIp",
            ip_type="v4",
            label="my-reserved-ip",
            region="sea")
        ```

        Attach a reserved IP to a instance:

        ```python
        import pulumi
        import ediri_vultr as vultr

        my_reserved_ip = vultr.ReservedIp("myReservedIp",
            instance_id="b9cc6fad-70b1-40ee-ab6a-4d622858962f",
            ip_type="v4",
            label="my-reserved-ip",
            region="sea")
        ```

        ## Import

        Reserved IPs can be imported using the reserved IP `ID`, e.g.

        ```sh
         $ pulumi import vultr:index/reservedIp:ReservedIp my_reserved_ip b9cc6fad-70b1-40ee-ab6a-4d622858962f
        ```

        :param str resource_name: The name of the resource.
        :param ReservedIpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReservedIpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 ip_type: Optional[pulumi.Input[str]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ReservedIpArgs.__new__(ReservedIpArgs)

            __props__.__dict__["instance_id"] = instance_id
            if ip_type is None and not opts.urn:
                raise TypeError("Missing required property 'ip_type'")
            __props__.__dict__["ip_type"] = ip_type
            __props__.__dict__["label"] = label
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            __props__.__dict__["subnet"] = None
            __props__.__dict__["subnet_size"] = None
        super(ReservedIp, __self__).__init__(
            'vultr:index/reservedIp:ReservedIp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            ip_type: Optional[pulumi.Input[str]] = None,
            label: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            subnet: Optional[pulumi.Input[str]] = None,
            subnet_size: Optional[pulumi.Input[int]] = None) -> 'ReservedIp':
        """
        Get an existing ReservedIp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: The VPS ID you want this reserved IP to be attached to.
        :param pulumi.Input[str] ip_type: The type of reserved IP that you want. Either "v4" or "v6".
        :param pulumi.Input[str] label: The label you want to give your reserved IP.
        :param pulumi.Input[str] region: The region ID that you want the reserved IP to be created in.
        :param pulumi.Input[str] subnet: The reserved IP's subnet.
        :param pulumi.Input[int] subnet_size: The reserved IP's subnet size.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ReservedIpState.__new__(_ReservedIpState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["ip_type"] = ip_type
        __props__.__dict__["label"] = label
        __props__.__dict__["region"] = region
        __props__.__dict__["subnet"] = subnet
        __props__.__dict__["subnet_size"] = subnet_size
        return ReservedIp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The VPS ID you want this reserved IP to be attached to.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="ipType")
    def ip_type(self) -> pulumi.Output[str]:
        """
        The type of reserved IP that you want. Either "v4" or "v6".
        """
        return pulumi.get(self, "ip_type")

    @property
    @pulumi.getter
    def label(self) -> pulumi.Output[Optional[str]]:
        """
        The label you want to give your reserved IP.
        """
        return pulumi.get(self, "label")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region ID that you want the reserved IP to be created in.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def subnet(self) -> pulumi.Output[str]:
        """
        The reserved IP's subnet.
        """
        return pulumi.get(self, "subnet")

    @property
    @pulumi.getter(name="subnetSize")
    def subnet_size(self) -> pulumi.Output[int]:
        """
        The reserved IP's subnet size.
        """
        return pulumi.get(self, "subnet_size")
