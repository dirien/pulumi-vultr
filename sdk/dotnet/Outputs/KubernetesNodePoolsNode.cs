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
    public sealed class KubernetesNodePoolsNode
    {
        /// <summary>
        /// Date node was created.
        /// </summary>
        public readonly string? DateCreated;
        /// <summary>
        /// ID of node.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The VKE clusters label.
        /// </summary>
        public readonly string? Label;
        /// <summary>
        /// Status of node.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private KubernetesNodePoolsNode(
            string? dateCreated,

            string? id,

            string? label,

            string? status)
        {
            DateCreated = dateCreated;
            Id = id;
            Label = label;
            Status = status;
        }
    }
}
