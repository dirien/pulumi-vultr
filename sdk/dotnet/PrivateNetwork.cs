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
    /// Deprecated: Use `vultr.Vpc` instead
    /// 
    /// Provides a Vultr private network resource. This can be used to create, read, and delete private networks on your Vultr account.
    /// 
    /// ## Example Usage
    /// 
    /// Create a new private network with an automatically generated CIDR block:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vultr = ediri.Vultr;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myNetwork = new Vultr.PrivateNetwork("myNetwork", new()
    ///     {
    ///         Description = "my private network",
    ///         Region = "ewr",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Create a new private network with a specified CIDR block:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vultr = ediri.Vultr;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myNetwork = new Vultr.PrivateNetwork("myNetwork", new()
    ///     {
    ///         Description = "my private network",
    ///         Region = "ewr",
    ///         V4Subnet = "10.0.0.0",
    ///         V4SubnetMask = 24,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Networks can be imported using the network `ID`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vultr:index/privateNetwork:PrivateNetwork my_network 0e04f918-575e-41cb-86f6-d729b354a5a1
    /// ```
    /// </summary>
    [VultrResourceType("vultr:index/privateNetwork:PrivateNetwork")]
    public partial class PrivateNetwork : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The date that the network was added to your Vultr account.
        /// </summary>
        [Output("dateCreated")]
        public Output<string> DateCreated { get; private set; } = null!;

        /// <summary>
        /// The description you want to give your network.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The region ID that you want the network to be created in.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The IPv4 subnet to be used when attaching instances to this network.
        /// </summary>
        [Output("v4Subnet")]
        public Output<string> V4Subnet { get; private set; } = null!;

        /// <summary>
        /// The number of bits for the netmask in CIDR notation. Example: 32
        /// </summary>
        [Output("v4SubnetMask")]
        public Output<int> V4SubnetMask { get; private set; } = null!;


        /// <summary>
        /// Create a PrivateNetwork resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateNetwork(string name, PrivateNetworkArgs args, CustomResourceOptions? options = null)
            : base("vultr:index/privateNetwork:PrivateNetwork", name, args ?? new PrivateNetworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateNetwork(string name, Input<string> id, PrivateNetworkState? state = null, CustomResourceOptions? options = null)
            : base("vultr:index/privateNetwork:PrivateNetwork", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PrivateNetwork resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateNetwork Get(string name, Input<string> id, PrivateNetworkState? state = null, CustomResourceOptions? options = null)
        {
            return new PrivateNetwork(name, id, state, options);
        }
    }

    public sealed class PrivateNetworkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description you want to give your network.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The region ID that you want the network to be created in.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// The IPv4 subnet to be used when attaching instances to this network.
        /// </summary>
        [Input("v4Subnet")]
        public Input<string>? V4Subnet { get; set; }

        /// <summary>
        /// The number of bits for the netmask in CIDR notation. Example: 32
        /// </summary>
        [Input("v4SubnetMask")]
        public Input<int>? V4SubnetMask { get; set; }

        public PrivateNetworkArgs()
        {
        }
        public static new PrivateNetworkArgs Empty => new PrivateNetworkArgs();
    }

    public sealed class PrivateNetworkState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date that the network was added to your Vultr account.
        /// </summary>
        [Input("dateCreated")]
        public Input<string>? DateCreated { get; set; }

        /// <summary>
        /// The description you want to give your network.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The region ID that you want the network to be created in.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The IPv4 subnet to be used when attaching instances to this network.
        /// </summary>
        [Input("v4Subnet")]
        public Input<string>? V4Subnet { get; set; }

        /// <summary>
        /// The number of bits for the netmask in CIDR notation. Example: 32
        /// </summary>
        [Input("v4SubnetMask")]
        public Input<int>? V4SubnetMask { get; set; }

        public PrivateNetworkState()
        {
        }
        public static new PrivateNetworkState Empty => new PrivateNetworkState();
    }
}
