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
    public static class GetPlan
    {
        /// <summary>
        /// Get information about a Vultr plan.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Get the information for a plan by `id`:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myPlan = Vultr.GetPlan.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Vultr.Inputs.GetPlanFilterInputArgs
        ///             {
        ///                 Name = "id",
        ///                 Values = new[]
        ///                 {
        ///                     "vc2-1c-1gb",
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
        public static Task<GetPlanResult> InvokeAsync(GetPlanArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPlanResult>("vultr:index/getPlan:getPlan", args ?? new GetPlanArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Vultr plan.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Get the information for a plan by `id`:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myPlan = Vultr.GetPlan.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Vultr.Inputs.GetPlanFilterInputArgs
        ///             {
        ///                 Name = "id",
        ///                 Values = new[]
        ///                 {
        ///                     "vc2-1c-1gb",
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
        public static Output<GetPlanResult> Invoke(GetPlanInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPlanResult>("vultr:index/getPlan:getPlan", args ?? new GetPlanInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPlanArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetPlanFilterArgs>? _filters;

        /// <summary>
        /// Query parameters for finding plans.
        /// </summary>
        public List<Inputs.GetPlanFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetPlanFilterArgs>());
            set => _filters = value;
        }

        public GetPlanArgs()
        {
        }
        public static new GetPlanArgs Empty => new GetPlanArgs();
    }

    public sealed class GetPlanInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetPlanFilterInputArgs>? _filters;

        /// <summary>
        /// Query parameters for finding plans.
        /// </summary>
        public InputList<Inputs.GetPlanFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetPlanFilterInputArgs>());
            set => _filters = value;
        }

        public GetPlanInvokeArgs()
        {
        }
        public static new GetPlanInvokeArgs Empty => new GetPlanInvokeArgs();
    }


    [OutputType]
    public sealed class GetPlanResult
    {
        /// <summary>
        /// The bandwidth available on the plan in GB.
        /// </summary>
        public readonly int Bandwidth;
        /// <summary>
        /// The amount of disk space in GB available on the plan.
        /// </summary>
        public readonly int Disk;
        /// <summary>
        /// The number of disks that this plan offers.
        /// </summary>
        public readonly int DiskCount;
        public readonly ImmutableArray<Outputs.GetPlanFilterResult> Filters;
        /// <summary>
        /// For GPU plans, the GPU card type.
        /// </summary>
        public readonly string GpuType;
        /// <summary>
        /// For GPU plans, the VRAM available in the plan.
        /// </summary>
        public readonly int GpuVram;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Locations;
        /// <summary>
        /// The price per month of the plan in USD.
        /// </summary>
        public readonly double MonthlyCost;
        /// <summary>
        /// The amount of memory available on the plan in MB.
        /// </summary>
        public readonly int Ram;
        /// <summary>
        /// The type of plan it is.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The number of virtual CPUs available on the plan.
        /// </summary>
        public readonly int VcpuCount;

        [OutputConstructor]
        private GetPlanResult(
            int bandwidth,

            int disk,

            int diskCount,

            ImmutableArray<Outputs.GetPlanFilterResult> filters,

            string gpuType,

            int gpuVram,

            string id,

            ImmutableArray<string> locations,

            double monthlyCost,

            int ram,

            string type,

            int vcpuCount)
        {
            Bandwidth = bandwidth;
            Disk = disk;
            DiskCount = diskCount;
            Filters = filters;
            GpuType = gpuType;
            GpuVram = gpuVram;
            Id = id;
            Locations = locations;
            MonthlyCost = monthlyCost;
            Ram = ram;
            Type = type;
            VcpuCount = vcpuCount;
        }
    }
}
