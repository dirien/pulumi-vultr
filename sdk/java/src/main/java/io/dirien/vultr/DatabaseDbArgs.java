// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DatabaseDbArgs extends com.pulumi.resources.ResourceArgs {

    public static final DatabaseDbArgs Empty = new DatabaseDbArgs();

    /**
     * The managed database ID you want to attach this logical DB to.
     * 
     */
    @Import(name="databaseId", required=true)
    private Output<String> databaseId;

    /**
     * @return The managed database ID you want to attach this logical DB to.
     * 
     */
    public Output<String> databaseId() {
        return this.databaseId;
    }

    /**
     * The name of the new managed database logical DB.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the new managed database logical DB.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private DatabaseDbArgs() {}

    private DatabaseDbArgs(DatabaseDbArgs $) {
        this.databaseId = $.databaseId;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DatabaseDbArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DatabaseDbArgs $;

        public Builder() {
            $ = new DatabaseDbArgs();
        }

        public Builder(DatabaseDbArgs defaults) {
            $ = new DatabaseDbArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param databaseId The managed database ID you want to attach this logical DB to.
         * 
         * @return builder
         * 
         */
        public Builder databaseId(Output<String> databaseId) {
            $.databaseId = databaseId;
            return this;
        }

        /**
         * @param databaseId The managed database ID you want to attach this logical DB to.
         * 
         * @return builder
         * 
         */
        public Builder databaseId(String databaseId) {
            return databaseId(Output.of(databaseId));
        }

        /**
         * @param name The name of the new managed database logical DB.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the new managed database logical DB.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public DatabaseDbArgs build() {
            $.databaseId = Objects.requireNonNull($.databaseId, "expected parameter 'databaseId' to be non-null");
            return $;
        }
    }

}
