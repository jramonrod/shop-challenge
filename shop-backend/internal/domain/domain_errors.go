package domain

import (
	"errors"
	"fmt"
)

// ErrInvalidUUID indicates that a provided UUID could not be parsed
var ErrInvalidUUID = errors.New("invalid UUID")
var ErrProductNotFound = errors.New("product not found")

// Wrapf enriches an existing error with additional context while preserving the original error for proper error
// unwrapping and classification.
//
// Example usage:
//
//	if err := someOperation(); err != nil {
//	    return Wrapf(err, "failed to execute operation for user %s", userID)
//	}
//
// This pattern ensures that errors remain informative at higher layers without losing their semantic meaning for
// lower-layer error detection.
func Wrapf(err error, fmtStr string, args ...interface{}) error {
	return fmt.Errorf("%s: %w", fmt.Sprintf(fmtStr, args...), err)
}
