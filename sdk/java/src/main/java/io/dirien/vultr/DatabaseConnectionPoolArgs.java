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


public final class DatabaseConnectionPoolArgs extends com.pulumi.resources.ResourceArgs {

    public static final DatabaseConnectionPoolArgs Empty = new DatabaseConnectionPoolArgs();

    @Import(name="database", required=true)
    private Output<String> database;

    public Output<String> database() {
        return this.database;
    }

    @Import(name="databaseId", required=true)
    private Output<String> databaseId;

    public Output<String> databaseId() {
        return this.databaseId;
    }

    @Import(name="mode", required=true)
    private Output<String> mode;

    public Output<String> mode() {
        return this.mode;
    }

    @Import(name="name")
    private @Nullable Output<String> name;

    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="size", required=true)
    private Output<Integer> size;

    public Output<Integer> size() {
        return this.size;
    }

    @Import(name="username", required=true)
    private Output<String> username;

    public Output<String> username() {
        return this.username;
    }

    private DatabaseConnectionPoolArgs() {}

    private DatabaseConnectionPoolArgs(DatabaseConnectionPoolArgs $) {
        this.database = $.database;
        this.databaseId = $.databaseId;
        this.mode = $.mode;
        this.name = $.name;
        this.size = $.size;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DatabaseConnectionPoolArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DatabaseConnectionPoolArgs $;

        public Builder() {
            $ = new DatabaseConnectionPoolArgs();
        }

        public Builder(DatabaseConnectionPoolArgs defaults) {
            $ = new DatabaseConnectionPoolArgs(Objects.requireNonNull(defaults));
        }

        public Builder database(Output<String> database) {
            $.database = database;
            return this;
        }

        public Builder database(String database) {
            return database(Output.of(database));
        }

        public Builder databaseId(Output<String> databaseId) {
            $.databaseId = databaseId;
            return this;
        }

        public Builder databaseId(String databaseId) {
            return databaseId(Output.of(databaseId));
        }

        public Builder mode(Output<String> mode) {
            $.mode = mode;
            return this;
        }

        public Builder mode(String mode) {
            return mode(Output.of(mode));
        }

        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder size(Output<Integer> size) {
            $.size = size;
            return this;
        }

        public Builder size(Integer size) {
            return size(Output.of(size));
        }

        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        public Builder username(String username) {
            return username(Output.of(username));
        }

        public DatabaseConnectionPoolArgs build() {
            $.database = Objects.requireNonNull($.database, "expected parameter 'database' to be non-null");
            $.databaseId = Objects.requireNonNull($.databaseId, "expected parameter 'databaseId' to be non-null");
            $.mode = Objects.requireNonNull($.mode, "expected parameter 'mode' to be non-null");
            $.size = Objects.requireNonNull($.size, "expected parameter 'size' to be non-null");
            $.username = Objects.requireNonNull($.username, "expected parameter 'username' to be non-null");
            return $;
        }
    }

}