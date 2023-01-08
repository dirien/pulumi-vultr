// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SSHKeyState extends com.pulumi.resources.ResourceArgs {

    public static final SSHKeyState Empty = new SSHKeyState();

    /**
     * The date the SSH key was added to your Vultr account.
     * 
     */
    @Import(name="dateCreated")
    private @Nullable Output<String> dateCreated;

    /**
     * @return The date the SSH key was added to your Vultr account.
     * 
     */
    public Optional<Output<String>> dateCreated() {
        return Optional.ofNullable(this.dateCreated);
    }

    /**
     * The name/label of the SSH key.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name/label of the SSH key.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The public SSH key.
     * 
     */
    @Import(name="sshKey")
    private @Nullable Output<String> sshKey;

    /**
     * @return The public SSH key.
     * 
     */
    public Optional<Output<String>> sshKey() {
        return Optional.ofNullable(this.sshKey);
    }

    private SSHKeyState() {}

    private SSHKeyState(SSHKeyState $) {
        this.dateCreated = $.dateCreated;
        this.name = $.name;
        this.sshKey = $.sshKey;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SSHKeyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SSHKeyState $;

        public Builder() {
            $ = new SSHKeyState();
        }

        public Builder(SSHKeyState defaults) {
            $ = new SSHKeyState(Objects.requireNonNull(defaults));
        }

        /**
         * @param dateCreated The date the SSH key was added to your Vultr account.
         * 
         * @return builder
         * 
         */
        public Builder dateCreated(@Nullable Output<String> dateCreated) {
            $.dateCreated = dateCreated;
            return this;
        }

        /**
         * @param dateCreated The date the SSH key was added to your Vultr account.
         * 
         * @return builder
         * 
         */
        public Builder dateCreated(String dateCreated) {
            return dateCreated(Output.of(dateCreated));
        }

        /**
         * @param name The name/label of the SSH key.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name/label of the SSH key.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param sshKey The public SSH key.
         * 
         * @return builder
         * 
         */
        public Builder sshKey(@Nullable Output<String> sshKey) {
            $.sshKey = sshKey;
            return this;
        }

        /**
         * @param sshKey The public SSH key.
         * 
         * @return builder
         * 
         */
        public Builder sshKey(String sshKey) {
            return sshKey(Output.of(sshKey));
        }

        public SSHKeyState build() {
            return $;
        }
    }

}