// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.outputs;

import com.pulumi.core.annotations.CustomType;
import io.dirien.vultr.outputs.GetPrivateNetworkFilter;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetPrivateNetworkResult {
    /**
     * @return The date the private network was added to your Vultr account.
     * 
     */
    private String dateCreated;
    /**
     * @return The private network&#39;s description.
     * 
     */
    private String description;
    private @Nullable List<GetPrivateNetworkFilter> filters;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The ID of the region that the private network is in.
     * 
     */
    private String region;
    /**
     * @return The IPv4 network address. For example: 10.1.1.0.
     * 
     */
    private String v4Subnet;
    /**
     * @return The number of bits for the netmask in CIDR notation. Example: 20
     * 
     */
    private Integer v4SubnetMask;

    private GetPrivateNetworkResult() {}
    /**
     * @return The date the private network was added to your Vultr account.
     * 
     */
    public String dateCreated() {
        return this.dateCreated;
    }
    /**
     * @return The private network&#39;s description.
     * 
     */
    public String description() {
        return this.description;
    }
    public List<GetPrivateNetworkFilter> filters() {
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
     * @return The ID of the region that the private network is in.
     * 
     */
    public String region() {
        return this.region;
    }
    /**
     * @return The IPv4 network address. For example: 10.1.1.0.
     * 
     */
    public String v4Subnet() {
        return this.v4Subnet;
    }
    /**
     * @return The number of bits for the netmask in CIDR notation. Example: 20
     * 
     */
    public Integer v4SubnetMask() {
        return this.v4SubnetMask;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPrivateNetworkResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String dateCreated;
        private String description;
        private @Nullable List<GetPrivateNetworkFilter> filters;
        private String id;
        private String region;
        private String v4Subnet;
        private Integer v4SubnetMask;
        public Builder() {}
        public Builder(GetPrivateNetworkResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dateCreated = defaults.dateCreated;
    	      this.description = defaults.description;
    	      this.filters = defaults.filters;
    	      this.id = defaults.id;
    	      this.region = defaults.region;
    	      this.v4Subnet = defaults.v4Subnet;
    	      this.v4SubnetMask = defaults.v4SubnetMask;
        }

        @CustomType.Setter
        public Builder dateCreated(String dateCreated) {
            this.dateCreated = Objects.requireNonNull(dateCreated);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder filters(@Nullable List<GetPrivateNetworkFilter> filters) {
            this.filters = filters;
            return this;
        }
        public Builder filters(GetPrivateNetworkFilter... filters) {
            return filters(List.of(filters));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            this.region = Objects.requireNonNull(region);
            return this;
        }
        @CustomType.Setter
        public Builder v4Subnet(String v4Subnet) {
            this.v4Subnet = Objects.requireNonNull(v4Subnet);
            return this;
        }
        @CustomType.Setter
        public Builder v4SubnetMask(Integer v4SubnetMask) {
            this.v4SubnetMask = Objects.requireNonNull(v4SubnetMask);
            return this;
        }
        public GetPrivateNetworkResult build() {
            final var o = new GetPrivateNetworkResult();
            o.dateCreated = dateCreated;
            o.description = description;
            o.filters = filters;
            o.id = id;
            o.region = region;
            o.v4Subnet = v4Subnet;
            o.v4SubnetMask = v4SubnetMask;
            return o;
        }
    }
}
