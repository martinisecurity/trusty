package stats

import (
	"testing"

	"github.com/ekspand/trusty/internal/db/cadb"
	"github.com/ekspand/trusty/internal/db/orgsdb"
	"github.com/ekspand/trusty/tests/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/dig"
)

const (
	projFolder = "../../../"
)

func TestFactory(t *testing.T) {

	cfg, err := testutils.LoadConfig(projFolder, "UNIT_TEST")
	require.NoError(t, err)

	cadbp, err := cadb.New(
		cfg.CaSQL.Driver,
		cfg.CaSQL.DataSource,
		cfg.CaSQL.MigrationsDir,
		testutils.IDGenerator().NextID,
	)
	require.NoError(t, err)
	defer cadbp.Close()

	odbp, err := orgsdb.New(
		cfg.OrgsSQL.Driver,
		cfg.OrgsSQL.DataSource,
		cfg.OrgsSQL.MigrationsDir,
		testutils.IDGenerator().NextID,
	)
	require.NoError(t, err)
	defer odbp.Close()

	c := dig.New()
	c.Provide(func() (cadb.CaReadonlyDb, orgsdb.OrgsReadOnlyDb) {
		return cadbp, odbp
	})
	require.NoError(t, err)

	scheduler := &testutils.MockTask{}

	f := Factory(scheduler, "test_run", "Every 30 minutes")
	require.NotNil(t, f)

	err = c.Invoke(f)
	require.NoError(t, err)

	require.Len(t, scheduler.Tasks, 1)
	executed := scheduler.Tasks[0].Run()
	assert.True(t, executed)
}
