// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"errors"
	"github.com/dirien/pulumi-vultr/sdk/v2/go/vultr/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Vultr database quota resource. This can be used to create, read, and delete quotas for a managed database on your Vultr account.
//
// ## Example Usage
//
// Create a new database quota:
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
//			_, err := vultr.NewDatabaseQuota(ctx, "myDatabaseQuota", &vultr.DatabaseQuotaArgs{
//				DatabaseId:        pulumi.Any(vultr_database.My_database.Id),
//				ClientId:          pulumi.String("my_database_quota"),
//				ConsumerByteRate:  pulumi.Int(3),
//				ProducerByteRate:  pulumi.Int(2),
//				RequestPercentage: pulumi.Int(120),
//				User:              pulumi.String("my_database_user"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DatabaseQuota struct {
	pulumi.CustomResourceState

	// The client ID for the new database quota.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// The consumer byte rate for the new managed database quota.
	ConsumerByteRate pulumi.IntOutput `pulumi:"consumerByteRate"`
	// The managed database ID you want to attach this quota to.
	DatabaseId pulumi.StringOutput `pulumi:"databaseId"`
	// The producer byte rate for the new managed database quota.
	ProducerByteRate pulumi.IntOutput `pulumi:"producerByteRate"`
	// The CPU request percentage for the new managed database quota.
	RequestPercentage pulumi.IntOutput `pulumi:"requestPercentage"`
	// The user for the new managed database quota.
	User pulumi.StringOutput `pulumi:"user"`
}

// NewDatabaseQuota registers a new resource with the given unique name, arguments, and options.
func NewDatabaseQuota(ctx *pulumi.Context,
	name string, args *DatabaseQuotaArgs, opts ...pulumi.ResourceOption) (*DatabaseQuota, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.ConsumerByteRate == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerByteRate'")
	}
	if args.DatabaseId == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseId'")
	}
	if args.ProducerByteRate == nil {
		return nil, errors.New("invalid value for required argument 'ProducerByteRate'")
	}
	if args.RequestPercentage == nil {
		return nil, errors.New("invalid value for required argument 'RequestPercentage'")
	}
	if args.User == nil {
		return nil, errors.New("invalid value for required argument 'User'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DatabaseQuota
	err := ctx.RegisterResource("vultr:index/databaseQuota:DatabaseQuota", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseQuota gets an existing DatabaseQuota resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseQuota(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseQuotaState, opts ...pulumi.ResourceOption) (*DatabaseQuota, error) {
	var resource DatabaseQuota
	err := ctx.ReadResource("vultr:index/databaseQuota:DatabaseQuota", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseQuota resources.
type databaseQuotaState struct {
	// The client ID for the new database quota.
	ClientId *string `pulumi:"clientId"`
	// The consumer byte rate for the new managed database quota.
	ConsumerByteRate *int `pulumi:"consumerByteRate"`
	// The managed database ID you want to attach this quota to.
	DatabaseId *string `pulumi:"databaseId"`
	// The producer byte rate for the new managed database quota.
	ProducerByteRate *int `pulumi:"producerByteRate"`
	// The CPU request percentage for the new managed database quota.
	RequestPercentage *int `pulumi:"requestPercentage"`
	// The user for the new managed database quota.
	User *string `pulumi:"user"`
}

type DatabaseQuotaState struct {
	// The client ID for the new database quota.
	ClientId pulumi.StringPtrInput
	// The consumer byte rate for the new managed database quota.
	ConsumerByteRate pulumi.IntPtrInput
	// The managed database ID you want to attach this quota to.
	DatabaseId pulumi.StringPtrInput
	// The producer byte rate for the new managed database quota.
	ProducerByteRate pulumi.IntPtrInput
	// The CPU request percentage for the new managed database quota.
	RequestPercentage pulumi.IntPtrInput
	// The user for the new managed database quota.
	User pulumi.StringPtrInput
}

func (DatabaseQuotaState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseQuotaState)(nil)).Elem()
}

type databaseQuotaArgs struct {
	// The client ID for the new database quota.
	ClientId string `pulumi:"clientId"`
	// The consumer byte rate for the new managed database quota.
	ConsumerByteRate int `pulumi:"consumerByteRate"`
	// The managed database ID you want to attach this quota to.
	DatabaseId string `pulumi:"databaseId"`
	// The producer byte rate for the new managed database quota.
	ProducerByteRate int `pulumi:"producerByteRate"`
	// The CPU request percentage for the new managed database quota.
	RequestPercentage int `pulumi:"requestPercentage"`
	// The user for the new managed database quota.
	User string `pulumi:"user"`
}

// The set of arguments for constructing a DatabaseQuota resource.
type DatabaseQuotaArgs struct {
	// The client ID for the new database quota.
	ClientId pulumi.StringInput
	// The consumer byte rate for the new managed database quota.
	ConsumerByteRate pulumi.IntInput
	// The managed database ID you want to attach this quota to.
	DatabaseId pulumi.StringInput
	// The producer byte rate for the new managed database quota.
	ProducerByteRate pulumi.IntInput
	// The CPU request percentage for the new managed database quota.
	RequestPercentage pulumi.IntInput
	// The user for the new managed database quota.
	User pulumi.StringInput
}

func (DatabaseQuotaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseQuotaArgs)(nil)).Elem()
}

type DatabaseQuotaInput interface {
	pulumi.Input

	ToDatabaseQuotaOutput() DatabaseQuotaOutput
	ToDatabaseQuotaOutputWithContext(ctx context.Context) DatabaseQuotaOutput
}

func (*DatabaseQuota) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseQuota)(nil)).Elem()
}

