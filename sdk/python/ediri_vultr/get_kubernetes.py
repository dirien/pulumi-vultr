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
    'GetKubernetesResult',
    'AwaitableGetKubernetesResult',
    'get_kubernetes',
    'get_kubernetes_output',
]

@pulumi.output_type
class GetKubernetesResult:
    """
    A collection of values returned by getKubernetes.
    """
    def __init__(__self__, client_certificate=None, client_key=None, cluster_ca_certificate=None, cluster_subnet=None, date_created=None, endpoint=None, filters=None, id=None, ip=None, kube_config=None, label=None, node_pools=None, region=None, service_subnet=None, status=None, version=None):
        if client_certificate and not isinstance(client_certificate, str):
            raise TypeError("Expected argument 'client_certificate' to be a str")
        pulumi.set(__self__, "client_certificate", client_certificate)
        if client_key and not isinstance(client_key, str):
            raise TypeError("Expected argument 'client_key' to be a str")
        pulumi.set(__self__, "client_key", client_key)
        if cluster_ca_certificate and not isinstance(cluster_ca_certificate, str):
            raise TypeError("Expected argument 'cluster_ca_certificate' to be a str")
        pulumi.set(__self__, "cluster_ca_certificate", cluster_ca_certificate)
        if cluster_subnet and not isinstance(cluster_subnet, str):
            raise TypeError("Expected argument 'cluster_subnet' to be a str")
        pulumi.set(__self__, "cluster_subnet", cluster_subnet)
        if date_created and not isinstance(date_created, str):
            raise TypeError("Expected argument 'date_created' to be a str")
        pulumi.set(__self__, "date_created", date_created)
        if endpoint and not isinstance(endpoint, str):
            raise TypeError("Expected argument 'endpoint' to be a str")
        pulumi.set(__self__, "endpoint", endpoint)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip and not isinstance(ip, str):
            raise TypeError("Expected argument 'ip' to be a str")
        pulumi.set(__self__, "ip", ip)
        if kube_config and not isinstance(kube_config, str):
            raise TypeError("Expected argument 'kube_config' to be a str")
        pulumi.set(__self__, "kube_config", kube_config)
        if label and not isinstance(label, str):
            raise TypeError("Expected argument 'label' to be a str")
        pulumi.set(__self__, "label", label)
        if node_pools and not isinstance(node_pools, list):
            raise TypeError("Expected argument 'node_pools' to be a list")
        pulumi.set(__self__, "node_pools", node_pools)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if service_subnet and not isinstance(service_subnet, str):
            raise TypeError("Expected argument 'service_subnet' to be a str")
        pulumi.set(__self__, "service_subnet", service_subnet)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="clientCertificate")
    def client_certificate(self) -> str:
        """
        The base64 encoded public certificate used by clients to access the cluster.
        """
        return pulumi.get(self, "client_certificate")

    @property
    @pulumi.getter(name="clientKey")
    def client_key(self) -> str:
        """
        The base64 encoded private key used by clients to access the cluster.
        """
        return pulumi.get(self, "client_key")

    @property
    @pulumi.getter(name="clusterCaCertificate")
    def cluster_ca_certificate(self) -> str:
        """
        The base64 encoded public certificate for the cluster's certificate authority.
        """
        return pulumi.get(self, "cluster_ca_certificate")

    @property
    @pulumi.getter(name="clusterSubnet")
    def cluster_subnet(self) -> str:
        """
        IP range that your pods will run on in this cluster.
        """
        return pulumi.get(self, "cluster_subnet")

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> str:
        """
        Date node was created.
        """
        return pulumi.get(self, "date_created")

    @property
    @pulumi.getter
    def endpoint(self) -> str:
        """
        Domain for your Kubernetes clusters control plane.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetKubernetesFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        ID of node.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ip(self) -> str:
        """
        IP address of VKE cluster control plane.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter(name="kubeConfig")
    def kube_config(self) -> str:
        """
        Base64 encoded Kubeconfig for this VKE cluster.
        """
        return pulumi.get(self, "kube_config")

    @property
    @pulumi.getter
    def label(self) -> str:
        """
        Label of node.
        """
        return pulumi.get(self, "label")

    @property
    @pulumi.getter(name="nodePools")
    def node_pools(self) -> Sequence['outputs.GetKubernetesNodePoolResult']:
        """
        Contains the default node pool that was deployed.
        """
        return pulumi.get(self, "node_pools")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        The region your VKE cluster is deployed in.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="serviceSubnet")
    def service_subnet(self) -> str:
        """
        IP range that services will run on this cluster.
        """
        return pulumi.get(self, "service_subnet")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Status of node.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        The current kubernetes version your VKE cluster is running on.
        """
        return pulumi.get(self, "version")


class AwaitableGetKubernetesResult(GetKubernetesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKubernetesResult(
            client_certificate=self.client_certificate,
            client_key=self.client_key,
            cluster_ca_certificate=self.cluster_ca_certificate,
            cluster_subnet=self.cluster_subnet,
            date_created=self.date_created,
            endpoint=self.endpoint,
            filters=self.filters,
            id=self.id,
            ip=self.ip,
            kube_config=self.kube_config,
            label=self.label,
            node_pools=self.node_pools,
            region=self.region,
            service_subnet=self.service_subnet,
            status=self.status,
            version=self.version)


def get_kubernetes(filters: Optional[Sequence[pulumi.InputType['GetKubernetesFilterArgs']]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKubernetesResult:
    """
    Get information about a Vultr Kubernetes Engine (VKE) Cluster.

    ## Example Usage

    Create a new VKE cluster:

    ```python
    import pulumi
    import pulumi_vultr as vultr

    my_vke = vultr.get_kubernetes(filters=[vultr.GetKubernetesFilterArgs(
        name="label",
        values=["my-lb-label"],
    )])
    ```


    :param Sequence[pulumi.InputType['GetKubernetesFilterArgs']] filters: Query parameters for finding VKE.
    """
    __args__ = dict()
    __args__['filters'] = filters
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vultr:index/getKubernetes:getKubernetes', __args__, opts=opts, typ=GetKubernetesResult).value

    return AwaitableGetKubernetesResult(
        client_certificate=pulumi.get(__ret__, 'client_certificate'),
        client_key=pulumi.get(__ret__, 'client_key'),
        cluster_ca_certificate=pulumi.get(__ret__, 'cluster_ca_certificate'),
        cluster_subnet=pulumi.get(__ret__, 'cluster_subnet'),
        date_created=pulumi.get(__ret__, 'date_created'),
        endpoint=pulumi.get(__ret__, 'endpoint'),
        filters=pulumi.get(__ret__, 'filters'),
        id=pulumi.get(__ret__, 'id'),
        ip=pulumi.get(__ret__, 'ip'),
        kube_config=pulumi.get(__ret__, 'kube_config'),
        label=pulumi.get(__ret__, 'label'),
        node_pools=pulumi.get(__ret__, 'node_pools'),
        region=pulumi.get(__ret__, 'region'),
        service_subnet=pulumi.get(__ret__, 'service_subnet'),
        status=pulumi.get(__ret__, 'status'),
        version=pulumi.get(__ret__, 'version'))


@_utilities.lift_output_func(get_kubernetes)
def get_kubernetes_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetKubernetesFilterArgs']]]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetKubernetesResult]:
    """
    Get information about a Vultr Kubernetes Engine (VKE) Cluster.

    ## Example Usage

    Create a new VKE cluster:

    ```python
    import pulumi
    import pulumi_vultr as vultr

    my_vke = vultr.get_kubernetes(filters=[vultr.GetKubernetesFilterArgs(
        name="label",
        values=["my-lb-label"],
    )])
    ```


    :param Sequence[pulumi.InputType['GetKubernetesFilterArgs']] filters: Query parameters for finding VKE.
    """
    ...
