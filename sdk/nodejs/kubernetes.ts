// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * Create a new VKE cluster:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@ediri/vultr";
 *
 * const k8 = new vultr.Kubernetes("k8", {
 *     label: "vke-test",
 *     nodePools: {
 *         autoScaler: true,
 *         label: "vke-nodepool",
 *         labels: {
 *             "my-label": "a-label-on-all-nodes",
 *             "my-second-label": "another-label-on-all-nodes",
 *         },
 *         maxNodes: 2,
 *         minNodes: 1,
 *         nodeQuantity: 1,
 *         plan: "vc2-1c-2gb",
 *         taints: [{
 *             effect: "NoExecute",
 *             key: "a-taint",
 *             value: "is-tainted",
 *         }],
 *     },
 *     region: "ewr",
 *     version: "v1.28.2+1",
 * });
 * ```
 *
 * A default node pool is required when first creating the resource but it can be removed at a later point so long as there is a separate `vultr.KubernetesNodePools` resource attached. For example:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@ediri/vultr";
 *
 * const k8 = new vultr.Kubernetes("k8", {
 *     region: "ewr",
 *     label: "vke-test",
 *     version: "v1.28.2+1",
 * });
 * // This resource must be created and attached to the cluster
 * // before removing the default node from the vultr_kubernetes resource
 * const np = new vultr.KubernetesNodePools("np", {
 *     clusterId: k8.id,
 *     nodeQuantity: 1,
 *     plan: "vc2-1c-2gb",
 *     label: "vke-nodepool",
 *     autoScaler: true,
 *     minNodes: 1,
 *     maxNodes: 2,
 * });
 * ```
 *
 * There is still a requirement that there be one node pool attached to the cluster but this should allow more flexibility about which node pool that is.
 *
 * ## Import
 *
 * A kubernetes cluster created outside of terraform can be imported into the
 *
 * terraform state using the UUID.  One thing to note is that all kubernetes
 *
 * resources have a default node pool with a tag of `tf-vke-default`. In order to
 *
 * avoid errors, ensure that there is a node pool with that tag set.
 *
 * ```sh
 * $ pulumi import vultr:index/kubernetes:Kubernetes my-k8s 7365a98b-5a43-450f-bd27-d768827100e5
 * ```
 */
export class Kubernetes extends pulumi.CustomResource {
    /**
     * Get an existing Kubernetes resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KubernetesState, opts?: pulumi.CustomResourceOptions): Kubernetes {
        return new Kubernetes(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vultr:index/kubernetes:Kubernetes';

    /**
     * Returns true if the given object is an instance of Kubernetes.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Kubernetes {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Kubernetes.__pulumiType;
    }

    /**
     * The base64 encoded public certificate used by clients to access the cluster.
     */
    public /*out*/ readonly clientCertificate!: pulumi.Output<string>;
    /**
     * The base64 encoded private key used by clients to access the cluster.
     */
    public /*out*/ readonly clientKey!: pulumi.Output<string>;
    /**
     * The base64 encoded public certificate for the cluster's certificate authority.
     */
    public /*out*/ readonly clusterCaCertificate!: pulumi.Output<string>;
    /**
     * IP range that your pods will run on in this cluster.
     */
    public /*out*/ readonly clusterSubnet!: pulumi.Output<string>;
    /**
     * Date node was created.
     */
    public /*out*/ readonly dateCreated!: pulumi.Output<string>;
    /**
     * Boolean indicating if the cluster should be created with a managed firewall.
     */
    public readonly enableFirewall!: pulumi.Output<boolean | undefined>;
    /**
     * Domain for your Kubernetes clusters control plane.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * The ID of the firewall group managed by this cluster.
     */
    public /*out*/ readonly firewallGroupId!: pulumi.Output<string>;
    /**
     * Boolean indicating if the cluster should be created with multiple, highly available controlplanes.
     */
    public readonly haControlplanes!: pulumi.Output<boolean | undefined>;
    /**
     * IP address of VKE cluster control plane.
     */
    public /*out*/ readonly ip!: pulumi.Output<string>;
    /**
     * Base64 encoded Kubeconfig for this VKE cluster.
     */
    public /*out*/ readonly kubeConfig!: pulumi.Output<string>;
    /**
     * The VKE clusters label.
     */
    public readonly label!: pulumi.Output<string>;
    /**
     * Contains the default node pool that was deployed.
     */
    public readonly nodePools!: pulumi.Output<outputs.KubernetesNodePools | undefined>;
    /**
     * The region your VKE cluster will be deployed in.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * IP range that services will run on this cluster.
     */
    public /*out*/ readonly serviceSubnet!: pulumi.Output<string>;
    /**
     * Status of node.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The version your VKE cluster you want deployed. [See Available Version](https://www.vultr.com/api/#operation/get-kubernetes-versions)
     */
    public readonly version!: pulumi.Output<string>;
    /**
     * The ID of the VPC to use when creating the cluster. If not provided a new VPC will be created instead.
     */
    public readonly vpcId!: pulumi.Output<string | undefined>;

