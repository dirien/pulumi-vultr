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
    public static class GetVpc
    {
        /// <summary>
        /// Get information about a Vultr VPC.
        /// 
        /// ## Example Usage
        /// 
        /// Get the information for a VPC by `description`:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myVpc = Vultr.GetVpc.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Vultr.Inputs.GetVpcFilterInputArgs
        ///             {
        ///                 Name = "description",
        ///                 Values = new[]
        ///                 {
        ///                     "my-vpc-description",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetVpcResult> InvokeAsync(GetVpcArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpcResult>("vultr:index/getVpc:getVpc", args ?? new GetVpcArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Vultr VPC.
        /// 
        /// ## Example Usage
        /// 
        /// Get the information for a VPC by `description`:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myVpc = Vultr.GetVpc.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Vultr.Inputs.GetVpcFilterInputArgs
        ///             {
        ///                 Name = "description",
        ///                 Values = new[]
        ///                 {
        ///                     "my-vpc-description",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetVpcResult> Invoke(GetVpcInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcResult>("vultr:index/getVpc:getVpc", args ?? new GetVpcInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetVpcFilterArgs>? _filters;

        /// <summary>
        /// Query parameters for finding VPCs.
        /// </summary>
        public List<Inputs.GetVpcFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetVpcFilterArgs>());
            set => _filters = value;
        }

        public GetVpcArgs()
        {
        }
        public static new GetVpcArgs Empty => new GetVpcArgs();
    }

    public sealed class GetVpcInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetVpcFilterInputArgs>? _filters;

        /// <summary>
        /// Query parameters for finding VPCs.
        /// </summary>
        public InputList<Inputs.GetVpcFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetVpcFilterInputArgs>());
            set => _filters = value;
        }

        public GetVpcInvokeArgs()
        {
        }
        public static new GetVpcInvokeArgs Empty => new GetVpcInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcResult
    {
        /// <summary>
        /// The date the VPC was added to your Vultr account.
        /// </summary>
        public readonly string DateCreated;
        /// <summary>
        /// The VPC's description.
        /// </summary>
        public readonly string Description;
        public readonly ImmutableArray<Outputs.GetVpcFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of the region that the VPC is in.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The IPv4 network address. For example: 10.1.1.0.
        /// </summary>
        public readonly string V4Subnet;
        /// <summary>
        /// The number of bits for the netmask in CIDR notation. Example: 20
        /// </summary>
        public readonly int V4SubnetMask;

        [OutputConstructor]
        private GetVpcResult(
            string dateCreated,

            string description,

            ImmutableArray<Outputs.GetVpcFilterResult> filters,

            string id,

            string region,

            string v4Subnet,

            int v4SubnetMask)
        {
            DateCreated = dateCreated;
            Description = description;
            Filters = filters;
            Id = id;
            Region = region;
            V4Subnet = v4Subnet;
            V4SubnetMask = v4SubnetMask;
        }
    }
}
