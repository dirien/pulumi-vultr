// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Vultr Firewall Rule resource. This can be used to create, read, modify, and delete Firewall rules.
//
// ## Example Usage
//
// # Create a Firewall Rule
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vultr/sdk/v2/go/vultr"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			myFirewallgroup, err := vultr.NewFirewallGroup(ctx, "myFirewallgroup", &vultr.FirewallGroupArgs{
//				Description: pulumi.String("base firewall"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vultr.NewFirewallRule(ctx, "myFirewallrule", &vultr.FirewallRuleArgs{
//				FirewallGroupId: myFirewallgroup.ID(),
//				Protocol:        pulumi.String("tcp"),
//				IpType:          pulumi.String("v4"),
//				Subnet:          pulumi.String("0.0.0.0"),
//				SubnetSize:      pulumi.Int(0),
//				Port:            pulumi.String("8090"),
//				Notes:           pulumi.String("my firewall rule"),
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
// Firewall Rules can be imported using the Firewall Group `ID` and Firewall Rule `ID`, e.g.
//
// ```sh
//
//	$ pulumi import vultr:index/firewallRule:FirewallRule my_rule b6a859c5-b299-49dd-8888-b1abbc517d08,1
//
// ```
type FirewallRule struct {
	pulumi.CustomResourceState

	// The firewall group that the firewall rule will belong to.
	FirewallGroupId pulumi.StringOutput `pulumi:"firewallGroupId"`
	// The type of ip for this firewall rule. Possible values (v4, v6) **Note** they must be lowercase
	IpType pulumi.StringOutput `pulumi:"ipType"`
	// A simple note for a given firewall rule
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// TCP/UDP only. This field can be a specific port or a colon separated port range.
	Port pulumi.StringPtrOutput `pulumi:"port"`
	// The type of protocol for this firewall rule. Possible values (icmp, tcp, udp, gre, esp, ah) **Note** they must be lowercase
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Possible values ("", cloudflare)
	Source pulumi.StringPtrOutput `pulumi:"source"`
	// IP address that you want to define for this firewall rule.
	Subnet pulumi.StringOutput `pulumi:"subnet"`
	// The number of bits for the subnet in CIDR notation. Example: 32.
	SubnetSize pulumi.IntOutput `pulumi:"subnetSize"`
}

// NewFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewFirewallRule(ctx *pulumi.Context,
	name string, args *FirewallRuleArgs, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FirewallGroupId == nil {
		return nil, errors.New("invalid value for required argument 'FirewallGroupId'")
	}
	if args.IpType == nil {
		return nil, errors.New("invalid value for required argument 'IpType'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.Subnet == nil {
		return nil, errors.New("invalid value for required argument 'Subnet'")
	}
	if args.SubnetSize == nil {
		return nil, errors.New("invalid value for required argument 'SubnetSize'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallRule
	err := ctx.RegisterResource("vultr:index/firewallRule:FirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallRule gets an existing FirewallRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallRuleState, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	var resource FirewallRule
	err := ctx.ReadResource("vultr:index/firewallRule:FirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallRule resources.
type firewallRuleState struct {
	// The firewall group that the firewall rule will belong to.
	FirewallGroupId *string `pulumi:"firewallGroupId"`
	// The type of ip for this firewall rule. Possible values (v4, v6) **Note** they must be lowercase
	IpType *string `pulumi:"ipType"`
	// A simple note for a given firewall rule
	Notes *string `pulumi:"notes"`
	// TCP/UDP only. This field can be a specific port or a colon separated port range.
	Port *string `pulumi:"port"`
	// The type of protocol for this firewall rule. Possible values (icmp, tcp, udp, gre, esp, ah) **Note** they must be lowercase
	Protocol *string `pulumi:"protocol"`
	// Possible values ("", cloudflare)
	Source *string `pulumi:"source"`
	// IP address that you want to define for this firewall rule.
	Subnet *string `pulumi:"subnet"`
	// The number of bits for the subnet in CIDR notation. Example: 32.
	SubnetSize *int `pulumi:"subnetSize"`
}

type FirewallRuleState struct {
	// The firewall group that the firewall rule will belong to.
	FirewallGroupId pulumi.StringPtrInput
	// The type of ip for this firewall rule. Possible values (v4, v6) **Note** they must be lowercase
	IpType pulumi.StringPtrInput
	// A simple note for a given firewall rule
	Notes pulumi.StringPtrInput
	// TCP/UDP only. This field can be a specific port or a colon separated port range.
	Port pulumi.StringPtrInput
	// The type of protocol for this firewall rule. Possible values (icmp, tcp, udp, gre, esp, ah) **Note** they must be lowercase
	Protocol pulumi.StringPtrInput
	// Possible values ("", cloudflare)
	Source pulumi.StringPtrInput
	// IP address that you want to define for this firewall rule.
	Subnet pulumi.StringPtrInput
	// The number of bits for the subnet in CIDR notation. Example: 32.
	SubnetSize pulumi.IntPtrInput
}

func (FirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleState)(nil)).Elem()
}

type firewallRuleArgs struct {
	// The firewall group that the firewall rule will belong to.
	FirewallGroupId string `pulumi:"firewallGroupId"`
	// The type of ip for this firewall rule. Possible values (v4, v6) **Note** they must be lowercase
	IpType string `pulumi:"ipType"`
	// A simple note for a given firewall rule
	Notes *string `pulumi:"notes"`
	// TCP/UDP only. This field can be a specific port or a colon separated port range.
	Port *string `pulumi:"port"`
	// The type of protocol for this firewall rule. Possible values (icmp, tcp, udp, gre, esp, ah) **Note** they must be lowercase
	Protocol string `pulumi:"protocol"`
	// Possible values ("", cloudflare)
	Source *string `pulumi:"source"`
	// IP address that you want to define for this firewall rule.
	Subnet string `pulumi:"subnet"`
	// The number of bits for the subnet in CIDR notation. Example: 32.
	SubnetSize int `pulumi:"subnetSize"`
}

// The set of arguments for constructing a FirewallRule resource.
type FirewallRuleArgs struct {
	// The firewall group that the firewall rule will belong to.
	FirewallGroupId pulumi.StringInput
	// The type of ip for this firewall rule. Possible values (v4, v6) **Note** they must be lowercase
	IpType pulumi.StringInput
	// A simple note for a given firewall rule
	Notes pulumi.StringPtrInput
	// TCP/UDP only. This field can be a specific port or a colon separated port range.
	Port pulumi.StringPtrInput
	// The type of protocol for this firewall rule. Possible values (icmp, tcp, udp, gre, esp, ah) **Note** they must be lowercase
	Protocol pulumi.StringInput
	// Possible values ("", cloudflare)
	Source pulumi.StringPtrInput
	// IP address that you want to define for this firewall rule.
	Subnet pulumi.StringInput
	// The number of bits for the subnet in CIDR notation. Example: 32.
	SubnetSize pulumi.IntInput
}

func (FirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleArgs)(nil)).Elem()
}

type FirewallRuleInput interface {
	pulumi.Input

	ToFirewallRuleOutput() FirewallRuleOutput
	ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput
}

func (*FirewallRule) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallRule)(nil)).Elem()
}

func (i *FirewallRule) ToFirewallRuleOutput() FirewallRuleOutput {
	return i.ToFirewallRuleOutputWithContext(context.Background())
}

func (i *FirewallRule) ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleOutput)
}

// FirewallRuleArrayInput is an input type that accepts FirewallRuleArray and FirewallRuleArrayOutput values.
// You can construct a concrete instance of `FirewallRuleArrayInput` via:
//
//	FirewallRuleArray{ FirewallRuleArgs{...} }
type FirewallRuleArrayInput interface {
	pulumi.Input

	ToFirewallRuleArrayOutput() FirewallRuleArrayOutput
	ToFirewallRuleArrayOutputWithContext(context.Context) FirewallRuleArrayOutput
}

type FirewallRuleArray []FirewallRuleInput

func (FirewallRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallRule)(nil)).Elem()
}

func (i FirewallRuleArray) ToFirewallRuleArrayOutput() FirewallRuleArrayOutput {
	return i.ToFirewallRuleArrayOutputWithContext(context.Background())
}

func (i FirewallRuleArray) ToFirewallRuleArrayOutputWithContext(ctx context.Context) FirewallRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleArrayOutput)
}

