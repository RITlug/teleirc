// Package mediashare provides a simple file hosting service with HTML5 player support.
package mediashare

import (
	"crypto/rand"
	"fmt"
	"io"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// FileInfo contains metadata about a stored file.
type FileInfo struct {
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

// Storage handles file storage operations.
type Storage struct {
	BasePath    string
	MaxFileSize int64
	db          *Database
}

// NewStorage creates a new Storage instance.
func NewStorage(basePath string, maxFileSize int64, db *Database) *Storage {
	return &Storage{
		BasePath:    basePath,
		MaxFileSize: maxFileSize,
		db:          db,
	}
}

// GenerateID creates a random 5-character alphanumeric ID.
func GenerateID() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 5)
	rand.Read(b)
	for i := range b {
		b[i] = charset[int(b[i])%len(charset)]
	}
	return string(b)
}

// GetDatePath returns a path in YYYY/MM/DD format for the current date.
func GetDatePath() string {
	now := time.Now()
	return fmt.Sprintf("%d/%02d/%02d", now.Year(), now.Month(), now.Day())
}

// Store saves a file and returns its FileInfo.
func (s *Storage) Store(filename string, username string, reader io.Reader) (*FileInfo, error) {
	// Generate unique ID (check database for collisions)
	var id string
	for {
		id = GenerateID()
		if s.db != nil {
			exists, err := s.db.Exists(id)
			if err != nil {
				return nil, fmt.Errorf("failed to check ID existence: %w", err)
			}
			if !exists {
				break
			}
		} else {
			break
		}
	}

	// Sanitize and get extension
	ext := strings.ToLower(filepath.Ext(filename))
	if ext == "" {
		ext = ".bin"
	}

	// Create date-based directory
	datePath := GetDatePath()
	dirPath := filepath.Join(s.BasePath, datePath)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	// Full path for the file
	storedFilename := id + ext
	fullPath := filepath.Join(dirPath, storedFilename)

	// Create file
	file, err := os.Create(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Copy content with size limit
	limitedReader := io.LimitReader(reader, s.MaxFileSize+1)
	written, err := io.Copy(file, limitedReader)
	if err != nil {
		os.Remove(fullPath)
		return nil, fmt.Errorf("failed to write file: %w", err)
	}

	if written > s.MaxFileSize {
		os.Remove(fullPath)
		return nil, fmt.Errorf("file exceeds maximum size of %d bytes", s.MaxFileSize)
	}

	// Detect content type
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	now := time.Now()
	relativePath := filepath.Join(datePath, storedFilename)

	info := &FileInfo{
		ID:          id,
		Filename:    filename,
		ContentType: contentType,
		Size:        written,
		Path:        relativePath,
		Username:    username,
		UploadedAt:  now,
	}

	// Save to database if available
	if s.db != nil {
		record := &FileRecord{
			ID:          id,
			Filename:    filename,
			ContentType: contentType,
			Size:        written,
			Path:        relativePath,
			Username:    username,
			UploadedAt:  now,
		}
		if err := s.db.Insert(record); err != nil {
			os.Remove(fullPath)
			return nil, fmt.Errorf("failed to save file record: %w", err)
		}
	}

	return info, nil
}

// Get retrieves file information by ID from the database.
func (s *Storage) Get(id string) (*FileInfo, error) {
	if s.db == nil {
		return nil, fmt.Errorf("database not available")
	}

	record, err := s.db.Get(id)
	if err != nil {
		return nil, err
	}
	if record == nil {
		return nil, nil
	}

	return &FileInfo{
		ID:           record.ID,
		Filename:     record.Filename,
		ContentType:  record.ContentType,
		Size:         record.Size,
		Path:         record.Path,
		Username:     record.Username,
		UploadedAt:   record.UploadedAt,
		LastOpenedAt: record.LastOpenedAt,
		OpenCount:    record.OpenCount,
	}, nil
}

// UpdateLastOpened updates the last opened timestamp for a file.
func (s *Storage) UpdateLastOpened(id string) error {
	if s.db == nil {
		return nil
	}
	return s.db.UpdateLastOpened(id)
}

// GetFullPath returns the full filesystem path for a file.
func (s *Storage) GetFullPath(relativePath string) string {
	return filepath.Join(s.BasePath, relativePath)
}

// Delete removes a file by ID.
func (s *Storage) Delete(id string) error {
	info, err := s.Get(id)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("file not found: %s", id)
	}

	// Remove from filesystem
	fullPath := s.GetFullPath(info.Path)
	if err := os.Remove(fullPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove file: %w", err)
	}

	// Remove from database
	if s.db != nil {
		if err := s.db.Delete(id); err != nil {
			return fmt.Errorf("failed to delete record: %w", err)
		}
	}

	return nil
}

// GetExpiredFiles returns files that haven't been opened within the retention period.
func (s *Storage) GetExpiredFiles(retentionHours int) ([]*FileInfo, error) {
	if s.db == nil {
		return nil, nil
	}

	cutoff := time.Now().Add(-time.Duration(retentionHours) * time.Hour)
	records, err := s.db.GetFilesNotOpenedSince(cutoff)
	if err != nil {
		return nil, err
	}

	var infos []*FileInfo
	for _, r := range records {
		infos = append(infos, &FileInfo{
			ID:           r.ID,
			Filename:     r.Filename,
			ContentType:  r.ContentType,
			Size:         r.Size,
			Path:         r.Path,
			Username:     r.Username,
			UploadedAt:   r.UploadedAt,
			LastOpenedAt: r.LastOpenedAt,
			OpenCount:    r.OpenCount,
		})
	}

	return infos, nil
}

// IsVideo checks if the content type is a video.
func IsVideo(contentType string) bool {
	return strings.HasPrefix(contentType, "video/")
}

// IsAudio checks if the content type is audio.
func IsAudio(contentType string) bool {
	return strings.HasPrefix(contentType, "audio/")
}

// IsImage checks if the content type is an image.
func IsImage(contentType string) bool {
	return strings.HasPrefix(contentType, "image/")
}
