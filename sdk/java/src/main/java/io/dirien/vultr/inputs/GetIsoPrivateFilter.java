// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class GetIsoPrivateFilter extends com.pulumi.resources.InvokeArgs {

    public static final GetIsoPrivateFilter Empty = new GetIsoPrivateFilter();

    /**
     * Attribute name to filter with.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return Attribute name to filter with.
     * 
     */
    public String name() {
        return this.name;
    }

    /**
     * One or more values filter with.
     * 
     */
    @Import(name="values", required=true)
    private List<String> values;

    /**
     * @return One or more values filter with.
     * 
     */
    public List<String> values() {
        return this.values;
    }

    private GetIsoPrivateFilter() {}

    private GetIsoPrivateFilter(GetIsoPrivateFilter $) {
        this.name = $.name;
        this.values = $.values;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIsoPrivateFilter defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIsoPrivateFilter $;

        public Builder() {
            $ = new GetIsoPrivateFilter();
        }

        public Builder(GetIsoPrivateFilter defaults) {
            $ = new GetIsoPrivateFilter(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Attribute name to filter with.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        /**
         * @param values One or more values filter with.
         * 
         * @return builder
         * 
         */
        public Builder values(List<String> values) {
            $.values = values;
            return this;
        }

        /**
         * @param values One or more values filter with.
         * 
         * @return builder
         * 
         */
        public Builder values(String... values) {
            return values(List.of(values));
        }

        public GetIsoPrivateFilter build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            $.values = Objects.requireNonNull($.values, "expected parameter 'values' to be non-null");
            return $;
        }
    }

}
