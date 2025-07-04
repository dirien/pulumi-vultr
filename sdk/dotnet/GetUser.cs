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
    public static class GetUser
    {
        /// <summary>
        /// Get information about a Vultr user associated with your account. This data source provides the name, email, access control list, and API status for a Vultr user associated with your account.
        /// 
        /// ## Example Usage
        /// 
        /// Get the information for a user by `email`:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myUser = Vultr.GetUser.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Vultr.Inputs.GetUserFilterInputArgs
        ///             {
        ///                 Name = "email",
        ///                 Values = new[]
        ///                 {
        ///                     "jdoe@example.com",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// Get the information for a user by `name`:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myUser = Vultr.GetUser.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Vultr.Inputs.GetUserFilterInputArgs
        ///             {
        ///                 Name = "name",
        ///                 Values = new[]
        ///                 {
        ///                     "John Doe",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetUserResult> InvokeAsync(GetUserArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserResult>("vultr:index/getUser:getUser", args ?? new GetUserArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Vultr user associated with your account. This data source provides the name, email, access control list, and API status for a Vultr user associated with your account.
        /// 
        /// ## Example Usage
        /// 
        /// Get the information for a user by `email`:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myUser = Vultr.GetUser.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Vultr.Inputs.GetUserFilterInputArgs
        ///             {
        ///                 Name = "email",
        ///                 Values = new[]
        ///                 {
        ///                     "jdoe@example.com",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// Get the information for a user by `name`:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myUser = Vultr.GetUser.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Vultr.Inputs.GetUserFilterInputArgs
        ///             {
        ///                 Name = "name",
        ///                 Values = new[]
        ///                 {
        ///                     "John Doe",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetUserResult> Invoke(GetUserInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserResult>("vultr:index/getUser:getUser", args ?? new GetUserInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Vultr user associated with your account. This data source provides the name, email, access control list, and API status for a Vultr user associated with your account.
        /// 
        /// ## Example Usage
        /// 
        /// Get the information for a user by `email`:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myUser = Vultr.GetUser.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Vultr.Inputs.GetUserFilterInputArgs
        ///             {
        ///                 Name = "email",
        ///                 Values = new[]
        ///                 {
        ///                     "jdoe@example.com",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// Get the information for a user by `name`:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myUser = Vultr.GetUser.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Vultr.Inputs.GetUserFilterInputArgs
        ///             {
        ///                 Name = "name",
        ///                 Values = new[]
        ///                 {
        ///                     "John Doe",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetUserResult> Invoke(GetUserInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserResult>("vultr:index/getUser:getUser", args ?? new GetUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetUserFilterArgs>? _filters;

        /// <summary>
        /// Query parameters for finding users.
        /// </summary>
        public List<Inputs.GetUserFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetUserFilterArgs>());
            set => _filters = value;
        }

        public GetUserArgs()
        {
        }
        public static new GetUserArgs Empty => new GetUserArgs();
    }

    public sealed class GetUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetUserFilterInputArgs>? _filters;

        /// <summary>
        /// Query parameters for finding users.
        /// </summary>
        public InputList<Inputs.GetUserFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetUserFilterInputArgs>());
            set => _filters = value;
        }

        public GetUserInvokeArgs()
        {
        }
        public static new GetUserInvokeArgs Empty => new GetUserInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserResult
    {
        /// <summary>
        /// The access control list for the user.
        /// </summary>
        public readonly ImmutableArray<string> Acls;
        /// <summary>
        /// Whether API is enabled for the user.
        /// </summary>
        public readonly bool ApiEnabled;
        /// <summary>
        /// The email of the user.
        /// </summary>
        public readonly string Email;
        public readonly ImmutableArray<Outputs.GetUserFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the user.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetUserResult(
            ImmutableArray<string> acls,

            bool apiEnabled,

            string email,

            ImmutableArray<Outputs.GetUserFilterResult> filters,

            string id,

            string name)
        {
            Acls = acls;
            ApiEnabled = apiEnabled;
            Email = email;
            Filters = filters;
            Id = id;
            Name = name;
        }
    }
}
