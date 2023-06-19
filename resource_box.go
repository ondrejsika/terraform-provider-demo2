package main

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceBoxCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	response, err := CreateBoxApi(m.(*Config).ApiOrigin, m.(*Config).Token, name)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	d.SetId(strconv.Itoa(response.BoxId))
	d.Set("name", name)
	return nil
}

func resourceBoxRead(d *schema.ResourceData, m interface{}) error {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	response, err := GetBoxApi(m.(*Config).ApiOrigin, m.(*Config).Token, id)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	d.Set("name", response.Name)
	return nil
}

func resourceBoxUpdate(d *schema.ResourceData, m interface{}) error {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	name := d.Get("name").(string)

	_, err = UpdateBoxApi(m.(*Config).ApiOrigin, m.(*Config).Token, id, name)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	d.Set("name", name)
	return nil
}

func resourceBoxDelete(d *schema.ResourceData, m interface{}) error {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	_, err = DeleteBoxApi(m.(*Config).ApiOrigin, m.(*Config).Token, id)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	return nil
}

func resourceBox() *schema.Resource {
	return &schema.Resource{
		Create: resourceBoxCreate,
		Update: resourceBoxUpdate,
		Read:   resourceBoxRead,
		Delete: resourceBoxDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
