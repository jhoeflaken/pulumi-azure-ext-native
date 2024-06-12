package certificate

import (
	"github.com/pulumi/pulumi-go-provider/infer"
)

type KeyVaultCertificate struct{}

// The following statements are not required. They are type assertions to indicate to Go that KeyVaultCertificate
// should implement the following interfaces. If the function signature doesn't match or isn't implemented,
// we get nice compile time errors at this location.
var _ = (infer.Annotated)((*KeyVaultCertificate)(nil))

// Annotate is a method on KeyVaultCertificate that provides a description of the resource.
func (c *KeyVaultCertificate) Annotate(a infer.Annotator) {
	a.Describe(&c, "An Azure key vault certificate.")
}

type KeyVaultCertificateArgs struct {
	VaultName                *string             `pulumi:"vaultName"`
	Base64EncodedCertificate *string             `pulumi:"base64EncodedCertificate"`
	Password                 *string             `pulumi:"password,optional"`
	Tags                     *map[string]*string `pulumi:"tags,optional"`
}

type KeyVaultCertificateState struct {
	KeyVaultCertificateArgs
	Name    string `pulumi:"name"`
	Version string `pulumi:"version"`
}
