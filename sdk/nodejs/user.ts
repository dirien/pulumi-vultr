// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a Vultr User resource. This can be used to create, read, modify, and delete Users.
 *
 * ## Example Usage
 *
 * Create a new User without any ACLs
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@ediri/vultr";
 *
 * const myUser = new vultr.User("myUser", {
 *     apiEnabled: true,
 *     email: "user@vultr.com",
 *     password: "myP@ssw0rd",
 * });
 * ```
 *
 * Create a new User with all ACLs
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@ediri/vultr";
 *
 * const myUser = new vultr.User("myUser", {
 *     acls: [
 *         "manage_users",
 *         "subscriptions",
 *         "provisioning",
 *         "billing",
 *         "support",
 *         "abuse",
 *         "dns",
 *         "upgrade",
 *     ],
 *     apiEnabled: true,
 *     email: "user@vultr.com",
 *     password: "myP@ssw0rd",
 * });
 * ```
 *
 * ## Import
 *
 * Users can be imported using the User `ID`, e.g.
 *
 * ```sh
 * $ pulumi import vultr:index/user:User myuser 1345fef0-8ed3-4a66-bd8c-822a7b7bd05a
 * ```
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vultr:index/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * The access control list for the user.
     */
    public readonly acls!: pulumi.Output<string[] | undefined>;
    /**
     * Whether API is enabled for the user. Default behavior is set to enabled.
     */
    public readonly apiEnabled!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly apiKey!: pulumi.Output<string>;
    /**
     * Email for this user.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * Name for this user.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Password for this user.
     */
    public readonly password!: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserState | undefined;
            resourceInputs["acls"] = state ? state.acls : undefined;
            resourceInputs["apiEnabled"] = state ? state.apiEnabled : undefined;
            resourceInputs["apiKey"] = state ? state.apiKey : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            resourceInputs["acls"] = args ? args.acls : undefined;
            resourceInputs["apiEnabled"] = args ? args.apiEnabled : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["apiKey"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(User.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * The access control list for the user.
     */
    acls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether API is enabled for the user. Default behavior is set to enabled.
     */
    apiEnabled?: pulumi.Input<boolean>;
    apiKey?: pulumi.Input<string>;
    /**
     * Email for this user.
     */
    email?: pulumi.Input<string>;
    /**
     * Name for this user.
     */
    name?: pulumi.Input<string>;
    /**
     * Password for this user.
     */
    password?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * The access control list for the user.
     */
    acls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether API is enabled for the user. Default behavior is set to enabled.
     */
    apiEnabled?: pulumi.Input<boolean>;
    /**
     * Email for this user.
     */
    email: pulumi.Input<string>;
    /**
     * Name for this user.
     */
    name?: pulumi.Input<string>;
    /**
     * Password for this user.
     */
    password: pulumi.Input<string>;
}
