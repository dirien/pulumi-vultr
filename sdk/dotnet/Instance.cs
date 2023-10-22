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
    /// Provides a Vultr instance resource. This can be used to create, read, modify, and delete instances on your Vultr account.
    /// 
    /// ## Example Usage
    /// 
    /// Create a new instance:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vultr = ediri.Vultr;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myInstance = new Vultr.Instance("myInstance", new()
    ///     {
    ///         OsId = 1743,
    ///         Plan = "vc2-1c-1gb",
    ///         Region = "sea",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Create a new instance with options:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vultr = ediri.Vultr;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myInstance = new Vultr.Instance("myInstance", new()
    ///     {
    ///         ActivationEmail = false,
    ///         Backups = "enabled",
    ///         BackupsSchedule = new Vultr.Inputs.InstanceBackupsScheduleArgs
    ///         {
    ///             Type = "daily",
    ///         },
    ///         DdosProtection = true,
    ///         EnableIpv6 = true,
    ///         Hostname = "my-instance-hostname",
    ///         Label = "my-instance-label",
    ///         OsId = 1743,
    ///         Plan = "vc2-1c-1gb",
    ///         Region = "sea",
    ///         Tags = new[]
    ///         {
    ///             "my-instance-tag",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Instances can be imported using the instance `ID`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import vultr:index/instance:Instance my_instance b6a859c5-b299-49dd-8888-b1abbc517d08
    /// ```
    /// </summary>
    [VultrResourceType("vultr:index/instance:Instance")]
    public partial class Instance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether an activation email will be sent when the server is ready.
        /// </summary>
        [Output("activationEmail")]
        public Output<bool?> ActivationEmail { get; private set; } = null!;

        /// <summary>
        /// The server's allowed bandwidth usage in GB.
        /// </summary>
        [Output("allowedBandwidth")]
        public Output<int> AllowedBandwidth { get; private set; } = null!;

        /// <summary>
        /// The ID of the Vultr application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications)
        /// </summary>
        [Output("appId")]
        public Output<int> AppId { get; private set; } = null!;

        /// <summary>
        /// Whether automatic backups will be enabled for this server (these have an extra charge associated with them). Values can be enabled or disabled.
        /// </summary>
        [Output("backups")]
        public Output<string?> Backups { get; private set; } = null!;

        /// <summary>
        /// A block that defines the way backups should be scheduled. While this is an optional field if `backups` are `enabled` this field is mandatory. The configuration of a `backups_schedule` is listed below.
        /// </summary>
        [Output("backupsSchedule")]
        public Output<Outputs.InstanceBackupsSchedule?> BackupsSchedule { get; private set; } = null!;

        /// <summary>
        /// The date the server was added to your Vultr account.
        /// </summary>
        [Output("dateCreated")]
        public Output<string> DateCreated { get; private set; } = null!;

        /// <summary>
        /// Whether DDOS protection will be enabled on the server (there is an additional charge for this).
        /// </summary>
        [Output("ddosProtection")]
        public Output<bool?> DdosProtection { get; private set; } = null!;

        /// <summary>
        /// The server's default password.
        /// </summary>
        [Output("defaultPassword")]
        public Output<string> DefaultPassword { get; private set; } = null!;

        /// <summary>
        /// The description of the disk(s) on the server.
        /// </summary>
        [Output("disk")]
        public Output<int> Disk { get; private set; } = null!;

        /// <summary>
        /// Whether the server has IPv6 networking activated.
        /// </summary>
        [Output("enableIpv6")]
        public Output<bool?> EnableIpv6 { get; private set; } = null!;

        /// <summary>
        /// Array of which features are enabled.
        /// </summary>
        [Output("features")]
        public Output<ImmutableArray<string>> Features { get; private set; } = null!;

        /// <summary>
        /// The ID of the firewall group to assign to the server.
        /// </summary>
        [Output("firewallGroupId")]
        public Output<string> FirewallGroupId { get; private set; } = null!;

        /// <summary>
        /// The server's IPv4 gateway.
        /// </summary>
        [Output("gatewayV4")]
        public Output<string> GatewayV4 { get; private set; } = null!;

        /// <summary>
        /// The hostname to assign to the server.
        /// </summary>
        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        /// <summary>
        /// The ID of the Vultr marketplace application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications) Note marketplace applications are denoted by type: `marketplace` and you must use the `image_id` not the id.
        /// </summary>
        [Output("imageId")]
        public Output<string> ImageId { get; private set; } = null!;

        /// <summary>
        /// The server's internal IP address.
        /// </summary>
        [Output("internalIp")]
        public Output<string> InternalIp { get; private set; } = null!;

        /// <summary>
        /// The ID of the ISO file to be installed on the server. [See List ISO](https://www.vultr.com/api/#operation/list-isos)
        /// </summary>
        [Output("isoId")]
        public Output<string?> IsoId { get; private set; } = null!;

        /// <summary>
        /// The server's current KVM URL. This URL will change periodically. It is not advised to cache this value.
        /// </summary>
        [Output("kvm")]
        public Output<string> Kvm { get; private set; } = null!;

        /// <summary>
        /// A label for the server.
        /// </summary>
        [Output("label")]
        public Output<string> Label { get; private set; } = null!;

        /// <summary>
        /// The server's main IP address.
        /// </summary>
        [Output("mainIp")]
        public Output<string> MainIp { get; private set; } = null!;

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

        /// <summary>
        /// The ID of the plan that you want the instance to subscribe to. [See List Plans](https://www.vultr.com/api/#tag/plans)
        /// </summary>
        [Output("plan")]
        public Output<string> Plan { get; private set; } = null!;

        /// <summary>
        /// Whether the server is powered on or not.
        /// </summary>
        [Output("powerStatus")]
        public Output<string> PowerStatus { get; private set; } = null!;

        /// <summary>
        /// (Deprecated: use `vpc_ids` instead) A list of private network IDs to be attached to the server.
        /// </summary>
        [Output("privateNetworkIds")]
        public Output<ImmutableArray<string>> PrivateNetworkIds { get; private set; } = null!;

        /// <summary>
        /// The amount of memory available on the server in MB.
        /// </summary>
        [Output("ram")]
        public Output<int> Ram { get; private set; } = null!;

        /// <summary>
        /// The ID of the region that the instance is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// ID of the floating IP to use as the main IP of this server.
        /// </summary>
        [Output("reservedIpId")]
        public Output<string> ReservedIpId { get; private set; } = null!;

        /// <summary>
        /// The ID of the startup script you want added to the server.
        /// </summary>
        [Output("scriptId")]
        public Output<string> ScriptId { get; private set; } = null!;

        /// <summary>
        /// A more detailed server status (none, locked, installingbooting, isomounting, ok).
        /// </summary>
        [Output("serverStatus")]
        public Output<string> ServerStatus { get; private set; } = null!;

        /// <summary>
        /// The ID of the Vultr snapshot that the server will restore for the initial installation. [See List Snapshots](https://www.vultr.com/api/#operation/list-snapshots)
        /// </summary>
        [Output("snapshotId")]
        public Output<string> SnapshotId { get; private set; } = null!;

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
        /// A list of tags to apply to the instance.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.
        /// </summary>
        [Output("userData")]
        public Output<string> UserData { get; private set; } = null!;

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
        /// The number of virtual CPUs available on the server.
        /// </summary>
        [Output("vcpuCount")]
        public Output<int> VcpuCount { get; private set; } = null!;

        /// <summary>
        /// A list of VPC 2.0 IDs to be attached to the server.
        /// </summary>
        [Output("vpc2Ids")]
        public Output<ImmutableArray<string>> Vpc2Ids { get; private set; } = null!;

        /// <summary>
        /// A list of VPC IDs to be attached to the server.
        /// </summary>
        [Output("vpcIds")]
        public Output<ImmutableArray<string>> VpcIds { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("vultr:index/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("vultr:index/instance:Instance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : global::Pulumi.ResourceArgs
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

        /// <summary>
        /// Whether automatic backups will be enabled for this server (these have an extra charge associated with them). Values can be enabled or disabled.
        /// </summary>
        [Input("backups")]
        public Input<string>? Backups { get; set; }

        /// <summary>
        /// A block that defines the way backups should be scheduled. While this is an optional field if `backups` are `enabled` this field is mandatory. The configuration of a `backups_schedule` is listed below.
        /// </summary>
        [Input("backupsSchedule")]
        public Input<Inputs.InstanceBackupsScheduleArgs>? BackupsSchedule { get; set; }

        /// <summary>
        /// Whether DDOS protection will be enabled on the server (there is an additional charge for this).
        /// </summary>
        [Input("ddosProtection")]
        public Input<bool>? DdosProtection { get; set; }

        /// <summary>
        /// Whether the server has IPv6 networking activated.
        /// </summary>
        [Input("enableIpv6")]
        public Input<bool>? EnableIpv6 { get; set; }

        /// <summary>
        /// The ID of the firewall group to assign to the server.
        /// </summary>
        [Input("firewallGroupId")]
        public Input<string>? FirewallGroupId { get; set; }

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
        /// The ID of the ISO file to be installed on the server. [See List ISO](https://www.vultr.com/api/#operation/list-isos)
        /// </summary>
        [Input("isoId")]
        public Input<string>? IsoId { get; set; }

        /// <summary>
        /// A label for the server.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// The ID of the operating system to be installed on the server. [See List OS](https://www.vultr.com/api/#operation/list-os)
        /// </summary>
        [Input("osId")]
        public Input<int>? OsId { get; set; }

        /// <summary>
        /// The ID of the plan that you want the instance to subscribe to. [See List Plans](https://www.vultr.com/api/#tag/plans)
        /// </summary>
        [Input("plan", required: true)]
        public Input<string> Plan { get; set; } = null!;

        [Input("privateNetworkIds")]
        private InputList<string>? _privateNetworkIds;

        /// <summary>
        /// (Deprecated: use `vpc_ids` instead) A list of private network IDs to be attached to the server.
        /// </summary>
        [Obsolete(@"private_network_ids has been deprecated and should no longer be used. Instead, use vpc_ids")]
        public InputList<string> PrivateNetworkIds
        {
            get => _privateNetworkIds ?? (_privateNetworkIds = new InputList<string>());
            set => _privateNetworkIds = value;
        }

        /// <summary>
        /// The ID of the region that the instance is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// ID of the floating IP to use as the main IP of this server.
        /// </summary>
        [Input("reservedIpId")]
        public Input<string>? ReservedIpId { get; set; }

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
        /// A list of tags to apply to the instance.
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

        [Input("vpcIds")]
        private InputList<string>? _vpcIds;

        /// <summary>
        /// A list of VPC IDs to be attached to the server.
        /// </summary>
        public InputList<string> VpcIds
        {
            get => _vpcIds ?? (_vpcIds = new InputList<string>());
            set => _vpcIds = value;
        }

        public InstanceArgs()
        {
        }
        public static new InstanceArgs Empty => new InstanceArgs();
    }

    public sealed class InstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether an activation email will be sent when the server is ready.
        /// </summary>
        [Input("activationEmail")]
        public Input<bool>? ActivationEmail { get; set; }

        /// <summary>
        /// The server's allowed bandwidth usage in GB.
        /// </summary>
        [Input("allowedBandwidth")]
        public Input<int>? AllowedBandwidth { get; set; }

        /// <summary>
        /// The ID of the Vultr application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications)
        /// </summary>
        [Input("appId")]
        public Input<int>? AppId { get; set; }

        /// <summary>
        /// Whether automatic backups will be enabled for this server (these have an extra charge associated with them). Values can be enabled or disabled.
        /// </summary>
        [Input("backups")]
        public Input<string>? Backups { get; set; }

        /// <summary>
        /// A block that defines the way backups should be scheduled. While this is an optional field if `backups` are `enabled` this field is mandatory. The configuration of a `backups_schedule` is listed below.
        /// </summary>
        [Input("backupsSchedule")]
        public Input<Inputs.InstanceBackupsScheduleGetArgs>? BackupsSchedule { get; set; }

        /// <summary>
        /// The date the server was added to your Vultr account.
        /// </summary>
        [Input("dateCreated")]
        public Input<string>? DateCreated { get; set; }

        /// <summary>
        /// Whether DDOS protection will be enabled on the server (there is an additional charge for this).
        /// </summary>
        [Input("ddosProtection")]
        public Input<bool>? DdosProtection { get; set; }

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
        public Input<int>? Disk { get; set; }

        /// <summary>
        /// Whether the server has IPv6 networking activated.
        /// </summary>
        [Input("enableIpv6")]
        public Input<bool>? EnableIpv6 { get; set; }

        [Input("features")]
        private InputList<string>? _features;

        /// <summary>
        /// Array of which features are enabled.
        /// </summary>
        public InputList<string> Features
        {
            get => _features ?? (_features = new InputList<string>());
            set => _features = value;
        }

        /// <summary>
        /// The ID of the firewall group to assign to the server.
        /// </summary>
        [Input("firewallGroupId")]
        public Input<string>? FirewallGroupId { get; set; }

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
        /// The server's internal IP address.
        /// </summary>
        [Input("internalIp")]
        public Input<string>? InternalIp { get; set; }

        /// <summary>
        /// The ID of the ISO file to be installed on the server. [See List ISO](https://www.vultr.com/api/#operation/list-isos)
        /// </summary>
        [Input("isoId")]
        public Input<string>? IsoId { get; set; }

        /// <summary>
        /// The server's current KVM URL. This URL will change periodically. It is not advised to cache this value.
        /// </summary>
        [Input("kvm")]
        public Input<string>? Kvm { get; set; }

        /// <summary>
        /// A label for the server.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// The server's main IP address.
        /// </summary>
        [Input("mainIp")]
        public Input<string>? MainIp { get; set; }

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

        /// <summary>
        /// The ID of the plan that you want the instance to subscribe to. [See List Plans](https://www.vultr.com/api/#tag/plans)
        /// </summary>
        [Input("plan")]
        public Input<string>? Plan { get; set; }

        /// <summary>
        /// Whether the server is powered on or not.
        /// </summary>
        [Input("powerStatus")]
        public Input<string>? PowerStatus { get; set; }

        [Input("privateNetworkIds")]
        private InputList<string>? _privateNetworkIds;

        /// <summary>
        /// (Deprecated: use `vpc_ids` instead) A list of private network IDs to be attached to the server.
        /// </summary>
        [Obsolete(@"private_network_ids has been deprecated and should no longer be used. Instead, use vpc_ids")]
        public InputList<string> PrivateNetworkIds
        {
            get => _privateNetworkIds ?? (_privateNetworkIds = new InputList<string>());
            set => _privateNetworkIds = value;
        }

        /// <summary>
        /// The amount of memory available on the server in MB.
        /// </summary>
        [Input("ram")]
        public Input<int>? Ram { get; set; }

        /// <summary>
        /// The ID of the region that the instance is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// ID of the floating IP to use as the main IP of this server.
        /// </summary>
        [Input("reservedIpId")]
        public Input<string>? ReservedIpId { get; set; }

        /// <summary>
        /// The ID of the startup script you want added to the server.
        /// </summary>
        [Input("scriptId")]
        public Input<string>? ScriptId { get; set; }

        /// <summary>
        /// A more detailed server status (none, locked, installingbooting, isomounting, ok).
        /// </summary>
        [Input("serverStatus")]
        public Input<string>? ServerStatus { get; set; }

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
        /// A list of tags to apply to the instance.
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

        /// <summary>
        /// The number of virtual CPUs available on the server.
        /// </summary>
        [Input("vcpuCount")]
        public Input<int>? VcpuCount { get; set; }

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

        [Input("vpcIds")]
        private InputList<string>? _vpcIds;

        /// <summary>
        /// A list of VPC IDs to be attached to the server.
        /// </summary>
        public InputList<string> VpcIds
        {
            get => _vpcIds ?? (_vpcIds = new InputList<string>());
            set => _vpcIds = value;
        }

        public InstanceState()
        {
        }
        public static new InstanceState Empty => new InstanceState();
    }
}
