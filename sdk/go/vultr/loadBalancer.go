// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Vultr load balancer.
//
// ## Example Usage
//
// Create a new load balancer:
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
//			_, err := vultr.NewLoadBalancer(ctx, "lb", &vultr.LoadBalancerArgs{
//				BalancingAlgorithm: pulumi.String("roundrobin"),
//				ForwardingRules: vultr.LoadBalancerForwardingRuleArray{
//					&vultr.LoadBalancerForwardingRuleArgs{
//						BackendPort:      pulumi.Int(81),
//						BackendProtocol:  pulumi.String("http"),
//						FrontendPort:     pulumi.Int(82),
//						FrontendProtocol: pulumi.String("http"),
//					},
//				},
//				HealthCheck: &vultr.LoadBalancerHealthCheckArgs{
//					CheckInterval:      pulumi.Int(3),
//					HealthyThreshold:   pulumi.Int(4),
//					Path:               pulumi.String("/test"),
//					Port:               pulumi.Int(8080),
//					Protocol:           pulumi.String("http"),
//					ResponseTimeout:    pulumi.Int(1),
//					UnhealthyThreshold: pulumi.Int(2),
//				},
//				Label:  pulumi.String("terraform lb example"),
//				Region: pulumi.String("ewr"),
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
// Load Balancers can be imported using the load balancer `ID`, e.g.
//
// ```sh
//
//	$ pulumi import vultr:index/loadBalancer:LoadBalancer lb b6a859c5-b299-49dd-8888-b1abbc517d08
//
// ```
type LoadBalancer struct {
	pulumi.CustomResourceState

	// Array of instances that are currently attached to the load balancer.
	AttachedInstances pulumi.StringArrayOutput `pulumi:"attachedInstances"`
	// The balancing algorithm for your load balancer. Options are `roundrobin` or `leastconn`. Default value is `roundrobin`
	BalancingAlgorithm pulumi.StringOutput `pulumi:"balancingAlgorithm"`
	// Name for your given sticky session.
	CookieName pulumi.StringPtrOutput `pulumi:"cookieName"`
	// Defines the firewall rules for a load balancer.
	FirewallRules LoadBalancerFirewallRuleArrayOutput `pulumi:"firewallRules"`
	// List of forwarding rules for a load balancer. The configuration of a `forwardingRules` is listened below.
	ForwardingRules LoadBalancerForwardingRuleArrayOutput `pulumi:"forwardingRules"`
	// Boolean value that indicates if SSL is enabled.
	HasSsl pulumi.BoolOutput `pulumi:"hasSsl"`
	// A block that defines the way load balancers should check for health. The configuration of a `healthCheck` is listed below.
	HealthCheck LoadBalancerHealthCheckOutput `pulumi:"healthCheck"`
	// IPv4 address for your load balancer.
	Ipv4 pulumi.StringOutput `pulumi:"ipv4"`
	// IPv6 address for your load balancer.
	Ipv6 pulumi.StringOutput `pulumi:"ipv6"`
	// The load balancer's label.
	Label pulumi.StringPtrOutput `pulumi:"label"`
	// A private network ID that the load balancer should be attached to.
	//
	// Deprecated: private_network is deprecated and should no longer be used. Instead, use vpc
	PrivateNetwork pulumi.StringPtrOutput `pulumi:"privateNetwork"`
	// Boolean value that indicates if Proxy Protocol is enabled.
	ProxyProtocol pulumi.BoolPtrOutput `pulumi:"proxyProtocol"`
	// The region your load balancer is deployed in.
	Region pulumi.StringOutput `pulumi:"region"`
	// A block that supplies your ssl configuration to be used with HTTPS. The configuration of a `ssl` is listed below.
	Ssl LoadBalancerSslPtrOutput `pulumi:"ssl"`
	// Boolean value that indicates if HTTP calls will be redirected to HTTPS.
	SslRedirect pulumi.BoolPtrOutput `pulumi:"sslRedirect"`
	// Current status for the load balancer
	Status pulumi.StringOutput `pulumi:"status"`
	// A VPC ID that the load balancer should be attached to.
	Vpc pulumi.StringPtrOutput `pulumi:"vpc"`
}

