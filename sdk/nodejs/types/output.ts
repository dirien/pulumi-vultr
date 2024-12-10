// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface DatabaseReadReplica {
    /**
     * The configured time zone for the Managed Database in TZ database format (e.g. `UTC`, `America/New_York`, `Europe/London`).
     */
    clusterTimeZone: string;
    /**
     * The database engine of the new managed database.
     */
    databaseEngine: string;
    /**
     * The database engine version of the new managed database.
     */
    databaseEngineVersion: string;
    /**
     * The date the managed database was added to your Vultr account.
     */
    dateCreated: string;
    /**
     * The managed database's default logical database.
     */
    dbname: string;
    /**
     * The configuration value for the data eviction policy on the managed database (Redis engine types only - `noeviction`, `allkeys-lru`, `volatile-lru`, `allkeys-random`, `volatile-random`, `volatile-ttl`, `volatile-lfu`, `allkeys-lfu`).
     */
    evictionPolicy: string;
    /**
     * An associated list of FerretDB connection credentials (FerretDB + PostgreSQL engine types only).
     */
    ferretdbCredentials: {[key: string]: string};
    /**
     * The hostname assigned to the managed database.
     */
    host: string;
    /**
     * The ID of the managed database.
     */
    id: string;
    /**
     * A label for the managed database.
     */
    label: string;
    /**
     * The date of the latest backup available on the managed database.
     */
    latestBackup: string;
    /**
     * The preferred maintenance day of week for the managed database.
     */
    maintenanceDow: string;
    /**
     * The preferred maintenance time for the managed database in 24-hour HH:00 format (e.g. `01:00`, `13:00`, `23:00`).
     */
    maintenanceTime: string;
    /**
     * The configuration value for the long query time (in seconds) on the managed database (MySQL engine types only).
     */
    mysqlLongQueryTime: number;
    /**
     * The configuration value for whether primary keys are required on the managed database (MySQL engine types only).
     */
    mysqlRequirePrimaryKey: boolean;
    /**
     * The configuration value for slow query logging on the managed database (MySQL engine types only).
     */
    mysqlSlowQueryLog: boolean;
    /**
     * A list of SQL modes to configure for the managed database (MySQL engine types only - `ALLOW_INVALID_DATES`, `ANSI`, `ANSI_QUOTES`, `ERROR_FOR_DIVISION_BY_ZERO`, `HIGH_NOT_PRECEDENCE`, `IGNORE_SPACE`, `NO_AUTO_VALUE_ON_ZERO`, `NO_DIR_IN_CREATE`, `NO_ENGINE_SUBSTITUTION`, `NO_UNSIGNED_SUBTRACTION`, `NO_ZERO_DATE`, `NO_ZERO_IN_DATE`, `ONLY_FULL_GROUP_BY`, `PIPES_AS_CONCAT`, `REAL_AS_FLOAT`, `STRICT_ALL_TABLES`, `STRICT_TRANS_TABLES`, `TIME_TRUNCATE_FRACTIONAL`, `TRADITIONAL`).
     */
    mysqlSqlModes: string[];
    /**
     * The password for the managed database's primary admin user.
     */
    password: string;
    /**
     * The ID of the plan that you want the managed database to subscribe to. [See List Managed Database Plans](https://www.vultr.com/api/#tag/managed-databases/operation/list-database-plans)
     */
    plan: string;
    /**
     * The description of the disk(s) on the managed database.
     */
    planDisk: number;
    /**
     * The amount of memory available on the managed database in MB.
     */
    planRam: number;
    /**
     * The number of standby nodes available on the managed database (excluded for Kafka engine types).
     */
    planReplicas: number;
    /**
     * The number of virtual CPUs available on the managed database.
     */
    planVcpus: number;
    /**
     * The connection port for the managed database.
     */
    port: string;
    /**
     * The public hostname assigned to the managed database (VPC-attached only).
     */
    publicHost: string;
    /**
     * The ID of the region that the managed database is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
     */
    region: string;
    /**
     * The current status of the managed database (poweroff, rebuilding, rebalancing, configuring, running).
     */
    status: string;
    /**
     * The tag to assign to the managed database.
     */
    tag: string;
    /**
     * A list of allowed IP addresses for the managed database.
     */
    trustedIps: string[];
    /**
     * The primary admin user for the managed database.
     */
    user: string;
    /**
     * The ID of the VPC Network to attach to the Managed Database.
     */
    vpcId: string;
}

