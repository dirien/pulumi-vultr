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
    public static class GetSnapshot
    {
        /// <summary>
        /// Get information about a Vultr snapshot.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Get the information for a snapshot by `description`:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mySnapshot = Vultr.GetSnapshot.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Vultr.Inputs.GetSnapshotFilterInputArgs
        ///             {
        ///                 Name = "description",
        ///                 Values = new[]
        ///                 {
        ///                     "my-snapshot-description",
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
        public static Task<GetSnapshotResult> InvokeAsync(GetSnapshotArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSnapshotResult>("vultr:index/getSnapshot:getSnapshot", args ?? new GetSnapshotArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Vultr snapshot.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Get the information for a snapshot by `description`:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mySnapshot = Vultr.GetSnapshot.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Vultr.Inputs.GetSnapshotFilterInputArgs
        ///             {
        ///                 Name = "description",
        ///                 Values = new[]
        ///                 {
        ///                     "my-snapshot-description",
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
        public static Output<GetSnapshotResult> Invoke(GetSnapshotInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSnapshotResult>("vultr:index/getSnapshot:getSnapshot", args ?? new GetSnapshotInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSnapshotArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetSnapshotFilterArgs>? _filters;

        /// <summary>
        /// Query parameters for finding snapshots.
        /// </summary>
        public List<Inputs.GetSnapshotFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetSnapshotFilterArgs>());
            set => _filters = value;
        }

        public GetSnapshotArgs()
        {
        }
        public static new GetSnapshotArgs Empty => new GetSnapshotArgs();
    }

    public sealed class GetSnapshotInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetSnapshotFilterInputArgs>? _filters;

        /// <summary>
        /// Query parameters for finding snapshots.
        /// </summary>
        public InputList<Inputs.GetSnapshotFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetSnapshotFilterInputArgs>());
            set => _filters = value;
        }

        public GetSnapshotInvokeArgs()
        {
        }
        public static new GetSnapshotInvokeArgs Empty => new GetSnapshotInvokeArgs();
    }


    [OutputType]
    public sealed class GetSnapshotResult
    {
        /// <summary>
        /// The application ID of the snapshot.
        /// </summary>
        public readonly int AppId;
        /// <summary>
        /// The date the snapshot was added to your Vultr account.
        /// </summary>
        public readonly string DateCreated;
        /// <summary>
        /// The description of the snapshot.
        /// </summary>
        public readonly string Description;
        public readonly ImmutableArray<Outputs.GetSnapshotFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The operating system ID of the snapshot.
        /// </summary>
        public readonly int OsId;
        /// <summary>
        /// The size of the snapshot in bytes.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// The status of the snapshot.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetSnapshotResult(
            int appId,

            string dateCreated,

            string description,

            ImmutableArray<Outputs.GetSnapshotFilterResult> filters,

            string id,

            int osId,

            int size,

            string status)
        {
            AppId = appId;
            DateCreated = dateCreated;
            Description = description;
            Filters = filters;
            Id = id;
            OsId = osId;
            Size = size;
            Status = status;
        }
    }
}
