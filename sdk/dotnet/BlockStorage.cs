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
    /// Provides a Vultr Block Storage resource. This can be used to create, read, modify, and delete Block Storage.
    /// 
    /// ## Example Usage
    /// 
    /// Create a new Block Storage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vultr = ediri.Vultr;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myBlockstorage = new Vultr.BlockStorage("myBlockstorage", new()
    ///     {
    ///         Label = "vultr-block-storage",
    ///         Region = "ewr",
    ///         SizeGb = 10,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Block Storage can be imported using the Block Storage `ID`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vultr:index/blockStorage:BlockStorage my_blockstorage e315835e-d466-4e89-9b4c-dfd8788d7685
    /// ```
    /// </summary>
    [VultrResourceType("vultr:index/blockStorage:BlockStorage")]
    public partial class BlockStorage : global::Pulumi.CustomResource
    {
        /// <summary>
        /// VPS ID that you want to have this block storage attached to.
        /// </summary>
        [Output("attachedToInstance")]
        public Output<string?> AttachedToInstance { get; private set; } = null!;

        /// <summary>
        /// Determines on the type of block storage volume that will be created. Soon to become a required parameter. Options are `high_perf` or `storage_opt`.
        /// </summary>
        [Output("blockType")]
        public Output<string> BlockType { get; private set; } = null!;

        /// <summary>
        /// The monthly cost of this block storage.
        /// </summary>
        [Output("cost")]
        public Output<double> Cost { get; private set; } = null!;

        /// <summary>
        /// The date this block storage was created.
        /// </summary>
        [Output("dateCreated")]
        public Output<string> DateCreated { get; private set; } = null!;

        /// <summary>
        /// Label that is given to your block storage.
        /// </summary>
        [Output("label")]
        public Output<string?> Label { get; private set; } = null!;

        /// <summary>
        /// Boolean value that will allow attachment of the volume to an instance without a restart. Default is false.
        /// </summary>
        [Output("live")]
        public Output<bool?> Live { get; private set; } = null!;

        /// <summary>
        /// An ID associated with the instance, when mounted the ID can be found in /dev/disk/by-id prefixed with virtio.
        /// </summary>
        [Output("mountId")]
        public Output<string> MountId { get; private set; } = null!;

        /// <summary>
        /// Region in which this block storage will reside in. (Currently only NJ/NY supported region "ewr")
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The size of the given block storage.
        /// </summary>
        [Output("sizeGb")]
        public Output<int> SizeGb { get; private set; } = null!;

        /// <summary>
        /// Current status of your block storage.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a BlockStorage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BlockStorage(string name, BlockStorageArgs args, CustomResourceOptions? options = null)
            : base("vultr:index/blockStorage:BlockStorage", name, args ?? new BlockStorageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BlockStorage(string name, Input<string> id, BlockStorageState? state = null, CustomResourceOptions? options = null)
            : base("vultr:index/blockStorage:BlockStorage", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BlockStorage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BlockStorage Get(string name, Input<string> id, BlockStorageState? state = null, CustomResourceOptions? options = null)
        {
            return new BlockStorage(name, id, state, options);
        }
    }

    public sealed class BlockStorageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// VPS ID that you want to have this block storage attached to.
        /// </summary>
        [Input("attachedToInstance")]
        public Input<string>? AttachedToInstance { get; set; }

        /// <summary>
        /// Determines on the type of block storage volume that will be created. Soon to become a required parameter. Options are `high_perf` or `storage_opt`.
        /// </summary>
        [Input("blockType")]
        public Input<string>? BlockType { get; set; }

        /// <summary>
        /// Label that is given to your block storage.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// Boolean value that will allow attachment of the volume to an instance without a restart. Default is false.
        /// </summary>
        [Input("live")]
        public Input<bool>? Live { get; set; }

        /// <summary>
        /// Region in which this block storage will reside in. (Currently only NJ/NY supported region "ewr")
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// The size of the given block storage.
        /// </summary>
        [Input("sizeGb", required: true)]
        public Input<int> SizeGb { get; set; } = null!;

        public BlockStorageArgs()
        {
        }
        public static new BlockStorageArgs Empty => new BlockStorageArgs();
    }

    public sealed class BlockStorageState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// VPS ID that you want to have this block storage attached to.
        /// </summary>
        [Input("attachedToInstance")]
        public Input<string>? AttachedToInstance { get; set; }

        /// <summary>
        /// Determines on the type of block storage volume that will be created. Soon to become a required parameter. Options are `high_perf` or `storage_opt`.
        /// </summary>
        [Input("blockType")]
        public Input<string>? BlockType { get; set; }

        /// <summary>
        /// The monthly cost of this block storage.
        /// </summary>
        [Input("cost")]
        public Input<double>? Cost { get; set; }

        /// <summary>
        /// The date this block storage was created.
        /// </summary>
        [Input("dateCreated")]
        public Input<string>? DateCreated { get; set; }

        /// <summary>
        /// Label that is given to your block storage.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// Boolean value that will allow attachment of the volume to an instance without a restart. Default is false.
        /// </summary>
        [Input("live")]
        public Input<bool>? Live { get; set; }

        /// <summary>
        /// An ID associated with the instance, when mounted the ID can be found in /dev/disk/by-id prefixed with virtio.
        /// </summary>
        [Input("mountId")]
        public Input<string>? MountId { get; set; }

        /// <summary>
        /// Region in which this block storage will reside in. (Currently only NJ/NY supported region "ewr")
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The size of the given block storage.
        /// </summary>
        [Input("sizeGb")]
        public Input<int>? SizeGb { get; set; }

        /// <summary>
        /// Current status of your block storage.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public BlockStorageState()
        {
        }
        public static new BlockStorageState Empty => new BlockStorageState();
    }
}
