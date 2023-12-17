// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides a Vultr database user resource. This can be used to create, read, modify, and delete users for a managed database on your Vultr account.
 *
 * ## Example Usage
 *
 * Create a new database user:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@ediri/vultr";
 *
 * const myDatabaseUser = new vultr.DatabaseUser("myDatabaseUser", {
 *     databaseId: vultr_database.my_database.id,
 *     username: "my_database_user",
 *     password: "randomTestPW40298",
 * });
 * ```
 */
export class DatabaseUser extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseUserState, opts?: pulumi.CustomResourceOptions): DatabaseUser {
        return new DatabaseUser(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vultr:index/databaseUser:DatabaseUser';

    /**
     * Returns true if the given object is an instance of DatabaseUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabaseUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabaseUser.__pulumiType;
    }

    public readonly accessControl!: pulumi.Output<outputs.DatabaseUserAccessControl>;
    /**
     * The managed database ID you want to attach this user to.
     */
    public readonly databaseId!: pulumi.Output<string>;
    /**
     * The encryption type of the new managed database user's password (MySQL engine types only - `cachingSha2Password`, `mysqlNativePassword`).
     */
    public readonly encryption!: pulumi.Output<string>;
    /**
     * The password of the new managed database user.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * The username of the new managed database user.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a DatabaseUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseUserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseUserArgs | DatabaseUserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatabaseUserState | undefined;
            resourceInputs["accessControl"] = state ? state.accessControl : undefined;
            resourceInputs["databaseId"] = state ? state.databaseId : undefined;
            resourceInputs["encryption"] = state ? state.encryption : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as DatabaseUserArgs | undefined;
            if ((!args || args.databaseId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseId'");
            }
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["accessControl"] = args ? args.accessControl : undefined;
            resourceInputs["databaseId"] = args ? args.databaseId : undefined;
            resourceInputs["encryption"] = args ? args.encryption : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DatabaseUser.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatabaseUser resources.
 */
export interface DatabaseUserState {
    accessControl?: pulumi.Input<inputs.DatabaseUserAccessControl>;
    /**
     * The managed database ID you want to attach this user to.
     */
    databaseId?: pulumi.Input<string>;
    /**
     * The encryption type of the new managed database user's password (MySQL engine types only - `cachingSha2Password`, `mysqlNativePassword`).
     */
    encryption?: pulumi.Input<string>;
    /**
     * The password of the new managed database user.
     */
    password?: pulumi.Input<string>;
    /**
     * The username of the new managed database user.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatabaseUser resource.
 */
export interface DatabaseUserArgs {
    accessControl?: pulumi.Input<inputs.DatabaseUserAccessControl>;
    /**
     * The managed database ID you want to attach this user to.
     */
    databaseId: pulumi.Input<string>;
    /**
     * The encryption type of the new managed database user's password (MySQL engine types only - `cachingSha2Password`, `mysqlNativePassword`).
     */
    encryption?: pulumi.Input<string>;
    /**
     * The password of the new managed database user.
     */
    password?: pulumi.Input<string>;
    /**
     * The username of the new managed database user.
     */
    username: pulumi.Input<string>;
}
