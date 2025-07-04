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
    /// Provides a Vultr User resource. This can be used to create, read, modify, and delete Users.
    /// 
    /// ## Example Usage
    /// 
    /// Create a new User without any ACLs
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vultr = ediri.Vultr;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myUser = new Vultr.User("myUser", new()
    ///     {
    ///         ApiEnabled = true,
    ///         Email = "user@vultr.com",
    ///         Password = "myP@ssw0rd",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Create a new User with all ACLs
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vultr = ediri.Vultr;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myUser = new Vultr.User("myUser", new()
    ///     {
    ///         Acls = new[]
    ///         {
    ///             "manage_users",
    ///             "subscriptions",
    ///             "provisioning",
    ///             "billing",
    ///             "support",
    ///             "abuse",
    ///             "dns",
    ///             "upgrade",
    ///         },
    ///         ApiEnabled = true,
    ///         Email = "user@vultr.com",
    ///         Password = "myP@ssw0rd",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Users can be imported using the User `ID`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vultr:index/user:User myuser 1345fef0-8ed3-4a66-bd8c-822a7b7bd05a
    /// ```
    /// </summary>
    [VultrResourceType("vultr:index/user:User")]
    public partial class User : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The access control list for the user.
        /// </summary>
        [Output("acls")]
        public Output<ImmutableArray<string>> Acls { get; private set; } = null!;

        /// <summary>
        /// Whether API is enabled for the user. Default behavior is set to enabled.
        /// </summary>
        [Output("apiEnabled")]
        public Output<bool?> ApiEnabled { get; private set; } = null!;

        [Output("apiKey")]
        public Output<string> ApiKey { get; private set; } = null!;

        /// <summary>
        /// Email for this user.
        /// </summary>
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        /// <summary>
        /// Name for this user.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Password for this user.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("vultr:index/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("vultr:index/user:User", name, state, MakeResourceOptions(options, id))
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
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : global::Pulumi.ResourceArgs
    {
        [Input("acls")]
        private InputList<string>? _acls;

        /// <summary>
        /// The access control list for the user.
        /// </summary>
        public InputList<string> Acls
        {
            get => _acls ?? (_acls = new InputList<string>());
            set => _acls = value;
        }

        /// <summary>
        /// Whether API is enabled for the user. Default behavior is set to enabled.
        /// </summary>
        [Input("apiEnabled")]
        public Input<bool>? ApiEnabled { get; set; }

        /// <summary>
        /// Email for this user.
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        /// <summary>
        /// Name for this user.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// Password for this user.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public UserArgs()
        {
        }
        public static new UserArgs Empty => new UserArgs();
    }

    public sealed class UserState : global::Pulumi.ResourceArgs
    {
        [Input("acls")]
        private InputList<string>? _acls;

        /// <summary>
        /// The access control list for the user.
        /// </summary>
        public InputList<string> Acls
        {
            get => _acls ?? (_acls = new InputList<string>());
            set => _acls = value;
        }

        /// <summary>
        /// Whether API is enabled for the user. Default behavior is set to enabled.
        /// </summary>
        [Input("apiEnabled")]
        public Input<bool>? ApiEnabled { get; set; }

        [Input("apiKey")]
        public Input<string>? ApiKey { get; set; }

        /// <summary>
        /// Email for this user.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// Name for this user.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password for this user.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public UserState()
        {
        }
        public static new UserState Empty => new UserState();
    }
}
