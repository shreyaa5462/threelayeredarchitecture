package migrations

import "gofr.dev/pkg/gofr/migration"

const createTableUserSQL = `
CREATE TABLE IF NOT EXISTS user (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);`

// CreateTableUser creates the 'users' table.
func CreateTableUser() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(createTableUserSQL)
			if err != nil {
				return err
			}
			return nil
		},
	}
}
