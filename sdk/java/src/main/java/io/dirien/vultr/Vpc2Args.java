// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class Vpc2Args extends com.pulumi.resources.ResourceArgs {

    public static final Vpc2Args Empty = new Vpc2Args();

    /**
     * The description you want to give your VPC 2.0.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description you want to give your VPC 2.0.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The IPv4 subnet to be used when attaching instances to this VPC 2.0.
     * 
     */
    @Import(name="ipBlock")
    private @Nullable Output<String> ipBlock;

    /**
     * @return The IPv4 subnet to be used when attaching instances to this VPC 2.0.
     * 
     */
    public Optional<Output<String>> ipBlock() {
        return Optional.ofNullable(this.ipBlock);
    }

    /**
     * Accepted values: `v4`.
     * 
     */
    @Import(name="ipType")
    private @Nullable Output<String> ipType;

    /**
     * @return Accepted values: `v4`.
     * 
     */
    public Optional<Output<String>> ipType() {
        return Optional.ofNullable(this.ipType);
    }

    /**
     * The number of bits for the netmask in CIDR notation. Example: 32
     * 
     */
    @Import(name="prefixLength")
    private @Nullable Output<Integer> prefixLength;

    /**
     * @return The number of bits for the netmask in CIDR notation. Example: 32
     * 
     */
    public Optional<Output<Integer>> prefixLength() {
        return Optional.ofNullable(this.prefixLength);
    }

    /**
     * The region ID that you want the VPC 2.0 to be created in.
     * 
     */
    @Import(name="region", required=true)
    private Output<String> region;

    /**
     * @return The region ID that you want the VPC 2.0 to be created in.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    private Vpc2Args() {}

    private Vpc2Args(Vpc2Args $) {
        this.description = $.description;
        this.ipBlock = $.ipBlock;
        this.ipType = $.ipType;
        this.prefixLength = $.prefixLength;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(Vpc2Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Vpc2Args $;

        public Builder() {
            $ = new Vpc2Args();
        }

        public Builder(Vpc2Args defaults) {
            $ = new Vpc2Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description you want to give your VPC 2.0.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description you want to give your VPC 2.0.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param ipBlock The IPv4 subnet to be used when attaching instances to this VPC 2.0.
         * 
         * @return builder
         * 
         */
        public Builder ipBlock(@Nullable Output<String> ipBlock) {
            $.ipBlock = ipBlock;
            return this;
        }

        /**
         * @param ipBlock The IPv4 subnet to be used when attaching instances to this VPC 2.0.
         * 
         * @return builder
         * 
         */
        public Builder ipBlock(String ipBlock) {
            return ipBlock(Output.of(ipBlock));
        }

        /**
         * @param ipType Accepted values: `v4`.
         * 
         * @return builder
         * 
         */
        public Builder ipType(@Nullable Output<String> ipType) {
            $.ipType = ipType;
            return this;
        }

        /**
         * @param ipType Accepted values: `v4`.
         * 
         * @return builder
         * 
         */
        public Builder ipType(String ipType) {
            return ipType(Output.of(ipType));
        }

        /**
         * @param prefixLength The number of bits for the netmask in CIDR notation. Example: 32
         * 
         * @return builder
         * 
         */
        public Builder prefixLength(@Nullable Output<Integer> prefixLength) {
            $.prefixLength = prefixLength;
            return this;
        }

        /**
         * @param prefixLength The number of bits for the netmask in CIDR notation. Example: 32
         * 
         * @return builder
         * 
         */
        public Builder prefixLength(Integer prefixLength) {
            return prefixLength(Output.of(prefixLength));
        }

        /**
         * @param region The region ID that you want the VPC 2.0 to be created in.
         * 
         * @return builder
         * 
         */
        public Builder region(Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region ID that you want the VPC 2.0 to be created in.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public Vpc2Args build() {
            $.region = Objects.requireNonNull($.region, "expected parameter 'region' to be non-null");
            return $;
        }
    }

}
