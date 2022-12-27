// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace dirien.Vultr.Inputs
{

    public sealed class LoadBalancerHealthCheckArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time in seconds to perform health check. Default value is 15.
        /// </summary>
        [Input("checkInterval", required: true)]
        public Input<int> CheckInterval { get; set; } = null!;

        /// <summary>
        /// Number of failed attempts encountered before failover. Default value is 5.
        /// </summary>
        [Input("healthyThreshold", required: true)]
        public Input<int> HealthyThreshold { get; set; } = null!;

        /// <summary>
        /// The path on the attached instances that the load balancer should check against. Default value is `/`
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The assigned port (integer) on the attached instances that the load balancer should check against. Default value is `80`.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// The protocol used to traffic requests to the load balancer. Possible values are `http`, or `tcp`. Default value is `http`.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// Time in seconds to wait for a health check response. Default value is 5.
        /// </summary>
        [Input("responseTimeout", required: true)]
        public Input<int> ResponseTimeout { get; set; } = null!;

        /// <summary>
        /// Number of failed attempts encountered before failover. Default value is 5.
        /// </summary>
        [Input("unhealthyThreshold", required: true)]
        public Input<int> UnhealthyThreshold { get; set; } = null!;

        public LoadBalancerHealthCheckArgs()
        {
        }
        public static new LoadBalancerHealthCheckArgs Empty => new LoadBalancerHealthCheckArgs();
    }
}
