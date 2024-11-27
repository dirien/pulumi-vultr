// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-vultr/sdk/v2/go/vultr/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Vultr Serverless Inference subscription.
//
// ## Example Usage
//
// Get the information for an inference subscription by `label`:
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
//			_, err := vultr.LookupInference(ctx, &vultr.LookupInferenceArgs{
//				Filters: []vultr.GetInferenceFilter{
//					{
//						Name: "label",
//						Values: []string{
//							"my_inference_label",
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
func LookupInference(ctx *pulumi.Context, args *LookupInferenceArgs, opts ...pulumi.InvokeOption) (*LookupInferenceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInferenceResult
	err := ctx.Invoke("vultr:index/getInference:getInference", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInference.
type LookupInferenceArgs struct {
	// Query parameters for finding inference subscriptions.
	Filters []GetInferenceFilter `pulumi:"filters"`
}

// A collection of values returned by getInference.
type LookupInferenceResult struct {
	// The inference subscription's API key for accessing the Vultr Inference API.
	ApiKey string `pulumi:"apiKey"`
	// The date the inference subscription was added to your Vultr account.
	DateCreated string               `pulumi:"dateCreated"`
	Filters     []GetInferenceFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The inference subscription's label.
	Label string            `pulumi:"label"`
	Usage map[string]string `pulumi:"usage"`
}

func LookupInferenceOutput(ctx *pulumi.Context, args LookupInferenceOutputArgs, opts ...pulumi.InvokeOption) LookupInferenceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInferenceResultOutput, error) {
			args := v.(LookupInferenceArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupInferenceResult
			secret, err := ctx.InvokePackageRaw("vultr:index/getInference:getInference", args, &rv, "", opts...)
			if err != nil {
				return LookupInferenceResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupInferenceResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupInferenceResultOutput), nil
			}
			return output, nil
		}).(LookupInferenceResultOutput)
}

// A collection of arguments for invoking getInference.
type LookupInferenceOutputArgs struct {
	// Query parameters for finding inference subscriptions.
	Filters GetInferenceFilterArrayInput `pulumi:"filters"`
}

func (LookupInferenceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInferenceArgs)(nil)).Elem()
}

// A collection of values returned by getInference.
type LookupInferenceResultOutput struct{ *pulumi.OutputState }

func (LookupInferenceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInferenceResult)(nil)).Elem()
}

func (o LookupInferenceResultOutput) ToLookupInferenceResultOutput() LookupInferenceResultOutput {
	return o
}

func (o LookupInferenceResultOutput) ToLookupInferenceResultOutputWithContext(ctx context.Context) LookupInferenceResultOutput {
	return o
}

// The inference subscription's API key for accessing the Vultr Inference API.
func (o LookupInferenceResultOutput) ApiKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInferenceResult) string { return v.ApiKey }).(pulumi.StringOutput)
}

// The date the inference subscription was added to your Vultr account.
func (o LookupInferenceResultOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInferenceResult) string { return v.DateCreated }).(pulumi.StringOutput)
}

func (o LookupInferenceResultOutput) Filters() GetInferenceFilterArrayOutput {
	return o.ApplyT(func(v LookupInferenceResult) []GetInferenceFilter { return v.Filters }).(GetInferenceFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupInferenceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInferenceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The inference subscription's label.
func (o LookupInferenceResultOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInferenceResult) string { return v.Label }).(pulumi.StringOutput)
}

func (o LookupInferenceResultOutput) Usage() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInferenceResult) map[string]string { return v.Usage }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInferenceResultOutput{})
}
