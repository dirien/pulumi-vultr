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
    public sealed class GetUserFilterResult
    {
        /// <summary>
        /// Attribute name to filter with.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// One or more values filter with.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetUserFilterResult(
            string name,

            ImmutableArray<string> values)
        {
            Name = name;
            Values = values;
        }
    }
}
