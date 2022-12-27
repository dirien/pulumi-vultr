// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import io.dirien.vultr.inputs.KubernetesNodePoolsArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class KubernetesArgs extends com.pulumi.resources.ResourceArgs {

    public static final KubernetesArgs Empty = new KubernetesArgs();

    /**
     * The VKE clusters label.
     * 
     */
    @Import(name="label", required=true)
    private Output<String> label;

    /**
     * @return The VKE clusters label.
     * 
     */
    public Output<String> label() {
        return this.label;
    }

    /**
     * Contains the default node pool that was deployed.
     * 
     */
    @Import(name="nodePools")
    private @Nullable Output<KubernetesNodePoolsArgs> nodePools;

    /**
     * @return Contains the default node pool that was deployed.
     * 
     */
    public Optional<Output<KubernetesNodePoolsArgs>> nodePools() {
        return Optional.ofNullable(this.nodePools);
    }

    /**
     * The region your VKE cluster will be deployed in. Currently, supported values are `ewr` and `lax`
     * 
     */
    @Import(name="region", required=true)
    private Output<String> region;

    /**
     * @return The region your VKE cluster will be deployed in. Currently, supported values are `ewr` and `lax`
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     * The version your VKE cluster you want deployed. [See Available Version](https://www.vultr.com/api/#operation/get-kubernetes-versions)
     * 
     */
    @Import(name="version", required=true)
    private Output<String> version;

    /**
     * @return The version your VKE cluster you want deployed. [See Available Version](https://www.vultr.com/api/#operation/get-kubernetes-versions)
     * 
     */
    public Output<String> version() {
        return this.version;
    }

    private KubernetesArgs() {}

    private KubernetesArgs(KubernetesArgs $) {
        this.label = $.label;
        this.nodePools = $.nodePools;
        this.region = $.region;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KubernetesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KubernetesArgs $;

        public Builder() {
            $ = new KubernetesArgs();
        }

        public Builder(KubernetesArgs defaults) {
            $ = new KubernetesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param label The VKE clusters label.
         * 
         * @return builder
         * 
         */
        public Builder label(Output<String> label) {
            $.label = label;
            return this;
        }

        /**
         * @param label The VKE clusters label.
         * 
         * @return builder
         * 
         */
        public Builder label(String label) {
            return label(Output.of(label));
        }

        /**
         * @param nodePools Contains the default node pool that was deployed.
         * 
         * @return builder
         * 
         */
        public Builder nodePools(@Nullable Output<KubernetesNodePoolsArgs> nodePools) {
            $.nodePools = nodePools;
            return this;
        }

        /**
         * @param nodePools Contains the default node pool that was deployed.
         * 
         * @return builder
         * 
         */
        public Builder nodePools(KubernetesNodePoolsArgs nodePools) {
            return nodePools(Output.of(nodePools));
        }

        /**
         * @param region The region your VKE cluster will be deployed in. Currently, supported values are `ewr` and `lax`
         * 
         * @return builder
         * 
         */
        public Builder region(Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region your VKE cluster will be deployed in. Currently, supported values are `ewr` and `lax`
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param version The version your VKE cluster you want deployed. [See Available Version](https://www.vultr.com/api/#operation/get-kubernetes-versions)
         * 
         * @return builder
         * 
         */
        public Builder version(Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version The version your VKE cluster you want deployed. [See Available Version](https://www.vultr.com/api/#operation/get-kubernetes-versions)
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        public KubernetesArgs build() {
            $.label = Objects.requireNonNull($.label, "expected parameter 'label' to be non-null");
            $.region = Objects.requireNonNull($.region, "expected parameter 'region' to be non-null");
            $.version = Objects.requireNonNull($.version, "expected parameter 'version' to be non-null");
            return $;
        }
    }

}
