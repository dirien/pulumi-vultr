// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Get information about a Vultr load balancer.
 *
 * ## Example Usage
 *
 * Get the information for a load balancer by `label`:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@pulumi/vultr";
 *
 * const myLb = vultr.getLoadBalancer({
 *     filters: [{
 *         name: "label",
 *         values: ["my-lb-label"],
 *     }],
 * });
 * ```
 */
export function getLoadBalancer(args?: GetLoadBalancerArgs, opts?: pulumi.InvokeOptions): Promise<GetLoadBalancerResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vultr:index/getLoadBalancer:getLoadBalancer", {
        "filters": args.filters,
        "proxyProtocol": args.proxyProtocol,
    }, opts);
}

/**
 * A collection of arguments for invoking getLoadBalancer.
 */
export interface GetLoadBalancerArgs {
    /**
     * Query parameters for finding load balancers.
     */
    filters?: inputs.GetLoadBalancerFilter[];
    /**
     * Boolean value that indicates if Proxy Protocol is enabled.
     */
    proxyProtocol?: boolean;
}

/**
 * A collection of values returned by getLoadBalancer.
 */
export interface GetLoadBalancerResult {
    /**
     * Array of instances that are currently attached to the load balancer.
     */
    readonly attachedInstances: string[];
    /**
     * The balancing algorithm for your load balancer.
     */
    readonly balancingAlgorithm: string;
    /**
     * Name for your given sticky session.
     */
    readonly cookieName: string;
    readonly dateCreated: string;
    readonly filters?: outputs.GetLoadBalancerFilter[];
    readonly firewallRules: {[key: string]: any}[];
    /**
     * Defines the forwarding rules for a load balancer. The configuration of a `forwardingRules` is listened below.
     */
    readonly forwardingRules: {[key: string]: any}[];
    /**
     * Boolean value that indicates if SSL is enabled.
     */
    readonly hasSsl: boolean;
    /**
     * Defines the way load balancers should check for health. The configuration of a `healthCheck` is listed below.
     */
    readonly healthCheck: {[key: string]: any};
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * IPv4 address for your load balancer.
     */
    readonly ipv4: string;
    /**
     * IPv6 address for your load balancer.
     */
    readonly ipv6: string;
    /**
     * The load balancers label.
     */
    readonly label: string;
    /**
     * (Deprecated: use `vpc` instead) Defines the private network the load balancer is attached to.
     */
    readonly privateNetwork: string;
    /**
     * Boolean value that indicates if Proxy Protocol is enabled.
     */
    readonly proxyProtocol?: boolean;
    /**
     * The region your load balancer is deployed in.
     */
    readonly region: string;
    readonly ssl: {[key: string]: any};
    /**
     * Boolean value that indicates if HTTP calls will be redirected to HTTPS.
     */
    readonly sslRedirect: boolean;
    /**
     * Current status for the load balancer
     */
    readonly status: string;
}
/**
 * Get information about a Vultr load balancer.
 *
 * ## Example Usage
 *
 * Get the information for a load balancer by `label`:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@pulumi/vultr";
 *
 * const myLb = vultr.getLoadBalancer({
 *     filters: [{
 *         name: "label",
 *         values: ["my-lb-label"],
 *     }],
 * });
 * ```
 */
export function getLoadBalancerOutput(args?: GetLoadBalancerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLoadBalancerResult> {
    return pulumi.output(args).apply((a: any) => getLoadBalancer(a, opts))
}

/**
 * A collection of arguments for invoking getLoadBalancer.
 */
export interface GetLoadBalancerOutputArgs {
    /**
     * Query parameters for finding load balancers.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.GetLoadBalancerFilterArgs>[]>;
    /**
     * Boolean value that indicates if Proxy Protocol is enabled.
     */
    proxyProtocol?: pulumi.Input<boolean>;
}
