package migrations

import "gofr.dev/pkg/gofr/migration"

const createTable = `
CREATE TABLE IF NOT EXISTS tasks (
    id INT PRIMARY KEY,
    description VARCHAR(255) NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT FALSE
);`

func createTableTasks() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			// Execute the SQL statement to create the table.
			_, err := d.SQL.Exec(createTable)
			if err != nil {
				return err
			}
			return nil
		},
	}
}