// NewLoadBalancer registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancer(ctx *pulumi.Context,
	name string, args *LoadBalancerArgs, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ForwardingRules == nil {
		return nil, errors.New("invalid value for required argument 'ForwardingRules'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource LoadBalancer
	err := ctx.RegisterResource("vultr:index/loadBalancer:LoadBalancer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadBalancer gets an existing LoadBalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancerState, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	var resource LoadBalancer
	err := ctx.ReadResource("vultr:index/loadBalancer:LoadBalancer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadBalancer resources.
type loadBalancerState struct {
	// Array of instances that are currently attached to the load balancer.
	AttachedInstances []string `pulumi:"attachedInstances"`
	// The balancing algorithm for your load balancer. Options are `roundrobin` or `leastconn`. Default value is `roundrobin`
	BalancingAlgorithm *string `pulumi:"balancingAlgorithm"`
	// Name for your given sticky session.
	CookieName *string `pulumi:"cookieName"`
	// Defines the firewall rules for a load balancer.
	FirewallRules []LoadBalancerFirewallRule `pulumi:"firewallRules"`
	// List of forwarding rules for a load balancer. The configuration of a `forwardingRules` is listened below.
	ForwardingRules []LoadBalancerForwardingRule `pulumi:"forwardingRules"`
	// Boolean value that indicates if SSL is enabled.
	HasSsl *bool `pulumi:"hasSsl"`
	// A block that defines the way load balancers should check for health. The configuration of a `healthCheck` is listed below.
	HealthCheck *LoadBalancerHealthCheck `pulumi:"healthCheck"`
	// IPv4 address for your load balancer.
	Ipv4 *string `pulumi:"ipv4"`
	// IPv6 address for your load balancer.
	Ipv6 *string `pulumi:"ipv6"`
	// The load balancer's label.
	Label *string `pulumi:"label"`
	// A private network ID that the load balancer should be attached to.
	//
	// Deprecated: private_network is deprecated and should no longer be used. Instead, use vpc
	PrivateNetwork *string `pulumi:"privateNetwork"`
	// Boolean value that indicates if Proxy Protocol is enabled.
	ProxyProtocol *bool `pulumi:"proxyProtocol"`
	// The region your load balancer is deployed in.
	Region *string `pulumi:"region"`
	// A block that supplies your ssl configuration to be used with HTTPS. The configuration of a `ssl` is listed below.
	Ssl *LoadBalancerSsl `pulumi:"ssl"`
	// Boolean value that indicates if HTTP calls will be redirected to HTTPS.
	SslRedirect *bool `pulumi:"sslRedirect"`
	// Current status for the load balancer
	Status *string `pulumi:"status"`
	// A VPC ID that the load balancer should be attached to.
	Vpc *string `pulumi:"vpc"`
}

type LoadBalancerState struct {
	// Array of instances that are currently attached to the load balancer.
	AttachedInstances pulumi.StringArrayInput
	// The balancing algorithm for your load balancer. Options are `roundrobin` or `leastconn`. Default value is `roundrobin`
	BalancingAlgorithm pulumi.StringPtrInput
	// Name for your given sticky session.
	CookieName pulumi.StringPtrInput
	// Defines the firewall rules for a load balancer.
	FirewallRules LoadBalancerFirewallRuleArrayInput
	// List of forwarding rules for a load balancer. The configuration of a `forwardingRules` is listened below.
	ForwardingRules LoadBalancerForwardingRuleArrayInput
	// Boolean value that indicates if SSL is enabled.
	HasSsl pulumi.BoolPtrInput
	// A block that defines the way load balancers should check for health. The configuration of a `healthCheck` is listed below.
	HealthCheck LoadBalancerHealthCheckPtrInput
	// IPv4 address for your load balancer.
	Ipv4 pulumi.StringPtrInput
	// IPv6 address for your load balancer.
	Ipv6 pulumi.StringPtrInput
	// The load balancer's label.
	Label pulumi.StringPtrInput
	// A private network ID that the load balancer should be attached to.
	//
	// Deprecated: private_network is deprecated and should no longer be used. Instead, use vpc
	PrivateNetwork pulumi.StringPtrInput
	// Boolean value that indicates if Proxy Protocol is enabled.
	ProxyProtocol pulumi.BoolPtrInput
	// The region your load balancer is deployed in.
	Region pulumi.StringPtrInput
	// A block that supplies your ssl configuration to be used with HTTPS. The configuration of a `ssl` is listed below.
	Ssl LoadBalancerSslPtrInput
	// Boolean value that indicates if HTTP calls will be redirected to HTTPS.
	SslRedirect pulumi.BoolPtrInput
	// Current status for the load balancer
	Status pulumi.StringPtrInput
	// A VPC ID that the load balancer should be attached to.
	Vpc pulumi.StringPtrInput
}

func (LoadBalancerState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerState)(nil)).Elem()
}

type loadBalancerArgs struct {
	// Array of instances that are currently attached to the load balancer.
	AttachedInstances []string `pulumi:"attachedInstances"`
	// The balancing algorithm for your load balancer. Options are `roundrobin` or `leastconn`. Default value is `roundrobin`
	BalancingAlgorithm *string `pulumi:"balancingAlgorithm"`
	// Name for your given sticky session.
	CookieName *string `pulumi:"cookieName"`
	// Defines the firewall rules for a load balancer.
	FirewallRules []LoadBalancerFirewallRule `pulumi:"firewallRules"`
	// List of forwarding rules for a load balancer. The configuration of a `forwardingRules` is listened below.
	ForwardingRules []LoadBalancerForwardingRule `pulumi:"forwardingRules"`
	// A block that defines the way load balancers should check for health. The configuration of a `healthCheck` is listed below.
	HealthCheck *LoadBalancerHealthCheck `pulumi:"healthCheck"`
	// The load balancer's label.
	Label *string `pulumi:"label"`
	// A private network ID that the load balancer should be attached to.
	//
	// Deprecated: private_network is deprecated and should no longer be used. Instead, use vpc
	PrivateNetwork *string `pulumi:"privateNetwork"`
	// Boolean value that indicates if Proxy Protocol is enabled.
	ProxyProtocol *bool `pulumi:"proxyProtocol"`
	// The region your load balancer is deployed in.
	Region string `pulumi:"region"`
	// A block that supplies your ssl configuration to be used with HTTPS. The configuration of a `ssl` is listed below.
	Ssl *LoadBalancerSsl `pulumi:"ssl"`
	// Boolean value that indicates if HTTP calls will be redirected to HTTPS.
	SslRedirect *bool `pulumi:"sslRedirect"`
	// A VPC ID that the load balancer should be attached to.
	Vpc *string `pulumi:"vpc"`
}

// The set of arguments for constructing a LoadBalancer resource.
type LoadBalancerArgs struct {
	// Array of instances that are currently attached to the load balancer.
	AttachedInstances pulumi.StringArrayInput
	// The balancing algorithm for your load balancer. Options are `roundrobin` or `leastconn`. Default value is `roundrobin`
	BalancingAlgorithm pulumi.StringPtrInput
	// Name for your given sticky session.
	CookieName pulumi.StringPtrInput
	// Defines the firewall rules for a load balancer.
	FirewallRules LoadBalancerFirewallRuleArrayInput
	// List of forwarding rules for a load balancer. The configuration of a `forwardingRules` is listened below.
	ForwardingRules LoadBalancerForwardingRuleArrayInput
	// A block that defines the way load balancers should check for health. The configuration of a `healthCheck` is listed below.
	HealthCheck LoadBalancerHealthCheckPtrInput
	// The load balancer's label.
	Label pulumi.StringPtrInput
	// A private network ID that the load balancer should be attached to.
	//
	// Deprecated: private_network is deprecated and should no longer be used. Instead, use vpc
	PrivateNetwork pulumi.StringPtrInput
	// Boolean value that indicates if Proxy Protocol is enabled.
	ProxyProtocol pulumi.BoolPtrInput
	// The region your load balancer is deployed in.
	Region pulumi.StringInput
	// A block that supplies your ssl configuration to be used with HTTPS. The configuration of a `ssl` is listed below.
	Ssl LoadBalancerSslPtrInput
	// Boolean value that indicates if HTTP calls will be redirected to HTTPS.
	SslRedirect pulumi.BoolPtrInput
	// A VPC ID that the load balancer should be attached to.
	Vpc pulumi.StringPtrInput
}

func (LoadBalancerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerArgs)(nil)).Elem()
}

