package provider

import (
	"fmt"
	"github.com/alexashley/terraform-provider-kong/kong/kong"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceKongConsumer() *schema.Resource {
	return &schema.Resource{
		Create: resourceKongConsumerCreate,
		Read:   resourceKongConsumerRead,
		Update: resourceKongConsumerUpdate,
		Delete: resourceKongConsumerDelete,
		Importer: &schema.ResourceImporter{
			State: importResourceIfUuidIsValid,
		},
		Schema: map[string]*schema.Schema{
			"username": {
				Description: "A unique username representing a consumer of the API.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"custom_id": {
				Description: "A unique identifier representing a user or service of your API. It can be used to map to existing users in your database.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func resourceKongConsumerCreate(data *schema.ResourceData, meta interface{}) error {
	client := meta.(*kong.KongClient)

	username := data.Get("username").(string)
	customId := data.Get("custom_id").(string)

	if username == "" && customId == "" {
		return fmt.Errorf("At least one of username or custom_id must be supplied")
	}

	consumer, err := client.CreateConsumer(mapToApiCustomer(data))

	if err != nil {
		return err
	}

	data.SetId(consumer.Id)

	return resourceKongConsumerRead(data, meta)
}

func resourceKongConsumerRead(data *schema.ResourceData, meta interface{}) error {
	kongClient := meta.(*kong.KongClient)

	consumer, err := kongClient.GetConsumer(data.Id())

	if err != nil {
		if resourceDoesNotExistError(err) {
			data.SetId("")
			return nil
		}

		return err
	}

	data.Set("username", consumer.Username)
	data.Set("custom_id", consumer.CustomId)

	return nil
}

func resourceKongConsumerUpdate(data *schema.ResourceData, meta interface{}) error {
	kongClient := meta.(*kong.KongClient)

	err := kongClient.UpdateConsumer(mapToApiCustomer(data))
	if err != nil {
		return err
	}

	return resourceKongConsumerRead(data, meta)
}

func resourceKongConsumerDelete(data *schema.ResourceData, meta interface{}) error {
	kongClient := meta.(*kong.KongClient)

	err := kongClient.DeleteConsumer(data.Id())

	if resourceDoesNotExistError(err) {
		return nil
	}

	return err
}

func mapToApiCustomer(data *schema.ResourceData) *kong.KongConsumer {
	return &kong.KongConsumer{
		Id:       data.Id(),
		CustomId: data.Get("custom_id").(string),
		Username: data.Get("username").(string),
	}
}
