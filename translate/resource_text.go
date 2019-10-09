package translate

import (
	"fmt"
	"hash/fnv"
	"strings"

	"cloud.google.com/go/translate"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"golang.org/x/text/language"
)

func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return fmt.Sprint(h.Sum32())
}

func resourceText() *schema.Resource {
	return &schema.Resource{
		Create: resourceTextCreate,
		Read:   resourceTextRead,
		Update: resourceTextUpdate,
		Delete: resourceTextDelete,

		Schema: map[string]*schema.Schema{
			"text": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"source_language": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"target_language": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"translated_text": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceTextCreate(d *schema.ResourceData, m interface{}) error {
	return resourceTextRead(d, m)
}

func resourceTextRead(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	source, err := language.Parse(d.Get("source_language").(string))
	if err != nil {
		return fmt.Errorf("source invalid: %v", err)
	}
	target, err := language.Parse(d.Get("target_language").(string))
	if err != nil {
		return fmt.Errorf("target invalid: %v", err)
	}

	options := &translate.Options{
		Source: source,
		Format: translate.Text,
	}

	text := d.Get("text").(string)
	body := strings.Split(text, "\n")

	translations, err := config.Client.Translate(config.Context, body, target, options)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("translate failed: %v", err)
	}
	d.SetId(hash(text))
	fullText := fullTranslation(translations)
	d.Set("translated_text", fullText)
	return nil
}

func fullTranslation(translations []translate.Translation) string {
	full := ""
	for _, t := range translations {
		full = full + t.Text + "\n"
	}
	return full
}

func resourceTextUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceTextRead(d, m)
}

func resourceTextDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
