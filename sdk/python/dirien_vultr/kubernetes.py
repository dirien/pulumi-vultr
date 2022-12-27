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

__all__ = ['KubernetesArgs', 'Kubernetes']

@pulumi.input_type
class KubernetesArgs:
    def __init__(__self__, *,
                 label: pulumi.Input[str],
                 region: pulumi.Input[str],
                 version: pulumi.Input[str],
                 node_pools: Optional[pulumi.Input['KubernetesNodePoolsArgs']] = None):
        """
        The set of arguments for constructing a Kubernetes resource.
        :param pulumi.Input[str] label: The VKE clusters label.
        :param pulumi.Input[str] region: The region your VKE cluster will be deployed in. Currently, supported values are `ewr` and `lax`
        :param pulumi.Input[str] version: The version your VKE cluster you want deployed. [See Available Version](https://www.vultr.com/api/#operation/get-kubernetes-versions)
        :param pulumi.Input['KubernetesNodePoolsArgs'] node_pools: Contains the default node pool that was deployed.
        """
        pulumi.set(__self__, "label", label)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "version", version)
        if node_pools is not None:
            pulumi.set(__self__, "node_pools", node_pools)

    @property
    @pulumi.getter
    def label(self) -> pulumi.Input[str]:
        """
        The VKE clusters label.
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: pulumi.Input[str]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        The region your VKE cluster will be deployed in. Currently, supported values are `ewr` and `lax`
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input[str]:
        """
        The version your VKE cluster you want deployed. [See Available Version](https://www.vultr.com/api/#operation/get-kubernetes-versions)
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[str]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter(name="nodePools")
    def node_pools(self) -> Optional[pulumi.Input['KubernetesNodePoolsArgs']]:
        """
        Contains the default node pool that was deployed.
        """
        return pulumi.get(self, "node_pools")

    @node_pools.setter
    def node_pools(self, value: Optional[pulumi.Input['KubernetesNodePoolsArgs']]):
        pulumi.set(self, "node_pools", value)


@pulumi.input_type
class _KubernetesState:
    def __init__(__self__, *,
                 cluster_subnet: Optional[pulumi.Input[str]] = None,
                 date_created: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 kube_config: Optional[pulumi.Input[str]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 node_pools: Optional[pulumi.Input['KubernetesNodePoolsArgs']] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 service_subnet: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Kubernetes resources.
        :param pulumi.Input[str] cluster_subnet: IP range that your pods will run on in this cluster.
        :param pulumi.Input[str] date_created: Date node was created.
        :param pulumi.Input[str] endpoint: Domain for your Kubernetes clusters control plane.
        :param pulumi.Input[str] ip: IP address of VKE cluster control plane.
        :param pulumi.Input[str] kube_config: Base64 encoded Kubeconfig for this VKE cluster.
        :param pulumi.Input[str] label: The VKE clusters label.
        :param pulumi.Input['KubernetesNodePoolsArgs'] node_pools: Contains the default node pool that was deployed.
        :param pulumi.Input[str] region: The region your VKE cluster will be deployed in. Currently, supported values are `ewr` and `lax`
        :param pulumi.Input[str] service_subnet: IP range that services will run on this cluster.
        :param pulumi.Input[str] status: Status of node.
        :param pulumi.Input[str] version: The version your VKE cluster you want deployed. [See Available Version](https://www.vultr.com/api/#operation/get-kubernetes-versions)
        """
        if cluster_subnet is not None:
            pulumi.set(__self__, "cluster_subnet", cluster_subnet)
        if date_created is not None:
            pulumi.set(__self__, "date_created", date_created)
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if kube_config is not None:
            pulumi.set(__self__, "kube_config", kube_config)
        if label is not None:
            pulumi.set(__self__, "label", label)
        if node_pools is not None:
            pulumi.set(__self__, "node_pools", node_pools)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if service_subnet is not None:
            pulumi.set(__self__, "service_subnet", service_subnet)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="clusterSubnet")
    def cluster_subnet(self) -> Optional[pulumi.Input[str]]:
        """
        IP range that your pods will run on in this cluster.
        """
        return pulumi.get(self, "cluster_subnet")

    @cluster_subnet.setter
    def cluster_subnet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_subnet", value)

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> Optional[pulumi.Input[str]]:
        """
        Date node was created.
        """
        return pulumi.get(self, "date_created")

    @date_created.setter
    def date_created(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "date_created", value)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        Domain for your Kubernetes clusters control plane.
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        IP address of VKE cluster control plane.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter(name="kubeConfig")
    def kube_config(self) -> Optional[pulumi.Input[str]]:
        """
        Base64 encoded Kubeconfig for this VKE cluster.
        """
        return pulumi.get(self, "kube_config")

    @kube_config.setter
    def kube_config(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kube_config", value)

    @property
    @pulumi.getter
    def label(self) -> Optional[pulumi.Input[str]]:
        """
        The VKE clusters label.
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter(name="nodePools")
    def node_pools(self) -> Optional[pulumi.Input['KubernetesNodePoolsArgs']]:
        """
        Contains the default node pool that was deployed.
        """
        return pulumi.get(self, "node_pools")

    @node_pools.setter
    def node_pools(self, value: Optional[pulumi.Input['KubernetesNodePoolsArgs']]):
        pulumi.set(self, "node_pools", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region your VKE cluster will be deployed in. Currently, supported values are `ewr` and `lax`
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="serviceSubnet")
    def service_subnet(self) -> Optional[pulumi.Input[str]]:
        """
        IP range that services will run on this cluster.
        """
        return pulumi.get(self, "service_subnet")

    @service_subnet.setter
    def service_subnet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_subnet", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of node.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        The version your VKE cluster you want deployed. [See Available Version](https://www.vultr.com/api/#operation/get-kubernetes-versions)
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class Kubernetes(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 node_pools: Optional[pulumi.Input[pulumi.InputType['KubernetesNodePoolsArgs']]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        Create a new VKE cluster:

        ```python
        import pulumi
        import dirien_vultr as vultr

        k8 = vultr.Kubernetes("k8",
            label="tf-test",
            node_pools=vultr.KubernetesNodePoolsArgs(
                auto_scaler=True,
                label="my-label",
                max_nodes=2,
                min_nodes=1,
                node_quantity=1,
                plan="vc2-2c-4gb",
            ),
            region="ewr",
            version="v1.23.5+1")
        ```

        A default node pool is required when first creating the resource but it can be removed at a later point so long as there is a separate `KubernetesNodePools` resource attached. For example:

        ```python
        import pulumi
        import dirien_vultr as vultr

        k8 = vultr.Kubernetes("k8",
            region="ewr",
            label="tf-test",
            version="v1.23.5+1")
        # This resource must be created and attached to the cluster
        # before removing the default node from the vultr_kubernetes resource
        np = vultr.KubernetesNodePools("np",
            cluster_id=k8.id,
            node_quantity=1,
            plan="vc2-2c-4gb",
            label="my-label",
            auto_scaler=True,
            min_nodes=1,
            max_nodes=2)
        ```

        There is still a requirement that there be one node pool attached to the cluster but this should allow more flexibility about which node pool that is.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] label: The VKE clusters label.
        :param pulumi.Input[pulumi.InputType['KubernetesNodePoolsArgs']] node_pools: Contains the default node pool that was deployed.
        :param pulumi.Input[str] region: The region your VKE cluster will be deployed in. Currently, supported values are `ewr` and `lax`
        :param pulumi.Input[str] version: The version your VKE cluster you want deployed. [See Available Version](https://www.vultr.com/api/#operation/get-kubernetes-versions)
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KubernetesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        Create a new VKE cluster:

        ```python
        import pulumi
        import dirien_vultr as vultr

        k8 = vultr.Kubernetes("k8",
            label="tf-test",
            node_pools=vultr.KubernetesNodePoolsArgs(
                auto_scaler=True,
                label="my-label",
                max_nodes=2,
                min_nodes=1,
                node_quantity=1,
                plan="vc2-2c-4gb",
            ),
            region="ewr",
            version="v1.23.5+1")
        ```

        A default node pool is required when first creating the resource but it can be removed at a later point so long as there is a separate `KubernetesNodePools` resource attached. For example:

        ```python
        import pulumi
        import dirien_vultr as vultr

        k8 = vultr.Kubernetes("k8",
            region="ewr",
            label="tf-test",
            version="v1.23.5+1")
        # This resource must be created and attached to the cluster
        # before removing the default node from the vultr_kubernetes resource
        np = vultr.KubernetesNodePools("np",
            cluster_id=k8.id,
            node_quantity=1,
            plan="vc2-2c-4gb",
            label="my-label",
            auto_scaler=True,
            min_nodes=1,
            max_nodes=2)
        ```

        There is still a requirement that there be one node pool attached to the cluster but this should allow more flexibility about which node pool that is.

        :param str resource_name: The name of the resource.
        :param KubernetesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KubernetesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 node_pools: Optional[pulumi.Input[pulumi.InputType['KubernetesNodePoolsArgs']]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KubernetesArgs.__new__(KubernetesArgs)

            if label is None and not opts.urn:
                raise TypeError("Missing required property 'label'")
            __props__.__dict__["label"] = label
            __props__.__dict__["node_pools"] = node_pools
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            if version is None and not opts.urn:
                raise TypeError("Missing required property 'version'")
            __props__.__dict__["version"] = version
            __props__.__dict__["cluster_subnet"] = None
            __props__.__dict__["date_created"] = None
            __props__.__dict__["endpoint"] = None
            __props__.__dict__["ip"] = None
            __props__.__dict__["kube_config"] = None
            __props__.__dict__["service_subnet"] = None
            __props__.__dict__["status"] = None
        super(Kubernetes, __self__).__init__(
            'vultr:index/kubernetes:Kubernetes',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_subnet: Optional[pulumi.Input[str]] = None,
            date_created: Optional[pulumi.Input[str]] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            ip: Optional[pulumi.Input[str]] = None,
            kube_config: Optional[pulumi.Input[str]] = None,
            label: Optional[pulumi.Input[str]] = None,
            node_pools: Optional[pulumi.Input[pulumi.InputType['KubernetesNodePoolsArgs']]] = None,
            region: Optional[pulumi.Input[str]] = None,
            service_subnet: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'Kubernetes':
        """
        Get an existing Kubernetes resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_subnet: IP range that your pods will run on in this cluster.
        :param pulumi.Input[str] date_created: Date node was created.
        :param pulumi.Input[str] endpoint: Domain for your Kubernetes clusters control plane.
        :param pulumi.Input[str] ip: IP address of VKE cluster control plane.
        :param pulumi.Input[str] kube_config: Base64 encoded Kubeconfig for this VKE cluster.
        :param pulumi.Input[str] label: The VKE clusters label.
        :param pulumi.Input[pulumi.InputType['KubernetesNodePoolsArgs']] node_pools: Contains the default node pool that was deployed.
        :param pulumi.Input[str] region: The region your VKE cluster will be deployed in. Currently, supported values are `ewr` and `lax`
        :param pulumi.Input[str] service_subnet: IP range that services will run on this cluster.
        :param pulumi.Input[str] status: Status of node.
        :param pulumi.Input[str] version: The version your VKE cluster you want deployed. [See Available Version](https://www.vultr.com/api/#operation/get-kubernetes-versions)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _KubernetesState.__new__(_KubernetesState)

        __props__.__dict__["cluster_subnet"] = cluster_subnet
        __props__.__dict__["date_created"] = date_created
        __props__.__dict__["endpoint"] = endpoint
        __props__.__dict__["ip"] = ip
        __props__.__dict__["kube_config"] = kube_config
        __props__.__dict__["label"] = label
        __props__.__dict__["node_pools"] = node_pools
        __props__.__dict__["region"] = region
        __props__.__dict__["service_subnet"] = service_subnet
        __props__.__dict__["status"] = status
        __props__.__dict__["version"] = version
        return Kubernetes(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterSubnet")
    def cluster_subnet(self) -> pulumi.Output[str]:
        """
        IP range that your pods will run on in this cluster.
        """
        return pulumi.get(self, "cluster_subnet")

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> pulumi.Output[str]:
        """
        Date node was created.
        """
        return pulumi.get(self, "date_created")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        Domain for your Kubernetes clusters control plane.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Output[str]:
        """
        IP address of VKE cluster control plane.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter(name="kubeConfig")
    def kube_config(self) -> pulumi.Output[str]:
        """
        Base64 encoded Kubeconfig for this VKE cluster.
        """
        return pulumi.get(self, "kube_config")

    @property
    @pulumi.getter
    def label(self) -> pulumi.Output[str]:
        """
        The VKE clusters label.
        """
        return pulumi.get(self, "label")

    @property
    @pulumi.getter(name="nodePools")
    def node_pools(self) -> pulumi.Output[Optional['outputs.KubernetesNodePools']]:
        """
        Contains the default node pool that was deployed.
        """
        return pulumi.get(self, "node_pools")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region your VKE cluster will be deployed in. Currently, supported values are `ewr` and `lax`
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="serviceSubnet")
    def service_subnet(self) -> pulumi.Output[str]:
        """
        IP range that services will run on this cluster.
        """
        return pulumi.get(self, "service_subnet")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status of node.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        The version your VKE cluster you want deployed. [See Available Version](https://www.vultr.com/api/#operation/get-kubernetes-versions)
        """
        return pulumi.get(self, "version")

