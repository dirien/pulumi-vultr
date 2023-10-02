// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace ediri.Vultr
{
    public static class Config
    {
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("vultr");

        private static readonly __Value<string?> _apiKey = new __Value<string?>(() => __config.Get("apiKey") ?? Utilities.GetEnv("VULTR_API_KEY"));
        /// <summary>
        /// The API Key that allows interaction with the API
        /// </summary>
        public static string? ApiKey
        {
            get => _apiKey.Get();
            set => _apiKey.Set(value);
        }

        private static readonly __Value<int?> _rateLimit = new __Value<int?>(() => __config.GetInt32("rateLimit") ?? 500);
        /// <summary>
        /// Allows users to set the speed of API calls to work with the Vultr Rate Limit
        /// </summary>
        public static int? RateLimit
        {
            get => _rateLimit.Get();
            set => _rateLimit.Set(value);
        }

        private static readonly __Value<int?> _retryLimit = new __Value<int?>(() => __config.GetInt32("retryLimit") ?? 3);
        /// <summary>
        /// Allows users to set the maximum number of retries allowed for a failed API call.
        /// </summary>
        public static int? RetryLimit
        {
            get => _retryLimit.Get();
            set => _retryLimit.Set(value);
        }

    }
}
