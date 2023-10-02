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

    /**
     * The managed database ID you want to attach this replica to.
     * 
     */
    @Import(name="databaseId", required=true)
    private Output<String> databaseId;

    /**
     * @return The managed database ID you want to attach this replica to.
     * 
     */
    public Output<String> databaseId() {
        return this.databaseId;
    }

    /**
     * A label for the managed database read replica.
     * 
     */
    @Import(name="label", required=true)
    private Output<String> label;

    /**
     * @return A label for the managed database read replica.
     * 
     */
    public Output<String> label() {
        return this.label;
    }

    /**
     * The configuration value for the long query time (in seconds) on the managed database read replica (MySQL engine types only).
     * 
     */
    @Import(name="mysqlLongQueryTime")
    private @Nullable Output<Integer> mysqlLongQueryTime;

    /**
     * @return The configuration value for the long query time (in seconds) on the managed database read replica (MySQL engine types only).
     * 
     */
    public Optional<Output<Integer>> mysqlLongQueryTime() {
        return Optional.ofNullable(this.mysqlLongQueryTime);
    }

    /**
     * The configuration value for whether primary keys are required on the managed database read replica (MySQL engine types only).
     * 
     */
    @Import(name="mysqlRequirePrimaryKey")
    private @Nullable Output<Boolean> mysqlRequirePrimaryKey;

    /**
     * @return The configuration value for whether primary keys are required on the managed database read replica (MySQL engine types only).
     * 
     */
    public Optional<Output<Boolean>> mysqlRequirePrimaryKey() {
        return Optional.ofNullable(this.mysqlRequirePrimaryKey);
    }

    /**
     * The configuration value for slow query logging on the managed database read replica (MySQL engine types only).
     * 
     */
    @Import(name="mysqlSlowQueryLog")
    private @Nullable Output<Boolean> mysqlSlowQueryLog;

    /**
     * @return The configuration value for slow query logging on the managed database read replica (MySQL engine types only).
     * 
     */
    public Optional<Output<Boolean>> mysqlSlowQueryLog() {
        return Optional.ofNullable(this.mysqlSlowQueryLog);
    }

    /**
     * A list of SQL modes currently configured for the managed database read replica (MySQL engine types only).
     * 
     */
    @Import(name="mysqlSqlModes")
    private @Nullable Output<List<String>> mysqlSqlModes;

    /**
     * @return A list of SQL modes currently configured for the managed database read replica (MySQL engine types only).
     * 
     */
    public Optional<Output<List<String>>> mysqlSqlModes() {
        return Optional.ofNullable(this.mysqlSqlModes);
    }

    /**
     * The description of the disk(s) on the managed database read replica.
     * 
     */
    @Import(name="planDisk")
    private @Nullable Output<Integer> planDisk;

    /**
     * @return The description of the disk(s) on the managed database read replica.
     * 
     */
    public Optional<Output<Integer>> planDisk() {
        return Optional.ofNullable(this.planDisk);
    }

    /**
     * The configuration value for the data eviction policy on the managed database read replica (Redis engine types only).
     * 
     */
    @Import(name="redisEvictionPolicy")
    private @Nullable Output<String> redisEvictionPolicy;

    /**
     * @return The configuration value for the data eviction policy on the managed database read replica (Redis engine types only).
     * 
     */
    public Optional<Output<String>> redisEvictionPolicy() {
        return Optional.ofNullable(this.redisEvictionPolicy);
    }

    /**
     * The ID of the region that the managed database read replica is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
     * 
     */
    @Import(name="region", required=true)
    private Output<String> region;

    /**
     * @return The ID of the region that the managed database read replica is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     * The tag to assign to the managed database read replica.
     * 
     */
    @Import(name="tag")
    private @Nullable Output<String> tag;

    /**
     * @return The tag to assign to the managed database read replica.
     * 
     */
    public Optional<Output<String>> tag() {
        return Optional.ofNullable(this.tag);
    }

    /**
     * A list of allowed IP addresses for the managed database read replica.
     * 
     */
    @Import(name="trustedIps")
    private @Nullable Output<List<String>> trustedIps;

    /**
     * @return A list of allowed IP addresses for the managed database read replica.
     * 
     */
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

        /**
         * @param databaseId The managed database ID you want to attach this replica to.
         * 
         * @return builder
         * 
         */
        public Builder databaseId(Output<String> databaseId) {
            $.databaseId = databaseId;
            return this;
        }

        /**
         * @param databaseId The managed database ID you want to attach this replica to.
         * 
         * @return builder
         * 
         */
        public Builder databaseId(String databaseId) {
            return databaseId(Output.of(databaseId));
        }

        /**
         * @param label A label for the managed database read replica.
         * 
         * @return builder
         * 
         */
        public Builder label(Output<String> label) {
            $.label = label;
            return this;
        }

        /**
         * @param label A label for the managed database read replica.
         * 
         * @return builder
         * 
         */
        public Builder label(String label) {
            return label(Output.of(label));
        }

        /**
         * @param mysqlLongQueryTime The configuration value for the long query time (in seconds) on the managed database read replica (MySQL engine types only).
         * 
         * @return builder
         * 
         */
        public Builder mysqlLongQueryTime(@Nullable Output<Integer> mysqlLongQueryTime) {
            $.mysqlLongQueryTime = mysqlLongQueryTime;
            return this;
        }

        /**
         * @param mysqlLongQueryTime The configuration value for the long query time (in seconds) on the managed database read replica (MySQL engine types only).
         * 
         * @return builder
         * 
         */
        public Builder mysqlLongQueryTime(Integer mysqlLongQueryTime) {
            return mysqlLongQueryTime(Output.of(mysqlLongQueryTime));
        }

        /**
         * @param mysqlRequirePrimaryKey The configuration value for whether primary keys are required on the managed database read replica (MySQL engine types only).
         * 
         * @return builder
         * 
         */
        public Builder mysqlRequirePrimaryKey(@Nullable Output<Boolean> mysqlRequirePrimaryKey) {
            $.mysqlRequirePrimaryKey = mysqlRequirePrimaryKey;
            return this;
        }

        /**
         * @param mysqlRequirePrimaryKey The configuration value for whether primary keys are required on the managed database read replica (MySQL engine types only).
         * 
         * @return builder
         * 
         */
        public Builder mysqlRequirePrimaryKey(Boolean mysqlRequirePrimaryKey) {
            return mysqlRequirePrimaryKey(Output.of(mysqlRequirePrimaryKey));
        }

        /**
         * @param mysqlSlowQueryLog The configuration value for slow query logging on the managed database read replica (MySQL engine types only).
         * 
         * @return builder
         * 
         */
        public Builder mysqlSlowQueryLog(@Nullable Output<Boolean> mysqlSlowQueryLog) {
            $.mysqlSlowQueryLog = mysqlSlowQueryLog;
            return this;
        }

        /**
         * @param mysqlSlowQueryLog The configuration value for slow query logging on the managed database read replica (MySQL engine types only).
         * 
         * @return builder
         * 
         */
        public Builder mysqlSlowQueryLog(Boolean mysqlSlowQueryLog) {
            return mysqlSlowQueryLog(Output.of(mysqlSlowQueryLog));
        }

        /**
         * @param mysqlSqlModes A list of SQL modes currently configured for the managed database read replica (MySQL engine types only).
         * 
         * @return builder
         * 
         */
        public Builder mysqlSqlModes(@Nullable Output<List<String>> mysqlSqlModes) {
            $.mysqlSqlModes = mysqlSqlModes;
            return this;
        }

        /**
         * @param mysqlSqlModes A list of SQL modes currently configured for the managed database read replica (MySQL engine types only).
         * 
         * @return builder
         * 
         */
        public Builder mysqlSqlModes(List<String> mysqlSqlModes) {
            return mysqlSqlModes(Output.of(mysqlSqlModes));
        }

        /**
         * @param mysqlSqlModes A list of SQL modes currently configured for the managed database read replica (MySQL engine types only).
         * 
         * @return builder
         * 
         */
        public Builder mysqlSqlModes(String... mysqlSqlModes) {
            return mysqlSqlModes(List.of(mysqlSqlModes));
        }

        /**
         * @param planDisk The description of the disk(s) on the managed database read replica.
         * 
         * @return builder
         * 
         */
        public Builder planDisk(@Nullable Output<Integer> planDisk) {
            $.planDisk = planDisk;
            return this;
        }

        /**
         * @param planDisk The description of the disk(s) on the managed database read replica.
         * 
         * @return builder
         * 
         */
        public Builder planDisk(Integer planDisk) {
            return planDisk(Output.of(planDisk));
        }

        /**
         * @param redisEvictionPolicy The configuration value for the data eviction policy on the managed database read replica (Redis engine types only).
         * 
         * @return builder
         * 
         */
        public Builder redisEvictionPolicy(@Nullable Output<String> redisEvictionPolicy) {
            $.redisEvictionPolicy = redisEvictionPolicy;
            return this;
        }

        /**
         * @param redisEvictionPolicy The configuration value for the data eviction policy on the managed database read replica (Redis engine types only).
         * 
         * @return builder
         * 
         */
        public Builder redisEvictionPolicy(String redisEvictionPolicy) {
            return redisEvictionPolicy(Output.of(redisEvictionPolicy));
        }

        /**
         * @param region The ID of the region that the managed database read replica is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
         * 
         * @return builder
         * 
         */
        public Builder region(Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The ID of the region that the managed database read replica is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param tag The tag to assign to the managed database read replica.
         * 
         * @return builder
         * 
         */
        public Builder tag(@Nullable Output<String> tag) {
            $.tag = tag;
            return this;
        }

        /**
         * @param tag The tag to assign to the managed database read replica.
         * 
         * @return builder
         * 
         */
        public Builder tag(String tag) {
            return tag(Output.of(tag));
        }

        /**
         * @param trustedIps A list of allowed IP addresses for the managed database read replica.
         * 
         * @return builder
         * 
         */
        public Builder trustedIps(@Nullable Output<List<String>> trustedIps) {
            $.trustedIps = trustedIps;
            return this;
        }

        /**
         * @param trustedIps A list of allowed IP addresses for the managed database read replica.
         * 
         * @return builder
         * 
         */
        public Builder trustedIps(List<String> trustedIps) {
            return trustedIps(Output.of(trustedIps));
        }

        /**
         * @param trustedIps A list of allowed IP addresses for the managed database read replica.
         * 
         * @return builder
         * 
         */
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
