// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class SnapshotFromUrlArgs extends com.pulumi.resources.ResourceArgs {

    public static final SnapshotFromUrlArgs Empty = new SnapshotFromUrlArgs();

    /**
     * URL of the given resource you want to create a snapshot from.
     * 
     */
    @Import(name="url", required=true)
    private Output<String> url;

    /**
     * @return URL of the given resource you want to create a snapshot from.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    private SnapshotFromUrlArgs() {}

    private SnapshotFromUrlArgs(SnapshotFromUrlArgs $) {
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SnapshotFromUrlArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SnapshotFromUrlArgs $;

        public Builder() {
            $ = new SnapshotFromUrlArgs();
        }

        public Builder(SnapshotFromUrlArgs defaults) {
            $ = new SnapshotFromUrlArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param url URL of the given resource you want to create a snapshot from.
         * 
         * @return builder
         * 
         */
        public Builder url(Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url URL of the given resource you want to create a snapshot from.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public SnapshotFromUrlArgs build() {
            $.url = Objects.requireNonNull($.url, "expected parameter 'url' to be non-null");
            return $;
        }
    }

}
