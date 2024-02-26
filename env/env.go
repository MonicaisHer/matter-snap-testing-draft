package env

import (
	"os"
	"strconv"
)

// Environment variables, used to override defaults
const (
	// Channel/Revision of the service snap (has default)
	envSnapChannel = "SNAP_CHANNEL"

	// Path to snap instead, used for testing a local snap instead of
	// downloading from the store
	envSnapPath = "SNAP_PATH"

	// Toggle the teardown operations during tests (has default)
	envTeardown = "TEARDOWN"
)

var (
	// Defaults
	snapChannel = "latest/edge"
	snapPath    = ""
	teardown    = true
)

func init() {
	loadEnvVars()
}

// Read environment variables and perform type conversion/casting
func loadEnvVars() {

	if v := os.Getenv(envSnapChannel); v != "" {
		snapChannel = v
	}

	if v := os.Getenv(envSnapPath); v != "" {
		snapPath = v
	}

	if v := os.Getenv(envTeardown); v != "" {
		var err error
		teardown, err = strconv.ParseBool(v)
		if err != nil {
			panic(err)
		}
	}
}

// SnapChannel returns the set snap channel
func SnapChannel() string {
	return snapChannel
}

// SnapPath returns the set path to a local snap
func SnapPath() string {
	return snapPath
}

// SkipTeardownRemoval return
func Teardown() (skip bool) {
	return teardown
}
