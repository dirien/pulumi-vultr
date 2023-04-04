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
    public static class GetAccount
    {
        /// <summary>
        /// Get information about your Vultr account. This data source provides the balance, pending charges, last payment date, and last payment amount for your Vultr account.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Get the information for an account:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vultr = Pulumi.Vultr;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myAccount = Vultr.GetAccount.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAccountResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccountResult>("vultr:index/getAccount:getAccount", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetAccountResult
    {
        /// <summary>
        /// The access control list on your Vultr account.
        /// </summary>
        public readonly ImmutableArray<string> Acls;
        /// <summary>
        /// The current balance on your Vultr account.
        /// </summary>
        public readonly double Balance;
        /// <summary>
        /// The email address on your Vultr account.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The amount of the last payment made on your Vultr account.
        /// </summary>
        public readonly double LastPaymentAmount;
        /// <summary>
        /// The date of the last payment made on your Vultr account.
        /// </summary>
        public readonly string LastPaymentDate;
        /// <summary>
        /// The name on your Vultr account.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The pending charges on your Vultr account.
        /// </summary>
        public readonly double PendingCharges;

        [OutputConstructor]
        private GetAccountResult(
            ImmutableArray<string> acls,

            double balance,

            string email,

            string id,

            double lastPaymentAmount,

            string lastPaymentDate,

            string name,

            double pendingCharges)
        {
            Acls = acls;
            Balance = balance;
            Email = email;
            Id = id;
            LastPaymentAmount = lastPaymentAmount;
            LastPaymentDate = lastPaymentDate;
            Name = name;
            PendingCharges = pendingCharges;
        }
    }
}
