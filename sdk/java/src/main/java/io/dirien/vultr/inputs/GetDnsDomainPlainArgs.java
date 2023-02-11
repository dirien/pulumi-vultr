// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetDnsDomainPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDnsDomainPlainArgs Empty = new GetDnsDomainPlainArgs();

    /**
     * The name you&#39;re searching for.
     * 
     */
    @Import(name="domain", required=true)
    private String domain;

    /**
     * @return The name you&#39;re searching for.
     * 
     */
    public String domain() {
        return this.domain;
    }

    private GetDnsDomainPlainArgs() {}

    private GetDnsDomainPlainArgs(GetDnsDomainPlainArgs $) {
        this.domain = $.domain;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDnsDomainPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDnsDomainPlainArgs $;

        public Builder() {
            $ = new GetDnsDomainPlainArgs();
        }

        public Builder(GetDnsDomainPlainArgs defaults) {
            $ = new GetDnsDomainPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param domain The name you&#39;re searching for.
         * 
         * @return builder
         * 
         */
        public Builder domain(String domain) {
            $.domain = domain;
            return this;
        }

        public GetDnsDomainPlainArgs build() {
            $.domain = Objects.requireNonNull($.domain, "expected parameter 'domain' to be non-null");
            return $;
        }
    }

}
