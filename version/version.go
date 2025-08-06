package version

import (
	"fmt"
)

const (
	Version        = "dev"
	CommitHash     = "unknown"
	BuildTimestamp = "unknown"
)

func BuildVersion() string {
	return fmt.Sprintf("%s-%s (%s)", Version, CommitHash, BuildTimestamp)
}
