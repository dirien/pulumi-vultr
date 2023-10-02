// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderArgs Empty = new ProviderArgs();

    /**
     * The API Key that allows interaction with the API
     * 
     */
    @Import(name="apiKey")
    private @Nullable Output<String> apiKey;

    /**
     * @return The API Key that allows interaction with the API
     * 
     */
    public Optional<Output<String>> apiKey() {
        return Optional.ofNullable(this.apiKey);
    }

    /**
     * Allows users to set the speed of API calls to work with the Vultr Rate Limit
     * 
     */
    @Import(name="rateLimit", json=true)
    private @Nullable Output<Integer> rateLimit;

    /**
     * @return Allows users to set the speed of API calls to work with the Vultr Rate Limit
     * 
     */
    public Optional<Output<Integer>> rateLimit() {
        return Optional.ofNullable(this.rateLimit);
    }

    /**
     * Allows users to set the maximum number of retries allowed for a failed API call.
     * 
     */
    @Import(name="retryLimit", json=true)
    private @Nullable Output<Integer> retryLimit;

    /**
     * @return Allows users to set the maximum number of retries allowed for a failed API call.
     * 
     */
    public Optional<Output<Integer>> retryLimit() {
        return Optional.ofNullable(this.retryLimit);
    }

    private ProviderArgs() {}

    private ProviderArgs(ProviderArgs $) {
        this.apiKey = $.apiKey;
        this.rateLimit = $.rateLimit;
        this.retryLimit = $.retryLimit;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderArgs $;

        public Builder() {
            $ = new ProviderArgs();
        }

        public Builder(ProviderArgs defaults) {
            $ = new ProviderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param apiKey The API Key that allows interaction with the API
         * 
         * @return builder
         * 
         */
        public Builder apiKey(@Nullable Output<String> apiKey) {
            $.apiKey = apiKey;
            return this;
        }

        /**
         * @param apiKey The API Key that allows interaction with the API
         * 
         * @return builder
         * 
         */
        public Builder apiKey(String apiKey) {
            return apiKey(Output.of(apiKey));
        }

        /**
         * @param rateLimit Allows users to set the speed of API calls to work with the Vultr Rate Limit
         * 
         * @return builder
         * 
         */
        public Builder rateLimit(@Nullable Output<Integer> rateLimit) {
            $.rateLimit = rateLimit;
            return this;
        }

        /**
         * @param rateLimit Allows users to set the speed of API calls to work with the Vultr Rate Limit
         * 
         * @return builder
         * 
         */
        public Builder rateLimit(Integer rateLimit) {
            return rateLimit(Output.of(rateLimit));
        }

        /**
         * @param retryLimit Allows users to set the maximum number of retries allowed for a failed API call.
         * 
         * @return builder
         * 
         */
        public Builder retryLimit(@Nullable Output<Integer> retryLimit) {
            $.retryLimit = retryLimit;
            return this;
        }

        /**
         * @param retryLimit Allows users to set the maximum number of retries allowed for a failed API call.
         * 
         * @return builder
         * 
         */
        public Builder retryLimit(Integer retryLimit) {
            return retryLimit(Output.of(retryLimit));
        }

        public ProviderArgs build() {
            $.apiKey = Codegen.stringProp("apiKey").secret().arg($.apiKey).env("VULTR_API_KEY").getNullable();
            $.rateLimit = Codegen.integerProp("rateLimit").output().arg($.rateLimit).def(500).getNullable();
            $.retryLimit = Codegen.integerProp("retryLimit").output().arg($.retryLimit).def(3).getNullable();
            return $;
        }
    }

}
