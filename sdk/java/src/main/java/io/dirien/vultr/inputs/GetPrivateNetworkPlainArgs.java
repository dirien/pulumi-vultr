// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.inputs;

import com.pulumi.core.annotations.Import;
import io.dirien.vultr.inputs.GetPrivateNetworkFilter;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPrivateNetworkPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPrivateNetworkPlainArgs Empty = new GetPrivateNetworkPlainArgs();

    /**
     * Query parameters for finding private networks.
     * 
     */
    @Import(name="filters")
    private @Nullable List<GetPrivateNetworkFilter> filters;

    /**
     * @return Query parameters for finding private networks.
     * 
     */
    public Optional<List<GetPrivateNetworkFilter>> filters() {
        return Optional.ofNullable(this.filters);
    }

    private GetPrivateNetworkPlainArgs() {}

    private GetPrivateNetworkPlainArgs(GetPrivateNetworkPlainArgs $) {
        this.filters = $.filters;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPrivateNetworkPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPrivateNetworkPlainArgs $;

        public Builder() {
            $ = new GetPrivateNetworkPlainArgs();
        }

        public Builder(GetPrivateNetworkPlainArgs defaults) {
            $ = new GetPrivateNetworkPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filters Query parameters for finding private networks.
         * 
         * @return builder
         * 
         */
        public Builder filters(@Nullable List<GetPrivateNetworkFilter> filters) {
            $.filters = filters;
            return this;
        }

        /**
         * @param filters Query parameters for finding private networks.
         * 
         * @return builder
         * 
         */
        public Builder filters(GetPrivateNetworkFilter... filters) {
            return filters(List.of(filters));
        }

        public GetPrivateNetworkPlainArgs build() {
            return $;
        }
    }

}
