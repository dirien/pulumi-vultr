// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Get information about a Vultr instance.
 *
 * ## Example Usage
 *
 * Get the information for a instance by `label`:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@pulumi/vultr";
 *
 * const myInstance = vultr.getInstance({
 *     filters: [{
 *         name: "label",
 *         values: ["my-instance-label"],
 *     }],
 * });
 * ```
 */
export function getInstance(args?: GetInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vultr:index/getInstance:getInstance", {
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstance.
 */
export interface GetInstanceArgs {
    /**
     * Query parameters for finding instances.
     */
    filters?: inputs.GetInstanceFilter[];
}

/**
 * A collection of values returned by getInstance.
 */
export interface GetInstanceResult {
    /**
     * The server's allowed bandwidth usage in GB.
     */
    readonly allowedBandwidth: number;
    /**
     * The server's application ID.
     */
    readonly appId: number;
    readonly backups: string;
    /**
     * The current configuration for backups
     */
    readonly backupsSchedule: {[key: string]: string};
    /**
     * The date the server was added to your Vultr account.
     */
    readonly dateCreated: string;
    /**
     * The description of the disk(s) on the server.
     */
    readonly disk: number;
    /**
     * Array of which features are enabled.
     */
    readonly features: string[];
    readonly filters?: outputs.GetInstanceFilter[];
    /**
     * The ID of the firewall group applied to this server.
     */
    readonly firewallGroupId: string;
    /**
     * The server's IPv4 gateway.
     */
    readonly gatewayV4: string;
    /**
     * The hostname assigned to the server.
     */
    readonly hostname: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The Marketplace ID for this application.
     */
    readonly imageId: string;
    /**
     * The server's internal IP address.
     */
    readonly internalIp: string;
    /**
     * The server's current KVM URL. This URL will change periodically. It is not advised to cache this value.
     */
    readonly kvm: string;
    /**
     * The server's label.
     */
    readonly label: string;
    readonly location: string;
    /**
     * The server's main IP address.
     */
    readonly mainIp: string;
    /**
     * The server's IPv4 netmask.
     */
    readonly netmaskV4: string;
    /**
     * The operating system of the instance.
     */
    readonly os: string;
    /**
     * The server's operating system ID.
     */
    readonly osId: number;
    /**
     * The server's plan ID.
     */
    readonly plan: string;
    /**
     * Whether the server is powered on or not.
     */
    readonly powerStatus: string;
    /**
     * The amount of memory available on the instance in MB.
     */
    readonly ram: number;
    /**
     * The region ID of the server.
     */
    readonly region: string;
    /**
     * A more detailed server status (none, locked, installingbooting, isomounting, ok).
     */
    readonly serverStatus: string;
    /**
     * The status of the server's subscription.
     */
    readonly status: string;
    /**
     * A list of tags applied to the instance.
     */
    readonly tags: string[];
    /**
     * The scheme used for the default user (linux servers only).
     */
    readonly userScheme: string;
    /**
     * The main IPv6 network address.
     */
    readonly v6MainIp: string;
    /**
     * The IPv6 subnet.
     */
    readonly v6Network: string;
    /**
     * The IPv6 network size in bits.
     */
    readonly v6NetworkSize: number;
    /**
     * The number of virtual CPUs available on the server.
     */
    readonly vcpuCount: number;
    /**
     * A list of VPC 2.0 IDs attached to the server.
     */
    readonly vpc2Ids: string[];
    readonly vpcIds: string[];
}
/**
 * Get information about a Vultr instance.
 *
 * ## Example Usage
 *
 * Get the information for a instance by `label`:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@pulumi/vultr";
 *
 * const myInstance = vultr.getInstance({
 *     filters: [{
 *         name: "label",
 *         values: ["my-instance-label"],
 *     }],
 * });
 * ```
 */
export function getInstanceOutput(args?: GetInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vultr:index/getInstance:getInstance", {
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstance.
 */
export interface GetInstanceOutputArgs {
    /**
     * Query parameters for finding instances.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.GetInstanceFilterArgs>[]>;
}
