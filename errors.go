package wallpapr

import (
	"errors"
	"fmt"
)

var errKeysNotInEnvVars = errors.New("authentication keys not found in os env")

type errUnexpectedStatusCode int

func (err errUnexpectedStatusCode) Error() string {
	return fmt.Sprintf("got unexpected status code: %d", err)
}
