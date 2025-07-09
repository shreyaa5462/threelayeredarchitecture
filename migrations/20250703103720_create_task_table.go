package migrations

import "gofr.dev/pkg/gofr/migration"

const createTableTaskSQL = `
CREATE TABLE IF NOT EXISTS task (
    id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
    description TEXT NOT NULL,
    completed BOOLEAN NOT NULL
);
`

// CreateTableTask creates the 'tasks' table.
func CreateTableTask() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(createTableTaskSQL)
			if err != nil {
				return err
			}
			return nil
		},
	}
}
