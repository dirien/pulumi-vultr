// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import io.dirien.vultr.inputs.InstanceBackupsScheduleArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceArgs Empty = new InstanceArgs();

    /**
     * Whether an activation email will be sent when the server is ready.
     * 
     */
    @Import(name="activationEmail")
    private @Nullable Output<Boolean> activationEmail;

    /**
     * @return Whether an activation email will be sent when the server is ready.
     * 
     */
    public Optional<Output<Boolean>> activationEmail() {
        return Optional.ofNullable(this.activationEmail);
    }

    /**
     * The ID of the Vultr application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications)
     * 
     */
    @Import(name="appId")
    private @Nullable Output<Integer> appId;

    /**
     * @return The ID of the Vultr application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications)
     * 
     */
    public Optional<Output<Integer>> appId() {
        return Optional.ofNullable(this.appId);
    }

    /**
     * Whether automatic backups will be enabled for this server (these have an extra charge associated with them). Values can be enabled or disabled.
     * 
     */
    @Import(name="backups")
    private @Nullable Output<String> backups;

    /**
     * @return Whether automatic backups will be enabled for this server (these have an extra charge associated with them). Values can be enabled or disabled.
     * 
     */
    public Optional<Output<String>> backups() {
        return Optional.ofNullable(this.backups);
    }

    /**
     * A block that defines the way backups should be scheduled. While this is an optional field if `backups` are `enabled` this field is mandatory. The configuration of a `backups_schedule` is listed below.
     * 
     */
    @Import(name="backupsSchedule")
    private @Nullable Output<InstanceBackupsScheduleArgs> backupsSchedule;

    /**
     * @return A block that defines the way backups should be scheduled. While this is an optional field if `backups` are `enabled` this field is mandatory. The configuration of a `backups_schedule` is listed below.
     * 
     */
    public Optional<Output<InstanceBackupsScheduleArgs>> backupsSchedule() {
        return Optional.ofNullable(this.backupsSchedule);
    }

    /**
     * Whether DDOS protection will be enabled on the server (there is an additional charge for this).
     * 
     */
    @Import(name="ddosProtection")
    private @Nullable Output<Boolean> ddosProtection;

    /**
     * @return Whether DDOS protection will be enabled on the server (there is an additional charge for this).
     * 
     */
    public Optional<Output<Boolean>> ddosProtection() {
        return Optional.ofNullable(this.ddosProtection);
    }

    /**
     * Whether the server has IPv6 networking activated.
     * 
     */
    @Import(name="enableIpv6")
    private @Nullable Output<Boolean> enableIpv6;

    /**
     * @return Whether the server has IPv6 networking activated.
     * 
     */
    public Optional<Output<Boolean>> enableIpv6() {
        return Optional.ofNullable(this.enableIpv6);
    }

    /**
     * The ID of the firewall group to assign to the server.
     * 
     */
    @Import(name="firewallGroupId")
    private @Nullable Output<String> firewallGroupId;

    /**
     * @return The ID of the firewall group to assign to the server.
     * 
     */
    public Optional<Output<String>> firewallGroupId() {
        return Optional.ofNullable(this.firewallGroupId);
    }

    /**
     * The hostname to assign to the server.
     * 
     */
    @Import(name="hostname")
    private @Nullable Output<String> hostname;

    /**
     * @return The hostname to assign to the server.
     * 
     */
    public Optional<Output<String>> hostname() {
        return Optional.ofNullable(this.hostname);
    }

    /**
     * The ID of the Vultr marketplace application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications) Note marketplace applications are denoted by type: `marketplace` and you must use the `image_id` not the id.
     * 
     */
    @Import(name="imageId")
    private @Nullable Output<String> imageId;

    /**
     * @return The ID of the Vultr marketplace application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications) Note marketplace applications are denoted by type: `marketplace` and you must use the `image_id` not the id.
     * 
     */
    public Optional<Output<String>> imageId() {
        return Optional.ofNullable(this.imageId);
    }

    /**
     * The ID of the ISO file to be installed on the server. [See List ISO](https://www.vultr.com/api/#operation/list-isos)
     * 
     */
    @Import(name="isoId")
    private @Nullable Output<String> isoId;

    /**
     * @return The ID of the ISO file to be installed on the server. [See List ISO](https://www.vultr.com/api/#operation/list-isos)
     * 
     */
    public Optional<Output<String>> isoId() {
        return Optional.ofNullable(this.isoId);
    }

    /**
     * A label for the server.
     * 
     */
    @Import(name="label")
    private @Nullable Output<String> label;

    /**
     * @return A label for the server.
     * 
     */
    public Optional<Output<String>> label() {
        return Optional.ofNullable(this.label);
    }

    /**
     * The ID of the operating system to be installed on the server. [See List OS](https://www.vultr.com/api/#operation/list-os)
     * 
     */
    @Import(name="osId")
    private @Nullable Output<Integer> osId;

    /**
     * @return The ID of the operating system to be installed on the server. [See List OS](https://www.vultr.com/api/#operation/list-os)
     * 
     */
    public Optional<Output<Integer>> osId() {
        return Optional.ofNullable(this.osId);
    }

    /**
     * The ID of the plan that you want the instance to subscribe to. [See List Plans](https://www.vultr.com/api/#tag/plans)
     * 
     */
    @Import(name="plan", required=true)
    private Output<String> plan;

    /**
     * @return The ID of the plan that you want the instance to subscribe to. [See List Plans](https://www.vultr.com/api/#tag/plans)
     * 
     */
    public Output<String> plan() {
        return this.plan;
    }

    /**
     * (Deprecated: use `vpc_ids` instead) A list of private network IDs to be attached to the server.
     * 
     * @deprecated
     * private_network_ids has been deprecated and should no longer be used. Instead, use vpc_ids
     * 
     */
    @Deprecated /* private_network_ids has been deprecated and should no longer be used. Instead, use vpc_ids */
    @Import(name="privateNetworkIds")
    private @Nullable Output<List<String>> privateNetworkIds;

    /**
     * @return (Deprecated: use `vpc_ids` instead) A list of private network IDs to be attached to the server.
     * 
     * @deprecated
     * private_network_ids has been deprecated and should no longer be used. Instead, use vpc_ids
     * 
     */
    @Deprecated /* private_network_ids has been deprecated and should no longer be used. Instead, use vpc_ids */
    public Optional<Output<List<String>>> privateNetworkIds() {
        return Optional.ofNullable(this.privateNetworkIds);
    }

    /**
     * The ID of the region that the instance is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
     * 
     */
    @Import(name="region", required=true)
    private Output<String> region;

    /**
     * @return The ID of the region that the instance is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     * ID of the floating IP to use as the main IP of this server.
     * 
     */
    @Import(name="reservedIpId")
    private @Nullable Output<String> reservedIpId;

    /**
     * @return ID of the floating IP to use as the main IP of this server.
     * 
     */
    public Optional<Output<String>> reservedIpId() {
        return Optional.ofNullable(this.reservedIpId);
    }

    /**
     * The ID of the startup script you want added to the server.
     * 
     */
    @Import(name="scriptId")
    private @Nullable Output<String> scriptId;

    /**
     * @return The ID of the startup script you want added to the server.
     * 
     */
    public Optional<Output<String>> scriptId() {
        return Optional.ofNullable(this.scriptId);
    }

    /**
     * The ID of the Vultr snapshot that the server will restore for the initial installation. [See List Snapshots](https://www.vultr.com/api/#operation/list-snapshots)
     * 
     */
    @Import(name="snapshotId")
    private @Nullable Output<String> snapshotId;

    /**
     * @return The ID of the Vultr snapshot that the server will restore for the initial installation. [See List Snapshots](https://www.vultr.com/api/#operation/list-snapshots)
     * 
     */
    public Optional<Output<String>> snapshotId() {
        return Optional.ofNullable(this.snapshotId);
    }

    /**
     * A list of SSH key IDs to apply to the server on install (only valid for Linux/FreeBSD).
     * 
     */
    @Import(name="sshKeyIds")
    private @Nullable Output<List<String>> sshKeyIds;

    /**
     * @return A list of SSH key IDs to apply to the server on install (only valid for Linux/FreeBSD).
     * 
     */
    public Optional<Output<List<String>>> sshKeyIds() {
        return Optional.ofNullable(this.sshKeyIds);
    }

    /**
     * A list of tags to apply to the instance.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return A list of tags to apply to the instance.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.
     * 
     */
    @Import(name="userData")
    private @Nullable Output<String> userData;

    /**
     * @return Generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.
     * 
     */
    public Optional<Output<String>> userData() {
        return Optional.ofNullable(this.userData);
    }

    /**
     * A list of VPC IDs to be attached to the server.
     * 
     */
    @Import(name="vpcIds")
    private @Nullable Output<List<String>> vpcIds;

    /**
     * @return A list of VPC IDs to be attached to the server.
     * 
     */
    public Optional<Output<List<String>>> vpcIds() {
        return Optional.ofNullable(this.vpcIds);
    }

    private InstanceArgs() {}

    private InstanceArgs(InstanceArgs $) {
        this.activationEmail = $.activationEmail;
        this.appId = $.appId;
        this.backups = $.backups;
        this.backupsSchedule = $.backupsSchedule;
        this.ddosProtection = $.ddosProtection;
        this.enableIpv6 = $.enableIpv6;
        this.firewallGroupId = $.firewallGroupId;
        this.hostname = $.hostname;
        this.imageId = $.imageId;
        this.isoId = $.isoId;
        this.label = $.label;
        this.osId = $.osId;
        this.plan = $.plan;
        this.privateNetworkIds = $.privateNetworkIds;
        this.region = $.region;
        this.reservedIpId = $.reservedIpId;
        this.scriptId = $.scriptId;
        this.snapshotId = $.snapshotId;
        this.sshKeyIds = $.sshKeyIds;
        this.tags = $.tags;
        this.userData = $.userData;
        this.vpcIds = $.vpcIds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceArgs $;

        public Builder() {
            $ = new InstanceArgs();
        }

        public Builder(InstanceArgs defaults) {
            $ = new InstanceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param activationEmail Whether an activation email will be sent when the server is ready.
         * 
         * @return builder
         * 
         */
        public Builder activationEmail(@Nullable Output<Boolean> activationEmail) {
            $.activationEmail = activationEmail;
            return this;
        }

        /**
         * @param activationEmail Whether an activation email will be sent when the server is ready.
         * 
         * @return builder
         * 
         */
        public Builder activationEmail(Boolean activationEmail) {
            return activationEmail(Output.of(activationEmail));
        }

        /**
         * @param appId The ID of the Vultr application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications)
         * 
         * @return builder
         * 
         */
        public Builder appId(@Nullable Output<Integer> appId) {
            $.appId = appId;
            return this;
        }

        /**
         * @param appId The ID of the Vultr application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications)
         * 
         * @return builder
         * 
         */
        public Builder appId(Integer appId) {
            return appId(Output.of(appId));
        }

        /**
         * @param backups Whether automatic backups will be enabled for this server (these have an extra charge associated with them). Values can be enabled or disabled.
         * 
         * @return builder
         * 
         */
        public Builder backups(@Nullable Output<String> backups) {
            $.backups = backups;
            return this;
        }

        /**
         * @param backups Whether automatic backups will be enabled for this server (these have an extra charge associated with them). Values can be enabled or disabled.
         * 
         * @return builder
         * 
         */
        public Builder backups(String backups) {
            return backups(Output.of(backups));
        }

        /**
         * @param backupsSchedule A block that defines the way backups should be scheduled. While this is an optional field if `backups` are `enabled` this field is mandatory. The configuration of a `backups_schedule` is listed below.
         * 
         * @return builder
         * 
         */
        public Builder backupsSchedule(@Nullable Output<InstanceBackupsScheduleArgs> backupsSchedule) {
            $.backupsSchedule = backupsSchedule;
            return this;
        }

        /**
         * @param backupsSchedule A block that defines the way backups should be scheduled. While this is an optional field if `backups` are `enabled` this field is mandatory. The configuration of a `backups_schedule` is listed below.
         * 
         * @return builder
         * 
         */
        public Builder backupsSchedule(InstanceBackupsScheduleArgs backupsSchedule) {
            return backupsSchedule(Output.of(backupsSchedule));
        }

        /**
         * @param ddosProtection Whether DDOS protection will be enabled on the server (there is an additional charge for this).
         * 
         * @return builder
         * 
         */
        public Builder ddosProtection(@Nullable Output<Boolean> ddosProtection) {
            $.ddosProtection = ddosProtection;
            return this;
        }

        /**
         * @param ddosProtection Whether DDOS protection will be enabled on the server (there is an additional charge for this).
         * 
         * @return builder
         * 
         */
        public Builder ddosProtection(Boolean ddosProtection) {
            return ddosProtection(Output.of(ddosProtection));
        }

        /**
         * @param enableIpv6 Whether the server has IPv6 networking activated.
         * 
         * @return builder
         * 
         */
        public Builder enableIpv6(@Nullable Output<Boolean> enableIpv6) {
            $.enableIpv6 = enableIpv6;
            return this;
        }

        /**
         * @param enableIpv6 Whether the server has IPv6 networking activated.
         * 
         * @return builder
         * 
         */
        public Builder enableIpv6(Boolean enableIpv6) {
            return enableIpv6(Output.of(enableIpv6));
        }

        /**
         * @param firewallGroupId The ID of the firewall group to assign to the server.
         * 
         * @return builder
         * 
         */
        public Builder firewallGroupId(@Nullable Output<String> firewallGroupId) {
            $.firewallGroupId = firewallGroupId;
            return this;
        }

        /**
         * @param firewallGroupId The ID of the firewall group to assign to the server.
         * 
         * @return builder
         * 
         */
        public Builder firewallGroupId(String firewallGroupId) {
            return firewallGroupId(Output.of(firewallGroupId));
        }

        /**
         * @param hostname The hostname to assign to the server.
         * 
         * @return builder
         * 
         */
        public Builder hostname(@Nullable Output<String> hostname) {
            $.hostname = hostname;
            return this;
        }

        /**
         * @param hostname The hostname to assign to the server.
         * 
         * @return builder
         * 
         */
        public Builder hostname(String hostname) {
            return hostname(Output.of(hostname));
        }

        /**
         * @param imageId The ID of the Vultr marketplace application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications) Note marketplace applications are denoted by type: `marketplace` and you must use the `image_id` not the id.
         * 
         * @return builder
         * 
         */
        public Builder imageId(@Nullable Output<String> imageId) {
            $.imageId = imageId;
            return this;
        }

        /**
         * @param imageId The ID of the Vultr marketplace application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications) Note marketplace applications are denoted by type: `marketplace` and you must use the `image_id` not the id.
         * 
         * @return builder
         * 
         */
        public Builder imageId(String imageId) {
            return imageId(Output.of(imageId));
        }

        /**
         * @param isoId The ID of the ISO file to be installed on the server. [See List ISO](https://www.vultr.com/api/#operation/list-isos)
         * 
         * @return builder
         * 
         */
        public Builder isoId(@Nullable Output<String> isoId) {
            $.isoId = isoId;
            return this;
        }

        /**
         * @param isoId The ID of the ISO file to be installed on the server. [See List ISO](https://www.vultr.com/api/#operation/list-isos)
         * 
         * @return builder
         * 
         */
        public Builder isoId(String isoId) {
            return isoId(Output.of(isoId));
        }

        /**
         * @param label A label for the server.
         * 
         * @return builder
         * 
         */
        public Builder label(@Nullable Output<String> label) {
            $.label = label;
            return this;
        }

        /**
         * @param label A label for the server.
         * 
         * @return builder
         * 
         */
        public Builder label(String label) {
            return label(Output.of(label));
        }

        /**
         * @param osId The ID of the operating system to be installed on the server. [See List OS](https://www.vultr.com/api/#operation/list-os)
         * 
         * @return builder
         * 
         */
        public Builder osId(@Nullable Output<Integer> osId) {
            $.osId = osId;
            return this;
        }

        /**
         * @param osId The ID of the operating system to be installed on the server. [See List OS](https://www.vultr.com/api/#operation/list-os)
         * 
         * @return builder
         * 
         */
        public Builder osId(Integer osId) {
            return osId(Output.of(osId));
        }

        /**
         * @param plan The ID of the plan that you want the instance to subscribe to. [See List Plans](https://www.vultr.com/api/#tag/plans)
         * 
         * @return builder
         * 
         */
        public Builder plan(Output<String> plan) {
            $.plan = plan;
            return this;
        }

        /**
         * @param plan The ID of the plan that you want the instance to subscribe to. [See List Plans](https://www.vultr.com/api/#tag/plans)
         * 
         * @return builder
         * 
         */
        public Builder plan(String plan) {
            return plan(Output.of(plan));
        }

        /**
         * @param privateNetworkIds (Deprecated: use `vpc_ids` instead) A list of private network IDs to be attached to the server.
         * 
         * @return builder
         * 
         * @deprecated
         * private_network_ids has been deprecated and should no longer be used. Instead, use vpc_ids
         * 
         */
        @Deprecated /* private_network_ids has been deprecated and should no longer be used. Instead, use vpc_ids */
        public Builder privateNetworkIds(@Nullable Output<List<String>> privateNetworkIds) {
            $.privateNetworkIds = privateNetworkIds;
            return this;
        }

        /**
         * @param privateNetworkIds (Deprecated: use `vpc_ids` instead) A list of private network IDs to be attached to the server.
         * 
         * @return builder
         * 
         * @deprecated
         * private_network_ids has been deprecated and should no longer be used. Instead, use vpc_ids
         * 
         */
        @Deprecated /* private_network_ids has been deprecated and should no longer be used. Instead, use vpc_ids */
        public Builder privateNetworkIds(List<String> privateNetworkIds) {
            return privateNetworkIds(Output.of(privateNetworkIds));
        }

        /**
         * @param privateNetworkIds (Deprecated: use `vpc_ids` instead) A list of private network IDs to be attached to the server.
         * 
         * @return builder
         * 
         * @deprecated
         * private_network_ids has been deprecated and should no longer be used. Instead, use vpc_ids
         * 
         */
        @Deprecated /* private_network_ids has been deprecated and should no longer be used. Instead, use vpc_ids */
        public Builder privateNetworkIds(String... privateNetworkIds) {
            return privateNetworkIds(List.of(privateNetworkIds));
        }

        /**
         * @param region The ID of the region that the instance is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
         * 
         * @return builder
         * 
         */
        public Builder region(Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The ID of the region that the instance is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param reservedIpId ID of the floating IP to use as the main IP of this server.
         * 
         * @return builder
         * 
         */
        public Builder reservedIpId(@Nullable Output<String> reservedIpId) {
            $.reservedIpId = reservedIpId;
            return this;
        }

        /**
         * @param reservedIpId ID of the floating IP to use as the main IP of this server.
         * 
         * @return builder
         * 
         */
        public Builder reservedIpId(String reservedIpId) {
            return reservedIpId(Output.of(reservedIpId));
        }

        /**
         * @param scriptId The ID of the startup script you want added to the server.
         * 
         * @return builder
         * 
         */
        public Builder scriptId(@Nullable Output<String> scriptId) {
            $.scriptId = scriptId;
            return this;
        }

        /**
         * @param scriptId The ID of the startup script you want added to the server.
         * 
         * @return builder
         * 
         */
        public Builder scriptId(String scriptId) {
            return scriptId(Output.of(scriptId));
        }

        /**
         * @param snapshotId The ID of the Vultr snapshot that the server will restore for the initial installation. [See List Snapshots](https://www.vultr.com/api/#operation/list-snapshots)
         * 
         * @return builder
         * 
         */
        public Builder snapshotId(@Nullable Output<String> snapshotId) {
            $.snapshotId = snapshotId;
            return this;
        }

        /**
         * @param snapshotId The ID of the Vultr snapshot that the server will restore for the initial installation. [See List Snapshots](https://www.vultr.com/api/#operation/list-snapshots)
         * 
         * @return builder
         * 
         */
        public Builder snapshotId(String snapshotId) {
            return snapshotId(Output.of(snapshotId));
        }

        /**
         * @param sshKeyIds A list of SSH key IDs to apply to the server on install (only valid for Linux/FreeBSD).
         * 
         * @return builder
         * 
         */
        public Builder sshKeyIds(@Nullable Output<List<String>> sshKeyIds) {
            $.sshKeyIds = sshKeyIds;
            return this;
        }

        /**
         * @param sshKeyIds A list of SSH key IDs to apply to the server on install (only valid for Linux/FreeBSD).
         * 
         * @return builder
         * 
         */
        public Builder sshKeyIds(List<String> sshKeyIds) {
            return sshKeyIds(Output.of(sshKeyIds));
        }

        /**
         * @param sshKeyIds A list of SSH key IDs to apply to the server on install (only valid for Linux/FreeBSD).
         * 
         * @return builder
         * 
         */
        public Builder sshKeyIds(String... sshKeyIds) {
            return sshKeyIds(List.of(sshKeyIds));
        }

        /**
         * @param tags A list of tags to apply to the instance.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A list of tags to apply to the instance.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags A list of tags to apply to the instance.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param userData Generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.
         * 
         * @return builder
         * 
         */
        public Builder userData(@Nullable Output<String> userData) {
            $.userData = userData;
            return this;
        }

        /**
         * @param userData Generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.
         * 
         * @return builder
         * 
         */
        public Builder userData(String userData) {
            return userData(Output.of(userData));
        }

        /**
         * @param vpcIds A list of VPC IDs to be attached to the server.
         * 
         * @return builder
         * 
         */
        public Builder vpcIds(@Nullable Output<List<String>> vpcIds) {
            $.vpcIds = vpcIds;
            return this;
        }

        /**
         * @param vpcIds A list of VPC IDs to be attached to the server.
         * 
         * @return builder
         * 
         */
        public Builder vpcIds(List<String> vpcIds) {
            return vpcIds(Output.of(vpcIds));
        }

        /**
         * @param vpcIds A list of VPC IDs to be attached to the server.
         * 
         * @return builder
         * 
         */
        public Builder vpcIds(String... vpcIds) {
            return vpcIds(List.of(vpcIds));
        }

        public InstanceArgs build() {
            $.plan = Objects.requireNonNull($.plan, "expected parameter 'plan' to be non-null");
            $.region = Objects.requireNonNull($.region, "expected parameter 'region' to be non-null");
            return $;
        }
    }

}