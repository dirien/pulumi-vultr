// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.outputs;

import com.pulumi.core.annotations.CustomType;
import io.dirien.vultr.outputs.GetIsoPublicFilter;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetIsoPublicResult {
    /**
     * @return The description of the ISO file.
     * 
     */
    private String description;
    private @Nullable List<GetIsoPublicFilter> filters;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The MD5Sum of the ISO file.
     * 
     */
    private String md5sum;
    /**
     * @return The ISO file&#39;s name.
     * 
     */
    private String name;

    private GetIsoPublicResult() {}
    /**
     * @return The description of the ISO file.
     * 
     */
    public String description() {
        return this.description;
    }
    public List<GetIsoPublicFilter> filters() {
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
     * @return The MD5Sum of the ISO file.
     * 
     */
    public String md5sum() {
        return this.md5sum;
    }
    /**
     * @return The ISO file&#39;s name.
     * 
     */
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIsoPublicResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private @Nullable List<GetIsoPublicFilter> filters;
        private String id;
        private String md5sum;
        private String name;
        public Builder() {}
        public Builder(GetIsoPublicResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.filters = defaults.filters;
    	      this.id = defaults.id;
    	      this.md5sum = defaults.md5sum;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder filters(@Nullable List<GetIsoPublicFilter> filters) {
            this.filters = filters;
            return this;
        }
        public Builder filters(GetIsoPublicFilter... filters) {
            return filters(List.of(filters));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder md5sum(String md5sum) {
            this.md5sum = Objects.requireNonNull(md5sum);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public GetIsoPublicResult build() {
            final var o = new GetIsoPublicResult();
            o.description = description;
            o.filters = filters;
            o.id = id;
            o.md5sum = md5sum;
            o.name = name;
            return o;
        }
    }
}
