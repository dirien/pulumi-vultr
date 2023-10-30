# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['Vpc2Args', 'Vpc2']

@pulumi.input_type
class Vpc2Args:
    def __init__(__self__, *,
                 region: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 ip_block: Optional[pulumi.Input[str]] = None,
                 ip_type: Optional[pulumi.Input[str]] = None,
                 prefix_length: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Vpc2 resource.
        :param pulumi.Input[str] region: The region ID that you want the VPC 2.0 to be created in.
        :param pulumi.Input[str] description: The description you want to give your VPC 2.0.
        :param pulumi.Input[str] ip_block: The IPv4 subnet to be used when attaching instances to this VPC 2.0.
        :param pulumi.Input[str] ip_type: Accepted values: `v4`.
        :param pulumi.Input[int] prefix_length: The number of bits for the netmask in CIDR notation. Example: 32
        """
        pulumi.set(__self__, "region", region)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ip_block is not None:
            pulumi.set(__self__, "ip_block", ip_block)
        if ip_type is not None:
            pulumi.set(__self__, "ip_type", ip_type)
        if prefix_length is not None:
            pulumi.set(__self__, "prefix_length", prefix_length)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        The region ID that you want the VPC 2.0 to be created in.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description you want to give your VPC 2.0.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="ipBlock")
    def ip_block(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv4 subnet to be used when attaching instances to this VPC 2.0.
        """
        return pulumi.get(self, "ip_block")

    @ip_block.setter
    def ip_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_block", value)

    @property
    @pulumi.getter(name="ipType")
    def ip_type(self) -> Optional[pulumi.Input[str]]:
        """
        Accepted values: `v4`.
        """
        return pulumi.get(self, "ip_type")

    @ip_type.setter
    def ip_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_type", value)

    @property
    @pulumi.getter(name="prefixLength")
    def prefix_length(self) -> Optional[pulumi.Input[int]]:
        """
        The number of bits for the netmask in CIDR notation. Example: 32
        """
        return pulumi.get(self, "prefix_length")

    @prefix_length.setter
    def prefix_length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "prefix_length", value)


