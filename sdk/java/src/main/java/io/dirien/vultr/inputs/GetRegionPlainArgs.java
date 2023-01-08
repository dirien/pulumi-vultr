// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.inputs;

import com.pulumi.core.annotations.Import;
import io.dirien.vultr.inputs.GetRegionFilter;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRegionPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRegionPlainArgs Empty = new GetRegionPlainArgs();

    /**
     * Query parameters for finding regions.
     * 
     */
    @Import(name="filters")
    private @Nullable List<GetRegionFilter> filters;

    /**
     * @return Query parameters for finding regions.
     * 
     */
    public Optional<List<GetRegionFilter>> filters() {
        return Optional.ofNullable(this.filters);
    }

    private GetRegionPlainArgs() {}

    private GetRegionPlainArgs(GetRegionPlainArgs $) {
        this.filters = $.filters;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRegionPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRegionPlainArgs $;

        public Builder() {
            $ = new GetRegionPlainArgs();
        }

        public Builder(GetRegionPlainArgs defaults) {
            $ = new GetRegionPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filters Query parameters for finding regions.
         * 
         * @return builder
         * 
         */
        public Builder filters(@Nullable List<GetRegionFilter> filters) {
            $.filters = filters;
            return this;
        }

        /**
         * @param filters Query parameters for finding regions.
         * 
         * @return builder
         * 
         */
        public Builder filters(GetRegionFilter... filters) {
            return filters(List.of(filters));
        }

        public GetRegionPlainArgs build() {
            return $;
        }
    }

}