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
    /// Provides a Vultr Reverse IPv6 resource. This can be used to create, read,
    /// modify, and delete reverse DNS records for IPv6 addresses. Upon success, DNS
    /// changes may take 6-12 hours to become active.
    /// </summary>
    [VultrResourceType("vultr:index/reverseIpv6:ReverseIpv6")]
    public partial class ReverseIpv6 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the server you want to set an IPv6
        /// reverse DNS record for.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The IPv6 address used in the reverse DNS record.
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// The hostname used in the IPv6 reverse DNS record.
        /// </summary>
        [Output("reverse")]
        public Output<string> Reverse { get; private set; } = null!;


        /// <summary>
        /// Create a ReverseIpv6 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReverseIpv6(string name, ReverseIpv6Args args, CustomResourceOptions? options = null)
            : base("vultr:index/reverseIpv6:ReverseIpv6", name, args ?? new ReverseIpv6Args(), MakeResourceOptions(options, ""))
        {
        }

        private ReverseIpv6(string name, Input<string> id, ReverseIpv6State? state = null, CustomResourceOptions? options = null)
            : base("vultr:index/reverseIpv6:ReverseIpv6", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReverseIpv6 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReverseIpv6 Get(string name, Input<string> id, ReverseIpv6State? state = null, CustomResourceOptions? options = null)
        {
            return new ReverseIpv6(name, id, state, options);
        }
    }

    public sealed class ReverseIpv6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the server you want to set an IPv6
        /// reverse DNS record for.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The IPv6 address used in the reverse DNS record.
        /// </summary>
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        /// <summary>
        /// The hostname used in the IPv6 reverse DNS record.
        /// </summary>
        [Input("reverse", required: true)]
        public Input<string> Reverse { get; set; } = null!;

        public ReverseIpv6Args()
        {
        }
        public static new ReverseIpv6Args Empty => new ReverseIpv6Args();
    }

    public sealed class ReverseIpv6State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the server you want to set an IPv6
        /// reverse DNS record for.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The IPv6 address used in the reverse DNS record.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The hostname used in the IPv6 reverse DNS record.
        /// </summary>
        [Input("reverse")]
        public Input<string>? Reverse { get; set; }

        public ReverseIpv6State()
        {
        }
        public static new ReverseIpv6State Empty => new ReverseIpv6State();
    }
}
