// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import io.dirien.vultr.inputs.GetObjectStorageClusterFilterArgs;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetObjectStorageClusterArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetObjectStorageClusterArgs Empty = new GetObjectStorageClusterArgs();

    /**
     * Query parameters for finding operating systems.
     * 
     */
    @Import(name="filters")
    private @Nullable Output<List<GetObjectStorageClusterFilterArgs>> filters;

    /**
     * @return Query parameters for finding operating systems.
     * 
     */
    public Optional<Output<List<GetObjectStorageClusterFilterArgs>>> filters() {
        return Optional.ofNullable(this.filters);
    }

    private GetObjectStorageClusterArgs() {}

    private GetObjectStorageClusterArgs(GetObjectStorageClusterArgs $) {
        this.filters = $.filters;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetObjectStorageClusterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetObjectStorageClusterArgs $;

        public Builder() {
            $ = new GetObjectStorageClusterArgs();
        }

        public Builder(GetObjectStorageClusterArgs defaults) {
            $ = new GetObjectStorageClusterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filters Query parameters for finding operating systems.
         * 
         * @return builder
         * 
         */
        public Builder filters(@Nullable Output<List<GetObjectStorageClusterFilterArgs>> filters) {
            $.filters = filters;
            return this;
        }

        /**
         * @param filters Query parameters for finding operating systems.
         * 
         * @return builder
         * 
         */
        public Builder filters(List<GetObjectStorageClusterFilterArgs> filters) {
            return filters(Output.of(filters));
        }

        /**
         * @param filters Query parameters for finding operating systems.
         * 
         * @return builder
         * 
         */
        public Builder filters(GetObjectStorageClusterFilterArgs... filters) {
            return filters(List.of(filters));
        }

        public GetObjectStorageClusterArgs build() {
            return $;
        }
    }

}