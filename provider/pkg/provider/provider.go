package provider

import (
	"github.com/jhoeflaken/pulumi-azure-ext-native/provider/pkg/provider/keyvault/certificate"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

func Provider() p.Provider {
	var resources []infer.InferredResource

	var resource = infer.Resource[*certificate.KeyVaultCertificate, certificate.KeyVaultCertificateArgs, certificate.KeyVaultCertificateState]()
	resources = append(resources, resource)

	var config = infer.Config[*Config]()

	var metadata = schema.Metadata{
		DisplayName: "Azure Extension",
		Description: "The Pulumi Azure Extension Provider enables you to manage resources not supported by the Pulumi azure Native Provider, such as vault certificates.",
		Keywords: []string{
			"azure",
			"extension",
		},
		Homepage:   "",
		License:    "apache-2.0",
		Repository: "",
		LogoURL:    "",
		LanguageMap: map[string]any{
			"csharp": map[string]any{
				"respectSchemaVersion": true,
				"packageReferences": map[string]any{
					"Pulumi": "3.*",
				},
			},
			"go": map[string]any{
				"respectSchemaVersion":           true,
				"generateResourceContainertypes": true,
				"importBasePath":                 "github.com/jhoeflaken/pulumi-azure-ext/provider",
			},
			"nodejs": map[string]any{
				"respectSchemaVersion": true,
				"dependencies": map[string]any{
					"@pulumi/pulumi": "^3.0.0",
				},
			},
			"python": map[string]any{
				"respectSchemaVersion": true,
				"requires": map[string]any{
					"pulumi": ">-3.0.0,<4.0.0",
				},
				"pyproject": map[string]any{
					"enabled": true,
				},
			},
			"java": map[string]any{
				"buildFiles":                      "gradle",
				"gradleNexusPublishPluginVersion": "1.1.0",
				"dependencies": map[string]any{
					"com.pulumi:pulumi": "0.12.0",
				},
			},
		},
	}

	return infer.Provider(infer.Options{
		Metadata:  metadata,
		Resources: resources,
		Config:    config,
	})
}

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
