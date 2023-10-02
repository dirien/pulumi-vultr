// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Get information about your Vultr account. This data source provides the balance, pending charges, last payment date, and last payment amount for your Vultr account.
 *
 * ## Example Usage
 *
 * Get the information for an account:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@pulumi/vultr";
 *
 * const myAccount = vultr.getAccount({});
 * ```
 */
export function getAccount(opts?: pulumi.InvokeOptions): Promise<GetAccountResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vultr:index/getAccount:getAccount", {
    }, opts);
}

/**
 * A collection of values returned by getAccount.
 */
export interface GetAccountResult {
    /**
     * The access control list on your Vultr account.
     */
    readonly acls: string[];
    /**
     * The current balance on your Vultr account.
     */
    readonly balance: number;
    /**
     * The email address on your Vultr account.
     */
    readonly email: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The amount of the last payment made on your Vultr account.
     */
    readonly lastPaymentAmount: number;
    /**
     * The date of the last payment made on your Vultr account.
     */
    readonly lastPaymentDate: string;
    /**
     * The name on your Vultr account.
     */
    readonly name: string;
    /**
     * The pending charges on your Vultr account.
     */
    readonly pendingCharges: number;
}
/**
 * Get information about your Vultr account. This data source provides the balance, pending charges, last payment date, and last payment amount for your Vultr account.
 *
 * ## Example Usage
 *
 * Get the information for an account:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@pulumi/vultr";
 *
 * const myAccount = vultr.getAccount({});
 * ```
 */
export function getAccountOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetAccountResult> {
    return pulumi.output(getAccount(opts))
}
