// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Get information about a Vultr database.
 *
 * ## Example Usage
 *
 * Get the information for a database by `label`:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@pulumi/vultr";
 *
 * const myDatabase = vultr.getDatabase({
 *     filters: [{
 *         name: "label",
 *         values: ["my-database-label"],
 *     }],
 * });
 * ```
 */
export function getDatabase(args?: GetDatabaseArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabaseResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vultr:index/getDatabase:getDatabase", {
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabase.
 */
export interface GetDatabaseArgs {
    /**
     * Query parameters for finding databases.
     */
    filters?: inputs.GetDatabaseFilter[];
}

/**
 * A collection of values returned by getDatabase.
 */
export interface GetDatabaseResult {
    /**
     * The certificate to authenticate the default user (Kafka engine types only).
     */
    readonly accessCert: string;
    /**
     * The private key to authenticate the default user (Kafka engine types only).
     */
    readonly accessKey: string;
    /**
     * The configured time zone for the Managed Database in TZ database format.
     */
    readonly clusterTimeZone: string;
    /**
     * The database engine of the managed database.
     */
    readonly databaseEngine: string;
    /**
     * The database engine version of the managed database.
     */
    readonly databaseEngineVersion: string;
    /**
     * The date the managed database was added to your Vultr account.
     */
    readonly dateCreated: string;
    /**
     * The managed database's default logical database.
     */
    readonly dbname: string;
    /**
     * The configuration value for the data eviction policy on the managed database (Redis engine types only).
     */
    readonly evictionPolicy: string;
    /**
     * An associated list of FerretDB connection credentials (FerretDB + PostgreSQL engine types only).
     */
    readonly ferretdbCredentials: {[key: string]: string};
    readonly filters?: outputs.GetDatabaseFilter[];
    /**
     * The hostname assigned to the managed database.
     */
    readonly host: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The managed database's label.
     */
    readonly label: string;
    /**
     * The date of the latest backup available on the managed database.
     */
    readonly latestBackup: string;
    /**
     * The preferred maintenance day of week for the managed database.
     */
    readonly maintenanceDow: string;
    /**
     * The preferred maintenance time for the managed database.
     */
    readonly maintenanceTime: string;
    /**
     * The configuration value for the long query time (in seconds) on the managed database (MySQL engine types only).
     */
    readonly mysqlLongQueryTime: number;
    /**
     * The configuration value for whether primary keys are required on the managed database (MySQL engine types only).
     */
    readonly mysqlRequirePrimaryKey: boolean;
    /**
     * The configuration value for slow query logging on the managed database (MySQL engine types only).
     */
    readonly mysqlSlowQueryLog: boolean;
    /**
     * A list of SQL modes currently configured for the managed database (MySQL engine types only).
     */
    readonly mysqlSqlModes: string[];
    /**
     * The password for the managed database's primary admin user.
     */
    readonly password: string;
    /**
     * The managed database's plan ID.
     */
    readonly plan: string;
    readonly planBrokers: number;
    /**
     * The description of the disk(s) on the managed database.
     */
    readonly planDisk: number;
    /**
     * The amount of memory available on the managed database in MB.
     */
    readonly planRam: number;
    /**
     * The number of standby nodes available on the managed database.
     */
    readonly planReplicas: number;
    /**
     * The number of virtual CPUs available on the managed database.
     */
    readonly planVcpus: number;
    /**
     * The connection port for the managed database.
     */
    readonly port: string;
    /**
     * The public hostname assigned to the managed database (VPC-attached only).
     */
    readonly publicHost: string;
    /**
     * A list of read replicas attached to the managed database.
     */
    readonly readReplicas: outputs.GetDatabaseReadReplica[];
    /**
     * The region ID of the managed database.
     */
    readonly region: string;
    /**
     * The SASL connection port for the managed database (Kafka engine types only).
     */
    readonly saslPort: string;
    /**
     * The current status of the managed database (poweroff, rebuilding, rebalancing, configuring, running).
     */
    readonly status: string;
    /**
     * The managed database's tag.
     */
    readonly tag: string;
    /**
     * A list of allowed IP addresses for the managed database.
     */
    readonly trustedIps: string[];
    /**
     * The primary admin user for the managed database.
     */
    readonly user: string;
    /**
     * The ID of the VPC Network attached to the Managed Database.
     */
    readonly vpcId: string;
}
/**
 * Get information about a Vultr database.
 *
 * ## Example Usage
 *
 * Get the information for a database by `label`:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vultr from "@pulumi/vultr";
 *
 * const myDatabase = vultr.getDatabase({
 *     filters: [{
 *         name: "label",
 *         values: ["my-database-label"],
 *     }],
 * });
 * ```
 */
export function getDatabaseOutput(args?: GetDatabaseOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatabaseResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vultr:index/getDatabase:getDatabase", {
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabase.
 */
export interface GetDatabaseOutputArgs {
    /**
     * Query parameters for finding databases.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.GetDatabaseFilterArgs>[]>;
}
