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

    public sealed class InstanceBackupsScheduleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Day of month to run. Use values between 1 and 28.
        /// </summary>
        [Input("dom")]
        public Input<int>? Dom { get; set; }

        /// <summary>
        /// Day of week to run. `1 = Sunday`, `2 = Monday`, `3 = Tuesday`, `4 = Wednesday`, `5 = Thursday`, `6 = Friday`, `7 = Saturday`
        /// </summary>
        [Input("dow")]
        public Input<int>? Dow { get; set; }

        /// <summary>
        /// Hour of day to run in UTC.
        /// </summary>
        [Input("hour")]
        public Input<int>? Hour { get; set; }

        /// <summary>
        /// Type of backup schedule Possible values are `daily`, `weekly`, `monthly`, `daily_alt_even`, or `daily_alt_odd`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public InstanceBackupsScheduleGetArgs()
        {
        }
        public static new InstanceBackupsScheduleGetArgs Empty => new InstanceBackupsScheduleGetArgs();
    }
}
