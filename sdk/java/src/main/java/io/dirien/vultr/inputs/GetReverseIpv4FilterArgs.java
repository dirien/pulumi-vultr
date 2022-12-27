// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class GetReverseIpv4FilterArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetReverseIpv4FilterArgs Empty = new GetReverseIpv4FilterArgs();

    /**
     * Attribute name to filter with.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Attribute name to filter with.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * One or more values to filter with.
     * 
     */
    @Import(name="values", required=true)
    private Output<List<String>> values;

    /**
     * @return One or more values to filter with.
     * 
     */
    public Output<List<String>> values() {
        return this.values;
    }

    private GetReverseIpv4FilterArgs() {}

    private GetReverseIpv4FilterArgs(GetReverseIpv4FilterArgs $) {
        this.name = $.name;
        this.values = $.values;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetReverseIpv4FilterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetReverseIpv4FilterArgs $;

        public Builder() {
            $ = new GetReverseIpv4FilterArgs();
        }

        public Builder(GetReverseIpv4FilterArgs defaults) {
            $ = new GetReverseIpv4FilterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Attribute name to filter with.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Attribute name to filter with.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param values One or more values to filter with.
         * 
         * @return builder
         * 
         */
        public Builder values(Output<List<String>> values) {
            $.values = values;
            return this;
        }

        /**
         * @param values One or more values to filter with.
         * 
         * @return builder
         * 
         */
        public Builder values(List<String> values) {
            return values(Output.of(values));
        }

        /**
         * @param values One or more values to filter with.
         * 
         * @return builder
         * 
         */
        public Builder values(String... values) {
            return values(List.of(values));
        }

        public GetReverseIpv4FilterArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            $.values = Objects.requireNonNull($.values, "expected parameter 'values' to be non-null");
            return $;
        }
    }

}
