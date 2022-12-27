// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Vultr SSH key resource. This can be used to create, read, modify, and delete SSH keys.
//
// ## Example Usage
//
// # Create an SSH key
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
//			_, err := vultr.NewSSHKey(ctx, "mySshKey", &vultr.SSHKeyArgs{
//				SshKey: pulumi.String("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCyVGaw1PuEl98f4/7Kq3O9ZIvDw2OFOSXAFVqilSFNkHlefm1iMtPeqsIBp2t9cbGUf55xNDULz/bD/4BCV43yZ5lh0cUYuXALg9NI29ui7PEGReXjSpNwUD6ceN/78YOK41KAcecq+SS0bJ4b4amKZIJG3JWmDKljtv1dmSBCrTmEAQaOorxqGGBYmZS7NQumRe4lav5r6wOs8OACMANE1ejkeZsGFzJFNqvr5DuHdDL5FAudW23me3BDmrM9ifUzzjl1Jwku3bnRaCcjaxH8oTumt1a00mWci/1qUlaVFft085yvVq7KZbF2OPPbl+erDW91+EZ2FgEi+v1/CSJ5 your_username@hostname"),
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
// SSH keys can be imported using the SSH key `ID`, e.g.
//
// ```sh
//
//	$ pulumi import vultr:index/sSHKey:SSHKey my_key 6b0876a7-f709-41ba-aed8-abed9d38ae45
//
// ```
type SSHKey struct {
	pulumi.CustomResourceState

	// The date the SSH key was added to your Vultr account.
	DateCreated pulumi.StringOutput `pulumi:"dateCreated"`
	// The name/label of the SSH key.
	Name pulumi.StringOutput `pulumi:"name"`
	// The public SSH key.
	SshKey pulumi.StringOutput `pulumi:"sshKey"`
}

// NewSSHKey registers a new resource with the given unique name, arguments, and options.
func NewSSHKey(ctx *pulumi.Context,
	name string, args *SSHKeyArgs, opts ...pulumi.ResourceOption) (*SSHKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SshKey == nil {
		return nil, errors.New("invalid value for required argument 'SshKey'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SSHKey
	err := ctx.RegisterResource("vultr:index/sSHKey:SSHKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSSHKey gets an existing SSHKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSSHKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SSHKeyState, opts ...pulumi.ResourceOption) (*SSHKey, error) {
	var resource SSHKey
	err := ctx.ReadResource("vultr:index/sSHKey:SSHKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SSHKey resources.
type sshkeyState struct {
	// The date the SSH key was added to your Vultr account.
	DateCreated *string `pulumi:"dateCreated"`
	// The name/label of the SSH key.
	Name *string `pulumi:"name"`
	// The public SSH key.
	SshKey *string `pulumi:"sshKey"`
}

type SSHKeyState struct {
	// The date the SSH key was added to your Vultr account.
	DateCreated pulumi.StringPtrInput
	// The name/label of the SSH key.
	Name pulumi.StringPtrInput
	// The public SSH key.
	SshKey pulumi.StringPtrInput
}

func (SSHKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sshkeyState)(nil)).Elem()
}

type sshkeyArgs struct {
	// The name/label of the SSH key.
	Name *string `pulumi:"name"`
	// The public SSH key.
	SshKey string `pulumi:"sshKey"`
}

// The set of arguments for constructing a SSHKey resource.
type SSHKeyArgs struct {
	// The name/label of the SSH key.
	Name pulumi.StringPtrInput
	// The public SSH key.
	SshKey pulumi.StringInput
}

func (SSHKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sshkeyArgs)(nil)).Elem()
}

type SSHKeyInput interface {
	pulumi.Input

	ToSSHKeyOutput() SSHKeyOutput
	ToSSHKeyOutputWithContext(ctx context.Context) SSHKeyOutput
}

func (*SSHKey) ElementType() reflect.Type {
	return reflect.TypeOf((**SSHKey)(nil)).Elem()
}

func (i *SSHKey) ToSSHKeyOutput() SSHKeyOutput {
	return i.ToSSHKeyOutputWithContext(context.Background())
}

func (i *SSHKey) ToSSHKeyOutputWithContext(ctx context.Context) SSHKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SSHKeyOutput)
}

