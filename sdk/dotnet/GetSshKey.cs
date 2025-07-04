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
    public static class GetSshKey
    {
        /// <summary>
        /// Get information about a Vultr SSH key. This data source provides the name, public SSH key, and the creation date for your Vultr SSH key.
        /// 
        /// ## Example Usage
        /// 
        /// Get the information for an SSH key by `name`:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mySshKey = Vultr.GetSshKey.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Vultr.Inputs.GetSshKeyFilterInputArgs
        ///             {
        ///                 Name = "name",
        ///                 Values = new[]
        ///                 {
        ///                     "my-ssh-key-name",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetSshKeyResult> InvokeAsync(GetSshKeyArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSshKeyResult>("vultr:index/getSshKey:getSshKey", args ?? new GetSshKeyArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Vultr SSH key. This data source provides the name, public SSH key, and the creation date for your Vultr SSH key.
        /// 
        /// ## Example Usage
        /// 
        /// Get the information for an SSH key by `name`:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mySshKey = Vultr.GetSshKey.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Vultr.Inputs.GetSshKeyFilterInputArgs
        ///             {
        ///                 Name = "name",
        ///                 Values = new[]
        ///                 {
        ///                     "my-ssh-key-name",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetSshKeyResult> Invoke(GetSshKeyInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSshKeyResult>("vultr:index/getSshKey:getSshKey", args ?? new GetSshKeyInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Vultr SSH key. This data source provides the name, public SSH key, and the creation date for your Vultr SSH key.
        /// 
        /// ## Example Usage
        /// 
        /// Get the information for an SSH key by `name`:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mySshKey = Vultr.GetSshKey.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Vultr.Inputs.GetSshKeyFilterInputArgs
        ///             {
        ///                 Name = "name",
        ///                 Values = new[]
        ///                 {
        ///                     "my-ssh-key-name",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetSshKeyResult> Invoke(GetSshKeyInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSshKeyResult>("vultr:index/getSshKey:getSshKey", args ?? new GetSshKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSshKeyArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetSshKeyFilterArgs>? _filters;

        /// <summary>
        /// Query parameters for finding SSH keys.
        /// </summary>
        public List<Inputs.GetSshKeyFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetSshKeyFilterArgs>());
            set => _filters = value;
        }

        public GetSshKeyArgs()
        {
        }
        public static new GetSshKeyArgs Empty => new GetSshKeyArgs();
    }

    public sealed class GetSshKeyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetSshKeyFilterInputArgs>? _filters;

        /// <summary>
        /// Query parameters for finding SSH keys.
        /// </summary>
        public InputList<Inputs.GetSshKeyFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetSshKeyFilterInputArgs>());
            set => _filters = value;
        }

        public GetSshKeyInvokeArgs()
        {
        }
        public static new GetSshKeyInvokeArgs Empty => new GetSshKeyInvokeArgs();
    }


    [OutputType]
    public sealed class GetSshKeyResult
    {
        /// <summary>
        /// The date the SSH key was added to your Vultr account.
        /// </summary>
        public readonly string DateCreated;
        public readonly ImmutableArray<Outputs.GetSshKeyFilterResult> Filters;
        /// <summary>
        /// The ID of the SSH key.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the SSH key.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The public SSH key.
        /// </summary>
        public readonly string SshKey;

        [OutputConstructor]
        private GetSshKeyResult(
            string dateCreated,

            ImmutableArray<Outputs.GetSshKeyFilterResult> filters,

            string id,

            string name,

            string sshKey)
        {
            DateCreated = dateCreated;
            Filters = filters;
            Id = id;
            Name = name;
            SshKey = sshKey;
        }
    }
}
