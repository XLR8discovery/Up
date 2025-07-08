// Copyright (C) 2025 XLR8discovery PBC
// See LICENSE for copying information.

package modify

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"xlr8d.io/oss-up/cmd/testdata"
)

func TestPersist(t *testing.T) {
	dir := t.TempDir()

	st, rt, err := testdata.InitCompose(dir)
	require.NoError(t, err)

	err = persist(st, rt, []string{"cockroach"})
	require.NoError(t, err)

	require.NoError(t, rt.Write())

	result, err := os.ReadFile(filepath.Join(dir, "docker-compose.yaml"))
	require.NoError(t, err)

	require.Contains(t, string(result), "source: "+filepath.Join(dir, "/cockroach/cockroach-data"))
	require.Contains(t, string(result), "target: /cockroach/cockroach-data")

}