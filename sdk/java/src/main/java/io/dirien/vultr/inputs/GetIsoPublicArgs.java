// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import io.dirien.vultr.inputs.GetIsoPublicFilterArgs;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetIsoPublicArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetIsoPublicArgs Empty = new GetIsoPublicArgs();

    /**
     * Query parameters for finding ISO files.
     * 
     */
    @Import(name="filters")
    private @Nullable Output<List<GetIsoPublicFilterArgs>> filters;

    /**
     * @return Query parameters for finding ISO files.
     * 
     */
    public Optional<Output<List<GetIsoPublicFilterArgs>>> filters() {
        return Optional.ofNullable(this.filters);
    }

    private GetIsoPublicArgs() {}

    private GetIsoPublicArgs(GetIsoPublicArgs $) {
        this.filters = $.filters;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIsoPublicArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIsoPublicArgs $;

        public Builder() {
            $ = new GetIsoPublicArgs();
        }

        public Builder(GetIsoPublicArgs defaults) {
            $ = new GetIsoPublicArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filters Query parameters for finding ISO files.
         * 
         * @return builder
         * 
         */
        public Builder filters(@Nullable Output<List<GetIsoPublicFilterArgs>> filters) {
            $.filters = filters;
            return this;
        }

        /**
         * @param filters Query parameters for finding ISO files.
         * 
         * @return builder
         * 
         */
        public Builder filters(List<GetIsoPublicFilterArgs> filters) {
            return filters(Output.of(filters));
        }

        /**
         * @param filters Query parameters for finding ISO files.
         * 
         * @return builder
         * 
         */
        public Builder filters(GetIsoPublicFilterArgs... filters) {
            return filters(List.of(filters));
        }

        public GetIsoPublicArgs build() {
            return $;
        }
    }

}
