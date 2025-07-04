// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class IsoPrivate extends pulumi.CustomResource {
    /**
     * Get an existing IsoPrivate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IsoPrivateState, opts?: pulumi.CustomResourceOptions): IsoPrivate {
        return new IsoPrivate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vultr:index/isoPrivate:IsoPrivate';

    /**
     * Returns true if the given object is an instance of IsoPrivate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IsoPrivate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IsoPrivate.__pulumiType;
    }

    public /*out*/ readonly dateCreated!: pulumi.Output<string>;
    public /*out*/ readonly filename!: pulumi.Output<string>;
    public /*out*/ readonly md5sum!: pulumi.Output<string>;
    public /*out*/ readonly sha512sum!: pulumi.Output<string>;
    public /*out*/ readonly size!: pulumi.Output<number>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a IsoPrivate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IsoPrivateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IsoPrivateArgs | IsoPrivateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IsoPrivateState | undefined;
            resourceInputs["dateCreated"] = state ? state.dateCreated : undefined;
            resourceInputs["filename"] = state ? state.filename : undefined;
            resourceInputs["md5sum"] = state ? state.md5sum : undefined;
            resourceInputs["sha512sum"] = state ? state.sha512sum : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as IsoPrivateArgs | undefined;
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["dateCreated"] = undefined /*out*/;
            resourceInputs["filename"] = undefined /*out*/;
            resourceInputs["md5sum"] = undefined /*out*/;
            resourceInputs["sha512sum"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IsoPrivate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IsoPrivate resources.
 */
export interface IsoPrivateState {
    dateCreated?: pulumi.Input<string>;
    filename?: pulumi.Input<string>;
    md5sum?: pulumi.Input<string>;
    sha512sum?: pulumi.Input<string>;
    size?: pulumi.Input<number>;
    status?: pulumi.Input<string>;
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IsoPrivate resource.
 */
export interface IsoPrivateArgs {
    url: pulumi.Input<string>;
}
