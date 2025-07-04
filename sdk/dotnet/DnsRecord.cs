// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
    /// Provides a Vultr DNS Record resource. This can be used to create, read, modify, and delete DNS Records.
    /// 
    /// ## Example Usage
    /// 
    /// Create a new DNS Record
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vultr = ediri.Vultr;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myDomain = new Vultr.DnsDomain("myDomain", new()
    ///     {
    ///         Domain = "domain.com",
    ///         Ip = "66.42.94.227",
    ///     });
    /// 
    ///     var myRecord = new Vultr.DnsRecord("myRecord", new()
    ///     {
    ///         Data = "66.42.94.227",
    ///         Domain = myDomain.Id,
    ///         Type = "A",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// DNS Records can be imported using the Dns Domain `domain` and DNS Record `ID` e.g.
    /// 
    /// ```sh
    /// $ pulumi import vultr:index/dnsRecord:DnsRecord rec domain.com,1a0019bd-7645-4310-81bd-03bc5906940f
    /// ```
    /// </summary>
    [VultrResourceType("vultr:index/dnsRecord:DnsRecord")]
    public partial class DnsRecord : global::Pulumi.CustomResource
    {
        /// <summary>
        /// IP Address of the instance the domain is associated with.
        /// </summary>
        [Output("data")]
        public Output<string> Data { get; private set; } = null!;

        /// <summary>
        /// Name of the DNS Domain this record will belong to.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Name (subdomain) for this record.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Priority of this record (only required for MX and SRV).
        /// </summary>
        [Output("priority")]
        public Output<int?> Priority { get; private set; } = null!;

        /// <summary>
        /// The time to live of this record.
        /// </summary>
        [Output("ttl")]
        public Output<int?> Ttl { get; private set; } = null!;

        /// <summary>
        /// Type of record.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a DnsRecord resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DnsRecord(string name, DnsRecordArgs args, CustomResourceOptions? options = null)
            : base("vultr:index/dnsRecord:DnsRecord", name, args ?? new DnsRecordArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DnsRecord(string name, Input<string> id, DnsRecordState? state = null, CustomResourceOptions? options = null)
            : base("vultr:index/dnsRecord:DnsRecord", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DnsRecord resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DnsRecord Get(string name, Input<string> id, DnsRecordState? state = null, CustomResourceOptions? options = null)
        {
            return new DnsRecord(name, id, state, options);
        }
    }

    public sealed class DnsRecordArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP Address of the instance the domain is associated with.
        /// </summary>
        [Input("data", required: true)]
        public Input<string> Data { get; set; } = null!;

        /// <summary>
        /// Name of the DNS Domain this record will belong to.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// Name (subdomain) for this record.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Priority of this record (only required for MX and SRV).
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The time to live of this record.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// Type of record.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public DnsRecordArgs()
        {
        }
        public static new DnsRecordArgs Empty => new DnsRecordArgs();
    }

    public sealed class DnsRecordState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP Address of the instance the domain is associated with.
        /// </summary>
        [Input("data")]
        public Input<string>? Data { get; set; }

        /// <summary>
        /// Name of the DNS Domain this record will belong to.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Name (subdomain) for this record.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Priority of this record (only required for MX and SRV).
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The time to live of this record.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// Type of record.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public DnsRecordState()
        {
        }
        public static new DnsRecordState Empty => new DnsRecordState();
    }
}