type LoadBalancerInput interface {
	pulumi.Input

	ToLoadBalancerOutput() LoadBalancerOutput
	ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput
}

func (*LoadBalancer) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancer)(nil)).Elem()
}

func (i *LoadBalancer) ToLoadBalancerOutput() LoadBalancerOutput {
	return i.ToLoadBalancerOutputWithContext(context.Background())
}

func (i *LoadBalancer) ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerOutput)
}

// LoadBalancerArrayInput is an input type that accepts LoadBalancerArray and LoadBalancerArrayOutput values.
// You can construct a concrete instance of `LoadBalancerArrayInput` via:
//
//	LoadBalancerArray{ LoadBalancerArgs{...} }
type LoadBalancerArrayInput interface {
	pulumi.Input

	ToLoadBalancerArrayOutput() LoadBalancerArrayOutput
	ToLoadBalancerArrayOutputWithContext(context.Context) LoadBalancerArrayOutput
}

type LoadBalancerArray []LoadBalancerInput

func (LoadBalancerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoadBalancer)(nil)).Elem()
}

func (i LoadBalancerArray) ToLoadBalancerArrayOutput() LoadBalancerArrayOutput {
	return i.ToLoadBalancerArrayOutputWithContext(context.Background())
}

func (i LoadBalancerArray) ToLoadBalancerArrayOutputWithContext(ctx context.Context) LoadBalancerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerArrayOutput)
}

