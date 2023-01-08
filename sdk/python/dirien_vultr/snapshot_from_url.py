# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SnapshotFromUrlArgs', 'SnapshotFromUrl']

@pulumi.input_type
class SnapshotFromUrlArgs:
    def __init__(__self__, *,
                 url: pulumi.Input[str]):
        """
        The set of arguments for constructing a SnapshotFromUrl resource.
        :param pulumi.Input[str] url: URL of the given resource you want to create a snapshot from.
        """
        pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def url(self) -> pulumi.Input[str]:
        """
        URL of the given resource you want to create a snapshot from.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)


@pulumi.input_type
class _SnapshotFromUrlState:
    def __init__(__self__, *,
                 app_id: Optional[pulumi.Input[int]] = None,
                 date_created: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 os_id: Optional[pulumi.Input[int]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SnapshotFromUrl resources.
        :param pulumi.Input[int] app_id: The app id which the snapshot is associated with.
        :param pulumi.Input[str] date_created: The date the snapshot was created.
        :param pulumi.Input[str] description: The description for the given snapshot.
        :param pulumi.Input[int] os_id: The os id which the snapshot is associated with.
        :param pulumi.Input[int] size: The size of the snapshot in Bytes.
        :param pulumi.Input[str] status: The status for the given snapshot.
        :param pulumi.Input[str] url: URL of the given resource you want to create a snapshot from.
        """
        if app_id is not None:
            pulumi.set(__self__, "app_id", app_id)
        if date_created is not None:
            pulumi.set(__self__, "date_created", date_created)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if os_id is not None:
            pulumi.set(__self__, "os_id", os_id)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> Optional[pulumi.Input[int]]:
        """
        The app id which the snapshot is associated with.
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> Optional[pulumi.Input[str]]:
        """
        The date the snapshot was created.
        """
        return pulumi.get(self, "date_created")

    @date_created.setter
    def date_created(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "date_created", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description for the given snapshot.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="osId")
    def os_id(self) -> Optional[pulumi.Input[int]]:
        """
        The os id which the snapshot is associated with.
        """
        return pulumi.get(self, "os_id")

    @os_id.setter
    def os_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "os_id", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[int]]:
        """
        The size of the snapshot in Bytes.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status for the given snapshot.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        URL of the given resource you want to create a snapshot from.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class SnapshotFromUrl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Vultr Snapshots from URL resource. This can be used to create, read, modify, and delete Snapshots from URL.

        ## Example Usage

        Create a new Snapshots from URL

        ```python
        import pulumi
        import dirien_vultr as vultr

        my_snapshot = vultr.SnapshotFromUrl("mySnapshot", url="http://dl-cdn.alpinelinux.org/alpine/v3.9/releases/x86_64/alpine-virt-3.9.1-x86_64.iso")
        ```

        ## Import

        Snapshots from Url can be imported using the Snapshot `ID`, e.g.

        ```sh
         $ pulumi import vultr:index/snapshotFromUrl:SnapshotFromUrl my_snapshot e60dc0a2-9313-4bab-bffc-57ffe33d99f6
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] url: URL of the given resource you want to create a snapshot from.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SnapshotFromUrlArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Vultr Snapshots from URL resource. This can be used to create, read, modify, and delete Snapshots from URL.

        ## Example Usage

        Create a new Snapshots from URL

        ```python
        import pulumi
        import dirien_vultr as vultr

        my_snapshot = vultr.SnapshotFromUrl("mySnapshot", url="http://dl-cdn.alpinelinux.org/alpine/v3.9/releases/x86_64/alpine-virt-3.9.1-x86_64.iso")
        ```

        ## Import

        Snapshots from Url can be imported using the Snapshot `ID`, e.g.

        ```sh
         $ pulumi import vultr:index/snapshotFromUrl:SnapshotFromUrl my_snapshot e60dc0a2-9313-4bab-bffc-57ffe33d99f6
        ```

        :param str resource_name: The name of the resource.
        :param SnapshotFromUrlArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SnapshotFromUrlArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SnapshotFromUrlArgs.__new__(SnapshotFromUrlArgs)

            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = url
            __props__.__dict__["app_id"] = None
            __props__.__dict__["date_created"] = None
            __props__.__dict__["description"] = None
            __props__.__dict__["os_id"] = None
            __props__.__dict__["size"] = None
            __props__.__dict__["status"] = None
        super(SnapshotFromUrl, __self__).__init__(
            'vultr:index/snapshotFromUrl:SnapshotFromUrl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_id: Optional[pulumi.Input[int]] = None,
            date_created: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            os_id: Optional[pulumi.Input[int]] = None,
            size: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'SnapshotFromUrl':
        """
        Get an existing SnapshotFromUrl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] app_id: The app id which the snapshot is associated with.
        :param pulumi.Input[str] date_created: The date the snapshot was created.
        :param pulumi.Input[str] description: The description for the given snapshot.
        :param pulumi.Input[int] os_id: The os id which the snapshot is associated with.
        :param pulumi.Input[int] size: The size of the snapshot in Bytes.
        :param pulumi.Input[str] status: The status for the given snapshot.
        :param pulumi.Input[str] url: URL of the given resource you want to create a snapshot from.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SnapshotFromUrlState.__new__(_SnapshotFromUrlState)

        __props__.__dict__["app_id"] = app_id
        __props__.__dict__["date_created"] = date_created
        __props__.__dict__["description"] = description
        __props__.__dict__["os_id"] = os_id
        __props__.__dict__["size"] = size
        __props__.__dict__["status"] = status
        __props__.__dict__["url"] = url
        return SnapshotFromUrl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Output[int]:
        """
        The app id which the snapshot is associated with.
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> pulumi.Output[str]:
        """
        The date the snapshot was created.
        """
        return pulumi.get(self, "date_created")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description for the given snapshot.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="osId")
    def os_id(self) -> pulumi.Output[int]:
        """
        The os id which the snapshot is associated with.
        """
        return pulumi.get(self, "os_id")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[int]:
        """
        The size of the snapshot in Bytes.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status for the given snapshot.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        URL of the given resource you want to create a snapshot from.
        """
        return pulumi.get(self, "url")
