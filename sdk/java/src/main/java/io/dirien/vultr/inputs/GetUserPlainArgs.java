// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.inputs;

import com.pulumi.core.annotations.Import;
import io.dirien.vultr.inputs.GetUserFilter;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetUserPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetUserPlainArgs Empty = new GetUserPlainArgs();

    /**
     * Query parameters for finding users.
     * 
     */
    @Import(name="filters")
    private @Nullable List<GetUserFilter> filters;

    /**
     * @return Query parameters for finding users.
     * 
     */
    public Optional<List<GetUserFilter>> filters() {
        return Optional.ofNullable(this.filters);
    }

    private GetUserPlainArgs() {}

    private GetUserPlainArgs(GetUserPlainArgs $) {
        this.filters = $.filters;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetUserPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetUserPlainArgs $;

        public Builder() {
            $ = new GetUserPlainArgs();
        }

        public Builder(GetUserPlainArgs defaults) {
            $ = new GetUserPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filters Query parameters for finding users.
         * 
         * @return builder
         * 
         */
        public Builder filters(@Nullable List<GetUserFilter> filters) {
            $.filters = filters;
            return this;
        }

        /**
         * @param filters Query parameters for finding users.
         * 
         * @return builder
         * 
         */
        public Builder filters(GetUserFilter... filters) {
            return filters(List.of(filters));
        }

        public GetUserPlainArgs build() {
            return $;
        }
    }

}
