// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DatabaseReplica struct {
	pulumi.CustomResourceState

	ClusterTimeZone        pulumi.StringOutput      `pulumi:"clusterTimeZone"`
	DatabaseEngine         pulumi.StringOutput      `pulumi:"databaseEngine"`
	DatabaseEngineVersion  pulumi.StringOutput      `pulumi:"databaseEngineVersion"`
	DatabaseId             pulumi.StringOutput      `pulumi:"databaseId"`
	DateCreated            pulumi.StringOutput      `pulumi:"dateCreated"`
	Dbname                 pulumi.StringOutput      `pulumi:"dbname"`
	Host                   pulumi.StringOutput      `pulumi:"host"`
	Label                  pulumi.StringOutput      `pulumi:"label"`
	LatestBackup           pulumi.StringOutput      `pulumi:"latestBackup"`
	MaintenanceDow         pulumi.StringOutput      `pulumi:"maintenanceDow"`
	MaintenanceTime        pulumi.StringOutput      `pulumi:"maintenanceTime"`
	MysqlLongQueryTime     pulumi.IntOutput         `pulumi:"mysqlLongQueryTime"`
	MysqlRequirePrimaryKey pulumi.BoolOutput        `pulumi:"mysqlRequirePrimaryKey"`
	MysqlSlowQueryLog      pulumi.BoolOutput        `pulumi:"mysqlSlowQueryLog"`
	MysqlSqlModes          pulumi.StringArrayOutput `pulumi:"mysqlSqlModes"`
	Password               pulumi.StringOutput      `pulumi:"password"`
	Plan                   pulumi.StringOutput      `pulumi:"plan"`
	PlanDisk               pulumi.IntOutput         `pulumi:"planDisk"`
	PlanRam                pulumi.IntOutput         `pulumi:"planRam"`
	PlanReplicas           pulumi.IntOutput         `pulumi:"planReplicas"`
	PlanVcpus              pulumi.IntOutput         `pulumi:"planVcpus"`
	Port                   pulumi.StringOutput      `pulumi:"port"`
	RedisEvictionPolicy    pulumi.StringOutput      `pulumi:"redisEvictionPolicy"`
	Region                 pulumi.StringOutput      `pulumi:"region"`
	Status                 pulumi.StringOutput      `pulumi:"status"`
	Tag                    pulumi.StringOutput      `pulumi:"tag"`
	TrustedIps             pulumi.StringArrayOutput `pulumi:"trustedIps"`
	User                   pulumi.StringOutput      `pulumi:"user"`
}

