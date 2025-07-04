// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Get information about a Vultr SSH key. This data source provides the name, public SSH key, and the creation date for your Vultr SSH key.
 *
 * ## Example Usage
 *
 * Get the information for an SSH key by `name`:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@ediri/vultr";
 *
 * const mySshKey = vultr.getSshKey({
 *     filters: [{
 *         name: "name",
 *         values: ["my-ssh-key-name"],
 *     }],
 * });
 * ```
 */
export function getSshKey(args?: GetSshKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetSshKeyResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vultr:index/getSshKey:getSshKey", {
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getSshKey.
 */
export interface GetSshKeyArgs {
    /**
     * Query parameters for finding SSH keys.
     */
    filters?: inputs.GetSshKeyFilter[];
}

/**
 * A collection of values returned by getSshKey.
 */
export interface GetSshKeyResult {
    /**
     * The date the SSH key was added to your Vultr account.
     */
    readonly dateCreated: string;
    readonly filters?: outputs.GetSshKeyFilter[];
    /**
     * The ID of the SSH key.
     */
    readonly id: string;
    /**
     * The name of the SSH key.
     */
    readonly name: string;
    /**
     * The public SSH key.
     */
    readonly sshKey: string;
}
/**
 * Get information about a Vultr SSH key. This data source provides the name, public SSH key, and the creation date for your Vultr SSH key.
 *
 * ## Example Usage
 *
 * Get the information for an SSH key by `name`:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@ediri/vultr";
 *
 * const mySshKey = vultr.getSshKey({
 *     filters: [{
 *         name: "name",
 *         values: ["my-ssh-key-name"],
 *     }],
 * });
 * ```
 */
export function getSshKeyOutput(args?: GetSshKeyOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetSshKeyResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vultr:index/getSshKey:getSshKey", {
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getSshKey.
 */
export interface GetSshKeyOutputArgs {
    /**
     * Query parameters for finding SSH keys.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.GetSshKeyFilterArgs>[]>;
}
