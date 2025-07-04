// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-vultr/sdk/v2/go/vultr/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Vultr user associated with your account. This data source provides the name, email, access control list, and API status for a Vultr user associated with your account.
//
// ## Example Usage
//
// Get the information for a user by `email`:
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
//			_, err := vultr.LookupUser(ctx, &vultr.LookupUserArgs{
//				Filters: []vultr.GetUserFilter{
//					{
//						Name: "email",
//						Values: []string{
//							"jdoe@example.com",
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
//
// Get the information for a user by `name`:
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
//			_, err := vultr.LookupUser(ctx, &vultr.LookupUserArgs{
//				Filters: []vultr.GetUserFilter{
//					{
//						Name: "name",
//						Values: []string{
//							"John Doe",
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
func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupUserResult
	err := ctx.Invoke("vultr:index/getUser:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUser.
type LookupUserArgs struct {
	// Query parameters for finding users.
	Filters []GetUserFilter `pulumi:"filters"`
}

// A collection of values returned by getUser.
type LookupUserResult struct {
	// The access control list for the user.
	Acls []string `pulumi:"acls"`
	// Whether API is enabled for the user.
	ApiEnabled bool `pulumi:"apiEnabled"`
	// The email of the user.
	Email   string          `pulumi:"email"`
	Filters []GetUserFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the user.
	Name string `pulumi:"name"`
}

func LookupUserOutput(ctx *pulumi.Context, args LookupUserOutputArgs, opts ...pulumi.InvokeOption) LookupUserResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupUserResultOutput, error) {
			args := v.(LookupUserArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vultr:index/getUser:getUser", args, LookupUserResultOutput{}, options).(LookupUserResultOutput), nil
		}).(LookupUserResultOutput)
}

// A collection of arguments for invoking getUser.
type LookupUserOutputArgs struct {
	// Query parameters for finding users.
	Filters GetUserFilterArrayInput `pulumi:"filters"`
}

func (LookupUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserArgs)(nil)).Elem()
}

// A collection of values returned by getUser.
type LookupUserResultOutput struct{ *pulumi.OutputState }

func (LookupUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserResult)(nil)).Elem()
}

func (o LookupUserResultOutput) ToLookupUserResultOutput() LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) ToLookupUserResultOutputWithContext(ctx context.Context) LookupUserResultOutput {
	return o
}

// The access control list for the user.
func (o LookupUserResultOutput) Acls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserResult) []string { return v.Acls }).(pulumi.StringArrayOutput)
}

// Whether API is enabled for the user.
func (o LookupUserResultOutput) ApiEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupUserResult) bool { return v.ApiEnabled }).(pulumi.BoolOutput)
}

// The email of the user.
func (o LookupUserResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Email }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) Filters() GetUserFilterArrayOutput {
	return o.ApplyT(func(v LookupUserResult) []GetUserFilter { return v.Filters }).(GetUserFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the user.
func (o LookupUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserResultOutput{})
}
