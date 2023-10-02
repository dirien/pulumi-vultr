// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Get information about a Vultr VPC 2.0.
 *
 * ## Example Usage
 *
 * Get the information for a VPC 2.0 by `description`:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@pulumi/vultr";
 *
 * const myVpc2 = vultr.getVpc2({
 *     filters: [{
 *         name: "description",
 *         values: ["my-vpc2-description"],
 *     }],
 * });
 * ```
 */
export function getVpc2(args?: GetVpc2Args, opts?: pulumi.InvokeOptions): Promise<GetVpc2Result> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vultr:index/getVpc2:getVpc2", {
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpc2.
 */
export interface GetVpc2Args {
    /**
     * Query parameters for finding VPCs 2.0.
     */
    filters?: inputs.GetVpc2Filter[];
}

/**
 * A collection of values returned by getVpc2.
 */
export interface GetVpc2Result {
    /**
     * The date the VPC 2.0 was added to your Vultr account.
     */
    readonly dateCreated: string;
    /**
     * The VPC 2.0's description.
     */
    readonly description: string;
    readonly filters?: outputs.GetVpc2Filter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The IPv4 network address. For example: 10.1.1.0.
     */
    readonly ipBlock: string;
    /**
     * The number of bits for the netmask in CIDR notation. Example: 20
     */
    readonly prefixLength: number;
    /**
     * The ID of the region that the VPC 2.0 is in.
     */
    readonly region: string;
}
/**
 * Get information about a Vultr VPC 2.0.
 *
 * ## Example Usage
 *
 * Get the information for a VPC 2.0 by `description`:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@pulumi/vultr";
 *
 * const myVpc2 = vultr.getVpc2({
 *     filters: [{
 *         name: "description",
 *         values: ["my-vpc2-description"],
 *     }],
 * });
 * ```
 */
export function getVpc2Output(args?: GetVpc2OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVpc2Result> {
    return pulumi.output(args).apply((a: any) => getVpc2(a, opts))
}

/**
 * A collection of arguments for invoking getVpc2.
 */
export interface GetVpc2OutputArgs {
    /**
     * Query parameters for finding VPCs 2.0.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.GetVpc2FilterArgs>[]>;
}