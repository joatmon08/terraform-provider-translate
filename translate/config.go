package translate

import (
	"context"
	"log"

	"cloud.google.com/go/translate"
)

type Config struct {
	Context context.Context
	Client  *translate.Client
}

func (c *Config) LoadAndValidate() error {
	c.Context = context.Background()
	client, err := translate.NewClient(c.Context)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return err
	}

	c.Client = client
	return nil
}
