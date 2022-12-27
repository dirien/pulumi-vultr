# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ObjectStorageArgs', 'ObjectStorage']

@pulumi.input_type
class ObjectStorageArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[int],
                 label: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ObjectStorage resource.
        :param pulumi.Input[int] cluster_id: The region ID that you want the network to be created in.
        :param pulumi.Input[str] label: The description you want to give your network.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        if label is not None:
            pulumi.set(__self__, "label", label)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[int]:
        """
        The region ID that you want the network to be created in.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def label(self) -> Optional[pulumi.Input[str]]:
        """
        The description you want to give your network.
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "label", value)


@pulumi.input_type
class _ObjectStorageState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[int]] = None,
                 date_created: Optional[pulumi.Input[str]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 s3_access_key: Optional[pulumi.Input[str]] = None,
                 s3_hostname: Optional[pulumi.Input[str]] = None,
                 s3_secret_key: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ObjectStorage resources.
        :param pulumi.Input[int] cluster_id: The region ID that you want the network to be created in.
        :param pulumi.Input[str] date_created: Date of creation for the object storage subscription.
        :param pulumi.Input[str] label: The description you want to give your network.
        :param pulumi.Input[str] location: The location which this subscription resides in.
        :param pulumi.Input[str] region: The region ID of the object storage subscription.
        :param pulumi.Input[str] s3_access_key: Your access key.
        :param pulumi.Input[str] s3_hostname: The hostname for this subscription.
        :param pulumi.Input[str] s3_secret_key: Your secret key.
        :param pulumi.Input[str] status: Current status of this object storage subscription.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if date_created is not None:
            pulumi.set(__self__, "date_created", date_created)
        if label is not None:
            pulumi.set(__self__, "label", label)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if s3_access_key is not None:
            pulumi.set(__self__, "s3_access_key", s3_access_key)
        if s3_hostname is not None:
            pulumi.set(__self__, "s3_hostname", s3_hostname)
        if s3_secret_key is not None:
            pulumi.set(__self__, "s3_secret_key", s3_secret_key)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[int]]:
        """
        The region ID that you want the network to be created in.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> Optional[pulumi.Input[str]]:
        """
        Date of creation for the object storage subscription.
        """
        return pulumi.get(self, "date_created")

    @date_created.setter
    def date_created(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "date_created", value)

    @property
    @pulumi.getter
    def label(self) -> Optional[pulumi.Input[str]]:
        """
        The description you want to give your network.
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The location which this subscription resides in.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region ID of the object storage subscription.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="s3AccessKey")
    def s3_access_key(self) -> Optional[pulumi.Input[str]]:
        """
        Your access key.
        """
        return pulumi.get(self, "s3_access_key")

    @s3_access_key.setter
    def s3_access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "s3_access_key", value)

    @property
    @pulumi.getter(name="s3Hostname")
    def s3_hostname(self) -> Optional[pulumi.Input[str]]:
        """
        The hostname for this subscription.
        """
        return pulumi.get(self, "s3_hostname")

    @s3_hostname.setter
    def s3_hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "s3_hostname", value)

    @property
    @pulumi.getter(name="s3SecretKey")
    def s3_secret_key(self) -> Optional[pulumi.Input[str]]:
        """
        Your secret key.
        """
        return pulumi.get(self, "s3_secret_key")

    @s3_secret_key.setter
    def s3_secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "s3_secret_key", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Current status of this object storage subscription.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class ObjectStorage(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[int]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Vultr private object storage resource. This can be used to create, read, update and delete object storage resources on your Vultr account.

        ## Example Usage

        Create a new object storage subscription.

        ```python
        import pulumi
        import dirien_vultr as vultr

        tf = vultr.ObjectStorage("tf",
            cluster_id=2,
            label="tf-label")
        ```

        ## Import

        Object Storage can be imported using the object storage `ID`, e.g.

        ```sh
         $ pulumi import vultr:index/objectStorage:ObjectStorage my_s3 0e04f918-575e-41cb-86f6-d729b354a5a1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] cluster_id: The region ID that you want the network to be created in.
        :param pulumi.Input[str] label: The description you want to give your network.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ObjectStorageArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Vultr private object storage resource. This can be used to create, read, update and delete object storage resources on your Vultr account.

        ## Example Usage

        Create a new object storage subscription.

        ```python
        import pulumi
        import dirien_vultr as vultr

        tf = vultr.ObjectStorage("tf",
            cluster_id=2,
            label="tf-label")
        ```

        ## Import

        Object Storage can be imported using the object storage `ID`, e.g.

        ```sh
         $ pulumi import vultr:index/objectStorage:ObjectStorage my_s3 0e04f918-575e-41cb-86f6-d729b354a5a1
        ```

        :param str resource_name: The name of the resource.
        :param ObjectStorageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ObjectStorageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[int]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ObjectStorageArgs.__new__(ObjectStorageArgs)

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            __props__.__dict__["label"] = label
            __props__.__dict__["date_created"] = None
            __props__.__dict__["location"] = None
            __props__.__dict__["region"] = None
            __props__.__dict__["s3_access_key"] = None
            __props__.__dict__["s3_hostname"] = None
            __props__.__dict__["s3_secret_key"] = None
            __props__.__dict__["status"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["s3AccessKey", "s3SecretKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ObjectStorage, __self__).__init__(
            'vultr:index/objectStorage:ObjectStorage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[int]] = None,
            date_created: Optional[pulumi.Input[str]] = None,
            label: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            s3_access_key: Optional[pulumi.Input[str]] = None,
            s3_hostname: Optional[pulumi.Input[str]] = None,
            s3_secret_key: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'ObjectStorage':
        """
        Get an existing ObjectStorage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] cluster_id: The region ID that you want the network to be created in.
        :param pulumi.Input[str] date_created: Date of creation for the object storage subscription.
        :param pulumi.Input[str] label: The description you want to give your network.
        :param pulumi.Input[str] location: The location which this subscription resides in.
        :param pulumi.Input[str] region: The region ID of the object storage subscription.
        :param pulumi.Input[str] s3_access_key: Your access key.
        :param pulumi.Input[str] s3_hostname: The hostname for this subscription.
        :param pulumi.Input[str] s3_secret_key: Your secret key.
        :param pulumi.Input[str] status: Current status of this object storage subscription.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ObjectStorageState.__new__(_ObjectStorageState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["date_created"] = date_created
        __props__.__dict__["label"] = label
        __props__.__dict__["location"] = location
        __props__.__dict__["region"] = region
        __props__.__dict__["s3_access_key"] = s3_access_key
        __props__.__dict__["s3_hostname"] = s3_hostname
        __props__.__dict__["s3_secret_key"] = s3_secret_key
        __props__.__dict__["status"] = status
        return ObjectStorage(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[int]:
        """
        The region ID that you want the network to be created in.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> pulumi.Output[str]:
        """
        Date of creation for the object storage subscription.
        """
        return pulumi.get(self, "date_created")

    @property
    @pulumi.getter
    def label(self) -> pulumi.Output[Optional[str]]:
        """
        The description you want to give your network.
        """
        return pulumi.get(self, "label")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The location which this subscription resides in.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region ID of the object storage subscription.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="s3AccessKey")
    def s3_access_key(self) -> pulumi.Output[str]:
        """
        Your access key.
        """
        return pulumi.get(self, "s3_access_key")

    @property
    @pulumi.getter(name="s3Hostname")
    def s3_hostname(self) -> pulumi.Output[str]:
        """
        The hostname for this subscription.
        """
        return pulumi.get(self, "s3_hostname")

    @property
    @pulumi.getter(name="s3SecretKey")
    def s3_secret_key(self) -> pulumi.Output[str]:
        """
        Your secret key.
        """
        return pulumi.get(self, "s3_secret_key")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Current status of this object storage subscription.
        """
        return pulumi.get(self, "status")

