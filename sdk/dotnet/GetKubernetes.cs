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
    public static class GetKubernetes
    {
        /// <summary>
        /// Get information about a Vultr Kubernetes Engine (VKE) Cluster.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Create a new VKE cluster:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myVke = Vultr.GetKubernetes.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Vultr.Inputs.GetKubernetesFilterInputArgs
        ///             {
        ///                 Name = "label",
        ///                 Values = new[]
        ///                 {
        ///                     "my-lb-label",
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
        public static Task<GetKubernetesResult> InvokeAsync(GetKubernetesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetKubernetesResult>("vultr:index/getKubernetes:getKubernetes", args ?? new GetKubernetesArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Vultr Kubernetes Engine (VKE) Cluster.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Create a new VKE cluster:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myVke = Vultr.GetKubernetes.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Vultr.Inputs.GetKubernetesFilterInputArgs
        ///             {
        ///                 Name = "label",
        ///                 Values = new[]
        ///                 {
        ///                     "my-lb-label",
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
        public static Output<GetKubernetesResult> Invoke(GetKubernetesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetKubernetesResult>("vultr:index/getKubernetes:getKubernetes", args ?? new GetKubernetesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKubernetesArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetKubernetesFilterArgs>? _filters;

        /// <summary>
        /// Query parameters for finding VKE.
        /// </summary>
        public List<Inputs.GetKubernetesFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetKubernetesFilterArgs>());
            set => _filters = value;
        }

        public GetKubernetesArgs()
        {
        }
        public static new GetKubernetesArgs Empty => new GetKubernetesArgs();
    }

    public sealed class GetKubernetesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetKubernetesFilterInputArgs>? _filters;

        /// <summary>
        /// Query parameters for finding VKE.
        /// </summary>
        public InputList<Inputs.GetKubernetesFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetKubernetesFilterInputArgs>());
            set => _filters = value;
        }

        public GetKubernetesInvokeArgs()
        {
        }
        public static new GetKubernetesInvokeArgs Empty => new GetKubernetesInvokeArgs();
    }


    [OutputType]
    public sealed class GetKubernetesResult
    {
        /// <summary>
        /// The base64 encoded public certificate used by clients to access the cluster.
        /// </summary>
        public readonly string ClientCertificate;
        /// <summary>
        /// The base64 encoded private key used by clients to access the cluster.
        /// </summary>
        public readonly string ClientKey;
        /// <summary>
        /// The base64 encoded public certificate for the cluster's certificate authority.
        /// </summary>
        public readonly string ClusterCaCertificate;
        /// <summary>
        /// IP range that your pods will run on in this cluster.
        /// </summary>
        public readonly string ClusterSubnet;
        /// <summary>
        /// Date node was created.
        /// </summary>
        public readonly string DateCreated;
        /// <summary>
        /// Domain for your Kubernetes clusters control plane.
        /// </summary>
        public readonly string Endpoint;
        public readonly ImmutableArray<Outputs.GetKubernetesFilterResult> Filters;
        /// <summary>
        /// ID of node.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// IP address of VKE cluster control plane.
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// Base64 encoded Kubeconfig for this VKE cluster.
        /// </summary>
        public readonly string KubeConfig;
        /// <summary>
        /// Label of node.
        /// </summary>
        public readonly string Label;
        /// <summary>
        /// Contains the default node pool that was deployed.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesNodePoolResult> NodePools;
        /// <summary>
        /// The region your VKE cluster is deployed in.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// IP range that services will run on this cluster.
        /// </summary>
        public readonly string ServiceSubnet;
        /// <summary>
        /// Status of node.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The current kubernetes version your VKE cluster is running on.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetKubernetesResult(
            string clientCertificate,

            string clientKey,

            string clusterCaCertificate,

            string clusterSubnet,

            string dateCreated,

            string endpoint,

            ImmutableArray<Outputs.GetKubernetesFilterResult> filters,

            string id,

            string ip,

            string kubeConfig,

            string label,

            ImmutableArray<Outputs.GetKubernetesNodePoolResult> nodePools,

            string region,

            string serviceSubnet,

            string status,

            string version)
        {
            ClientCertificate = clientCertificate;
            ClientKey = clientKey;
            ClusterCaCertificate = clusterCaCertificate;
            ClusterSubnet = clusterSubnet;
            DateCreated = dateCreated;
            Endpoint = endpoint;
            Filters = filters;
            Id = id;
            Ip = ip;
            KubeConfig = kubeConfig;
            Label = label;
            NodePools = nodePools;
            Region = region;
            ServiceSubnet = serviceSubnet;
            Status = status;
            Version = version;
        }
    }
}
