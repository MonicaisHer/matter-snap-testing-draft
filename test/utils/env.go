package utils

import (
	"os"
	"strconv"
)

const (
	// environment variables
	// used to override defaults
	serviceChannelEnv    = "SERVICE_CHANNEL"     // channel/revision of the service snap (has default)
	localServiceSnapEnv  = "LOCAL_SERVICE_SNAP"  // path to local service snap to be tested instead of downloading from a channel

	fullConfigTestEnv      = "FULL_CONFIG_TEST"      // toggle full config tests (has default)
	skipTeardownRemovalEnv = "SKIP_TEARDOWN_REMOVAL" // skip the removal of snaps during teardown
)

var (
	// global defaults
	ServiceChannel        = "latest/edge"
	LocalServiceSnapPath  = ""
	FullConfigTest        = false
	SkipTeardownRemoval   = false
)

func init() {
	if v := os.Getenv(serviceChannelEnv); v != "" {
		ServiceChannel = v
	}

	if v := os.Getenv(localServiceSnapEnv); v != "" {
		LocalServiceSnapPath = v
	}

	if v := os.Getenv(fullConfigTestEnv); v != "" {
		var err error
		FullConfigTest, err = strconv.ParseBool(v)
		if err != nil {
			panic(err)
		}
	}

	if v := os.Getenv(skipTeardownRemovalEnv); v != "" {
		var err error
		SkipTeardownRemoval, err = strconv.ParseBool(v)
		if err != nil {
			panic(err)
		}
	}
}
