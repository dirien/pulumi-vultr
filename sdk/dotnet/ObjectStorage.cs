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
    /// Provides a Vultr private object storage resource. This can be used to create, read, update and delete object storage resources on your Vultr account.
    /// 
    /// ## Example Usage
    /// 
    /// Create a new object storage subscription.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vultr = ediri.Vultr;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var tf = new Vultr.ObjectStorage("tf", new()
    ///     {
    ///         ClusterId = 2,
    ///         Label = "tf-label",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Object Storage can be imported using the object storage `ID`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import vultr:index/objectStorage:ObjectStorage my_s3 0e04f918-575e-41cb-86f6-d729b354a5a1
    /// ```
    /// </summary>
    [VultrResourceType("vultr:index/objectStorage:ObjectStorage")]
    public partial class ObjectStorage : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The region ID that you want the network to be created in.
        /// </summary>
        [Output("clusterId")]
        public Output<int> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Date of creation for the object storage subscription.
        /// </summary>
        [Output("dateCreated")]
        public Output<string> DateCreated { get; private set; } = null!;

        /// <summary>
        /// The description you want to give your network.
        /// </summary>
        [Output("label")]
        public Output<string?> Label { get; private set; } = null!;

        /// <summary>
        /// The location which this subscription resides in.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The region ID of the object storage subscription.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Your access key.
        /// </summary>
        [Output("s3AccessKey")]
        public Output<string> S3AccessKey { get; private set; } = null!;

        /// <summary>
        /// The hostname for this subscription.
        /// </summary>
        [Output("s3Hostname")]
        public Output<string> S3Hostname { get; private set; } = null!;

        /// <summary>
        /// Your secret key.
        /// </summary>
        [Output("s3SecretKey")]
        public Output<string> S3SecretKey { get; private set; } = null!;

        /// <summary>
        /// Current status of this object storage subscription.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a ObjectStorage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ObjectStorage(string name, ObjectStorageArgs args, CustomResourceOptions? options = null)
            : base("vultr:index/objectStorage:ObjectStorage", name, args ?? new ObjectStorageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ObjectStorage(string name, Input<string> id, ObjectStorageState? state = null, CustomResourceOptions? options = null)
            : base("vultr:index/objectStorage:ObjectStorage", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/dirien/pulumi-vultr",
                AdditionalSecretOutputs =
                {
                    "s3AccessKey",
                    "s3SecretKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ObjectStorage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ObjectStorage Get(string name, Input<string> id, ObjectStorageState? state = null, CustomResourceOptions? options = null)
        {
            return new ObjectStorage(name, id, state, options);
        }
    }

    public sealed class ObjectStorageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The region ID that you want the network to be created in.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<int> ClusterId { get; set; } = null!;

        /// <summary>
        /// The description you want to give your network.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        public ObjectStorageArgs()
        {
        }
        public static new ObjectStorageArgs Empty => new ObjectStorageArgs();
    }

    public sealed class ObjectStorageState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The region ID that you want the network to be created in.
        /// </summary>
        [Input("clusterId")]
        public Input<int>? ClusterId { get; set; }

        /// <summary>
        /// Date of creation for the object storage subscription.
        /// </summary>
        [Input("dateCreated")]
        public Input<string>? DateCreated { get; set; }

        /// <summary>
        /// The description you want to give your network.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// The location which this subscription resides in.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The region ID of the object storage subscription.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("s3AccessKey")]
        private Input<string>? _s3AccessKey;

        /// <summary>
        /// Your access key.
        /// </summary>
        public Input<string>? S3AccessKey
        {
            get => _s3AccessKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _s3AccessKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The hostname for this subscription.
        /// </summary>
        [Input("s3Hostname")]
        public Input<string>? S3Hostname { get; set; }

        [Input("s3SecretKey")]
        private Input<string>? _s3SecretKey;

        /// <summary>
        /// Your secret key.
        /// </summary>
        public Input<string>? S3SecretKey
        {
            get => _s3SecretKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _s3SecretKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Current status of this object storage subscription.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ObjectStorageState()
        {
        }
        public static new ObjectStorageState Empty => new ObjectStorageState();
    }
}