export interface DatabaseUserAccessControl {
    /**
     * List of command category rules for this managed database user (Redis engine types only).
     */
    aclCategories: string[];
    /**
     * List of publish/subscribe channel patterns for this managed database user (Redis engine types only).
     */
    aclChannels: string[];
    /**
     * List of individual command rules for this managed database user (Redis engine types only).
     */
    aclCommands: string[];
    /**
     * List of access rules for this managed database user (Redis engine types only).
     */
    aclKeys: string[];
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

export interface GetContainerRegistryFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
}

export interface GetContainerRegistryRepository {
    /**
     * A count of the artifacts in the repository.
     */
    artifactCount: number;
    /**
     * A date-time of when the root user was created.
     */
    dateCreated: string;
    /**
     * The date-time that the repository was last updated.
     */
    dateModified: string;
    /**
     * A description of the repo, if set.
     */
    description: string;
    /**
     * The image name in the repository.
     */
    image: string;
    /**
     * The name of the repository.
     */
    name: string;
    /**
     * A count of the number of pulls against the repository.
     */
    pullCount: number;
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

export interface GetDatabaseReadReplica {
    /**
     * The configured time zone for the Managed Database in TZ database format.
     */
    clusterTimeZone: string;
    /**
     * The database engine of the managed database.
     */
    databaseEngine: string;
    /**
     * The database engine version of the managed database.
     */
    databaseEngineVersion: string;
    /**
     * The date the managed database was added to your Vultr account.
     */
    dateCreated: string;
    /**
     * The managed database's default logical database.
     */
    dbname: string;
    /**
     * The configuration value for the data eviction policy on the managed database (Redis engine types only).
     */
    evictionPolicy: string;
    /**
     * An associated list of FerretDB connection credentials (FerretDB + PostgreSQL engine types only).
     */
    ferretdbCredentials: {[key: string]: string};
    /**
     * The hostname assigned to the managed database.
     */
    host: string;
    id: string;
    /**
     * The managed database's label.
     */
    label: string;
    /**
     * The date of the latest backup available on the managed database.
     */
    latestBackup: string;
    /**
     * The preferred maintenance day of week for the managed database.
     */
    maintenanceDow: string;
    /**
     * The preferred maintenance time for the managed database.
     */
    maintenanceTime: string;
    /**
     * The configuration value for the long query time (in seconds) on the managed database (MySQL engine types only).
     */
    mysqlLongQueryTime: number;
    /**
     * The configuration value for whether primary keys are required on the managed database (MySQL engine types only).
     */
    mysqlRequirePrimaryKey: boolean;
    /**
     * The configuration value for slow query logging on the managed database (MySQL engine types only).
     */
    mysqlSlowQueryLog: boolean;
    /**
     * A list of SQL modes currently configured for the managed database (MySQL engine types only).
     */
    mysqlSqlModes: string[];
    /**
     * The password for the managed database's primary admin user.
     */
    password: string;
    /**
     * The managed database's plan ID.
     */
    plan: string;
    /**
     * The description of the disk(s) on the managed database.
     */
    planDisk: number;
    /**
     * The amount of memory available on the managed database in MB.
     */
    planRam: number;
    /**
     * The number of standby nodes available on the managed database.
     */
    planReplicas: number;
    /**
     * The number of virtual CPUs available on the managed database.
     */
    planVcpus: number;
    /**
     * The connection port for the managed database.
     */
    port: string;
    /**
     * The public hostname assigned to the managed database (VPC-attached only).
     */
    publicHost: string;
    /**
     * The region ID of the managed database.
     */
    region: string;
    /**
     * The current status of the managed database (poweroff, rebuilding, rebalancing, configuring, running).
     */
    status: string;
    /**
     * The managed database's tag.
     */
    tag: string;
    /**
     * A list of allowed IP addresses for the managed database.
     */
    trustedIps: string[];
    /**
     * The primary admin user for the managed database.
     */
    user: string;
    /**
     * The ID of the VPC Network attached to the Managed Database.
     */
    vpcId: string;
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

export interface GetInferenceFilter {
    /**
     * Attribute name to filter with.
     */
    name: string;
    /**
     * One or more values filter with.
     */
    values: string[];
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

export interface GetInstancesInstance {
    /**
     * The server's allowed bandwidth usage in GB.
     */
    allowedBandwidth: number;
    /**
     * The server's application ID.
     */
    appId: number;
    backups: string;
    /**
     * The current configuration for backups
     */
    backupsSchedule: {[key: string]: string};
    /**
     * The date the server was added to your Vultr account.
     */
    dateCreated: string;
    /**
     * The description of the disk(s) on the server.
     */
    disk: number;
    /**
     * Array of which features are enabled.
     */
    features: string[];
    /**
     * The ID of the firewall group applied to this server.
     */
    firewallGroupId: string;
    /**
     * The server's IPv4 gateway.
     */
    gatewayV4: string;
    /**
     * The hostname assigned to the server.
     */
    hostname: string;
    id: string;
    /**
     * The Marketplace ID for this application.
     */
    imageId: string;
    /**
     * The server's internal IP address.
     */
    internalIp: string;
    /**
     * The server's current KVM URL. This URL will change periodically. It is not advised to cache this value.
     */
    kvm: string;
    /**
     * The server's label.
     */
    label: string;
    location: string;
    /**
     * The server's main IP address.
     */
    mainIp: string;
    /**
     * The server's IPv4 netmask.
     */
    netmaskV4: string;
    /**
     * The operating system of the instance.
     */
    os: string;
    /**
     * The server's operating system ID.
     */
    osId: number;
    /**
     * The server's plan ID.
     */
    plan: string;
    /**
     * Whether the server is powered on or not.
     */
    powerStatus: string;
    privateNetworkIds: string[];
    /**
     * The amount of memory available on the instance in MB.
     */
    ram: number;
    /**
     * The region ID of the server.
     */
    region: string;
    /**
     * A more detailed server status (none, locked, installingbooting, isomounting, ok).
     */
    serverStatus: string;
    /**
     * The status of the server's subscription.
     */
    status: string;
    /**
     * A list of tags applied to the instance.
     */
    tags: string[];
    /**
     * The scheme used for the default user (linux servers only).
     */
    userScheme: string;
    /**
     * The main IPv6 network address.
     */
    v6MainIp: string;
    /**
     * The IPv6 subnet.
     */
    v6Network: string;
    /**
     * The IPv6 network size in bits.
     */
    v6NetworkSize: number;
    /**
     * The number of virtual CPUs available on the server.
     */
    vcpuCount: number;
    vpcIds: string[];
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

export interface GetKubernetesNodePool {
    /**
     * Boolean indicating if the auto scaler for the default node pool is active.
     */
    autoScaler?: boolean;
    /**
     * Date node was created.
     */
    dateCreated: string;
    /**
     * Date of node pool updates.
     */
    dateUpdated: string;
    /**
     * ID of node.
     */
    id: string;
    /**
     * Label of node.
     */
    label: string;
    /**
     * The maximum number of nodes used by the auto scaler.
     */
    maxNodes?: number;
    /**
     * The minimum number of nodes used by the auto scaler.
     */
    minNodes?: number;
    /**
     * Number of nodes within node pool.
     */
    nodeQuantity: number;
    /**
     * Array that contains information about nodes within this node pool.
     */
    nodes: outputs.GetKubernetesNodePoolNode[];
    /**
     * Node plan that nodes are using within this node pool.
     */
    plan: string;
    /**
     * Status of node.
     */
    status: string;
    /**
     * Tag for node pool.
     */
    tag: string;
}

export interface GetKubernetesNodePoolNode {
    /**
     * Date node was created.
     */
    dateCreated: string;
    /**
     * ID of node.
     */
    id: string;
    /**
     * Label of node.
     */
    label: string;
    /**
     * Status of node.
     */
    status: string;
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

export interface InstanceBackupsSchedule {
    /**
     * Day of month to run. Use values between 1 and 28.
     */
    dom: number;
    /**
     * Day of week to run. `1 = Sunday`, `2 = Monday`, `3 = Tuesday`, `4 = Wednesday`, `5 = Thursday`, `6 = Friday`, `7 = Saturday`
     */
    dow: number;
    /**
     * Hour of day to run in UTC.
     */
    hour: number;
    /**
     * Type of backup schedule Possible values are `daily`, `weekly`, `monthly`, `dailyAltEven`, or `dailyAltOdd`.
     */
    type: string;
}

export interface KubernetesNodePools {
    /**
     * Enable the auto scaler for the default node pool.
     */
    autoScaler?: boolean;
    /**
     * Date node was created.
     */
    dateCreated: string;
    /**
     * Date of node pool updates.
     */
    dateUpdated: string;
    /**
     * ID of node.
     */
    id: string;
    /**
     * The label to be used as a prefix for nodes in this node pool.
     */
    label: string;
    /**
     * The maximum number of nodes to use with the auto scaler.
     */
    maxNodes?: number;
    /**
     * The minimum number of nodes to use with the auto scaler.
     */
    minNodes?: number;
    /**
     * The number of nodes in this node pool.
     */
    nodeQuantity: number;
    /**
     * Array that contains information about nodes within this node pool.
     */
    nodes: outputs.KubernetesNodePoolsNode[];
    /**
     * The plan to be used in this node pool. [See Plans List](https://www.vultr.com/api/#operation/list-plans) Note the minimum plan requirements must have at least 1 core and 2 gbs of memory.
     */
    plan: string;
    /**
     * Status of node.
     */
    status: string;
    /**
     * Tag for node pool.
     */
    tag: string;
}

export interface KubernetesNodePoolsNode {
    /**
     * Date node was created.
     */
    dateCreated: string;
    /**
     * ID of node.
     */
    id: string;
    /**
     * The label to be used as a prefix for nodes in this node pool.
     */
    label: string;
    /**
     * Status of node.
     */
    status: string;
}

export interface LoadBalancerFirewallRule {
    /**
     * The load balancer ID.
     */
    id: string;
    /**
     * The type of ip this rule is - may be either v4 or v6.
     */
    ipType: string;
    /**
     * The assigned port (integer) on the attached instances that the load balancer should check against. Default value is `80`.
     */
    port: number;
    /**
     * IP address with subnet that is allowed through the firewall. You may also pass in `cloudflare` which will allow only CloudFlares IP range.
     */
    source: string;
}

export interface LoadBalancerForwardingRule {
    /**
     * Port on instance side.
     */
    backendPort: number;
    /**
     * Protocol on instance side. Possible values: "http", "https", "tcp".
     */
    backendProtocol: string;
    /**
     * Port on load balancer side.
     */
    frontendPort: number;
    /**
     * Protocol on load balancer side. Possible values: "http", "https", "tcp".
     */
    frontendProtocol: string;
    ruleId: string;
}

export interface LoadBalancerHealthCheck {
    /**
     * Time in seconds to perform health check. Default value is 15.
     */
    checkInterval?: number;
    /**
     * Number of failed attempts encountered before failover. Default value is 5.
     */
    healthyThreshold?: number;
    /**
     * The path on the attached instances that the load balancer should check against. Default value is `/`
     */
    path?: string;
    /**
     * The assigned port (integer) on the attached instances that the load balancer should check against. Default value is `80`.
     */
    port: number;
    /**
     * The protocol used to traffic requests to the load balancer. Possible values are `http`, or `tcp`. Default value is `http`.
     */
    protocol: string;
    /**
     * Time in seconds to wait for a health check response. Default value is 5.
     */
    responseTimeout?: number;
    /**
     * Number of failed attempts encountered before failover. Default value is 5.
     */
    unhealthyThreshold?: number;
}

export interface LoadBalancerSsl {
    /**
     * The SSL Certificate.
     */
    certificate: string;
    /**
     * The SSL certificate chain.
     */
    chain?: string;
    /**
     * The SSL certificates private key.
     */
    privateKey: string;
}

