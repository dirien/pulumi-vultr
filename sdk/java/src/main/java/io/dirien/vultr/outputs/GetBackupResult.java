// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.outputs;

import com.pulumi.core.annotations.CustomType;
import io.dirien.vultr.outputs.GetBackupFilter;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetBackupResult {
    private List<Map<String,Object>> backups;
    private @Nullable List<GetBackupFilter> filters;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;

    private GetBackupResult() {}
    public List<Map<String,Object>> backups() {
        return this.backups;
    }
    public List<GetBackupFilter> filters() {
        return this.filters == null ? List.of() : this.filters;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBackupResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<Map<String,Object>> backups;
        private @Nullable List<GetBackupFilter> filters;
        private String id;
        public Builder() {}
        public Builder(GetBackupResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.backups = defaults.backups;
    	      this.filters = defaults.filters;
    	      this.id = defaults.id;
        }

        @CustomType.Setter
        public Builder backups(List<Map<String,Object>> backups) {
            this.backups = Objects.requireNonNull(backups);
            return this;
        }
        @CustomType.Setter
        public Builder filters(@Nullable List<GetBackupFilter> filters) {
            this.filters = filters;
            return this;
        }
        public Builder filters(GetBackupFilter... filters) {
            return filters(List.of(filters));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public GetBackupResult build() {
            final var o = new GetBackupResult();
            o.backups = backups;
            o.filters = filters;
            o.id = id;
            return o;
        }
    }
}
