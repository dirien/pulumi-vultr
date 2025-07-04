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
    /// Provides a Vultr instance IPv4 resource. This can be used to create, read, and
    /// modify a IPv4 address. instance is rebooted by default.
    /// 
    /// ## Example Usage
    /// 
    /// Create a new IPv4 address for a instance:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vultr = ediri.Vultr;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myInstance = new Vultr.Instance("myInstance", new()
    ///     {
    ///         EnableIpv6 = true,
    ///         OsId = 167,
    ///         Plan = "vc2-1c-1gb",
    ///         Region = "ewr",
    ///     });
    /// 
    ///     var myInstanceIpv4 = new Vultr.InstanceIpv4("myInstanceIpv4", new()
    ///     {
    ///         InstanceId = myInstance.Id,
    ///         Reboot = false,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [VultrResourceType("vultr:index/instanceIpv4:InstanceIpv4")]
    public partial class InstanceIpv4 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The gateway IP address.
        /// </summary>
        [Output("gateway")]
        public Output<string> Gateway { get; private set; } = null!;

        /// <summary>
        /// The ID of the instance to be assigned the IPv4 address.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The IPv4 address in canonical format.
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// The IPv4 netmask in dot-decimal notation.
        /// </summary>
        [Output("netmask")]
        public Output<string> Netmask { get; private set; } = null!;

        /// <summary>
        /// Default true. Determines whether or not the server is rebooted after adding the IPv4 address.
        /// </summary>
        [Output("reboot")]
        public Output<bool?> Reboot { get; private set; } = null!;

        /// <summary>
        /// The reverse DNS information for this IP address.
        /// </summary>
        [Output("reverse")]
        public Output<string> Reverse { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceIpv4 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceIpv4(string name, InstanceIpv4Args args, CustomResourceOptions? options = null)
            : base("vultr:index/instanceIpv4:InstanceIpv4", name, args ?? new InstanceIpv4Args(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceIpv4(string name, Input<string> id, InstanceIpv4State? state = null, CustomResourceOptions? options = null)
            : base("vultr:index/instanceIpv4:InstanceIpv4", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceIpv4 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceIpv4 Get(string name, Input<string> id, InstanceIpv4State? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceIpv4(name, id, state, options);
        }
    }

    public sealed class InstanceIpv4Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the instance to be assigned the IPv4 address.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Default true. Determines whether or not the server is rebooted after adding the IPv4 address.
        /// </summary>
        [Input("reboot")]
        public Input<bool>? Reboot { get; set; }

        public InstanceIpv4Args()
        {
        }
        public static new InstanceIpv4Args Empty => new InstanceIpv4Args();
    }

    public sealed class InstanceIpv4State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The gateway IP address.
        /// </summary>
        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        /// <summary>
        /// The ID of the instance to be assigned the IPv4 address.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The IPv4 address in canonical format.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The IPv4 netmask in dot-decimal notation.
        /// </summary>
        [Input("netmask")]
        public Input<string>? Netmask { get; set; }

        /// <summary>
        /// Default true. Determines whether or not the server is rebooted after adding the IPv4 address.
        /// </summary>
        [Input("reboot")]
        public Input<bool>? Reboot { get; set; }

        /// <summary>
        /// The reverse DNS information for this IP address.
        /// </summary>
        [Input("reverse")]
        public Input<string>? Reverse { get; set; }

        public InstanceIpv4State()
        {
        }
        public static new InstanceIpv4State Empty => new InstanceIpv4State();
    }
}
