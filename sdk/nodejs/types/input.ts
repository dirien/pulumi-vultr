// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface DatabaseReadReplica {
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
     * An associated list of FerretDB connection credentials (FerretDB + PostgreSQL engine types only).
     */
    ferretdbCredentials?: pulumi.Input<{[key: string]: any}>;
    /**
     * The hostname assigned to the managed database.
     */
    host?: pulumi.Input<string>;
    /**
     * The ID of the managed database.
     */
    id?: pulumi.Input<string>;
    /**
     * A label for the managed database.
     */
    label: pulumi.Input<string>;
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
     * The configuration value for the data eviction policy on the managed database (Redis engine types only - `noeviction`, `allkeys-lru`, `volatile-lru`, `allkeys-random`, `volatile-random`, `volatile-ttl`, `volatile-lfu`, `allkeys-lfu`).
     */
    redisEvictionPolicy?: pulumi.Input<string>;
    /**
     * The ID of the region that the managed database is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
     */
    region: pulumi.Input<string>;
    /**
     * The current status of the managed database (poweroff, rebuilding, rebalancing, configuring, running).
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

export interface DatabaseUserAccessControl {
    /**
     * The list of command category rules for this managed database user.
     */
    redisAclCategories: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of publish/subscribe channel patterns for this managed database user.
     */
    redisAclChannels: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of individual command rules for this managed database user.
     */
    redisAclCommands: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of access rules for this managed database user.
     */
    redisAclKeys: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetApplicationFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetApplicationFilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetBackupFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetBackupFilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetBareMetalPlanFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetBareMetalPlanFilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetBareMetalServerFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetBareMetalServerFilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetBlockStorageFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetBlockStorageFilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetDatabaseFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetDatabaseFilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetFirewallGroupFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetFirewallGroupFilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetInstanceFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetInstanceFilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetInstanceIpv4Filter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values to filter with.
     */
    values: string[];
}

export interface GetInstanceIpv4FilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values to filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetInstancesFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetInstancesFilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetIsoPrivateFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetIsoPrivateFilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetIsoPublicFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetIsoPublicFilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetKubernetesFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetKubernetesFilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetLoadBalancerFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetLoadBalancerFilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetObjectStorageClusterFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetObjectStorageClusterFilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetObjectStorageFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetObjectStorageFilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetOsFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetOsFilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetPlanFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetPlanFilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetPrivateNetworkFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetPrivateNetworkFilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetRegionFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetRegionFilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetReservedIpFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetReservedIpFilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetReverseIpv4Filter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values to filter with.
     */
    values: string[];
}

export interface GetReverseIpv4FilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values to filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetReverseIpv6Filter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values to filter with.
     */
    values: string[];
}

export interface GetReverseIpv6FilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values to filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetSnapshotFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetSnapshotFilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetSshKeyFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetSshKeyFilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetStartupScriptFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetStartupScriptFilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetUserFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetUserFilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetVpc2Filter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetVpc2FilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetVpcFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetVpcFilterArgs {
    /**
     * Attribute name to filter with.
     */
    name: pulumi.Input<string>;
    /**
     * One or more values filter with.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface InstanceBackupsSchedule {
    /**
     * Day of month to run. Use values between 1 and 28.
     */
    dom?: pulumi.Input<number>;
    /**
     * Day of week to run. `1 = Sunday`, `2 = Monday`, `3 = Tuesday`, `4 = Wednesday`, `5 = Thursday`, `6 = Friday`, `7 = Saturday`
     */
    dow?: pulumi.Input<number>;
    /**
     * Hour of day to run in UTC.
     */
    hour?: pulumi.Input<number>;
    /**
     * Type of backup schedule Possible values are `daily`, `weekly`, `monthly`, `dailyAltEven`, or `dailyAltOdd`.
     */
    type: pulumi.Input<string>;
}

