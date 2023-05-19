// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DatabaseReplicaArgs extends com.pulumi.resources.ResourceArgs {

    public static final DatabaseReplicaArgs Empty = new DatabaseReplicaArgs();

    @Import(name="databaseId", required=true)
    private Output<String> databaseId;

    public Output<String> databaseId() {
        return this.databaseId;
    }

    @Import(name="label", required=true)
    private Output<String> label;

    public Output<String> label() {
        return this.label;
    }

    @Import(name="mysqlLongQueryTime")
    private @Nullable Output<Integer> mysqlLongQueryTime;

    public Optional<Output<Integer>> mysqlLongQueryTime() {
        return Optional.ofNullable(this.mysqlLongQueryTime);
    }

    @Import(name="mysqlRequirePrimaryKey")
    private @Nullable Output<Boolean> mysqlRequirePrimaryKey;

    public Optional<Output<Boolean>> mysqlRequirePrimaryKey() {
        return Optional.ofNullable(this.mysqlRequirePrimaryKey);
    }

    @Import(name="mysqlSlowQueryLog")
    private @Nullable Output<Boolean> mysqlSlowQueryLog;

    public Optional<Output<Boolean>> mysqlSlowQueryLog() {
        return Optional.ofNullable(this.mysqlSlowQueryLog);
    }

    @Import(name="mysqlSqlModes")
    private @Nullable Output<List<String>> mysqlSqlModes;

    public Optional<Output<List<String>>> mysqlSqlModes() {
        return Optional.ofNullable(this.mysqlSqlModes);
    }

    @Import(name="planDisk")
    private @Nullable Output<Integer> planDisk;

    public Optional<Output<Integer>> planDisk() {
        return Optional.ofNullable(this.planDisk);
    }

    @Import(name="redisEvictionPolicy")
    private @Nullable Output<String> redisEvictionPolicy;

    public Optional<Output<String>> redisEvictionPolicy() {
        return Optional.ofNullable(this.redisEvictionPolicy);
    }

    @Import(name="region", required=true)
    private Output<String> region;

    public Output<String> region() {
        return this.region;
    }

    @Import(name="tag")
    private @Nullable Output<String> tag;

    public Optional<Output<String>> tag() {
        return Optional.ofNullable(this.tag);
    }

    @Import(name="trustedIps")
    private @Nullable Output<List<String>> trustedIps;

    public Optional<Output<List<String>>> trustedIps() {
        return Optional.ofNullable(this.trustedIps);
    }

    private DatabaseReplicaArgs() {}

    private DatabaseReplicaArgs(DatabaseReplicaArgs $) {
        this.databaseId = $.databaseId;
        this.label = $.label;
        this.mysqlLongQueryTime = $.mysqlLongQueryTime;
        this.mysqlRequirePrimaryKey = $.mysqlRequirePrimaryKey;
        this.mysqlSlowQueryLog = $.mysqlSlowQueryLog;
        this.mysqlSqlModes = $.mysqlSqlModes;
        this.planDisk = $.planDisk;
        this.redisEvictionPolicy = $.redisEvictionPolicy;
        this.region = $.region;
        this.tag = $.tag;
        this.trustedIps = $.trustedIps;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DatabaseReplicaArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DatabaseReplicaArgs $;

        public Builder() {
            $ = new DatabaseReplicaArgs();
        }

        public Builder(DatabaseReplicaArgs defaults) {
            $ = new DatabaseReplicaArgs(Objects.requireNonNull(defaults));
        }

        public Builder databaseId(Output<String> databaseId) {
            $.databaseId = databaseId;
            return this;
        }

        public Builder databaseId(String databaseId) {
            return databaseId(Output.of(databaseId));
        }

        public Builder label(Output<String> label) {
            $.label = label;
            return this;
        }

        public Builder label(String label) {
            return label(Output.of(label));
        }

        public Builder mysqlLongQueryTime(@Nullable Output<Integer> mysqlLongQueryTime) {
            $.mysqlLongQueryTime = mysqlLongQueryTime;
            return this;
        }

        public Builder mysqlLongQueryTime(Integer mysqlLongQueryTime) {
            return mysqlLongQueryTime(Output.of(mysqlLongQueryTime));
        }

        public Builder mysqlRequirePrimaryKey(@Nullable Output<Boolean> mysqlRequirePrimaryKey) {
            $.mysqlRequirePrimaryKey = mysqlRequirePrimaryKey;
            return this;
        }

        public Builder mysqlRequirePrimaryKey(Boolean mysqlRequirePrimaryKey) {
            return mysqlRequirePrimaryKey(Output.of(mysqlRequirePrimaryKey));
        }

        public Builder mysqlSlowQueryLog(@Nullable Output<Boolean> mysqlSlowQueryLog) {
            $.mysqlSlowQueryLog = mysqlSlowQueryLog;
            return this;
        }

        public Builder mysqlSlowQueryLog(Boolean mysqlSlowQueryLog) {
            return mysqlSlowQueryLog(Output.of(mysqlSlowQueryLog));
        }

        public Builder mysqlSqlModes(@Nullable Output<List<String>> mysqlSqlModes) {
            $.mysqlSqlModes = mysqlSqlModes;
            return this;
        }

        public Builder mysqlSqlModes(List<String> mysqlSqlModes) {
            return mysqlSqlModes(Output.of(mysqlSqlModes));
        }

        public Builder mysqlSqlModes(String... mysqlSqlModes) {
            return mysqlSqlModes(List.of(mysqlSqlModes));
        }

        public Builder planDisk(@Nullable Output<Integer> planDisk) {
            $.planDisk = planDisk;
            return this;
        }

        public Builder planDisk(Integer planDisk) {
            return planDisk(Output.of(planDisk));
        }

        public Builder redisEvictionPolicy(@Nullable Output<String> redisEvictionPolicy) {
            $.redisEvictionPolicy = redisEvictionPolicy;
            return this;
        }

        public Builder redisEvictionPolicy(String redisEvictionPolicy) {
            return redisEvictionPolicy(Output.of(redisEvictionPolicy));
        }

        public Builder region(Output<String> region) {
            $.region = region;
            return this;
        }

        public Builder region(String region) {
            return region(Output.of(region));
        }

        public Builder tag(@Nullable Output<String> tag) {
            $.tag = tag;
            return this;
        }

        public Builder tag(String tag) {
            return tag(Output.of(tag));
        }

        public Builder trustedIps(@Nullable Output<List<String>> trustedIps) {
            $.trustedIps = trustedIps;
            return this;
        }

        public Builder trustedIps(List<String> trustedIps) {
            return trustedIps(Output.of(trustedIps));
        }

        public Builder trustedIps(String... trustedIps) {
            return trustedIps(List.of(trustedIps));
        }

        public DatabaseReplicaArgs build() {
            $.databaseId = Objects.requireNonNull($.databaseId, "expected parameter 'databaseId' to be non-null");
            $.label = Objects.requireNonNull($.label, "expected parameter 'label' to be non-null");
            $.region = Objects.requireNonNull($.region, "expected parameter 'region' to be non-null");
            return $;
        }
    }

}