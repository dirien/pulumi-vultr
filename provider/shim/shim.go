package shim

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vultr/terraform-provider-vultr/vultr"
)

func NewProvider() *schema.Provider {
	p := vultr.Provider()
	return p
}
