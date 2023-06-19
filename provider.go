package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type Config struct {
	ApiOrigin string
	Token     string
}

func Provider() *schema.Provider {
	p := &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"demo2_box":  resourceBox(),
			"demo2_text": resourceText(),
		},
		Schema: map[string]*schema.Schema{
			"api_origin": {
				Type:     schema.TypeString,
				Required: true,
			},
			"token": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
	p.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		config := Config{
			ApiOrigin: d.Get("api_origin").(string),
			Token:     d.Get("token").(string),
		}
		return &config, nil
	}
	return p
}
