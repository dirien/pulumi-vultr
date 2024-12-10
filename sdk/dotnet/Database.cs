// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Vultr
{
    /// <summary>
    /// Provides a Vultr database resource. This can be used to create, read, modify, and delete managed databases on your Vultr account.
    /// 
    /// ## Example Usage
    /// 
    /// Create a new database:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vultr = ediri.Vultr;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myDatabase = new Vultr.Database("myDatabase", new()
    ///     {
    ///         DatabaseEngine = "pg",
    ///         DatabaseEngineVersion = "15",
    ///         Label = "my_database_label",
    ///         Plan = "vultr-dbaas-startup-cc-1-55-2",
    ///         Region = "ewr",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Create a new database with options:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vultr = ediri.Vultr;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myDatabase = new Vultr.Database("myDatabase", new()
    ///     {
    ///         ClusterTimeZone = "America/New_York",
    ///         DatabaseEngine = "pg",
    ///         DatabaseEngineVersion = "15",
    ///         Label = "my_database_label",
    ///         MaintenanceDow = "sunday",
    ///         MaintenanceTime = "01:00",
    ///         Plan = "vultr-dbaas-startup-cc-1-55-2",
    ///         Region = "ewr",
    ///         Tag = "some tag",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Database can be imported using the database `ID`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vultr:index/database:Database my_database b6a859c5-b299-49dd-8888-b1abbc517d08
    /// ```
    /// </summary>
    [VultrResourceType("vultr:index/database:Database")]
    public partial class Database : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The certificate to authenticate the default user (Kafka engine types only).
        /// </summary>
        [Output("accessCert")]
        public Output<string> AccessCert { get; private set; } = null!;

        /// <summary>
        /// The private key to authenticate the default user (Kafka engine types only).
        /// </summary>
        [Output("accessKey")]
        public Output<string> AccessKey { get; private set; } = null!;

        /// <summary>
        /// The configured time zone for the Managed Database in TZ database format (e.g. `UTC`, `America/New_York`, `Europe/London`).
        /// </summary>
        [Output("clusterTimeZone")]
        public Output<string> ClusterTimeZone { get; private set; } = null!;

        /// <summary>
        /// The database engine of the new managed database.
        /// </summary>
        [Output("databaseEngine")]
        public Output<string> DatabaseEngine { get; private set; } = null!;

        /// <summary>
        /// The database engine version of the new managed database.
        /// </summary>
        [Output("databaseEngineVersion")]
        public Output<string> DatabaseEngineVersion { get; private set; } = null!;

        /// <summary>
        /// The date the managed database was added to your Vultr account.
        /// </summary>
        [Output("dateCreated")]
        public Output<string> DateCreated { get; private set; } = null!;

        /// <summary>
        /// The managed database's default logical database.
        /// </summary>
        [Output("dbname")]
        public Output<string> Dbname { get; private set; } = null!;

        /// <summary>
        /// The configuration value for the data eviction policy on the managed database (Redis engine types only - `noeviction`, `allkeys-lru`, `volatile-lru`, `allkeys-random`, `volatile-random`, `volatile-ttl`, `volatile-lfu`, `allkeys-lfu`).
        /// </summary>
        [Output("evictionPolicy")]
        public Output<string> EvictionPolicy { get; private set; } = null!;

        /// <summary>
        /// An associated list of FerretDB connection credentials (FerretDB + PostgreSQL engine types only).
        /// </summary>
        [Output("ferretdbCredentials")]
        public Output<ImmutableDictionary<string, string>> FerretdbCredentials { get; private set; } = null!;

        /// <summary>
        /// The hostname assigned to the managed database.
        /// </summary>
        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        /// <summary>
        /// A label for the managed database.
        /// </summary>
        [Output("label")]
        public Output<string> Label { get; private set; } = null!;

        /// <summary>
        /// The date of the latest backup available on the managed database.
        /// </summary>
        [Output("latestBackup")]
        public Output<string> LatestBackup { get; private set; } = null!;

        /// <summary>
        /// The preferred maintenance day of week for the managed database.
        /// </summary>
        [Output("maintenanceDow")]
        public Output<string> MaintenanceDow { get; private set; } = null!;

        /// <summary>
        /// The preferred maintenance time for the managed database in 24-hour HH:00 format (e.g. `01:00`, `13:00`, `23:00`).
        /// </summary>
        [Output("maintenanceTime")]
        public Output<string> MaintenanceTime { get; private set; } = null!;

        /// <summary>
        /// The configuration value for the long query time (in seconds) on the managed database (MySQL engine types only).
        /// </summary>
        [Output("mysqlLongQueryTime")]
        public Output<int> MysqlLongQueryTime { get; private set; } = null!;

        /// <summary>
        /// The configuration value for whether primary keys are required on the managed database (MySQL engine types only).
        /// </summary>
        [Output("mysqlRequirePrimaryKey")]
        public Output<bool?> MysqlRequirePrimaryKey { get; private set; } = null!;

        /// <summary>
        /// The configuration value for slow query logging on the managed database (MySQL engine types only).
        /// </summary>
        [Output("mysqlSlowQueryLog")]
        public Output<bool> MysqlSlowQueryLog { get; private set; } = null!;

        /// <summary>
        /// A list of SQL modes to configure for the managed database (MySQL engine types only - `ALLOW_INVALID_DATES`, `ANSI`, `ANSI_QUOTES`, `ERROR_FOR_DIVISION_BY_ZERO`, `HIGH_NOT_PRECEDENCE`, `IGNORE_SPACE`, `NO_AUTO_VALUE_ON_ZERO`, `NO_DIR_IN_CREATE`, `NO_ENGINE_SUBSTITUTION`, `NO_UNSIGNED_SUBTRACTION`, `NO_ZERO_DATE`, `NO_ZERO_IN_DATE`, `ONLY_FULL_GROUP_BY`, `PIPES_AS_CONCAT`, `REAL_AS_FLOAT`, `STRICT_ALL_TABLES`, `STRICT_TRANS_TABLES`, `TIME_TRUNCATE_FRACTIONAL`, `TRADITIONAL`).
        /// </summary>
        [Output("mysqlSqlModes")]
        public Output<ImmutableArray<string>> MysqlSqlModes { get; private set; } = null!;

        /// <summary>
        /// The password for the managed database's primary admin user.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// The ID of the plan that you want the managed database to subscribe to. [See List Managed Database Plans](https://www.vultr.com/api/#tag/managed-databases/operation/list-database-plans)
        /// </summary>
        [Output("plan")]
        public Output<string> Plan { get; private set; } = null!;

        /// <summary>
        /// The number of brokers available on the managed database (Kafka engine types only).
        /// </summary>
        [Output("planBrokers")]
        public Output<int> PlanBrokers { get; private set; } = null!;

        /// <summary>
        /// The description of the disk(s) on the managed database.
        /// </summary>
        [Output("planDisk")]
        public Output<int> PlanDisk { get; private set; } = null!;

        /// <summary>
        /// The amount of memory available on the managed database in MB.
        /// </summary>
        [Output("planRam")]
        public Output<int> PlanRam { get; private set; } = null!;

        /// <summary>
        /// The number of standby nodes available on the managed database (excluded for Kafka engine types).
        /// </summary>
        [Output("planReplicas")]
        public Output<int> PlanReplicas { get; private set; } = null!;

        /// <summary>
        /// The number of virtual CPUs available on the managed database.
        /// </summary>
        [Output("planVcpus")]
        public Output<int> PlanVcpus { get; private set; } = null!;

        /// <summary>
        /// The connection port for the managed database.
        /// </summary>
        [Output("port")]
        public Output<string> Port { get; private set; } = null!;

        /// <summary>
        /// The public hostname assigned to the managed database (VPC-attached only).
        /// </summary>
        [Output("publicHost")]
        public Output<string> PublicHost { get; private set; } = null!;

        /// <summary>
        /// A list of read replicas attached to the managed database.
        /// </summary>
        [Output("readReplicas")]
        public Output<ImmutableArray<Outputs.DatabaseReadReplica>> ReadReplicas { get; private set; } = null!;

        /// <summary>
        /// The ID of the region that the managed database is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The SASL connection port for the managed database (Kafka engine types only).
        /// </summary>
        [Output("saslPort")]
        public Output<string> SaslPort { get; private set; } = null!;

        /// <summary>
        /// The current status of the managed database (poweroff, rebuilding, rebalancing, configuring, running).
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The tag to assign to the managed database.
        /// </summary>
        [Output("tag")]
        public Output<string?> Tag { get; private set; } = null!;

        /// <summary>
        /// A list of allowed IP addresses for the managed database.
        /// </summary>
        [Output("trustedIps")]
        public Output<ImmutableArray<string>> TrustedIps { get; private set; } = null!;

        /// <summary>
        /// The primary admin user for the managed database.
        /// </summary>
        [Output("user")]
        public Output<string> User { get; private set; } = null!;

        /// <summary>
        /// The ID of the VPC Network to attach to the Managed Database.
        /// </summary>
        [Output("vpcId")]
        public Output<string?> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a Database resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Database(string name, DatabaseArgs args, CustomResourceOptions? options = null)
            : base("vultr:index/database:Database", name, args ?? new DatabaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Database(string name, Input<string> id, DatabaseState? state = null, CustomResourceOptions? options = null)
            : base("vultr:index/database:Database", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/dirien/pulumi-vultr",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Database resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Database Get(string name, Input<string> id, DatabaseState? state = null, CustomResourceOptions? options = null)
        {
            return new Database(name, id, state, options);
        }
    }

    public sealed class DatabaseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The certificate to authenticate the default user (Kafka engine types only).
        /// </summary>
        [Input("accessCert")]
        public Input<string>? AccessCert { get; set; }

        /// <summary>
        /// The private key to authenticate the default user (Kafka engine types only).
        /// </summary>
        [Input("accessKey")]
        public Input<string>? AccessKey { get; set; }

        /// <summary>
        /// The configured time zone for the Managed Database in TZ database format (e.g. `UTC`, `America/New_York`, `Europe/London`).
        /// </summary>
        [Input("clusterTimeZone")]
        public Input<string>? ClusterTimeZone { get; set; }

        /// <summary>
        /// The database engine of the new managed database.
        /// </summary>
        [Input("databaseEngine", required: true)]
        public Input<string> DatabaseEngine { get; set; } = null!;

        /// <summary>
        /// The database engine version of the new managed database.
        /// </summary>
        [Input("databaseEngineVersion", required: true)]
        public Input<string> DatabaseEngineVersion { get; set; } = null!;

        /// <summary>
        /// The configuration value for the data eviction policy on the managed database (Redis engine types only - `noeviction`, `allkeys-lru`, `volatile-lru`, `allkeys-random`, `volatile-random`, `volatile-ttl`, `volatile-lfu`, `allkeys-lfu`).
        /// </summary>
        [Input("evictionPolicy")]
        public Input<string>? EvictionPolicy { get; set; }

        [Input("ferretdbCredentials")]
        private InputMap<string>? _ferretdbCredentials;

        /// <summary>
        /// An associated list of FerretDB connection credentials (FerretDB + PostgreSQL engine types only).
        /// </summary>
        public InputMap<string> FerretdbCredentials
        {
            get => _ferretdbCredentials ?? (_ferretdbCredentials = new InputMap<string>());
            set => _ferretdbCredentials = value;
        }

        /// <summary>
        /// A label for the managed database.
        /// </summary>
        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        /// <summary>
        /// The preferred maintenance day of week for the managed database.
        /// </summary>
        [Input("maintenanceDow")]
        public Input<string>? MaintenanceDow { get; set; }

        /// <summary>
        /// The preferred maintenance time for the managed database in 24-hour HH:00 format (e.g. `01:00`, `13:00`, `23:00`).
        /// </summary>
        [Input("maintenanceTime")]
        public Input<string>? MaintenanceTime { get; set; }

        /// <summary>
        /// The configuration value for the long query time (in seconds) on the managed database (MySQL engine types only).
        /// </summary>
        [Input("mysqlLongQueryTime")]
        public Input<int>? MysqlLongQueryTime { get; set; }

        /// <summary>
        /// The configuration value for whether primary keys are required on the managed database (MySQL engine types only).
        /// </summary>
        [Input("mysqlRequirePrimaryKey")]
        public Input<bool>? MysqlRequirePrimaryKey { get; set; }

        /// <summary>
        /// The configuration value for slow query logging on the managed database (MySQL engine types only).
        /// </summary>
        [Input("mysqlSlowQueryLog")]
        public Input<bool>? MysqlSlowQueryLog { get; set; }

        [Input("mysqlSqlModes")]
        private InputList<string>? _mysqlSqlModes;

        /// <summary>
        /// A list of SQL modes to configure for the managed database (MySQL engine types only - `ALLOW_INVALID_DATES`, `ANSI`, `ANSI_QUOTES`, `ERROR_FOR_DIVISION_BY_ZERO`, `HIGH_NOT_PRECEDENCE`, `IGNORE_SPACE`, `NO_AUTO_VALUE_ON_ZERO`, `NO_DIR_IN_CREATE`, `NO_ENGINE_SUBSTITUTION`, `NO_UNSIGNED_SUBTRACTION`, `NO_ZERO_DATE`, `NO_ZERO_IN_DATE`, `ONLY_FULL_GROUP_BY`, `PIPES_AS_CONCAT`, `REAL_AS_FLOAT`, `STRICT_ALL_TABLES`, `STRICT_TRANS_TABLES`, `TIME_TRUNCATE_FRACTIONAL`, `TRADITIONAL`).
        /// </summary>
        public InputList<string> MysqlSqlModes
        {
            get => _mysqlSqlModes ?? (_mysqlSqlModes = new InputList<string>());
            set => _mysqlSqlModes = value;
        }

        /// <summary>
        /// The password for the managed database's primary admin user.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The ID of the plan that you want the managed database to subscribe to. [See List Managed Database Plans](https://www.vultr.com/api/#tag/managed-databases/operation/list-database-plans)
        /// </summary>
        [Input("plan", required: true)]
        public Input<string> Plan { get; set; } = null!;

        /// <summary>
        /// The number of brokers available on the managed database (Kafka engine types only).
        /// </summary>
        [Input("planBrokers")]
        public Input<int>? PlanBrokers { get; set; }

        /// <summary>
        /// The description of the disk(s) on the managed database.
        /// </summary>
        [Input("planDisk")]
        public Input<int>? PlanDisk { get; set; }

        /// <summary>
        /// The number of standby nodes available on the managed database (excluded for Kafka engine types).
        /// </summary>
        [Input("planReplicas")]
        public Input<int>? PlanReplicas { get; set; }

        /// <summary>
        /// The public hostname assigned to the managed database (VPC-attached only).
        /// </summary>
        [Input("publicHost")]
        public Input<string>? PublicHost { get; set; }

        [Input("readReplicas")]
        private InputList<Inputs.DatabaseReadReplicaArgs>? _readReplicas;

        /// <summary>
        /// A list of read replicas attached to the managed database.
        /// </summary>
        public InputList<Inputs.DatabaseReadReplicaArgs> ReadReplicas
        {
            get => _readReplicas ?? (_readReplicas = new InputList<Inputs.DatabaseReadReplicaArgs>());
            set => _readReplicas = value;
        }

        /// <summary>
        /// The ID of the region that the managed database is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// The SASL connection port for the managed database (Kafka engine types only).
        /// </summary>
        [Input("saslPort")]
        public Input<string>? SaslPort { get; set; }

        /// <summary>
        /// The tag to assign to the managed database.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        [Input("trustedIps")]
        private InputList<string>? _trustedIps;

        /// <summary>
        /// A list of allowed IP addresses for the managed database.
        /// </summary>
        public InputList<string> TrustedIps
        {
            get => _trustedIps ?? (_trustedIps = new InputList<string>());
            set => _trustedIps = value;
        }

        /// <summary>
        /// The ID of the VPC Network to attach to the Managed Database.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public DatabaseArgs()
        {
        }
        public static new DatabaseArgs Empty => new DatabaseArgs();
    }

    public sealed class DatabaseState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The certificate to authenticate the default user (Kafka engine types only).
        /// </summary>
        [Input("accessCert")]
        public Input<string>? AccessCert { get; set; }

        /// <summary>
        /// The private key to authenticate the default user (Kafka engine types only).
        /// </summary>
        [Input("accessKey")]
        public Input<string>? AccessKey { get; set; }

        /// <summary>
        /// The configured time zone for the Managed Database in TZ database format (e.g. `UTC`, `America/New_York`, `Europe/London`).
        /// </summary>
        [Input("clusterTimeZone")]
        public Input<string>? ClusterTimeZone { get; set; }

        /// <summary>
        /// The database engine of the new managed database.
        /// </summary>
        [Input("databaseEngine")]
        public Input<string>? DatabaseEngine { get; set; }

        /// <summary>
        /// The database engine version of the new managed database.
        /// </summary>
        [Input("databaseEngineVersion")]
        public Input<string>? DatabaseEngineVersion { get; set; }

        /// <summary>
        /// The date the managed database was added to your Vultr account.
        /// </summary>
        [Input("dateCreated")]
        public Input<string>? DateCreated { get; set; }

        /// <summary>
        /// The managed database's default logical database.
        /// </summary>
        [Input("dbname")]
        public Input<string>? Dbname { get; set; }

        /// <summary>
        /// The configuration value for the data eviction policy on the managed database (Redis engine types only - `noeviction`, `allkeys-lru`, `volatile-lru`, `allkeys-random`, `volatile-random`, `volatile-ttl`, `volatile-lfu`, `allkeys-lfu`).
        /// </summary>
        [Input("evictionPolicy")]
        public Input<string>? EvictionPolicy { get; set; }

        [Input("ferretdbCredentials")]
        private InputMap<string>? _ferretdbCredentials;

        /// <summary>
        /// An associated list of FerretDB connection credentials (FerretDB + PostgreSQL engine types only).
        /// </summary>
        public InputMap<string> FerretdbCredentials
        {
            get => _ferretdbCredentials ?? (_ferretdbCredentials = new InputMap<string>());
            set => _ferretdbCredentials = value;
        }

        /// <summary>
        /// The hostname assigned to the managed database.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// A label for the managed database.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// The date of the latest backup available on the managed database.
        /// </summary>
        [Input("latestBackup")]
        public Input<string>? LatestBackup { get; set; }

        /// <summary>
        /// The preferred maintenance day of week for the managed database.
        /// </summary>
        [Input("maintenanceDow")]
        public Input<string>? MaintenanceDow { get; set; }

        /// <summary>
        /// The preferred maintenance time for the managed database in 24-hour HH:00 format (e.g. `01:00`, `13:00`, `23:00`).
        /// </summary>
        [Input("maintenanceTime")]
        public Input<string>? MaintenanceTime { get; set; }

        /// <summary>
        /// The configuration value for the long query time (in seconds) on the managed database (MySQL engine types only).
        /// </summary>
        [Input("mysqlLongQueryTime")]
        public Input<int>? MysqlLongQueryTime { get; set; }

        /// <summary>
        /// The configuration value for whether primary keys are required on the managed database (MySQL engine types only).
        /// </summary>
        [Input("mysqlRequirePrimaryKey")]
        public Input<bool>? MysqlRequirePrimaryKey { get; set; }

        /// <summary>
        /// The configuration value for slow query logging on the managed database (MySQL engine types only).
        /// </summary>
        [Input("mysqlSlowQueryLog")]
        public Input<bool>? MysqlSlowQueryLog { get; set; }

        [Input("mysqlSqlModes")]
        private InputList<string>? _mysqlSqlModes;

        /// <summary>
        /// A list of SQL modes to configure for the managed database (MySQL engine types only - `ALLOW_INVALID_DATES`, `ANSI`, `ANSI_QUOTES`, `ERROR_FOR_DIVISION_BY_ZERO`, `HIGH_NOT_PRECEDENCE`, `IGNORE_SPACE`, `NO_AUTO_VALUE_ON_ZERO`, `NO_DIR_IN_CREATE`, `NO_ENGINE_SUBSTITUTION`, `NO_UNSIGNED_SUBTRACTION`, `NO_ZERO_DATE`, `NO_ZERO_IN_DATE`, `ONLY_FULL_GROUP_BY`, `PIPES_AS_CONCAT`, `REAL_AS_FLOAT`, `STRICT_ALL_TABLES`, `STRICT_TRANS_TABLES`, `TIME_TRUNCATE_FRACTIONAL`, `TRADITIONAL`).
        /// </summary>
        public InputList<string> MysqlSqlModes
        {
            get => _mysqlSqlModes ?? (_mysqlSqlModes = new InputList<string>());
            set => _mysqlSqlModes = value;
        }

        /// <summary>
        /// The password for the managed database's primary admin user.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The ID of the plan that you want the managed database to subscribe to. [See List Managed Database Plans](https://www.vultr.com/api/#tag/managed-databases/operation/list-database-plans)
        /// </summary>
        [Input("plan")]
        public Input<string>? Plan { get; set; }

        /// <summary>
        /// The number of brokers available on the managed database (Kafka engine types only).
        /// </summary>
        [Input("planBrokers")]
        public Input<int>? PlanBrokers { get; set; }

        /// <summary>
        /// The description of the disk(s) on the managed database.
        /// </summary>
        [Input("planDisk")]
        public Input<int>? PlanDisk { get; set; }

        /// <summary>
        /// The amount of memory available on the managed database in MB.
        /// </summary>
        [Input("planRam")]
        public Input<int>? PlanRam { get; set; }

        /// <summary>
        /// The number of standby nodes available on the managed database (excluded for Kafka engine types).
        /// </summary>
        [Input("planReplicas")]
        public Input<int>? PlanReplicas { get; set; }

        /// <summary>
        /// The number of virtual CPUs available on the managed database.
        /// </summary>
        [Input("planVcpus")]
        public Input<int>? PlanVcpus { get; set; }

        /// <summary>
        /// The connection port for the managed database.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// The public hostname assigned to the managed database (VPC-attached only).
        /// </summary>
        [Input("publicHost")]
        public Input<string>? PublicHost { get; set; }

        [Input("readReplicas")]
        private InputList<Inputs.DatabaseReadReplicaGetArgs>? _readReplicas;

        /// <summary>
        /// A list of read replicas attached to the managed database.
        /// </summary>
        public InputList<Inputs.DatabaseReadReplicaGetArgs> ReadReplicas
        {
            get => _readReplicas ?? (_readReplicas = new InputList<Inputs.DatabaseReadReplicaGetArgs>());
            set => _readReplicas = value;
        }

        /// <summary>
        /// The ID of the region that the managed database is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The SASL connection port for the managed database (Kafka engine types only).
        /// </summary>
        [Input("saslPort")]
        public Input<string>? SaslPort { get; set; }

        /// <summary>
        /// The current status of the managed database (poweroff, rebuilding, rebalancing, configuring, running).
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The tag to assign to the managed database.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        [Input("trustedIps")]
        private InputList<string>? _trustedIps;

        /// <summary>
        /// A list of allowed IP addresses for the managed database.
        /// </summary>
        public InputList<string> TrustedIps
        {
            get => _trustedIps ?? (_trustedIps = new InputList<string>());
            set => _trustedIps = value;
        }

        /// <summary>
        /// The primary admin user for the managed database.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        /// <summary>
        /// The ID of the VPC Network to attach to the Managed Database.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public DatabaseState()
        {
        }
        public static new DatabaseState Empty => new DatabaseState();
    }
}
