// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Vultr.Outputs
{

    [OutputType]
    public sealed class DatabaseReadReplica
    {
        /// <summary>
        /// The configured time zone for the Managed Database in TZ database format (e.g. `UTC`, `America/New_York`, `Europe/London`).
        /// </summary>
        public readonly string? ClusterTimeZone;
        /// <summary>
        /// The database engine of the new managed database.
        /// </summary>
        public readonly string? DatabaseEngine;
        /// <summary>
        /// The database engine version of the new managed database.
        /// </summary>
        public readonly string? DatabaseEngineVersion;
        /// <summary>
        /// The date the managed database was added to your Vultr account.
        /// </summary>
        public readonly string? DateCreated;
        /// <summary>
        /// The managed database's default logical database.
        /// </summary>
        public readonly string? Dbname;
        /// <summary>
        /// An associated list of FerretDB connection credentials (FerretDB + PostgreSQL engine types only).
        /// </summary>
        public readonly ImmutableDictionary<string, object>? FerretdbCredentials;
        /// <summary>
        /// The hostname assigned to the managed database.
        /// </summary>
        public readonly string? Host;
        /// <summary>
        /// The ID of the managed database.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// A label for the managed database.
        /// </summary>
        public readonly string Label;
        /// <summary>
        /// The date of the latest backup available on the managed database.
        /// </summary>
        public readonly string? LatestBackup;
        /// <summary>
        /// The preferred maintenance day of week for the managed database.
        /// </summary>
        public readonly string? MaintenanceDow;
        /// <summary>
        /// The preferred maintenance time for the managed database in 24-hour HH:00 format (e.g. `01:00`, `13:00`, `23:00`).
        /// </summary>
        public readonly string? MaintenanceTime;
        /// <summary>
        /// The configuration value for the long query time (in seconds) on the managed database (MySQL engine types only).
        /// </summary>
        public readonly int? MysqlLongQueryTime;
        /// <summary>
        /// The configuration value for whether primary keys are required on the managed database (MySQL engine types only).
        /// </summary>
        public readonly bool? MysqlRequirePrimaryKey;
        /// <summary>
        /// The configuration value for slow query logging on the managed database (MySQL engine types only).
        /// </summary>
        public readonly bool? MysqlSlowQueryLog;
        /// <summary>
        /// A list of SQL modes to configure for the managed database (MySQL engine types only - `ALLOW_INVALID_DATES`, `ANSI`, `ANSI_QUOTES`, `ERROR_FOR_DIVISION_BY_ZERO`, `HIGH_NOT_PRECEDENCE`, `IGNORE_SPACE`, `NO_AUTO_VALUE_ON_ZERO`, `NO_DIR_IN_CREATE`, `NO_ENGINE_SUBSTITUTION`, `NO_UNSIGNED_SUBTRACTION`, `NO_ZERO_DATE`, `NO_ZERO_IN_DATE`, `ONLY_FULL_GROUP_BY`, `PIPES_AS_CONCAT`, `REAL_AS_FLOAT`, `STRICT_ALL_TABLES`, `STRICT_TRANS_TABLES`, `TIME_TRUNCATE_FRACTIONAL`, `TRADITIONAL`).
        /// </summary>
        public readonly ImmutableArray<string> MysqlSqlModes;
        /// <summary>
        /// The password for the managed database's primary admin user.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// The ID of the plan that you want the managed database to subscribe to. [See List Managed Database Plans](https://www.vultr.com/api/#tag/managed-databases/operation/list-database-plans)
        /// </summary>
        public readonly string? Plan;
        /// <summary>
        /// The description of the disk(s) on the managed database.
        /// </summary>
        public readonly int? PlanDisk;
        /// <summary>
        /// The amount of memory available on the managed database in MB.
        /// </summary>
        public readonly int? PlanRam;
        /// <summary>
        /// The number of standby nodes available on the managed database.
        /// </summary>
        public readonly int? PlanReplicas;
        /// <summary>
        /// The number of virtual CPUs available on the managed database.
        /// </summary>
        public readonly int? PlanVcpus;
        /// <summary>
        /// The connection port for the managed database.
        /// </summary>
        public readonly string? Port;
        /// <summary>
        /// The public hostname assigned to the managed database (VPC-attached only).
        /// </summary>
        public readonly string? PublicHost;
        /// <summary>
        /// The configuration value for the data eviction policy on the managed database (Redis engine types only - `noeviction`, `allkeys-lru`, `volatile-lru`, `allkeys-random`, `volatile-random`, `volatile-ttl`, `volatile-lfu`, `allkeys-lfu`).
        /// </summary>
        public readonly string? RedisEvictionPolicy;
        /// <summary>
        /// The ID of the region that the managed database is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The current status of the managed database (poweroff, rebuilding, rebalancing, configuring, running).
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// The tag to assign to the managed database.
        /// </summary>
        public readonly string? Tag;
        /// <summary>
        /// A list of allowed IP addresses for the managed database.
        /// </summary>
        public readonly ImmutableArray<string> TrustedIps;
        /// <summary>
        /// The primary admin user for the managed database.
        /// </summary>
        public readonly string? User;
        /// <summary>
        /// The ID of the VPC Network to attach to the Managed Database.
        /// </summary>
        public readonly string? VpcId;

        [OutputConstructor]
        private DatabaseReadReplica(
            string? clusterTimeZone,

            string? databaseEngine,

            string? databaseEngineVersion,

            string? dateCreated,

            string? dbname,

            ImmutableDictionary<string, object>? ferretdbCredentials,

            string? host,

            string? id,

            string label,

            string? latestBackup,

            string? maintenanceDow,

            string? maintenanceTime,

            int? mysqlLongQueryTime,

            bool? mysqlRequirePrimaryKey,

            bool? mysqlSlowQueryLog,

            ImmutableArray<string> mysqlSqlModes,

            string? password,

            string? plan,

            int? planDisk,

            int? planRam,

            int? planReplicas,

            int? planVcpus,

            string? port,

            string? publicHost,

            string? redisEvictionPolicy,

            string region,

            string? status,

            string? tag,

            ImmutableArray<string> trustedIps,

            string? user,

            string? vpcId)
        {
            ClusterTimeZone = clusterTimeZone;
            DatabaseEngine = databaseEngine;
            DatabaseEngineVersion = databaseEngineVersion;
            DateCreated = dateCreated;
            Dbname = dbname;
            FerretdbCredentials = ferretdbCredentials;
            Host = host;
            Id = id;
            Label = label;
            LatestBackup = latestBackup;
            MaintenanceDow = maintenanceDow;
            MaintenanceTime = maintenanceTime;
            MysqlLongQueryTime = mysqlLongQueryTime;
            MysqlRequirePrimaryKey = mysqlRequirePrimaryKey;
            MysqlSlowQueryLog = mysqlSlowQueryLog;
            MysqlSqlModes = mysqlSqlModes;
            Password = password;
            Plan = plan;
            PlanDisk = planDisk;
            PlanRam = planRam;
            PlanReplicas = planReplicas;
            PlanVcpus = planVcpus;
            Port = port;
            PublicHost = publicHost;
            RedisEvictionPolicy = redisEvictionPolicy;
            Region = region;
            Status = status;
            Tag = tag;
            TrustedIps = trustedIps;
            User = user;
            VpcId = vpcId;
        }
    }
}
