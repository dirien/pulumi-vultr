// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about Object Storage Clusters on Vultr.
//
// ## Example Usage
//
// Get the information for an object storage cluster by `region`:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vultr/sdk/go/vultr"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vultr.GetObjectStorageCluster(ctx, &vultr.GetObjectStorageClusterArgs{
//				Filters: []vultr.GetObjectStorageClusterFilter{
//					{
//						Name: "region",
//						Values: []string{
//							"ewr",
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
func GetObjectStorageCluster(ctx *pulumi.Context, args *GetObjectStorageClusterArgs, opts ...pulumi.InvokeOption) (*GetObjectStorageClusterResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetObjectStorageClusterResult
	err := ctx.Invoke("vultr:index/getObjectStorageCluster:getObjectStorageCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getObjectStorageCluster.
type GetObjectStorageClusterArgs struct {
	// Query parameters for finding operating systems.
	Filters []GetObjectStorageClusterFilter `pulumi:"filters"`
}

// A collection of values returned by getObjectStorageCluster.
type GetObjectStorageClusterResult struct {
	// The Cluster is eligible for Object Storage deployment. (yes or no)
	Deploy  string                          `pulumi:"deploy"`
	Filters []GetObjectStorageClusterFilter `pulumi:"filters"`
	// The cluster hostname.
	Hostname string `pulumi:"hostname"`
	// The identifying cluster ID.
	Id int `pulumi:"id"`
	// The region ID of the object storage cluster.
	Region string `pulumi:"region"`
}

func GetObjectStorageClusterOutput(ctx *pulumi.Context, args GetObjectStorageClusterOutputArgs, opts ...pulumi.InvokeOption) GetObjectStorageClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetObjectStorageClusterResult, error) {
			args := v.(GetObjectStorageClusterArgs)
			r, err := GetObjectStorageCluster(ctx, &args, opts...)
			var s GetObjectStorageClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetObjectStorageClusterResultOutput)
}

// A collection of arguments for invoking getObjectStorageCluster.
type GetObjectStorageClusterOutputArgs struct {
	// Query parameters for finding operating systems.
	Filters GetObjectStorageClusterFilterArrayInput `pulumi:"filters"`
}

func (GetObjectStorageClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetObjectStorageClusterArgs)(nil)).Elem()
}

// A collection of values returned by getObjectStorageCluster.
type GetObjectStorageClusterResultOutput struct{ *pulumi.OutputState }

func (GetObjectStorageClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetObjectStorageClusterResult)(nil)).Elem()
}

func (o GetObjectStorageClusterResultOutput) ToGetObjectStorageClusterResultOutput() GetObjectStorageClusterResultOutput {
	return o
}

func (o GetObjectStorageClusterResultOutput) ToGetObjectStorageClusterResultOutputWithContext(ctx context.Context) GetObjectStorageClusterResultOutput {
	return o
}

// The Cluster is eligible for Object Storage deployment. (yes or no)
func (o GetObjectStorageClusterResultOutput) Deploy() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectStorageClusterResult) string { return v.Deploy }).(pulumi.StringOutput)
}

func (o GetObjectStorageClusterResultOutput) Filters() GetObjectStorageClusterFilterArrayOutput {
	return o.ApplyT(func(v GetObjectStorageClusterResult) []GetObjectStorageClusterFilter { return v.Filters }).(GetObjectStorageClusterFilterArrayOutput)
}

// The cluster hostname.
func (o GetObjectStorageClusterResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectStorageClusterResult) string { return v.Hostname }).(pulumi.StringOutput)
}

// The identifying cluster ID.
func (o GetObjectStorageClusterResultOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v GetObjectStorageClusterResult) int { return v.Id }).(pulumi.IntOutput)
}

// The region ID of the object storage cluster.
func (o GetObjectStorageClusterResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectStorageClusterResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetObjectStorageClusterResultOutput{})
}
