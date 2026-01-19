package mediashare

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateID(t *testing.T) {
	id1 := GenerateID()
	id2 := GenerateID()

	if len(id1) != 5 {
		t.Errorf("Expected ID length 5, got %d", len(id1))
	}

	if id1 == id2 {
		t.Error("Generated IDs should be unique")
	}

	// Check alphanumeric
	for _, c := range id1 {
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')) {
			t.Errorf("Invalid character in ID: %c", c)
		}
	}
}

func setupTestDB(t *testing.T, tmpDir string) *Database {
	dbPath := filepath.Join(tmpDir, "test.db")
	db, err := NewDatabase(dbPath)
	if err != nil {
		t.Fatal(err)
	}
	return db
}

func TestStorage_Store(t *testing.T) {
	// Create temp directory
	tmpDir, err := os.MkdirTemp("", "mediashare_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	db := setupTestDB(t, tmpDir)
	defer db.Close()

	storage := NewStorage(tmpDir, 10*1024*1024, db)

	// Test storing a file
	content := "test content"
	info, err := storage.Store("test.txt", "testuser", strings.NewReader(content))
	if err != nil {
		t.Fatalf("Failed to store file: %v", err)
	}

	if info.ID == "" {
		t.Error("ID should not be empty")
	}

	if len(info.ID) != 5 {
		t.Errorf("Expected ID length 5, got %d", len(info.ID))
	}

	if info.Filename != "test.txt" {
		t.Errorf("Expected filename test.txt, got %s", info.Filename)
	}

	if info.Size != int64(len(content)) {
		t.Errorf("Expected size %d, got %d", len(content), info.Size)
	}

	if info.ContentType != "text/plain; charset=utf-8" {
		t.Errorf("Expected text/plain content type, got %s", info.ContentType)
	}

	if info.Username != "testuser" {
		t.Errorf("Expected username testuser, got %s", info.Username)
	}

	// Verify file exists
	fullPath := filepath.Join(tmpDir, info.Path)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		t.Error("File should exist on disk")
	}
}

func TestStorage_Get(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "mediashare_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	db := setupTestDB(t, tmpDir)
	defer db.Close()

	storage := NewStorage(tmpDir, 10*1024*1024, db)

	// Store a file first
	content := "test content for get"
	info, err := storage.Store("gettest.txt", "testuser", strings.NewReader(content))
	if err != nil {
		t.Fatalf("Failed to store file: %v", err)
	}

	// Now retrieve it
	retrieved, err := storage.Get(info.ID)
	if err != nil {
		t.Fatalf("Failed to get file: %v", err)
	}

	if retrieved.ID != info.ID {
		t.Errorf("ID mismatch: expected %s, got %s", info.ID, retrieved.ID)
	}

	if retrieved.Size != info.Size {
		t.Errorf("Size mismatch: expected %d, got %d", info.Size, retrieved.Size)
	}

	if retrieved.Username != "testuser" {
		t.Errorf("Username mismatch: expected testuser, got %s", retrieved.Username)
	}
}

func TestStorage_MaxFileSize(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "mediashare_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	db := setupTestDB(t, tmpDir)
	defer db.Close()

	// Set max size to 10 bytes
	storage := NewStorage(tmpDir, 10, db)

	// Try to store a file larger than limit
	content := "this content is definitely longer than 10 bytes"
	_, err = storage.Store("big.txt", "testuser", strings.NewReader(content))
	if err == nil {
		t.Error("Expected error for file exceeding max size")
	}
}

func TestIsVideo(t *testing.T) {
	tests := []struct {
		contentType string
		expected    bool
	}{
		{"video/mp4", true},
		{"video/webm", true},
		{"audio/mp3", false},
		{"image/png", false},
	}

	for _, tt := range tests {
		result := IsVideo(tt.contentType)
		if result != tt.expected {
			t.Errorf("IsVideo(%s) = %v, expected %v", tt.contentType, result, tt.expected)
		}
	}
}

func TestIsAudio(t *testing.T) {
	tests := []struct {
		contentType string
		expected    bool
	}{
		{"audio/mp3", true},
		{"audio/ogg", true},
		{"video/mp4", false},
		{"image/png", false},
	}

	for _, tt := range tests {
		result := IsAudio(tt.contentType)
		if result != tt.expected {
			t.Errorf("IsAudio(%s) = %v, expected %v", tt.contentType, result, tt.expected)
		}
	}
}

func TestIsImage(t *testing.T) {
	tests := []struct {
		contentType string
		expected    bool
	}{
		{"image/png", true},
		{"image/jpeg", true},
		{"video/mp4", false},
		{"audio/mp3", false},
	}

	for _, tt := range tests {
		result := IsImage(tt.contentType)
		if result != tt.expected {
			t.Errorf("IsImage(%s) = %v, expected %v", tt.contentType, result, tt.expected)
		}
	}
}
