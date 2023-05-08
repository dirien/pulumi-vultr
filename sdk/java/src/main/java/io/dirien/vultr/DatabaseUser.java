// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import io.dirien.vultr.DatabaseUserArgs;
import io.dirien.vultr.Utilities;
import io.dirien.vultr.inputs.DatabaseUserState;
import java.lang.String;
import javax.annotation.Nullable;

@ResourceType(type="vultr:index/databaseUser:DatabaseUser")
public class DatabaseUser extends com.pulumi.resources.CustomResource {
    @Export(name="databaseId", refs={String.class}, tree="[0]")
    private Output<String> databaseId;

    public Output<String> databaseId() {
        return this.databaseId;
    }
    @Export(name="encryption", refs={String.class}, tree="[0]")
    private Output<String> encryption;

    public Output<String> encryption() {
        return this.encryption;
    }
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output<String> password;

    public Output<String> password() {
        return this.password;
    }
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output<String> username;

    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DatabaseUser(String name) {
        this(name, DatabaseUserArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DatabaseUser(String name, DatabaseUserArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DatabaseUser(String name, DatabaseUserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vultr:index/databaseUser:DatabaseUser", name, args == null ? DatabaseUserArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DatabaseUser(String name, Output<String> id, @Nullable DatabaseUserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vultr:index/databaseUser:DatabaseUser", name, state, makeResourceOptions(options, id));
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
    public static DatabaseUser get(String name, Output<String> id, @Nullable DatabaseUserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DatabaseUser(name, id, state, options);
    }
}
