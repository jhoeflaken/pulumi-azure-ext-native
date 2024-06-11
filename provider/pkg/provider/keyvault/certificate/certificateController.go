package certificate

import (
	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azcertificates"
	cfg "github.com/jhoeflaken/pulumi-azure-ext-native/provider/pkg/provider/config"
	"github.com/pkg/errors"
	p "github.com/pulumi/pulumi-go-provider"

	"github.com/pulumi/pulumi-go-provider/infer"
)

// The following statements are not required. They are type assertions to indicate to Go that Command implements the following interfaces.
// If the function signature doesn't match or isn't implemented, we get nice compile time errors at this location.

// They would normally be included in the commandController.go file, but they're located here for instructive purposes.
var _ = (infer.CustomResource[KeyVaultCertificateArgs, KeyVaultCertificateState])((*KeyVaultCertificate)(nil))
var _ = (infer.CustomUpdate[KeyVaultCertificateArgs, KeyVaultCertificateState])((*KeyVaultCertificate)(nil))
var _ = (infer.CustomDelete[KeyVaultCertificateState])((*KeyVaultCertificate)(nil))

// The Create method will be run on every create.
func (c *KeyVaultCertificate) Create(ctx p.Context, name string, args KeyVaultCertificateArgs, preview bool) (string, KeyVaultCertificateState, error) {
	client, err := newClient(ctx, args.VaultName)
	if err != nil {
		return "", KeyVaultCertificateState{}, errors.Wrap(err, "Failed to create key vault client.")
	}

	if preview {
		return "", KeyVaultCertificateState{}, nil
	}

	resp, err := client.ImportCertificate(ctx, name, azcertificates.ImportCertificateParameters{}, nil)
	if err != nil {
		return "", KeyVaultCertificateState{}, errors.Wrap(err, "Failed to import certificate.")
	}

	return resp.ID.Name(), KeyVaultCertificateState{
		Name:    resp.ID.Name(),
		Version: resp.ID.Version(),
	}, nil
}

// The Update method will be run on every update.
func (c *KeyVaultCertificate) Update(ctx p.Context, name string, olds KeyVaultCertificateState, news KeyVaultCertificateArgs, preview bool) (KeyVaultCertificateState, error) {
	client, err := newClient(ctx, news.VaultName)
	if err != nil {
		return KeyVaultCertificateState{}, errors.Wrap(err, "Failed to create key vault client.")
	}

	if preview {
		return KeyVaultCertificateState{}, nil
	}

	resp, err := client.ImportCertificate(ctx, name, azcertificates.ImportCertificateParameters{}, nil)
	if err != nil {
		return KeyVaultCertificateState{}, errors.Wrap(err, "Failed to import certificate.")
	}

	return KeyVaultCertificateState{
		Name:    resp.ID.Name(),
		Version: resp.ID.Version(),
	}, nil
}

// The Delete method will run when the resource is deleted.
func (c *KeyVaultCertificate) Delete(ctx p.Context, id string, props KeyVaultCertificateState) error {
	client, err := newClient(ctx, props.VaultName)
	if err != nil {
		return errors.Wrap(err, "Failed to create key vault client.")
	}

	_, err = client.DeleteCertificate(ctx, id, nil)
	if err != nil {
		return errors.Wrap(err, "Failed to delete certificate.")
	}

	return nil
}

// Create a new client for managing the Azure Key Vault Certificates.
func newClient(ctx p.Context, vaultName string) (*azcertificates.Client, error) {
	var config = infer.GetConfig[*cfg.Config](ctx)
	client, err := azcertificates.NewClient(vaultName, config.Credential, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