// NewDatabaseReplica registers a new resource with the given unique name, arguments, and options.
func NewDatabaseReplica(ctx *pulumi.Context,
	name string, args *DatabaseReplicaArgs, opts ...pulumi.ResourceOption) (*DatabaseReplica, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseId == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseId'")
	}
	if args.Label == nil {
		return nil, errors.New("invalid value for required argument 'Label'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DatabaseReplica
	err := ctx.RegisterResource("vultr:index/databaseReplica:DatabaseReplica", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseReplica gets an existing DatabaseReplica resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseReplica(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseReplicaState, opts ...pulumi.ResourceOption) (*DatabaseReplica, error) {
	var resource DatabaseReplica
	err := ctx.ReadResource("vultr:index/databaseReplica:DatabaseReplica", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseReplica resources.
type databaseReplicaState struct {
	ClusterTimeZone        *string  `pulumi:"clusterTimeZone"`
	DatabaseEngine         *string  `pulumi:"databaseEngine"`
	DatabaseEngineVersion  *string  `pulumi:"databaseEngineVersion"`
	DatabaseId             *string  `pulumi:"databaseId"`
	DateCreated            *string  `pulumi:"dateCreated"`
	Dbname                 *string  `pulumi:"dbname"`
	Host                   *string  `pulumi:"host"`
	Label                  *string  `pulumi:"label"`
	LatestBackup           *string  `pulumi:"latestBackup"`
	MaintenanceDow         *string  `pulumi:"maintenanceDow"`
	MaintenanceTime        *string  `pulumi:"maintenanceTime"`
	MysqlLongQueryTime     *int     `pulumi:"mysqlLongQueryTime"`
	MysqlRequirePrimaryKey *bool    `pulumi:"mysqlRequirePrimaryKey"`
	MysqlSlowQueryLog      *bool    `pulumi:"mysqlSlowQueryLog"`
	MysqlSqlModes          []string `pulumi:"mysqlSqlModes"`
	Password               *string  `pulumi:"password"`
	Plan                   *string  `pulumi:"plan"`
	PlanDisk               *int     `pulumi:"planDisk"`
	PlanRam                *int     `pulumi:"planRam"`
	PlanReplicas           *int     `pulumi:"planReplicas"`
	PlanVcpus              *int     `pulumi:"planVcpus"`
	Port                   *string  `pulumi:"port"`
	RedisEvictionPolicy    *string  `pulumi:"redisEvictionPolicy"`
	Region                 *string  `pulumi:"region"`
	Status                 *string  `pulumi:"status"`
	Tag                    *string  `pulumi:"tag"`
	TrustedIps             []string `pulumi:"trustedIps"`
	User                   *string  `pulumi:"user"`
}

type DatabaseReplicaState struct {
	ClusterTimeZone        pulumi.StringPtrInput
	DatabaseEngine         pulumi.StringPtrInput
	DatabaseEngineVersion  pulumi.StringPtrInput
	DatabaseId             pulumi.StringPtrInput
	DateCreated            pulumi.StringPtrInput
	Dbname                 pulumi.StringPtrInput
	Host                   pulumi.StringPtrInput
	Label                  pulumi.StringPtrInput
	LatestBackup           pulumi.StringPtrInput
	MaintenanceDow         pulumi.StringPtrInput
	MaintenanceTime        pulumi.StringPtrInput
	MysqlLongQueryTime     pulumi.IntPtrInput
	MysqlRequirePrimaryKey pulumi.BoolPtrInput
	MysqlSlowQueryLog      pulumi.BoolPtrInput
	MysqlSqlModes          pulumi.StringArrayInput
	Password               pulumi.StringPtrInput
	Plan                   pulumi.StringPtrInput
	PlanDisk               pulumi.IntPtrInput
	PlanRam                pulumi.IntPtrInput
	PlanReplicas           pulumi.IntPtrInput
	PlanVcpus              pulumi.IntPtrInput
	Port                   pulumi.StringPtrInput
	RedisEvictionPolicy    pulumi.StringPtrInput
	Region                 pulumi.StringPtrInput
	Status                 pulumi.StringPtrInput
	Tag                    pulumi.StringPtrInput
	TrustedIps             pulumi.StringArrayInput
	User                   pulumi.StringPtrInput
}

func (DatabaseReplicaState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseReplicaState)(nil)).Elem()
}

type databaseReplicaArgs struct {
	DatabaseId             string   `pulumi:"databaseId"`
	Label                  string   `pulumi:"label"`
	MysqlLongQueryTime     *int     `pulumi:"mysqlLongQueryTime"`
	MysqlRequirePrimaryKey *bool    `pulumi:"mysqlRequirePrimaryKey"`
	MysqlSlowQueryLog      *bool    `pulumi:"mysqlSlowQueryLog"`
	MysqlSqlModes          []string `pulumi:"mysqlSqlModes"`
	PlanDisk               *int     `pulumi:"planDisk"`
	RedisEvictionPolicy    *string  `pulumi:"redisEvictionPolicy"`
	Region                 string   `pulumi:"region"`
	Tag                    *string  `pulumi:"tag"`
	TrustedIps             []string `pulumi:"trustedIps"`
}

// The set of arguments for constructing a DatabaseReplica resource.
type DatabaseReplicaArgs struct {
	DatabaseId             pulumi.StringInput
	Label                  pulumi.StringInput
	MysqlLongQueryTime     pulumi.IntPtrInput
	MysqlRequirePrimaryKey pulumi.BoolPtrInput
	MysqlSlowQueryLog      pulumi.BoolPtrInput
	MysqlSqlModes          pulumi.StringArrayInput
	PlanDisk               pulumi.IntPtrInput
	RedisEvictionPolicy    pulumi.StringPtrInput
	Region                 pulumi.StringInput
	Tag                    pulumi.StringPtrInput
	TrustedIps             pulumi.StringArrayInput
}

func (DatabaseReplicaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseReplicaArgs)(nil)).Elem()
}

type DatabaseReplicaInput interface {
	pulumi.Input

	ToDatabaseReplicaOutput() DatabaseReplicaOutput
	ToDatabaseReplicaOutputWithContext(ctx context.Context) DatabaseReplicaOutput
}

func (*DatabaseReplica) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseReplica)(nil)).Elem()
}

