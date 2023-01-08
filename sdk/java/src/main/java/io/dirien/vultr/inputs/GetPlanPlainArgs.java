// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.inputs;

import com.pulumi.core.annotations.Import;
import io.dirien.vultr.inputs.GetPlanFilter;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPlanPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPlanPlainArgs Empty = new GetPlanPlainArgs();

    /**
     * Query parameters for finding plans.
     * 
     */
    @Import(name="filters")
    private @Nullable List<GetPlanFilter> filters;

    /**
     * @return Query parameters for finding plans.
     * 
     */
    public Optional<List<GetPlanFilter>> filters() {
        return Optional.ofNullable(this.filters);
    }

    private GetPlanPlainArgs() {}

    private GetPlanPlainArgs(GetPlanPlainArgs $) {
        this.filters = $.filters;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPlanPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPlanPlainArgs $;

        public Builder() {
            $ = new GetPlanPlainArgs();
        }

        public Builder(GetPlanPlainArgs defaults) {
            $ = new GetPlanPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filters Query parameters for finding plans.
         * 
         * @return builder
         * 
         */
        public Builder filters(@Nullable List<GetPlanFilter> filters) {
            $.filters = filters;
            return this;
        }

        /**
         * @param filters Query parameters for finding plans.
         * 
         * @return builder
         * 
         */
        public Builder filters(GetPlanFilter... filters) {
            return filters(List.of(filters));
        }

        public GetPlanPlainArgs build() {
            return $;
        }
    }

}