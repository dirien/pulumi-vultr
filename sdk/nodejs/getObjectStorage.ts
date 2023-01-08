// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Get information about an Object Storage subscription on Vultr.
 *
 * ## Example Usage
 *
 * Get the information for an object storage subscription by `label`:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@pulumi/vultr";
 *
 * const s3 = vultr.getObjectStorage({
 *     filters: [{
 *         name: "label",
 *         values: ["my-s3"],
 *     }],
 * });
 * ```
 */
export function getObjectStorage(args?: GetObjectStorageArgs, opts?: pulumi.InvokeOptions): Promise<GetObjectStorageResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vultr:index/getObjectStorage:getObjectStorage", {
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getObjectStorage.
 */
export interface GetObjectStorageArgs {
    /**
     * Query parameters for finding operating systems.
     */
    filters?: inputs.GetObjectStorageFilter[];
}

/**
 * A collection of values returned by getObjectStorage.
 */
export interface GetObjectStorageResult {
    /**
     * The identifying cluster ID.
     */
    readonly clusterId: number;
    /**
     * Date of creation for the object storage subscription.
     */
    readonly dateCreated: string;
    readonly filters?: outputs.GetObjectStorageFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The label of the object storage subscription.
     */
    readonly label: string;
    /**
     * The location which this subscription resides in.
     */
    readonly location: string;
    /**
     * The region ID of the object storage subscription.
     */
    readonly region: string;
    /**
     * Your access key.
     */
    readonly s3AccessKey: string;
    /**
     * The hostname for this subscription.
     */
    readonly s3Hostname: string;
    /**
     * Your secret key.
     */
    readonly s3SecretKey: string;
    /**
     * Current status of this object storage subscription.
     */
    readonly status: string;
}
/**
 * Get information about an Object Storage subscription on Vultr.
 *
 * ## Example Usage
 *
 * Get the information for an object storage subscription by `label`:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@pulumi/vultr";
 *
 * const s3 = vultr.getObjectStorage({
 *     filters: [{
 *         name: "label",
 *         values: ["my-s3"],
 *     }],
 * });
 * ```
 */
export function getObjectStorageOutput(args?: GetObjectStorageOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetObjectStorageResult> {
    return pulumi.output(args).apply((a: any) => getObjectStorage(a, opts))
}

/**
 * A collection of arguments for invoking getObjectStorage.
 */
export interface GetObjectStorageOutputArgs {
    /**
     * Query parameters for finding operating systems.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.GetObjectStorageFilterArgs>[]>;
}