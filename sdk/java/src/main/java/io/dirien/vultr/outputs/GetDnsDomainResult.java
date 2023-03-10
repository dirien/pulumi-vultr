// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetDnsDomainResult {
    /**
     * @return The date the DNS domain was added to your Vultr account.
     * 
     */
    private String dateCreated;
    /**
     * @return The Domain&#39;s DNSSEC status
     * 
     */
    private String dnsSec;
    /**
     * @return Name of domain.
     * 
     */
    private String domain;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;

    private GetDnsDomainResult() {}
    /**
     * @return The date the DNS domain was added to your Vultr account.
     * 
     */
    public String dateCreated() {
        return this.dateCreated;
    }
    /**
     * @return The Domain&#39;s DNSSEC status
     * 
     */
    public String dnsSec() {
        return this.dnsSec;
    }
    /**
     * @return Name of domain.
     * 
     */
    public String domain() {
        return this.domain;
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

    public static Builder builder(GetDnsDomainResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String dateCreated;
        private String dnsSec;
        private String domain;
        private String id;
        public Builder() {}
        public Builder(GetDnsDomainResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dateCreated = defaults.dateCreated;
    	      this.dnsSec = defaults.dnsSec;
    	      this.domain = defaults.domain;
    	      this.id = defaults.id;
        }

        @CustomType.Setter
        public Builder dateCreated(String dateCreated) {
            this.dateCreated = Objects.requireNonNull(dateCreated);
            return this;
        }
        @CustomType.Setter
        public Builder dnsSec(String dnsSec) {
            this.dnsSec = Objects.requireNonNull(dnsSec);
            return this;
        }
        @CustomType.Setter
        public Builder domain(String domain) {
            this.domain = Objects.requireNonNull(domain);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public GetDnsDomainResult build() {
            final var o = new GetDnsDomainResult();
            o.dateCreated = dateCreated;
            o.dnsSec = dnsSec;
            o.domain = domain;
            o.id = id;
            return o;
        }
    }
}
