package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"terraform-provider-eod/db"
	"time"
)

func resourceRucheDb() *schema.Resource {
	return &schema.Resource{
		Create: resourceRucheDBCreate,
		Read:   resourceRucheDBRead,
		Delete: resourceRucheDBDelete,
		Exists: resourceRucheDBExists,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"expires_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"retention": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceRucheDBDelete(data *schema.ResourceData, i interface{}) error {
	dbClient := i.(*db.Client)
	return dbClient.Delete(data.Get("name").(string))
}

func resourceRucheDBExists(data *schema.ResourceData, i interface{}) (bool, error) {
	dbClient := i.(*db.Client)
	_, err := dbClient.Read(data.Get("name").(string))
	if err != nil {
		switch err.(type) {
		case *db.NotFound:
			return false, nil
		default:
			return true, err
		}
	}
	return true, nil
}

func resourceRucheDBRead(data *schema.ResourceData, i interface{}) error {
	dbClient := i.(*db.Client)
	env, err := dbClient.Read(data.Get("name").(string))
	if err != nil {
		return err
	}
	data.Set("port", env.Port)
	data.Set("expires_at", env.ExpiresAt)
	data.Set("status", env.Status)
	return nil
}

func resourceRucheDBCreate(data *schema.ResourceData, i interface{}) error {
	dbClient := i.(*db.Client)
	env, err := dbClient.Create(data.Get("retention").(int))
	if err != nil {
		return err
	}
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	data.Set("name", env.ID)
	return resourceRucheDBRead(data, i)
}