// FirewallRuleMapInput is an input type that accepts FirewallRuleMap and FirewallRuleMapOutput values.
// You can construct a concrete instance of `FirewallRuleMapInput` via:
//
//	FirewallRuleMap{ "key": FirewallRuleArgs{...} }
type FirewallRuleMapInput interface {
	pulumi.Input

	ToFirewallRuleMapOutput() FirewallRuleMapOutput
	ToFirewallRuleMapOutputWithContext(context.Context) FirewallRuleMapOutput
}

type FirewallRuleMap map[string]FirewallRuleInput

func (FirewallRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallRule)(nil)).Elem()
}

func (i FirewallRuleMap) ToFirewallRuleMapOutput() FirewallRuleMapOutput {
	return i.ToFirewallRuleMapOutputWithContext(context.Background())
}

func (i FirewallRuleMap) ToFirewallRuleMapOutputWithContext(ctx context.Context) FirewallRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleMapOutput)
}

type FirewallRuleOutput struct{ *pulumi.OutputState }

func (FirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallRule)(nil)).Elem()
}

func (o FirewallRuleOutput) ToFirewallRuleOutput() FirewallRuleOutput {
	return o
}

func (o FirewallRuleOutput) ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput {
	return o
}

// The firewall group that the firewall rule will belong to.
func (o FirewallRuleOutput) FirewallGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.FirewallGroupId }).(pulumi.StringOutput)
}

// The type of ip for this firewall rule. Possible values (v4, v6) **Note** they must be lowercase
func (o FirewallRuleOutput) IpType() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.IpType }).(pulumi.StringOutput)
}

// A simple note for a given firewall rule
func (o FirewallRuleOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

// TCP/UDP only. This field can be a specific port or a colon separated port range.
func (o FirewallRuleOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringPtrOutput { return v.Port }).(pulumi.StringPtrOutput)
}

// The type of protocol for this firewall rule. Possible values (icmp, tcp, udp, gre, esp, ah) **Note** they must be lowercase
func (o FirewallRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// Possible values ("", cloudflare)
func (o FirewallRuleOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringPtrOutput { return v.Source }).(pulumi.StringPtrOutput)
}

// IP address that you want to define for this firewall rule.
func (o FirewallRuleOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.Subnet }).(pulumi.StringOutput)
}

// The number of bits for the subnet in CIDR notation. Example: 32.
func (o FirewallRuleOutput) SubnetSize() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.IntOutput { return v.SubnetSize }).(pulumi.IntOutput)
}

type FirewallRuleArrayOutput struct{ *pulumi.OutputState }

func (FirewallRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallRule)(nil)).Elem()
}

func (o FirewallRuleArrayOutput) ToFirewallRuleArrayOutput() FirewallRuleArrayOutput {
	return o
}

func (o FirewallRuleArrayOutput) ToFirewallRuleArrayOutputWithContext(ctx context.Context) FirewallRuleArrayOutput {
	return o
}

func (o FirewallRuleArrayOutput) Index(i pulumi.IntInput) FirewallRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallRule {
		return vs[0].([]*FirewallRule)[vs[1].(int)]
	}).(FirewallRuleOutput)
}

type FirewallRuleMapOutput struct{ *pulumi.OutputState }

func (FirewallRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallRule)(nil)).Elem()
}

func (o FirewallRuleMapOutput) ToFirewallRuleMapOutput() FirewallRuleMapOutput {
	return o
}

func (o FirewallRuleMapOutput) ToFirewallRuleMapOutputWithContext(ctx context.Context) FirewallRuleMapOutput {
	return o
}

func (o FirewallRuleMapOutput) MapIndex(k pulumi.StringInput) FirewallRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallRule {
		return vs[0].(map[string]*FirewallRule)[vs[1].(string)]
	}).(FirewallRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallRuleInput)(nil)).Elem(), &FirewallRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallRuleArrayInput)(nil)).Elem(), FirewallRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallRuleMapInput)(nil)).Elem(), FirewallRuleMap{})
	pulumi.RegisterOutputType(FirewallRuleOutput{})
	pulumi.RegisterOutputType(FirewallRuleArrayOutput{})
	pulumi.RegisterOutputType(FirewallRuleMapOutput{})
}
