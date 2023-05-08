// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DatabaseReadReplica {
    private @Nullable String clusterTimeZone;
    private @Nullable String databaseEngine;
    private @Nullable String databaseEngineVersion;
    private @Nullable String dateCreated;
    private @Nullable String dbname;
    private @Nullable String host;
    private @Nullable String id;
    private String label;
    private @Nullable String latestBackup;
    private @Nullable String maintenanceDow;
    private @Nullable String maintenanceTime;
    private @Nullable Integer mysqlLongQueryTime;
    private @Nullable Boolean mysqlRequirePrimaryKey;
    private @Nullable Boolean mysqlSlowQueryLog;
    private @Nullable List<String> mysqlSqlModes;
    private @Nullable String password;
    private @Nullable String plan;
    private @Nullable Integer planDisk;
    private @Nullable Integer planRam;
    private @Nullable Integer planReplicas;
    private @Nullable Integer planVcpus;
    private @Nullable String port;
    private @Nullable String redisEvictionPolicy;
    private String region;
    private @Nullable String status;
    private @Nullable String tag;
    private @Nullable List<String> trustedIps;
    private @Nullable String user;

    private DatabaseReadReplica() {}
    public Optional<String> clusterTimeZone() {
        return Optional.ofNullable(this.clusterTimeZone);
    }
    public Optional<String> databaseEngine() {
        return Optional.ofNullable(this.databaseEngine);
    }
    public Optional<String> databaseEngineVersion() {
        return Optional.ofNullable(this.databaseEngineVersion);
    }
    public Optional<String> dateCreated() {
        return Optional.ofNullable(this.dateCreated);
    }
    public Optional<String> dbname() {
        return Optional.ofNullable(this.dbname);
    }
    public Optional<String> host() {
        return Optional.ofNullable(this.host);
    }
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }
    public String label() {
        return this.label;
    }
    public Optional<String> latestBackup() {
        return Optional.ofNullable(this.latestBackup);
    }
    public Optional<String> maintenanceDow() {
        return Optional.ofNullable(this.maintenanceDow);
    }
    public Optional<String> maintenanceTime() {
        return Optional.ofNullable(this.maintenanceTime);
    }
    public Optional<Integer> mysqlLongQueryTime() {
        return Optional.ofNullable(this.mysqlLongQueryTime);
    }
    public Optional<Boolean> mysqlRequirePrimaryKey() {
        return Optional.ofNullable(this.mysqlRequirePrimaryKey);
    }
    public Optional<Boolean> mysqlSlowQueryLog() {
        return Optional.ofNullable(this.mysqlSlowQueryLog);
    }
    public List<String> mysqlSqlModes() {
        return this.mysqlSqlModes == null ? List.of() : this.mysqlSqlModes;
    }
    public Optional<String> password() {
        return Optional.ofNullable(this.password);
    }
    public Optional<String> plan() {
        return Optional.ofNullable(this.plan);
    }
    public Optional<Integer> planDisk() {
        return Optional.ofNullable(this.planDisk);
    }
    public Optional<Integer> planRam() {
        return Optional.ofNullable(this.planRam);
    }
    public Optional<Integer> planReplicas() {
        return Optional.ofNullable(this.planReplicas);
    }
    public Optional<Integer> planVcpus() {
        return Optional.ofNullable(this.planVcpus);
    }
    public Optional<String> port() {
        return Optional.ofNullable(this.port);
    }
    public Optional<String> redisEvictionPolicy() {
        return Optional.ofNullable(this.redisEvictionPolicy);
    }
    public String region() {
        return this.region;
    }
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    public Optional<String> tag() {
        return Optional.ofNullable(this.tag);
    }
    public List<String> trustedIps() {
        return this.trustedIps == null ? List.of() : this.trustedIps;
    }
    public Optional<String> user() {
        return Optional.ofNullable(this.user);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DatabaseReadReplica defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String clusterTimeZone;
        private @Nullable String databaseEngine;
        private @Nullable String databaseEngineVersion;
        private @Nullable String dateCreated;
        private @Nullable String dbname;
        private @Nullable String host;
        private @Nullable String id;
        private String label;
        private @Nullable String latestBackup;
        private @Nullable String maintenanceDow;
        private @Nullable String maintenanceTime;
        private @Nullable Integer mysqlLongQueryTime;
        private @Nullable Boolean mysqlRequirePrimaryKey;
        private @Nullable Boolean mysqlSlowQueryLog;
        private @Nullable List<String> mysqlSqlModes;
        private @Nullable String password;
        private @Nullable String plan;
        private @Nullable Integer planDisk;
        private @Nullable Integer planRam;
        private @Nullable Integer planReplicas;
        private @Nullable Integer planVcpus;
        private @Nullable String port;
        private @Nullable String redisEvictionPolicy;
        private String region;
        private @Nullable String status;
        private @Nullable String tag;
        private @Nullable List<String> trustedIps;
        private @Nullable String user;
        public Builder() {}
        public Builder(DatabaseReadReplica defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterTimeZone = defaults.clusterTimeZone;
    	      this.databaseEngine = defaults.databaseEngine;
    	      this.databaseEngineVersion = defaults.databaseEngineVersion;
    	      this.dateCreated = defaults.dateCreated;
    	      this.dbname = defaults.dbname;
    	      this.host = defaults.host;
    	      this.id = defaults.id;
    	      this.label = defaults.label;
    	      this.latestBackup = defaults.latestBackup;
    	      this.maintenanceDow = defaults.maintenanceDow;
    	      this.maintenanceTime = defaults.maintenanceTime;
    	      this.mysqlLongQueryTime = defaults.mysqlLongQueryTime;
    	      this.mysqlRequirePrimaryKey = defaults.mysqlRequirePrimaryKey;
    	      this.mysqlSlowQueryLog = defaults.mysqlSlowQueryLog;
    	      this.mysqlSqlModes = defaults.mysqlSqlModes;
    	      this.password = defaults.password;
    	      this.plan = defaults.plan;
    	      this.planDisk = defaults.planDisk;
    	      this.planRam = defaults.planRam;
    	      this.planReplicas = defaults.planReplicas;
    	      this.planVcpus = defaults.planVcpus;
    	      this.port = defaults.port;
    	      this.redisEvictionPolicy = defaults.redisEvictionPolicy;
    	      this.region = defaults.region;
    	      this.status = defaults.status;
    	      this.tag = defaults.tag;
    	      this.trustedIps = defaults.trustedIps;
    	      this.user = defaults.user;
        }

        @CustomType.Setter
        public Builder clusterTimeZone(@Nullable String clusterTimeZone) {
            this.clusterTimeZone = clusterTimeZone;
            return this;
        }
        @CustomType.Setter
        public Builder databaseEngine(@Nullable String databaseEngine) {
            this.databaseEngine = databaseEngine;
            return this;
        }
        @CustomType.Setter
        public Builder databaseEngineVersion(@Nullable String databaseEngineVersion) {
            this.databaseEngineVersion = databaseEngineVersion;
            return this;
        }
        @CustomType.Setter
        public Builder dateCreated(@Nullable String dateCreated) {
            this.dateCreated = dateCreated;
            return this;
        }
        @CustomType.Setter
        public Builder dbname(@Nullable String dbname) {
            this.dbname = dbname;
            return this;
        }
        @CustomType.Setter
        public Builder host(@Nullable String host) {
            this.host = host;
            return this;
        }
        @CustomType.Setter
        public Builder id(@Nullable String id) {
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder label(String label) {
            this.label = Objects.requireNonNull(label);
            return this;
        }
        @CustomType.Setter
        public Builder latestBackup(@Nullable String latestBackup) {
            this.latestBackup = latestBackup;
            return this;
        }
        @CustomType.Setter
        public Builder maintenanceDow(@Nullable String maintenanceDow) {
            this.maintenanceDow = maintenanceDow;
            return this;
        }
        @CustomType.Setter
        public Builder maintenanceTime(@Nullable String maintenanceTime) {
            this.maintenanceTime = maintenanceTime;
            return this;
        }
        @CustomType.Setter
        public Builder mysqlLongQueryTime(@Nullable Integer mysqlLongQueryTime) {
            this.mysqlLongQueryTime = mysqlLongQueryTime;
            return this;
        }
        @CustomType.Setter
        public Builder mysqlRequirePrimaryKey(@Nullable Boolean mysqlRequirePrimaryKey) {
            this.mysqlRequirePrimaryKey = mysqlRequirePrimaryKey;
            return this;
        }
        @CustomType.Setter
        public Builder mysqlSlowQueryLog(@Nullable Boolean mysqlSlowQueryLog) {
            this.mysqlSlowQueryLog = mysqlSlowQueryLog;
            return this;
        }
        @CustomType.Setter
        public Builder mysqlSqlModes(@Nullable List<String> mysqlSqlModes) {
            this.mysqlSqlModes = mysqlSqlModes;
            return this;
        }
        public Builder mysqlSqlModes(String... mysqlSqlModes) {
            return mysqlSqlModes(List.of(mysqlSqlModes));
        }
        @CustomType.Setter
        public Builder password(@Nullable String password) {
            this.password = password;
            return this;
        }
        @CustomType.Setter
        public Builder plan(@Nullable String plan) {
            this.plan = plan;
            return this;
        }
        @CustomType.Setter
        public Builder planDisk(@Nullable Integer planDisk) {
            this.planDisk = planDisk;
            return this;
        }
        @CustomType.Setter
        public Builder planRam(@Nullable Integer planRam) {
            this.planRam = planRam;
            return this;
        }
        @CustomType.Setter
        public Builder planReplicas(@Nullable Integer planReplicas) {
            this.planReplicas = planReplicas;
            return this;
        }
        @CustomType.Setter
        public Builder planVcpus(@Nullable Integer planVcpus) {
            this.planVcpus = planVcpus;
            return this;
        }
        @CustomType.Setter
        public Builder port(@Nullable String port) {
            this.port = port;
            return this;
        }
        @CustomType.Setter
        public Builder redisEvictionPolicy(@Nullable String redisEvictionPolicy) {
            this.redisEvictionPolicy = redisEvictionPolicy;
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            this.region = Objects.requireNonNull(region);
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder tag(@Nullable String tag) {
            this.tag = tag;
            return this;
        }
        @CustomType.Setter
        public Builder trustedIps(@Nullable List<String> trustedIps) {
            this.trustedIps = trustedIps;
            return this;
        }
        public Builder trustedIps(String... trustedIps) {
            return trustedIps(List.of(trustedIps));
        }
        @CustomType.Setter
        public Builder user(@Nullable String user) {
            this.user = user;
            return this;
        }
        public DatabaseReadReplica build() {
            final var o = new DatabaseReadReplica();
            o.clusterTimeZone = clusterTimeZone;
            o.databaseEngine = databaseEngine;
            o.databaseEngineVersion = databaseEngineVersion;
            o.dateCreated = dateCreated;
            o.dbname = dbname;
            o.host = host;
            o.id = id;
            o.label = label;
            o.latestBackup = latestBackup;
            o.maintenanceDow = maintenanceDow;
            o.maintenanceTime = maintenanceTime;
            o.mysqlLongQueryTime = mysqlLongQueryTime;
            o.mysqlRequirePrimaryKey = mysqlRequirePrimaryKey;
            o.mysqlSlowQueryLog = mysqlSlowQueryLog;
            o.mysqlSqlModes = mysqlSqlModes;
            o.password = password;
            o.plan = plan;
            o.planDisk = planDisk;
            o.planRam = planRam;
            o.planReplicas = planReplicas;
            o.planVcpus = planVcpus;
            o.port = port;
            o.redisEvictionPolicy = redisEvictionPolicy;
            o.region = region;
            o.status = status;
            o.tag = tag;
            o.trustedIps = trustedIps;
            o.user = user;
            return o;
        }
    }
}
