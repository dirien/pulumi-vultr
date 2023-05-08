// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DatabaseDbState extends com.pulumi.resources.ResourceArgs {

    public static final DatabaseDbState Empty = new DatabaseDbState();

    @Import(name="databaseId")
    private @Nullable Output<String> databaseId;

    public Optional<Output<String>> databaseId() {
        return Optional.ofNullable(this.databaseId);
    }

    @Import(name="name")
    private @Nullable Output<String> name;

    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private DatabaseDbState() {}

    private DatabaseDbState(DatabaseDbState $) {
        this.databaseId = $.databaseId;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DatabaseDbState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DatabaseDbState $;

        public Builder() {
            $ = new DatabaseDbState();
        }

        public Builder(DatabaseDbState defaults) {
            $ = new DatabaseDbState(Objects.requireNonNull(defaults));
        }

        public Builder databaseId(@Nullable Output<String> databaseId) {
            $.databaseId = databaseId;
            return this;
        }

        public Builder databaseId(String databaseId) {
            return databaseId(Output.of(databaseId));
        }

        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public DatabaseDbState build() {
            return $;
        }
    }

}