// LoadBalancerMapInput is an input type that accepts LoadBalancerMap and LoadBalancerMapOutput values.
// You can construct a concrete instance of `LoadBalancerMapInput` via:
//
//	LoadBalancerMap{ "key": LoadBalancerArgs{...} }
type LoadBalancerMapInput interface {
	pulumi.Input

	ToLoadBalancerMapOutput() LoadBalancerMapOutput
	ToLoadBalancerMapOutputWithContext(context.Context) LoadBalancerMapOutput
}

type LoadBalancerMap map[string]LoadBalancerInput

func (LoadBalancerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoadBalancer)(nil)).Elem()
}

func (i LoadBalancerMap) ToLoadBalancerMapOutput() LoadBalancerMapOutput {
	return i.ToLoadBalancerMapOutputWithContext(context.Background())
}

func (i LoadBalancerMap) ToLoadBalancerMapOutputWithContext(ctx context.Context) LoadBalancerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerMapOutput)
}

type LoadBalancerOutput struct{ *pulumi.OutputState }

func (LoadBalancerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancer)(nil)).Elem()
}

func (o LoadBalancerOutput) ToLoadBalancerOutput() LoadBalancerOutput {
	return o
}

func (o LoadBalancerOutput) ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput {
	return o
}

// Array of instances that are currently attached to the load balancer.
func (o LoadBalancerOutput) AttachedInstances() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringArrayOutput { return v.AttachedInstances }).(pulumi.StringArrayOutput)
}

// The balancing algorithm for your load balancer. Options are `roundrobin` or `leastconn`. Default value is `roundrobin`
func (o LoadBalancerOutput) BalancingAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.BalancingAlgorithm }).(pulumi.StringOutput)
}

// Name for your given sticky session.
func (o LoadBalancerOutput) CookieName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringPtrOutput { return v.CookieName }).(pulumi.StringPtrOutput)
}

// Defines the firewall rules for a load balancer.
func (o LoadBalancerOutput) FirewallRules() LoadBalancerFirewallRuleArrayOutput {
	return o.ApplyT(func(v *LoadBalancer) LoadBalancerFirewallRuleArrayOutput { return v.FirewallRules }).(LoadBalancerFirewallRuleArrayOutput)
}

// List of forwarding rules for a load balancer. The configuration of a `forwardingRules` is listened below.
func (o LoadBalancerOutput) ForwardingRules() LoadBalancerForwardingRuleArrayOutput {
	return o.ApplyT(func(v *LoadBalancer) LoadBalancerForwardingRuleArrayOutput { return v.ForwardingRules }).(LoadBalancerForwardingRuleArrayOutput)
}

