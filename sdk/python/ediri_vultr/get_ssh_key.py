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
    'GetSshKeyResult',
    'AwaitableGetSshKeyResult',
    'get_ssh_key',
    'get_ssh_key_output',
]

@pulumi.output_type
class GetSshKeyResult:
    """
    A collection of values returned by getSshKey.
    """
    def __init__(__self__, date_created=None, filters=None, id=None, name=None, ssh_key=None):
        if date_created and not isinstance(date_created, str):
            raise TypeError("Expected argument 'date_created' to be a str")
        pulumi.set(__self__, "date_created", date_created)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if ssh_key and not isinstance(ssh_key, str):
            raise TypeError("Expected argument 'ssh_key' to be a str")
        pulumi.set(__self__, "ssh_key", ssh_key)

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> str:
        """
        The date the SSH key was added to your Vultr account.
        """
        return pulumi.get(self, "date_created")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetSshKeyFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the SSH key.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="sshKey")
    def ssh_key(self) -> str:
        """
        The public SSH key.
        """
        return pulumi.get(self, "ssh_key")


class AwaitableGetSshKeyResult(GetSshKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSshKeyResult(
            date_created=self.date_created,
            filters=self.filters,
            id=self.id,
            name=self.name,
            ssh_key=self.ssh_key)


def get_ssh_key(filters: Optional[Sequence[pulumi.InputType['GetSshKeyFilterArgs']]] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSshKeyResult:
    """
    Get information about a Vultr SSH key. This data source provides the name, public SSH key, and the creation date for your Vultr SSH key.

    ## Example Usage

    Get the information for an SSH key by `name`:

    ```python
    import pulumi
    import pulumi_vultr as vultr

    my_ssh_key = vultr.get_ssh_key(filters=[vultr.GetSshKeyFilterArgs(
        name="name",
        values=["my-ssh-key-name"],
    )])
    ```


    :param Sequence[pulumi.InputType['GetSshKeyFilterArgs']] filters: Query parameters for finding SSH keys.
    """
    __args__ = dict()
    __args__['filters'] = filters
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vultr:index/getSshKey:getSshKey', __args__, opts=opts, typ=GetSshKeyResult).value

    return AwaitableGetSshKeyResult(
        date_created=pulumi.get(__ret__, 'date_created'),
        filters=pulumi.get(__ret__, 'filters'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        ssh_key=pulumi.get(__ret__, 'ssh_key'))


@_utilities.lift_output_func(get_ssh_key)
def get_ssh_key_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetSshKeyFilterArgs']]]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSshKeyResult]:
    """
    Get information about a Vultr SSH key. This data source provides the name, public SSH key, and the creation date for your Vultr SSH key.

    ## Example Usage

    Get the information for an SSH key by `name`:

    ```python
    import pulumi
    import pulumi_vultr as vultr

    my_ssh_key = vultr.get_ssh_key(filters=[vultr.GetSshKeyFilterArgs(
        name="name",
        values=["my-ssh-key-name"],
    )])
    ```


    :param Sequence[pulumi.InputType['GetSshKeyFilterArgs']] filters: Query parameters for finding SSH keys.
    """
    ...
