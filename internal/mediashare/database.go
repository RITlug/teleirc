package mediashare

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

// FileRecord represents a file entry in the database.
type FileRecord struct {
	ID           string
	Filename     string
	ContentType  string
	Size         int64
	Path         string
	Username     string
	UploadedAt   time.Time
	LastOpenedAt *time.Time
	OpenCount    int
}

// Database wraps SQLite operations for MediaShare.
type Database struct {
	db *sql.DB
}

// NewDatabase creates a new database connection and initializes schema.
func NewDatabase(dbPath string) (*Database, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Enable WAL mode for better concurrent access
	if _, err := db.Exec("PRAGMA journal_mode=WAL"); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to enable WAL mode: %w", err)
	}

	database := &Database{db: db}
	if err := database.initSchema(); err != nil {
		db.Close()
		return nil, err
	}

	return database, nil
}

// initSchema creates the files table if it doesn't exist.
func (d *Database) initSchema() error {
	schema := `
	CREATE TABLE IF NOT EXISTS files (
		id TEXT PRIMARY KEY,
		filename TEXT NOT NULL,
		content_type TEXT NOT NULL,
		size INTEGER NOT NULL,
		path TEXT NOT NULL,
		username TEXT,
		uploaded_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		last_opened_at DATETIME,
		open_count INTEGER DEFAULT 0
	);
	CREATE INDEX IF NOT EXISTS idx_last_opened ON files(last_opened_at);
	CREATE INDEX IF NOT EXISTS idx_uploaded_at ON files(uploaded_at);
	`
	_, err := d.db.Exec(schema)
	if err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}
	return nil
}

// Insert adds a new file record to the database.
func (d *Database) Insert(record *FileRecord) error {
	query := `
	INSERT INTO files (id, filename, content_type, size, path, username, uploaded_at, last_opened_at, open_count)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := d.db.Exec(query,
		record.ID,
		record.Filename,
		record.ContentType,
		record.Size,
		record.Path,
		record.Username,
		record.UploadedAt,
		record.LastOpenedAt,
		record.OpenCount,
	)
	if err != nil {
		return fmt.Errorf("failed to insert file record: %w", err)
	}
	return nil
}

// Get retrieves a file record by ID.
func (d *Database) Get(id string) (*FileRecord, error) {
	query := `
	SELECT id, filename, content_type, size, path, username, uploaded_at, last_opened_at, open_count
	FROM files WHERE id = ?
	`
	row := d.db.QueryRow(query, id)

	var record FileRecord
	var lastOpenedAt sql.NullTime
	var username sql.NullString

	err := row.Scan(
		&record.ID,
		&record.Filename,
		&record.ContentType,
		&record.Size,
		&record.Path,
		&username,
		&record.UploadedAt,
		&lastOpenedAt,
		&record.OpenCount,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get file record: %w", err)
	}

	if lastOpenedAt.Valid {
		record.LastOpenedAt = &lastOpenedAt.Time
	}
	if username.Valid {
		record.Username = username.String
	}

	return &record, nil
}

// UpdateLastOpened updates the last_opened_at timestamp and increments open_count.
func (d *Database) UpdateLastOpened(id string) error {
	query := `
	UPDATE files
	SET last_opened_at = CURRENT_TIMESTAMP, open_count = open_count + 1
	WHERE id = ?
	`
	_, err := d.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to update last opened: %w", err)
	}
	return nil
}

// GetFilesNotOpenedSince returns files that haven't been opened since the cutoff time.
// If a file was never opened, it uses uploaded_at for comparison.
func (d *Database) GetFilesNotOpenedSince(cutoff time.Time) ([]*FileRecord, error) {
	query := `
	SELECT id, filename, content_type, size, path, username, uploaded_at, last_opened_at, open_count
	FROM files
	WHERE (last_opened_at IS NOT NULL AND last_opened_at < ?)
	   OR (last_opened_at IS NULL AND uploaded_at < ?)
	`
	rows, err := d.db.Query(query, cutoff, cutoff)
	if err != nil {
		return nil, fmt.Errorf("failed to query expired files: %w", err)
	}
	defer rows.Close()

	var records []*FileRecord
	for rows.Next() {
		var record FileRecord
		var lastOpenedAt sql.NullTime
		var username sql.NullString

		err := rows.Scan(
			&record.ID,
			&record.Filename,
			&record.ContentType,
			&record.Size,
			&record.Path,
			&username,
			&record.UploadedAt,
			&lastOpenedAt,
			&record.OpenCount,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan file record: %w", err)
		}

		if lastOpenedAt.Valid {
			record.LastOpenedAt = &lastOpenedAt.Time
		}
		if username.Valid {
			record.Username = username.String
		}

		records = append(records, &record)
	}

	return records, nil
}

// Delete removes a file record by ID.
func (d *Database) Delete(id string) error {
	query := `DELETE FROM files WHERE id = ?`
	_, err := d.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete file record: %w", err)
	}
	return nil
}

// Exists checks if a file with given ID exists in the database.
func (d *Database) Exists(id string) (bool, error) {
	query := `SELECT 1 FROM files WHERE id = ? LIMIT 1`
	row := d.db.QueryRow(query, id)
	var exists int
	err := row.Scan(&exists)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("failed to check existence: %w", err)
	}
	return true, nil
}

// GetRecentFiles returns the most recent files, ordered by upload date descending.
func (d *Database) GetRecentFiles(limit int) ([]*FileRecord, error) {
	query := `
	SELECT id, filename, content_type, size, path, username, uploaded_at, last_opened_at, open_count
	FROM files
	ORDER BY uploaded_at DESC
	LIMIT ?
	`
	rows, err := d.db.Query(query, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to query recent files: %w", err)
	}
	defer rows.Close()

	var records []*FileRecord
	for rows.Next() {
		var record FileRecord
		var lastOpenedAt sql.NullTime
		var username sql.NullString

		err := rows.Scan(
			&record.ID,
			&record.Filename,
			&record.ContentType,
			&record.Size,
			&record.Path,
			&username,
			&record.UploadedAt,
			&lastOpenedAt,
			&record.OpenCount,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan file record: %w", err)
		}

		if lastOpenedAt.Valid {
			record.LastOpenedAt = &lastOpenedAt.Time
		}
		if username.Valid {
			record.Username = username.String
		}

		records = append(records, &record)
	}

	return records, nil
}

// Close closes the database connection.
func (d *Database) Close() error {
	return d.db.Close()
}
