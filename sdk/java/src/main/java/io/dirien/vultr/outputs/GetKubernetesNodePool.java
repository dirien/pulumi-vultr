// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.outputs;

import com.pulumi.core.annotations.CustomType;
import io.dirien.vultr.outputs.GetKubernetesNodePoolNode;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetKubernetesNodePool {
    /**
     * @return Boolean indicating if the auto scaler for the default node pool is active.
     * 
     */
    private @Nullable Boolean autoScaler;
    /**
     * @return Date node was created.
     * 
     */
    private String dateCreated;
    /**
     * @return Date of node pool updates.
     * 
     */
    private String dateUpdated;
    /**
     * @return ID of node.
     * 
     */
    private String id;
    /**
     * @return Label of node.
     * 
     */
    private String label;
    /**
     * @return The maximum number of nodes used by the auto scaler.
     * 
     */
    private @Nullable Integer maxNodes;
    /**
     * @return The minimum number of nodes used by the auto scaler.
     * 
     */
    private @Nullable Integer minNodes;
    /**
     * @return Number of nodes within node pool.
     * 
     */
    private Integer nodeQuantity;
    /**
     * @return Array that contains information about nodes within this node pool.
     * 
     */
    private List<GetKubernetesNodePoolNode> nodes;
    /**
     * @return Node plan that nodes are using within this node pool.
     * 
     */
    private String plan;
    /**
     * @return Status of node.
     * 
     */
    private String status;
    /**
     * @return Tag for node pool.
     * 
     */
    private String tag;

    private GetKubernetesNodePool() {}
    /**
     * @return Boolean indicating if the auto scaler for the default node pool is active.
     * 
     */
    public Optional<Boolean> autoScaler() {
        return Optional.ofNullable(this.autoScaler);
    }
    /**
     * @return Date node was created.
     * 
     */
    public String dateCreated() {
        return this.dateCreated;
    }
    /**
     * @return Date of node pool updates.
     * 
     */
    public String dateUpdated() {
        return this.dateUpdated;
    }
    /**
     * @return ID of node.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Label of node.
     * 
     */
    public String label() {
        return this.label;
    }
    /**
     * @return The maximum number of nodes used by the auto scaler.
     * 
     */
    public Optional<Integer> maxNodes() {
        return Optional.ofNullable(this.maxNodes);
    }
    /**
     * @return The minimum number of nodes used by the auto scaler.
     * 
     */
    public Optional<Integer> minNodes() {
        return Optional.ofNullable(this.minNodes);
    }
    /**
     * @return Number of nodes within node pool.
     * 
     */
    public Integer nodeQuantity() {
        return this.nodeQuantity;
    }
    /**
     * @return Array that contains information about nodes within this node pool.
     * 
     */
    public List<GetKubernetesNodePoolNode> nodes() {
        return this.nodes;
    }
    /**
     * @return Node plan that nodes are using within this node pool.
     * 
     */
    public String plan() {
        return this.plan;
    }
    /**
     * @return Status of node.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return Tag for node pool.
     * 
     */
    public String tag() {
        return this.tag;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetKubernetesNodePool defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean autoScaler;
        private String dateCreated;
        private String dateUpdated;
        private String id;
        private String label;
        private @Nullable Integer maxNodes;
        private @Nullable Integer minNodes;
        private Integer nodeQuantity;
        private List<GetKubernetesNodePoolNode> nodes;
        private String plan;
        private String status;
        private String tag;
        public Builder() {}
        public Builder(GetKubernetesNodePool defaults) {
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
        public Builder dateCreated(String dateCreated) {
            this.dateCreated = Objects.requireNonNull(dateCreated);
            return this;
        }
        @CustomType.Setter
        public Builder dateUpdated(String dateUpdated) {
            this.dateUpdated = Objects.requireNonNull(dateUpdated);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
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
        public Builder nodes(List<GetKubernetesNodePoolNode> nodes) {
            this.nodes = Objects.requireNonNull(nodes);
            return this;
        }
        public Builder nodes(GetKubernetesNodePoolNode... nodes) {
            return nodes(List.of(nodes));
        }
        @CustomType.Setter
        public Builder plan(String plan) {
            this.plan = Objects.requireNonNull(plan);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder tag(String tag) {
            this.tag = Objects.requireNonNull(tag);
            return this;
        }
        public GetKubernetesNodePool build() {
            final var o = new GetKubernetesNodePool();
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
