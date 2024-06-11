package config

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/pkg/errors"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Config struct {
	Credential *azidentity.DefaultAzureCredential
}

var _ = (infer.CustomConfigure)((*Config)(nil))

// Configure the provider with the given configuration.
func (c *Config) Configure(p.Context) error {
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return errors.Wrap(err, "Failed to create default azure credential.")
	}

	c.Credential = credential
	return nil
}
