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
    public sealed class GetDatabaseReadReplicaResult
    {
        public readonly string ClusterTimeZone;
        public readonly string DatabaseEngine;
        public readonly string DatabaseEngineVersion;
        public readonly string DateCreated;
        public readonly string Dbname;
        public readonly string Host;
        public readonly string Id;
        public readonly string Label;
        public readonly string LatestBackup;
        public readonly string MaintenanceDow;
        public readonly string MaintenanceTime;
        public readonly int MysqlLongQueryTime;
        public readonly bool MysqlRequirePrimaryKey;
        public readonly bool MysqlSlowQueryLog;
        public readonly ImmutableArray<string> MysqlSqlModes;
        public readonly string Password;
        public readonly string Plan;
        public readonly int PlanDisk;
        public readonly int PlanRam;
        public readonly int PlanReplicas;
        public readonly int PlanVcpus;
        public readonly string Port;
        public readonly string RedisEvictionPolicy;
        public readonly string Region;
        public readonly string Status;
        public readonly string Tag;
        public readonly ImmutableArray<string> TrustedIps;
        public readonly string User;

        [OutputConstructor]
        private GetDatabaseReadReplicaResult(
            string clusterTimeZone,

            string databaseEngine,

            string databaseEngineVersion,

            string dateCreated,

            string dbname,

            string host,

            string id,

            string label,

            string latestBackup,

            string maintenanceDow,

            string maintenanceTime,

            int mysqlLongQueryTime,

            bool mysqlRequirePrimaryKey,

            bool mysqlSlowQueryLog,

            ImmutableArray<string> mysqlSqlModes,

            string password,

            string plan,

            int planDisk,

            int planRam,

            int planReplicas,

            int planVcpus,

            string port,

            string redisEvictionPolicy,

            string region,

            string status,

            string tag,

            ImmutableArray<string> trustedIps,

            string user)
        {
            ClusterTimeZone = clusterTimeZone;
            DatabaseEngine = databaseEngine;
            DatabaseEngineVersion = databaseEngineVersion;
            DateCreated = dateCreated;
            Dbname = dbname;
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
            RedisEvictionPolicy = redisEvictionPolicy;
            Region = region;
            Status = status;
            Tag = tag;
            TrustedIps = trustedIps;
            User = user;
        }
    }
}