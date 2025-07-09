package migrations

import (
	"gofr.dev/pkg/gofr/migration"
)

func All() map[int64]migration.Migrate {
	return map[int64]migration.Migrate{

		20250703110104: createTableTasks(),
		20250703110122: createTableUsers(),
	}
}
