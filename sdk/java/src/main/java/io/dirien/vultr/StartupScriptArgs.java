// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class StartupScriptArgs extends com.pulumi.resources.ResourceArgs {

    public static final StartupScriptArgs Empty = new StartupScriptArgs();

    /**
     * Name of the given script.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the given script.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Contents of the startup script base64 encoded.
     * 
     */
    @Import(name="script", required=true)
    private Output<String> script;

    /**
     * @return Contents of the startup script base64 encoded.
     * 
     */
    public Output<String> script() {
        return this.script;
    }

    /**
     * Type of startup script. Possible values are boot or pxe - default is boot.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Type of startup script. Possible values are boot or pxe - default is boot.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private StartupScriptArgs() {}

    private StartupScriptArgs(StartupScriptArgs $) {
        this.name = $.name;
        this.script = $.script;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(StartupScriptArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private StartupScriptArgs $;

        public Builder() {
            $ = new StartupScriptArgs();
        }

        public Builder(StartupScriptArgs defaults) {
            $ = new StartupScriptArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Name of the given script.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the given script.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param script Contents of the startup script base64 encoded.
         * 
         * @return builder
         * 
         */
        public Builder script(Output<String> script) {
            $.script = script;
            return this;
        }

        /**
         * @param script Contents of the startup script base64 encoded.
         * 
         * @return builder
         * 
         */
        public Builder script(String script) {
            return script(Output.of(script));
        }

        /**
         * @param type Type of startup script. Possible values are boot or pxe - default is boot.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Type of startup script. Possible values are boot or pxe - default is boot.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public StartupScriptArgs build() {
            $.script = Objects.requireNonNull($.script, "expected parameter 'script' to be non-null");
            return $;
        }
    }

}
