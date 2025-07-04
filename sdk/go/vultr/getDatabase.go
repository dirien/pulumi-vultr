// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-vultr/sdk/v2/go/vultr/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Vultr database.
//
// ## Example Usage
//
// Get the information for a database by `label`:
//
// ```go
// package main
//
// import (
//
//	"github.com/dirien/pulumi-vultr/sdk/v2/go/vultr"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vultr.LookupDatabase(ctx, &vultr.LookupDatabaseArgs{
//				Filters: []vultr.GetDatabaseFilter{
//					{
//						Name: "label",
//						Values: []string{
//							"my-database-label",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupDatabase(ctx *pulumi.Context, args *LookupDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDatabaseResult
	err := ctx.Invoke("vultr:index/getDatabase:getDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabase.
type LookupDatabaseArgs struct {
	// Query parameters for finding databases.
	Filters []GetDatabaseFilter `pulumi:"filters"`
}

// A collection of values returned by getDatabase.
type LookupDatabaseResult struct {
	// The certificate to authenticate the default user (Kafka engine types only).
	AccessCert string `pulumi:"accessCert"`
	// The private key to authenticate the default user (Kafka engine types only).
	AccessKey string `pulumi:"accessKey"`
	// The configured time zone for the Managed Database in TZ database format.
	ClusterTimeZone string `pulumi:"clusterTimeZone"`
	// The database engine of the managed database.
	DatabaseEngine string `pulumi:"databaseEngine"`
	// The database engine version of the managed database.
	DatabaseEngineVersion string `pulumi:"databaseEngineVersion"`
	// The date the managed database was added to your Vultr account.
	DateCreated string `pulumi:"dateCreated"`
	// The managed database's default logical database.
	Dbname string `pulumi:"dbname"`
	// The configuration value for the data eviction policy on the managed database (Valkey engine types only).
	EvictionPolicy string `pulumi:"evictionPolicy"`
	// An associated list of FerretDB connection credentials (FerretDB + PostgreSQL engine types only).
	FerretdbCredentials map[string]string   `pulumi:"ferretdbCredentials"`
	Filters             []GetDatabaseFilter `pulumi:"filters"`
	// The hostname assigned to the managed database.
	Host string `pulumi:"host"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The managed database's label.
	Label string `pulumi:"label"`
	// The date of the latest backup available on the managed database.
	LatestBackup string `pulumi:"latestBackup"`
	// The preferred maintenance day of week for the managed database.
	MaintenanceDow string `pulumi:"maintenanceDow"`
	// The preferred maintenance time for the managed database.
	MaintenanceTime string `pulumi:"maintenanceTime"`
	// The configuration value for the long query time (in seconds) on the managed database (MySQL engine types only).
	MysqlLongQueryTime int `pulumi:"mysqlLongQueryTime"`
	// The configuration value for whether primary keys are required on the managed database (MySQL engine types only).
	MysqlRequirePrimaryKey bool `pulumi:"mysqlRequirePrimaryKey"`
	// The configuration value for slow query logging on the managed database (MySQL engine types only).
	MysqlSlowQueryLog bool `pulumi:"mysqlSlowQueryLog"`
	// A list of SQL modes currently configured for the managed database (MySQL engine types only).
	MysqlSqlModes []string `pulumi:"mysqlSqlModes"`
	// The password for the managed database's primary admin user.
	Password string `pulumi:"password"`
	// The managed database's plan ID.
	Plan        string `pulumi:"plan"`
	PlanBrokers int    `pulumi:"planBrokers"`
	// The description of the disk(s) on the managed database.
	PlanDisk int `pulumi:"planDisk"`
	// The amount of memory available on the managed database in MB.
	PlanRam int `pulumi:"planRam"`
	// The number of standby nodes available on the managed database.
	PlanReplicas int `pulumi:"planReplicas"`
	// The number of virtual CPUs available on the managed database.
	PlanVcpus int `pulumi:"planVcpus"`
	// The connection port for the managed database.
	Port string `pulumi:"port"`
	// The public hostname assigned to the managed database (VPC-attached only).
	PublicHost string `pulumi:"publicHost"`
	// A list of read replicas attached to the managed database.
	ReadReplicas []GetDatabaseReadReplica `pulumi:"readReplicas"`
	// The region ID of the managed database.
	Region string `pulumi:"region"`
	// The SASL connection port for the managed database (Kafka engine types only).
	SaslPort string `pulumi:"saslPort"`
	// The current status of the managed database (poweroff, rebuilding, rebalancing, configuring, running).
	Status string `pulumi:"status"`
	// The managed database's tag.
	Tag string `pulumi:"tag"`
	// A list of allowed IP addresses for the managed database.
	TrustedIps []string `pulumi:"trustedIps"`
	// The primary admin user for the managed database.
	User string `pulumi:"user"`
	// The ID of the VPC Network attached to the Managed Database.
	VpcId string `pulumi:"vpcId"`
}

func LookupDatabaseOutput(ctx *pulumi.Context, args LookupDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDatabaseResultOutput, error) {
			args := v.(LookupDatabaseArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vultr:index/getDatabase:getDatabase", args, LookupDatabaseResultOutput{}, options).(LookupDatabaseResultOutput), nil
		}).(LookupDatabaseResultOutput)
}

// A collection of arguments for invoking getDatabase.
type LookupDatabaseOutputArgs struct {
	// Query parameters for finding databases.
	Filters GetDatabaseFilterArrayInput `pulumi:"filters"`
}

func (LookupDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseArgs)(nil)).Elem()
}

// A collection of values returned by getDatabase.
type LookupDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseResult)(nil)).Elem()
}

func (o LookupDatabaseResultOutput) ToLookupDatabaseResultOutput() LookupDatabaseResultOutput {
	return o
}

func (o LookupDatabaseResultOutput) ToLookupDatabaseResultOutputWithContext(ctx context.Context) LookupDatabaseResultOutput {
	return o
}

// The certificate to authenticate the default user (Kafka engine types only).
func (o LookupDatabaseResultOutput) AccessCert() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.AccessCert }).(pulumi.StringOutput)
}

// The private key to authenticate the default user (Kafka engine types only).
func (o LookupDatabaseResultOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.AccessKey }).(pulumi.StringOutput)
}

// The configured time zone for the Managed Database in TZ database format.
func (o LookupDatabaseResultOutput) ClusterTimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.ClusterTimeZone }).(pulumi.StringOutput)
}

// The database engine of the managed database.
func (o LookupDatabaseResultOutput) DatabaseEngine() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.DatabaseEngine }).(pulumi.StringOutput)
}

// The database engine version of the managed database.
func (o LookupDatabaseResultOutput) DatabaseEngineVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.DatabaseEngineVersion }).(pulumi.StringOutput)
}

// The date the managed database was added to your Vultr account.
func (o LookupDatabaseResultOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.DateCreated }).(pulumi.StringOutput)
}

// The managed database's default logical database.
func (o LookupDatabaseResultOutput) Dbname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Dbname }).(pulumi.StringOutput)
}

// The configuration value for the data eviction policy on the managed database (Valkey engine types only).
func (o LookupDatabaseResultOutput) EvictionPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.EvictionPolicy }).(pulumi.StringOutput)
}

// An associated list of FerretDB connection credentials (FerretDB + PostgreSQL engine types only).
func (o LookupDatabaseResultOutput) FerretdbCredentials() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatabaseResult) map[string]string { return v.FerretdbCredentials }).(pulumi.StringMapOutput)
}

func (o LookupDatabaseResultOutput) Filters() GetDatabaseFilterArrayOutput {
	return o.ApplyT(func(v LookupDatabaseResult) []GetDatabaseFilter { return v.Filters }).(GetDatabaseFilterArrayOutput)
}

// The hostname assigned to the managed database.
func (o LookupDatabaseResultOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Host }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

// The managed database's label.
func (o LookupDatabaseResultOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Label }).(pulumi.StringOutput)
}

// The date of the latest backup available on the managed database.
func (o LookupDatabaseResultOutput) LatestBackup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.LatestBackup }).(pulumi.StringOutput)
}

// The preferred maintenance day of week for the managed database.
func (o LookupDatabaseResultOutput) MaintenanceDow() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.MaintenanceDow }).(pulumi.StringOutput)
}

// The preferred maintenance time for the managed database.
func (o LookupDatabaseResultOutput) MaintenanceTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.MaintenanceTime }).(pulumi.StringOutput)
}

// The configuration value for the long query time (in seconds) on the managed database (MySQL engine types only).
func (o LookupDatabaseResultOutput) MysqlLongQueryTime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDatabaseResult) int { return v.MysqlLongQueryTime }).(pulumi.IntOutput)
}

// The configuration value for whether primary keys are required on the managed database (MySQL engine types only).
func (o LookupDatabaseResultOutput) MysqlRequirePrimaryKey() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDatabaseResult) bool { return v.MysqlRequirePrimaryKey }).(pulumi.BoolOutput)
}

// The configuration value for slow query logging on the managed database (MySQL engine types only).
func (o LookupDatabaseResultOutput) MysqlSlowQueryLog() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDatabaseResult) bool { return v.MysqlSlowQueryLog }).(pulumi.BoolOutput)
}

// A list of SQL modes currently configured for the managed database (MySQL engine types only).
func (o LookupDatabaseResultOutput) MysqlSqlModes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDatabaseResult) []string { return v.MysqlSqlModes }).(pulumi.StringArrayOutput)
}

// The password for the managed database's primary admin user.
func (o LookupDatabaseResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Password }).(pulumi.StringOutput)
}

// The managed database's plan ID.
func (o LookupDatabaseResultOutput) Plan() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Plan }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) PlanBrokers() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDatabaseResult) int { return v.PlanBrokers }).(pulumi.IntOutput)
}

// The description of the disk(s) on the managed database.
func (o LookupDatabaseResultOutput) PlanDisk() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDatabaseResult) int { return v.PlanDisk }).(pulumi.IntOutput)
}

// The amount of memory available on the managed database in MB.
func (o LookupDatabaseResultOutput) PlanRam() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDatabaseResult) int { return v.PlanRam }).(pulumi.IntOutput)
}

// The number of standby nodes available on the managed database.
func (o LookupDatabaseResultOutput) PlanReplicas() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDatabaseResult) int { return v.PlanReplicas }).(pulumi.IntOutput)
}

// The number of virtual CPUs available on the managed database.
func (o LookupDatabaseResultOutput) PlanVcpus() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDatabaseResult) int { return v.PlanVcpus }).(pulumi.IntOutput)
}

// The connection port for the managed database.
func (o LookupDatabaseResultOutput) Port() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Port }).(pulumi.StringOutput)
}

// The public hostname assigned to the managed database (VPC-attached only).
func (o LookupDatabaseResultOutput) PublicHost() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.PublicHost }).(pulumi.StringOutput)
}

// A list of read replicas attached to the managed database.
func (o LookupDatabaseResultOutput) ReadReplicas() GetDatabaseReadReplicaArrayOutput {
	return o.ApplyT(func(v LookupDatabaseResult) []GetDatabaseReadReplica { return v.ReadReplicas }).(GetDatabaseReadReplicaArrayOutput)
}

// The region ID of the managed database.
func (o LookupDatabaseResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Region }).(pulumi.StringOutput)
}

// The SASL connection port for the managed database (Kafka engine types only).
func (o LookupDatabaseResultOutput) SaslPort() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.SaslPort }).(pulumi.StringOutput)
}

// The current status of the managed database (poweroff, rebuilding, rebalancing, configuring, running).
func (o LookupDatabaseResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Status }).(pulumi.StringOutput)
}

// The managed database's tag.
func (o LookupDatabaseResultOutput) Tag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Tag }).(pulumi.StringOutput)
}

// A list of allowed IP addresses for the managed database.
func (o LookupDatabaseResultOutput) TrustedIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDatabaseResult) []string { return v.TrustedIps }).(pulumi.StringArrayOutput)
}

// The primary admin user for the managed database.
func (o LookupDatabaseResultOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.User }).(pulumi.StringOutput)
}

// The ID of the VPC Network attached to the Managed Database.
func (o LookupDatabaseResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseResultOutput{})
}
