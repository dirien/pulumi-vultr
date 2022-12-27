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
    public static class GetInstance
    {
        /// <summary>
        /// Get information about a Vultr instance.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Get the information for a instance by `label`:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myInstance = Vultr.GetInstance.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Vultr.Inputs.GetInstanceFilterInputArgs
        ///             {
        ///                 Name = "label",
        ///                 Values = new[]
        ///                 {
        ///                     "my-instance-label",
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
        public static Task<GetInstanceResult> InvokeAsync(GetInstanceArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstanceResult>("vultr:index/getInstance:getInstance", args ?? new GetInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Vultr instance.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Get the information for a instance by `label`:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myInstance = Vultr.GetInstance.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Vultr.Inputs.GetInstanceFilterInputArgs
        ///             {
        ///                 Name = "label",
        ///                 Values = new[]
        ///                 {
        ///                     "my-instance-label",
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
        public static Output<GetInstanceResult> Invoke(GetInstanceInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceResult>("vultr:index/getInstance:getInstance", args ?? new GetInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetInstanceFilterArgs>? _filters;

        /// <summary>
        /// Query parameters for finding instances.
        /// </summary>
        public List<Inputs.GetInstanceFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetInstanceFilterArgs>());
            set => _filters = value;
        }

        public GetInstanceArgs()
        {
        }
        public static new GetInstanceArgs Empty => new GetInstanceArgs();
    }

    public sealed class GetInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetInstanceFilterInputArgs>? _filters;

        /// <summary>
        /// Query parameters for finding instances.
        /// </summary>
        public InputList<Inputs.GetInstanceFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetInstanceFilterInputArgs>());
            set => _filters = value;
        }

        public GetInstanceInvokeArgs()
        {
        }
        public static new GetInstanceInvokeArgs Empty => new GetInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceResult
    {
        /// <summary>
        /// The server's allowed bandwidth usage in GB.
        /// </summary>
        public readonly int AllowedBandwidth;
        /// <summary>
        /// The server's application ID.
        /// </summary>
        public readonly int AppId;
        public readonly string Backups;
        /// <summary>
        /// The current configuration for backups
        /// </summary>
        public readonly ImmutableDictionary<string, object> BackupsSchedule;
        /// <summary>
        /// The date the server was added to your Vultr account.
        /// </summary>
        public readonly string DateCreated;
        /// <summary>
        /// The description of the disk(s) on the server.
        /// </summary>
        public readonly int Disk;
        /// <summary>
        /// Array of which features are enabled.
        /// </summary>
        public readonly ImmutableArray<string> Features;
        public readonly ImmutableArray<Outputs.GetInstanceFilterResult> Filters;
        /// <summary>
        /// The ID of the firewall group applied to this server.
        /// </summary>
        public readonly string FirewallGroupId;
        /// <summary>
        /// The server's IPv4 gateway.
        /// </summary>
        public readonly string GatewayV4;
        /// <summary>
        /// The hostname assigned to the server.
        /// </summary>
        public readonly string Hostname;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Marketplace ID for this application.
        /// </summary>
        public readonly string ImageId;
        /// <summary>
        /// The server's internal IP address.
        /// </summary>
        public readonly string InternalIp;
        /// <summary>
        /// The server's current KVM URL. This URL will change periodically. It is not advised to cache this value.
        /// </summary>
        public readonly string Kvm;
        /// <summary>
        /// The server's label.
        /// </summary>
        public readonly string Label;
        public readonly string Location;
        /// <summary>
        /// The server's main IP address.
        /// </summary>
        public readonly string MainIp;
        /// <summary>
        /// The server's IPv4 netmask.
        /// </summary>
        public readonly string NetmaskV4;
        /// <summary>
        /// The operating system of the instance.
        /// </summary>
        public readonly string Os;
        /// <summary>
        /// The server's operating system ID.
        /// </summary>
        public readonly int OsId;
        /// <summary>
        /// The server's plan ID.
        /// </summary>
        public readonly string Plan;
        /// <summary>
        /// Whether the server is powered on or not.
        /// </summary>
        public readonly string PowerStatus;
        public readonly ImmutableArray<string> PrivateNetworkIds;
        /// <summary>
        /// The amount of memory available on the instance in MB.
        /// </summary>
        public readonly int Ram;
        /// <summary>
        /// The region ID of the server.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// A more detailed server status (none, locked, installingbooting, isomounting, ok).
        /// </summary>
        public readonly string ServerStatus;
        /// <summary>
        /// The status of the server's subscription.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// A list of tags applied to the instance.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The main IPv6 network address.
        /// </summary>
        public readonly string V6MainIp;
        /// <summary>
        /// The IPv6 subnet.
        /// </summary>
        public readonly string V6Network;
        /// <summary>
        /// The IPv6 network size in bits.
        /// </summary>
        public readonly int V6NetworkSize;
        /// <summary>
        /// The number of virtual CPUs available on the server.
        /// </summary>
        public readonly int VcpuCount;
        public readonly ImmutableArray<string> VpcIds;

        [OutputConstructor]
        private GetInstanceResult(
            int allowedBandwidth,

            int appId,

            string backups,

            ImmutableDictionary<string, object> backupsSchedule,

            string dateCreated,

            int disk,

            ImmutableArray<string> features,

            ImmutableArray<Outputs.GetInstanceFilterResult> filters,

            string firewallGroupId,

            string gatewayV4,

            string hostname,

            string id,

            string imageId,

            string internalIp,

            string kvm,

            string label,

            string location,

            string mainIp,

            string netmaskV4,

            string os,

            int osId,

            string plan,

            string powerStatus,

            ImmutableArray<string> privateNetworkIds,

            int ram,

            string region,

            string serverStatus,

            string status,

            ImmutableArray<string> tags,

            string v6MainIp,

            string v6Network,

            int v6NetworkSize,

            int vcpuCount,

            ImmutableArray<string> vpcIds)
        {
            AllowedBandwidth = allowedBandwidth;
            AppId = appId;
            Backups = backups;
            BackupsSchedule = backupsSchedule;
            DateCreated = dateCreated;
            Disk = disk;
            Features = features;
            Filters = filters;
            FirewallGroupId = firewallGroupId;
            GatewayV4 = gatewayV4;
            Hostname = hostname;
            Id = id;
            ImageId = imageId;
            InternalIp = internalIp;
            Kvm = kvm;
            Label = label;
            Location = location;
            MainIp = mainIp;
            NetmaskV4 = netmaskV4;
            Os = os;
            OsId = osId;
            Plan = plan;
            PowerStatus = powerStatus;
            PrivateNetworkIds = privateNetworkIds;
            Ram = ram;
            Region = region;
            ServerStatus = serverStatus;
            Status = status;
            Tags = tags;
            V6MainIp = v6MainIp;
            V6Network = v6Network;
            V6NetworkSize = v6NetworkSize;
            VcpuCount = vcpuCount;
            VpcIds = vpcIds;
        }
    }
}
