package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareArgoTunnelSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: "The account identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"secret": {
			Type:      schema.TypeString,
			Required:  true,
			Sensitive: true,
			ForceNew:  true,
		},
		"cname": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"tunnel_token": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
