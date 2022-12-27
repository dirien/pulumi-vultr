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
    public static class GetApplication
    {
        /// <summary>
        /// Get information about applications that can be launched when creating a Vultr VPS.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Get the information for an application by `deploy_name`:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var docker = Vultr.GetApplication.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Vultr.Inputs.GetApplicationFilterInputArgs
        ///             {
        ///                 Name = "deploy_name",
        ///                 Values = new[]
        ///                 {
        ///                     "Docker on CentOS 7 x64",
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
        public static Task<GetApplicationResult> InvokeAsync(GetApplicationArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApplicationResult>("vultr:index/getApplication:getApplication", args ?? new GetApplicationArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about applications that can be launched when creating a Vultr VPS.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Get the information for an application by `deploy_name`:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var docker = Vultr.GetApplication.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Vultr.Inputs.GetApplicationFilterInputArgs
        ///             {
        ///                 Name = "deploy_name",
        ///                 Values = new[]
        ///                 {
        ///                     "Docker on CentOS 7 x64",
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
        public static Output<GetApplicationResult> Invoke(GetApplicationInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApplicationResult>("vultr:index/getApplication:getApplication", args ?? new GetApplicationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApplicationArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetApplicationFilterArgs>? _filters;

        /// <summary>
        /// Query parameters for finding applications.
        /// </summary>
        public List<Inputs.GetApplicationFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetApplicationFilterArgs>());
            set => _filters = value;
        }

        public GetApplicationArgs()
        {
        }
        public static new GetApplicationArgs Empty => new GetApplicationArgs();
    }

    public sealed class GetApplicationInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetApplicationFilterInputArgs>? _filters;

        /// <summary>
        /// Query parameters for finding applications.
        /// </summary>
        public InputList<Inputs.GetApplicationFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetApplicationFilterInputArgs>());
            set => _filters = value;
        }

        public GetApplicationInvokeArgs()
        {
        }
        public static new GetApplicationInvokeArgs Empty => new GetApplicationInvokeArgs();
    }


    [OutputType]
    public sealed class GetApplicationResult
    {
        /// <summary>
        /// The deploy name of the application.
        /// </summary>
        public readonly string DeployName;
        public readonly ImmutableArray<Outputs.GetApplicationFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Marketplace ID for this application.
        /// </summary>
        public readonly string ImageId;
        /// <summary>
        /// The name of the application.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The short name of the application.
        /// </summary>
        public readonly string ShortName;
        /// <summary>
        /// This defines if an application is `one-click` or `marketplace` .
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The vendor who authored the application.
        /// </summary>
        public readonly string Vendor;

        [OutputConstructor]
        private GetApplicationResult(
            string deployName,

            ImmutableArray<Outputs.GetApplicationFilterResult> filters,

            string id,

            string imageId,

            string name,

            string shortName,

            string type,

            string vendor)
        {
            DeployName = deployName;
            Filters = filters;
            Id = id;
            ImageId = imageId;
            Name = name;
            ShortName = shortName;
            Type = type;
            Vendor = vendor;
        }
    }
}
