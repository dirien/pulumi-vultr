// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.outputs;

import com.pulumi.core.annotations.CustomType;
import io.dirien.vultr.outputs.KubernetesNodePoolsNode;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class KubernetesNodePools {
    /**
     * @return Enable the auto scaler for the default node pool.
     * 
     */
    private @Nullable Boolean autoScaler;
    /**
     * @return Date node was created.
     * 
     */
    private @Nullable String dateCreated;
    /**
     * @return Date of node pool updates.
     * 
     */
    private @Nullable String dateUpdated;
    /**
     * @return ID of node.
     * 
     */
    private @Nullable String id;
    /**
     * @return The label to be used as a prefix for nodes in this node pool.
     * 
     */
    private String label;
    /**
     * @return The maximum number of nodes to use with the auto scaler.
     * 
     */
    private @Nullable Integer maxNodes;
    /**
     * @return The minimum number of nodes to use with the auto scaler.
     * 
     */
    private @Nullable Integer minNodes;
    /**
     * @return The number of nodes in this node pool.
     * 
     */
    private Integer nodeQuantity;
    /**
     * @return Array that contains information about nodes within this node pool.
     * 
     */
    private @Nullable List<KubernetesNodePoolsNode> nodes;
    /**
     * @return The plan to be used in this node pool. [See Plans List](https://www.vultr.com/api/#operation/list-plans) Note the minimum plan requirements must have at least 1 core and 2 gbs of memory.
     * 
     */
    private String plan;
    /**
     * @return Status of node.
     * 
     */
    private @Nullable String status;
    /**
     * @return Tag for node pool.
     * 
     */
    private @Nullable String tag;

    private KubernetesNodePools() {}
    /**
     * @return Enable the auto scaler for the default node pool.
     * 
     */
    public Optional<Boolean> autoScaler() {
        return Optional.ofNullable(this.autoScaler);
    }
    /**
     * @return Date node was created.
     * 
     */
    public Optional<String> dateCreated() {
        return Optional.ofNullable(this.dateCreated);
    }
    /**
     * @return Date of node pool updates.
     * 
     */
    public Optional<String> dateUpdated() {
        return Optional.ofNullable(this.dateUpdated);
    }
    /**
     * @return ID of node.
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }
    /**
     * @return The label to be used as a prefix for nodes in this node pool.
     * 
     */
    public String label() {
        return this.label;
    }
    /**
     * @return The maximum number of nodes to use with the auto scaler.
     * 
     */
    public Optional<Integer> maxNodes() {
        return Optional.ofNullable(this.maxNodes);
    }
    /**
     * @return The minimum number of nodes to use with the auto scaler.
     * 
     */
    public Optional<Integer> minNodes() {
        return Optional.ofNullable(this.minNodes);
    }
    /**
     * @return The number of nodes in this node pool.
     * 
     */
    public Integer nodeQuantity() {
        return this.nodeQuantity;
    }
    /**
     * @return Array that contains information about nodes within this node pool.
     * 
     */
    public List<KubernetesNodePoolsNode> nodes() {
        return this.nodes == null ? List.of() : this.nodes;
    }
    /**
     * @return The plan to be used in this node pool. [See Plans List](https://www.vultr.com/api/#operation/list-plans) Note the minimum plan requirements must have at least 1 core and 2 gbs of memory.
     * 
     */
    public String plan() {
        return this.plan;
    }
    /**
     * @return Status of node.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    /**
     * @return Tag for node pool.
     * 
     */
    public Optional<String> tag() {
        return Optional.ofNullable(this.tag);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(KubernetesNodePools defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean autoScaler;
        private @Nullable String dateCreated;
        private @Nullable String dateUpdated;
        private @Nullable String id;
        private String label;
        private @Nullable Integer maxNodes;
        private @Nullable Integer minNodes;
        private Integer nodeQuantity;
        private @Nullable List<KubernetesNodePoolsNode> nodes;
        private String plan;
        private @Nullable String status;
        private @Nullable String tag;
        public Builder() {}
        public Builder(KubernetesNodePools defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.autoScaler = defaults.autoScaler;
    	      this.dateCreated = defaults.dateCreated;
    	      this.dateUpdated = defaults.dateUpdated;
    	      this.id = defaults.id;
    	      this.label = defaults.label;
    	      this.maxNodes = defaults.maxNodes;
    	      this.minNodes = defaults.minNodes;
    	      this.nodeQuantity = defaults.nodeQuantity;
    	      this.nodes = defaults.nodes;
    	      this.plan = defaults.plan;
    	      this.status = defaults.status;
    	      this.tag = defaults.tag;
        }

        @CustomType.Setter
        public Builder autoScaler(@Nullable Boolean autoScaler) {
            this.autoScaler = autoScaler;
            return this;
        }
        @CustomType.Setter
        public Builder dateCreated(@Nullable String dateCreated) {
            this.dateCreated = dateCreated;
            return this;
        }
        @CustomType.Setter
        public Builder dateUpdated(@Nullable String dateUpdated) {
            this.dateUpdated = dateUpdated;
            return this;
        }
        @CustomType.Setter
        public Builder id(@Nullable String id) {
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder label(String label) {
            this.label = Objects.requireNonNull(label);
            return this;
        }
        @CustomType.Setter
        public Builder maxNodes(@Nullable Integer maxNodes) {
            this.maxNodes = maxNodes;
            return this;
        }
        @CustomType.Setter
        public Builder minNodes(@Nullable Integer minNodes) {
            this.minNodes = minNodes;
            return this;
        }
        @CustomType.Setter
        public Builder nodeQuantity(Integer nodeQuantity) {
            this.nodeQuantity = Objects.requireNonNull(nodeQuantity);
            return this;
        }
        @CustomType.Setter
        public Builder nodes(@Nullable List<KubernetesNodePoolsNode> nodes) {
            this.nodes = nodes;
            return this;
        }
        public Builder nodes(KubernetesNodePoolsNode... nodes) {
            return nodes(List.of(nodes));
        }
        @CustomType.Setter
        public Builder plan(String plan) {
            this.plan = Objects.requireNonNull(plan);
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder tag(@Nullable String tag) {
            this.tag = tag;
            return this;
        }
        public KubernetesNodePools build() {
            final var o = new KubernetesNodePools();
            o.autoScaler = autoScaler;
            o.dateCreated = dateCreated;
            o.dateUpdated = dateUpdated;
            o.id = id;
            o.label = label;
            o.maxNodes = maxNodes;
            o.minNodes = minNodes;
            o.nodeQuantity = nodeQuantity;
            o.nodes = nodes;
            o.plan = plan;
            o.status = status;
            o.tag = tag;
            return o;
        }
    }
}
