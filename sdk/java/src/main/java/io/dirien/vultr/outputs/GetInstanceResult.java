// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.outputs;

import com.pulumi.core.annotations.CustomType;
import io.dirien.vultr.outputs.GetInstanceFilter;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetInstanceResult {
    /**
     * @return The server&#39;s allowed bandwidth usage in GB.
     * 
     */
    private Integer allowedBandwidth;
    /**
     * @return The server&#39;s application ID.
     * 
     */
    private Integer appId;
    private String backups;
    /**
     * @return The current configuration for backups
     * 
     */
    private Map<String,Object> backupsSchedule;
    /**
     * @return The date the server was added to your Vultr account.
     * 
     */
    private String dateCreated;
    /**
     * @return The description of the disk(s) on the server.
     * 
     */
    private Integer disk;
    /**
     * @return Array of which features are enabled.
     * 
     */
    private List<String> features;
    private @Nullable List<GetInstanceFilter> filters;
    /**
     * @return The ID of the firewall group applied to this server.
     * 
     */
    private String firewallGroupId;
    /**
     * @return The server&#39;s IPv4 gateway.
     * 
     */
    private String gatewayV4;
    /**
     * @return The hostname assigned to the server.
     * 
     */
    private String hostname;
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
     * @return The server&#39;s internal IP address.
     * 
     */
    private String internalIp;
    /**
     * @return The server&#39;s current KVM URL. This URL will change periodically. It is not advised to cache this value.
     * 
     */
    private String kvm;
    /**
     * @return The server&#39;s label.
     * 
     */
    private String label;
    private String location;
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
     * @return The operating system of the instance.
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
     * @return Whether the server is powered on or not.
     * 
     */
    private String powerStatus;
    private List<String> privateNetworkIds;
    /**
     * @return The amount of memory available on the instance in MB.
     * 
     */
    private Integer ram;
    /**
     * @return The region ID of the server.
     * 
     */
    private String region;
    /**
     * @return A more detailed server status (none, locked, installingbooting, isomounting, ok).
     * 
     */
    private String serverStatus;
    /**
     * @return The status of the server&#39;s subscription.
     * 
     */
    private String status;
    /**
     * @return A list of tags applied to the instance.
     * 
     */
    private List<String> tags;
    /**
     * @return The main IPv6 network address.
     * 
     */
    private String v6MainIp;
    /**
     * @return The IPv6 subnet.
     * 
     */
    private String v6Network;
    /**
     * @return The IPv6 network size in bits.
     * 
     */
    private Integer v6NetworkSize;
    /**
     * @return The number of virtual CPUs available on the server.
     * 
     */
    private Integer vcpuCount;
    /**
     * @return A list of VPC 2.0 IDs attached to the server.
     * 
     */
    private List<String> vpc2Ids;
    private List<String> vpcIds;

    private GetInstanceResult() {}
    /**
     * @return The server&#39;s allowed bandwidth usage in GB.
     * 
     */
    public Integer allowedBandwidth() {
        return this.allowedBandwidth;
    }
    /**
     * @return The server&#39;s application ID.
     * 
     */
    public Integer appId() {
        return this.appId;
    }
    public String backups() {
        return this.backups;
    }
    /**
     * @return The current configuration for backups
     * 
     */
    public Map<String,Object> backupsSchedule() {
        return this.backupsSchedule;
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
    public Integer disk() {
        return this.disk;
    }
    /**
     * @return Array of which features are enabled.
     * 
     */
    public List<String> features() {
        return this.features;
    }
    public List<GetInstanceFilter> filters() {
        return this.filters == null ? List.of() : this.filters;
    }
    /**
     * @return The ID of the firewall group applied to this server.
     * 
     */
    public String firewallGroupId() {
        return this.firewallGroupId;
    }
    /**
     * @return The server&#39;s IPv4 gateway.
     * 
     */
    public String gatewayV4() {
        return this.gatewayV4;
    }
    /**
     * @return The hostname assigned to the server.
     * 
     */
    public String hostname() {
        return this.hostname;
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
     * @return The server&#39;s internal IP address.
     * 
     */
    public String internalIp() {
        return this.internalIp;
    }
    /**
     * @return The server&#39;s current KVM URL. This URL will change periodically. It is not advised to cache this value.
     * 
     */
    public String kvm() {
        return this.kvm;
    }
    /**
     * @return The server&#39;s label.
     * 
     */
    public String label() {
        return this.label;
    }
    public String location() {
        return this.location;
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
     * @return The operating system of the instance.
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
     * @return Whether the server is powered on or not.
     * 
     */
    public String powerStatus() {
        return this.powerStatus;
    }
    public List<String> privateNetworkIds() {
        return this.privateNetworkIds;
    }
    /**
     * @return The amount of memory available on the instance in MB.
     * 
     */
    public Integer ram() {
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
     * @return A more detailed server status (none, locked, installingbooting, isomounting, ok).
     * 
     */
    public String serverStatus() {
        return this.serverStatus;
    }
    /**
     * @return The status of the server&#39;s subscription.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return A list of tags applied to the instance.
     * 
     */
    public List<String> tags() {
        return this.tags;
    }
    /**
     * @return The main IPv6 network address.
     * 
     */
    public String v6MainIp() {
        return this.v6MainIp;
    }
    /**
     * @return The IPv6 subnet.
     * 
     */
    public String v6Network() {
        return this.v6Network;
    }
    /**
     * @return The IPv6 network size in bits.
     * 
     */
    public Integer v6NetworkSize() {
        return this.v6NetworkSize;
    }
    /**
     * @return The number of virtual CPUs available on the server.
     * 
     */
    public Integer vcpuCount() {
        return this.vcpuCount;
    }
    /**
     * @return A list of VPC 2.0 IDs attached to the server.
     * 
     */
    public List<String> vpc2Ids() {
        return this.vpc2Ids;
    }
    public List<String> vpcIds() {
        return this.vpcIds;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstanceResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer allowedBandwidth;
        private Integer appId;
        private String backups;
        private Map<String,Object> backupsSchedule;
        private String dateCreated;
        private Integer disk;
        private List<String> features;
        private @Nullable List<GetInstanceFilter> filters;
        private String firewallGroupId;
        private String gatewayV4;
        private String hostname;
        private String id;
        private String imageId;
        private String internalIp;
        private String kvm;
        private String label;
        private String location;
        private String mainIp;
        private String netmaskV4;
        private String os;
        private Integer osId;
        private String plan;
        private String powerStatus;
        private List<String> privateNetworkIds;
        private Integer ram;
        private String region;
        private String serverStatus;
        private String status;
        private List<String> tags;
        private String v6MainIp;
        private String v6Network;
        private Integer v6NetworkSize;
        private Integer vcpuCount;
        private List<String> vpc2Ids;
        private List<String> vpcIds;
        public Builder() {}
        public Builder(GetInstanceResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowedBandwidth = defaults.allowedBandwidth;
    	      this.appId = defaults.appId;
    	      this.backups = defaults.backups;
    	      this.backupsSchedule = defaults.backupsSchedule;
    	      this.dateCreated = defaults.dateCreated;
    	      this.disk = defaults.disk;
    	      this.features = defaults.features;
    	      this.filters = defaults.filters;
    	      this.firewallGroupId = defaults.firewallGroupId;
    	      this.gatewayV4 = defaults.gatewayV4;
    	      this.hostname = defaults.hostname;
    	      this.id = defaults.id;
    	      this.imageId = defaults.imageId;
    	      this.internalIp = defaults.internalIp;
    	      this.kvm = defaults.kvm;
    	      this.label = defaults.label;
    	      this.location = defaults.location;
    	      this.mainIp = defaults.mainIp;
    	      this.netmaskV4 = defaults.netmaskV4;
    	      this.os = defaults.os;
    	      this.osId = defaults.osId;
    	      this.plan = defaults.plan;
    	      this.powerStatus = defaults.powerStatus;
    	      this.privateNetworkIds = defaults.privateNetworkIds;
    	      this.ram = defaults.ram;
    	      this.region = defaults.region;
    	      this.serverStatus = defaults.serverStatus;
    	      this.status = defaults.status;
    	      this.tags = defaults.tags;
    	      this.v6MainIp = defaults.v6MainIp;
    	      this.v6Network = defaults.v6Network;
    	      this.v6NetworkSize = defaults.v6NetworkSize;
    	      this.vcpuCount = defaults.vcpuCount;
    	      this.vpc2Ids = defaults.vpc2Ids;
    	      this.vpcIds = defaults.vpcIds;
        }

        @CustomType.Setter
        public Builder allowedBandwidth(Integer allowedBandwidth) {
            this.allowedBandwidth = Objects.requireNonNull(allowedBandwidth);
            return this;
        }
        @CustomType.Setter
        public Builder appId(Integer appId) {
            this.appId = Objects.requireNonNull(appId);
            return this;
        }
        @CustomType.Setter
        public Builder backups(String backups) {
            this.backups = Objects.requireNonNull(backups);
            return this;
        }
        @CustomType.Setter
        public Builder backupsSchedule(Map<String,Object> backupsSchedule) {
            this.backupsSchedule = Objects.requireNonNull(backupsSchedule);
            return this;
        }
        @CustomType.Setter
        public Builder dateCreated(String dateCreated) {
            this.dateCreated = Objects.requireNonNull(dateCreated);
            return this;
        }
        @CustomType.Setter
        public Builder disk(Integer disk) {
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
        public Builder filters(@Nullable List<GetInstanceFilter> filters) {
            this.filters = filters;
            return this;
        }
        public Builder filters(GetInstanceFilter... filters) {
            return filters(List.of(filters));
        }
        @CustomType.Setter
        public Builder firewallGroupId(String firewallGroupId) {
            this.firewallGroupId = Objects.requireNonNull(firewallGroupId);
            return this;
        }
        @CustomType.Setter
        public Builder gatewayV4(String gatewayV4) {
            this.gatewayV4 = Objects.requireNonNull(gatewayV4);
            return this;
        }
        @CustomType.Setter
        public Builder hostname(String hostname) {
            this.hostname = Objects.requireNonNull(hostname);
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
        public Builder internalIp(String internalIp) {
            this.internalIp = Objects.requireNonNull(internalIp);
            return this;
        }
        @CustomType.Setter
        public Builder kvm(String kvm) {
            this.kvm = Objects.requireNonNull(kvm);
            return this;
        }
        @CustomType.Setter
        public Builder label(String label) {
            this.label = Objects.requireNonNull(label);
            return this;
        }
        @CustomType.Setter
        public Builder location(String location) {
            this.location = Objects.requireNonNull(location);
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
        public Builder powerStatus(String powerStatus) {
            this.powerStatus = Objects.requireNonNull(powerStatus);
            return this;
        }
        @CustomType.Setter
        public Builder privateNetworkIds(List<String> privateNetworkIds) {
            this.privateNetworkIds = Objects.requireNonNull(privateNetworkIds);
            return this;
        }
        public Builder privateNetworkIds(String... privateNetworkIds) {
            return privateNetworkIds(List.of(privateNetworkIds));
        }
        @CustomType.Setter
        public Builder ram(Integer ram) {
            this.ram = Objects.requireNonNull(ram);
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            this.region = Objects.requireNonNull(region);
            return this;
        }
        @CustomType.Setter
        public Builder serverStatus(String serverStatus) {
            this.serverStatus = Objects.requireNonNull(serverStatus);
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
        public Builder vcpuCount(Integer vcpuCount) {
            this.vcpuCount = Objects.requireNonNull(vcpuCount);
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
        @CustomType.Setter
        public Builder vpcIds(List<String> vpcIds) {
            this.vpcIds = Objects.requireNonNull(vpcIds);
            return this;
        }
        public Builder vpcIds(String... vpcIds) {
            return vpcIds(List.of(vpcIds));
        }
        public GetInstanceResult build() {
            final var o = new GetInstanceResult();
            o.allowedBandwidth = allowedBandwidth;
            o.appId = appId;
            o.backups = backups;
            o.backupsSchedule = backupsSchedule;
            o.dateCreated = dateCreated;
            o.disk = disk;
            o.features = features;
            o.filters = filters;
            o.firewallGroupId = firewallGroupId;
            o.gatewayV4 = gatewayV4;
            o.hostname = hostname;
            o.id = id;
            o.imageId = imageId;
            o.internalIp = internalIp;
            o.kvm = kvm;
            o.label = label;
            o.location = location;
            o.mainIp = mainIp;
            o.netmaskV4 = netmaskV4;
            o.os = os;
            o.osId = osId;
            o.plan = plan;
            o.powerStatus = powerStatus;
            o.privateNetworkIds = privateNetworkIds;
            o.ram = ram;
            o.region = region;
            o.serverStatus = serverStatus;
            o.status = status;
            o.tags = tags;
            o.v6MainIp = v6MainIp;
            o.v6Network = v6Network;
            o.v6NetworkSize = v6NetworkSize;
            o.vcpuCount = vcpuCount;
            o.vpc2Ids = vpc2Ids;
            o.vpcIds = vpcIds;
            return o;
        }
    }
}