export interface KubernetesNodePools {
    /**
     * Enable the auto scaler for the default node pool.
     */
    autoScaler?: pulumi.Input<boolean>;
    /**
     * Date node was created.
     */
    dateCreated?: pulumi.Input<string>;
    /**
     * Date of node pool updates.
     */
    dateUpdated?: pulumi.Input<string>;
    /**
     * ID of node.
     */
    id?: pulumi.Input<string>;
    /**
     * The label to be used as a prefix for nodes in this node pool.
     */
    label: pulumi.Input<string>;
    /**
     * The maximum number of nodes to use with the auto scaler.
     */
    maxNodes?: pulumi.Input<number>;
    /**
     * The minimum number of nodes to use with the auto scaler.
     */
    minNodes?: pulumi.Input<number>;
    /**
     * The number of nodes in this node pool.
     */
    nodeQuantity: pulumi.Input<number>;
    /**
     * Array that contains information about nodes within this node pool.
     */
    nodes?: pulumi.Input<pulumi.Input<inputs.KubernetesNodePoolsNode>[]>;
    /**
     * The plan to be used in this node pool. [See Plans List](https://www.vultr.com/api/#operation/list-plans) Note the minimum plan requirements must have at least 1 core and 2 gbs of memory.
     */
    plan: pulumi.Input<string>;
    /**
     * Status of node.
     */
    status?: pulumi.Input<string>;
    /**
     * Tag for node pool.
     */
    tag?: pulumi.Input<string>;
}

export interface KubernetesNodePoolsNode {
    /**
     * Date node was created.
     */
    dateCreated?: pulumi.Input<string>;
    /**
     * ID of node.
     */
    id?: pulumi.Input<string>;
    /**
     * The label to be used as a prefix for nodes in this node pool.
     */
    label?: pulumi.Input<string>;
    /**
     * Status of node.
     */
    status?: pulumi.Input<string>;
}

export interface LoadBalancerFirewallRule {
    /**
     * The load balancer ID.
     */
    id?: pulumi.Input<string>;
    /**
     * The type of ip this rule is - may be either v4 or v6.
     */
    ipType: pulumi.Input<string>;
    port: pulumi.Input<number>;
    /**
     * IP address with subnet that is allowed through the firewall. You may also pass in `cloudflare` which will allow only CloudFlares IP range.
     */
    source: pulumi.Input<string>;
}

export interface LoadBalancerForwardingRule {
    /**
     * Port on instance side.
     */
    backendPort: pulumi.Input<number>;
    /**
     * Protocol on instance side. Possible values: "http", "https", "tcp".
     */
    backendProtocol: pulumi.Input<string>;
    /**
     * Port on load balancer side.
     */
    frontendPort: pulumi.Input<number>;
    /**
     * Protocol on load balancer side. Possible values: "http", "https", "tcp".
     */
    frontendProtocol: pulumi.Input<string>;
    ruleId?: pulumi.Input<string>;
}

export interface LoadBalancerHealthCheck {
    /**
     * Time in seconds to perform health check. Default value is 15.
     */
    checkInterval?: pulumi.Input<number>;
    /**
     * Number of failed attempts encountered before failover. Default value is 5.
     */
    healthyThreshold?: pulumi.Input<number>;
    /**
     * The path on the attached instances that the load balancer should check against. Default value is `/`
     */
    path?: pulumi.Input<string>;
    /**
     * The assigned port (integer) on the attached instances that the load balancer should check against. Default value is `80`.
     */
    port: pulumi.Input<number>;
    /**
     * The protocol used to traffic requests to the load balancer. Possible values are `http`, or `tcp`. Default value is `http`.
     */
    protocol: pulumi.Input<string>;
    /**
     * Time in seconds to wait for a health check response. Default value is 5.
     */
    responseTimeout?: pulumi.Input<number>;
    /**
     * Number of failed attempts encountered before failover. Default value is 5.
     */
    unhealthyThreshold?: pulumi.Input<number>;
}

export interface LoadBalancerSsl {
    /**
     * The SSL Certificate.
     */
    certificate: pulumi.Input<string>;
    /**
     * The SSL certificate chain.
     */
    chain?: pulumi.Input<string>;
    /**
     * The SSL certificates private key.
     */
    privateKey: pulumi.Input<string>;
}
