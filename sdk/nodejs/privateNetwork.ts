// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Deprecated: Use `vultr.Vpc` instead
 *
 * Provides a Vultr private network resource. This can be used to create, read, and delete private networks on your Vultr account.
 *
 * ## Example Usage
 *
 * Create a new private network with an automatically generated CIDR block:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@ediri/vultr";
 *
 * const myNetwork = new vultr.PrivateNetwork("myNetwork", {
 *     description: "my private network",
 *     region: "ewr",
 * });
 * ```
 *
 * Create a new private network with a specified CIDR block:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@ediri/vultr";
 *
 * const myNetwork = new vultr.PrivateNetwork("myNetwork", {
 *     description: "my private network",
 *     region: "ewr",
 *     v4Subnet: "10.0.0.0",
 *     v4SubnetMask: 24,
 * });
 * ```
 *
 * ## Import
 *
 * Networks can be imported using the network `ID`, e.g.
 *
 * ```sh
 *  $ pulumi import vultr:index/privateNetwork:PrivateNetwork my_network 0e04f918-575e-41cb-86f6-d729b354a5a1
 * ```
 */
export class PrivateNetwork extends pulumi.CustomResource {
    /**
     * Get an existing PrivateNetwork resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PrivateNetworkState, opts?: pulumi.CustomResourceOptions): PrivateNetwork {
        return new PrivateNetwork(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vultr:index/privateNetwork:PrivateNetwork';

    /**
     * Returns true if the given object is an instance of PrivateNetwork.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivateNetwork {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivateNetwork.__pulumiType;
    }

    /**
     * The date that the network was added to your Vultr account.
     */
    public /*out*/ readonly dateCreated!: pulumi.Output<string>;
    /**
     * The description you want to give your network.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The region ID that you want the network to be created in.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The IPv4 subnet to be used when attaching instances to this network.
     */
    public readonly v4Subnet!: pulumi.Output<string>;
    /**
     * The number of bits for the netmask in CIDR notation. Example: 32
     */
    public readonly v4SubnetMask!: pulumi.Output<number>;

    /**
     * Create a PrivateNetwork resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivateNetworkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrivateNetworkArgs | PrivateNetworkState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PrivateNetworkState | undefined;
            resourceInputs["dateCreated"] = state ? state.dateCreated : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["v4Subnet"] = state ? state.v4Subnet : undefined;
            resourceInputs["v4SubnetMask"] = state ? state.v4SubnetMask : undefined;
        } else {
            const args = argsOrState as PrivateNetworkArgs | undefined;
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["v4Subnet"] = args ? args.v4Subnet : undefined;
            resourceInputs["v4SubnetMask"] = args ? args.v4SubnetMask : undefined;
            resourceInputs["dateCreated"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PrivateNetwork.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PrivateNetwork resources.
 */
export interface PrivateNetworkState {
    /**
     * The date that the network was added to your Vultr account.
     */
    dateCreated?: pulumi.Input<string>;
    /**
     * The description you want to give your network.
     */
    description?: pulumi.Input<string>;
    /**
     * The region ID that you want the network to be created in.
     */
    region?: pulumi.Input<string>;
    /**
     * The IPv4 subnet to be used when attaching instances to this network.
     */
    v4Subnet?: pulumi.Input<string>;
    /**
     * The number of bits for the netmask in CIDR notation. Example: 32
     */
    v4SubnetMask?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a PrivateNetwork resource.
 */
export interface PrivateNetworkArgs {
    /**
     * The description you want to give your network.
     */
    description?: pulumi.Input<string>;
    /**
     * The region ID that you want the network to be created in.
     */
    region: pulumi.Input<string>;
    /**
     * The IPv4 subnet to be used when attaching instances to this network.
     */
    v4Subnet?: pulumi.Input<string>;
    /**
     * The number of bits for the netmask in CIDR notation. Example: 32
     */
    v4SubnetMask?: pulumi.Input<number>;
}
