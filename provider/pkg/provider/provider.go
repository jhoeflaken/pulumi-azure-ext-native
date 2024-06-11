package provider

import (
	cfg "github.com/jhoeflaken/pulumi-azure-ext-native/provider/pkg/provider/config"
	"github.com/jhoeflaken/pulumi-azure-ext-native/provider/pkg/provider/keyvault/certificate"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
)

func Provider() p.Provider {
	var resources []infer.InferredResource

	var resource = infer.Resource[*certificate.KeyVaultCertificate, certificate.KeyVaultCertificateArgs, certificate.KeyVaultCertificateState]()
	resources = append(resources, resource)

	var config = infer.Config[*cfg.Config]()

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
