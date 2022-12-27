// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.outputs;

import com.pulumi.core.annotations.CustomType;
import io.dirien.vultr.outputs.GetReverseIpv6Filter;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetReverseIpv6Result {
    private @Nullable List<GetReverseIpv6Filter> filters;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The ID of the instance the IPv6 reverse DNS record was set for.
     * 
     */
    private String instanceId;
    /**
     * @return The IPv6 address in canonical format used in the reverse DNS record.
     * 
     */
    private String ip;
    /**
     * @return The hostname used in the IPv6 reverse DNS record.
     * 
     */
    private String reverse;

    private GetReverseIpv6Result() {}
    public List<GetReverseIpv6Filter> filters() {
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
     * @return The ID of the instance the IPv6 reverse DNS record was set for.
     * 
     */
    public String instanceId() {
        return this.instanceId;
    }
    /**
     * @return The IPv6 address in canonical format used in the reverse DNS record.
     * 
     */
    public String ip() {
        return this.ip;
    }
    /**
     * @return The hostname used in the IPv6 reverse DNS record.
     * 
     */
    public String reverse() {
        return this.reverse;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetReverseIpv6Result defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<GetReverseIpv6Filter> filters;
        private String id;
        private String instanceId;
        private String ip;
        private String reverse;
        public Builder() {}
        public Builder(GetReverseIpv6Result defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.filters = defaults.filters;
    	      this.id = defaults.id;
    	      this.instanceId = defaults.instanceId;
    	      this.ip = defaults.ip;
    	      this.reverse = defaults.reverse;
        }

        @CustomType.Setter
        public Builder filters(@Nullable List<GetReverseIpv6Filter> filters) {
            this.filters = filters;
            return this;
        }
        public Builder filters(GetReverseIpv6Filter... filters) {
            return filters(List.of(filters));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder instanceId(String instanceId) {
            this.instanceId = Objects.requireNonNull(instanceId);
            return this;
        }
        @CustomType.Setter
        public Builder ip(String ip) {
            this.ip = Objects.requireNonNull(ip);
            return this;
        }
        @CustomType.Setter
        public Builder reverse(String reverse) {
            this.reverse = Objects.requireNonNull(reverse);
            return this;
        }
        public GetReverseIpv6Result build() {
            final var o = new GetReverseIpv6Result();
            o.filters = filters;
            o.id = id;
            o.instanceId = instanceId;
            o.ip = ip;
            o.reverse = reverse;
            return o;
        }
    }
}
