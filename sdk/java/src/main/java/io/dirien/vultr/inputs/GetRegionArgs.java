// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import io.dirien.vultr.inputs.GetRegionFilterArgs;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRegionArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRegionArgs Empty = new GetRegionArgs();

    /**
     * Query parameters for finding regions.
     * 
     */
    @Import(name="filters")
    private @Nullable Output<List<GetRegionFilterArgs>> filters;

    /**
     * @return Query parameters for finding regions.
     * 
     */
    public Optional<Output<List<GetRegionFilterArgs>>> filters() {
        return Optional.ofNullable(this.filters);
    }

    private GetRegionArgs() {}

    private GetRegionArgs(GetRegionArgs $) {
        this.filters = $.filters;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRegionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRegionArgs $;

        public Builder() {
            $ = new GetRegionArgs();
        }

        public Builder(GetRegionArgs defaults) {
            $ = new GetRegionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filters Query parameters for finding regions.
         * 
         * @return builder
         * 
         */
        public Builder filters(@Nullable Output<List<GetRegionFilterArgs>> filters) {
            $.filters = filters;
            return this;
        }

        /**
         * @param filters Query parameters for finding regions.
         * 
         * @return builder
         * 
         */
        public Builder filters(List<GetRegionFilterArgs> filters) {
            return filters(Output.of(filters));
        }

        /**
         * @param filters Query parameters for finding regions.
         * 
         * @return builder
         * 
         */
        public Builder filters(GetRegionFilterArgs... filters) {
            return filters(List.of(filters));
        }

        public GetRegionArgs build() {
            return $;
        }
    }

}
