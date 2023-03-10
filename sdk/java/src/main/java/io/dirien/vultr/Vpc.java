// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import io.dirien.vultr.Utilities;
import io.dirien.vultr.VpcArgs;
import io.dirien.vultr.inputs.VpcState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Vultr VPC resource. This can be used to create, read, and delete VPCs on your Vultr account.
 * 
 * ## Example Usage
 * 
 * Create a new VPC with an automatically generated CIDR block:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vultr.Vpc;
 * import com.pulumi.vultr.VpcArgs;
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
 *         var myVpc = new Vpc(&#34;myVpc&#34;, VpcArgs.builder()        
 *             .description(&#34;my vpc&#34;)
 *             .region(&#34;ewr&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * Create a new VPC with a specified CIDR block:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vultr.Vpc;
 * import com.pulumi.vultr.VpcArgs;
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
 *         var myVpc = new Vpc(&#34;myVpc&#34;, VpcArgs.builder()        
 *             .description(&#34;my private vpc&#34;)
 *             .region(&#34;ewr&#34;)
 *             .v4Subnet(&#34;10.0.0.0&#34;)
 *             .v4SubnetMask(24)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * VPCs can be imported using the VPC `ID`, e.g.
 * 
 * ```sh
 *  $ pulumi import vultr:index/vpc:Vpc my_vpc 0e04f918-575e-41cb-86f6-d729b354a5a1
 * ```
 * 
 */
@ResourceType(type="vultr:index/vpc:Vpc")
public class Vpc extends com.pulumi.resources.CustomResource {
    /**
     * The date that the VPC was added to your Vultr account.
     * 
     */
    @Export(name="dateCreated", refs={String.class}, tree="[0]")
    private Output<String> dateCreated;

    /**
     * @return The date that the VPC was added to your Vultr account.
     * 
     */
    public Output<String> dateCreated() {
        return this.dateCreated;
    }
    /**
     * The description you want to give your VPC.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description you want to give your VPC.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The region ID that you want the VPC to be created in.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region ID that you want the VPC to be created in.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The IPv4 subnet to be used when attaching instances to this VPC.
     * 
     */
    @Export(name="v4Subnet", refs={String.class}, tree="[0]")
    private Output<String> v4Subnet;

    /**
     * @return The IPv4 subnet to be used when attaching instances to this VPC.
     * 
     */
    public Output<String> v4Subnet() {
        return this.v4Subnet;
    }
    /**
     * The number of bits for the netmask in CIDR notation. Example: 32
     * 
     */
    @Export(name="v4SubnetMask", refs={Integer.class}, tree="[0]")
    private Output<Integer> v4SubnetMask;

    /**
     * @return The number of bits for the netmask in CIDR notation. Example: 32
     * 
     */
    public Output<Integer> v4SubnetMask() {
        return this.v4SubnetMask;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Vpc(String name) {
        this(name, VpcArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Vpc(String name, VpcArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Vpc(String name, VpcArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vultr:index/vpc:Vpc", name, args == null ? VpcArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Vpc(String name, Output<String> id, @Nullable VpcState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vultr:index/vpc:Vpc", name, state, makeResourceOptions(options, id));
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
    public static Vpc get(String name, Output<String> id, @Nullable VpcState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Vpc(name, id, state, options);
    }
}
