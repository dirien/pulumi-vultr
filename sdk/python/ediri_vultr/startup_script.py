# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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

__all__ = ['StartupScriptArgs', 'StartupScript']

@pulumi.input_type
class StartupScriptArgs:
    def __init__(__self__, *,
                 script: pulumi.Input[builtins.str],
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a StartupScript resource.
        :param pulumi.Input[builtins.str] script: Contents of the startup script base64 encoded.
        :param pulumi.Input[builtins.str] name: Name of the given script.
        :param pulumi.Input[builtins.str] type: Type of startup script. Possible values are boot or pxe - default is boot.
        """
        pulumi.set(__self__, "script", script)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def script(self) -> pulumi.Input[builtins.str]:
        """
        Contents of the startup script base64 encoded.
        """
        return pulumi.get(self, "script")

    @script.setter
    def script(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "script", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the given script.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Type of startup script. Possible values are boot or pxe - default is boot.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _StartupScriptState:
    def __init__(__self__, *,
                 date_created: Optional[pulumi.Input[builtins.str]] = None,
                 date_modified: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 script: Optional[pulumi.Input[builtins.str]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering StartupScript resources.
        :param pulumi.Input[builtins.str] date_created: Date the script was created.
        :param pulumi.Input[builtins.str] date_modified: Date the script was last modified.
        :param pulumi.Input[builtins.str] name: Name of the given script.
        :param pulumi.Input[builtins.str] script: Contents of the startup script base64 encoded.
        :param pulumi.Input[builtins.str] type: Type of startup script. Possible values are boot or pxe - default is boot.
        """
        if date_created is not None:
            pulumi.set(__self__, "date_created", date_created)
        if date_modified is not None:
            pulumi.set(__self__, "date_modified", date_modified)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if script is not None:
            pulumi.set(__self__, "script", script)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Date the script was created.
        """
        return pulumi.get(self, "date_created")

    @date_created.setter
    def date_created(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "date_created", value)

    @property
    @pulumi.getter(name="dateModified")
    def date_modified(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Date the script was last modified.
        """
        return pulumi.get(self, "date_modified")

    @date_modified.setter
    def date_modified(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "date_modified", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the given script.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def script(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Contents of the startup script base64 encoded.
        """
        return pulumi.get(self, "script")

    @script.setter
    def script(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "script", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Type of startup script. Possible values are boot or pxe - default is boot.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "type", value)


@pulumi.type_token("vultr:index/startupScript:StartupScript")
class StartupScript(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 script: Optional[pulumi.Input[builtins.str]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Provides a Vultr Startup Script resource. This can be used to create, read, modify, and delete Startup Scripts.

        ## Example Usage

        Create a new Startup Script

        ```python
        import pulumi
        import ediri_vultr as vultr

        my_script = vultr.StartupScript("myScript", script="ZWNobyAkUEFUSAo=")
        ```

        ## Import

        Startup Scripts can be imported using the Startup Scripts `ID`, e.g.

        ```sh
        $ pulumi import vultr:index/startupScript:StartupScript my_script ff8f36a8-eb86-4b8d-8667-b9d5459b6390
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] name: Name of the given script.
        :param pulumi.Input[builtins.str] script: Contents of the startup script base64 encoded.
        :param pulumi.Input[builtins.str] type: Type of startup script. Possible values are boot or pxe - default is boot.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StartupScriptArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Vultr Startup Script resource. This can be used to create, read, modify, and delete Startup Scripts.

        ## Example Usage

        Create a new Startup Script

        ```python
        import pulumi
        import ediri_vultr as vultr

        my_script = vultr.StartupScript("myScript", script="ZWNobyAkUEFUSAo=")
        ```

        ## Import

        Startup Scripts can be imported using the Startup Scripts `ID`, e.g.

        ```sh
        $ pulumi import vultr:index/startupScript:StartupScript my_script ff8f36a8-eb86-4b8d-8667-b9d5459b6390
        ```

        :param str resource_name: The name of the resource.
        :param StartupScriptArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StartupScriptArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 script: Optional[pulumi.Input[builtins.str]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StartupScriptArgs.__new__(StartupScriptArgs)

            __props__.__dict__["name"] = name
            if script is None and not opts.urn:
                raise TypeError("Missing required property 'script'")
            __props__.__dict__["script"] = script
            __props__.__dict__["type"] = type
            __props__.__dict__["date_created"] = None
            __props__.__dict__["date_modified"] = None
        super(StartupScript, __self__).__init__(
            'vultr:index/startupScript:StartupScript',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            date_created: Optional[pulumi.Input[builtins.str]] = None,
            date_modified: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            script: Optional[pulumi.Input[builtins.str]] = None,
            type: Optional[pulumi.Input[builtins.str]] = None) -> 'StartupScript':
        """
        Get an existing StartupScript resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] date_created: Date the script was created.
        :param pulumi.Input[builtins.str] date_modified: Date the script was last modified.
        :param pulumi.Input[builtins.str] name: Name of the given script.
        :param pulumi.Input[builtins.str] script: Contents of the startup script base64 encoded.
        :param pulumi.Input[builtins.str] type: Type of startup script. Possible values are boot or pxe - default is boot.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StartupScriptState.__new__(_StartupScriptState)

        __props__.__dict__["date_created"] = date_created
        __props__.__dict__["date_modified"] = date_modified
        __props__.__dict__["name"] = name
        __props__.__dict__["script"] = script
        __props__.__dict__["type"] = type
        return StartupScript(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> pulumi.Output[builtins.str]:
        """
        Date the script was created.
        """
        return pulumi.get(self, "date_created")

    @property
    @pulumi.getter(name="dateModified")
    def date_modified(self) -> pulumi.Output[builtins.str]:
        """
        Date the script was last modified.
        """
        return pulumi.get(self, "date_modified")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Name of the given script.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def script(self) -> pulumi.Output[builtins.str]:
        """
        Contents of the startup script base64 encoded.
        """
        return pulumi.get(self, "script")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Type of startup script. Possible values are boot or pxe - default is boot.
        """
        return pulumi.get(self, "type")

