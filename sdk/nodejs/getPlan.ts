// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Get information about a Vultr plan.
 *
 * ## Example Usage
 *
 * Get the information for a plan by `id`:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@pulumi/vultr";
 *
 * const myPlan = vultr.getPlan({
 *     filters: [{
 *         name: "id",
 *         values: ["vc2-1c-2gb"],
 *     }],
 * });
 * ```
 */
export function getPlan(args?: GetPlanArgs, opts?: pulumi.InvokeOptions): Promise<GetPlanResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vultr:index/getPlan:getPlan", {
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getPlan.
 */
export interface GetPlanArgs {
    /**
     * Query parameters for finding plans.
     */
    filters?: inputs.GetPlanFilter[];
}

/**
 * A collection of values returned by getPlan.
 */
export interface GetPlanResult {
    /**
     * The bandwidth available on the plan in GB.
     */
    readonly bandwidth: number;
    /**
     * The amount of disk space in GB available on the plan.
     */
    readonly disk: number;
    /**
     * The number of disks that this plan offers.
     */
    readonly diskCount: number;
    readonly filters?: outputs.GetPlanFilter[];
    /**
     * For GPU plans, the GPU card type.
     */
    readonly gpuType: string;
    /**
     * For GPU plans, the VRAM available in the plan.
     */
    readonly gpuVram: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly locations: string[];
    /**
     * The price per month of the plan in USD.
     */
    readonly monthlyCost: number;
    /**
     * The amount of memory available on the plan in MB.
     */
    readonly ram: number;
    /**
     * The type of plan it is.
     */
    readonly type: string;
    /**
     * The number of virtual CPUs available on the plan.
     */
    readonly vcpuCount: number;
}
/**
 * Get information about a Vultr plan.
 *
 * ## Example Usage
 *
 * Get the information for a plan by `id`:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@pulumi/vultr";
 *
 * const myPlan = vultr.getPlan({
 *     filters: [{
 *         name: "id",
 *         values: ["vc2-1c-2gb"],
 *     }],
 * });
 * ```
 */
export function getPlanOutput(args?: GetPlanOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPlanResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vultr:index/getPlan:getPlan", {
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getPlan.
 */
export interface GetPlanOutputArgs {
    /**
     * Query parameters for finding plans.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.GetPlanFilterArgs>[]>;
}
