// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SnapshotFromUrlState extends com.pulumi.resources.ResourceArgs {

    public static final SnapshotFromUrlState Empty = new SnapshotFromUrlState();

    /**
     * The app id which the snapshot is associated with.
     * 
     */
    @Import(name="appId")
    private @Nullable Output<Integer> appId;

    /**
     * @return The app id which the snapshot is associated with.
     * 
     */
    public Optional<Output<Integer>> appId() {
        return Optional.ofNullable(this.appId);
    }

    /**
     * The date the snapshot was created.
     * 
     */
    @Import(name="dateCreated")
    private @Nullable Output<String> dateCreated;

    /**
     * @return The date the snapshot was created.
     * 
     */
    public Optional<Output<String>> dateCreated() {
        return Optional.ofNullable(this.dateCreated);
    }

    /**
     * The description for the given snapshot.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description for the given snapshot.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The os id which the snapshot is associated with.
     * 
     */
    @Import(name="osId")
    private @Nullable Output<Integer> osId;

    /**
     * @return The os id which the snapshot is associated with.
     * 
     */
    public Optional<Output<Integer>> osId() {
        return Optional.ofNullable(this.osId);
    }

    /**
     * The size of the snapshot in Bytes.
     * 
     */
    @Import(name="size")
    private @Nullable Output<Integer> size;

    /**
     * @return The size of the snapshot in Bytes.
     * 
     */
    public Optional<Output<Integer>> size() {
        return Optional.ofNullable(this.size);
    }

    /**
     * The status for the given snapshot.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status for the given snapshot.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * URL of the given resource you want to create a snapshot from.
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return URL of the given resource you want to create a snapshot from.
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    private SnapshotFromUrlState() {}

    private SnapshotFromUrlState(SnapshotFromUrlState $) {
        this.appId = $.appId;
        this.dateCreated = $.dateCreated;
        this.description = $.description;
        this.osId = $.osId;
        this.size = $.size;
        this.status = $.status;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SnapshotFromUrlState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SnapshotFromUrlState $;

        public Builder() {
            $ = new SnapshotFromUrlState();
        }

        public Builder(SnapshotFromUrlState defaults) {
            $ = new SnapshotFromUrlState(Objects.requireNonNull(defaults));
        }

        /**
         * @param appId The app id which the snapshot is associated with.
         * 
         * @return builder
         * 
         */
        public Builder appId(@Nullable Output<Integer> appId) {
            $.appId = appId;
            return this;
        }

        /**
         * @param appId The app id which the snapshot is associated with.
         * 
         * @return builder
         * 
         */
        public Builder appId(Integer appId) {
            return appId(Output.of(appId));
        }

        /**
         * @param dateCreated The date the snapshot was created.
         * 
         * @return builder
         * 
         */
        public Builder dateCreated(@Nullable Output<String> dateCreated) {
            $.dateCreated = dateCreated;
            return this;
        }

        /**
         * @param dateCreated The date the snapshot was created.
         * 
         * @return builder
         * 
         */
        public Builder dateCreated(String dateCreated) {
            return dateCreated(Output.of(dateCreated));
        }

        /**
         * @param description The description for the given snapshot.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description for the given snapshot.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param osId The os id which the snapshot is associated with.
         * 
         * @return builder
         * 
         */
        public Builder osId(@Nullable Output<Integer> osId) {
            $.osId = osId;
            return this;
        }

        /**
         * @param osId The os id which the snapshot is associated with.
         * 
         * @return builder
         * 
         */
        public Builder osId(Integer osId) {
            return osId(Output.of(osId));
        }

        /**
         * @param size The size of the snapshot in Bytes.
         * 
         * @return builder
         * 
         */
        public Builder size(@Nullable Output<Integer> size) {
            $.size = size;
            return this;
        }

        /**
         * @param size The size of the snapshot in Bytes.
         * 
         * @return builder
         * 
         */
        public Builder size(Integer size) {
            return size(Output.of(size));
        }

        /**
         * @param status The status for the given snapshot.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status for the given snapshot.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param url URL of the given resource you want to create a snapshot from.
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
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

        public SnapshotFromUrlState build() {
            return $;
        }
    }

}
