// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import io.dirien.vultr.KubernetesNodePoolsArgs;
import io.dirien.vultr.Utilities;
import io.dirien.vultr.inputs.KubernetesNodePoolsState;
import io.dirien.vultr.outputs.KubernetesNodePoolsNode;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Deploy additional node pools to an existing Vultr Kubernetes Engine (VKE) cluster.
 * 
 * ## Example Usage
 * 
 * Create a new VKE cluster:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vultr.KubernetesNodePools;
 * import com.pulumi.vultr.KubernetesNodePoolsArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var np_1 = new KubernetesNodePools(&#34;np-1&#34;, KubernetesNodePoolsArgs.builder()        
 *             .clusterId(vultr_kubernetes.k8().id())
 *             .nodeQuantity(1)
 *             .plan(&#34;vc2-2c-4gb&#34;)
 *             .label(&#34;my-label&#34;)
 *             .tag(&#34;my-tag&#34;)
 *             .autoScaler(true)
 *             .minNodes(1)
 *             .maxNodes(2)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="vultr:index/kubernetesNodePools:KubernetesNodePools")
public class KubernetesNodePools extends com.pulumi.resources.CustomResource {
    /**
     * Enable the auto scaler for the default node pool.
     * 
     */
    @Export(name="autoScaler", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> autoScaler;

    /**
     * @return Enable the auto scaler for the default node pool.
     * 
     */
    public Output<Optional<Boolean>> autoScaler() {
        return Codegen.optional(this.autoScaler);
    }
    /**
     * The VKE cluster ID you want to attach this nodepool to.
     * 
     */
    @Export(name="clusterId", refs={String.class}, tree="[0]")
    private Output<String> clusterId;

    /**
     * @return The VKE cluster ID you want to attach this nodepool to.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }
    /**
     * Date node was created.
     * 
     */
    @Export(name="dateCreated", refs={String.class}, tree="[0]")
    private Output<String> dateCreated;

    /**
     * @return Date node was created.
     * 
     */
    public Output<String> dateCreated() {
        return this.dateCreated;
    }
    /**
     * Date of node pool updates.
     * 
     */
    @Export(name="dateUpdated", refs={String.class}, tree="[0]")
    private Output<String> dateUpdated;

    /**
     * @return Date of node pool updates.
     * 
     */
    public Output<String> dateUpdated() {
        return this.dateUpdated;
    }
    /**
     * The label to be used as a prefix for nodes in this node pool.
     * 
     */
    @Export(name="label", refs={String.class}, tree="[0]")
    private Output<String> label;

    /**
     * @return The label to be used as a prefix for nodes in this node pool.
     * 
     */
    public Output<String> label() {
        return this.label;
    }
    /**
     * The maximum number of nodes to use with the auto scaler.
     * 
     */
    @Export(name="maxNodes", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxNodes;

    /**
     * @return The maximum number of nodes to use with the auto scaler.
     * 
     */
    public Output<Optional<Integer>> maxNodes() {
        return Codegen.optional(this.maxNodes);
    }
    /**
     * The minimum number of nodes to use with the auto scaler.
     * 
     */
    @Export(name="minNodes", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> minNodes;

    /**
     * @return The minimum number of nodes to use with the auto scaler.
     * 
     */
    public Output<Optional<Integer>> minNodes() {
        return Codegen.optional(this.minNodes);
    }
    /**
     * The number of nodes in this node pool.
     * 
     */
    @Export(name="nodeQuantity", refs={Integer.class}, tree="[0]")
    private Output<Integer> nodeQuantity;

    /**
     * @return The number of nodes in this node pool.
     * 
     */
    public Output<Integer> nodeQuantity() {
        return this.nodeQuantity;
    }
    /**
     * Array that contains information about nodes within this node pool.
     * 
     */
    @Export(name="nodes", refs={List.class,KubernetesNodePoolsNode.class}, tree="[0,1]")
    private Output<List<KubernetesNodePoolsNode>> nodes;

    /**
     * @return Array that contains information about nodes within this node pool.
     * 
     */
    public Output<List<KubernetesNodePoolsNode>> nodes() {
        return this.nodes;
    }
    /**
     * The plan to be used in this node pool. [See Plans List](https://www.vultr.com/api/#operation/list-plans) Note the minimum plan requirements must have at least 1 core and 2 gbs of memory.
     * 
     */
    @Export(name="plan", refs={String.class}, tree="[0]")
    private Output<String> plan;

    /**
     * @return The plan to be used in this node pool. [See Plans List](https://www.vultr.com/api/#operation/list-plans) Note the minimum plan requirements must have at least 1 core and 2 gbs of memory.
     * 
     */
    public Output<String> plan() {
        return this.plan;
    }
    /**
     * Status of node.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Status of node.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * A tag that is assigned to this node pool.
     * 
     */
    @Export(name="tag", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tag;

    /**
     * @return A tag that is assigned to this node pool.
     * 
     */
    public Output<Optional<String>> tag() {
        return Codegen.optional(this.tag);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public KubernetesNodePools(String name) {
        this(name, KubernetesNodePoolsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public KubernetesNodePools(String name, KubernetesNodePoolsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public KubernetesNodePools(String name, KubernetesNodePoolsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vultr:index/kubernetesNodePools:KubernetesNodePools", name, args == null ? KubernetesNodePoolsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private KubernetesNodePools(String name, Output<String> id, @Nullable KubernetesNodePoolsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vultr:index/kubernetesNodePools:KubernetesNodePools", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static KubernetesNodePools get(String name, Output<String> id, @Nullable KubernetesNodePoolsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new KubernetesNodePools(name, id, state, options);
    }
}
