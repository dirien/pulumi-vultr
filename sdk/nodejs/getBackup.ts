// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Get information about a Vultr backup. This data source provides a list of backups which contain the description, size, status, and the creation date for your Vultr backup.
 *
 * ## Example Usage
 *
 * Get the information for a backup by `description`:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@ediri/vultr";
 *
 * const myBackup = vultr.getBackup({
 *     filters: [{
 *         name: "description",
 *         values: ["my-backup-description"],
 *     }],
 * });
 * ```
 */
export function getBackup(args?: GetBackupArgs, opts?: pulumi.InvokeOptions): Promise<GetBackupResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vultr:index/getBackup:getBackup", {
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getBackup.
 */
export interface GetBackupArgs {
    /**
     * Query parameters for finding backups.
     */
    filters?: inputs.GetBackupFilter[];
}

/**
 * A collection of values returned by getBackup.
 */
export interface GetBackupResult {
    readonly backups: {[key: string]: string}[];
    readonly filters?: outputs.GetBackupFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
/**
 * Get information about a Vultr backup. This data source provides a list of backups which contain the description, size, status, and the creation date for your Vultr backup.
 *
 * ## Example Usage
 *
 * Get the information for a backup by `description`:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@ediri/vultr";
 *
 * const myBackup = vultr.getBackup({
 *     filters: [{
 *         name: "description",
 *         values: ["my-backup-description"],
 *     }],
 * });
 * ```
 */
export function getBackupOutput(args?: GetBackupOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetBackupResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vultr:index/getBackup:getBackup", {
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getBackup.
 */
export interface GetBackupOutputArgs {
    /**
     * Query parameters for finding backups.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.GetBackupFilterArgs>[]>;
}