func (i *DatabaseReplica) ToDatabaseReplicaOutput() DatabaseReplicaOutput {
	return i.ToDatabaseReplicaOutputWithContext(context.Background())
}

func (i *DatabaseReplica) ToDatabaseReplicaOutputWithContext(ctx context.Context) DatabaseReplicaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseReplicaOutput)
}

// DatabaseReplicaArrayInput is an input type that accepts DatabaseReplicaArray and DatabaseReplicaArrayOutput values.
// You can construct a concrete instance of `DatabaseReplicaArrayInput` via:
//
//	DatabaseReplicaArray{ DatabaseReplicaArgs{...} }
type DatabaseReplicaArrayInput interface {
	pulumi.Input

	ToDatabaseReplicaArrayOutput() DatabaseReplicaArrayOutput
	ToDatabaseReplicaArrayOutputWithContext(context.Context) DatabaseReplicaArrayOutput
}

type DatabaseReplicaArray []DatabaseReplicaInput

func (DatabaseReplicaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseReplica)(nil)).Elem()
}

func (i DatabaseReplicaArray) ToDatabaseReplicaArrayOutput() DatabaseReplicaArrayOutput {
	return i.ToDatabaseReplicaArrayOutputWithContext(context.Background())
}

func (i DatabaseReplicaArray) ToDatabaseReplicaArrayOutputWithContext(ctx context.Context) DatabaseReplicaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseReplicaArrayOutput)
}

// DatabaseReplicaMapInput is an input type that accepts DatabaseReplicaMap and DatabaseReplicaMapOutput values.
// You can construct a concrete instance of `DatabaseReplicaMapInput` via:
//
//	DatabaseReplicaMap{ "key": DatabaseReplicaArgs{...} }
type DatabaseReplicaMapInput interface {
	pulumi.Input

	ToDatabaseReplicaMapOutput() DatabaseReplicaMapOutput
	ToDatabaseReplicaMapOutputWithContext(context.Context) DatabaseReplicaMapOutput
}

type DatabaseReplicaMap map[string]DatabaseReplicaInput

func (DatabaseReplicaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseReplica)(nil)).Elem()
}

func (i DatabaseReplicaMap) ToDatabaseReplicaMapOutput() DatabaseReplicaMapOutput {
	return i.ToDatabaseReplicaMapOutputWithContext(context.Background())
}

func (i DatabaseReplicaMap) ToDatabaseReplicaMapOutputWithContext(ctx context.Context) DatabaseReplicaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseReplicaMapOutput)
}

type DatabaseReplicaOutput struct{ *pulumi.OutputState }

func (DatabaseReplicaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseReplica)(nil)).Elem()
}

func (o DatabaseReplicaOutput) ToDatabaseReplicaOutput() DatabaseReplicaOutput {
	return o
}

func (o DatabaseReplicaOutput) ToDatabaseReplicaOutputWithContext(ctx context.Context) DatabaseReplicaOutput {
	return o
}

func (o DatabaseReplicaOutput) ClusterTimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.StringOutput { return v.ClusterTimeZone }).(pulumi.StringOutput)
}

func (o DatabaseReplicaOutput) DatabaseEngine() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.StringOutput { return v.DatabaseEngine }).(pulumi.StringOutput)
}

func (o DatabaseReplicaOutput) DatabaseEngineVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.StringOutput { return v.DatabaseEngineVersion }).(pulumi.StringOutput)
}

func (o DatabaseReplicaOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.StringOutput { return v.DatabaseId }).(pulumi.StringOutput)
}

func (o DatabaseReplicaOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.StringOutput { return v.DateCreated }).(pulumi.StringOutput)
}

func (o DatabaseReplicaOutput) Dbname() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.StringOutput { return v.Dbname }).(pulumi.StringOutput)
}

func (o DatabaseReplicaOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

func (o DatabaseReplicaOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.StringOutput { return v.Label }).(pulumi.StringOutput)
}

func (o DatabaseReplicaOutput) LatestBackup() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.StringOutput { return v.LatestBackup }).(pulumi.StringOutput)
}

func (o DatabaseReplicaOutput) MaintenanceDow() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.StringOutput { return v.MaintenanceDow }).(pulumi.StringOutput)
}

