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
    /// Provides a Vultr database quota resource. This can be used to create, read, and delete quotas for a managed database on your Vultr account.
    /// 
    /// ## Example Usage
    /// 
    /// Create a new database quota:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vultr = ediri.Vultr;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myDatabaseQuota = new Vultr.DatabaseQuota("myDatabaseQuota", new()
    ///     {
    ///         DatabaseId = vultr_database.My_database.Id,
    ///         ClientId = "my_database_quota",
    ///         ConsumerByteRate = 3,
    ///         ProducerByteRate = 2,
    ///         RequestPercentage = 120,
    ///         User = "my_database_user",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [VultrResourceType("vultr:index/databaseQuota:DatabaseQuota")]
    public partial class DatabaseQuota : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The client ID for the new database quota.
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// The consumer byte rate for the new managed database quota.
        /// </summary>
        [Output("consumerByteRate")]
        public Output<int> ConsumerByteRate { get; private set; } = null!;

        /// <summary>
        /// The managed database ID you want to attach this quota to.
        /// </summary>
        [Output("databaseId")]
        public Output<string> DatabaseId { get; private set; } = null!;

        /// <summary>
        /// The producer byte rate for the new managed database quota.
        /// </summary>
        [Output("producerByteRate")]
        public Output<int> ProducerByteRate { get; private set; } = null!;

        /// <summary>
        /// The CPU request percentage for the new managed database quota.
        /// </summary>
        [Output("requestPercentage")]
        public Output<int> RequestPercentage { get; private set; } = null!;

        /// <summary>
        /// The user for the new managed database quota.
        /// </summary>
        [Output("user")]
        public Output<string> User { get; private set; } = null!;


        /// <summary>
        /// Create a DatabaseQuota resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatabaseQuota(string name, DatabaseQuotaArgs args, CustomResourceOptions? options = null)
            : base("vultr:index/databaseQuota:DatabaseQuota", name, args ?? new DatabaseQuotaArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatabaseQuota(string name, Input<string> id, DatabaseQuotaState? state = null, CustomResourceOptions? options = null)
            : base("vultr:index/databaseQuota:DatabaseQuota", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DatabaseQuota resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatabaseQuota Get(string name, Input<string> id, DatabaseQuotaState? state = null, CustomResourceOptions? options = null)
        {
            return new DatabaseQuota(name, id, state, options);
        }
    }

    public sealed class DatabaseQuotaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The client ID for the new database quota.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// The consumer byte rate for the new managed database quota.
        /// </summary>
        [Input("consumerByteRate", required: true)]
        public Input<int> ConsumerByteRate { get; set; } = null!;

        /// <summary>
        /// The managed database ID you want to attach this quota to.
        /// </summary>
        [Input("databaseId", required: true)]
        public Input<string> DatabaseId { get; set; } = null!;

        /// <summary>
        /// The producer byte rate for the new managed database quota.
        /// </summary>
        [Input("producerByteRate", required: true)]
        public Input<int> ProducerByteRate { get; set; } = null!;

        /// <summary>
        /// The CPU request percentage for the new managed database quota.
        /// </summary>
        [Input("requestPercentage", required: true)]
        public Input<int> RequestPercentage { get; set; } = null!;

        /// <summary>
        /// The user for the new managed database quota.
        /// </summary>
        [Input("user", required: true)]
        public Input<string> User { get; set; } = null!;

        public DatabaseQuotaArgs()
        {
        }
        public static new DatabaseQuotaArgs Empty => new DatabaseQuotaArgs();
    }

    public sealed class DatabaseQuotaState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The client ID for the new database quota.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The consumer byte rate for the new managed database quota.
        /// </summary>
        [Input("consumerByteRate")]
        public Input<int>? ConsumerByteRate { get; set; }

        /// <summary>
        /// The managed database ID you want to attach this quota to.
        /// </summary>
        [Input("databaseId")]
        public Input<string>? DatabaseId { get; set; }

        /// <summary>
        /// The producer byte rate for the new managed database quota.
        /// </summary>
        [Input("producerByteRate")]
        public Input<int>? ProducerByteRate { get; set; }

        /// <summary>
        /// The CPU request percentage for the new managed database quota.
        /// </summary>
        [Input("requestPercentage")]
        public Input<int>? RequestPercentage { get; set; }

        /// <summary>
        /// The user for the new managed database quota.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public DatabaseQuotaState()
        {
        }
        public static new DatabaseQuotaState Empty => new DatabaseQuotaState();
    }
}
