// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.inputs;

import com.pulumi.core.annotations.Import;
import io.dirien.vultr.inputs.GetOsFilter;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetOsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetOsPlainArgs Empty = new GetOsPlainArgs();

    /**
     * Query parameters for finding operating systems.
     * 
     */
    @Import(name="filters")
    private @Nullable List<GetOsFilter> filters;

    /**
     * @return Query parameters for finding operating systems.
     * 
     */
    public Optional<List<GetOsFilter>> filters() {
        return Optional.ofNullable(this.filters);
    }

    private GetOsPlainArgs() {}

    private GetOsPlainArgs(GetOsPlainArgs $) {
        this.filters = $.filters;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetOsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetOsPlainArgs $;

        public Builder() {
            $ = new GetOsPlainArgs();
        }

        public Builder(GetOsPlainArgs defaults) {
            $ = new GetOsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filters Query parameters for finding operating systems.
         * 
         * @return builder
         * 
         */
        public Builder filters(@Nullable List<GetOsFilter> filters) {
            $.filters = filters;
            return this;
        }

        /**
         * @param filters Query parameters for finding operating systems.
         * 
         * @return builder
         * 
         */
        public Builder filters(GetOsFilter... filters) {
            return filters(List.of(filters));
        }

        public GetOsPlainArgs build() {
            return $;
        }
    }

}
