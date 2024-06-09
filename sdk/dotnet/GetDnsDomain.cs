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
    public static class GetDnsDomain
    {
        /// <summary>
        /// Get information about a DNS domain associated with your Vultr account.
        /// 
        /// ## Example Usage
        /// 
        /// Get the information for a DNS domain:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myDomain = Vultr.GetDnsDomain.Invoke(new()
        ///     {
        ///         Domain = "example.com",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetDnsDomainResult> InvokeAsync(GetDnsDomainArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDnsDomainResult>("vultr:index/getDnsDomain:getDnsDomain", args ?? new GetDnsDomainArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a DNS domain associated with your Vultr account.
        /// 
        /// ## Example Usage
        /// 
        /// Get the information for a DNS domain:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myDomain = Vultr.GetDnsDomain.Invoke(new()
        ///     {
        ///         Domain = "example.com",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDnsDomainResult> Invoke(GetDnsDomainInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDnsDomainResult>("vultr:index/getDnsDomain:getDnsDomain", args ?? new GetDnsDomainInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDnsDomainArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name you're searching for.
        /// </summary>
        [Input("domain", required: true)]
        public string Domain { get; set; } = null!;

        public GetDnsDomainArgs()
        {
        }
        public static new GetDnsDomainArgs Empty => new GetDnsDomainArgs();
    }

    public sealed class GetDnsDomainInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name you're searching for.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        public GetDnsDomainInvokeArgs()
        {
        }
        public static new GetDnsDomainInvokeArgs Empty => new GetDnsDomainInvokeArgs();
    }


    [OutputType]
    public sealed class GetDnsDomainResult
    {
        /// <summary>
        /// The date the DNS domain was added to your Vultr account.
        /// </summary>
        public readonly string DateCreated;
        /// <summary>
        /// The Domain's DNSSEC status
        /// </summary>
        public readonly string DnsSec;
        /// <summary>
        /// Name of domain.
        /// </summary>
        public readonly string Domain;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetDnsDomainResult(
            string dateCreated,

            string dnsSec,

            string domain,

            string id)
        {
            DateCreated = dateCreated;
            DnsSec = dnsSec;
            Domain = domain;
            Id = id;
        }
    }
}
