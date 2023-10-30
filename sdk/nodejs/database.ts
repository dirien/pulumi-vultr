// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides a Vultr database resource. This can be used to create, read, modify, and delete managed databases on your Vultr account.
 *
 * ## Example Usage
 *
 * Create a new database:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@ediri/vultr";
 *
 * const myDatabase = new vultr.Database("myDatabase", {
 *     databaseEngine: "pg",
 *     databaseEngineVersion: "15",
 *     label: "my_database_label",
 *     plan: "vultr-dbaas-startup-cc-1-55-2",
 *     region: "ewr",
 * });
 * ```
 *
 * Create a new database with options:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@ediri/vultr";
 *
 * const myDatabase = new vultr.Database("myDatabase", {
 *     clusterTimeZone: "America/New_York",
 *     databaseEngine: "pg",
 *     databaseEngineVersion: "15",
 *     label: "my_database_label",
 *     maintenanceDow: "sunday",
 *     maintenanceTime: "01:00",
 *     plan: "vultr-dbaas-startup-cc-1-55-2",
 *     region: "ewr",
 *     tag: "some tag",
 * });
 * ```
 *
 * ## Import
 *
 * Database can be imported using the database `ID`, e.g.
 *
 * ```sh
 *  $ pulumi import vultr:index/database:Database my_database b6a859c5-b299-49dd-8888-b1abbc517d08
 * ```
 */
export class Database extends pulumi.CustomResource {
    /**
     * Get an existing Database resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseState, opts?: pulumi.CustomResourceOptions): Database {
        return new Database(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vultr:index/database:Database';

    /**
     * Returns true if the given object is an instance of Database.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Database {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Database.__pulumiType;
    }

    /**
     * The configured time zone for the Managed Database in TZ database format (e.g. `UTC`, `America/New_York`, `Europe/London`).
     */
    public readonly clusterTimeZone!: pulumi.Output<string>;
    /**
     * The database engine of the new managed database.
     */
    public readonly databaseEngine!: pulumi.Output<string>;
    /**
     * The database engine version of the new managed database.
     */
    public readonly databaseEngineVersion!: pulumi.Output<string>;
    /**
     * The date the managed database was added to your Vultr account.
     */
    public /*out*/ readonly dateCreated!: pulumi.Output<string>;
    /**
     * The managed database's default logical database.
     */
    public /*out*/ readonly dbname!: pulumi.Output<string>;
    /**
     * The hostname assigned to the managed database.
     */
    public /*out*/ readonly host!: pulumi.Output<string>;
    /**
     * A label for the managed database.
     */
    public readonly label!: pulumi.Output<string>;
    /**
     * The date of the latest backup available on the managed database.
     */
    public /*out*/ readonly latestBackup!: pulumi.Output<string>;
    /**
     * The preferred maintenance day of week for the managed database.
     */
    public readonly maintenanceDow!: pulumi.Output<string>;
    /**
     * The preferred maintenance time for the managed database in 24-hour HH:00 format (e.g. `01:00`, `13:00`, `23:00`).
     */
    public readonly maintenanceTime!: pulumi.Output<string>;
    /**
     * The configuration value for the long query time (in seconds) on the managed database (MySQL engine types only).
     */
    public readonly mysqlLongQueryTime!: pulumi.Output<number | undefined>;
    /**
     * The configuration value for whether primary keys are required on the managed database (MySQL engine types only).
     */
    public readonly mysqlRequirePrimaryKey!: pulumi.Output<boolean | undefined>;
    /**
     * The configuration value for slow query logging on the managed database (MySQL engine types only).
     */
    public readonly mysqlSlowQueryLog!: pulumi.Output<boolean | undefined>;
    /**
     * A list of SQL modes to configure for the managed database (MySQL engine types only - `ALLOW_INVALID_DATES`, `ANSI`, `ANSI_QUOTES`, `ERROR_FOR_DIVISION_BY_ZERO`, `HIGH_NOT_PRECEDENCE`, `IGNORE_SPACE`, `NO_AUTO_VALUE_ON_ZERO`, `NO_DIR_IN_CREATE`, `NO_ENGINE_SUBSTITUTION`, `NO_UNSIGNED_SUBTRACTION`, `NO_ZERO_DATE`, `NO_ZERO_IN_DATE`, `ONLY_FULL_GROUP_BY`, `PIPES_AS_CONCAT`, `REAL_AS_FLOAT`, `STRICT_ALL_TABLES`, `STRICT_TRANS_TABLES`, `TIME_TRUNCATE_FRACTIONAL`, `TRADITIONAL`).
     */
    public readonly mysqlSqlModes!: pulumi.Output<string[] | undefined>;
    /**
     * The password for the managed database's primary admin user.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * The ID of the plan that you want the managed database to subscribe to. [See List Managed Database Plans](https://www.vultr.com/api/#tag/managed-databases/operation/list-database-plans)
     */
    public readonly plan!: pulumi.Output<string>;
    /**
     * The description of the disk(s) on the managed database.
     */
    public readonly planDisk!: pulumi.Output<number>;
    /**
     * The amount of memory available on the managed database in MB.
     */
    public /*out*/ readonly planRam!: pulumi.Output<number>;
    /**
     * The number of standby nodes available on the managed database.
     */
    public /*out*/ readonly planReplicas!: pulumi.Output<number>;
    /**
     * The number of virtual CPUs available on the managed database.
     */
    public /*out*/ readonly planVcpus!: pulumi.Output<number>;
    /**
     * The connection port for the managed database.
     */
    public /*out*/ readonly port!: pulumi.Output<string>;
    /**
     * The public hostname assigned to the managed database (VPC-attached only).
     */
    public readonly publicHost!: pulumi.Output<string>;
    /**
     * A list of read replicas attached to the managed database.
     */
    public readonly readReplicas!: pulumi.Output<outputs.DatabaseReadReplica[]>;
    /**
     * The configuration value for the data eviction policy on the managed database (Redis engine types only - `noeviction`, `allkeys-lru`, `volatile-lru`, `allkeys-random`, `volatile-random`, `volatile-ttl`, `volatile-lfu`, `allkeys-lfu`).
     */
    public readonly redisEvictionPolicy!: pulumi.Output<string | undefined>;
    /**
     * The ID of the region that the managed database is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The current status of the managed database (poweroff, rebuilding, rebalancing, running).
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The tag to assign to the managed database.
     */
    public readonly tag!: pulumi.Output<string | undefined>;
    /**
     * A list of allowed IP addresses for the managed database.
     */
    public readonly trustedIps!: pulumi.Output<string[] | undefined>;
    /**
     * The primary admin user for the managed database.
     */
    public /*out*/ readonly user!: pulumi.Output<string>;
    /**
     * The ID of the VPC Network to attach to the Managed Database.
     */
    public readonly vpcId!: pulumi.Output<string | undefined>;

