// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.outputs;

import com.pulumi.core.annotations.CustomType;
import io.dirien.vultr.outputs.GetUserFilter;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetUserResult {
    /**
     * @return The access control list for the user.
     * 
     */
    private List<String> acls;
    /**
     * @return Whether API is enabled for the user.
     * 
     */
    private Boolean apiEnabled;
    /**
     * @return The email of the user.
     * 
     */
    private String email;
    private @Nullable List<GetUserFilter> filters;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The name of the user.
     * 
     */
    private String name;

    private GetUserResult() {}
    /**
     * @return The access control list for the user.
     * 
     */
    public List<String> acls() {
        return this.acls;
    }
    /**
     * @return Whether API is enabled for the user.
     * 
     */
    public Boolean apiEnabled() {
        return this.apiEnabled;
    }
    /**
     * @return The email of the user.
     * 
     */
    public String email() {
        return this.email;
    }
    public List<GetUserFilter> filters() {
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
     * @return The name of the user.
     * 
     */
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetUserResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> acls;
        private Boolean apiEnabled;
        private String email;
        private @Nullable List<GetUserFilter> filters;
        private String id;
        private String name;
        public Builder() {}
        public Builder(GetUserResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.acls = defaults.acls;
    	      this.apiEnabled = defaults.apiEnabled;
    	      this.email = defaults.email;
    	      this.filters = defaults.filters;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder acls(List<String> acls) {
            this.acls = Objects.requireNonNull(acls);
            return this;
        }
        public Builder acls(String... acls) {
            return acls(List.of(acls));
        }
        @CustomType.Setter
        public Builder apiEnabled(Boolean apiEnabled) {
            this.apiEnabled = Objects.requireNonNull(apiEnabled);
            return this;
        }
        @CustomType.Setter
        public Builder email(String email) {
            this.email = Objects.requireNonNull(email);
            return this;
        }
        @CustomType.Setter
        public Builder filters(@Nullable List<GetUserFilter> filters) {
            this.filters = filters;
            return this;
        }
        public Builder filters(GetUserFilter... filters) {
            return filters(List.of(filters));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public GetUserResult build() {
            final var o = new GetUserResult();
            o.acls = acls;
            o.apiEnabled = apiEnabled;
            o.email = email;
            o.filters = filters;
            o.id = id;
            o.name = name;
            return o;
        }
    }
}
