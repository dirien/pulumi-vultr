// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class LoadBalancerFirewallRule {
    /**
     * @return The load balancer ID.
     * 
     */
    private @Nullable String id;
    /**
     * @return The type of ip this rule is - may be either v4 or v6.
     * 
     */
    private String ipType;
    /**
     * @return The assigned port (integer) on the attached instances that the load balancer should check against. Default value is `80`.
     * 
     */
    private Integer port;
    /**
     * @return IP address with subnet that is allowed through the firewall. You may also pass in `cloudflare` which will allow only CloudFlares IP range.
     * 
     */
    private String source;

    private LoadBalancerFirewallRule() {}
    /**
     * @return The load balancer ID.
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }
    /**
     * @return The type of ip this rule is - may be either v4 or v6.
     * 
     */
    public String ipType() {
        return this.ipType;
    }
    /**
     * @return The assigned port (integer) on the attached instances that the load balancer should check against. Default value is `80`.
     * 
     */
    public Integer port() {
        return this.port;
    }
    /**
     * @return IP address with subnet that is allowed through the firewall. You may also pass in `cloudflare` which will allow only CloudFlares IP range.
     * 
     */
    public String source() {
        return this.source;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(LoadBalancerFirewallRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String id;
        private String ipType;
        private Integer port;
        private String source;
        public Builder() {}
        public Builder(LoadBalancerFirewallRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.ipType = defaults.ipType;
    	      this.port = defaults.port;
    	      this.source = defaults.source;
        }

        @CustomType.Setter
        public Builder id(@Nullable String id) {
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder ipType(String ipType) {
            this.ipType = Objects.requireNonNull(ipType);
            return this;
        }
        @CustomType.Setter
        public Builder port(Integer port) {
            this.port = Objects.requireNonNull(port);
            return this;
        }
        @CustomType.Setter
        public Builder source(String source) {
            this.source = Objects.requireNonNull(source);
            return this;
        }
        public LoadBalancerFirewallRule build() {
            final var o = new LoadBalancerFirewallRule();
            o.id = id;
            o.ipType = ipType;
            o.port = port;
            o.source = source;
            return o;
        }
    }
}
