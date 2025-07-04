// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/dirien/pulumi-vultr/sdk/v2/go/vultr/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "vultr:index/bareMetalServer:BareMetalServer":
		r = &BareMetalServer{}
	case "vultr:index/blockStorage:BlockStorage":
		r = &BlockStorage{}
	case "vultr:index/containerRegistry:ContainerRegistry":
		r = &ContainerRegistry{}
	case "vultr:index/database:Database":
		r = &Database{}
	case "vultr:index/databaseConnectionPool:DatabaseConnectionPool":
		r = &DatabaseConnectionPool{}
	case "vultr:index/databaseDb:DatabaseDb":
		r = &DatabaseDb{}
	case "vultr:index/databaseQuota:DatabaseQuota":
		r = &DatabaseQuota{}
	case "vultr:index/databaseReplica:DatabaseReplica":
		r = &DatabaseReplica{}
	case "vultr:index/databaseTopic:DatabaseTopic":
		r = &DatabaseTopic{}
	case "vultr:index/databaseUser:DatabaseUser":
		r = &DatabaseUser{}
	case "vultr:index/dnsDomain:DnsDomain":
		r = &DnsDomain{}
	case "vultr:index/dnsRecord:DnsRecord":
		r = &DnsRecord{}
	case "vultr:index/firewallGroup:FirewallGroup":
		r = &FirewallGroup{}
	case "vultr:index/firewallRule:FirewallRule":
		r = &FirewallRule{}
	case "vultr:index/inference:Inference":
		r = &Inference{}
	case "vultr:index/instance:Instance":
		r = &Instance{}
	case "vultr:index/instanceIpv4:InstanceIpv4":
		r = &InstanceIpv4{}
	case "vultr:index/isoPrivate:IsoPrivate":
		r = &IsoPrivate{}
	case "vultr:index/kubernetes:Kubernetes":
		r = &Kubernetes{}
	case "vultr:index/kubernetesNodePools:KubernetesNodePools":
		r = &KubernetesNodePools{}
	case "vultr:index/loadBalancer:LoadBalancer":
		r = &LoadBalancer{}
	case "vultr:index/objectStorage:ObjectStorage":
		r = &ObjectStorage{}
	case "vultr:index/reservedIp:ReservedIp":
		r = &ReservedIp{}
	case "vultr:index/reverseIpv4:ReverseIpv4":
		r = &ReverseIpv4{}
	case "vultr:index/reverseIpv6:ReverseIpv6":
		r = &ReverseIpv6{}
	case "vultr:index/sSHKey:SSHKey":
		r = &SSHKey{}
	case "vultr:index/snapshot:Snapshot":
		r = &Snapshot{}
	case "vultr:index/snapshotFromUrl:SnapshotFromUrl":
		r = &SnapshotFromUrl{}
	case "vultr:index/startupScript:StartupScript":
		r = &StartupScript{}
	case "vultr:index/user:User":
		r = &User{}
	case "vultr:index/virtualFileSystemStorage:VirtualFileSystemStorage":
		r = &VirtualFileSystemStorage{}
	case "vultr:index/vpc2:Vpc2":
		r = &Vpc2{}
	case "vultr:index/vpc:Vpc":
		r = &Vpc{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:vultr" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"vultr",
		"index/bareMetalServer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/blockStorage",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/containerRegistry",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/database",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/databaseConnectionPool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/databaseDb",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/databaseQuota",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/databaseReplica",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/databaseTopic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/databaseUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/dnsDomain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/dnsRecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/firewallGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/firewallRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/inference",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/instance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/instanceIpv4",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/isoPrivate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/kubernetes",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/kubernetesNodePools",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/loadBalancer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/objectStorage",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/reservedIp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/reverseIpv4",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/reverseIpv6",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/sSHKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/snapshot",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/snapshotFromUrl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/startupScript",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/user",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/virtualFileSystemStorage",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/vpc",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vultr",
		"index/vpc2",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"vultr",
		&pkg{version},
	)
}
