# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FirewallGroupArgs', 'FirewallGroup']

@pulumi.input_type
class FirewallGroupArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallGroup resource.
        :param pulumi.Input[str] description: Description of the firewall group.
        """
        FirewallGroupArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            description=description,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             description: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):

        if description is not None:
            _setter("description", description)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the firewall group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class _FirewallGroupState:
    def __init__(__self__, *,
                 date_created: Optional[pulumi.Input[str]] = None,
                 date_modified: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_count: Optional[pulumi.Input[int]] = None,
                 max_rule_count: Optional[pulumi.Input[int]] = None,
                 rule_count: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering FirewallGroup resources.
        :param pulumi.Input[str] date_created: The date the firewall group was created.
        :param pulumi.Input[str] date_modified: The date the firewall group was modified.
        :param pulumi.Input[str] description: Description of the firewall group.
        :param pulumi.Input[int] instance_count: The number of instances that are currently using this firewall group.
        :param pulumi.Input[int] max_rule_count: The number of max firewall rules this group can have.
        :param pulumi.Input[int] rule_count: The number of firewall rules this group currently has.
        """
        _FirewallGroupState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            date_created=date_created,
            date_modified=date_modified,
            description=description,
            instance_count=instance_count,
            max_rule_count=max_rule_count,
            rule_count=rule_count,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             date_created: Optional[pulumi.Input[str]] = None,
             date_modified: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             instance_count: Optional[pulumi.Input[int]] = None,
             max_rule_count: Optional[pulumi.Input[int]] = None,
             rule_count: Optional[pulumi.Input[int]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'dateCreated' in kwargs:
            date_created = kwargs['dateCreated']
        if 'dateModified' in kwargs:
            date_modified = kwargs['dateModified']
        if 'instanceCount' in kwargs:
            instance_count = kwargs['instanceCount']
        if 'maxRuleCount' in kwargs:
            max_rule_count = kwargs['maxRuleCount']
        if 'ruleCount' in kwargs:
            rule_count = kwargs['ruleCount']

        if date_created is not None:
            _setter("date_created", date_created)
        if date_modified is not None:
            _setter("date_modified", date_modified)
        if description is not None:
            _setter("description", description)
        if instance_count is not None:
            _setter("instance_count", instance_count)
        if max_rule_count is not None:
            _setter("max_rule_count", max_rule_count)
        if rule_count is not None:
            _setter("rule_count", rule_count)

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> Optional[pulumi.Input[str]]:
        """
        The date the firewall group was created.
        """
        return pulumi.get(self, "date_created")

    @date_created.setter
    def date_created(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "date_created", value)

    @property
    @pulumi.getter(name="dateModified")
    def date_modified(self) -> Optional[pulumi.Input[str]]:
        """
        The date the firewall group was modified.
        """
        return pulumi.get(self, "date_modified")

    @date_modified.setter
    def date_modified(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "date_modified", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the firewall group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="instanceCount")
    def instance_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of instances that are currently using this firewall group.
        """
        return pulumi.get(self, "instance_count")

    @instance_count.setter
    def instance_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_count", value)

    @property
    @pulumi.getter(name="maxRuleCount")
    def max_rule_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of max firewall rules this group can have.
        """
        return pulumi.get(self, "max_rule_count")

    @max_rule_count.setter
    def max_rule_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_rule_count", value)

    @property
    @pulumi.getter(name="ruleCount")
    def rule_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of firewall rules this group currently has.
        """
        return pulumi.get(self, "rule_count")

    @rule_count.setter
    def rule_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rule_count", value)


class FirewallGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Vultr Firewall Group resource. This can be used to create, read, modify, and delete Firewall Group.

        ## Example Usage

        Create a new Firewall group

        ```python
        import pulumi
        import ediri_vultr as vultr

        my_firewallgroup = vultr.FirewallGroup("myFirewallgroup", description="base firewall")
        ```

        ## Import

        Firewall Groups can be imported using the Firewall Group `FIREWALLGROUPID`, e.g.

        ```sh
         $ pulumi import vultr:index/firewallGroup:FirewallGroup my_firewallgroup c342f929
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the firewall group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FirewallGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Vultr Firewall Group resource. This can be used to create, read, modify, and delete Firewall Group.

        ## Example Usage

        Create a new Firewall group

        ```python
        import pulumi
        import ediri_vultr as vultr

        my_firewallgroup = vultr.FirewallGroup("myFirewallgroup", description="base firewall")
        ```

        ## Import

        Firewall Groups can be imported using the Firewall Group `FIREWALLGROUPID`, e.g.

        ```sh
         $ pulumi import vultr:index/firewallGroup:FirewallGroup my_firewallgroup c342f929
        ```

        :param str resource_name: The name of the resource.
        :param FirewallGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            FirewallGroupArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallGroupArgs.__new__(FirewallGroupArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["date_created"] = None
            __props__.__dict__["date_modified"] = None
            __props__.__dict__["instance_count"] = None
            __props__.__dict__["max_rule_count"] = None
            __props__.__dict__["rule_count"] = None
        super(FirewallGroup, __self__).__init__(
            'vultr:index/firewallGroup:FirewallGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            date_created: Optional[pulumi.Input[str]] = None,
            date_modified: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            instance_count: Optional[pulumi.Input[int]] = None,
            max_rule_count: Optional[pulumi.Input[int]] = None,
            rule_count: Optional[pulumi.Input[int]] = None) -> 'FirewallGroup':
        """
        Get an existing FirewallGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] date_created: The date the firewall group was created.
        :param pulumi.Input[str] date_modified: The date the firewall group was modified.
        :param pulumi.Input[str] description: Description of the firewall group.
        :param pulumi.Input[int] instance_count: The number of instances that are currently using this firewall group.
        :param pulumi.Input[int] max_rule_count: The number of max firewall rules this group can have.
        :param pulumi.Input[int] rule_count: The number of firewall rules this group currently has.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallGroupState.__new__(_FirewallGroupState)

        __props__.__dict__["date_created"] = date_created
        __props__.__dict__["date_modified"] = date_modified
        __props__.__dict__["description"] = description
        __props__.__dict__["instance_count"] = instance_count
        __props__.__dict__["max_rule_count"] = max_rule_count
        __props__.__dict__["rule_count"] = rule_count
        return FirewallGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> pulumi.Output[str]:
        """
        The date the firewall group was created.
        """
        return pulumi.get(self, "date_created")

    @property
    @pulumi.getter(name="dateModified")
    def date_modified(self) -> pulumi.Output[str]:
        """
        The date the firewall group was modified.
        """
        return pulumi.get(self, "date_modified")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the firewall group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="instanceCount")
    def instance_count(self) -> pulumi.Output[int]:
        """
        The number of instances that are currently using this firewall group.
        """
        return pulumi.get(self, "instance_count")

    @property
    @pulumi.getter(name="maxRuleCount")
    def max_rule_count(self) -> pulumi.Output[int]:
        """
        The number of max firewall rules this group can have.
        """
        return pulumi.get(self, "max_rule_count")

    @property
    @pulumi.getter(name="ruleCount")
    def rule_count(self) -> pulumi.Output[int]:
        """
        The number of firewall rules this group currently has.
        """
        return pulumi.get(self, "rule_count")

