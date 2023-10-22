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

__all__ = ['KubernetesNodePoolsInitArgs', 'KubernetesNodePools']

@pulumi.input_type
class KubernetesNodePoolsInitArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 label: pulumi.Input[str],
                 node_quantity: pulumi.Input[int],
                 plan: pulumi.Input[str],
                 auto_scaler: Optional[pulumi.Input[bool]] = None,
                 max_nodes: Optional[pulumi.Input[int]] = None,
                 min_nodes: Optional[pulumi.Input[int]] = None,
                 tag: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a KubernetesNodePools resource.
        :param pulumi.Input[str] cluster_id: The VKE cluster ID you want to attach this nodepool to.
        :param pulumi.Input[str] label: The label to be used as a prefix for nodes in this node pool.
        :param pulumi.Input[int] node_quantity: The number of nodes in this node pool.
        :param pulumi.Input[str] plan: The plan to be used in this node pool. [See Plans List](https://www.vultr.com/api/#operation/list-plans) Note the minimum plan requirements must have at least 1 core and 2 gbs of memory.
        :param pulumi.Input[bool] auto_scaler: Enable the auto scaler for the default node pool.
        :param pulumi.Input[int] max_nodes: The maximum number of nodes to use with the auto scaler.
        :param pulumi.Input[int] min_nodes: The minimum number of nodes to use with the auto scaler.
        :param pulumi.Input[str] tag: A tag that is assigned to this node pool.
        """
        KubernetesNodePoolsInitArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            cluster_id=cluster_id,
            label=label,
            node_quantity=node_quantity,
            plan=plan,
            auto_scaler=auto_scaler,
            max_nodes=max_nodes,
            min_nodes=min_nodes,
            tag=tag,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             cluster_id: pulumi.Input[str],
             label: pulumi.Input[str],
             node_quantity: pulumi.Input[int],
             plan: pulumi.Input[str],
             auto_scaler: Optional[pulumi.Input[bool]] = None,
             max_nodes: Optional[pulumi.Input[int]] = None,
             min_nodes: Optional[pulumi.Input[int]] = None,
             tag: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'clusterId' in kwargs:
            cluster_id = kwargs['clusterId']
        if 'nodeQuantity' in kwargs:
            node_quantity = kwargs['nodeQuantity']
        if 'autoScaler' in kwargs:
            auto_scaler = kwargs['autoScaler']
        if 'maxNodes' in kwargs:
            max_nodes = kwargs['maxNodes']
        if 'minNodes' in kwargs:
            min_nodes = kwargs['minNodes']

        _setter("cluster_id", cluster_id)
        _setter("label", label)
        _setter("node_quantity", node_quantity)
        _setter("plan", plan)
        if auto_scaler is not None:
            _setter("auto_scaler", auto_scaler)
        if max_nodes is not None:
            _setter("max_nodes", max_nodes)
        if min_nodes is not None:
            _setter("min_nodes", min_nodes)
        if tag is not None:
            _setter("tag", tag)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        The VKE cluster ID you want to attach this nodepool to.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def label(self) -> pulumi.Input[str]:
        """
        The label to be used as a prefix for nodes in this node pool.
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: pulumi.Input[str]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter(name="nodeQuantity")
    def node_quantity(self) -> pulumi.Input[int]:
        """
        The number of nodes in this node pool.
        """
        return pulumi.get(self, "node_quantity")

    @node_quantity.setter
    def node_quantity(self, value: pulumi.Input[int]):
        pulumi.set(self, "node_quantity", value)

    @property
    @pulumi.getter
    def plan(self) -> pulumi.Input[str]:
        """
        The plan to be used in this node pool. [See Plans List](https://www.vultr.com/api/#operation/list-plans) Note the minimum plan requirements must have at least 1 core and 2 gbs of memory.
        """
        return pulumi.get(self, "plan")

    @plan.setter
    def plan(self, value: pulumi.Input[str]):
        pulumi.set(self, "plan", value)

    @property
    @pulumi.getter(name="autoScaler")
    def auto_scaler(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable the auto scaler for the default node pool.
        """
        return pulumi.get(self, "auto_scaler")

    @auto_scaler.setter
    def auto_scaler(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_scaler", value)

    @property
    @pulumi.getter(name="maxNodes")
    def max_nodes(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum number of nodes to use with the auto scaler.
        """
        return pulumi.get(self, "max_nodes")

    @max_nodes.setter
    def max_nodes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_nodes", value)

    @property
    @pulumi.getter(name="minNodes")
    def min_nodes(self) -> Optional[pulumi.Input[int]]:
        """
        The minimum number of nodes to use with the auto scaler.
        """
        return pulumi.get(self, "min_nodes")

    @min_nodes.setter
    def min_nodes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_nodes", value)

    @property
    @pulumi.getter
    def tag(self) -> Optional[pulumi.Input[str]]:
        """
        A tag that is assigned to this node pool.
        """
        return pulumi.get(self, "tag")

    @tag.setter
    def tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag", value)


@pulumi.input_type
class _KubernetesNodePoolsState:
    def __init__(__self__, *,
                 auto_scaler: Optional[pulumi.Input[bool]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 date_created: Optional[pulumi.Input[str]] = None,
                 date_updated: Optional[pulumi.Input[str]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 max_nodes: Optional[pulumi.Input[int]] = None,
                 min_nodes: Optional[pulumi.Input[int]] = None,
                 node_quantity: Optional[pulumi.Input[int]] = None,
                 nodes: Optional[pulumi.Input[Sequence[pulumi.Input['KubernetesNodePoolsNodeArgs']]]] = None,
                 plan: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tag: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering KubernetesNodePools resources.
        :param pulumi.Input[bool] auto_scaler: Enable the auto scaler for the default node pool.
        :param pulumi.Input[str] cluster_id: The VKE cluster ID you want to attach this nodepool to.
        :param pulumi.Input[str] date_created: Date node was created.
        :param pulumi.Input[str] date_updated: Date of node pool updates.
        :param pulumi.Input[str] label: The label to be used as a prefix for nodes in this node pool.
        :param pulumi.Input[int] max_nodes: The maximum number of nodes to use with the auto scaler.
        :param pulumi.Input[int] min_nodes: The minimum number of nodes to use with the auto scaler.
        :param pulumi.Input[int] node_quantity: The number of nodes in this node pool.
        :param pulumi.Input[Sequence[pulumi.Input['KubernetesNodePoolsNodeArgs']]] nodes: Array that contains information about nodes within this node pool.
        :param pulumi.Input[str] plan: The plan to be used in this node pool. [See Plans List](https://www.vultr.com/api/#operation/list-plans) Note the minimum plan requirements must have at least 1 core and 2 gbs of memory.
        :param pulumi.Input[str] status: Status of node.
        :param pulumi.Input[str] tag: A tag that is assigned to this node pool.
        """
        _KubernetesNodePoolsState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            auto_scaler=auto_scaler,
            cluster_id=cluster_id,
            date_created=date_created,
            date_updated=date_updated,
            label=label,
            max_nodes=max_nodes,
            min_nodes=min_nodes,
            node_quantity=node_quantity,
            nodes=nodes,
            plan=plan,
            status=status,
            tag=tag,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             auto_scaler: Optional[pulumi.Input[bool]] = None,
             cluster_id: Optional[pulumi.Input[str]] = None,
             date_created: Optional[pulumi.Input[str]] = None,
             date_updated: Optional[pulumi.Input[str]] = None,
             label: Optional[pulumi.Input[str]] = None,
             max_nodes: Optional[pulumi.Input[int]] = None,
             min_nodes: Optional[pulumi.Input[int]] = None,
             node_quantity: Optional[pulumi.Input[int]] = None,
             nodes: Optional[pulumi.Input[Sequence[pulumi.Input['KubernetesNodePoolsNodeArgs']]]] = None,
             plan: Optional[pulumi.Input[str]] = None,
             status: Optional[pulumi.Input[str]] = None,
             tag: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'autoScaler' in kwargs:
            auto_scaler = kwargs['autoScaler']
        if 'clusterId' in kwargs:
            cluster_id = kwargs['clusterId']
        if 'dateCreated' in kwargs:
            date_created = kwargs['dateCreated']
        if 'dateUpdated' in kwargs:
            date_updated = kwargs['dateUpdated']
        if 'maxNodes' in kwargs:
            max_nodes = kwargs['maxNodes']
        if 'minNodes' in kwargs:
            min_nodes = kwargs['minNodes']
        if 'nodeQuantity' in kwargs:
            node_quantity = kwargs['nodeQuantity']

        if auto_scaler is not None:
            _setter("auto_scaler", auto_scaler)
        if cluster_id is not None:
            _setter("cluster_id", cluster_id)
        if date_created is not None:
            _setter("date_created", date_created)
        if date_updated is not None:
            _setter("date_updated", date_updated)
        if label is not None:
            _setter("label", label)
        if max_nodes is not None:
            _setter("max_nodes", max_nodes)
        if min_nodes is not None:
            _setter("min_nodes", min_nodes)
        if node_quantity is not None:
            _setter("node_quantity", node_quantity)
        if nodes is not None:
            _setter("nodes", nodes)
        if plan is not None:
            _setter("plan", plan)
        if status is not None:
            _setter("status", status)
        if tag is not None:
            _setter("tag", tag)

    @property
    @pulumi.getter(name="autoScaler")
    def auto_scaler(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable the auto scaler for the default node pool.
        """
        return pulumi.get(self, "auto_scaler")

    @auto_scaler.setter
    def auto_scaler(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_scaler", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        The VKE cluster ID you want to attach this nodepool to.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

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
    @pulumi.getter(name="dateUpdated")
    def date_updated(self) -> Optional[pulumi.Input[str]]:
        """
        Date of node pool updates.
        """
        return pulumi.get(self, "date_updated")

    @date_updated.setter
    def date_updated(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "date_updated", value)

    @property
    @pulumi.getter
    def label(self) -> Optional[pulumi.Input[str]]:
        """
        The label to be used as a prefix for nodes in this node pool.
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter(name="maxNodes")
    def max_nodes(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum number of nodes to use with the auto scaler.
        """
        return pulumi.get(self, "max_nodes")

    @max_nodes.setter
    def max_nodes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_nodes", value)

    @property
    @pulumi.getter(name="minNodes")
    def min_nodes(self) -> Optional[pulumi.Input[int]]:
        """
        The minimum number of nodes to use with the auto scaler.
        """
        return pulumi.get(self, "min_nodes")

    @min_nodes.setter
    def min_nodes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_nodes", value)

    @property
    @pulumi.getter(name="nodeQuantity")
    def node_quantity(self) -> Optional[pulumi.Input[int]]:
        """
        The number of nodes in this node pool.
        """
        return pulumi.get(self, "node_quantity")

    @node_quantity.setter
    def node_quantity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "node_quantity", value)

    @property
    @pulumi.getter
    def nodes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['KubernetesNodePoolsNodeArgs']]]]:
        """
        Array that contains information about nodes within this node pool.
        """
        return pulumi.get(self, "nodes")

    @nodes.setter
    def nodes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['KubernetesNodePoolsNodeArgs']]]]):
        pulumi.set(self, "nodes", value)

    @property
    @pulumi.getter
    def plan(self) -> Optional[pulumi.Input[str]]:
        """
        The plan to be used in this node pool. [See Plans List](https://www.vultr.com/api/#operation/list-plans) Note the minimum plan requirements must have at least 1 core and 2 gbs of memory.
        """
        return pulumi.get(self, "plan")

    @plan.setter
    def plan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plan", value)

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
    def tag(self) -> Optional[pulumi.Input[str]]:
        """
        A tag that is assigned to this node pool.
        """
        return pulumi.get(self, "tag")

    @tag.setter
    def tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag", value)


