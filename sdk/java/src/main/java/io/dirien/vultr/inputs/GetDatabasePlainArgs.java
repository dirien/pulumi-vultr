// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.inputs;

import com.pulumi.core.annotations.Import;
import io.dirien.vultr.inputs.GetDatabaseFilter;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDatabasePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDatabasePlainArgs Empty = new GetDatabasePlainArgs();

    @Import(name="filters")
    private @Nullable List<GetDatabaseFilter> filters;

    public Optional<List<GetDatabaseFilter>> filters() {
        return Optional.ofNullable(this.filters);
    }

    private GetDatabasePlainArgs() {}

    private GetDatabasePlainArgs(GetDatabasePlainArgs $) {
        this.filters = $.filters;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDatabasePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDatabasePlainArgs $;

        public Builder() {
            $ = new GetDatabasePlainArgs();
        }

        public Builder(GetDatabasePlainArgs defaults) {
            $ = new GetDatabasePlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder filters(@Nullable List<GetDatabaseFilter> filters) {
            $.filters = filters;
            return this;
        }

        public Builder filters(GetDatabaseFilter... filters) {
            return filters(List.of(filters));
        }

        public GetDatabasePlainArgs build() {
            return $;
        }
    }

}
