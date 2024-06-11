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
	VaultName                string
	Base64EncodedCertificate *string
	Password                 *string
	Tags                     map[string]*string
}

type KeyVaultCertificateState struct {
	KeyVaultCertificateArgs
	name    string
	version string
}
