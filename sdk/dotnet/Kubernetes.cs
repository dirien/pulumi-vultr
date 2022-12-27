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
    /// <summary>
    /// ## Example Usage
    /// 
    /// Create a new VKE cluster:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Vultr = ediri.Vultr;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var k8 = new Vultr.Kubernetes("k8", new()
    ///     {
    ///         Label = "tf-test",
    ///         NodePools = new Vultr.Inputs.KubernetesNodePoolsArgs
    ///         {
    ///             AutoScaler = true,
    ///             Label = "my-label",
    ///             MaxNodes = 2,
    ///             MinNodes = 1,
    ///             NodeQuantity = 1,
    ///             Plan = "vc2-2c-4gb",
    ///         },
    ///         Region = "ewr",
    ///         Version = "v1.23.5+1",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// A default node pool is required when first creating the resource but it can be removed at a later point so long as there is a separate `vultr.KubernetesNodePools` resource attached. For example:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Vultr = ediri.Vultr;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var k8 = new Vultr.Kubernetes("k8", new()
    ///     {
    ///         Region = "ewr",
    ///         Label = "tf-test",
    ///         Version = "v1.23.5+1",
    ///     });
    /// 
    ///     // This resource must be created and attached to the cluster
    ///     // before removing the default node from the vultr_kubernetes resource
    ///     var np = new Vultr.KubernetesNodePools("np", new()
    ///     {
    ///         ClusterId = k8.Id,
    ///         NodeQuantity = 1,
    ///         Plan = "vc2-2c-4gb",
    ///         Label = "my-label",
    ///         AutoScaler = true,
    ///         MinNodes = 1,
    ///         MaxNodes = 2,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// There is still a requirement that there be one node pool attached to the cluster but this should allow more flexibility about which node pool that is.
    /// </summary>
    [VultrResourceType("vultr:index/kubernetes:Kubernetes")]
    public partial class Kubernetes : global::Pulumi.CustomResource
    {
        /// <summary>
        /// IP range that your pods will run on in this cluster.
        /// </summary>
        [Output("clusterSubnet")]
        public Output<string> ClusterSubnet { get; private set; } = null!;

        /// <summary>
        /// Date node was created.
        /// </summary>
        [Output("dateCreated")]
        public Output<string> DateCreated { get; private set; } = null!;

        /// <summary>
        /// Domain for your Kubernetes clusters control plane.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// IP address of VKE cluster control plane.
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// Base64 encoded Kubeconfig for this VKE cluster.
        /// </summary>
        [Output("kubeConfig")]
        public Output<string> KubeConfig { get; private set; } = null!;

        /// <summary>
        /// The VKE clusters label.
        /// </summary>
        [Output("label")]
        public Output<string> Label { get; private set; } = null!;

        /// <summary>
        /// Contains the default node pool that was deployed.
        /// </summary>
        [Output("nodePools")]
        public Output<Outputs.KubernetesNodePools?> NodePools { get; private set; } = null!;

        /// <summary>
        /// The region your VKE cluster will be deployed in. Currently, supported values are `ewr` and `lax`
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// IP range that services will run on this cluster.
        /// </summary>
        [Output("serviceSubnet")]
        public Output<string> ServiceSubnet { get; private set; } = null!;

        /// <summary>
        /// Status of node.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The version your VKE cluster you want deployed. [See Available Version](https://www.vultr.com/api/#operation/get-kubernetes-versions)
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Kubernetes resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Kubernetes(string name, KubernetesArgs args, CustomResourceOptions? options = null)
            : base("vultr:index/kubernetes:Kubernetes", name, args ?? new KubernetesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Kubernetes(string name, Input<string> id, KubernetesState? state = null, CustomResourceOptions? options = null)
            : base("vultr:index/kubernetes:Kubernetes", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/dirien/pulumi-vultr",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Kubernetes resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Kubernetes Get(string name, Input<string> id, KubernetesState? state = null, CustomResourceOptions? options = null)
        {
            return new Kubernetes(name, id, state, options);
        }
    }

    public sealed class KubernetesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The VKE clusters label.
        /// </summary>
        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        /// <summary>
        /// Contains the default node pool that was deployed.
        /// </summary>
        [Input("nodePools")]
        public Input<Inputs.KubernetesNodePoolsArgs>? NodePools { get; set; }

        /// <summary>
        /// The region your VKE cluster will be deployed in. Currently, supported values are `ewr` and `lax`
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// The version your VKE cluster you want deployed. [See Available Version](https://www.vultr.com/api/#operation/get-kubernetes-versions)
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public KubernetesArgs()
        {
        }
        public static new KubernetesArgs Empty => new KubernetesArgs();
    }

    public sealed class KubernetesState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP range that your pods will run on in this cluster.
        /// </summary>
        [Input("clusterSubnet")]
        public Input<string>? ClusterSubnet { get; set; }

        /// <summary>
        /// Date node was created.
        /// </summary>
        [Input("dateCreated")]
        public Input<string>? DateCreated { get; set; }

        /// <summary>
        /// Domain for your Kubernetes clusters control plane.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// IP address of VKE cluster control plane.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Base64 encoded Kubeconfig for this VKE cluster.
        /// </summary>
        [Input("kubeConfig")]
        public Input<string>? KubeConfig { get; set; }

        /// <summary>
        /// The VKE clusters label.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// Contains the default node pool that was deployed.
        /// </summary>
        [Input("nodePools")]
        public Input<Inputs.KubernetesNodePoolsGetArgs>? NodePools { get; set; }

        /// <summary>
        /// The region your VKE cluster will be deployed in. Currently, supported values are `ewr` and `lax`
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// IP range that services will run on this cluster.
        /// </summary>
        [Input("serviceSubnet")]
        public Input<string>? ServiceSubnet { get; set; }

        /// <summary>
        /// Status of node.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The version your VKE cluster you want deployed. [See Available Version](https://www.vultr.com/api/#operation/get-kubernetes-versions)
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public KubernetesState()
        {
        }
        public static new KubernetesState Empty => new KubernetesState();
    }
}
