// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ReverseIpv6State extends com.pulumi.resources.ResourceArgs {

    public static final ReverseIpv6State Empty = new ReverseIpv6State();

    /**
     * The ID of the server you want to set an IPv6
     * reverse DNS record for.
     * 
     */
    @Import(name="instanceId")
    private @Nullable Output<String> instanceId;

    /**
     * @return The ID of the server you want to set an IPv6
     * reverse DNS record for.
     * 
     */
    public Optional<Output<String>> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * The IPv6 address used in the reverse DNS record.
     * 
     */
    @Import(name="ip")
    private @Nullable Output<String> ip;

    /**
     * @return The IPv6 address used in the reverse DNS record.
     * 
     */
    public Optional<Output<String>> ip() {
        return Optional.ofNullable(this.ip);
    }

    /**
     * The hostname used in the IPv6 reverse DNS record.
     * 
     */
    @Import(name="reverse")
    private @Nullable Output<String> reverse;

    /**
     * @return The hostname used in the IPv6 reverse DNS record.
     * 
     */
    public Optional<Output<String>> reverse() {
        return Optional.ofNullable(this.reverse);
    }

    private ReverseIpv6State() {}

    private ReverseIpv6State(ReverseIpv6State $) {
        this.instanceId = $.instanceId;
        this.ip = $.ip;
        this.reverse = $.reverse;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ReverseIpv6State defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ReverseIpv6State $;

        public Builder() {
            $ = new ReverseIpv6State();
        }

        public Builder(ReverseIpv6State defaults) {
            $ = new ReverseIpv6State(Objects.requireNonNull(defaults));
        }

        /**
         * @param instanceId The ID of the server you want to set an IPv6
         * reverse DNS record for.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(@Nullable Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The ID of the server you want to set an IPv6
         * reverse DNS record for.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param ip The IPv6 address used in the reverse DNS record.
         * 
         * @return builder
         * 
         */
        public Builder ip(@Nullable Output<String> ip) {
            $.ip = ip;
            return this;
        }

        /**
         * @param ip The IPv6 address used in the reverse DNS record.
         * 
         * @return builder
         * 
         */
        public Builder ip(String ip) {
            return ip(Output.of(ip));
        }

        /**
         * @param reverse The hostname used in the IPv6 reverse DNS record.
         * 
         * @return builder
         * 
         */
        public Builder reverse(@Nullable Output<String> reverse) {
            $.reverse = reverse;
            return this;
        }

        /**
         * @param reverse The hostname used in the IPv6 reverse DNS record.
         * 
         * @return builder
         * 
         */
        public Builder reverse(String reverse) {
            return reverse(Output.of(reverse));
        }

        public ReverseIpv6State build() {
            return $;
        }
    }

}
