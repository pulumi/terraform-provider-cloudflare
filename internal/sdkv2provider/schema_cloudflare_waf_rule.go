package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWAFRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"rule_id": {
			Type:     schema.TypeString,
			Required: true,
		},

		consts.ZoneIDSchemaKey: {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
		},

		"group_id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"package_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"mode": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}
