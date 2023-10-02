// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.outputs;

import com.pulumi.core.annotations.CustomType;
import io.dirien.vultr.outputs.GetBareMetalServerFilter;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetBareMetalServerResult {
    /**
     * @return The server&#39;s application ID.
     * 
     */
    private Integer appId;
    /**
     * @return The number of CPUs available on the server.
     * 
     */
    private Integer cpuCount;
    /**
     * @return The date the server was added to your Vultr account.
     * 
     */
    private String dateCreated;
    /**
     * @return The description of the disk(s) on the server.
     * 
     */
    private String disk;
    private List<String> features;
    private @Nullable List<GetBareMetalServerFilter> filters;
    /**
     * @return The server&#39;s IPv4 gateway.
     * 
     */
    private String gatewayV4;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The Marketplace ID for this application.
     * 
     */
    private String imageId;
    /**
     * @return The server&#39;s label.
     * 
     */
    private String label;
    private Integer macAddress;
    /**
     * @return The server&#39;s main IP address.
     * 
     */
    private String mainIp;
    /**
     * @return The server&#39;s IPv4 netmask.
     * 
     */
    private String netmaskV4;
    /**
     * @return The operating system of the server.
     * 
     */
    private String os;
    /**
     * @return The server&#39;s operating system ID.
     * 
     */
    private Integer osId;
    /**
     * @return The server&#39;s plan ID.
     * 
     */
    private String plan;
    /**
     * @return The amount of memory available on the server in MB.
     * 
     */
    private String ram;
    /**
     * @return The region ID of the server.
     * 
     */
    private String region;
    /**
     * @return The status of the server&#39;s subscription.
     * 
     */
    private String status;
    /**
     * @return A list of tags applied to the server.
     * 
     */
    private List<String> tags;
    private String v6MainIp;
    private String v6Network;
    private Integer v6NetworkSize;
    /**
     * @return A list of VPC 2.0 IDs attached to the server.
     * 
     */
    private List<String> vpc2Ids;

    private GetBareMetalServerResult() {}
    /**
     * @return The server&#39;s application ID.
     * 
     */
    public Integer appId() {
        return this.appId;
    }
    /**
     * @return The number of CPUs available on the server.
     * 
     */
    public Integer cpuCount() {
        return this.cpuCount;
    }
    /**
     * @return The date the server was added to your Vultr account.
     * 
     */
    public String dateCreated() {
        return this.dateCreated;
    }
    /**
     * @return The description of the disk(s) on the server.
     * 
     */
    public String disk() {
        return this.disk;
    }
    public List<String> features() {
        return this.features;
    }
    public List<GetBareMetalServerFilter> filters() {
        return this.filters == null ? List.of() : this.filters;
    }
    /**
     * @return The server&#39;s IPv4 gateway.
     * 
     */
    public String gatewayV4() {
        return this.gatewayV4;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The Marketplace ID for this application.
     * 
     */
    public String imageId() {
        return this.imageId;
    }
    /**
     * @return The server&#39;s label.
     * 
     */
    public String label() {
        return this.label;
    }
    public Integer macAddress() {
        return this.macAddress;
    }
    /**
     * @return The server&#39;s main IP address.
     * 
     */
    public String mainIp() {
        return this.mainIp;
    }
    /**
     * @return The server&#39;s IPv4 netmask.
     * 
     */
    public String netmaskV4() {
        return this.netmaskV4;
    }
    /**
     * @return The operating system of the server.
     * 
     */
    public String os() {
        return this.os;
    }
    /**
     * @return The server&#39;s operating system ID.
     * 
     */
    public Integer osId() {
        return this.osId;
    }
    /**
     * @return The server&#39;s plan ID.
     * 
     */
    public String plan() {
        return this.plan;
    }
    /**
     * @return The amount of memory available on the server in MB.
     * 
     */
    public String ram() {
        return this.ram;
    }
    /**
     * @return The region ID of the server.
     * 
     */
    public String region() {
        return this.region;
    }
    /**
     * @return The status of the server&#39;s subscription.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return A list of tags applied to the server.
     * 
     */
    public List<String> tags() {
        return this.tags;
    }
    public String v6MainIp() {
        return this.v6MainIp;
    }
    public String v6Network() {
        return this.v6Network;
    }
    public Integer v6NetworkSize() {
        return this.v6NetworkSize;
    }
    /**
     * @return A list of VPC 2.0 IDs attached to the server.
     * 
     */
    public List<String> vpc2Ids() {
        return this.vpc2Ids;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBareMetalServerResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer appId;
        private Integer cpuCount;
        private String dateCreated;
        private String disk;
        private List<String> features;
        private @Nullable List<GetBareMetalServerFilter> filters;
        private String gatewayV4;
        private String id;
        private String imageId;
        private String label;
        private Integer macAddress;
        private String mainIp;
        private String netmaskV4;
        private String os;
        private Integer osId;
        private String plan;
        private String ram;
        private String region;
        private String status;
        private List<String> tags;
        private String v6MainIp;
        private String v6Network;
        private Integer v6NetworkSize;
        private List<String> vpc2Ids;
        public Builder() {}
        public Builder(GetBareMetalServerResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.appId = defaults.appId;
    	      this.cpuCount = defaults.cpuCount;
    	      this.dateCreated = defaults.dateCreated;
    	      this.disk = defaults.disk;
    	      this.features = defaults.features;
    	      this.filters = defaults.filters;
    	      this.gatewayV4 = defaults.gatewayV4;
    	      this.id = defaults.id;
    	      this.imageId = defaults.imageId;
    	      this.label = defaults.label;
    	      this.macAddress = defaults.macAddress;
    	      this.mainIp = defaults.mainIp;
    	      this.netmaskV4 = defaults.netmaskV4;
    	      this.os = defaults.os;
    	      this.osId = defaults.osId;
    	      this.plan = defaults.plan;
    	      this.ram = defaults.ram;
    	      this.region = defaults.region;
    	      this.status = defaults.status;
    	      this.tags = defaults.tags;
    	      this.v6MainIp = defaults.v6MainIp;
    	      this.v6Network = defaults.v6Network;
    	      this.v6NetworkSize = defaults.v6NetworkSize;
    	      this.vpc2Ids = defaults.vpc2Ids;
        }

        @CustomType.Setter
        public Builder appId(Integer appId) {
            this.appId = Objects.requireNonNull(appId);
            return this;
        }
        @CustomType.Setter
        public Builder cpuCount(Integer cpuCount) {
            this.cpuCount = Objects.requireNonNull(cpuCount);
            return this;
        }
        @CustomType.Setter
        public Builder dateCreated(String dateCreated) {
            this.dateCreated = Objects.requireNonNull(dateCreated);
            return this;
        }
        @CustomType.Setter
        public Builder disk(String disk) {
            this.disk = Objects.requireNonNull(disk);
            return this;
        }
        @CustomType.Setter
        public Builder features(List<String> features) {
            this.features = Objects.requireNonNull(features);
            return this;
        }
        public Builder features(String... features) {
            return features(List.of(features));
        }
        @CustomType.Setter
        public Builder filters(@Nullable List<GetBareMetalServerFilter> filters) {
            this.filters = filters;
            return this;
        }
        public Builder filters(GetBareMetalServerFilter... filters) {
            return filters(List.of(filters));
        }
        @CustomType.Setter
        public Builder gatewayV4(String gatewayV4) {
            this.gatewayV4 = Objects.requireNonNull(gatewayV4);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder imageId(String imageId) {
            this.imageId = Objects.requireNonNull(imageId);
            return this;
        }
        @CustomType.Setter
        public Builder label(String label) {
            this.label = Objects.requireNonNull(label);
            return this;
        }
        @CustomType.Setter
        public Builder macAddress(Integer macAddress) {
            this.macAddress = Objects.requireNonNull(macAddress);
            return this;
        }
        @CustomType.Setter
        public Builder mainIp(String mainIp) {
            this.mainIp = Objects.requireNonNull(mainIp);
            return this;
        }
        @CustomType.Setter
        public Builder netmaskV4(String netmaskV4) {
            this.netmaskV4 = Objects.requireNonNull(netmaskV4);
            return this;
        }
        @CustomType.Setter
        public Builder os(String os) {
            this.os = Objects.requireNonNull(os);
            return this;
        }
        @CustomType.Setter
        public Builder osId(Integer osId) {
            this.osId = Objects.requireNonNull(osId);
            return this;
        }
        @CustomType.Setter
        public Builder plan(String plan) {
            this.plan = Objects.requireNonNull(plan);
            return this;
        }
        @CustomType.Setter
        public Builder ram(String ram) {
            this.ram = Objects.requireNonNull(ram);
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            this.region = Objects.requireNonNull(region);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder tags(List<String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }
        @CustomType.Setter
        public Builder v6MainIp(String v6MainIp) {
            this.v6MainIp = Objects.requireNonNull(v6MainIp);
            return this;
        }
        @CustomType.Setter
        public Builder v6Network(String v6Network) {
            this.v6Network = Objects.requireNonNull(v6Network);
            return this;
        }
        @CustomType.Setter
        public Builder v6NetworkSize(Integer v6NetworkSize) {
            this.v6NetworkSize = Objects.requireNonNull(v6NetworkSize);
            return this;
        }
        @CustomType.Setter
        public Builder vpc2Ids(List<String> vpc2Ids) {
            this.vpc2Ids = Objects.requireNonNull(vpc2Ids);
            return this;
        }
        public Builder vpc2Ids(String... vpc2Ids) {
            return vpc2Ids(List.of(vpc2Ids));
        }
        public GetBareMetalServerResult build() {
            final var o = new GetBareMetalServerResult();
            o.appId = appId;
            o.cpuCount = cpuCount;
            o.dateCreated = dateCreated;
            o.disk = disk;
            o.features = features;
            o.filters = filters;
            o.gatewayV4 = gatewayV4;
            o.id = id;
            o.imageId = imageId;
            o.label = label;
            o.macAddress = macAddress;
            o.mainIp = mainIp;
            o.netmaskV4 = netmaskV4;
            o.os = os;
            o.osId = osId;
            o.plan = plan;
            o.ram = ram;
            o.region = region;
            o.status = status;
            o.tags = tags;
            o.v6MainIp = v6MainIp;
            o.v6Network = v6Network;
            o.v6NetworkSize = v6NetworkSize;
            o.vpc2Ids = vpc2Ids;
            return o;
        }
    }
}
