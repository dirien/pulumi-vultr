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


public final class LoadBalancerHealthCheckArgs extends com.pulumi.resources.ResourceArgs {

    public static final LoadBalancerHealthCheckArgs Empty = new LoadBalancerHealthCheckArgs();

    /**
     * Time in seconds to perform health check. Default value is 15.
     * 
     */
    @Import(name="checkInterval", required=true)
    private Output<Integer> checkInterval;

    /**
     * @return Time in seconds to perform health check. Default value is 15.
     * 
     */
    public Output<Integer> checkInterval() {
        return this.checkInterval;
    }

    /**
     * Number of failed attempts encountered before failover. Default value is 5.
     * 
     */
    @Import(name="healthyThreshold", required=true)
    private Output<Integer> healthyThreshold;

    /**
     * @return Number of failed attempts encountered before failover. Default value is 5.
     * 
     */
    public Output<Integer> healthyThreshold() {
        return this.healthyThreshold;
    }

    /**
     * The path on the attached instances that the load balancer should check against. Default value is `/`
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return The path on the attached instances that the load balancer should check against. Default value is `/`
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * The assigned port (integer) on the attached instances that the load balancer should check against. Default value is `80`.
     * 
     */
    @Import(name="port", required=true)
    private Output<Integer> port;

    /**
     * @return The assigned port (integer) on the attached instances that the load balancer should check against. Default value is `80`.
     * 
     */
    public Output<Integer> port() {
        return this.port;
    }

    /**
     * The protocol used to traffic requests to the load balancer. Possible values are `http`, or `tcp`. Default value is `http`.
     * 
     */
    @Import(name="protocol", required=true)
    private Output<String> protocol;

    /**
     * @return The protocol used to traffic requests to the load balancer. Possible values are `http`, or `tcp`. Default value is `http`.
     * 
     */
    public Output<String> protocol() {
        return this.protocol;
    }

    /**
     * Time in seconds to wait for a health check response. Default value is 5.
     * 
     */
    @Import(name="responseTimeout", required=true)
    private Output<Integer> responseTimeout;

    /**
     * @return Time in seconds to wait for a health check response. Default value is 5.
     * 
     */
    public Output<Integer> responseTimeout() {
        return this.responseTimeout;
    }

    /**
     * Number of failed attempts encountered before failover. Default value is 5.
     * 
     */
    @Import(name="unhealthyThreshold", required=true)
    private Output<Integer> unhealthyThreshold;

    /**
     * @return Number of failed attempts encountered before failover. Default value is 5.
     * 
     */
    public Output<Integer> unhealthyThreshold() {
        return this.unhealthyThreshold;
    }

    private LoadBalancerHealthCheckArgs() {}

    private LoadBalancerHealthCheckArgs(LoadBalancerHealthCheckArgs $) {
        this.checkInterval = $.checkInterval;
        this.healthyThreshold = $.healthyThreshold;
        this.path = $.path;
        this.port = $.port;
        this.protocol = $.protocol;
        this.responseTimeout = $.responseTimeout;
        this.unhealthyThreshold = $.unhealthyThreshold;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LoadBalancerHealthCheckArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LoadBalancerHealthCheckArgs $;

        public Builder() {
            $ = new LoadBalancerHealthCheckArgs();
        }

        public Builder(LoadBalancerHealthCheckArgs defaults) {
            $ = new LoadBalancerHealthCheckArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param checkInterval Time in seconds to perform health check. Default value is 15.
         * 
         * @return builder
         * 
         */
        public Builder checkInterval(Output<Integer> checkInterval) {
            $.checkInterval = checkInterval;
            return this;
        }

        /**
         * @param checkInterval Time in seconds to perform health check. Default value is 15.
         * 
         * @return builder
         * 
         */
        public Builder checkInterval(Integer checkInterval) {
            return checkInterval(Output.of(checkInterval));
        }

        /**
         * @param healthyThreshold Number of failed attempts encountered before failover. Default value is 5.
         * 
         * @return builder
         * 
         */
        public Builder healthyThreshold(Output<Integer> healthyThreshold) {
            $.healthyThreshold = healthyThreshold;
            return this;
        }

        /**
         * @param healthyThreshold Number of failed attempts encountered before failover. Default value is 5.
         * 
         * @return builder
         * 
         */
        public Builder healthyThreshold(Integer healthyThreshold) {
            return healthyThreshold(Output.of(healthyThreshold));
        }

        /**
         * @param path The path on the attached instances that the load balancer should check against. Default value is `/`
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The path on the attached instances that the load balancer should check against. Default value is `/`
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param port The assigned port (integer) on the attached instances that the load balancer should check against. Default value is `80`.
         * 
         * @return builder
         * 
         */
        public Builder port(Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The assigned port (integer) on the attached instances that the load balancer should check against. Default value is `80`.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param protocol The protocol used to traffic requests to the load balancer. Possible values are `http`, or `tcp`. Default value is `http`.
         * 
         * @return builder
         * 
         */
        public Builder protocol(Output<String> protocol) {
            $.protocol = protocol;
            return this;
        }

        /**
         * @param protocol The protocol used to traffic requests to the load balancer. Possible values are `http`, or `tcp`. Default value is `http`.
         * 
         * @return builder
         * 
         */
        public Builder protocol(String protocol) {
            return protocol(Output.of(protocol));
        }

        /**
         * @param responseTimeout Time in seconds to wait for a health check response. Default value is 5.
         * 
         * @return builder
         * 
         */
        public Builder responseTimeout(Output<Integer> responseTimeout) {
            $.responseTimeout = responseTimeout;
            return this;
        }

        /**
         * @param responseTimeout Time in seconds to wait for a health check response. Default value is 5.
         * 
         * @return builder
         * 
         */
        public Builder responseTimeout(Integer responseTimeout) {
            return responseTimeout(Output.of(responseTimeout));
        }

        /**
         * @param unhealthyThreshold Number of failed attempts encountered before failover. Default value is 5.
         * 
         * @return builder
         * 
         */
        public Builder unhealthyThreshold(Output<Integer> unhealthyThreshold) {
            $.unhealthyThreshold = unhealthyThreshold;
            return this;
        }

        /**
         * @param unhealthyThreshold Number of failed attempts encountered before failover. Default value is 5.
         * 
         * @return builder
         * 
         */
        public Builder unhealthyThreshold(Integer unhealthyThreshold) {
            return unhealthyThreshold(Output.of(unhealthyThreshold));
        }

        public LoadBalancerHealthCheckArgs build() {
            $.checkInterval = Objects.requireNonNull($.checkInterval, "expected parameter 'checkInterval' to be non-null");
            $.healthyThreshold = Objects.requireNonNull($.healthyThreshold, "expected parameter 'healthyThreshold' to be non-null");
            $.port = Objects.requireNonNull($.port, "expected parameter 'port' to be non-null");
            $.protocol = Objects.requireNonNull($.protocol, "expected parameter 'protocol' to be non-null");
            $.responseTimeout = Objects.requireNonNull($.responseTimeout, "expected parameter 'responseTimeout' to be non-null");
            $.unhealthyThreshold = Objects.requireNonNull($.unhealthyThreshold, "expected parameter 'unhealthyThreshold' to be non-null");
            return $;
        }
    }

}
