package time

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	"time"
)

// The following statements are not required. They are type assertions to indicate to Go that Command implements the following interfaces.
// If the function signature doesn't match or isn't implemented, we get nice compile time errors at this location.

// They would normally be included in the commandController.go file, but they're located here for instructive purposes.
var _ = (infer.CustomResource[SleepArgs, SleepState])((*Sleep)(nil))

// The Create method will be run on every create.
func (c *Sleep) Create(ctx p.Context, name string, args SleepArgs, preview bool) (string, SleepState, error) {
	duration, err := time.ParseDuration(*args.Duration)
	if err != nil {
		return "", SleepState{}, errors.Wrap(err, "Error parsing sleep duration")
	}

	if preview {
		return "", SleepState{}, nil
	}

	logging.V(3).Infof("Waiting for %s to continue", *args.Duration)
	time.Sleep(duration)

	return uuid.NewString(), SleepState{Duration: args.Duration}, nil
}
