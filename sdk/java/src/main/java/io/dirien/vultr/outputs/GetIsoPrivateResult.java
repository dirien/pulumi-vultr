// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.outputs;

import com.pulumi.core.annotations.CustomType;
import io.dirien.vultr.outputs.GetIsoPrivateFilter;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetIsoPrivateResult {
    /**
     * @return The date the ISO file was added to your Vultr account.
     * 
     */
    private String dateCreated;
    /**
     * @return The ISO file&#39;s filename.
     * 
     */
    private String filename;
    private @Nullable List<GetIsoPrivateFilter> filters;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The md5 hash of the ISO file.
     * 
     */
    private String md5sum;
    /**
     * @return The sha512 hash of the ISO file.
     * 
     */
    private String sha512sum;
    /**
     * @return The size of the ISO file in bytes.
     * 
     */
    private Integer size;
    /**
     * @return The status of the ISO file.
     * 
     */
    private String status;

    private GetIsoPrivateResult() {}
    /**
     * @return The date the ISO file was added to your Vultr account.
     * 
     */
    public String dateCreated() {
        return this.dateCreated;
    }
    /**
     * @return The ISO file&#39;s filename.
     * 
     */
    public String filename() {
        return this.filename;
    }
    public List<GetIsoPrivateFilter> filters() {
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
     * @return The md5 hash of the ISO file.
     * 
     */
    public String md5sum() {
        return this.md5sum;
    }
    /**
     * @return The sha512 hash of the ISO file.
     * 
     */
    public String sha512sum() {
        return this.sha512sum;
    }
    /**
     * @return The size of the ISO file in bytes.
     * 
     */
    public Integer size() {
        return this.size;
    }
    /**
     * @return The status of the ISO file.
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIsoPrivateResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String dateCreated;
        private String filename;
        private @Nullable List<GetIsoPrivateFilter> filters;
        private String id;
        private String md5sum;
        private String sha512sum;
        private Integer size;
        private String status;
        public Builder() {}
        public Builder(GetIsoPrivateResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dateCreated = defaults.dateCreated;
    	      this.filename = defaults.filename;
    	      this.filters = defaults.filters;
    	      this.id = defaults.id;
    	      this.md5sum = defaults.md5sum;
    	      this.sha512sum = defaults.sha512sum;
    	      this.size = defaults.size;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder dateCreated(String dateCreated) {
            this.dateCreated = Objects.requireNonNull(dateCreated);
            return this;
        }
        @CustomType.Setter
        public Builder filename(String filename) {
            this.filename = Objects.requireNonNull(filename);
            return this;
        }
        @CustomType.Setter
        public Builder filters(@Nullable List<GetIsoPrivateFilter> filters) {
            this.filters = filters;
            return this;
        }
        public Builder filters(GetIsoPrivateFilter... filters) {
            return filters(List.of(filters));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder md5sum(String md5sum) {
            this.md5sum = Objects.requireNonNull(md5sum);
            return this;
        }
        @CustomType.Setter
        public Builder sha512sum(String sha512sum) {
            this.sha512sum = Objects.requireNonNull(sha512sum);
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
        public GetIsoPrivateResult build() {
            final var o = new GetIsoPrivateResult();
            o.dateCreated = dateCreated;
            o.filename = filename;
            o.filters = filters;
            o.id = id;
            o.md5sum = md5sum;
            o.sha512sum = sha512sum;
            o.size = size;
            o.status = status;
            return o;
        }
    }
}
