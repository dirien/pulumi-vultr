// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Vultr.Inputs
{

    public sealed class LoadBalancerFirewallRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The load balancer ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The type of ip this rule is - may be either v4 or v6.
        /// </summary>
        [Input("ipType", required: true)]
        public Input<string> IpType { get; set; } = null!;

        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// IP address with subnet that is allowed through the firewall. You may also pass in `cloudflare` which will allow only CloudFlares IP range.
        /// </summary>
        [Input("source", required: true)]
        public Input<string> Source { get; set; } = null!;

        public LoadBalancerFirewallRuleArgs()
        {
        }
        public static new LoadBalancerFirewallRuleArgs Empty => new LoadBalancerFirewallRuleArgs();
    }
}
