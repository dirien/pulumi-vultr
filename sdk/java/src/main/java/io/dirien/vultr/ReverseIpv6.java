// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import io.dirien.vultr.ReverseIpv6Args;
import io.dirien.vultr.Utilities;
import io.dirien.vultr.inputs.ReverseIpv6State;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Vultr Reverse IPv6 resource. This can be used to create, read,
 * modify, and delete reverse DNS records for IPv6 addresses. Upon success, DNS
 * changes may take 6-12 hours to become active.
 * 
 */
@ResourceType(type="vultr:index/reverseIpv6:ReverseIpv6")
public class ReverseIpv6 extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the server you want to set an IPv6
     * reverse DNS record for.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return The ID of the server you want to set an IPv6
     * reverse DNS record for.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The IPv6 address used in the reverse DNS record.
     * 
     */
    @Export(name="ip", refs={String.class}, tree="[0]")
    private Output<String> ip;

    /**
     * @return The IPv6 address used in the reverse DNS record.
     * 
     */
    public Output<String> ip() {
        return this.ip;
    }
    /**
     * The hostname used in the IPv6 reverse DNS record.
     * 
     */
    @Export(name="reverse", refs={String.class}, tree="[0]")
    private Output<String> reverse;

    /**
     * @return The hostname used in the IPv6 reverse DNS record.
     * 
     */
    public Output<String> reverse() {
        return this.reverse;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ReverseIpv6(String name) {
        this(name, ReverseIpv6Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ReverseIpv6(String name, ReverseIpv6Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ReverseIpv6(String name, ReverseIpv6Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vultr:index/reverseIpv6:ReverseIpv6", name, args == null ? ReverseIpv6Args.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ReverseIpv6(String name, Output<String> id, @Nullable ReverseIpv6State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vultr:index/reverseIpv6:ReverseIpv6", name, state, makeResourceOptions(options, id));
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
    public static ReverseIpv6 get(String name, Output<String> id, @Nullable ReverseIpv6State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ReverseIpv6(name, id, state, options);
    }
}
