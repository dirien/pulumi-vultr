// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Get information about Object Storage Clusters on Vultr.
 *
 * ## Example Usage
 *
 * Get the information for an object storage cluster by `region`:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@pulumi/vultr";
 *
 * const s3 = vultr.getObjectStorageCluster({
 *     filters: [{
 *         name: "region",
 *         values: ["ewr"],
 *     }],
 * });
 * ```
 */
export function getObjectStorageCluster(args?: GetObjectStorageClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetObjectStorageClusterResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vultr:index/getObjectStorageCluster:getObjectStorageCluster", {
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getObjectStorageCluster.
 */
export interface GetObjectStorageClusterArgs {
    /**
     * Query parameters for finding operating systems.
     */
    filters?: inputs.GetObjectStorageClusterFilter[];
}

/**
 * A collection of values returned by getObjectStorageCluster.
 */
export interface GetObjectStorageClusterResult {
    /**
     * The Cluster is eligible for Object Storage deployment. (yes or no)
     */
    readonly deploy: string;
    readonly filters?: outputs.GetObjectStorageClusterFilter[];
    /**
     * The cluster hostname.
     */
    readonly hostname: string;
    /**
     * The identifying cluster ID.
     */
    readonly id: number;
    /**
     * The region ID of the object storage cluster.
     */
    readonly region: string;
}
/**
 * Get information about Object Storage Clusters on Vultr.
 *
 * ## Example Usage
 *
 * Get the information for an object storage cluster by `region`:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@pulumi/vultr";
 *
 * const s3 = vultr.getObjectStorageCluster({
 *     filters: [{
 *         name: "region",
 *         values: ["ewr"],
 *     }],
 * });
 * ```
 */
export function getObjectStorageClusterOutput(args?: GetObjectStorageClusterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetObjectStorageClusterResult> {
    return pulumi.output(args).apply((a: any) => getObjectStorageCluster(a, opts))
}

/**
 * A collection of arguments for invoking getObjectStorageCluster.
 */
export interface GetObjectStorageClusterOutputArgs {
    /**
     * Query parameters for finding operating systems.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.GetObjectStorageClusterFilterArgs>[]>;
}