// SSHKeyArrayInput is an input type that accepts SSHKeyArray and SSHKeyArrayOutput values.
// You can construct a concrete instance of `SSHKeyArrayInput` via:
//
//	SSHKeyArray{ SSHKeyArgs{...} }
type SSHKeyArrayInput interface {
	pulumi.Input

	ToSSHKeyArrayOutput() SSHKeyArrayOutput
	ToSSHKeyArrayOutputWithContext(context.Context) SSHKeyArrayOutput
}

type SSHKeyArray []SSHKeyInput

func (SSHKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SSHKey)(nil)).Elem()
}

func (i SSHKeyArray) ToSSHKeyArrayOutput() SSHKeyArrayOutput {
	return i.ToSSHKeyArrayOutputWithContext(context.Background())
}

func (i SSHKeyArray) ToSSHKeyArrayOutputWithContext(ctx context.Context) SSHKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SSHKeyArrayOutput)
}

// SSHKeyMapInput is an input type that accepts SSHKeyMap and SSHKeyMapOutput values.
// You can construct a concrete instance of `SSHKeyMapInput` via:
//
//	SSHKeyMap{ "key": SSHKeyArgs{...} }
type SSHKeyMapInput interface {
	pulumi.Input

	ToSSHKeyMapOutput() SSHKeyMapOutput
	ToSSHKeyMapOutputWithContext(context.Context) SSHKeyMapOutput
}

type SSHKeyMap map[string]SSHKeyInput

func (SSHKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SSHKey)(nil)).Elem()
}

func (i SSHKeyMap) ToSSHKeyMapOutput() SSHKeyMapOutput {
	return i.ToSSHKeyMapOutputWithContext(context.Background())
}

func (i SSHKeyMap) ToSSHKeyMapOutputWithContext(ctx context.Context) SSHKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SSHKeyMapOutput)
}

type SSHKeyOutput struct{ *pulumi.OutputState }

func (SSHKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SSHKey)(nil)).Elem()
}

func (o SSHKeyOutput) ToSSHKeyOutput() SSHKeyOutput {
	return o
}

func (o SSHKeyOutput) ToSSHKeyOutputWithContext(ctx context.Context) SSHKeyOutput {
	return o
}

// The date the SSH key was added to your Vultr account.
func (o SSHKeyOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *SSHKey) pulumi.StringOutput { return v.DateCreated }).(pulumi.StringOutput)
}

// The name/label of the SSH key.
func (o SSHKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SSHKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The public SSH key.
func (o SSHKeyOutput) SshKey() pulumi.StringOutput {
	return o.ApplyT(func(v *SSHKey) pulumi.StringOutput { return v.SshKey }).(pulumi.StringOutput)
}

type SSHKeyArrayOutput struct{ *pulumi.OutputState }

func (SSHKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SSHKey)(nil)).Elem()
}

func (o SSHKeyArrayOutput) ToSSHKeyArrayOutput() SSHKeyArrayOutput {
	return o
}

func (o SSHKeyArrayOutput) ToSSHKeyArrayOutputWithContext(ctx context.Context) SSHKeyArrayOutput {
	return o
}

func (o SSHKeyArrayOutput) Index(i pulumi.IntInput) SSHKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SSHKey {
		return vs[0].([]*SSHKey)[vs[1].(int)]
	}).(SSHKeyOutput)
}

type SSHKeyMapOutput struct{ *pulumi.OutputState }

func (SSHKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SSHKey)(nil)).Elem()
}

func (o SSHKeyMapOutput) ToSSHKeyMapOutput() SSHKeyMapOutput {
	return o
}

func (o SSHKeyMapOutput) ToSSHKeyMapOutputWithContext(ctx context.Context) SSHKeyMapOutput {
	return o
}

func (o SSHKeyMapOutput) MapIndex(k pulumi.StringInput) SSHKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SSHKey {
		return vs[0].(map[string]*SSHKey)[vs[1].(string)]
	}).(SSHKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SSHKeyInput)(nil)).Elem(), &SSHKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*SSHKeyArrayInput)(nil)).Elem(), SSHKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SSHKeyMapInput)(nil)).Elem(), SSHKeyMap{})
	pulumi.RegisterOutputType(SSHKeyOutput{})
	pulumi.RegisterOutputType(SSHKeyArrayOutput{})
	pulumi.RegisterOutputType(SSHKeyMapOutput{})
}
