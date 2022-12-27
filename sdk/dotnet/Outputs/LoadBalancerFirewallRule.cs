// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace dirien.Vultr.Outputs
{

    [OutputType]
    public sealed class LoadBalancerFirewallRule
    {
        /// <summary>
        /// The load balancer ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The type of ip this rule is - may be either v4 or v6.
        /// </summary>
        public readonly string IpType;
        /// <summary>
        /// The assigned port (integer) on the attached instances that the load balancer should check against. Default value is `80`.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// IP address with subnet that is allowed through the firewall. You may also pass in `cloudflare` which will allow only CloudFlares IP range.
        /// </summary>
        public readonly string Source;

        [OutputConstructor]
        private LoadBalancerFirewallRule(
            string? id,

            string ipType,

            int port,

            string source)
        {
            Id = id;
            IpType = ipType;
            Port = port;
            Source = source;
        }
    }
}
