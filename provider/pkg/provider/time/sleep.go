package time

import (
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Sleep struct{}

// The following statements are not required. They are type assertions to indicate to Go that KeyVaultCertificate
// should implement the following interfaces. If the function signature doesn't match or isn't implemented,
// we get nice compile time errors at this location.
var _ = (infer.Annotated)((*Sleep)(nil))

// Annotate is a method on KeyVaultCertificate that provides a description of the resource.
func (c *Sleep) Annotate(a infer.Annotator) {
	a.Describe(&c, "Sleep for a specified amount of time.")
}

type SleepArgs struct {
	Duration *string `pulumi:"duration"`
}

type SleepState struct {
	Duration *string `pulumi:"duration"`
}
