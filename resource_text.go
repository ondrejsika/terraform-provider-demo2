package main

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceTextCreate(d *schema.ResourceData, m interface{}) error {
	box_id := d.Get("box_id").(int)
	text := d.Get("text").(string)

	response, err := CreateTextApi(m.(*Config).ApiOrigin, m.(*Config).Token, box_id, text)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	d.SetId(strconv.Itoa(response.TextId))
	d.Set("box_id", box_id)
	d.Set("text", text)
	return nil
}

func resourceTextRead(d *schema.ResourceData, m interface{}) error {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	box_id := d.Get("box_id").(int)

	response, err := GetTextApi(m.(*Config).ApiOrigin, m.(*Config).Token, box_id, id)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	d.Set("text", response.Text)
	return nil
}

func resourceTextUpdate(d *schema.ResourceData, m interface{}) error {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	box_id := d.Get("box_id").(int)
	text := d.Get("text").(string)

	_, err = UpdateTextApi(m.(*Config).ApiOrigin, m.(*Config).Token, box_id, id, text)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	d.Set("text", text)
	return nil
}

func resourceTextDelete(d *schema.ResourceData, m interface{}) error {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	box_id := d.Get("box_id").(int)

	_, err = DeleteTextApi(m.(*Config).ApiOrigin, m.(*Config).Token, box_id, id)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	return nil
}

func resourceText() *schema.Resource {
	return &schema.Resource{
		Create: resourceTextCreate,
		Update: resourceTextUpdate,
		Read:   resourceTextRead,
		Delete: resourceTextDelete,

		Schema: map[string]*schema.Schema{
			"box_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"text": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
