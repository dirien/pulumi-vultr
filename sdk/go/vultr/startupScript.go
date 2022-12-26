// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Vultr Startup Script resource. This can be used to create, read, modify, and delete Startup Scripts.
//
// ## Example Usage
//
// # Create a new Startup Script
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
//			_, err := vultr.NewStartupScript(ctx, "myScript", &vultr.StartupScriptArgs{
//				Script: pulumi.String("ZWNobyAkUEFUSAo="),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Startup Scripts can be imported using the Startup Scripts `ID`, e.g.
//
// ```sh
//
//	$ pulumi import vultr:index/startupScript:StartupScript my_script ff8f36a8-eb86-4b8d-8667-b9d5459b6390
//
// ```
type StartupScript struct {
	pulumi.CustomResourceState

	// Date the script was created.
	DateCreated pulumi.StringOutput `pulumi:"dateCreated"`
	// Date the script was last modified.
	DateModified pulumi.StringOutput `pulumi:"dateModified"`
	// Name of the given script.
	Name pulumi.StringOutput `pulumi:"name"`
	// Contents of the startup script base64 encoded.
	Script pulumi.StringOutput `pulumi:"script"`
	// Type of startup script. Possible values are boot or pxe - default is boot.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewStartupScript registers a new resource with the given unique name, arguments, and options.
func NewStartupScript(ctx *pulumi.Context,
	name string, args *StartupScriptArgs, opts ...pulumi.ResourceOption) (*StartupScript, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Script == nil {
		return nil, errors.New("invalid value for required argument 'Script'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource StartupScript
	err := ctx.RegisterResource("vultr:index/startupScript:StartupScript", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStartupScript gets an existing StartupScript resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStartupScript(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StartupScriptState, opts ...pulumi.ResourceOption) (*StartupScript, error) {
	var resource StartupScript
	err := ctx.ReadResource("vultr:index/startupScript:StartupScript", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StartupScript resources.
type startupScriptState struct {
	// Date the script was created.
	DateCreated *string `pulumi:"dateCreated"`
	// Date the script was last modified.
	DateModified *string `pulumi:"dateModified"`
	// Name of the given script.
	Name *string `pulumi:"name"`
	// Contents of the startup script base64 encoded.
	Script *string `pulumi:"script"`
	// Type of startup script. Possible values are boot or pxe - default is boot.
	Type *string `pulumi:"type"`
}

type StartupScriptState struct {
	// Date the script was created.
	DateCreated pulumi.StringPtrInput
	// Date the script was last modified.
	DateModified pulumi.StringPtrInput
	// Name of the given script.
	Name pulumi.StringPtrInput
	// Contents of the startup script base64 encoded.
	Script pulumi.StringPtrInput
	// Type of startup script. Possible values are boot or pxe - default is boot.
	Type pulumi.StringPtrInput
}

func (StartupScriptState) ElementType() reflect.Type {
	return reflect.TypeOf((*startupScriptState)(nil)).Elem()
}

type startupScriptArgs struct {
	// Name of the given script.
	Name *string `pulumi:"name"`
	// Contents of the startup script base64 encoded.
	Script string `pulumi:"script"`
	// Type of startup script. Possible values are boot or pxe - default is boot.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a StartupScript resource.
type StartupScriptArgs struct {
	// Name of the given script.
	Name pulumi.StringPtrInput
	// Contents of the startup script base64 encoded.
	Script pulumi.StringInput
	// Type of startup script. Possible values are boot or pxe - default is boot.
	Type pulumi.StringPtrInput
}

func (StartupScriptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*startupScriptArgs)(nil)).Elem()
}

type StartupScriptInput interface {
	pulumi.Input

	ToStartupScriptOutput() StartupScriptOutput
	ToStartupScriptOutputWithContext(ctx context.Context) StartupScriptOutput
}

func (*StartupScript) ElementType() reflect.Type {
	return reflect.TypeOf((**StartupScript)(nil)).Elem()
}

func (i *StartupScript) ToStartupScriptOutput() StartupScriptOutput {
	return i.ToStartupScriptOutputWithContext(context.Background())
}

func (i *StartupScript) ToStartupScriptOutputWithContext(ctx context.Context) StartupScriptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StartupScriptOutput)
}

// StartupScriptArrayInput is an input type that accepts StartupScriptArray and StartupScriptArrayOutput values.
// You can construct a concrete instance of `StartupScriptArrayInput` via:
//
//	StartupScriptArray{ StartupScriptArgs{...} }
type StartupScriptArrayInput interface {
	pulumi.Input

	ToStartupScriptArrayOutput() StartupScriptArrayOutput
	ToStartupScriptArrayOutputWithContext(context.Context) StartupScriptArrayOutput
}

type StartupScriptArray []StartupScriptInput

func (StartupScriptArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StartupScript)(nil)).Elem()
}

func (i StartupScriptArray) ToStartupScriptArrayOutput() StartupScriptArrayOutput {
	return i.ToStartupScriptArrayOutputWithContext(context.Background())
}

func (i StartupScriptArray) ToStartupScriptArrayOutputWithContext(ctx context.Context) StartupScriptArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StartupScriptArrayOutput)
}