class KubernetesNodePools(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_scaler: Optional[pulumi.Input[bool]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 max_nodes: Optional[pulumi.Input[int]] = None,
                 min_nodes: Optional[pulumi.Input[int]] = None,
                 node_quantity: Optional[pulumi.Input[int]] = None,
                 plan: Optional[pulumi.Input[str]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Deploy additional node pools to an existing Vultr Kubernetes Engine (VKE) cluster.

        ## Example Usage

        Create a new VKE cluster:

        ```python
        import pulumi
        import ediri_vultr as vultr

        np_1 = vultr.KubernetesNodePools("np-1",
            cluster_id=vultr_kubernetes["k8"]["id"],
            node_quantity=1,
            plan="vc2-2c-4gb",
            label="my-label",
            tag="my-tag",
            auto_scaler=True,
            min_nodes=1,
            max_nodes=2)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_scaler: Enable the auto scaler for the default node pool.
        :param pulumi.Input[str] cluster_id: The VKE cluster ID you want to attach this nodepool to.
        :param pulumi.Input[str] label: The label to be used as a prefix for nodes in this node pool.
        :param pulumi.Input[int] max_nodes: The maximum number of nodes to use with the auto scaler.
        :param pulumi.Input[int] min_nodes: The minimum number of nodes to use with the auto scaler.
        :param pulumi.Input[int] node_quantity: The number of nodes in this node pool.
        :param pulumi.Input[str] plan: The plan to be used in this node pool. [See Plans List](https://www.vultr.com/api/#operation/list-plans) Note the minimum plan requirements must have at least 1 core and 2 gbs of memory.
        :param pulumi.Input[str] tag: A tag that is assigned to this node pool.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KubernetesNodePoolsInitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Deploy additional node pools to an existing Vultr Kubernetes Engine (VKE) cluster.

        ## Example Usage

        Create a new VKE cluster:

        ```python
        import pulumi
        import ediri_vultr as vultr

        np_1 = vultr.KubernetesNodePools("np-1",
            cluster_id=vultr_kubernetes["k8"]["id"],
            node_quantity=1,
            plan="vc2-2c-4gb",
            label="my-label",
            tag="my-tag",
            auto_scaler=True,
            min_nodes=1,
            max_nodes=2)
        ```

        :param str resource_name: The name of the resource.
        :param KubernetesNodePoolsInitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KubernetesNodePoolsInitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            KubernetesNodePoolsInitArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_scaler: Optional[pulumi.Input[bool]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 max_nodes: Optional[pulumi.Input[int]] = None,
                 min_nodes: Optional[pulumi.Input[int]] = None,
                 node_quantity: Optional[pulumi.Input[int]] = None,
                 plan: Optional[pulumi.Input[str]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KubernetesNodePoolsInitArgs.__new__(KubernetesNodePoolsInitArgs)

            __props__.__dict__["auto_scaler"] = auto_scaler
            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            if label is None and not opts.urn:
                raise TypeError("Missing required property 'label'")
            __props__.__dict__["label"] = label
            __props__.__dict__["max_nodes"] = max_nodes
            __props__.__dict__["min_nodes"] = min_nodes
            if node_quantity is None and not opts.urn:
                raise TypeError("Missing required property 'node_quantity'")
            __props__.__dict__["node_quantity"] = node_quantity
            if plan is None and not opts.urn:
                raise TypeError("Missing required property 'plan'")
            __props__.__dict__["plan"] = plan
            __props__.__dict__["tag"] = tag
            __props__.__dict__["date_created"] = None
            __props__.__dict__["date_updated"] = None
            __props__.__dict__["nodes"] = None
            __props__.__dict__["status"] = None
        super(KubernetesNodePools, __self__).__init__(
            'vultr:index/kubernetesNodePools:KubernetesNodePools',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_scaler: Optional[pulumi.Input[bool]] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            date_created: Optional[pulumi.Input[str]] = None,
            date_updated: Optional[pulumi.Input[str]] = None,
            label: Optional[pulumi.Input[str]] = None,
            max_nodes: Optional[pulumi.Input[int]] = None,
            min_nodes: Optional[pulumi.Input[int]] = None,
            node_quantity: Optional[pulumi.Input[int]] = None,
            nodes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KubernetesNodePoolsNodeArgs']]]]] = None,
            plan: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tag: Optional[pulumi.Input[str]] = None) -> 'KubernetesNodePools':
        """
        Get an existing KubernetesNodePools resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_scaler: Enable the auto scaler for the default node pool.
        :param pulumi.Input[str] cluster_id: The VKE cluster ID you want to attach this nodepool to.
        :param pulumi.Input[str] date_created: Date node was created.
        :param pulumi.Input[str] date_updated: Date of node pool updates.
        :param pulumi.Input[str] label: The label to be used as a prefix for nodes in this node pool.
        :param pulumi.Input[int] max_nodes: The maximum number of nodes to use with the auto scaler.
        :param pulumi.Input[int] min_nodes: The minimum number of nodes to use with the auto scaler.
        :param pulumi.Input[int] node_quantity: The number of nodes in this node pool.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KubernetesNodePoolsNodeArgs']]]] nodes: Array that contains information about nodes within this node pool.
        :param pulumi.Input[str] plan: The plan to be used in this node pool. [See Plans List](https://www.vultr.com/api/#operation/list-plans) Note the minimum plan requirements must have at least 1 core and 2 gbs of memory.
        :param pulumi.Input[str] status: Status of node.
        :param pulumi.Input[str] tag: A tag that is assigned to this node pool.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _KubernetesNodePoolsState.__new__(_KubernetesNodePoolsState)

        __props__.__dict__["auto_scaler"] = auto_scaler
        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["date_created"] = date_created
        __props__.__dict__["date_updated"] = date_updated
        __props__.__dict__["label"] = label
        __props__.__dict__["max_nodes"] = max_nodes
        __props__.__dict__["min_nodes"] = min_nodes
        __props__.__dict__["node_quantity"] = node_quantity
        __props__.__dict__["nodes"] = nodes
        __props__.__dict__["plan"] = plan
        __props__.__dict__["status"] = status
        __props__.__dict__["tag"] = tag
        return KubernetesNodePools(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoScaler")
    def auto_scaler(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable the auto scaler for the default node pool.
        """
        return pulumi.get(self, "auto_scaler")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        The VKE cluster ID you want to attach this nodepool to.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> pulumi.Output[str]:
        """
        Date node was created.
        """
        return pulumi.get(self, "date_created")

    @property
    @pulumi.getter(name="dateUpdated")
    def date_updated(self) -> pulumi.Output[str]:
        """
        Date of node pool updates.
        """
        return pulumi.get(self, "date_updated")

    @property
    @pulumi.getter
    def label(self) -> pulumi.Output[str]:
        """
        The label to be used as a prefix for nodes in this node pool.
        """
        return pulumi.get(self, "label")

    @property
    @pulumi.getter(name="maxNodes")
    def max_nodes(self) -> pulumi.Output[Optional[int]]:
        """
        The maximum number of nodes to use with the auto scaler.
        """
        return pulumi.get(self, "max_nodes")

    @property
    @pulumi.getter(name="minNodes")
    def min_nodes(self) -> pulumi.Output[Optional[int]]:
        """
        The minimum number of nodes to use with the auto scaler.
        """
        return pulumi.get(self, "min_nodes")

    @property
    @pulumi.getter(name="nodeQuantity")
    def node_quantity(self) -> pulumi.Output[int]:
        """
        The number of nodes in this node pool.
        """
        return pulumi.get(self, "node_quantity")

    @property
    @pulumi.getter
    def nodes(self) -> pulumi.Output[Sequence['outputs.KubernetesNodePoolsNode']]:
        """
        Array that contains information about nodes within this node pool.
        """
        return pulumi.get(self, "nodes")

    @property
    @pulumi.getter
    def plan(self) -> pulumi.Output[str]:
        """
        The plan to be used in this node pool. [See Plans List](https://www.vultr.com/api/#operation/list-plans) Note the minimum plan requirements must have at least 1 core and 2 gbs of memory.
        """
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status of node.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tag(self) -> pulumi.Output[Optional[str]]:
        """
        A tag that is assigned to this node pool.
        """
        return pulumi.get(self, "tag")