func (i *DatabaseQuota) ToDatabaseQuotaOutput() DatabaseQuotaOutput {
	return i.ToDatabaseQuotaOutputWithContext(context.Background())
}

func (i *DatabaseQuota) ToDatabaseQuotaOutputWithContext(ctx context.Context) DatabaseQuotaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseQuotaOutput)
}

// DatabaseQuotaArrayInput is an input type that accepts DatabaseQuotaArray and DatabaseQuotaArrayOutput values.
// You can construct a concrete instance of `DatabaseQuotaArrayInput` via:
//
//	DatabaseQuotaArray{ DatabaseQuotaArgs{...} }
type DatabaseQuotaArrayInput interface {
	pulumi.Input

	ToDatabaseQuotaArrayOutput() DatabaseQuotaArrayOutput
	ToDatabaseQuotaArrayOutputWithContext(context.Context) DatabaseQuotaArrayOutput
}

type DatabaseQuotaArray []DatabaseQuotaInput

func (DatabaseQuotaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseQuota)(nil)).Elem()
}

func (i DatabaseQuotaArray) ToDatabaseQuotaArrayOutput() DatabaseQuotaArrayOutput {
	return i.ToDatabaseQuotaArrayOutputWithContext(context.Background())
}

func (i DatabaseQuotaArray) ToDatabaseQuotaArrayOutputWithContext(ctx context.Context) DatabaseQuotaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseQuotaArrayOutput)
}

// DatabaseQuotaMapInput is an input type that accepts DatabaseQuotaMap and DatabaseQuotaMapOutput values.
// You can construct a concrete instance of `DatabaseQuotaMapInput` via:
//
//	DatabaseQuotaMap{ "key": DatabaseQuotaArgs{...} }
type DatabaseQuotaMapInput interface {
	pulumi.Input

	ToDatabaseQuotaMapOutput() DatabaseQuotaMapOutput
	ToDatabaseQuotaMapOutputWithContext(context.Context) DatabaseQuotaMapOutput
}

type DatabaseQuotaMap map[string]DatabaseQuotaInput

func (DatabaseQuotaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseQuota)(nil)).Elem()
}

