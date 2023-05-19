// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.vultr.outputs;

import com.pulumi.core.annotations.CustomType;
import io.dirien.vultr.outputs.GetDatabaseFilter;
import io.dirien.vultr.outputs.GetDatabaseReadReplica;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetDatabaseResult {
    private String clusterTimeZone;
    private String databaseEngine;
    private String databaseEngineVersion;
    private String dateCreated;
    private String dbname;
    private @Nullable List<GetDatabaseFilter> filters;
    private String host;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String label;
    private String latestBackup;
    private String maintenanceDow;
    private String maintenanceTime;
    private Integer mysqlLongQueryTime;
    private Boolean mysqlRequirePrimaryKey;
    private Boolean mysqlSlowQueryLog;
    private List<String> mysqlSqlModes;
    private String password;
    private String plan;
    private Integer planDisk;
    private Integer planRam;
    private Integer planReplicas;
    private Integer planVcpus;
    private String port;
    private List<GetDatabaseReadReplica> readReplicas;
    private String redisEvictionPolicy;
    private String region;
    private String status;
    private String tag;
    private List<String> trustedIps;
    private String user;

    private GetDatabaseResult() {}
    public String clusterTimeZone() {
        return this.clusterTimeZone;
    }
    public String databaseEngine() {
        return this.databaseEngine;
    }
    public String databaseEngineVersion() {
        return this.databaseEngineVersion;
    }
    public String dateCreated() {
        return this.dateCreated;
    }
    public String dbname() {
        return this.dbname;
    }
    public List<GetDatabaseFilter> filters() {
        return this.filters == null ? List.of() : this.filters;
    }
    public String host() {
        return this.host;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String label() {
        return this.label;
    }
    public String latestBackup() {
        return this.latestBackup;
    }
    public String maintenanceDow() {
        return this.maintenanceDow;
    }
    public String maintenanceTime() {
        return this.maintenanceTime;
    }
    public Integer mysqlLongQueryTime() {
        return this.mysqlLongQueryTime;
    }
    public Boolean mysqlRequirePrimaryKey() {
        return this.mysqlRequirePrimaryKey;
    }
    public Boolean mysqlSlowQueryLog() {
        return this.mysqlSlowQueryLog;
    }
    public List<String> mysqlSqlModes() {
        return this.mysqlSqlModes;
    }
    public String password() {
        return this.password;
    }
    public String plan() {
        return this.plan;
    }
    public Integer planDisk() {
        return this.planDisk;
    }
    public Integer planRam() {
        return this.planRam;
    }
    public Integer planReplicas() {
        return this.planReplicas;
    }
    public Integer planVcpus() {
        return this.planVcpus;
    }
    public String port() {
        return this.port;
    }
    public List<GetDatabaseReadReplica> readReplicas() {
        return this.readReplicas;
    }
    public String redisEvictionPolicy() {
        return this.redisEvictionPolicy;
    }
    public String region() {
        return this.region;
    }
    public String status() {
        return this.status;
    }
    public String tag() {
        return this.tag;
    }
    public List<String> trustedIps() {
        return this.trustedIps;
    }
    public String user() {
        return this.user;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDatabaseResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String clusterTimeZone;
        private String databaseEngine;
        private String databaseEngineVersion;
        private String dateCreated;
        private String dbname;
        private @Nullable List<GetDatabaseFilter> filters;
        private String host;
        private String id;
        private String label;
        private String latestBackup;
        private String maintenanceDow;
        private String maintenanceTime;
        private Integer mysqlLongQueryTime;
        private Boolean mysqlRequirePrimaryKey;
        private Boolean mysqlSlowQueryLog;
        private List<String> mysqlSqlModes;
        private String password;
        private String plan;
        private Integer planDisk;
        private Integer planRam;
        private Integer planReplicas;
        private Integer planVcpus;
        private String port;
        private List<GetDatabaseReadReplica> readReplicas;
        private String redisEvictionPolicy;
        private String region;
        private String status;
        private String tag;
        private List<String> trustedIps;
        private String user;
        public Builder() {}
        public Builder(GetDatabaseResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterTimeZone = defaults.clusterTimeZone;
    	      this.databaseEngine = defaults.databaseEngine;
    	      this.databaseEngineVersion = defaults.databaseEngineVersion;
    	      this.dateCreated = defaults.dateCreated;
    	      this.dbname = defaults.dbname;
    	      this.filters = defaults.filters;
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
    	      this.readReplicas = defaults.readReplicas;
    	      this.redisEvictionPolicy = defaults.redisEvictionPolicy;
    	      this.region = defaults.region;
    	      this.status = defaults.status;
    	      this.tag = defaults.tag;
    	      this.trustedIps = defaults.trustedIps;
    	      this.user = defaults.user;
        }

        @CustomType.Setter
        public Builder clusterTimeZone(String clusterTimeZone) {
            this.clusterTimeZone = Objects.requireNonNull(clusterTimeZone);
            return this;
        }
        @CustomType.Setter
        public Builder databaseEngine(String databaseEngine) {
            this.databaseEngine = Objects.requireNonNull(databaseEngine);
            return this;
        }
        @CustomType.Setter
        public Builder databaseEngineVersion(String databaseEngineVersion) {
            this.databaseEngineVersion = Objects.requireNonNull(databaseEngineVersion);
            return this;
        }
        @CustomType.Setter
        public Builder dateCreated(String dateCreated) {
            this.dateCreated = Objects.requireNonNull(dateCreated);
            return this;
        }
        @CustomType.Setter
        public Builder dbname(String dbname) {
            this.dbname = Objects.requireNonNull(dbname);
            return this;
        }
        @CustomType.Setter
        public Builder filters(@Nullable List<GetDatabaseFilter> filters) {
            this.filters = filters;
            return this;
        }
        public Builder filters(GetDatabaseFilter... filters) {
            return filters(List.of(filters));
        }
        @CustomType.Setter
        public Builder host(String host) {
            this.host = Objects.requireNonNull(host);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder label(String label) {
            this.label = Objects.requireNonNull(label);
            return this;
        }
        @CustomType.Setter
        public Builder latestBackup(String latestBackup) {
            this.latestBackup = Objects.requireNonNull(latestBackup);
            return this;
        }
        @CustomType.Setter
        public Builder maintenanceDow(String maintenanceDow) {
            this.maintenanceDow = Objects.requireNonNull(maintenanceDow);
            return this;
        }
        @CustomType.Setter
        public Builder maintenanceTime(String maintenanceTime) {
            this.maintenanceTime = Objects.requireNonNull(maintenanceTime);
            return this;
        }
        @CustomType.Setter
        public Builder mysqlLongQueryTime(Integer mysqlLongQueryTime) {
            this.mysqlLongQueryTime = Objects.requireNonNull(mysqlLongQueryTime);
            return this;
        }
        @CustomType.Setter
        public Builder mysqlRequirePrimaryKey(Boolean mysqlRequirePrimaryKey) {
            this.mysqlRequirePrimaryKey = Objects.requireNonNull(mysqlRequirePrimaryKey);
            return this;
        }
        @CustomType.Setter
        public Builder mysqlSlowQueryLog(Boolean mysqlSlowQueryLog) {
            this.mysqlSlowQueryLog = Objects.requireNonNull(mysqlSlowQueryLog);
            return this;
        }
        @CustomType.Setter
        public Builder mysqlSqlModes(List<String> mysqlSqlModes) {
            this.mysqlSqlModes = Objects.requireNonNull(mysqlSqlModes);
            return this;
        }
        public Builder mysqlSqlModes(String... mysqlSqlModes) {
            return mysqlSqlModes(List.of(mysqlSqlModes));
        }
        @CustomType.Setter
        public Builder password(String password) {
            this.password = Objects.requireNonNull(password);
            return this;
        }
        @CustomType.Setter
        public Builder plan(String plan) {
            this.plan = Objects.requireNonNull(plan);
            return this;
        }
        @CustomType.Setter
        public Builder planDisk(Integer planDisk) {
            this.planDisk = Objects.requireNonNull(planDisk);
            return this;
        }
        @CustomType.Setter
        public Builder planRam(Integer planRam) {
            this.planRam = Objects.requireNonNull(planRam);
            return this;
        }
        @CustomType.Setter
        public Builder planReplicas(Integer planReplicas) {
            this.planReplicas = Objects.requireNonNull(planReplicas);
            return this;
        }
        @CustomType.Setter
        public Builder planVcpus(Integer planVcpus) {
            this.planVcpus = Objects.requireNonNull(planVcpus);
            return this;
        }
        @CustomType.Setter
        public Builder port(String port) {
            this.port = Objects.requireNonNull(port);
            return this;
        }
        @CustomType.Setter
        public Builder readReplicas(List<GetDatabaseReadReplica> readReplicas) {
            this.readReplicas = Objects.requireNonNull(readReplicas);
            return this;
        }
        public Builder readReplicas(GetDatabaseReadReplica... readReplicas) {
            return readReplicas(List.of(readReplicas));
        }
        @CustomType.Setter
        public Builder redisEvictionPolicy(String redisEvictionPolicy) {
            this.redisEvictionPolicy = Objects.requireNonNull(redisEvictionPolicy);
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            this.region = Objects.requireNonNull(region);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder tag(String tag) {
            this.tag = Objects.requireNonNull(tag);
            return this;
        }
        @CustomType.Setter
        public Builder trustedIps(List<String> trustedIps) {
            this.trustedIps = Objects.requireNonNull(trustedIps);
            return this;
        }
        public Builder trustedIps(String... trustedIps) {
            return trustedIps(List.of(trustedIps));
        }
        @CustomType.Setter
        public Builder user(String user) {
            this.user = Objects.requireNonNull(user);
            return this;
        }
        public GetDatabaseResult build() {
            final var o = new GetDatabaseResult();
            o.clusterTimeZone = clusterTimeZone;
            o.databaseEngine = databaseEngine;
            o.databaseEngineVersion = databaseEngineVersion;
            o.dateCreated = dateCreated;
            o.dbname = dbname;
            o.filters = filters;
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
            o.readReplicas = readReplicas;
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