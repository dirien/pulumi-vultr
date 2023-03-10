// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import io.dirien.vultr.inputs.GetOsFilterArgs;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetOsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetOsArgs Empty = new GetOsArgs();

    /**
     * Query parameters for finding operating systems.
     * 
     */
    @Import(name="filters")
    private @Nullable Output<List<GetOsFilterArgs>> filters;

    /**
     * @return Query parameters for finding operating systems.
     * 
     */
    public Optional<Output<List<GetOsFilterArgs>>> filters() {
        return Optional.ofNullable(this.filters);
    }

    private GetOsArgs() {}

    private GetOsArgs(GetOsArgs $) {
        this.filters = $.filters;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetOsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetOsArgs $;

        public Builder() {
            $ = new GetOsArgs();
        }

        public Builder(GetOsArgs defaults) {
            $ = new GetOsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filters Query parameters for finding operating systems.
         * 
         * @return builder
         * 
         */
        public Builder filters(@Nullable Output<List<GetOsFilterArgs>> filters) {
            $.filters = filters;
            return this;
        }

        /**
         * @param filters Query parameters for finding operating systems.
         * 
         * @return builder
         * 
         */
        public Builder filters(List<GetOsFilterArgs> filters) {
            return filters(Output.of(filters));
        }

        /**
         * @param filters Query parameters for finding operating systems.
         * 
         * @return builder
         * 
         */
        public Builder filters(GetOsFilterArgs... filters) {
            return filters(List.of(filters));
        }

        public GetOsArgs build() {
            return $;
        }
    }

}
