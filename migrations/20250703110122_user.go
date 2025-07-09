package migrations

import "gofr.dev/pkg/gofr/migration"

const createTableUsersSQL = `
CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);`

// createTableUsers returns a migration.Migrate object for creating the users table.
func createTableUsers() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			// Execute the SQL statement to create the table.
			_, err := d.SQL.Exec(createTableUsersSQL)
			if err != nil {
				return err
			}
			return nil
		},
		// DOWN: func(d migration.Datasource) error {
		// 	_, err := d.SQL.Exec("DROP TABLE IF EXISTS users;")
		// 	return err
		// },
	}
}
