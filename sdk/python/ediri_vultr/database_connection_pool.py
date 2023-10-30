# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['DatabaseConnectionPoolArgs', 'DatabaseConnectionPool']

@pulumi.input_type
class DatabaseConnectionPoolArgs:
    def __init__(__self__, *,
                 database: pulumi.Input[str],
                 database_id: pulumi.Input[str],
                 mode: pulumi.Input[str],
                 size: pulumi.Input[int],
                 username: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DatabaseConnectionPool resource.
        :param pulumi.Input[str] database: The logical database to use for the new managed database connection pool.
        :param pulumi.Input[str] database_id: The managed database ID you want to attach this connection pool to.
        :param pulumi.Input[str] mode: The mode to configure for the new managed database connection pool (`session`, `transaction`, `statement`).
        :param pulumi.Input[int] size: The size of the new managed database connection pool.
        :param pulumi.Input[str] username: The database user to use for the new managed database connection pool.
        :param pulumi.Input[str] name: The name of the new managed database connection pool.
        """
        pulumi.set(__self__, "database", database)
        pulumi.set(__self__, "database_id", database_id)
        pulumi.set(__self__, "mode", mode)
        pulumi.set(__self__, "size", size)
        pulumi.set(__self__, "username", username)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def database(self) -> pulumi.Input[str]:
        """
        The logical database to use for the new managed database connection pool.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: pulumi.Input[str]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> pulumi.Input[str]:
        """
        The managed database ID you want to attach this connection pool to.
        """
        return pulumi.get(self, "database_id")

    @database_id.setter
    def database_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "database_id", value)

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Input[str]:
        """
        The mode to configure for the new managed database connection pool (`session`, `transaction`, `statement`).
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def size(self) -> pulumi.Input[int]:
        """
        The size of the new managed database connection pool.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: pulumi.Input[int]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        """
        The database user to use for the new managed database connection pool.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the new managed database connection pool.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _DatabaseConnectionPoolState:
    def __init__(__self__, *,
                 database: Optional[pulumi.Input[str]] = None,
                 database_id: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DatabaseConnectionPool resources.
        :param pulumi.Input[str] database: The logical database to use for the new managed database connection pool.
        :param pulumi.Input[str] database_id: The managed database ID you want to attach this connection pool to.
        :param pulumi.Input[str] mode: The mode to configure for the new managed database connection pool (`session`, `transaction`, `statement`).
        :param pulumi.Input[str] name: The name of the new managed database connection pool.
        :param pulumi.Input[int] size: The size of the new managed database connection pool.
        :param pulumi.Input[str] username: The database user to use for the new managed database connection pool.
        """
        if database is not None:
            pulumi.set(__self__, "database", database)
        if database_id is not None:
            pulumi.set(__self__, "database_id", database_id)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def database(self) -> Optional[pulumi.Input[str]]:
        """
        The logical database to use for the new managed database connection pool.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> Optional[pulumi.Input[str]]:
        """
        The managed database ID you want to attach this connection pool to.
        """
        return pulumi.get(self, "database_id")

    @database_id.setter
    def database_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_id", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        The mode to configure for the new managed database connection pool (`session`, `transaction`, `statement`).
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the new managed database connection pool.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[int]]:
        """
        The size of the new managed database connection pool.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The database user to use for the new managed database connection pool.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class DatabaseConnectionPool(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 database_id: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Vultr database connection pool resource. This can be used to create, read, modify, and delete connection pools for a PostgreSQL managed database on your Vultr account.

        ## Example Usage

        Create a new database connection pool:

        ```python
        import pulumi
        import ediri_vultr as vultr

        my_database_connection_pool = vultr.DatabaseConnectionPool("myDatabaseConnectionPool",
            database_id=vultr_database["my_database"]["id"],
            database="defaultdb",
            username="vultradmin",
            mode="transaction",
            size=3)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database: The logical database to use for the new managed database connection pool.
        :param pulumi.Input[str] database_id: The managed database ID you want to attach this connection pool to.
        :param pulumi.Input[str] mode: The mode to configure for the new managed database connection pool (`session`, `transaction`, `statement`).
        :param pulumi.Input[str] name: The name of the new managed database connection pool.
        :param pulumi.Input[int] size: The size of the new managed database connection pool.
        :param pulumi.Input[str] username: The database user to use for the new managed database connection pool.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatabaseConnectionPoolArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Vultr database connection pool resource. This can be used to create, read, modify, and delete connection pools for a PostgreSQL managed database on your Vultr account.

        ## Example Usage

        Create a new database connection pool:

        ```python
        import pulumi
        import ediri_vultr as vultr

        my_database_connection_pool = vultr.DatabaseConnectionPool("myDatabaseConnectionPool",
            database_id=vultr_database["my_database"]["id"],
            database="defaultdb",
            username="vultradmin",
            mode="transaction",
            size=3)
        ```

        :param str resource_name: The name of the resource.
        :param DatabaseConnectionPoolArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatabaseConnectionPoolArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 database_id: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DatabaseConnectionPoolArgs.__new__(DatabaseConnectionPoolArgs)

            if database is None and not opts.urn:
                raise TypeError("Missing required property 'database'")
            __props__.__dict__["database"] = database
            if database_id is None and not opts.urn:
                raise TypeError("Missing required property 'database_id'")
            __props__.__dict__["database_id"] = database_id
            if mode is None and not opts.urn:
                raise TypeError("Missing required property 'mode'")
            __props__.__dict__["mode"] = mode
            __props__.__dict__["name"] = name
            if size is None and not opts.urn:
                raise TypeError("Missing required property 'size'")
            __props__.__dict__["size"] = size
            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__.__dict__["username"] = username
        super(DatabaseConnectionPool, __self__).__init__(
            'vultr:index/databaseConnectionPool:DatabaseConnectionPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            database: Optional[pulumi.Input[str]] = None,
            database_id: Optional[pulumi.Input[str]] = None,
            mode: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            size: Optional[pulumi.Input[int]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'DatabaseConnectionPool':
        """
        Get an existing DatabaseConnectionPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database: The logical database to use for the new managed database connection pool.
        :param pulumi.Input[str] database_id: The managed database ID you want to attach this connection pool to.
        :param pulumi.Input[str] mode: The mode to configure for the new managed database connection pool (`session`, `transaction`, `statement`).
        :param pulumi.Input[str] name: The name of the new managed database connection pool.
        :param pulumi.Input[int] size: The size of the new managed database connection pool.
        :param pulumi.Input[str] username: The database user to use for the new managed database connection pool.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DatabaseConnectionPoolState.__new__(_DatabaseConnectionPoolState)

        __props__.__dict__["database"] = database
        __props__.__dict__["database_id"] = database_id
        __props__.__dict__["mode"] = mode
        __props__.__dict__["name"] = name
        __props__.__dict__["size"] = size
        __props__.__dict__["username"] = username
        return DatabaseConnectionPool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def database(self) -> pulumi.Output[str]:
        """
        The logical database to use for the new managed database connection pool.
        """
        return pulumi.get(self, "database")

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> pulumi.Output[str]:
        """
        The managed database ID you want to attach this connection pool to.
        """
        return pulumi.get(self, "database_id")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[str]:
        """
        The mode to configure for the new managed database connection pool (`session`, `transaction`, `statement`).
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the new managed database connection pool.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[int]:
        """
        The size of the new managed database connection pool.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        The database user to use for the new managed database connection pool.
        """
        return pulumi.get(self, "username")

