// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a Vultr private object storage resource. This can be used to create, read, update and delete object storage resources on your Vultr account.
 *
 * ## Example Usage
 *
 * Create a new object storage subscription.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@ediri/vultr";
 *
 * const tf = new vultr.ObjectStorage("tf", {
 *     clusterId: 9,
 *     label: "vultr-object-storage",
 *     tierId: 4,
 * });
 * ```
 *
 * ## Import
 *
 * Object Storage can be imported using the object storage `ID`, e.g.
 *
 * ```sh
 * $ pulumi import vultr:index/objectStorage:ObjectStorage my_s3 0e04f918-575e-41cb-86f6-d729b354a5a1
 * ```
 */
export class ObjectStorage extends pulumi.CustomResource {
    /**
     * Get an existing ObjectStorage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ObjectStorageState, opts?: pulumi.CustomResourceOptions): ObjectStorage {
        return new ObjectStorage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vultr:index/objectStorage:ObjectStorage';

    /**
     * Returns true if the given object is an instance of ObjectStorage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ObjectStorage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ObjectStorage.__pulumiType;
    }

    /**
     * The ID of the region that you want the object storage to be deployed in.
     */
    public readonly clusterId!: pulumi.Output<number>;
    /**
     * Date of creation for the object storage subscription.
     */
    public /*out*/ readonly dateCreated!: pulumi.Output<string>;
    /**
     * The description you want to give your object storage.
     */
    public readonly label!: pulumi.Output<string | undefined>;
    /**
     * The location which this subscription resides in.
     */
    public /*out*/ readonly location!: pulumi.Output<string>;
    /**
     * The region ID of the object storage subscription.
     */
    public /*out*/ readonly region!: pulumi.Output<string>;
    /**
     * Your access key.
     */
    public /*out*/ readonly s3AccessKey!: pulumi.Output<string>;
    /**
     * The hostname for this subscription.
     */
    public /*out*/ readonly s3Hostname!: pulumi.Output<string>;
    /**
     * Your secret key.
     */
    public /*out*/ readonly s3SecretKey!: pulumi.Output<string>;
    /**
     * Current status of this object storage subscription.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The ID of the tier to deploy the storage under.
     */
    public readonly tierId!: pulumi.Output<number>;

    /**
     * Create a ObjectStorage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ObjectStorageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ObjectStorageArgs | ObjectStorageState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ObjectStorageState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["dateCreated"] = state ? state.dateCreated : undefined;
            resourceInputs["label"] = state ? state.label : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["s3AccessKey"] = state ? state.s3AccessKey : undefined;
            resourceInputs["s3Hostname"] = state ? state.s3Hostname : undefined;
            resourceInputs["s3SecretKey"] = state ? state.s3SecretKey : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tierId"] = state ? state.tierId : undefined;
        } else {
            const args = argsOrState as ObjectStorageArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.tierId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tierId'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["label"] = args ? args.label : undefined;
            resourceInputs["tierId"] = args ? args.tierId : undefined;
            resourceInputs["dateCreated"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["s3AccessKey"] = undefined /*out*/;
            resourceInputs["s3Hostname"] = undefined /*out*/;
            resourceInputs["s3SecretKey"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["s3AccessKey", "s3SecretKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ObjectStorage.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ObjectStorage resources.
 */
export interface ObjectStorageState {
    /**
     * The ID of the region that you want the object storage to be deployed in.
     */
    clusterId?: pulumi.Input<number>;
    /**
     * Date of creation for the object storage subscription.
     */
    dateCreated?: pulumi.Input<string>;
    /**
     * The description you want to give your object storage.
     */
    label?: pulumi.Input<string>;
    /**
     * The location which this subscription resides in.
     */
    location?: pulumi.Input<string>;
    /**
     * The region ID of the object storage subscription.
     */
    region?: pulumi.Input<string>;
    /**
     * Your access key.
     */
    s3AccessKey?: pulumi.Input<string>;
    /**
     * The hostname for this subscription.
     */
    s3Hostname?: pulumi.Input<string>;
    /**
     * Your secret key.
     */
    s3SecretKey?: pulumi.Input<string>;
    /**
     * Current status of this object storage subscription.
     */
    status?: pulumi.Input<string>;
    /**
     * The ID of the tier to deploy the storage under.
     */
    tierId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ObjectStorage resource.
 */
export interface ObjectStorageArgs {
    /**
     * The ID of the region that you want the object storage to be deployed in.
     */
    clusterId: pulumi.Input<number>;
    /**
     * The description you want to give your object storage.
     */
    label?: pulumi.Input<string>;
    /**
     * The ID of the tier to deploy the storage under.
     */
    tierId: pulumi.Input<number>;
}
