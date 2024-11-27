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
    /// Provides a Vultr bare metal server resource. This can be used to create, read, modify, and delete bare metal servers on your Vultr account.
    /// 
    /// ## Example Usage
    /// 
    /// Create a new bare metal server:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vultr = ediri.Vultr;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myServer = new Vultr.BareMetalServer("myServer", new()
    ///     {
    ///         OsId = 1743,
    ///         Plan = "vbm-4c-32gb",
    ///         Region = "ewr",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Create a new bare metal server with options:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vultr = ediri.Vultr;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myServer = new Vultr.BareMetalServer("myServer", new()
    ///     {
    ///         ActivationEmail = false,
    ///         EnableIpv6 = true,
    ///         Hostname = "my-server-hostname",
    ///         Label = "my-server-label",
    ///         OsId = 1743,
    ///         Plan = "vbm-4c-32gb",
    ///         Region = "ewr",
    ///         Tags = new[]
    ///         {
    ///             "my-server-tag",
    ///         },
    ///         UserData = "this is my user data",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Bare Metal Servers can be imported using the server `ID`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vultr:index/bareMetalServer:BareMetalServer my_server b6a859c5-b299-49dd-8888-b1abbc517d08
    /// ```
    /// </summary>
    [VultrResourceType("vultr:index/bareMetalServer:BareMetalServer")]
    public partial class BareMetalServer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether an activation email will be sent when the server is ready.
        /// </summary>
        [Output("activationEmail")]
        public Output<bool?> ActivationEmail { get; private set; } = null!;

        /// <summary>
        /// The ID of the Vultr application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications)
        /// </summary>
        [Output("appId")]
        public Output<int> AppId { get; private set; } = null!;

        /// <summary>
        /// A map of user-supplied variable keys and values for Vultr Marketplace apps. [See List Marketplace App Variables](https://www.vultr.com/api/#tag/marketplace/operation/list-marketplace-app-variables)
        /// </summary>
        [Output("appVariables")]
        public Output<ImmutableDictionary<string, string>?> AppVariables { get; private set; } = null!;

        /// <summary>
        /// The number of CPUs available on the server.
        /// </summary>
        [Output("cpuCount")]
        public Output<int> CpuCount { get; private set; } = null!;

        /// <summary>
        /// The date the server was added to your Vultr account.
        /// </summary>
        [Output("dateCreated")]
        public Output<string> DateCreated { get; private set; } = null!;

        /// <summary>
        /// The server's default password.
        /// </summary>
        [Output("defaultPassword")]
        public Output<string> DefaultPassword { get; private set; } = null!;

        /// <summary>
        /// The description of the disk(s) on the server.
        /// </summary>
        [Output("disk")]
        public Output<string> Disk { get; private set; } = null!;

        /// <summary>
        /// Whether the server has IPv6 networking activated.
        /// </summary>
        [Output("enableIpv6")]
        public Output<bool?> EnableIpv6 { get; private set; } = null!;

        /// <summary>
        /// The server's IPv4 gateway.
        /// </summary>
        [Output("gatewayV4")]
        public Output<string> GatewayV4 { get; private set; } = null!;

        /// <summary>
        /// The hostname to assign to the server.
        /// </summary>
        [Output("hostname")]
        public Output<string?> Hostname { get; private set; } = null!;

        /// <summary>
        /// The ID of the Vultr marketplace application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications) Note marketplace applications are denoted by type: `marketplace` and you must use the `image_id` not the id.
        /// </summary>
        [Output("imageId")]
        public Output<string> ImageId { get; private set; } = null!;

        /// <summary>
        /// A label for the server.
        /// </summary>
        [Output("label")]
        public Output<string?> Label { get; private set; } = null!;

        /// <summary>
        /// The MAC address associated with the server.
        /// </summary>
        [Output("macAddress")]
        public Output<int> MacAddress { get; private set; } = null!;

        /// <summary>
        /// The server's main IP address.
        /// </summary>
        [Output("mainIp")]
        public Output<string> MainIp { get; private set; } = null!;

        [Output("mdiskMode")]
        public Output<string?> MdiskMode { get; private set; } = null!;

        /// <summary>
        /// The server's IPv4 netmask.
        /// </summary>
        [Output("netmaskV4")]
        public Output<string> NetmaskV4 { get; private set; } = null!;

        /// <summary>
        /// The string description of the operating system installed on the server.
        /// </summary>
        [Output("os")]
        public Output<string> Os { get; private set; } = null!;

        /// <summary>
        /// The ID of the operating system to be installed on the server. [See List OS](https://www.vultr.com/api/#operation/list-os)
        /// </summary>
        [Output("osId")]
        public Output<int> OsId { get; private set; } = null!;

        [Output("persistentPxe")]
        public Output<bool?> PersistentPxe { get; private set; } = null!;

        /// <summary>
        /// The ID of the plan that you want the server to subscribe to. [See List Plans](https://www.vultr.com/api/#tag/plans)
        /// </summary>
        [Output("plan")]
        public Output<string> Plan { get; private set; } = null!;

        /// <summary>
        /// The amount of memory available on the server in MB.
        /// </summary>
        [Output("ram")]
        public Output<string> Ram { get; private set; } = null!;

        /// <summary>
        /// The ID of the region that the server is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The ID of the floating IP to use as the main IP of this server. [See Reserved IPs](https://www.vultr.com/api/#operation/list-reserved-ips)
        /// </summary>
        [Output("reservedIpv4")]
        public Output<string> ReservedIpv4 { get; private set; } = null!;

        /// <summary>
        /// The ID of the startup script you want added to the server.
        /// </summary>
        [Output("scriptId")]
        public Output<string?> ScriptId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Vultr snapshot that the server will restore for the initial installation. [See List Snapshots](https://www.vultr.com/api/#operation/list-snapshots)
        /// </summary>
        [Output("snapshotId")]
        public Output<string?> SnapshotId { get; private set; } = null!;

        /// <summary>
        /// A list of SSH key IDs to apply to the server on install (only valid for Linux/FreeBSD).
        /// </summary>
        [Output("sshKeyIds")]
        public Output<ImmutableArray<string>> SshKeyIds { get; private set; } = null!;

        /// <summary>
        /// The status of the server's subscription.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A list of tags to apply to the servier.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.
        /// </summary>
        [Output("userData")]
        public Output<string> UserData { get; private set; } = null!;

        /// <summary>
        /// The scheme used for the default user. Possible values are `root` or `limited` (linux servers only).
        /// </summary>
        [Output("userScheme")]
        public Output<string?> UserScheme { get; private set; } = null!;

        /// <summary>
        /// The main IPv6 network address.
        /// </summary>
        [Output("v6MainIp")]
        public Output<string> V6MainIp { get; private set; } = null!;

        /// <summary>
        /// The IPv6 subnet.
        /// </summary>
        [Output("v6Network")]
        public Output<string> V6Network { get; private set; } = null!;

        /// <summary>
        /// The IPv6 network size in bits.
        /// </summary>
        [Output("v6NetworkSize")]
        public Output<int> V6NetworkSize { get; private set; } = null!;

        /// <summary>
        /// A list of VPC 2.0 IDs to be attached to the server.
        /// </summary>
        [Output("vpc2Ids")]
        public Output<ImmutableArray<string>> Vpc2Ids { get; private set; } = null!;


        /// <summary>
        /// Create a BareMetalServer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BareMetalServer(string name, BareMetalServerArgs args, CustomResourceOptions? options = null)
            : base("vultr:index/bareMetalServer:BareMetalServer", name, args ?? new BareMetalServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BareMetalServer(string name, Input<string> id, BareMetalServerState? state = null, CustomResourceOptions? options = null)
            : base("vultr:index/bareMetalServer:BareMetalServer", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/dirien/pulumi-vultr",
                AdditionalSecretOutputs =
                {
                    "defaultPassword",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BareMetalServer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BareMetalServer Get(string name, Input<string> id, BareMetalServerState? state = null, CustomResourceOptions? options = null)
        {
            return new BareMetalServer(name, id, state, options);
        }
    }

    public sealed class BareMetalServerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether an activation email will be sent when the server is ready.
        /// </summary>
        [Input("activationEmail")]
        public Input<bool>? ActivationEmail { get; set; }

        /// <summary>
        /// The ID of the Vultr application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications)
        /// </summary>
        [Input("appId")]
        public Input<int>? AppId { get; set; }

        [Input("appVariables")]
        private InputMap<string>? _appVariables;

        /// <summary>
        /// A map of user-supplied variable keys and values for Vultr Marketplace apps. [See List Marketplace App Variables](https://www.vultr.com/api/#tag/marketplace/operation/list-marketplace-app-variables)
        /// </summary>
        public InputMap<string> AppVariables
        {
            get => _appVariables ?? (_appVariables = new InputMap<string>());
            set => _appVariables = value;
        }

        /// <summary>
        /// Whether the server has IPv6 networking activated.
        /// </summary>
        [Input("enableIpv6")]
        public Input<bool>? EnableIpv6 { get; set; }

        /// <summary>
        /// The hostname to assign to the server.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// The ID of the Vultr marketplace application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications) Note marketplace applications are denoted by type: `marketplace` and you must use the `image_id` not the id.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// A label for the server.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        [Input("mdiskMode")]
        public Input<string>? MdiskMode { get; set; }

        /// <summary>
        /// The ID of the operating system to be installed on the server. [See List OS](https://www.vultr.com/api/#operation/list-os)
        /// </summary>
        [Input("osId")]
        public Input<int>? OsId { get; set; }

        [Input("persistentPxe")]
        public Input<bool>? PersistentPxe { get; set; }

        /// <summary>
        /// The ID of the plan that you want the server to subscribe to. [See List Plans](https://www.vultr.com/api/#tag/plans)
        /// </summary>
        [Input("plan", required: true)]
        public Input<string> Plan { get; set; } = null!;

        /// <summary>
        /// The ID of the region that the server is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// The ID of the floating IP to use as the main IP of this server. [See Reserved IPs](https://www.vultr.com/api/#operation/list-reserved-ips)
        /// </summary>
        [Input("reservedIpv4")]
        public Input<string>? ReservedIpv4 { get; set; }

        /// <summary>
        /// The ID of the startup script you want added to the server.
        /// </summary>
        [Input("scriptId")]
        public Input<string>? ScriptId { get; set; }

        /// <summary>
        /// The ID of the Vultr snapshot that the server will restore for the initial installation. [See List Snapshots](https://www.vultr.com/api/#operation/list-snapshots)
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        [Input("sshKeyIds")]
        private InputList<string>? _sshKeyIds;

        /// <summary>
        /// A list of SSH key IDs to apply to the server on install (only valid for Linux/FreeBSD).
        /// </summary>
        public InputList<string> SshKeyIds
        {
            get => _sshKeyIds ?? (_sshKeyIds = new InputList<string>());
            set => _sshKeyIds = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tags to apply to the servier.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

        /// <summary>
        /// The scheme used for the default user. Possible values are `root` or `limited` (linux servers only).
        /// </summary>
        [Input("userScheme")]
        public Input<string>? UserScheme { get; set; }

        [Input("vpc2Ids")]
        private InputList<string>? _vpc2Ids;

        /// <summary>
        /// A list of VPC 2.0 IDs to be attached to the server.
        /// </summary>
        public InputList<string> Vpc2Ids
        {
            get => _vpc2Ids ?? (_vpc2Ids = new InputList<string>());
            set => _vpc2Ids = value;
        }

        public BareMetalServerArgs()
        {
        }
        public static new BareMetalServerArgs Empty => new BareMetalServerArgs();
    }

    public sealed class BareMetalServerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether an activation email will be sent when the server is ready.
        /// </summary>
        [Input("activationEmail")]
        public Input<bool>? ActivationEmail { get; set; }

        /// <summary>
        /// The ID of the Vultr application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications)
        /// </summary>
        [Input("appId")]
        public Input<int>? AppId { get; set; }

        [Input("appVariables")]
        private InputMap<string>? _appVariables;

        /// <summary>
        /// A map of user-supplied variable keys and values for Vultr Marketplace apps. [See List Marketplace App Variables](https://www.vultr.com/api/#tag/marketplace/operation/list-marketplace-app-variables)
        /// </summary>
        public InputMap<string> AppVariables
        {
            get => _appVariables ?? (_appVariables = new InputMap<string>());
            set => _appVariables = value;
        }

        /// <summary>
        /// The number of CPUs available on the server.
        /// </summary>
        [Input("cpuCount")]
        public Input<int>? CpuCount { get; set; }

        /// <summary>
        /// The date the server was added to your Vultr account.
        /// </summary>
        [Input("dateCreated")]
        public Input<string>? DateCreated { get; set; }

        [Input("defaultPassword")]
        private Input<string>? _defaultPassword;

        /// <summary>
        /// The server's default password.
        /// </summary>
        public Input<string>? DefaultPassword
        {
            get => _defaultPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _defaultPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The description of the disk(s) on the server.
        /// </summary>
        [Input("disk")]
        public Input<string>? Disk { get; set; }

        /// <summary>
        /// Whether the server has IPv6 networking activated.
        /// </summary>
        [Input("enableIpv6")]
        public Input<bool>? EnableIpv6 { get; set; }

        /// <summary>
        /// The server's IPv4 gateway.
        /// </summary>
        [Input("gatewayV4")]
        public Input<string>? GatewayV4 { get; set; }

        /// <summary>
        /// The hostname to assign to the server.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// The ID of the Vultr marketplace application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications) Note marketplace applications are denoted by type: `marketplace` and you must use the `image_id` not the id.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// A label for the server.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// The MAC address associated with the server.
        /// </summary>
        [Input("macAddress")]
        public Input<int>? MacAddress { get; set; }

        /// <summary>
        /// The server's main IP address.
        /// </summary>
        [Input("mainIp")]
        public Input<string>? MainIp { get; set; }

        [Input("mdiskMode")]
        public Input<string>? MdiskMode { get; set; }

        /// <summary>
        /// The server's IPv4 netmask.
        /// </summary>
        [Input("netmaskV4")]
        public Input<string>? NetmaskV4 { get; set; }

        /// <summary>
        /// The string description of the operating system installed on the server.
        /// </summary>
        [Input("os")]
        public Input<string>? Os { get; set; }

        /// <summary>
        /// The ID of the operating system to be installed on the server. [See List OS](https://www.vultr.com/api/#operation/list-os)
        /// </summary>
        [Input("osId")]
        public Input<int>? OsId { get; set; }

        [Input("persistentPxe")]
        public Input<bool>? PersistentPxe { get; set; }

        /// <summary>
        /// The ID of the plan that you want the server to subscribe to. [See List Plans](https://www.vultr.com/api/#tag/plans)
        /// </summary>
        [Input("plan")]
        public Input<string>? Plan { get; set; }

        /// <summary>
        /// The amount of memory available on the server in MB.
        /// </summary>
        [Input("ram")]
        public Input<string>? Ram { get; set; }

        /// <summary>
        /// The ID of the region that the server is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The ID of the floating IP to use as the main IP of this server. [See Reserved IPs](https://www.vultr.com/api/#operation/list-reserved-ips)
        /// </summary>
        [Input("reservedIpv4")]
        public Input<string>? ReservedIpv4 { get; set; }

        /// <summary>
        /// The ID of the startup script you want added to the server.
        /// </summary>
        [Input("scriptId")]
        public Input<string>? ScriptId { get; set; }

        /// <summary>
        /// The ID of the Vultr snapshot that the server will restore for the initial installation. [See List Snapshots](https://www.vultr.com/api/#operation/list-snapshots)
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        [Input("sshKeyIds")]
        private InputList<string>? _sshKeyIds;

        /// <summary>
        /// A list of SSH key IDs to apply to the server on install (only valid for Linux/FreeBSD).
        /// </summary>
        public InputList<string> SshKeyIds
        {
            get => _sshKeyIds ?? (_sshKeyIds = new InputList<string>());
            set => _sshKeyIds = value;
        }

        /// <summary>
        /// The status of the server's subscription.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tags to apply to the servier.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

        /// <summary>
        /// The scheme used for the default user. Possible values are `root` or `limited` (linux servers only).
        /// </summary>
        [Input("userScheme")]
        public Input<string>? UserScheme { get; set; }

        /// <summary>
        /// The main IPv6 network address.
        /// </summary>
        [Input("v6MainIp")]
        public Input<string>? V6MainIp { get; set; }

        /// <summary>
        /// The IPv6 subnet.
        /// </summary>
        [Input("v6Network")]
        public Input<string>? V6Network { get; set; }

        /// <summary>
        /// The IPv6 network size in bits.
        /// </summary>
        [Input("v6NetworkSize")]
        public Input<int>? V6NetworkSize { get; set; }

        [Input("vpc2Ids")]
        private InputList<string>? _vpc2Ids;

        /// <summary>
        /// A list of VPC 2.0 IDs to be attached to the server.
        /// </summary>
        public InputList<string> Vpc2Ids
        {
            get => _vpc2Ids ?? (_vpc2Ids = new InputList<string>());
            set => _vpc2Ids = value;
        }

        public BareMetalServerState()
        {
        }
        public static new BareMetalServerState Empty => new BareMetalServerState();
    }
}
