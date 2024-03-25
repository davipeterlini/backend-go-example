package repository // Generic Repository for CRUD

import (
	"context"
	"database/sql"
	"fmt"
	"sales/vehicle/pkg/config" // config for .env file

	_ "github.com/lib/pq" // PostgreSQL driver
)

type BaseRepository struct {
	db *sql.DB
}

func (repo *BaseRepository) DB() *sql.DB {
	return repo.db
}

func NewBaseRepository() (*BaseRepository, error) {
	dbConfig := config.LoadDatabaseConfig()
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("could not connect to the database: %v", err)
	}
	// Test the database connection
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping the database: %v", err)
	}

	return &BaseRepository{db: db}, nil
}

func (r *BaseRepository) Create(ctx context.Context, table string, columns []string, values []interface{}) error {
	placeholders := make([]string, len(values))
	for i := range values {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		table,
		join(columns, ", "),
		join(placeholders, ", "))
	_, err := r.db.ExecContext(ctx, query, values...)
	return err
}

func (r *BaseRepository) List(ctx context.Context, table string, columns []string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT %s FROM %s", join(columns, ", "), table)
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *BaseRepository) Read(ctx context.Context, table, idColumn string, idValue interface{}) (*sql.Row, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = $1", table, idColumn)
	return r.db.QueryRowContext(ctx, query, idValue), nil
}

func (r *BaseRepository) Update(ctx context.Context, table, idColumn string, idValue interface{}, columns []string, values []interface{}) error {
	setClauses := make([]string, len(columns))
	for i, col := range columns {
		setClauses[i] = fmt.Sprintf("%s = $%d", col, i+2) // Starting from $2 because $1 is reserved for the idValue
	}
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s = $1",
		table,
		join(setClauses, ", "),
		idColumn)
	_, err := r.db.ExecContext(ctx, query, append([]interface{}{idValue}, values...)...)
	return err
}

func (r *BaseRepository) Delete(ctx context.Context, table, idColumn string, idValue interface{}) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = $1", table, idColumn)
	_, err := r.db.ExecContext(ctx, query, idValue)
	return err
}

// Helper function to join elements of a slice into a single string separated by a delimiter.
func join(items []string, delimiter string) string {
	result := ""
	for i, item := range items {
		if i > 0 {
			result += delimiter
		}
		result += item
	}
	return result
}
