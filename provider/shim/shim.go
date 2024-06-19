package shim

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	prov "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vultr/terraform-provider-vultr/vultr"
)

func NewProvider() *schema.Provider {
	p := vultr.Provider()
	return p
}

var _ prov.Provider = &VultrProvider{}

type VultrProvider struct{}

func (a VultrProvider) Metadata(ctx context.Context, request prov.MetadataRequest, response *prov.MetadataResponse) {
}

func (a VultrProvider) Schema(ctx context.Context, request prov.SchemaRequest, response *prov.SchemaResponse) {

}

func (a VultrProvider) Configure(ctx context.Context, request prov.ConfigureRequest, response *prov.ConfigureResponse) {

}

func (a VultrProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (a VultrProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func Framework() func() prov.Provider {
	return func() prov.Provider {
		return &VultrProvider{}
	}
}
