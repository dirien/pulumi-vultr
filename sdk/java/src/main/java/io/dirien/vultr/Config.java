// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr;

import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;

public final class Config {

    private static final com.pulumi.Config config = com.pulumi.Config.of("vultr");
/**
 * The API Key that allows interaction with the API
 * 
 */
    public String apiKey() {
        return Codegen.stringProp("apiKey").config(config).require();
    }
/**
 * Allows users to set the speed of API calls to work with the Vultr Rate Limit
 * 
 */
    public Optional<Integer> rateLimit() {
        return Codegen.integerProp("rateLimit").config(config).get();
    }
/**
 * Allows users to set the maximum number of retries allowed for a failed API call.
 * 
 */
    public Optional<Integer> retryLimit() {
        return Codegen.integerProp("retryLimit").config(config).get();
    }
}
