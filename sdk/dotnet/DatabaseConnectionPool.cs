// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
    /// Provides a Vultr database connection pool resource. This can be used to create, read, modify, and delete connection pools for a PostgreSQL managed database on your Vultr account.
    /// 
    /// ## Example Usage
    /// 
    /// Create a new database connection pool:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vultr = ediri.Vultr;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myDatabaseConnectionPool = new Vultr.DatabaseConnectionPool("myDatabaseConnectionPool", new()
    ///     {
    ///         DatabaseId = vultr_database.My_database.Id,
    ///         Database = "defaultdb",
    ///         Username = "vultradmin",
    ///         Mode = "transaction",
    ///         Size = 3,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [VultrResourceType("vultr:index/databaseConnectionPool:DatabaseConnectionPool")]
    public partial class DatabaseConnectionPool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The logical database to use for the new managed database connection pool.
        /// </summary>
        [Output("database")]
        public Output<string> Database { get; private set; } = null!;

        /// <summary>
        /// The managed database ID you want to attach this connection pool to.
        /// </summary>
        [Output("databaseId")]
        public Output<string> DatabaseId { get; private set; } = null!;

        /// <summary>
        /// The mode to configure for the new managed database connection pool (`session`, `transaction`, `statement`).
        /// </summary>
        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        /// <summary>
        /// The name of the new managed database connection pool.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The size of the new managed database connection pool.
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// The database user to use for the new managed database connection pool.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a DatabaseConnectionPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatabaseConnectionPool(string name, DatabaseConnectionPoolArgs args, CustomResourceOptions? options = null)
            : base("vultr:index/databaseConnectionPool:DatabaseConnectionPool", name, args ?? new DatabaseConnectionPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatabaseConnectionPool(string name, Input<string> id, DatabaseConnectionPoolState? state = null, CustomResourceOptions? options = null)
            : base("vultr:index/databaseConnectionPool:DatabaseConnectionPool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DatabaseConnectionPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatabaseConnectionPool Get(string name, Input<string> id, DatabaseConnectionPoolState? state = null, CustomResourceOptions? options = null)
        {
            return new DatabaseConnectionPool(name, id, state, options);
        }
    }

    public sealed class DatabaseConnectionPoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The logical database to use for the new managed database connection pool.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// The managed database ID you want to attach this connection pool to.
        /// </summary>
        [Input("databaseId", required: true)]
        public Input<string> DatabaseId { get; set; } = null!;

        /// <summary>
        /// The mode to configure for the new managed database connection pool (`session`, `transaction`, `statement`).
        /// </summary>
        [Input("mode", required: true)]
        public Input<string> Mode { get; set; } = null!;

        /// <summary>
        /// The name of the new managed database connection pool.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The size of the new managed database connection pool.
        /// </summary>
        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        /// <summary>
        /// The database user to use for the new managed database connection pool.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public DatabaseConnectionPoolArgs()
        {
        }
        public static new DatabaseConnectionPoolArgs Empty => new DatabaseConnectionPoolArgs();
    }

    public sealed class DatabaseConnectionPoolState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The logical database to use for the new managed database connection pool.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// The managed database ID you want to attach this connection pool to.
        /// </summary>
        [Input("databaseId")]
        public Input<string>? DatabaseId { get; set; }

        /// <summary>
        /// The mode to configure for the new managed database connection pool (`session`, `transaction`, `statement`).
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// The name of the new managed database connection pool.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The size of the new managed database connection pool.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// The database user to use for the new managed database connection pool.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public DatabaseConnectionPoolState()
        {
        }
        public static new DatabaseConnectionPoolState Empty => new DatabaseConnectionPoolState();
    }
}
