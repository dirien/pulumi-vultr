// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Vultr.Outputs
{

    [OutputType]
    public sealed class InstanceBackupsSchedule
    {
        /// <summary>
        /// Day of month to run. Use values between 1 and 28.
        /// </summary>
        public readonly int? Dom;
        /// <summary>
        /// Day of week to run. `1 = Sunday`, `2 = Monday`, `3 = Tuesday`, `4 = Wednesday`, `5 = Thursday`, `6 = Friday`, `7 = Saturday`
        /// </summary>
        public readonly int? Dow;
        /// <summary>
        /// Hour of day to run in UTC.
        /// </summary>
        public readonly int? Hour;
        /// <summary>
        /// Type of backup schedule Possible values are `daily`, `weekly`, `monthly`, `daily_alt_even`, or `daily_alt_odd`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private InstanceBackupsSchedule(
            int? dom,

            int? dow,

            int? hour,

            string type)
        {
            Dom = dom;
            Dow = dow;
            Hour = hour;
            Type = type;
        }
    }
}
