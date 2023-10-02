# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetDatabaseResult',
    'AwaitableGetDatabaseResult',
    'get_database',
    'get_database_output',
]

@pulumi.output_type
class GetDatabaseResult:
    """
    A collection of values returned by getDatabase.
    """
    def __init__(__self__, cluster_time_zone=None, database_engine=None, database_engine_version=None, date_created=None, dbname=None, filters=None, host=None, id=None, label=None, latest_backup=None, maintenance_dow=None, maintenance_time=None, mysql_long_query_time=None, mysql_require_primary_key=None, mysql_slow_query_log=None, mysql_sql_modes=None, password=None, plan=None, plan_disk=None, plan_ram=None, plan_replicas=None, plan_vcpus=None, port=None, read_replicas=None, redis_eviction_policy=None, region=None, status=None, tag=None, trusted_ips=None, user=None, vpc_id=None):
        if cluster_time_zone and not isinstance(cluster_time_zone, str):
            raise TypeError("Expected argument 'cluster_time_zone' to be a str")
        pulumi.set(__self__, "cluster_time_zone", cluster_time_zone)
        if database_engine and not isinstance(database_engine, str):
            raise TypeError("Expected argument 'database_engine' to be a str")
        pulumi.set(__self__, "database_engine", database_engine)
        if database_engine_version and not isinstance(database_engine_version, str):
            raise TypeError("Expected argument 'database_engine_version' to be a str")
        pulumi.set(__self__, "database_engine_version", database_engine_version)
        if date_created and not isinstance(date_created, str):
            raise TypeError("Expected argument 'date_created' to be a str")
        pulumi.set(__self__, "date_created", date_created)
        if dbname and not isinstance(dbname, str):
            raise TypeError("Expected argument 'dbname' to be a str")
        pulumi.set(__self__, "dbname", dbname)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if host and not isinstance(host, str):
            raise TypeError("Expected argument 'host' to be a str")
        pulumi.set(__self__, "host", host)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if label and not isinstance(label, str):
            raise TypeError("Expected argument 'label' to be a str")
        pulumi.set(__self__, "label", label)
        if latest_backup and not isinstance(latest_backup, str):
            raise TypeError("Expected argument 'latest_backup' to be a str")
        pulumi.set(__self__, "latest_backup", latest_backup)
        if maintenance_dow and not isinstance(maintenance_dow, str):
            raise TypeError("Expected argument 'maintenance_dow' to be a str")
        pulumi.set(__self__, "maintenance_dow", maintenance_dow)
        if maintenance_time and not isinstance(maintenance_time, str):
            raise TypeError("Expected argument 'maintenance_time' to be a str")
        pulumi.set(__self__, "maintenance_time", maintenance_time)
        if mysql_long_query_time and not isinstance(mysql_long_query_time, int):
            raise TypeError("Expected argument 'mysql_long_query_time' to be a int")
        pulumi.set(__self__, "mysql_long_query_time", mysql_long_query_time)
        if mysql_require_primary_key and not isinstance(mysql_require_primary_key, bool):
            raise TypeError("Expected argument 'mysql_require_primary_key' to be a bool")
        pulumi.set(__self__, "mysql_require_primary_key", mysql_require_primary_key)
        if mysql_slow_query_log and not isinstance(mysql_slow_query_log, bool):
            raise TypeError("Expected argument 'mysql_slow_query_log' to be a bool")
        pulumi.set(__self__, "mysql_slow_query_log", mysql_slow_query_log)
        if mysql_sql_modes and not isinstance(mysql_sql_modes, list):
            raise TypeError("Expected argument 'mysql_sql_modes' to be a list")
        pulumi.set(__self__, "mysql_sql_modes", mysql_sql_modes)
        if password and not isinstance(password, str):
            raise TypeError("Expected argument 'password' to be a str")
        pulumi.set(__self__, "password", password)
        if plan and not isinstance(plan, str):
            raise TypeError("Expected argument 'plan' to be a str")
        pulumi.set(__self__, "plan", plan)
        if plan_disk and not isinstance(plan_disk, int):
            raise TypeError("Expected argument 'plan_disk' to be a int")
        pulumi.set(__self__, "plan_disk", plan_disk)
        if plan_ram and not isinstance(plan_ram, int):
            raise TypeError("Expected argument 'plan_ram' to be a int")
        pulumi.set(__self__, "plan_ram", plan_ram)
        if plan_replicas and not isinstance(plan_replicas, int):
            raise TypeError("Expected argument 'plan_replicas' to be a int")
        pulumi.set(__self__, "plan_replicas", plan_replicas)
        if plan_vcpus and not isinstance(plan_vcpus, int):
            raise TypeError("Expected argument 'plan_vcpus' to be a int")
        pulumi.set(__self__, "plan_vcpus", plan_vcpus)
        if port and not isinstance(port, str):
            raise TypeError("Expected argument 'port' to be a str")
        pulumi.set(__self__, "port", port)
        if read_replicas and not isinstance(read_replicas, list):
            raise TypeError("Expected argument 'read_replicas' to be a list")
        pulumi.set(__self__, "read_replicas", read_replicas)
        if redis_eviction_policy and not isinstance(redis_eviction_policy, str):
            raise TypeError("Expected argument 'redis_eviction_policy' to be a str")
        pulumi.set(__self__, "redis_eviction_policy", redis_eviction_policy)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tag and not isinstance(tag, str):
            raise TypeError("Expected argument 'tag' to be a str")
        pulumi.set(__self__, "tag", tag)
        if trusted_ips and not isinstance(trusted_ips, list):
            raise TypeError("Expected argument 'trusted_ips' to be a list")
        pulumi.set(__self__, "trusted_ips", trusted_ips)
        if user and not isinstance(user, str):
            raise TypeError("Expected argument 'user' to be a str")
        pulumi.set(__self__, "user", user)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="clusterTimeZone")
    def cluster_time_zone(self) -> str:
        """
        The configured time zone for the Managed Database in TZ database format.
        """
        return pulumi.get(self, "cluster_time_zone")

    @property
    @pulumi.getter(name="databaseEngine")
    def database_engine(self) -> str:
        """
        The database engine of the managed database.
        """
        return pulumi.get(self, "database_engine")

    @property
    @pulumi.getter(name="databaseEngineVersion")
    def database_engine_version(self) -> str:
        """
        The database engine version of the managed database.
        """
        return pulumi.get(self, "database_engine_version")

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> str:
        """
        The date the managed database was added to your Vultr account.
        """
        return pulumi.get(self, "date_created")

    @property
    @pulumi.getter
    def dbname(self) -> str:
        """
        The managed database's default logical database.
        """
        return pulumi.get(self, "dbname")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetDatabaseFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def host(self) -> str:
        """
        The hostname assigned to the managed database.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def label(self) -> str:
        """
        The managed database's label.
        """
        return pulumi.get(self, "label")

    @property
    @pulumi.getter(name="latestBackup")
    def latest_backup(self) -> str:
        """
        The date of the latest backup available on the managed database.
        """
        return pulumi.get(self, "latest_backup")

    @property
    @pulumi.getter(name="maintenanceDow")
    def maintenance_dow(self) -> str:
        """
        The preferred maintenance day of week for the managed database.
        """
        return pulumi.get(self, "maintenance_dow")

    @property
    @pulumi.getter(name="maintenanceTime")
    def maintenance_time(self) -> str:
        """
        The preferred maintenance time for the managed database.
        """
        return pulumi.get(self, "maintenance_time")

    @property
    @pulumi.getter(name="mysqlLongQueryTime")
    def mysql_long_query_time(self) -> int:
        """
        The configuration value for the long query time (in seconds) on the managed database (MySQL engine types only).
        """
        return pulumi.get(self, "mysql_long_query_time")

    @property
    @pulumi.getter(name="mysqlRequirePrimaryKey")
    def mysql_require_primary_key(self) -> bool:
        """
        The configuration value for whether primary keys are required on the managed database (MySQL engine types only).
        """
        return pulumi.get(self, "mysql_require_primary_key")

    @property
    @pulumi.getter(name="mysqlSlowQueryLog")
    def mysql_slow_query_log(self) -> bool:
        """
        The configuration value for slow query logging on the managed database (MySQL engine types only).
        """
        return pulumi.get(self, "mysql_slow_query_log")

    @property
    @pulumi.getter(name="mysqlSqlModes")
    def mysql_sql_modes(self) -> Sequence[str]:
        """
        A list of SQL modes currently configured for the managed database (MySQL engine types only).
        """
        return pulumi.get(self, "mysql_sql_modes")

    @property
    @pulumi.getter
    def password(self) -> str:
        """
        The password for the managed database's primary admin user.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def plan(self) -> str:
        """
        The managed database's plan ID.
        """
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter(name="planDisk")
    def plan_disk(self) -> int:
        """
        The description of the disk(s) on the managed database.
        """
        return pulumi.get(self, "plan_disk")

    @property
    @pulumi.getter(name="planRam")
    def plan_ram(self) -> int:
        """
        The amount of memory available on the managed database in MB.
        """
        return pulumi.get(self, "plan_ram")

    @property
    @pulumi.getter(name="planReplicas")
    def plan_replicas(self) -> int:
        """
        The number of standby nodes available on the managed database.
        """
        return pulumi.get(self, "plan_replicas")

    @property
    @pulumi.getter(name="planVcpus")
    def plan_vcpus(self) -> int:
        """
        The number of virtual CPUs available on the managed database.
        """
        return pulumi.get(self, "plan_vcpus")

    @property
    @pulumi.getter
    def port(self) -> str:
        """
        The connection port for the managed database.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="readReplicas")
    def read_replicas(self) -> Sequence['outputs.GetDatabaseReadReplicaResult']:
        """
        A list of read replicas attached to the managed database.
        """
        return pulumi.get(self, "read_replicas")

    @property
    @pulumi.getter(name="redisEvictionPolicy")
    def redis_eviction_policy(self) -> str:
        """
        The configuration value for the data eviction policy on the managed database (Redis engine types only).
        """
        return pulumi.get(self, "redis_eviction_policy")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        The region ID of the managed database.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The current status of the managed database (poweroff, rebuilding, rebalancing, running).
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tag(self) -> str:
        """
        The managed database's tag.
        """
        return pulumi.get(self, "tag")

    @property
    @pulumi.getter(name="trustedIps")
    def trusted_ips(self) -> Sequence[str]:
        """
        A list of allowed IP addresses for the managed database.
        """
        return pulumi.get(self, "trusted_ips")

    @property
    @pulumi.getter
    def user(self) -> str:
        """
        The primary admin user for the managed database.
        """
        return pulumi.get(self, "user")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        return pulumi.get(self, "vpc_id")


class AwaitableGetDatabaseResult(GetDatabaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseResult(
            cluster_time_zone=self.cluster_time_zone,
            database_engine=self.database_engine,
            database_engine_version=self.database_engine_version,
            date_created=self.date_created,
            dbname=self.dbname,
            filters=self.filters,
            host=self.host,
            id=self.id,
            label=self.label,
            latest_backup=self.latest_backup,
            maintenance_dow=self.maintenance_dow,
            maintenance_time=self.maintenance_time,
            mysql_long_query_time=self.mysql_long_query_time,
            mysql_require_primary_key=self.mysql_require_primary_key,
            mysql_slow_query_log=self.mysql_slow_query_log,
            mysql_sql_modes=self.mysql_sql_modes,
            password=self.password,
            plan=self.plan,
            plan_disk=self.plan_disk,
            plan_ram=self.plan_ram,
            plan_replicas=self.plan_replicas,
            plan_vcpus=self.plan_vcpus,
            port=self.port,
            read_replicas=self.read_replicas,
            redis_eviction_policy=self.redis_eviction_policy,
            region=self.region,
            status=self.status,
            tag=self.tag,
            trusted_ips=self.trusted_ips,
            user=self.user,
            vpc_id=self.vpc_id)


def get_database(filters: Optional[Sequence[pulumi.InputType['GetDatabaseFilterArgs']]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseResult:
    """
    Get information about a Vultr database.

    ## Example Usage

    Get the information for a database by `label`:

    ```python
    import pulumi
    import pulumi_vultr as vultr

    my_database = vultr.get_database(filters=[vultr.GetDatabaseFilterArgs(
        name="label",
        values=["my-database-label"],
    )])
    ```


    :param Sequence[pulumi.InputType['GetDatabaseFilterArgs']] filters: Query parameters for finding databases.
    """
    __args__ = dict()
    __args__['filters'] = filters
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vultr:index/getDatabase:getDatabase', __args__, opts=opts, typ=GetDatabaseResult).value

    return AwaitableGetDatabaseResult(
        cluster_time_zone=pulumi.get(__ret__, 'cluster_time_zone'),
        database_engine=pulumi.get(__ret__, 'database_engine'),
        database_engine_version=pulumi.get(__ret__, 'database_engine_version'),
        date_created=pulumi.get(__ret__, 'date_created'),
        dbname=pulumi.get(__ret__, 'dbname'),
        filters=pulumi.get(__ret__, 'filters'),
        host=pulumi.get(__ret__, 'host'),
        id=pulumi.get(__ret__, 'id'),
        label=pulumi.get(__ret__, 'label'),
        latest_backup=pulumi.get(__ret__, 'latest_backup'),
        maintenance_dow=pulumi.get(__ret__, 'maintenance_dow'),
        maintenance_time=pulumi.get(__ret__, 'maintenance_time'),
        mysql_long_query_time=pulumi.get(__ret__, 'mysql_long_query_time'),
        mysql_require_primary_key=pulumi.get(__ret__, 'mysql_require_primary_key'),
        mysql_slow_query_log=pulumi.get(__ret__, 'mysql_slow_query_log'),
        mysql_sql_modes=pulumi.get(__ret__, 'mysql_sql_modes'),
        password=pulumi.get(__ret__, 'password'),
        plan=pulumi.get(__ret__, 'plan'),
        plan_disk=pulumi.get(__ret__, 'plan_disk'),
        plan_ram=pulumi.get(__ret__, 'plan_ram'),
        plan_replicas=pulumi.get(__ret__, 'plan_replicas'),
        plan_vcpus=pulumi.get(__ret__, 'plan_vcpus'),
        port=pulumi.get(__ret__, 'port'),
        read_replicas=pulumi.get(__ret__, 'read_replicas'),
        redis_eviction_policy=pulumi.get(__ret__, 'redis_eviction_policy'),
        region=pulumi.get(__ret__, 'region'),
        status=pulumi.get(__ret__, 'status'),
        tag=pulumi.get(__ret__, 'tag'),
        trusted_ips=pulumi.get(__ret__, 'trusted_ips'),
        user=pulumi.get(__ret__, 'user'),
        vpc_id=pulumi.get(__ret__, 'vpc_id'))


@_utilities.lift_output_func(get_database)
def get_database_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetDatabaseFilterArgs']]]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatabaseResult]:
    """
    Get information about a Vultr database.

    ## Example Usage

    Get the information for a database by `label`:

    ```python
    import pulumi
    import pulumi_vultr as vultr

    my_database = vultr.get_database(filters=[vultr.GetDatabaseFilterArgs(
        name="label",
        values=["my-database-label"],
    )])
    ```


    :param Sequence[pulumi.InputType['GetDatabaseFilterArgs']] filters: Query parameters for finding databases.
    """
    ...
