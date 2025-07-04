// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-vultr/sdk/v2/go/vultr/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Vultr virtual file system storage subscription.
//
// ## Example Usage
//
// Get the information for a virtual file system storage subscription by `label`:
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
//			_, err := vultr.LookupVirtualFileSystemStorage(ctx, &vultr.LookupVirtualFileSystemStorageArgs{
//				Filters: []vultr.GetVirtualFileSystemStorageFilter{
//					{
//						Name: "label",
//						Values: []string{
//							"my-vfs-storage-label",
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
func LookupVirtualFileSystemStorage(ctx *pulumi.Context, args *LookupVirtualFileSystemStorageArgs, opts ...pulumi.InvokeOption) (*LookupVirtualFileSystemStorageResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualFileSystemStorageResult
	err := ctx.Invoke("vultr:index/getVirtualFileSystemStorage:getVirtualFileSystemStorage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualFileSystemStorage.
type LookupVirtualFileSystemStorageArgs struct {
	// Query parameters for finding virtual file system storage subscriptions.
	Filters []GetVirtualFileSystemStorageFilter `pulumi:"filters"`
}

// A collection of values returned by getVirtualFileSystemStorage.
type LookupVirtualFileSystemStorageResult struct {
	// A list of instance IDs currently attached to the virtual file system storage.
	AttachedInstances []string `pulumi:"attachedInstances"`
	// A list of attchment states for instances currently attached to the virtual file system storage.
	Attachments []GetVirtualFileSystemStorageAttachment `pulumi:"attachments"`
	// The current pending charges for the virtual file system storage subscription in USD.
	Charges float64 `pulumi:"charges"`
	// The cost per month of the virtual file system storage subscription in USD.
	Cost float64 `pulumi:"cost"`
	// The date the virtual file system storage subscription was added to your Vultr account.
	DateCreated string `pulumi:"dateCreated"`
	// The underlying disk type used by the virtual file system storage subscription.
	DiskType string                              `pulumi:"diskType"`
	Filters  []GetVirtualFileSystemStorageFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The label of the virtual file system storage subscription.
	Label string `pulumi:"label"`
	// The region ID of the virtual file system storage subscription.
	Region string `pulumi:"region"`
	// The size of the virtual file system storage subscription in GB.
	SizeGb int `pulumi:"sizeGb"`
	// The status of the virtual file system storage subscription.
	Status string `pulumi:"status"`
	// A list of tags used on the virtual file system storage subscription.
	Tags []string `pulumi:"tags"`
}

func LookupVirtualFileSystemStorageOutput(ctx *pulumi.Context, args LookupVirtualFileSystemStorageOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualFileSystemStorageResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVirtualFileSystemStorageResultOutput, error) {
			args := v.(LookupVirtualFileSystemStorageArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vultr:index/getVirtualFileSystemStorage:getVirtualFileSystemStorage", args, LookupVirtualFileSystemStorageResultOutput{}, options).(LookupVirtualFileSystemStorageResultOutput), nil
		}).(LookupVirtualFileSystemStorageResultOutput)
}

// A collection of arguments for invoking getVirtualFileSystemStorage.
type LookupVirtualFileSystemStorageOutputArgs struct {
	// Query parameters for finding virtual file system storage subscriptions.
	Filters GetVirtualFileSystemStorageFilterArrayInput `pulumi:"filters"`
}

func (LookupVirtualFileSystemStorageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualFileSystemStorageArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualFileSystemStorage.
type LookupVirtualFileSystemStorageResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualFileSystemStorageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualFileSystemStorageResult)(nil)).Elem()
}

func (o LookupVirtualFileSystemStorageResultOutput) ToLookupVirtualFileSystemStorageResultOutput() LookupVirtualFileSystemStorageResultOutput {
	return o
}

func (o LookupVirtualFileSystemStorageResultOutput) ToLookupVirtualFileSystemStorageResultOutputWithContext(ctx context.Context) LookupVirtualFileSystemStorageResultOutput {
	return o
}

// A list of instance IDs currently attached to the virtual file system storage.
func (o LookupVirtualFileSystemStorageResultOutput) AttachedInstances() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualFileSystemStorageResult) []string { return v.AttachedInstances }).(pulumi.StringArrayOutput)
}

// A list of attchment states for instances currently attached to the virtual file system storage.
func (o LookupVirtualFileSystemStorageResultOutput) Attachments() GetVirtualFileSystemStorageAttachmentArrayOutput {
	return o.ApplyT(func(v LookupVirtualFileSystemStorageResult) []GetVirtualFileSystemStorageAttachment {
		return v.Attachments
	}).(GetVirtualFileSystemStorageAttachmentArrayOutput)
}

// The current pending charges for the virtual file system storage subscription in USD.
func (o LookupVirtualFileSystemStorageResultOutput) Charges() pulumi.Float64Output {
	return o.ApplyT(func(v LookupVirtualFileSystemStorageResult) float64 { return v.Charges }).(pulumi.Float64Output)
}

// The cost per month of the virtual file system storage subscription in USD.
func (o LookupVirtualFileSystemStorageResultOutput) Cost() pulumi.Float64Output {
	return o.ApplyT(func(v LookupVirtualFileSystemStorageResult) float64 { return v.Cost }).(pulumi.Float64Output)
}

// The date the virtual file system storage subscription was added to your Vultr account.
func (o LookupVirtualFileSystemStorageResultOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualFileSystemStorageResult) string { return v.DateCreated }).(pulumi.StringOutput)
}

// The underlying disk type used by the virtual file system storage subscription.
func (o LookupVirtualFileSystemStorageResultOutput) DiskType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualFileSystemStorageResult) string { return v.DiskType }).(pulumi.StringOutput)
}

func (o LookupVirtualFileSystemStorageResultOutput) Filters() GetVirtualFileSystemStorageFilterArrayOutput {
	return o.ApplyT(func(v LookupVirtualFileSystemStorageResult) []GetVirtualFileSystemStorageFilter { return v.Filters }).(GetVirtualFileSystemStorageFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVirtualFileSystemStorageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualFileSystemStorageResult) string { return v.Id }).(pulumi.StringOutput)
}

// The label of the virtual file system storage subscription.
func (o LookupVirtualFileSystemStorageResultOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualFileSystemStorageResult) string { return v.Label }).(pulumi.StringOutput)
}

// The region ID of the virtual file system storage subscription.
func (o LookupVirtualFileSystemStorageResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualFileSystemStorageResult) string { return v.Region }).(pulumi.StringOutput)
}

// The size of the virtual file system storage subscription in GB.
func (o LookupVirtualFileSystemStorageResultOutput) SizeGb() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVirtualFileSystemStorageResult) int { return v.SizeGb }).(pulumi.IntOutput)
}

// The status of the virtual file system storage subscription.
func (o LookupVirtualFileSystemStorageResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualFileSystemStorageResult) string { return v.Status }).(pulumi.StringOutput)
}

// A list of tags used on the virtual file system storage subscription.
func (o LookupVirtualFileSystemStorageResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualFileSystemStorageResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualFileSystemStorageResultOutput{})
}
