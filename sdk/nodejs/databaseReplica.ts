// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class DatabaseReplica extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseReplica resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseReplicaState, opts?: pulumi.CustomResourceOptions): DatabaseReplica {
        return new DatabaseReplica(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vultr:index/databaseReplica:DatabaseReplica';

    /**
     * Returns true if the given object is an instance of DatabaseReplica.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabaseReplica {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabaseReplica.__pulumiType;
    }

    public /*out*/ readonly clusterTimeZone!: pulumi.Output<string>;
    public /*out*/ readonly databaseEngine!: pulumi.Output<string>;
    public /*out*/ readonly databaseEngineVersion!: pulumi.Output<string>;
    public readonly databaseId!: pulumi.Output<string>;
    public /*out*/ readonly dateCreated!: pulumi.Output<string>;
    public /*out*/ readonly dbname!: pulumi.Output<string>;
    public /*out*/ readonly host!: pulumi.Output<string>;
    public readonly label!: pulumi.Output<string>;
    public /*out*/ readonly latestBackup!: pulumi.Output<string>;
    public /*out*/ readonly maintenanceDow!: pulumi.Output<string>;
    public /*out*/ readonly maintenanceTime!: pulumi.Output<string>;
    public readonly mysqlLongQueryTime!: pulumi.Output<number>;
    public readonly mysqlRequirePrimaryKey!: pulumi.Output<boolean>;
    public readonly mysqlSlowQueryLog!: pulumi.Output<boolean>;
    public readonly mysqlSqlModes!: pulumi.Output<string[]>;
    public /*out*/ readonly password!: pulumi.Output<string>;
    public /*out*/ readonly plan!: pulumi.Output<string>;
    public readonly planDisk!: pulumi.Output<number>;
    public /*out*/ readonly planRam!: pulumi.Output<number>;
    public /*out*/ readonly planReplicas!: pulumi.Output<number>;
    public /*out*/ readonly planVcpus!: pulumi.Output<number>;
    public /*out*/ readonly port!: pulumi.Output<string>;
    public readonly redisEvictionPolicy!: pulumi.Output<string>;
    public readonly region!: pulumi.Output<string>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly tag!: pulumi.Output<string>;
    public readonly trustedIps!: pulumi.Output<string[]>;
    public /*out*/ readonly user!: pulumi.Output<string>;

    /**
     * Create a DatabaseReplica resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseReplicaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseReplicaArgs | DatabaseReplicaState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatabaseReplicaState | undefined;
            resourceInputs["clusterTimeZone"] = state ? state.clusterTimeZone : undefined;
            resourceInputs["databaseEngine"] = state ? state.databaseEngine : undefined;
            resourceInputs["databaseEngineVersion"] = state ? state.databaseEngineVersion : undefined;
            resourceInputs["databaseId"] = state ? state.databaseId : undefined;
            resourceInputs["dateCreated"] = state ? state.dateCreated : undefined;
            resourceInputs["dbname"] = state ? state.dbname : undefined;
            resourceInputs["host"] = state ? state.host : undefined;
            resourceInputs["label"] = state ? state.label : undefined;
            resourceInputs["latestBackup"] = state ? state.latestBackup : undefined;
            resourceInputs["maintenanceDow"] = state ? state.maintenanceDow : undefined;
            resourceInputs["maintenanceTime"] = state ? state.maintenanceTime : undefined;
            resourceInputs["mysqlLongQueryTime"] = state ? state.mysqlLongQueryTime : undefined;
            resourceInputs["mysqlRequirePrimaryKey"] = state ? state.mysqlRequirePrimaryKey : undefined;
            resourceInputs["mysqlSlowQueryLog"] = state ? state.mysqlSlowQueryLog : undefined;
            resourceInputs["mysqlSqlModes"] = state ? state.mysqlSqlModes : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["plan"] = state ? state.plan : undefined;
            resourceInputs["planDisk"] = state ? state.planDisk : undefined;
            resourceInputs["planRam"] = state ? state.planRam : undefined;
            resourceInputs["planReplicas"] = state ? state.planReplicas : undefined;
            resourceInputs["planVcpus"] = state ? state.planVcpus : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["redisEvictionPolicy"] = state ? state.redisEvictionPolicy : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tag"] = state ? state.tag : undefined;
            resourceInputs["trustedIps"] = state ? state.trustedIps : undefined;
            resourceInputs["user"] = state ? state.user : undefined;
        } else {
            const args = argsOrState as DatabaseReplicaArgs | undefined;
            if ((!args || args.databaseId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseId'");
            }
            if ((!args || args.label === undefined) && !opts.urn) {
                throw new Error("Missing required property 'label'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["databaseId"] = args ? args.databaseId : undefined;
            resourceInputs["label"] = args ? args.label : undefined;
            resourceInputs["mysqlLongQueryTime"] = args ? args.mysqlLongQueryTime : undefined;
            resourceInputs["mysqlRequirePrimaryKey"] = args ? args.mysqlRequirePrimaryKey : undefined;
            resourceInputs["mysqlSlowQueryLog"] = args ? args.mysqlSlowQueryLog : undefined;
            resourceInputs["mysqlSqlModes"] = args ? args.mysqlSqlModes : undefined;
            resourceInputs["planDisk"] = args ? args.planDisk : undefined;
            resourceInputs["redisEvictionPolicy"] = args ? args.redisEvictionPolicy : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tag"] = args ? args.tag : undefined;
            resourceInputs["trustedIps"] = args ? args.trustedIps : undefined;
            resourceInputs["clusterTimeZone"] = undefined /*out*/;
            resourceInputs["databaseEngine"] = undefined /*out*/;
            resourceInputs["databaseEngineVersion"] = undefined /*out*/;
            resourceInputs["dateCreated"] = undefined /*out*/;
            resourceInputs["dbname"] = undefined /*out*/;
            resourceInputs["host"] = undefined /*out*/;
            resourceInputs["latestBackup"] = undefined /*out*/;
            resourceInputs["maintenanceDow"] = undefined /*out*/;
            resourceInputs["maintenanceTime"] = undefined /*out*/;
            resourceInputs["password"] = undefined /*out*/;
            resourceInputs["plan"] = undefined /*out*/;
            resourceInputs["planRam"] = undefined /*out*/;
            resourceInputs["planReplicas"] = undefined /*out*/;
            resourceInputs["planVcpus"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["user"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DatabaseReplica.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatabaseReplica resources.
 */
export interface DatabaseReplicaState {
    clusterTimeZone?: pulumi.Input<string>;
    databaseEngine?: pulumi.Input<string>;
    databaseEngineVersion?: pulumi.Input<string>;
    databaseId?: pulumi.Input<string>;
    dateCreated?: pulumi.Input<string>;
    dbname?: pulumi.Input<string>;
    host?: pulumi.Input<string>;
    label?: pulumi.Input<string>;
    latestBackup?: pulumi.Input<string>;
    maintenanceDow?: pulumi.Input<string>;
    maintenanceTime?: pulumi.Input<string>;
    mysqlLongQueryTime?: pulumi.Input<number>;
    mysqlRequirePrimaryKey?: pulumi.Input<boolean>;
    mysqlSlowQueryLog?: pulumi.Input<boolean>;
    mysqlSqlModes?: pulumi.Input<pulumi.Input<string>[]>;
    password?: pulumi.Input<string>;
    plan?: pulumi.Input<string>;
    planDisk?: pulumi.Input<number>;
    planRam?: pulumi.Input<number>;
    planReplicas?: pulumi.Input<number>;
    planVcpus?: pulumi.Input<number>;
    port?: pulumi.Input<string>;
    redisEvictionPolicy?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    tag?: pulumi.Input<string>;
    trustedIps?: pulumi.Input<pulumi.Input<string>[]>;
    user?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatabaseReplica resource.
 */
export interface DatabaseReplicaArgs {
    databaseId: pulumi.Input<string>;
    label: pulumi.Input<string>;
    mysqlLongQueryTime?: pulumi.Input<number>;
    mysqlRequirePrimaryKey?: pulumi.Input<boolean>;
    mysqlSlowQueryLog?: pulumi.Input<boolean>;
    mysqlSqlModes?: pulumi.Input<pulumi.Input<string>[]>;
    planDisk?: pulumi.Input<number>;
    redisEvictionPolicy?: pulumi.Input<string>;
    region: pulumi.Input<string>;
    tag?: pulumi.Input<string>;
    trustedIps?: pulumi.Input<pulumi.Input<string>[]>;
}