    /**
     * Create a Database resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseArgs | DatabaseState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatabaseState | undefined;
            resourceInputs["clusterTimeZone"] = state ? state.clusterTimeZone : undefined;
            resourceInputs["databaseEngine"] = state ? state.databaseEngine : undefined;
            resourceInputs["databaseEngineVersion"] = state ? state.databaseEngineVersion : undefined;
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
            resourceInputs["publicHost"] = state ? state.publicHost : undefined;
            resourceInputs["readReplicas"] = state ? state.readReplicas : undefined;
            resourceInputs["redisEvictionPolicy"] = state ? state.redisEvictionPolicy : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tag"] = state ? state.tag : undefined;
            resourceInputs["trustedIps"] = state ? state.trustedIps : undefined;
            resourceInputs["user"] = state ? state.user : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as DatabaseArgs | undefined;
            if ((!args || args.databaseEngine === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseEngine'");
            }
            if ((!args || args.databaseEngineVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseEngineVersion'");
            }
            if ((!args || args.label === undefined) && !opts.urn) {
                throw new Error("Missing required property 'label'");
            }
            if ((!args || args.plan === undefined) && !opts.urn) {
                throw new Error("Missing required property 'plan'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["clusterTimeZone"] = args ? args.clusterTimeZone : undefined;
            resourceInputs["databaseEngine"] = args ? args.databaseEngine : undefined;
            resourceInputs["databaseEngineVersion"] = args ? args.databaseEngineVersion : undefined;
            resourceInputs["label"] = args ? args.label : undefined;
            resourceInputs["maintenanceDow"] = args ? args.maintenanceDow : undefined;
            resourceInputs["maintenanceTime"] = args ? args.maintenanceTime : undefined;
            resourceInputs["mysqlLongQueryTime"] = args ? args.mysqlLongQueryTime : undefined;
            resourceInputs["mysqlRequirePrimaryKey"] = args ? args.mysqlRequirePrimaryKey : undefined;
            resourceInputs["mysqlSlowQueryLog"] = args ? args.mysqlSlowQueryLog : undefined;
            resourceInputs["mysqlSqlModes"] = args ? args.mysqlSqlModes : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["plan"] = args ? args.plan : undefined;
            resourceInputs["planDisk"] = args ? args.planDisk : undefined;
            resourceInputs["publicHost"] = args ? args.publicHost : undefined;
            resourceInputs["readReplicas"] = args ? args.readReplicas : undefined;
            resourceInputs["redisEvictionPolicy"] = args ? args.redisEvictionPolicy : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tag"] = args ? args.tag : undefined;
            resourceInputs["trustedIps"] = args ? args.trustedIps : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["dateCreated"] = undefined /*out*/;
            resourceInputs["dbname"] = undefined /*out*/;
            resourceInputs["host"] = undefined /*out*/;
            resourceInputs["latestBackup"] = undefined /*out*/;
            resourceInputs["planRam"] = undefined /*out*/;
            resourceInputs["planReplicas"] = undefined /*out*/;
            resourceInputs["planVcpus"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["user"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Database.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Database resources.
 */
export interface DatabaseState {
    /**
     * The configured time zone for the Managed Database in TZ database format (e.g. `UTC`, `America/New_York`, `Europe/London`).
     */
    clusterTimeZone?: pulumi.Input<string>;
    /**
     * The database engine of the new managed database.
     */
    databaseEngine?: pulumi.Input<string>;
    /**
     * The database engine version of the new managed database.
     */
    databaseEngineVersion?: pulumi.Input<string>;
    /**
     * The date the managed database was added to your Vultr account.
     */
    dateCreated?: pulumi.Input<string>;
    /**
     * The managed database's default logical database.
     */
    dbname?: pulumi.Input<string>;
    /**
     * The hostname assigned to the managed database.
     */
    host?: pulumi.Input<string>;
    /**
     * A label for the managed database.
     */
    label?: pulumi.Input<string>;
    /**
     * The date of the latest backup available on the managed database.
     */
    latestBackup?: pulumi.Input<string>;
    /**
     * The preferred maintenance day of week for the managed database.
     */
    maintenanceDow?: pulumi.Input<string>;
    /**
     * The preferred maintenance time for the managed database in 24-hour HH:00 format (e.g. `01:00`, `13:00`, `23:00`).
     */
    maintenanceTime?: pulumi.Input<string>;
    /**
     * The configuration value for the long query time (in seconds) on the managed database (MySQL engine types only).
     */
    mysqlLongQueryTime?: pulumi.Input<number>;
    /**
     * The configuration value for whether primary keys are required on the managed database (MySQL engine types only).
     */
    mysqlRequirePrimaryKey?: pulumi.Input<boolean>;
    /**
     * The configuration value for slow query logging on the managed database (MySQL engine types only).
     */
    mysqlSlowQueryLog?: pulumi.Input<boolean>;
    /**
     * A list of SQL modes to configure for the managed database (MySQL engine types only - `ALLOW_INVALID_DATES`, `ANSI`, `ANSI_QUOTES`, `ERROR_FOR_DIVISION_BY_ZERO`, `HIGH_NOT_PRECEDENCE`, `IGNORE_SPACE`, `NO_AUTO_VALUE_ON_ZERO`, `NO_DIR_IN_CREATE`, `NO_ENGINE_SUBSTITUTION`, `NO_UNSIGNED_SUBTRACTION`, `NO_ZERO_DATE`, `NO_ZERO_IN_DATE`, `ONLY_FULL_GROUP_BY`, `PIPES_AS_CONCAT`, `REAL_AS_FLOAT`, `STRICT_ALL_TABLES`, `STRICT_TRANS_TABLES`, `TIME_TRUNCATE_FRACTIONAL`, `TRADITIONAL`).
     */
    mysqlSqlModes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The password for the managed database's primary admin user.
     */
    password?: pulumi.Input<string>;
    /**
     * The ID of the plan that you want the managed database to subscribe to. [See List Managed Database Plans](https://www.vultr.com/api/#tag/managed-databases/operation/list-database-plans)
     */
    plan?: pulumi.Input<string>;
    /**
     * The description of the disk(s) on the managed database.
     */
    planDisk?: pulumi.Input<number>;
    /**
     * The amount of memory available on the managed database in MB.
     */
    planRam?: pulumi.Input<number>;
    /**
     * The number of standby nodes available on the managed database.
     */
    planReplicas?: pulumi.Input<number>;
    /**
     * The number of virtual CPUs available on the managed database.
     */
    planVcpus?: pulumi.Input<number>;
    /**
     * The connection port for the managed database.
     */
    port?: pulumi.Input<string>;
    /**
     * The public hostname assigned to the managed database (VPC-attached only).
     */
    publicHost?: pulumi.Input<string>;
    /**
     * A list of read replicas attached to the managed database.
     */
    readReplicas?: pulumi.Input<pulumi.Input<inputs.DatabaseReadReplica>[]>;
    /**
     * The configuration value for the data eviction policy on the managed database (Redis engine types only - `noeviction`, `allkeys-lru`, `volatile-lru`, `allkeys-random`, `volatile-random`, `volatile-ttl`, `volatile-lfu`, `allkeys-lfu`).
     */
    redisEvictionPolicy?: pulumi.Input<string>;
    /**
     * The ID of the region that the managed database is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
     */
    region?: pulumi.Input<string>;
    /**
     * The current status of the managed database (poweroff, rebuilding, rebalancing, running).
     */
    status?: pulumi.Input<string>;
    /**
     * The tag to assign to the managed database.
     */
    tag?: pulumi.Input<string>;
    /**
     * A list of allowed IP addresses for the managed database.
     */
    trustedIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The primary admin user for the managed database.
     */
    user?: pulumi.Input<string>;
    /**
     * The ID of the VPC Network to attach to the Managed Database.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Database resource.
 */
export interface DatabaseArgs {
    /**
     * The configured time zone for the Managed Database in TZ database format (e.g. `UTC`, `America/New_York`, `Europe/London`).
     */
    clusterTimeZone?: pulumi.Input<string>;
    /**
     * The database engine of the new managed database.
     */
    databaseEngine: pulumi.Input<string>;
    /**
     * The database engine version of the new managed database.
     */
    databaseEngineVersion: pulumi.Input<string>;
    /**
     * A label for the managed database.
     */
    label: pulumi.Input<string>;
    /**
     * The preferred maintenance day of week for the managed database.
     */
    maintenanceDow?: pulumi.Input<string>;
    /**
     * The preferred maintenance time for the managed database in 24-hour HH:00 format (e.g. `01:00`, `13:00`, `23:00`).
     */
    maintenanceTime?: pulumi.Input<string>;
    /**
     * The configuration value for the long query time (in seconds) on the managed database (MySQL engine types only).
     */
    mysqlLongQueryTime?: pulumi.Input<number>;
    /**
     * The configuration value for whether primary keys are required on the managed database (MySQL engine types only).
     */
    mysqlRequirePrimaryKey?: pulumi.Input<boolean>;
    /**
     * The configuration value for slow query logging on the managed database (MySQL engine types only).
     */
    mysqlSlowQueryLog?: pulumi.Input<boolean>;
    /**
     * A list of SQL modes to configure for the managed database (MySQL engine types only - `ALLOW_INVALID_DATES`, `ANSI`, `ANSI_QUOTES`, `ERROR_FOR_DIVISION_BY_ZERO`, `HIGH_NOT_PRECEDENCE`, `IGNORE_SPACE`, `NO_AUTO_VALUE_ON_ZERO`, `NO_DIR_IN_CREATE`, `NO_ENGINE_SUBSTITUTION`, `NO_UNSIGNED_SUBTRACTION`, `NO_ZERO_DATE`, `NO_ZERO_IN_DATE`, `ONLY_FULL_GROUP_BY`, `PIPES_AS_CONCAT`, `REAL_AS_FLOAT`, `STRICT_ALL_TABLES`, `STRICT_TRANS_TABLES`, `TIME_TRUNCATE_FRACTIONAL`, `TRADITIONAL`).
     */
    mysqlSqlModes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The password for the managed database's primary admin user.
     */
    password?: pulumi.Input<string>;
    /**
     * The ID of the plan that you want the managed database to subscribe to. [See List Managed Database Plans](https://www.vultr.com/api/#tag/managed-databases/operation/list-database-plans)
     */
    plan: pulumi.Input<string>;
    /**
     * The description of the disk(s) on the managed database.
     */
    planDisk?: pulumi.Input<number>;
    /**
     * The public hostname assigned to the managed database (VPC-attached only).
     */
    publicHost?: pulumi.Input<string>;
    /**
     * A list of read replicas attached to the managed database.
     */
    readReplicas?: pulumi.Input<pulumi.Input<inputs.DatabaseReadReplica>[]>;
    /**
     * The configuration value for the data eviction policy on the managed database (Redis engine types only - `noeviction`, `allkeys-lru`, `volatile-lru`, `allkeys-random`, `volatile-random`, `volatile-ttl`, `volatile-lfu`, `allkeys-lfu`).
     */
    redisEvictionPolicy?: pulumi.Input<string>;
    /**
     * The ID of the region that the managed database is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
     */
    region: pulumi.Input<string>;
    /**
     * The tag to assign to the managed database.
     */
    tag?: pulumi.Input<string>;
    /**
     * A list of allowed IP addresses for the managed database.
     */
    trustedIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the VPC Network to attach to the Managed Database.
     */
    vpcId?: pulumi.Input<string>;
}
