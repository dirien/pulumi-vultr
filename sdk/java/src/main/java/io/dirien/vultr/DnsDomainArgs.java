// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DnsDomainArgs extends com.pulumi.resources.ResourceArgs {

    public static final DnsDomainArgs Empty = new DnsDomainArgs();

    /**
     * The Domain&#39;s DNSSEC status. Valid options are `enabled` or `disabled`. Note `disabled` is default
     * 
     */
    @Import(name="dnsSec")
    private @Nullable Output<String> dnsSec;

    /**
     * @return The Domain&#39;s DNSSEC status. Valid options are `enabled` or `disabled`. Note `disabled` is default
     * 
     */
    public Optional<Output<String>> dnsSec() {
        return Optional.ofNullable(this.dnsSec);
    }

    /**
     * Name of domain.
     * 
     */
    @Import(name="domain", required=true)
    private Output<String> domain;

    /**
     * @return Name of domain.
     * 
     */
    public Output<String> domain() {
        return this.domain;
    }

    /**
     * Instance IP you want associated to domain. If omitted this will create a domain with no records.
     * 
     */
    @Import(name="ip")
    private @Nullable Output<String> ip;

    /**
     * @return Instance IP you want associated to domain. If omitted this will create a domain with no records.
     * 
     */
    public Optional<Output<String>> ip() {
        return Optional.ofNullable(this.ip);
    }

    private DnsDomainArgs() {}

    private DnsDomainArgs(DnsDomainArgs $) {
        this.dnsSec = $.dnsSec;
        this.domain = $.domain;
        this.ip = $.ip;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DnsDomainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DnsDomainArgs $;

        public Builder() {
            $ = new DnsDomainArgs();
        }

        public Builder(DnsDomainArgs defaults) {
            $ = new DnsDomainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dnsSec The Domain&#39;s DNSSEC status. Valid options are `enabled` or `disabled`. Note `disabled` is default
         * 
         * @return builder
         * 
         */
        public Builder dnsSec(@Nullable Output<String> dnsSec) {
            $.dnsSec = dnsSec;
            return this;
        }

        /**
         * @param dnsSec The Domain&#39;s DNSSEC status. Valid options are `enabled` or `disabled`. Note `disabled` is default
         * 
         * @return builder
         * 
         */
        public Builder dnsSec(String dnsSec) {
            return dnsSec(Output.of(dnsSec));
        }

        /**
         * @param domain Name of domain.
         * 
         * @return builder
         * 
         */
        public Builder domain(Output<String> domain) {
            $.domain = domain;
            return this;
        }

        /**
         * @param domain Name of domain.
         * 
         * @return builder
         * 
         */
        public Builder domain(String domain) {
            return domain(Output.of(domain));
        }

        /**
         * @param ip Instance IP you want associated to domain. If omitted this will create a domain with no records.
         * 
         * @return builder
         * 
         */
        public Builder ip(@Nullable Output<String> ip) {
            $.ip = ip;
            return this;
        }

        /**
         * @param ip Instance IP you want associated to domain. If omitted this will create a domain with no records.
         * 
         * @return builder
         * 
         */
        public Builder ip(String ip) {
            return ip(Output.of(ip));
        }

        public DnsDomainArgs build() {
            $.domain = Objects.requireNonNull($.domain, "expected parameter 'domain' to be non-null");
            return $;
        }
    }

}
