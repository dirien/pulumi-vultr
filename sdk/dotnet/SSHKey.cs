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
    /// Provides a Vultr SSH key resource. This can be used to create, read, modify, and delete SSH keys.
    /// 
    /// ## Example Usage
    /// 
    /// Create an SSH key
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Vultr = ediri.Vultr;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mySshKey = new Vultr.SSHKey("mySshKey", new()
    ///     {
    ///         SshKey = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCyVGaw1PuEl98f4/7Kq3O9ZIvDw2OFOSXAFVqilSFNkHlefm1iMtPeqsIBp2t9cbGUf55xNDULz/bD/4BCV43yZ5lh0cUYuXALg9NI29ui7PEGReXjSpNwUD6ceN/78YOK41KAcecq+SS0bJ4b4amKZIJG3JWmDKljtv1dmSBCrTmEAQaOorxqGGBYmZS7NQumRe4lav5r6wOs8OACMANE1ejkeZsGFzJFNqvr5DuHdDL5FAudW23me3BDmrM9ifUzzjl1Jwku3bnRaCcjaxH8oTumt1a00mWci/1qUlaVFft085yvVq7KZbF2OPPbl+erDW91+EZ2FgEi+v1/CSJ5 your_username@hostname",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// SSH keys can be imported using the SSH key `ID`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import vultr:index/sSHKey:SSHKey my_key 6b0876a7-f709-41ba-aed8-abed9d38ae45
    /// ```
    /// </summary>
    [VultrResourceType("vultr:index/sSHKey:SSHKey")]
    public partial class SSHKey : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The date the SSH key was added to your Vultr account.
        /// </summary>
        [Output("dateCreated")]
        public Output<string> DateCreated { get; private set; } = null!;

        /// <summary>
        /// The name/label of the SSH key.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The public SSH key.
        /// </summary>
        [Output("sshKey")]
        public Output<string> SshKey { get; private set; } = null!;


        /// <summary>
        /// Create a SSHKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SSHKey(string name, SSHKeyArgs args, CustomResourceOptions? options = null)
            : base("vultr:index/sSHKey:SSHKey", name, args ?? new SSHKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SSHKey(string name, Input<string> id, SSHKeyState? state = null, CustomResourceOptions? options = null)
            : base("vultr:index/sSHKey:SSHKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SSHKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SSHKey Get(string name, Input<string> id, SSHKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new SSHKey(name, id, state, options);
        }
    }

    public sealed class SSHKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name/label of the SSH key.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The public SSH key.
        /// </summary>
        [Input("sshKey", required: true)]
        public Input<string> SshKey { get; set; } = null!;

        public SSHKeyArgs()
        {
        }
        public static new SSHKeyArgs Empty => new SSHKeyArgs();
    }

    public sealed class SSHKeyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date the SSH key was added to your Vultr account.
        /// </summary>
        [Input("dateCreated")]
        public Input<string>? DateCreated { get; set; }

        /// <summary>
        /// The name/label of the SSH key.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The public SSH key.
        /// </summary>
        [Input("sshKey")]
        public Input<string>? SshKey { get; set; }

        public SSHKeyState()
        {
        }
        public static new SSHKeyState Empty => new SSHKeyState();
    }
}