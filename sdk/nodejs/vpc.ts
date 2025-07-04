// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a Vultr VPC resource. This can be used to create, read, and delete VPCs on your Vultr account.
 *
 * ## Example Usage
 *
 * Create a new VPC with an automatically generated CIDR block:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@ediri/vultr";
 *
 * const myVpc = new vultr.Vpc("myVpc", {
 *     description: "my vpc",
 *     region: "ewr",
 * });
 * ```
 *
 * Create a new VPC with a specified CIDR block:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@ediri/vultr";
 *
 * const myVpc = new vultr.Vpc("myVpc", {
 *     description: "my private vpc",
 *     region: "ewr",
 *     v4Subnet: "10.0.0.0",
 *     v4SubnetMask: 24,
 * });
 * ```
 *
 * ## Import
 *
 * VPCs can be imported using the VPC `ID`, e.g.
 *
 * ```sh
 * $ pulumi import vultr:index/vpc:Vpc my_vpc 0e04f918-575e-41cb-86f6-d729b354a5a1
 * ```
 */
export class Vpc extends pulumi.CustomResource {
    /**
     * Get an existing Vpc resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcState, opts?: pulumi.CustomResourceOptions): Vpc {
        return new Vpc(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vultr:index/vpc:Vpc';

    /**
     * Returns true if the given object is an instance of Vpc.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Vpc {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Vpc.__pulumiType;
    }

    /**
     * The date that the VPC was added to your Vultr account.
     */
    public /*out*/ readonly dateCreated!: pulumi.Output<string>;
    /**
     * The description you want to give your VPC.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The region ID that you want the VPC to be created in.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The IPv4 subnet to be used when attaching instances to this VPC.
     */
    public readonly v4Subnet!: pulumi.Output<string>;
    /**
     * The number of bits for the netmask in CIDR notation. Example: 32
     */
    public readonly v4SubnetMask!: pulumi.Output<number>;

    /**
     * Create a Vpc resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcArgs | VpcState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcState | undefined;
            resourceInputs["dateCreated"] = state ? state.dateCreated : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["v4Subnet"] = state ? state.v4Subnet : undefined;
            resourceInputs["v4SubnetMask"] = state ? state.v4SubnetMask : undefined;
        } else {
            const args = argsOrState as VpcArgs | undefined;
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
        super(Vpc.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Vpc resources.
 */
export interface VpcState {
    /**
     * The date that the VPC was added to your Vultr account.
     */
    dateCreated?: pulumi.Input<string>;
    /**
     * The description you want to give your VPC.
     */
    description?: pulumi.Input<string>;
    /**
     * The region ID that you want the VPC to be created in.
     */
    region?: pulumi.Input<string>;
    /**
     * The IPv4 subnet to be used when attaching instances to this VPC.
     */
    v4Subnet?: pulumi.Input<string>;
    /**
     * The number of bits for the netmask in CIDR notation. Example: 32
     */
    v4SubnetMask?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Vpc resource.
 */
export interface VpcArgs {
    /**
     * The description you want to give your VPC.
     */
    description?: pulumi.Input<string>;
    /**
     * The region ID that you want the VPC to be created in.
     */
    region: pulumi.Input<string>;
    /**
     * The IPv4 subnet to be used when attaching instances to this VPC.
     */
    v4Subnet?: pulumi.Input<string>;
    /**
     * The number of bits for the netmask in CIDR notation. Example: 32
     */
    v4SubnetMask?: pulumi.Input<number>;
}
