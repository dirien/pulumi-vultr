// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace dirien.Vultr
{
    public static class GetInstanceIpv4
    {
        /// <summary>
        /// Get information about a Vultr instance IPv4.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Get the information for an IPv4 address by `instance_id`:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myInstanceIpv4 = Vultr.GetInstanceIpv4.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Vultr.Inputs.GetInstanceIpv4FilterInputArgs
        ///             {
        ///                 Name = "ip",
        ///                 Values = new[]
        ///                 {
        ///                     "123.123.123.123",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstanceIpv4Result> InvokeAsync(GetInstanceIpv4Args? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstanceIpv4Result>("vultr:index/getInstanceIpv4:getInstanceIpv4", args ?? new GetInstanceIpv4Args(), options.WithDefaults());

        /// <summary>
        /// Get information about a Vultr instance IPv4.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Get the information for an IPv4 address by `instance_id`:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myInstanceIpv4 = Vultr.GetInstanceIpv4.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Vultr.Inputs.GetInstanceIpv4FilterInputArgs
        ///             {
        ///                 Name = "ip",
        ///                 Values = new[]
        ///                 {
        ///                     "123.123.123.123",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetInstanceIpv4Result> Invoke(GetInstanceIpv4InvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceIpv4Result>("vultr:index/getInstanceIpv4:getInstanceIpv4", args ?? new GetInstanceIpv4InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceIpv4Args : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetInstanceIpv4FilterArgs>? _filters;

        /// <summary>
        /// Query parameters for finding IPv4 address.
        /// </summary>
        public List<Inputs.GetInstanceIpv4FilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetInstanceIpv4FilterArgs>());
            set => _filters = value;
        }

        public GetInstanceIpv4Args()
        {
        }
        public static new GetInstanceIpv4Args Empty => new GetInstanceIpv4Args();
    }

    public sealed class GetInstanceIpv4InvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetInstanceIpv4FilterInputArgs>? _filters;

        /// <summary>
        /// Query parameters for finding IPv4 address.
        /// </summary>
        public InputList<Inputs.GetInstanceIpv4FilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetInstanceIpv4FilterInputArgs>());
            set => _filters = value;
        }

        public GetInstanceIpv4InvokeArgs()
        {
        }
        public static new GetInstanceIpv4InvokeArgs Empty => new GetInstanceIpv4InvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceIpv4Result
    {
        public readonly ImmutableArray<Outputs.GetInstanceIpv4FilterResult> Filters;
        /// <summary>
        /// The gateway IP address.
        /// </summary>
        public readonly string Gateway;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of the instance the IPv4 address.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// The IPv4 address in canonical format.
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// The IPv4 netmask in dot-decimal notation.
        /// </summary>
        public readonly string Netmask;
        /// <summary>
        /// The reverse DNS information for this IP address.
        /// </summary>
        public readonly string Reverse;

        [OutputConstructor]
        private GetInstanceIpv4Result(
            ImmutableArray<Outputs.GetInstanceIpv4FilterResult> filters,

            string gateway,

            string id,

            string instanceId,

            string ip,

            string netmask,

            string reverse)
        {
            Filters = filters;
            Gateway = gateway;
            Id = id;
            InstanceId = instanceId;
            Ip = ip;
            Netmask = netmask;
            Reverse = reverse;
        }
    }
}