func (i DatabaseQuotaMap) ToDatabaseQuotaMapOutput() DatabaseQuotaMapOutput {
	return i.ToDatabaseQuotaMapOutputWithContext(context.Background())
}

func (i DatabaseQuotaMap) ToDatabaseQuotaMapOutputWithContext(ctx context.Context) DatabaseQuotaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseQuotaMapOutput)
}

type DatabaseQuotaOutput struct{ *pulumi.OutputState }

func (DatabaseQuotaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseQuota)(nil)).Elem()
}

func (o DatabaseQuotaOutput) ToDatabaseQuotaOutput() DatabaseQuotaOutput {
	return o
}

func (o DatabaseQuotaOutput) ToDatabaseQuotaOutputWithContext(ctx context.Context) DatabaseQuotaOutput {
	return o
}

// The client ID for the new database quota.
func (o DatabaseQuotaOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseQuota) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// The consumer byte rate for the new managed database quota.
func (o DatabaseQuotaOutput) ConsumerByteRate() pulumi.IntOutput {
	return o.ApplyT(func(v *DatabaseQuota) pulumi.IntOutput { return v.ConsumerByteRate }).(pulumi.IntOutput)
}

// The managed database ID you want to attach this quota to.
func (o DatabaseQuotaOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseQuota) pulumi.StringOutput { return v.DatabaseId }).(pulumi.StringOutput)
}

// The producer byte rate for the new managed database quota.
func (o DatabaseQuotaOutput) ProducerByteRate() pulumi.IntOutput {
	return o.ApplyT(func(v *DatabaseQuota) pulumi.IntOutput { return v.ProducerByteRate }).(pulumi.IntOutput)
}

// The CPU request percentage for the new managed database quota.
func (o DatabaseQuotaOutput) RequestPercentage() pulumi.IntOutput {
	return o.ApplyT(func(v *DatabaseQuota) pulumi.IntOutput { return v.RequestPercentage }).(pulumi.IntOutput)
}

// The user for the new managed database quota.
func (o DatabaseQuotaOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseQuota) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

type DatabaseQuotaArrayOutput struct{ *pulumi.OutputState }

func (DatabaseQuotaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseQuota)(nil)).Elem()
}

func (o DatabaseQuotaArrayOutput) ToDatabaseQuotaArrayOutput() DatabaseQuotaArrayOutput {
	return o
}

func (o DatabaseQuotaArrayOutput) ToDatabaseQuotaArrayOutputWithContext(ctx context.Context) DatabaseQuotaArrayOutput {
	return o
}

func (o DatabaseQuotaArrayOutput) Index(i pulumi.IntInput) DatabaseQuotaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatabaseQuota {
		return vs[0].([]*DatabaseQuota)[vs[1].(int)]
	}).(DatabaseQuotaOutput)
}

type DatabaseQuotaMapOutput struct{ *pulumi.OutputState }

func (DatabaseQuotaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseQuota)(nil)).Elem()
}

func (o DatabaseQuotaMapOutput) ToDatabaseQuotaMapOutput() DatabaseQuotaMapOutput {
	return o
}

func (o DatabaseQuotaMapOutput) ToDatabaseQuotaMapOutputWithContext(ctx context.Context) DatabaseQuotaMapOutput {
	return o
}

func (o DatabaseQuotaMapOutput) MapIndex(k pulumi.StringInput) DatabaseQuotaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatabaseQuota {
		return vs[0].(map[string]*DatabaseQuota)[vs[1].(string)]
	}).(DatabaseQuotaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseQuotaInput)(nil)).Elem(), &DatabaseQuota{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseQuotaArrayInput)(nil)).Elem(), DatabaseQuotaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseQuotaMapInput)(nil)).Elem(), DatabaseQuotaMap{})
	pulumi.RegisterOutputType(DatabaseQuotaOutput{})
	pulumi.RegisterOutputType(DatabaseQuotaArrayOutput{})
	pulumi.RegisterOutputType(DatabaseQuotaMapOutput{})
}