// StartupScriptMapInput is an input type that accepts StartupScriptMap and StartupScriptMapOutput values.
// You can construct a concrete instance of `StartupScriptMapInput` via:
//
//	StartupScriptMap{ "key": StartupScriptArgs{...} }
type StartupScriptMapInput interface {
	pulumi.Input

	ToStartupScriptMapOutput() StartupScriptMapOutput
	ToStartupScriptMapOutputWithContext(context.Context) StartupScriptMapOutput
}

type StartupScriptMap map[string]StartupScriptInput

func (StartupScriptMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StartupScript)(nil)).Elem()
}

func (i StartupScriptMap) ToStartupScriptMapOutput() StartupScriptMapOutput {
	return i.ToStartupScriptMapOutputWithContext(context.Background())
}

func (i StartupScriptMap) ToStartupScriptMapOutputWithContext(ctx context.Context) StartupScriptMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StartupScriptMapOutput)
}

type StartupScriptOutput struct{ *pulumi.OutputState }

func (StartupScriptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StartupScript)(nil)).Elem()
}

func (o StartupScriptOutput) ToStartupScriptOutput() StartupScriptOutput {
	return o
}

func (o StartupScriptOutput) ToStartupScriptOutputWithContext(ctx context.Context) StartupScriptOutput {
	return o
}

// Date the script was created.
func (o StartupScriptOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *StartupScript) pulumi.StringOutput { return v.DateCreated }).(pulumi.StringOutput)
}

// Date the script was last modified.
func (o StartupScriptOutput) DateModified() pulumi.StringOutput {
	return o.ApplyT(func(v *StartupScript) pulumi.StringOutput { return v.DateModified }).(pulumi.StringOutput)
}

// Name of the given script.
func (o StartupScriptOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StartupScript) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Contents of the startup script base64 encoded.
func (o StartupScriptOutput) Script() pulumi.StringOutput {
	return o.ApplyT(func(v *StartupScript) pulumi.StringOutput { return v.Script }).(pulumi.StringOutput)
}

// Type of startup script. Possible values are boot or pxe - default is boot.
func (o StartupScriptOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StartupScript) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type StartupScriptArrayOutput struct{ *pulumi.OutputState }

func (StartupScriptArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StartupScript)(nil)).Elem()
}

func (o StartupScriptArrayOutput) ToStartupScriptArrayOutput() StartupScriptArrayOutput {
	return o
}

func (o StartupScriptArrayOutput) ToStartupScriptArrayOutputWithContext(ctx context.Context) StartupScriptArrayOutput {
	return o
}

func (o StartupScriptArrayOutput) Index(i pulumi.IntInput) StartupScriptOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StartupScript {
		return vs[0].([]*StartupScript)[vs[1].(int)]
	}).(StartupScriptOutput)
}

type StartupScriptMapOutput struct{ *pulumi.OutputState }

func (StartupScriptMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StartupScript)(nil)).Elem()
}

func (o StartupScriptMapOutput) ToStartupScriptMapOutput() StartupScriptMapOutput {
	return o
}

func (o StartupScriptMapOutput) ToStartupScriptMapOutputWithContext(ctx context.Context) StartupScriptMapOutput {
	return o
}

func (o StartupScriptMapOutput) MapIndex(k pulumi.StringInput) StartupScriptOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StartupScript {
		return vs[0].(map[string]*StartupScript)[vs[1].(string)]
	}).(StartupScriptOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StartupScriptInput)(nil)).Elem(), &StartupScript{})
	pulumi.RegisterInputType(reflect.TypeOf((*StartupScriptArrayInput)(nil)).Elem(), StartupScriptArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StartupScriptMapInput)(nil)).Elem(), StartupScriptMap{})
	pulumi.RegisterOutputType(StartupScriptOutput{})
	pulumi.RegisterOutputType(StartupScriptArrayOutput{})
	pulumi.RegisterOutputType(StartupScriptMapOutput{})
}