// Boolean value that indicates if SSL is enabled.
func (o LoadBalancerOutput) HasSsl() pulumi.BoolOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.BoolOutput { return v.HasSsl }).(pulumi.BoolOutput)
}

// A block that defines the way load balancers should check for health. The configuration of a `healthCheck` is listed below.
func (o LoadBalancerOutput) HealthCheck() LoadBalancerHealthCheckOutput {
	return o.ApplyT(func(v *LoadBalancer) LoadBalancerHealthCheckOutput { return v.HealthCheck }).(LoadBalancerHealthCheckOutput)
}

// IPv4 address for your load balancer.
func (o LoadBalancerOutput) Ipv4() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.Ipv4 }).(pulumi.StringOutput)
}

// IPv6 address for your load balancer.
func (o LoadBalancerOutput) Ipv6() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.Ipv6 }).(pulumi.StringOutput)
}

// The load balancer's label.
func (o LoadBalancerOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringPtrOutput { return v.Label }).(pulumi.StringPtrOutput)
}

// A private network ID that the load balancer should be attached to.
//
// Deprecated: private_network is deprecated and should no longer be used. Instead, use vpc
func (o LoadBalancerOutput) PrivateNetwork() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringPtrOutput { return v.PrivateNetwork }).(pulumi.StringPtrOutput)
}

// Boolean value that indicates if Proxy Protocol is enabled.
func (o LoadBalancerOutput) ProxyProtocol() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.BoolPtrOutput { return v.ProxyProtocol }).(pulumi.BoolPtrOutput)
}

// The region your load balancer is deployed in.
func (o LoadBalancerOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// A block that supplies your ssl configuration to be used with HTTPS. The configuration of a `ssl` is listed below.
func (o LoadBalancerOutput) Ssl() LoadBalancerSslPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) LoadBalancerSslPtrOutput { return v.Ssl }).(LoadBalancerSslPtrOutput)
}

// Boolean value that indicates if HTTP calls will be redirected to HTTPS.
func (o LoadBalancerOutput) SslRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.BoolPtrOutput { return v.SslRedirect }).(pulumi.BoolPtrOutput)
}

// Current status for the load balancer
func (o LoadBalancerOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A VPC ID that the load balancer should be attached to.
func (o LoadBalancerOutput) Vpc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringPtrOutput { return v.Vpc }).(pulumi.StringPtrOutput)
}

type LoadBalancerArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoadBalancer)(nil)).Elem()
}

func (o LoadBalancerArrayOutput) ToLoadBalancerArrayOutput() LoadBalancerArrayOutput {
	return o
}

func (o LoadBalancerArrayOutput) ToLoadBalancerArrayOutputWithContext(ctx context.Context) LoadBalancerArrayOutput {
	return o
}

func (o LoadBalancerArrayOutput) Index(i pulumi.IntInput) LoadBalancerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LoadBalancer {
		return vs[0].([]*LoadBalancer)[vs[1].(int)]
	}).(LoadBalancerOutput)
}

type LoadBalancerMapOutput struct{ *pulumi.OutputState }

func (LoadBalancerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoadBalancer)(nil)).Elem()
}

func (o LoadBalancerMapOutput) ToLoadBalancerMapOutput() LoadBalancerMapOutput {
	return o
}

func (o LoadBalancerMapOutput) ToLoadBalancerMapOutputWithContext(ctx context.Context) LoadBalancerMapOutput {
	return o
}

func (o LoadBalancerMapOutput) MapIndex(k pulumi.StringInput) LoadBalancerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LoadBalancer {
		return vs[0].(map[string]*LoadBalancer)[vs[1].(string)]
	}).(LoadBalancerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerInput)(nil)).Elem(), &LoadBalancer{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerArrayInput)(nil)).Elem(), LoadBalancerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerMapInput)(nil)).Elem(), LoadBalancerMap{})
	pulumi.RegisterOutputType(LoadBalancerOutput{})
	pulumi.RegisterOutputType(LoadBalancerArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancerMapOutput{})
}
