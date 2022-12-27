// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.outputs;

import com.pulumi.core.annotations.CustomType;
import io.dirien.vultr.outputs.GetOsFilter;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetOsResult {
    /**
     * @return The architecture of the operating system.
     * 
     */
    private String arch;
    /**
     * @return The family of the operating system.
     * 
     */
    private String family;
    private @Nullable List<GetOsFilter> filters;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The name of the operating system.
     * 
     */
    private String name;

    private GetOsResult() {}
    /**
     * @return The architecture of the operating system.
     * 
     */
    public String arch() {
        return this.arch;
    }
    /**
     * @return The family of the operating system.
     * 
     */
    public String family() {
        return this.family;
    }
    public List<GetOsFilter> filters() {
        return this.filters == null ? List.of() : this.filters;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The name of the operating system.
     * 
     */
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String arch;
        private String family;
        private @Nullable List<GetOsFilter> filters;
        private String id;
        private String name;
        public Builder() {}
        public Builder(GetOsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.arch = defaults.arch;
    	      this.family = defaults.family;
    	      this.filters = defaults.filters;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder arch(String arch) {
            this.arch = Objects.requireNonNull(arch);
            return this;
        }
        @CustomType.Setter
        public Builder family(String family) {
            this.family = Objects.requireNonNull(family);
            return this;
        }
        @CustomType.Setter
        public Builder filters(@Nullable List<GetOsFilter> filters) {
            this.filters = filters;
            return this;
        }
        public Builder filters(GetOsFilter... filters) {
            return filters(List.of(filters));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public GetOsResult build() {
            final var o = new GetOsResult();
            o.arch = arch;
            o.family = family;
            o.filters = filters;
            o.id = id;
            o.name = name;
            return o;
        }
    }
}
