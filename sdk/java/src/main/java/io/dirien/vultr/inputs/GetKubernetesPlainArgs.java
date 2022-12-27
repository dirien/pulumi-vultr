// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.inputs;

import com.pulumi.core.annotations.Import;
import io.dirien.vultr.inputs.GetKubernetesFilter;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetKubernetesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetKubernetesPlainArgs Empty = new GetKubernetesPlainArgs();

    /**
     * Query parameters for finding VKE.
     * 
     */
    @Import(name="filters")
    private @Nullable List<GetKubernetesFilter> filters;

    /**
     * @return Query parameters for finding VKE.
     * 
     */
    public Optional<List<GetKubernetesFilter>> filters() {
        return Optional.ofNullable(this.filters);
    }

    private GetKubernetesPlainArgs() {}

    private GetKubernetesPlainArgs(GetKubernetesPlainArgs $) {
        this.filters = $.filters;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetKubernetesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetKubernetesPlainArgs $;

        public Builder() {
            $ = new GetKubernetesPlainArgs();
        }

        public Builder(GetKubernetesPlainArgs defaults) {
            $ = new GetKubernetesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filters Query parameters for finding VKE.
         * 
         * @return builder
         * 
         */
        public Builder filters(@Nullable List<GetKubernetesFilter> filters) {
            $.filters = filters;
            return this;
        }

        /**
         * @param filters Query parameters for finding VKE.
         * 
         * @return builder
         * 
         */
        public Builder filters(GetKubernetesFilter... filters) {
            return filters(List.of(filters));
        }

        public GetKubernetesPlainArgs build() {
            return $;
        }
    }

}