@pulumi.input_type
class _Vpc2State:
    def __init__(__self__, *,
                 date_created: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_block: Optional[pulumi.Input[str]] = None,
                 ip_type: Optional[pulumi.Input[str]] = None,
                 prefix_length: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Vpc2 resources.
        :param pulumi.Input[str] date_created: The date that the VPC 2.0 was added to your Vultr account.
        :param pulumi.Input[str] description: The description you want to give your VPC 2.0.
        :param pulumi.Input[str] ip_block: The IPv4 subnet to be used when attaching instances to this VPC 2.0.
        :param pulumi.Input[str] ip_type: Accepted values: `v4`.
        :param pulumi.Input[int] prefix_length: The number of bits for the netmask in CIDR notation. Example: 32
        :param pulumi.Input[str] region: The region ID that you want the VPC 2.0 to be created in.
        """
        if date_created is not None:
            pulumi.set(__self__, "date_created", date_created)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ip_block is not None:
            pulumi.set(__self__, "ip_block", ip_block)
        if ip_type is not None:
            pulumi.set(__self__, "ip_type", ip_type)
        if prefix_length is not None:
            pulumi.set(__self__, "prefix_length", prefix_length)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> Optional[pulumi.Input[str]]:
        """
        The date that the VPC 2.0 was added to your Vultr account.
        """
        return pulumi.get(self, "date_created")

    @date_created.setter
    def date_created(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "date_created", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description you want to give your VPC 2.0.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="ipBlock")
    def ip_block(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv4 subnet to be used when attaching instances to this VPC 2.0.
        """
        return pulumi.get(self, "ip_block")

    @ip_block.setter
    def ip_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_block", value)

    @property
    @pulumi.getter(name="ipType")
    def ip_type(self) -> Optional[pulumi.Input[str]]:
        """
        Accepted values: `v4`.
        """
        return pulumi.get(self, "ip_type")

    @ip_type.setter
    def ip_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_type", value)

    @property
    @pulumi.getter(name="prefixLength")
    def prefix_length(self) -> Optional[pulumi.Input[int]]:
        """
        The number of bits for the netmask in CIDR notation. Example: 32
        """
        return pulumi.get(self, "prefix_length")

    @prefix_length.setter
    def prefix_length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "prefix_length", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region ID that you want the VPC 2.0 to be created in.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class Vpc2(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_block: Optional[pulumi.Input[str]] = None,
                 ip_type: Optional[pulumi.Input[str]] = None,
                 prefix_length: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Vultr VPC 2.0 resource. This can be used to create, read, and delete VPCs 2.0 on your Vultr account.

        ## Example Usage

        Create a new VPC 2.0 with an automatically generated CIDR block:

        ```python
        import pulumi
        import ediri_vultr as vultr

        my_vpc2 = vultr.Vpc2("myVpc2",
            description="my vpc2",
            region="ewr")
        ```

        Create a new VPC 2.0 with a specified CIDR block:

        ```python
        import pulumi
        import ediri_vultr as vultr

        my_vpc2 = vultr.Vpc2("myVpc2",
            description="my private vpc2",
            ip_block="10.0.0.0",
            prefix_length=24,
            region="ewr")
        ```

        ## Import

        VPCs 2.0 can be imported using the VPC 2.0 `ID`, e.g.

        ```sh
         $ pulumi import vultr:index/vpc2:Vpc2 my_vpc2 0e04f918-575e-41cb-86f6-d729b354a5a1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description you want to give your VPC 2.0.
        :param pulumi.Input[str] ip_block: The IPv4 subnet to be used when attaching instances to this VPC 2.0.
        :param pulumi.Input[str] ip_type: Accepted values: `v4`.
        :param pulumi.Input[int] prefix_length: The number of bits for the netmask in CIDR notation. Example: 32
        :param pulumi.Input[str] region: The region ID that you want the VPC 2.0 to be created in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Vpc2Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Vultr VPC 2.0 resource. This can be used to create, read, and delete VPCs 2.0 on your Vultr account.

        ## Example Usage

        Create a new VPC 2.0 with an automatically generated CIDR block:

        ```python
        import pulumi
        import ediri_vultr as vultr

        my_vpc2 = vultr.Vpc2("myVpc2",
            description="my vpc2",
            region="ewr")
        ```

        Create a new VPC 2.0 with a specified CIDR block:

        ```python
        import pulumi
        import ediri_vultr as vultr

        my_vpc2 = vultr.Vpc2("myVpc2",
            description="my private vpc2",
            ip_block="10.0.0.0",
            prefix_length=24,
            region="ewr")
        ```

        ## Import

        VPCs 2.0 can be imported using the VPC 2.0 `ID`, e.g.

        ```sh
         $ pulumi import vultr:index/vpc2:Vpc2 my_vpc2 0e04f918-575e-41cb-86f6-d729b354a5a1
        ```

        :param str resource_name: The name of the resource.
        :param Vpc2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(Vpc2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_block: Optional[pulumi.Input[str]] = None,
                 ip_type: Optional[pulumi.Input[str]] = None,
                 prefix_length: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = Vpc2Args.__new__(Vpc2Args)

            __props__.__dict__["description"] = description
            __props__.__dict__["ip_block"] = ip_block
            __props__.__dict__["ip_type"] = ip_type
            __props__.__dict__["prefix_length"] = prefix_length
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            __props__.__dict__["date_created"] = None
        super(Vpc2, __self__).__init__(
            'vultr:index/vpc2:Vpc2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            date_created: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            ip_block: Optional[pulumi.Input[str]] = None,
            ip_type: Optional[pulumi.Input[str]] = None,
            prefix_length: Optional[pulumi.Input[int]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'Vpc2':
        """
        Get an existing Vpc2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] date_created: The date that the VPC 2.0 was added to your Vultr account.
        :param pulumi.Input[str] description: The description you want to give your VPC 2.0.
        :param pulumi.Input[str] ip_block: The IPv4 subnet to be used when attaching instances to this VPC 2.0.
        :param pulumi.Input[str] ip_type: Accepted values: `v4`.
        :param pulumi.Input[int] prefix_length: The number of bits for the netmask in CIDR notation. Example: 32
        :param pulumi.Input[str] region: The region ID that you want the VPC 2.0 to be created in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _Vpc2State.__new__(_Vpc2State)

        __props__.__dict__["date_created"] = date_created
        __props__.__dict__["description"] = description
        __props__.__dict__["ip_block"] = ip_block
        __props__.__dict__["ip_type"] = ip_type
        __props__.__dict__["prefix_length"] = prefix_length
        __props__.__dict__["region"] = region
        return Vpc2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> pulumi.Output[str]:
        """
        The date that the VPC 2.0 was added to your Vultr account.
        """
        return pulumi.get(self, "date_created")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description you want to give your VPC 2.0.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ipBlock")
    def ip_block(self) -> pulumi.Output[Optional[str]]:
        """
        The IPv4 subnet to be used when attaching instances to this VPC 2.0.
        """
        return pulumi.get(self, "ip_block")

    @property
    @pulumi.getter(name="ipType")
    def ip_type(self) -> pulumi.Output[Optional[str]]:
        """
        Accepted values: `v4`.
        """
        return pulumi.get(self, "ip_type")

    @property
    @pulumi.getter(name="prefixLength")
    def prefix_length(self) -> pulumi.Output[Optional[int]]:
        """
        The number of bits for the netmask in CIDR notation. Example: 32
        """
        return pulumi.get(self, "prefix_length")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region ID that you want the VPC 2.0 to be created in.
        """
        return pulumi.get(self, "region")

