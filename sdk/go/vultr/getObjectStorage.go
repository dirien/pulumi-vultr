// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about an Object Storage subscription on Vultr.
//
// ## Example Usage
//
// Get the information for an object storage subscription by `label`:
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
//			_, err := vultr.LookupObjectStorage(ctx, &vultr.LookupObjectStorageArgs{
//				Filters: []vultr.GetObjectStorageFilter{
//					{
//						Name: "label",
//						Values: []string{
//							"my-s3",
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
func LookupObjectStorage(ctx *pulumi.Context, args *LookupObjectStorageArgs, opts ...pulumi.InvokeOption) (*LookupObjectStorageResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupObjectStorageResult
	err := ctx.Invoke("vultr:index/getObjectStorage:getObjectStorage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getObjectStorage.
type LookupObjectStorageArgs struct {
	// Query parameters for finding operating systems.
	Filters []GetObjectStorageFilter `pulumi:"filters"`
}

// A collection of values returned by getObjectStorage.
type LookupObjectStorageResult struct {
	// The identifying cluster ID.
	ClusterId int `pulumi:"clusterId"`
	// Date of creation for the object storage subscription.
	DateCreated string                   `pulumi:"dateCreated"`
	Filters     []GetObjectStorageFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The label of the object storage subscription.
	Label string `pulumi:"label"`
	// The location which this subscription resides in.
	Location string `pulumi:"location"`
	// The region ID of the object storage subscription.
	Region string `pulumi:"region"`
	// Your access key.
	S3AccessKey string `pulumi:"s3AccessKey"`
	// The hostname for this subscription.
	S3Hostname string `pulumi:"s3Hostname"`
	// Your secret key.
	S3SecretKey string `pulumi:"s3SecretKey"`
	// Current status of this object storage subscription.
	Status string `pulumi:"status"`
}

func LookupObjectStorageOutput(ctx *pulumi.Context, args LookupObjectStorageOutputArgs, opts ...pulumi.InvokeOption) LookupObjectStorageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupObjectStorageResult, error) {
			args := v.(LookupObjectStorageArgs)
			r, err := LookupObjectStorage(ctx, &args, opts...)
			var s LookupObjectStorageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupObjectStorageResultOutput)
}

// A collection of arguments for invoking getObjectStorage.
type LookupObjectStorageOutputArgs struct {
	// Query parameters for finding operating systems.
	Filters GetObjectStorageFilterArrayInput `pulumi:"filters"`
}

func (LookupObjectStorageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupObjectStorageArgs)(nil)).Elem()
}

// A collection of values returned by getObjectStorage.
type LookupObjectStorageResultOutput struct{ *pulumi.OutputState }

func (LookupObjectStorageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupObjectStorageResult)(nil)).Elem()
}

func (o LookupObjectStorageResultOutput) ToLookupObjectStorageResultOutput() LookupObjectStorageResultOutput {
	return o
}

func (o LookupObjectStorageResultOutput) ToLookupObjectStorageResultOutputWithContext(ctx context.Context) LookupObjectStorageResultOutput {
	return o
}

// The identifying cluster ID.
func (o LookupObjectStorageResultOutput) ClusterId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupObjectStorageResult) int { return v.ClusterId }).(pulumi.IntOutput)
}

// Date of creation for the object storage subscription.
func (o LookupObjectStorageResultOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectStorageResult) string { return v.DateCreated }).(pulumi.StringOutput)
}

func (o LookupObjectStorageResultOutput) Filters() GetObjectStorageFilterArrayOutput {
	return o.ApplyT(func(v LookupObjectStorageResult) []GetObjectStorageFilter { return v.Filters }).(GetObjectStorageFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupObjectStorageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectStorageResult) string { return v.Id }).(pulumi.StringOutput)
}

// The label of the object storage subscription.
func (o LookupObjectStorageResultOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectStorageResult) string { return v.Label }).(pulumi.StringOutput)
}

// The location which this subscription resides in.
func (o LookupObjectStorageResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectStorageResult) string { return v.Location }).(pulumi.StringOutput)
}

// The region ID of the object storage subscription.
func (o LookupObjectStorageResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectStorageResult) string { return v.Region }).(pulumi.StringOutput)
}

// Your access key.
func (o LookupObjectStorageResultOutput) S3AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectStorageResult) string { return v.S3AccessKey }).(pulumi.StringOutput)
}

// The hostname for this subscription.
func (o LookupObjectStorageResultOutput) S3Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectStorageResult) string { return v.S3Hostname }).(pulumi.StringOutput)
}

// Your secret key.
func (o LookupObjectStorageResultOutput) S3SecretKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectStorageResult) string { return v.S3SecretKey }).(pulumi.StringOutput)
}

// Current status of this object storage subscription.
func (o LookupObjectStorageResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectStorageResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupObjectStorageResultOutput{})
}
