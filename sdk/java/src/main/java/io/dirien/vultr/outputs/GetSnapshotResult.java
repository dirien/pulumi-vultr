// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.outputs;

import com.pulumi.core.annotations.CustomType;
import io.dirien.vultr.outputs.GetSnapshotFilter;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetSnapshotResult {
    /**
     * @return The application ID of the snapshot.
     * 
     */
    private Integer appId;
    /**
     * @return The date the snapshot was added to your Vultr account.
     * 
     */
    private String dateCreated;
    /**
     * @return The description of the snapshot.
     * 
     */
    private String description;
    private @Nullable List<GetSnapshotFilter> filters;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The operating system ID of the snapshot.
     * 
     */
    private Integer osId;
    /**
     * @return The size of the snapshot in bytes.
     * 
     */
    private Integer size;
    /**
     * @return The status of the snapshot.
     * 
     */
    private String status;

    private GetSnapshotResult() {}
    /**
     * @return The application ID of the snapshot.
     * 
     */
    public Integer appId() {
        return this.appId;
    }
    /**
     * @return The date the snapshot was added to your Vultr account.
     * 
     */
    public String dateCreated() {
        return this.dateCreated;
    }
    /**
     * @return The description of the snapshot.
     * 
     */
    public String description() {
        return this.description;
    }
    public List<GetSnapshotFilter> filters() {
        return this.filters == null ? List.of() : this.filters;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The operating system ID of the snapshot.
     * 
     */
    public Integer osId() {
        return this.osId;
    }
    /**
     * @return The size of the snapshot in bytes.
     * 
     */
    public Integer size() {
        return this.size;
    }
    /**
     * @return The status of the snapshot.
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSnapshotResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer appId;
        private String dateCreated;
        private String description;
        private @Nullable List<GetSnapshotFilter> filters;
        private String id;
        private Integer osId;
        private Integer size;
        private String status;
        public Builder() {}
        public Builder(GetSnapshotResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.appId = defaults.appId;
    	      this.dateCreated = defaults.dateCreated;
    	      this.description = defaults.description;
    	      this.filters = defaults.filters;
    	      this.id = defaults.id;
    	      this.osId = defaults.osId;
    	      this.size = defaults.size;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder appId(Integer appId) {
            this.appId = Objects.requireNonNull(appId);
            return this;
        }
        @CustomType.Setter
        public Builder dateCreated(String dateCreated) {
            this.dateCreated = Objects.requireNonNull(dateCreated);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder filters(@Nullable List<GetSnapshotFilter> filters) {
            this.filters = filters;
            return this;
        }
        public Builder filters(GetSnapshotFilter... filters) {
            return filters(List.of(filters));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder osId(Integer osId) {
            this.osId = Objects.requireNonNull(osId);
            return this;
        }
        @CustomType.Setter
        public Builder size(Integer size) {
            this.size = Objects.requireNonNull(size);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        public GetSnapshotResult build() {
            final var o = new GetSnapshotResult();
            o.appId = appId;
            o.dateCreated = dateCreated;
            o.description = description;
            o.filters = filters;
            o.id = id;
            o.osId = osId;
            o.size = size;
            o.status = status;
            return o;
        }
    }
}