func (o DatabaseReplicaOutput) MaintenanceTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.StringOutput { return v.MaintenanceTime }).(pulumi.StringOutput)
}

func (o DatabaseReplicaOutput) MysqlLongQueryTime() pulumi.IntOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.IntOutput { return v.MysqlLongQueryTime }).(pulumi.IntOutput)
}

func (o DatabaseReplicaOutput) MysqlRequirePrimaryKey() pulumi.BoolOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.BoolOutput { return v.MysqlRequirePrimaryKey }).(pulumi.BoolOutput)
}

func (o DatabaseReplicaOutput) MysqlSlowQueryLog() pulumi.BoolOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.BoolOutput { return v.MysqlSlowQueryLog }).(pulumi.BoolOutput)
}

func (o DatabaseReplicaOutput) MysqlSqlModes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.StringArrayOutput { return v.MysqlSqlModes }).(pulumi.StringArrayOutput)
}

func (o DatabaseReplicaOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

func (o DatabaseReplicaOutput) Plan() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.StringOutput { return v.Plan }).(pulumi.StringOutput)
}

func (o DatabaseReplicaOutput) PlanDisk() pulumi.IntOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.IntOutput { return v.PlanDisk }).(pulumi.IntOutput)
}

func (o DatabaseReplicaOutput) PlanRam() pulumi.IntOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.IntOutput { return v.PlanRam }).(pulumi.IntOutput)
}

func (o DatabaseReplicaOutput) PlanReplicas() pulumi.IntOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.IntOutput { return v.PlanReplicas }).(pulumi.IntOutput)
}

func (o DatabaseReplicaOutput) PlanVcpus() pulumi.IntOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.IntOutput { return v.PlanVcpus }).(pulumi.IntOutput)
}

func (o DatabaseReplicaOutput) Port() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.StringOutput { return v.Port }).(pulumi.StringOutput)
}

func (o DatabaseReplicaOutput) RedisEvictionPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.StringOutput { return v.RedisEvictionPolicy }).(pulumi.StringOutput)
}

func (o DatabaseReplicaOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o DatabaseReplicaOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o DatabaseReplicaOutput) Tag() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.StringOutput { return v.Tag }).(pulumi.StringOutput)
}

func (o DatabaseReplicaOutput) TrustedIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.StringArrayOutput { return v.TrustedIps }).(pulumi.StringArrayOutput)
}

func (o DatabaseReplicaOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseReplica) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

type DatabaseReplicaArrayOutput struct{ *pulumi.OutputState }

func (DatabaseReplicaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseReplica)(nil)).Elem()
}

func (o DatabaseReplicaArrayOutput) ToDatabaseReplicaArrayOutput() DatabaseReplicaArrayOutput {
	return o
}

func (o DatabaseReplicaArrayOutput) ToDatabaseReplicaArrayOutputWithContext(ctx context.Context) DatabaseReplicaArrayOutput {
	return o
}

func (o DatabaseReplicaArrayOutput) Index(i pulumi.IntInput) DatabaseReplicaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatabaseReplica {
		return vs[0].([]*DatabaseReplica)[vs[1].(int)]
	}).(DatabaseReplicaOutput)
}

type DatabaseReplicaMapOutput struct{ *pulumi.OutputState }

func (DatabaseReplicaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseReplica)(nil)).Elem()
}

func (o DatabaseReplicaMapOutput) ToDatabaseReplicaMapOutput() DatabaseReplicaMapOutput {
	return o
}

func (o DatabaseReplicaMapOutput) ToDatabaseReplicaMapOutputWithContext(ctx context.Context) DatabaseReplicaMapOutput {
	return o
}

func (o DatabaseReplicaMapOutput) MapIndex(k pulumi.StringInput) DatabaseReplicaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatabaseReplica {
		return vs[0].(map[string]*DatabaseReplica)[vs[1].(string)]
	}).(DatabaseReplicaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseReplicaInput)(nil)).Elem(), &DatabaseReplica{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseReplicaArrayInput)(nil)).Elem(), DatabaseReplicaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseReplicaMapInput)(nil)).Elem(), DatabaseReplicaMap{})
	pulumi.RegisterOutputType(DatabaseReplicaOutput{})
	pulumi.RegisterOutputType(DatabaseReplicaArrayOutput{})
	pulumi.RegisterOutputType(DatabaseReplicaMapOutput{})
}