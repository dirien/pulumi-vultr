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
    'GetFirewallGroupResult',
    'AwaitableGetFirewallGroupResult',
    'get_firewall_group',
    'get_firewall_group_output',
]

@pulumi.output_type
class GetFirewallGroupResult:
    """
    A collection of values returned by getFirewallGroup.
    """
    def __init__(__self__, date_created=None, date_modified=None, description=None, filters=None, id=None, instance_count=None, max_rule_count=None, rule_count=None):
        if date_created and not isinstance(date_created, str):
            raise TypeError("Expected argument 'date_created' to be a str")
        pulumi.set(__self__, "date_created", date_created)
        if date_modified and not isinstance(date_modified, str):
            raise TypeError("Expected argument 'date_modified' to be a str")
        pulumi.set(__self__, "date_modified", date_modified)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_count and not isinstance(instance_count, int):
            raise TypeError("Expected argument 'instance_count' to be a int")
        pulumi.set(__self__, "instance_count", instance_count)
        if max_rule_count and not isinstance(max_rule_count, int):
            raise TypeError("Expected argument 'max_rule_count' to be a int")
        pulumi.set(__self__, "max_rule_count", max_rule_count)
        if rule_count and not isinstance(rule_count, int):
            raise TypeError("Expected argument 'rule_count' to be a int")
        pulumi.set(__self__, "rule_count", rule_count)

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> str:
        """
        The date the firewall group was added to your Vultr account.
        """
        return pulumi.get(self, "date_created")

    @property
    @pulumi.getter(name="dateModified")
    def date_modified(self) -> str:
        """
        The date the firewall group was last modified.
        """
        return pulumi.get(self, "date_modified")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the firewall group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetFirewallGroupFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceCount")
    def instance_count(self) -> int:
        """
        The number of instances this firewall group is applied to.
        """
        return pulumi.get(self, "instance_count")

    @property
    @pulumi.getter(name="maxRuleCount")
    def max_rule_count(self) -> int:
        """
        The maximum number of rules this firewall group can have.
        """
        return pulumi.get(self, "max_rule_count")

    @property
    @pulumi.getter(name="ruleCount")
    def rule_count(self) -> int:
        """
        The number of rules added to this firewall group.
        """
        return pulumi.get(self, "rule_count")


class AwaitableGetFirewallGroupResult(GetFirewallGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallGroupResult(
            date_created=self.date_created,
            date_modified=self.date_modified,
            description=self.description,
            filters=self.filters,
            id=self.id,
            instance_count=self.instance_count,
            max_rule_count=self.max_rule_count,
            rule_count=self.rule_count)


def get_firewall_group(filters: Optional[Sequence[pulumi.InputType['GetFirewallGroupFilterArgs']]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallGroupResult:
    """
    Get information about a firewall group on your Vultr account.

    ## Example Usage

    Get the information for a firewall group by `description`:

    ```python
    import pulumi
    import pulumi_vultr as vultr

    my_fwg = vultr.get_firewall_group(filters=[vultr.GetFirewallGroupFilterArgs(
        name="description",
        values=["fwg-description"],
    )])
    ```


    :param Sequence[pulumi.InputType['GetFirewallGroupFilterArgs']] filters: Query parameters for finding firewall groups.
    """
    __args__ = dict()
    __args__['filters'] = filters
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vultr:index/getFirewallGroup:getFirewallGroup', __args__, opts=opts, typ=GetFirewallGroupResult).value

    return AwaitableGetFirewallGroupResult(
        date_created=pulumi.get(__ret__, 'date_created'),
        date_modified=pulumi.get(__ret__, 'date_modified'),
        description=pulumi.get(__ret__, 'description'),
        filters=pulumi.get(__ret__, 'filters'),
        id=pulumi.get(__ret__, 'id'),
        instance_count=pulumi.get(__ret__, 'instance_count'),
        max_rule_count=pulumi.get(__ret__, 'max_rule_count'),
        rule_count=pulumi.get(__ret__, 'rule_count'))


@_utilities.lift_output_func(get_firewall_group)
def get_firewall_group_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetFirewallGroupFilterArgs']]]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFirewallGroupResult]:
    """
    Get information about a firewall group on your Vultr account.

    ## Example Usage

    Get the information for a firewall group by `description`:

    ```python
    import pulumi
    import pulumi_vultr as vultr

    my_fwg = vultr.get_firewall_group(filters=[vultr.GetFirewallGroupFilterArgs(
        name="description",
        values=["fwg-description"],
    )])
    ```


    :param Sequence[pulumi.InputType['GetFirewallGroupFilterArgs']] filters: Query parameters for finding firewall groups.
    """
    ...
