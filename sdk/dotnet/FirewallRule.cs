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
    /// Provides a Vultr Firewall Rule resource. This can be used to create, read, modify, and delete Firewall rules.
    /// 
    /// ## Example Usage
    /// 
    /// Create a Firewall Rule
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vultr = ediri.Vultr;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myFirewallgroup = new Vultr.FirewallGroup("myFirewallgroup", new()
    ///     {
    ///         Description = "base firewall",
    ///     });
    /// 
    ///     var myFirewallrule = new Vultr.FirewallRule("myFirewallrule", new()
    ///     {
    ///         FirewallGroupId = myFirewallgroup.Id,
    ///         Protocol = "tcp",
    ///         IpType = "v4",
    ///         Subnet = "0.0.0.0",
    ///         SubnetSize = 0,
    ///         Port = "8090",
    ///         Notes = "my firewall rule",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Firewall Rules can be imported using the Firewall Group `ID` and Firewall Rule `ID`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vultr:index/firewallRule:FirewallRule my_rule b6a859c5-b299-49dd-8888-b1abbc517d08,1
    /// ```
    /// </summary>
    [VultrResourceType("vultr:index/firewallRule:FirewallRule")]
    public partial class FirewallRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The firewall group that the firewall rule will belong to.
        /// </summary>
        [Output("firewallGroupId")]
        public Output<string> FirewallGroupId { get; private set; } = null!;

        /// <summary>
        /// The type of ip for this firewall rule. Possible values (v4, v6) **Note** they must be lowercase
        /// </summary>
        [Output("ipType")]
        public Output<string> IpType { get; private set; } = null!;

        /// <summary>
        /// A simple note for a given firewall rule
        /// </summary>
        [Output("notes")]
        public Output<string?> Notes { get; private set; } = null!;

        /// <summary>
        /// TCP/UDP only. This field can be a specific port or a colon separated port range.
        /// </summary>
        [Output("port")]
        public Output<string?> Port { get; private set; } = null!;

        /// <summary>
        /// The type of protocol for this firewall rule. Possible values (icmp, tcp, udp, gre, esp, ah) **Note** they must be lowercase
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// Possible values ("", cloudflare)
        /// </summary>
        [Output("source")]
        public Output<string?> Source { get; private set; } = null!;

        /// <summary>
        /// IP address that you want to define for this firewall rule.
        /// </summary>
        [Output("subnet")]
        public Output<string> Subnet { get; private set; } = null!;

        /// <summary>
        /// The number of bits for the subnet in CIDR notation. Example: 32.
        /// </summary>
        [Output("subnetSize")]
        public Output<int> SubnetSize { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallRule(string name, FirewallRuleArgs args, CustomResourceOptions? options = null)
            : base("vultr:index/firewallRule:FirewallRule", name, args ?? new FirewallRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallRule(string name, Input<string> id, FirewallRuleState? state = null, CustomResourceOptions? options = null)
            : base("vultr:index/firewallRule:FirewallRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FirewallRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallRule Get(string name, Input<string> id, FirewallRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallRule(name, id, state, options);
        }
    }

    public sealed class FirewallRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The firewall group that the firewall rule will belong to.
        /// </summary>
        [Input("firewallGroupId", required: true)]
        public Input<string> FirewallGroupId { get; set; } = null!;

        /// <summary>
        /// The type of ip for this firewall rule. Possible values (v4, v6) **Note** they must be lowercase
        /// </summary>
        [Input("ipType", required: true)]
        public Input<string> IpType { get; set; } = null!;

        /// <summary>
        /// A simple note for a given firewall rule
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        /// <summary>
        /// TCP/UDP only. This field can be a specific port or a colon separated port range.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// The type of protocol for this firewall rule. Possible values (icmp, tcp, udp, gre, esp, ah) **Note** they must be lowercase
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// Possible values ("", cloudflare)
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// IP address that you want to define for this firewall rule.
        /// </summary>
        [Input("subnet", required: true)]
        public Input<string> Subnet { get; set; } = null!;

        /// <summary>
        /// The number of bits for the subnet in CIDR notation. Example: 32.
        /// </summary>
        [Input("subnetSize", required: true)]
        public Input<int> SubnetSize { get; set; } = null!;

        public FirewallRuleArgs()
        {
        }
        public static new FirewallRuleArgs Empty => new FirewallRuleArgs();
    }

    public sealed class FirewallRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The firewall group that the firewall rule will belong to.
        /// </summary>
        [Input("firewallGroupId")]
        public Input<string>? FirewallGroupId { get; set; }

        /// <summary>
        /// The type of ip for this firewall rule. Possible values (v4, v6) **Note** they must be lowercase
        /// </summary>
        [Input("ipType")]
        public Input<string>? IpType { get; set; }

        /// <summary>
        /// A simple note for a given firewall rule
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        /// <summary>
        /// TCP/UDP only. This field can be a specific port or a colon separated port range.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// The type of protocol for this firewall rule. Possible values (icmp, tcp, udp, gre, esp, ah) **Note** they must be lowercase
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Possible values ("", cloudflare)
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// IP address that you want to define for this firewall rule.
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        /// <summary>
        /// The number of bits for the subnet in CIDR notation. Example: 32.
        /// </summary>
        [Input("subnetSize")]
        public Input<int>? SubnetSize { get; set; }

        public FirewallRuleState()
        {
        }
        public static new FirewallRuleState Empty => new FirewallRuleState();
    }
}
