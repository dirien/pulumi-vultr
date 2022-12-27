// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Get information about a Vultr Reverse IPv4.
 *
 * ## Example Usage
 *
 * Get the information for an IPv4 reverse DNS record by `reverse`:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@pulumi/vultr";
 *
 * const myReverseIpv4 = vultr.getReverseIpv4({
 *     filters: [{
 *         name: "reverse",
 *         values: ["host.example.com"],
 *     }],
 * });
 * ```
 */
export function getReverseIpv4(args?: GetReverseIpv4Args, opts?: pulumi.InvokeOptions): Promise<GetReverseIpv4Result> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vultr:index/getReverseIpv4:getReverseIpv4", {
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getReverseIpv4.
 */
export interface GetReverseIpv4Args {
    /**
     * Query parameters for finding IPv4 reverse DNS records.
     */
    filters?: inputs.GetReverseIpv4Filter[];
}

/**
 * A collection of values returned by getReverseIpv4.
 */
export interface GetReverseIpv4Result {
    readonly filters?: outputs.GetReverseIpv4Filter[];
    /**
     * The gateway IP address.
     */
    readonly gateway: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The ID of the instance the IPv4 reverse DNS record was set for.
     */
    readonly instanceId: string;
    /**
     * The IPv4 address in canonical format used in the reverse DNS record.
     */
    readonly ip: string;
    /**
     * The IPv4 netmask in dot-decimal notation.
     */
    readonly netmask: string;
    /**
     * The hostname used in the IPv4 reverse DNS record.
     */
    readonly reverse: string;
}
/**
 * Get information about a Vultr Reverse IPv4.
 *
 * ## Example Usage
 *
 * Get the information for an IPv4 reverse DNS record by `reverse`:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@pulumi/vultr";
 *
 * const myReverseIpv4 = vultr.getReverseIpv4({
 *     filters: [{
 *         name: "reverse",
 *         values: ["host.example.com"],
 *     }],
 * });
 * ```
 */
export function getReverseIpv4Output(args?: GetReverseIpv4OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetReverseIpv4Result> {
    return pulumi.output(args).apply((a: any) => getReverseIpv4(a, opts))
}

/**
 * A collection of arguments for invoking getReverseIpv4.
 */
export interface GetReverseIpv4OutputArgs {
    /**
     * Query parameters for finding IPv4 reverse DNS records.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.GetReverseIpv4FilterArgs>[]>;
}