    /**
     * Create a Kubernetes resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KubernetesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KubernetesArgs | KubernetesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KubernetesState | undefined;
            resourceInputs["clientCertificate"] = state ? state.clientCertificate : undefined;
            resourceInputs["clientKey"] = state ? state.clientKey : undefined;
            resourceInputs["clusterCaCertificate"] = state ? state.clusterCaCertificate : undefined;
            resourceInputs["clusterSubnet"] = state ? state.clusterSubnet : undefined;
            resourceInputs["dateCreated"] = state ? state.dateCreated : undefined;
            resourceInputs["enableFirewall"] = state ? state.enableFirewall : undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["firewallGroupId"] = state ? state.firewallGroupId : undefined;
            resourceInputs["haControlplanes"] = state ? state.haControlplanes : undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["kubeConfig"] = state ? state.kubeConfig : undefined;
            resourceInputs["label"] = state ? state.label : undefined;
            resourceInputs["nodePools"] = state ? state.nodePools : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["serviceSubnet"] = state ? state.serviceSubnet : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as KubernetesArgs | undefined;
            if ((!args || args.label === undefined) && !opts.urn) {
                throw new Error("Missing required property 'label'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            resourceInputs["enableFirewall"] = args ? args.enableFirewall : undefined;
            resourceInputs["haControlplanes"] = args ? args.haControlplanes : undefined;
            resourceInputs["label"] = args ? args.label : undefined;
            resourceInputs["nodePools"] = args ? args.nodePools : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["clientCertificate"] = undefined /*out*/;
            resourceInputs["clientKey"] = undefined /*out*/;
            resourceInputs["clusterCaCertificate"] = undefined /*out*/;
            resourceInputs["clusterSubnet"] = undefined /*out*/;
            resourceInputs["dateCreated"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
            resourceInputs["firewallGroupId"] = undefined /*out*/;
            resourceInputs["ip"] = undefined /*out*/;
            resourceInputs["kubeConfig"] = undefined /*out*/;
            resourceInputs["serviceSubnet"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["clientCertificate", "clientKey", "clusterCaCertificate", "kubeConfig"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Kubernetes.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Kubernetes resources.
 */
export interface KubernetesState {
    /**
     * The base64 encoded public certificate used by clients to access the cluster.
     */
    clientCertificate?: pulumi.Input<string>;
    /**
     * The base64 encoded private key used by clients to access the cluster.
     */
    clientKey?: pulumi.Input<string>;
    /**
     * The base64 encoded public certificate for the cluster's certificate authority.
     */
    clusterCaCertificate?: pulumi.Input<string>;
    /**
     * IP range that your pods will run on in this cluster.
     */
    clusterSubnet?: pulumi.Input<string>;
    /**
     * Date node was created.
     */
    dateCreated?: pulumi.Input<string>;
    /**
     * Boolean indicating if the cluster should be created with a managed firewall.
     */
    enableFirewall?: pulumi.Input<boolean>;
    /**
     * Domain for your Kubernetes clusters control plane.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * The ID of the firewall group managed by this cluster.
     */
    firewallGroupId?: pulumi.Input<string>;
    /**
     * Boolean indicating if the cluster should be created with multiple, highly available controlplanes.
     */
    haControlplanes?: pulumi.Input<boolean>;
    /**
     * IP address of VKE cluster control plane.
     */
    ip?: pulumi.Input<string>;
    /**
     * Base64 encoded Kubeconfig for this VKE cluster.
     */
    kubeConfig?: pulumi.Input<string>;
    /**
     * The VKE clusters label.
     */
    label?: pulumi.Input<string>;
    /**
     * Contains the default node pool that was deployed.
     */
    nodePools?: pulumi.Input<inputs.KubernetesNodePools>;
    /**
     * The region your VKE cluster will be deployed in.
     */
    region?: pulumi.Input<string>;
    /**
     * IP range that services will run on this cluster.
     */
    serviceSubnet?: pulumi.Input<string>;
    /**
     * Status of node.
     */
    status?: pulumi.Input<string>;
    /**
     * The version your VKE cluster you want deployed. [See Available Version](https://www.vultr.com/api/#operation/get-kubernetes-versions)
     */
    version?: pulumi.Input<string>;
    /**
     * The ID of the VPC to use when creating the cluster. If not provided a new VPC will be created instead.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Kubernetes resource.
 */
export interface KubernetesArgs {
    /**
     * Boolean indicating if the cluster should be created with a managed firewall.
     */
    enableFirewall?: pulumi.Input<boolean>;
    /**
     * Boolean indicating if the cluster should be created with multiple, highly available controlplanes.
     */
    haControlplanes?: pulumi.Input<boolean>;
    /**
     * The VKE clusters label.
     */
    label: pulumi.Input<string>;
    /**
     * Contains the default node pool that was deployed.
     */
    nodePools?: pulumi.Input<inputs.KubernetesNodePools>;
    /**
     * The region your VKE cluster will be deployed in.
     */
    region: pulumi.Input<string>;
    /**
     * The version your VKE cluster you want deployed. [See Available Version](https://www.vultr.com/api/#operation/get-kubernetes-versions)
     */
    version: pulumi.Input<string>;
    /**
     * The ID of the VPC to use when creating the cluster. If not provided a new VPC will be created instead.
     */
    vpcId?: pulumi.Input<string>;
}
