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
    public static class GetOs
    {
        /// <summary>
        /// Get information about operating systems that can be launched when creating a Vultr VPS.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Get the information for an operating system by `name`:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var centos = Vultr.GetOs.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Vultr.Inputs.GetOsFilterInputArgs
        ///             {
        ///                 Name = "name",
        ///                 Values = new[]
        ///                 {
        ///                     "CentOS 7 x64",
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
        public static Task<GetOsResult> InvokeAsync(GetOsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOsResult>("vultr:index/getOs:getOs", args ?? new GetOsArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about operating systems that can be launched when creating a Vultr VPS.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Get the information for an operating system by `name`:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var centos = Vultr.GetOs.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Vultr.Inputs.GetOsFilterInputArgs
        ///             {
        ///                 Name = "name",
        ///                 Values = new[]
        ///                 {
        ///                     "CentOS 7 x64",
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
        public static Output<GetOsResult> Invoke(GetOsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOsResult>("vultr:index/getOs:getOs", args ?? new GetOsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOsArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetOsFilterArgs>? _filters;

        /// <summary>
        /// Query parameters for finding operating systems.
        /// </summary>
        public List<Inputs.GetOsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetOsFilterArgs>());
            set => _filters = value;
        }

        public GetOsArgs()
        {
        }
        public static new GetOsArgs Empty => new GetOsArgs();
    }

    public sealed class GetOsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetOsFilterInputArgs>? _filters;

        /// <summary>
        /// Query parameters for finding operating systems.
        /// </summary>
        public InputList<Inputs.GetOsFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetOsFilterInputArgs>());
            set => _filters = value;
        }

        public GetOsInvokeArgs()
        {
        }
        public static new GetOsInvokeArgs Empty => new GetOsInvokeArgs();
    }


    [OutputType]
    public sealed class GetOsResult
    {
        /// <summary>
        /// The architecture of the operating system.
        /// </summary>
        public readonly string Arch;
        /// <summary>
        /// The family of the operating system.
        /// </summary>
        public readonly string Family;
        public readonly ImmutableArray<Outputs.GetOsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the operating system.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetOsResult(
            string arch,

            string family,

            ImmutableArray<Outputs.GetOsFilterResult> filters,

            string id,

            string name)
        {
            Arch = arch;
            Family = family;
            Filters = filters;
            Id = id;
            Name = name;
        }
    }
}
