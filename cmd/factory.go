// Copyright (C) 2025 XLR8discovery PBC
// See LICENSE for copying information.

package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/zeebo/errs/v2"

	"xlr8d.io/oss-up/pkg/runtime/compose"
	"xlr8d.io/oss-up/pkg/runtime/runtime"
	"xlr8d.io/oss-up/pkg/runtime/standalone"
)

// FromDir creates the right runtime based on available file names in the directory.
func FromDir(dir string) (runtime.Runtime, error) {
	_, err := os.Stat(filepath.Join(dir, "docker-compose.yaml"))
	if err == nil {
		return compose.NewCompose(dir)
	}

	_, err = os.Stat(filepath.Join(dir, "supervisord.conf"))
	if err == nil {
		ossProjectDir := os.Getenv("OSS_PROJECT_DIR")
		if ossProjectDir == "" {
			return nil, errs.Errorf("Please set \"OSS_PROJECT_DIR\" environment variable with the location of your checked out oss/oss project. (Required to use web resources")
		}
		gatewayProjectDir := os.Getenv("GATEWAY_PROJECT_DIR")
		if gatewayProjectDir == "" {
			fmt.Println("WARNING: \"GATEWAY_PROJECT_DIR\" environment variable not set! Please set or add -g flag with the location of your checked out oss/gateway-mt project to use web resources.")
			gatewayProjectDir = "/tmp"
		}
		return standalone.NewStandalone(standalone.Paths{
			ScriptDir:  dir,
			OSSDir:   ossProjectDir,
			GatewayDir: gatewayProjectDir,
			CleanDir:   false,
		})
	}

	return nil, errors.New("directory doesn't contain supported deployment descriptor")
}