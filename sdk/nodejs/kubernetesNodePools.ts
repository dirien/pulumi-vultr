// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Deploy additional node pools to an existing Vultr Kubernetes Engine (VKE) cluster.
 *
 * ## Example Usage
 *
 * Create a new VKE cluster:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@ediri/vultr";
 *
 * const np_1 = new vultr.KubernetesNodePools("np-1", {
 *     clusterId: vultr_kubernetes.k8.id,
 *     nodeQuantity: 1,
 *     plan: "vc2-4c-8gb",
 *     label: "my-label",
 *     tag: "my-tag",
 *     autoScaler: true,
 *     minNodes: 1,
 *     maxNodes: 2,
 *     labels: {
 *         "my-label": "a-label-on-all-nodes",
 *         "my-second-label": "another-label-on-all-nodes",
 *     },
 *     taints: [
 *         {
 *             key: "a-taint",
 *             value: "is-tainted",
 *             effect: "NoExecute",
 *         },
 *         {
 *             key: "another-taint",
 *             value: "is-tainted",
 *             effect: "NoSchedule",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Node pool resources are able to be imported into terraform state like other
 *
 * resources, however, since they rely on a kubernetes cluster, the import state
 *
 * requires the UUID of the cluster as well. With that in mind, format the second
 *
 * argument to the `pulumi import` command as a space delimited string of
 *
 * UUIDs, the first is the cluster ID, the second is the node pool ID. It will
 *
 * look like this:
 *
 * "clusterID nodePoolID"
 *
 * ```sh
 * $ pulumi import vultr:index/kubernetesNodePools:KubernetesNodePools my-k8s-np "7365a98b-5a43-450f-bd27-d768827100e5 ec330340-4f50-4526-858f-a39199f568ac"
 * ```
 */
export class KubernetesNodePools extends pulumi.CustomResource {
    /**
     * Get an existing KubernetesNodePools resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KubernetesNodePoolsState, opts?: pulumi.CustomResourceOptions): KubernetesNodePools {
        return new KubernetesNodePools(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vultr:index/kubernetesNodePools:KubernetesNodePools';

    /**
     * Returns true if the given object is an instance of KubernetesNodePools.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KubernetesNodePools {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KubernetesNodePools.__pulumiType;
    }

    /**
     * Enable the auto scaler for the default node pool.
     */
    public readonly autoScaler!: pulumi.Output<boolean | undefined>;
    /**
     * The VKE cluster ID you want to attach this nodepool to.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Date node was created.
     */
    public /*out*/ readonly dateCreated!: pulumi.Output<string>;
    /**
     * Date of node pool updates.
     */
    public /*out*/ readonly dateUpdated!: pulumi.Output<string>;
    /**
     * The label to be used as a prefix for nodes in this node pool.
     */
    public readonly label!: pulumi.Output<string>;
    /**
     * A map of key/value pairs for Kubernetes node labels.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The maximum number of nodes to use with the auto scaler.
     */
    public readonly maxNodes!: pulumi.Output<number | undefined>;
    /**
     * The minimum number of nodes to use with the auto scaler.
     */
    public readonly minNodes!: pulumi.Output<number | undefined>;
    /**
     * The number of nodes in this node pool.
     */
    public readonly nodeQuantity!: pulumi.Output<number>;
    /**
     * Array that contains information about nodes within this node pool.
     */
    public /*out*/ readonly nodes!: pulumi.Output<outputs.KubernetesNodePoolsNode[]>;
    /**
     * The plan to be used in this node pool. [See Plans List](https://www.vultr.com/api/#operation/list-plans) Note the minimum plan requirements must have at least 1 core and 2 gbs of memory.
     */
    public readonly plan!: pulumi.Output<string>;
    /**
     * Status of node.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A tag that is assigned to this node pool.
     */
    public readonly tag!: pulumi.Output<string | undefined>;
    /**
     * Taints to apply to the nodes in the node pool. Should contain `key`, `value` and `effect`.  The `effect` should be one of `NoSchedule`, `PreferNoSchedule` or `NoExecute`.
     */
    public readonly taints!: pulumi.Output<outputs.KubernetesNodePoolsTaint[] | undefined>;

    /**
     * Create a KubernetesNodePools resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KubernetesNodePoolsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KubernetesNodePoolsArgs | KubernetesNodePoolsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KubernetesNodePoolsState | undefined;
            resourceInputs["autoScaler"] = state ? state.autoScaler : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["dateCreated"] = state ? state.dateCreated : undefined;
            resourceInputs["dateUpdated"] = state ? state.dateUpdated : undefined;
            resourceInputs["label"] = state ? state.label : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["maxNodes"] = state ? state.maxNodes : undefined;
            resourceInputs["minNodes"] = state ? state.minNodes : undefined;
            resourceInputs["nodeQuantity"] = state ? state.nodeQuantity : undefined;
            resourceInputs["nodes"] = state ? state.nodes : undefined;
            resourceInputs["plan"] = state ? state.plan : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tag"] = state ? state.tag : undefined;
            resourceInputs["taints"] = state ? state.taints : undefined;
        } else {
            const args = argsOrState as KubernetesNodePoolsArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.label === undefined) && !opts.urn) {
                throw new Error("Missing required property 'label'");
            }
            if ((!args || args.nodeQuantity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeQuantity'");
            }
            if ((!args || args.plan === undefined) && !opts.urn) {
                throw new Error("Missing required property 'plan'");
            }
            resourceInputs["autoScaler"] = args ? args.autoScaler : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["label"] = args ? args.label : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["maxNodes"] = args ? args.maxNodes : undefined;
            resourceInputs["minNodes"] = args ? args.minNodes : undefined;
            resourceInputs["nodeQuantity"] = args ? args.nodeQuantity : undefined;
            resourceInputs["plan"] = args ? args.plan : undefined;
            resourceInputs["tag"] = args ? args.tag : undefined;
            resourceInputs["taints"] = args ? args.taints : undefined;
            resourceInputs["dateCreated"] = undefined /*out*/;
            resourceInputs["dateUpdated"] = undefined /*out*/;
            resourceInputs["nodes"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(KubernetesNodePools.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KubernetesNodePools resources.
 */
export interface KubernetesNodePoolsState {
    /**
     * Enable the auto scaler for the default node pool.
     */
    autoScaler?: pulumi.Input<boolean>;
    /**
     * The VKE cluster ID you want to attach this nodepool to.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Date node was created.
     */
    dateCreated?: pulumi.Input<string>;
    /**
     * Date of node pool updates.
     */
    dateUpdated?: pulumi.Input<string>;
    /**
     * The label to be used as a prefix for nodes in this node pool.
     */
    label?: pulumi.Input<string>;
    /**
     * A map of key/value pairs for Kubernetes node labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The maximum number of nodes to use with the auto scaler.
     */
    maxNodes?: pulumi.Input<number>;
    /**
     * The minimum number of nodes to use with the auto scaler.
     */
    minNodes?: pulumi.Input<number>;
    /**
     * The number of nodes in this node pool.
     */
    nodeQuantity?: pulumi.Input<number>;
    /**
     * Array that contains information about nodes within this node pool.
     */
    nodes?: pulumi.Input<pulumi.Input<inputs.KubernetesNodePoolsNode>[]>;
    /**
     * The plan to be used in this node pool. [See Plans List](https://www.vultr.com/api/#operation/list-plans) Note the minimum plan requirements must have at least 1 core and 2 gbs of memory.
     */
    plan?: pulumi.Input<string>;
    /**
     * Status of node.
     */
    status?: pulumi.Input<string>;
    /**
     * A tag that is assigned to this node pool.
     */
    tag?: pulumi.Input<string>;
    /**
     * Taints to apply to the nodes in the node pool. Should contain `key`, `value` and `effect`.  The `effect` should be one of `NoSchedule`, `PreferNoSchedule` or `NoExecute`.
     */
    taints?: pulumi.Input<pulumi.Input<inputs.KubernetesNodePoolsTaint>[]>;
}

/**
 * The set of arguments for constructing a KubernetesNodePools resource.
 */
export interface KubernetesNodePoolsArgs {
    /**
     * Enable the auto scaler for the default node pool.
     */
    autoScaler?: pulumi.Input<boolean>;
    /**
     * The VKE cluster ID you want to attach this nodepool to.
     */
    clusterId: pulumi.Input<string>;
    /**
     * The label to be used as a prefix for nodes in this node pool.
     */
    label: pulumi.Input<string>;
    /**
     * A map of key/value pairs for Kubernetes node labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The maximum number of nodes to use with the auto scaler.
     */
    maxNodes?: pulumi.Input<number>;
    /**
     * The minimum number of nodes to use with the auto scaler.
     */
    minNodes?: pulumi.Input<number>;
    /**
     * The number of nodes in this node pool.
     */
    nodeQuantity: pulumi.Input<number>;
    /**
     * The plan to be used in this node pool. [See Plans List](https://www.vultr.com/api/#operation/list-plans) Note the minimum plan requirements must have at least 1 core and 2 gbs of memory.
     */
    plan: pulumi.Input<string>;
    /**
     * A tag that is assigned to this node pool.
     */
    tag?: pulumi.Input<string>;
    /**
     * Taints to apply to the nodes in the node pool. Should contain `key`, `value` and `effect`.  The `effect` should be one of `NoSchedule`, `PreferNoSchedule` or `NoExecute`.
     */
    taints?: pulumi.Input<pulumi.Input<inputs.KubernetesNodePoolsTaint>[]>;
}
