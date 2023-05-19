// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DatabaseUserArgs extends com.pulumi.resources.ResourceArgs {

    public static final DatabaseUserArgs Empty = new DatabaseUserArgs();

    @Import(name="databaseId", required=true)
    private Output<String> databaseId;

    public Output<String> databaseId() {
        return this.databaseId;
    }

    @Import(name="encryption")
    private @Nullable Output<String> encryption;

    public Optional<Output<String>> encryption() {
        return Optional.ofNullable(this.encryption);
    }

    @Import(name="password")
    private @Nullable Output<String> password;

    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    @Import(name="username", required=true)
    private Output<String> username;

    public Output<String> username() {
        return this.username;
    }

    private DatabaseUserArgs() {}

    private DatabaseUserArgs(DatabaseUserArgs $) {
        this.databaseId = $.databaseId;
        this.encryption = $.encryption;
        this.password = $.password;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DatabaseUserArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DatabaseUserArgs $;

        public Builder() {
            $ = new DatabaseUserArgs();
        }

        public Builder(DatabaseUserArgs defaults) {
            $ = new DatabaseUserArgs(Objects.requireNonNull(defaults));
        }

        public Builder databaseId(Output<String> databaseId) {
            $.databaseId = databaseId;
            return this;
        }

        public Builder databaseId(String databaseId) {
            return databaseId(Output.of(databaseId));
        }

        public Builder encryption(@Nullable Output<String> encryption) {
            $.encryption = encryption;
            return this;
        }

        public Builder encryption(String encryption) {
            return encryption(Output.of(encryption));
        }

        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        public Builder password(String password) {
            return password(Output.of(password));
        }

        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        public Builder username(String username) {
            return username(Output.of(username));
        }

        public DatabaseUserArgs build() {
            $.databaseId = Objects.requireNonNull($.databaseId, "expected parameter 'databaseId' to be non-null");
            $.username = Objects.requireNonNull($.username, "expected parameter 'username' to be non-null");
            return $;
        }
    }

}