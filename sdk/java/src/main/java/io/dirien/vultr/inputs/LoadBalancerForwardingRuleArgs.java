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


public final class LoadBalancerForwardingRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final LoadBalancerForwardingRuleArgs Empty = new LoadBalancerForwardingRuleArgs();

    /**
     * Port on instance side.
     * 
     */
    @Import(name="backendPort", required=true)
    private Output<Integer> backendPort;

    /**
     * @return Port on instance side.
     * 
     */
    public Output<Integer> backendPort() {
        return this.backendPort;
    }

    /**
     * Protocol on instance side. Possible values: &#34;http&#34;, &#34;https&#34;, &#34;tcp&#34;.
     * 
     */
    @Import(name="backendProtocol", required=true)
    private Output<String> backendProtocol;

    /**
     * @return Protocol on instance side. Possible values: &#34;http&#34;, &#34;https&#34;, &#34;tcp&#34;.
     * 
     */
    public Output<String> backendProtocol() {
        return this.backendProtocol;
    }

    /**
     * Port on load balancer side.
     * 
     */
    @Import(name="frontendPort", required=true)
    private Output<Integer> frontendPort;

    /**
     * @return Port on load balancer side.
     * 
     */
    public Output<Integer> frontendPort() {
        return this.frontendPort;
    }

    /**
     * Protocol on load balancer side. Possible values: &#34;http&#34;, &#34;https&#34;, &#34;tcp&#34;.
     * 
     */
    @Import(name="frontendProtocol", required=true)
    private Output<String> frontendProtocol;

    /**
     * @return Protocol on load balancer side. Possible values: &#34;http&#34;, &#34;https&#34;, &#34;tcp&#34;.
     * 
     */
    public Output<String> frontendProtocol() {
        return this.frontendProtocol;
    }

    @Import(name="ruleId")
    private @Nullable Output<String> ruleId;

    public Optional<Output<String>> ruleId() {
        return Optional.ofNullable(this.ruleId);
    }

    private LoadBalancerForwardingRuleArgs() {}

    private LoadBalancerForwardingRuleArgs(LoadBalancerForwardingRuleArgs $) {
        this.backendPort = $.backendPort;
        this.backendProtocol = $.backendProtocol;
        this.frontendPort = $.frontendPort;
        this.frontendProtocol = $.frontendProtocol;
        this.ruleId = $.ruleId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LoadBalancerForwardingRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LoadBalancerForwardingRuleArgs $;

        public Builder() {
            $ = new LoadBalancerForwardingRuleArgs();
        }

        public Builder(LoadBalancerForwardingRuleArgs defaults) {
            $ = new LoadBalancerForwardingRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backendPort Port on instance side.
         * 
         * @return builder
         * 
         */
        public Builder backendPort(Output<Integer> backendPort) {
            $.backendPort = backendPort;
            return this;
        }

        /**
         * @param backendPort Port on instance side.
         * 
         * @return builder
         * 
         */
        public Builder backendPort(Integer backendPort) {
            return backendPort(Output.of(backendPort));
        }

        /**
         * @param backendProtocol Protocol on instance side. Possible values: &#34;http&#34;, &#34;https&#34;, &#34;tcp&#34;.
         * 
         * @return builder
         * 
         */
        public Builder backendProtocol(Output<String> backendProtocol) {
            $.backendProtocol = backendProtocol;
            return this;
        }

        /**
         * @param backendProtocol Protocol on instance side. Possible values: &#34;http&#34;, &#34;https&#34;, &#34;tcp&#34;.
         * 
         * @return builder
         * 
         */
        public Builder backendProtocol(String backendProtocol) {
            return backendProtocol(Output.of(backendProtocol));
        }

        /**
         * @param frontendPort Port on load balancer side.
         * 
         * @return builder
         * 
         */
        public Builder frontendPort(Output<Integer> frontendPort) {
            $.frontendPort = frontendPort;
            return this;
        }

        /**
         * @param frontendPort Port on load balancer side.
         * 
         * @return builder
         * 
         */
        public Builder frontendPort(Integer frontendPort) {
            return frontendPort(Output.of(frontendPort));
        }

        /**
         * @param frontendProtocol Protocol on load balancer side. Possible values: &#34;http&#34;, &#34;https&#34;, &#34;tcp&#34;.
         * 
         * @return builder
         * 
         */
        public Builder frontendProtocol(Output<String> frontendProtocol) {
            $.frontendProtocol = frontendProtocol;
            return this;
        }

        /**
         * @param frontendProtocol Protocol on load balancer side. Possible values: &#34;http&#34;, &#34;https&#34;, &#34;tcp&#34;.
         * 
         * @return builder
         * 
         */
        public Builder frontendProtocol(String frontendProtocol) {
            return frontendProtocol(Output.of(frontendProtocol));
        }

        public Builder ruleId(@Nullable Output<String> ruleId) {
            $.ruleId = ruleId;
            return this;
        }

        public Builder ruleId(String ruleId) {
            return ruleId(Output.of(ruleId));
        }

        public LoadBalancerForwardingRuleArgs build() {
            $.backendPort = Objects.requireNonNull($.backendPort, "expected parameter 'backendPort' to be non-null");
            $.backendProtocol = Objects.requireNonNull($.backendProtocol, "expected parameter 'backendProtocol' to be non-null");
            $.frontendPort = Objects.requireNonNull($.frontendPort, "expected parameter 'frontendPort' to be non-null");
            $.frontendProtocol = Objects.requireNonNull($.frontendProtocol, "expected parameter 'frontendProtocol' to be non-null");
            return $;
        }
    }

}
