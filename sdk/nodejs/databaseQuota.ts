// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a Vultr database quota resource. This can be used to create, read, and delete quotas for a managed database on your Vultr account.
 *
 * ## Example Usage
 *
 * Create a new database quota:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@ediri/vultr";
 *
 * const myDatabaseQuota = new vultr.DatabaseQuota("myDatabaseQuota", {
 *     databaseId: vultr_database.my_database.id,
 *     clientId: "my_database_quota",
 *     consumerByteRate: 3,
 *     producerByteRate: 2,
 *     requestPercentage: 120,
 *     user: "my_database_user",
 * });
 * ```
 */
export class DatabaseQuota extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseQuota resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseQuotaState, opts?: pulumi.CustomResourceOptions): DatabaseQuota {
        return new DatabaseQuota(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vultr:index/databaseQuota:DatabaseQuota';

    /**
     * Returns true if the given object is an instance of DatabaseQuota.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabaseQuota {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabaseQuota.__pulumiType;
    }

    /**
     * The client ID for the new database quota.
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * The consumer byte rate for the new managed database quota.
     */
    public readonly consumerByteRate!: pulumi.Output<number>;
    /**
     * The managed database ID you want to attach this quota to.
     */
    public readonly databaseId!: pulumi.Output<string>;
    /**
     * The producer byte rate for the new managed database quota.
     */
    public readonly producerByteRate!: pulumi.Output<number>;
    /**
     * The CPU request percentage for the new managed database quota.
     */
    public readonly requestPercentage!: pulumi.Output<number>;
    /**
     * The user for the new managed database quota.
     */
    public readonly user!: pulumi.Output<string>;

    /**
     * Create a DatabaseQuota resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseQuotaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseQuotaArgs | DatabaseQuotaState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatabaseQuotaState | undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["consumerByteRate"] = state ? state.consumerByteRate : undefined;
            resourceInputs["databaseId"] = state ? state.databaseId : undefined;
            resourceInputs["producerByteRate"] = state ? state.producerByteRate : undefined;
            resourceInputs["requestPercentage"] = state ? state.requestPercentage : undefined;
            resourceInputs["user"] = state ? state.user : undefined;
        } else {
            const args = argsOrState as DatabaseQuotaArgs | undefined;
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            if ((!args || args.consumerByteRate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'consumerByteRate'");
            }
            if ((!args || args.databaseId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseId'");
            }
            if ((!args || args.producerByteRate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'producerByteRate'");
            }
            if ((!args || args.requestPercentage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'requestPercentage'");
            }
            if ((!args || args.user === undefined) && !opts.urn) {
                throw new Error("Missing required property 'user'");
            }
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["consumerByteRate"] = args ? args.consumerByteRate : undefined;
            resourceInputs["databaseId"] = args ? args.databaseId : undefined;
            resourceInputs["producerByteRate"] = args ? args.producerByteRate : undefined;
            resourceInputs["requestPercentage"] = args ? args.requestPercentage : undefined;
            resourceInputs["user"] = args ? args.user : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DatabaseQuota.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatabaseQuota resources.
 */
export interface DatabaseQuotaState {
    /**
     * The client ID for the new database quota.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The consumer byte rate for the new managed database quota.
     */
    consumerByteRate?: pulumi.Input<number>;
    /**
     * The managed database ID you want to attach this quota to.
     */
    databaseId?: pulumi.Input<string>;
    /**
     * The producer byte rate for the new managed database quota.
     */
    producerByteRate?: pulumi.Input<number>;
    /**
     * The CPU request percentage for the new managed database quota.
     */
    requestPercentage?: pulumi.Input<number>;
    /**
     * The user for the new managed database quota.
     */
    user?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatabaseQuota resource.
 */
export interface DatabaseQuotaArgs {
    /**
     * The client ID for the new database quota.
     */
    clientId: pulumi.Input<string>;
    /**
     * The consumer byte rate for the new managed database quota.
     */
    consumerByteRate: pulumi.Input<number>;
    /**
     * The managed database ID you want to attach this quota to.
     */
    databaseId: pulumi.Input<string>;
    /**
     * The producer byte rate for the new managed database quota.
     */
    producerByteRate: pulumi.Input<number>;
    /**
     * The CPU request percentage for the new managed database quota.
     */
    requestPercentage: pulumi.Input<number>;
    /**
     * The user for the new managed database quota.
     */
    user: pulumi.Input<string>;
}