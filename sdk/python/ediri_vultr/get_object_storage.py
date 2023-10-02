# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetObjectStorageResult',
    'AwaitableGetObjectStorageResult',
    'get_object_storage',
    'get_object_storage_output',
]

@pulumi.output_type
class GetObjectStorageResult:
    """
    A collection of values returned by getObjectStorage.
    """
    def __init__(__self__, cluster_id=None, date_created=None, filters=None, id=None, label=None, location=None, region=None, s3_access_key=None, s3_hostname=None, s3_secret_key=None, status=None):
        if cluster_id and not isinstance(cluster_id, int):
            raise TypeError("Expected argument 'cluster_id' to be a int")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if date_created and not isinstance(date_created, str):
            raise TypeError("Expected argument 'date_created' to be a str")
        pulumi.set(__self__, "date_created", date_created)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if label and not isinstance(label, str):
            raise TypeError("Expected argument 'label' to be a str")
        pulumi.set(__self__, "label", label)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if s3_access_key and not isinstance(s3_access_key, str):
            raise TypeError("Expected argument 's3_access_key' to be a str")
        pulumi.set(__self__, "s3_access_key", s3_access_key)
        if s3_hostname and not isinstance(s3_hostname, str):
            raise TypeError("Expected argument 's3_hostname' to be a str")
        pulumi.set(__self__, "s3_hostname", s3_hostname)
        if s3_secret_key and not isinstance(s3_secret_key, str):
            raise TypeError("Expected argument 's3_secret_key' to be a str")
        pulumi.set(__self__, "s3_secret_key", s3_secret_key)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> int:
        """
        The identifying cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> str:
        """
        Date of creation for the object storage subscription.
        """
        return pulumi.get(self, "date_created")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetObjectStorageFilterResult']]:
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
    def label(self) -> str:
        """
        The label of the object storage subscription.
        """
        return pulumi.get(self, "label")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The location which this subscription resides in.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        The region ID of the object storage subscription.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="s3AccessKey")
    def s3_access_key(self) -> str:
        """
        Your access key.
        """
        return pulumi.get(self, "s3_access_key")

    @property
    @pulumi.getter(name="s3Hostname")
    def s3_hostname(self) -> str:
        """
        The hostname for this subscription.
        """
        return pulumi.get(self, "s3_hostname")

    @property
    @pulumi.getter(name="s3SecretKey")
    def s3_secret_key(self) -> str:
        """
        Your secret key.
        """
        return pulumi.get(self, "s3_secret_key")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Current status of this object storage subscription.
        """
        return pulumi.get(self, "status")


class AwaitableGetObjectStorageResult(GetObjectStorageResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetObjectStorageResult(
            cluster_id=self.cluster_id,
            date_created=self.date_created,
            filters=self.filters,
            id=self.id,
            label=self.label,
            location=self.location,
            region=self.region,
            s3_access_key=self.s3_access_key,
            s3_hostname=self.s3_hostname,
            s3_secret_key=self.s3_secret_key,
            status=self.status)


def get_object_storage(filters: Optional[Sequence[pulumi.InputType['GetObjectStorageFilterArgs']]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetObjectStorageResult:
    """
    Get information about an Object Storage subscription on Vultr.

    ## Example Usage

    Get the information for an object storage subscription by `label`:

    ```python
    import pulumi
    import pulumi_vultr as vultr

    s3 = vultr.get_object_storage(filters=[vultr.GetObjectStorageFilterArgs(
        name="label",
        values=["my-s3"],
    )])
    ```


    :param Sequence[pulumi.InputType['GetObjectStorageFilterArgs']] filters: Query parameters for finding operating systems.
    """
    __args__ = dict()
    __args__['filters'] = filters
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vultr:index/getObjectStorage:getObjectStorage', __args__, opts=opts, typ=GetObjectStorageResult).value

    return AwaitableGetObjectStorageResult(
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        date_created=pulumi.get(__ret__, 'date_created'),
        filters=pulumi.get(__ret__, 'filters'),
        id=pulumi.get(__ret__, 'id'),
        label=pulumi.get(__ret__, 'label'),
        location=pulumi.get(__ret__, 'location'),
        region=pulumi.get(__ret__, 'region'),
        s3_access_key=pulumi.get(__ret__, 's3_access_key'),
        s3_hostname=pulumi.get(__ret__, 's3_hostname'),
        s3_secret_key=pulumi.get(__ret__, 's3_secret_key'),
        status=pulumi.get(__ret__, 'status'))


@_utilities.lift_output_func(get_object_storage)
def get_object_storage_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetObjectStorageFilterArgs']]]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetObjectStorageResult]:
    """
    Get information about an Object Storage subscription on Vultr.

    ## Example Usage

    Get the information for an object storage subscription by `label`:

    ```python
    import pulumi
    import pulumi_vultr as vultr

    s3 = vultr.get_object_storage(filters=[vultr.GetObjectStorageFilterArgs(
        name="label",
        values=["my-s3"],
    )])
    ```


    :param Sequence[pulumi.InputType['GetObjectStorageFilterArgs']] filters: Query parameters for finding operating systems.
    """
    ...
