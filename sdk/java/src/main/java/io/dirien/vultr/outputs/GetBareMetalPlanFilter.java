// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetBareMetalPlanFilter {
    /**
     * @return Attribute name to filter with.
     * 
     */
    private String name;
    /**
     * @return One or more values filter with.
     * 
     */
    private List<String> values;

    private GetBareMetalPlanFilter() {}
    /**
     * @return Attribute name to filter with.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return One or more values filter with.
     * 
     */
    public List<String> values() {
        return this.values;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBareMetalPlanFilter defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String name;
        private List<String> values;
        public Builder() {}
        public Builder(GetBareMetalPlanFilter defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.values = defaults.values;
        }

        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder values(List<String> values) {
            this.values = Objects.requireNonNull(values);
            return this;
        }
        public Builder values(String... values) {
            return values(List.of(values));
        }
        public GetBareMetalPlanFilter build() {
            final var o = new GetBareMetalPlanFilter();
            o.name = name;
            o.values = values;
            return o;
        }
    }
}
