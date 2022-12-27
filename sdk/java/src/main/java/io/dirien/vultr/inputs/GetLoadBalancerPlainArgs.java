// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.inputs;

import com.pulumi.core.annotations.Import;
import io.dirien.vultr.inputs.GetLoadBalancerFilter;
import java.lang.Boolean;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetLoadBalancerPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetLoadBalancerPlainArgs Empty = new GetLoadBalancerPlainArgs();

    /**
     * Query parameters for finding load balancers.
     * 
     */
    @Import(name="filters")
    private @Nullable List<GetLoadBalancerFilter> filters;

    /**
     * @return Query parameters for finding load balancers.
     * 
     */
    public Optional<List<GetLoadBalancerFilter>> filters() {
        return Optional.ofNullable(this.filters);
    }

    /**
     * Boolean value that indicates if Proxy Protocol is enabled.
     * 
     */
    @Import(name="proxyProtocol")
    private @Nullable Boolean proxyProtocol;

    /**
     * @return Boolean value that indicates if Proxy Protocol is enabled.
     * 
     */
    public Optional<Boolean> proxyProtocol() {
        return Optional.ofNullable(this.proxyProtocol);
    }

    private GetLoadBalancerPlainArgs() {}

    private GetLoadBalancerPlainArgs(GetLoadBalancerPlainArgs $) {
        this.filters = $.filters;
        this.proxyProtocol = $.proxyProtocol;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetLoadBalancerPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetLoadBalancerPlainArgs $;

        public Builder() {
            $ = new GetLoadBalancerPlainArgs();
        }

        public Builder(GetLoadBalancerPlainArgs defaults) {
            $ = new GetLoadBalancerPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filters Query parameters for finding load balancers.
         * 
         * @return builder
         * 
         */
        public Builder filters(@Nullable List<GetLoadBalancerFilter> filters) {
            $.filters = filters;
            return this;
        }

        /**
         * @param filters Query parameters for finding load balancers.
         * 
         * @return builder
         * 
         */
        public Builder filters(GetLoadBalancerFilter... filters) {
            return filters(List.of(filters));
        }

        /**
         * @param proxyProtocol Boolean value that indicates if Proxy Protocol is enabled.
         * 
         * @return builder
         * 
         */
        public Builder proxyProtocol(@Nullable Boolean proxyProtocol) {
            $.proxyProtocol = proxyProtocol;
            return this;
        }

        public GetLoadBalancerPlainArgs build() {
            return $;
        }
    }

}
