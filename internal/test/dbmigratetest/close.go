package dbmigratetest

import (
	"github.com/pennsieve/repo-service/internal/dbmigrate"
	"github.com/pennsieve/repo-service/internal/test"
	"github.com/stretchr/testify/require"
)

func Close(t require.TestingT, migrator *dbmigrate.DatabaseMigrator) {
	test.Helper(t)
	srcErr, dbErr := migrator.Close()
	require.NoError(t, srcErr)
	require.NoError(t, dbErr)
}
