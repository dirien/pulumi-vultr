// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UserArgs extends com.pulumi.resources.ResourceArgs {

    public static final UserArgs Empty = new UserArgs();

    /**
     * The access control list for the user.
     * 
     */
    @Import(name="acls")
    private @Nullable Output<List<String>> acls;

    /**
     * @return The access control list for the user.
     * 
     */
    public Optional<Output<List<String>>> acls() {
        return Optional.ofNullable(this.acls);
    }

    /**
     * Whether API is enabled for the user. Default behavior is set to enabled.
     * 
     */
    @Import(name="apiEnabled")
    private @Nullable Output<Boolean> apiEnabled;

    /**
     * @return Whether API is enabled for the user. Default behavior is set to enabled.
     * 
     */
    public Optional<Output<Boolean>> apiEnabled() {
        return Optional.ofNullable(this.apiEnabled);
    }

    /**
     * Email for this user.
     * 
     */
    @Import(name="email", required=true)
    private Output<String> email;

    /**
     * @return Email for this user.
     * 
     */
    public Output<String> email() {
        return this.email;
    }

    /**
     * Name for this user.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name for this user.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Password for this user.
     * 
     */
    @Import(name="password", required=true)
    private Output<String> password;

    /**
     * @return Password for this user.
     * 
     */
    public Output<String> password() {
        return this.password;
    }

    private UserArgs() {}

    private UserArgs(UserArgs $) {
        this.acls = $.acls;
        this.apiEnabled = $.apiEnabled;
        this.email = $.email;
        this.name = $.name;
        this.password = $.password;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserArgs $;

        public Builder() {
            $ = new UserArgs();
        }

        public Builder(UserArgs defaults) {
            $ = new UserArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param acls The access control list for the user.
         * 
         * @return builder
         * 
         */
        public Builder acls(@Nullable Output<List<String>> acls) {
            $.acls = acls;
            return this;
        }

        /**
         * @param acls The access control list for the user.
         * 
         * @return builder
         * 
         */
        public Builder acls(List<String> acls) {
            return acls(Output.of(acls));
        }

        /**
         * @param acls The access control list for the user.
         * 
         * @return builder
         * 
         */
        public Builder acls(String... acls) {
            return acls(List.of(acls));
        }

        /**
         * @param apiEnabled Whether API is enabled for the user. Default behavior is set to enabled.
         * 
         * @return builder
         * 
         */
        public Builder apiEnabled(@Nullable Output<Boolean> apiEnabled) {
            $.apiEnabled = apiEnabled;
            return this;
        }

        /**
         * @param apiEnabled Whether API is enabled for the user. Default behavior is set to enabled.
         * 
         * @return builder
         * 
         */
        public Builder apiEnabled(Boolean apiEnabled) {
            return apiEnabled(Output.of(apiEnabled));
        }

        /**
         * @param email Email for this user.
         * 
         * @return builder
         * 
         */
        public Builder email(Output<String> email) {
            $.email = email;
            return this;
        }

        /**
         * @param email Email for this user.
         * 
         * @return builder
         * 
         */
        public Builder email(String email) {
            return email(Output.of(email));
        }

        /**
         * @param name Name for this user.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name for this user.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param password Password for this user.
         * 
         * @return builder
         * 
         */
        public Builder password(Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password Password for this user.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        public UserArgs build() {
            $.email = Objects.requireNonNull($.email, "expected parameter 'email' to be non-null");
            $.password = Objects.requireNonNull($.password, "expected parameter 'password' to be non-null");
            return $;
        }
    }

}