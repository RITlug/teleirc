package mediashare

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func setupTestServer(t *testing.T) (*Server, string) {
	tmpDir, err := os.MkdirTemp("", "mediashare_server_test")
	if err != nil {
		t.Fatal(err)
	}

	dbPath := tmpDir + "/test.db"
	config := &Config{
		Port:           8080,
		APIKey:         "testkey",
		BaseURL:        "http://localhost:8080",
		StoragePath:    tmpDir,
		DBPath:         dbPath,
		MaxFileSize:    10 * 1024 * 1024,
		RetentionHours: 72,
		ServiceName:    "MediaShareTest",
		Language:       "en",
	}

	server, err := NewServer(config)
	if err != nil {
		os.RemoveAll(tmpDir)
		t.Fatal(err)
	}
	return server, tmpDir
}

func TestHealthEndpoint(t *testing.T) {
	server, tmpDir := setupTestServer(t)
	defer os.RemoveAll(tmpDir)

	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response map[string]string
	json.NewDecoder(w.Body).Decode(&response)

	if response["status"] != "ok" {
		t.Errorf("Expected status ok, got %s", response["status"])
	}
}

func TestUploadWithoutAPIKey(t *testing.T) {
	server, tmpDir := setupTestServer(t)
	defer os.RemoveAll(tmpDir)

	// Create multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "test.txt")
	part.Write([]byte("test content"))
	writer.Close()

	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", w.Code)
	}
}

func TestUploadWithAPIKey(t *testing.T) {
	server, tmpDir := setupTestServer(t)
	defer os.RemoveAll(tmpDir)

	// Create multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "test.txt")
	part.Write([]byte("test content"))
	writer.Close()

	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-API-Key", "testkey")
	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response UploadResponse
	json.NewDecoder(w.Body).Decode(&response)

	if !response.Success {
		t.Error("Expected success to be true")
	}

	if response.ID == "" {
		t.Error("Expected ID to be set")
	}

	if response.URL == "" {
		t.Error("Expected URL to be set")
	}
}

func TestUploadWithQueryKey(t *testing.T) {
	server, tmpDir := setupTestServer(t)
	defer os.RemoveAll(tmpDir)

	// Create multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "test.txt")
	part.Write([]byte("test content"))
	writer.Close()

	req := httptest.NewRequest("POST", "/upload?key=testkey", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestViewEndpoint(t *testing.T) {
	server, tmpDir := setupTestServer(t)
	defer os.RemoveAll(tmpDir)

	// First upload a file
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "test.mp4")
	part.Write([]byte("fake video content"))
	writer.Close()

	uploadReq := httptest.NewRequest("POST", "/upload", body)
	uploadReq.Header.Set("Content-Type", writer.FormDataContentType())
	uploadReq.Header.Set("X-API-Key", "testkey")
	uploadW := httptest.NewRecorder()

	server.ServeHTTP(uploadW, uploadReq)

	var uploadResponse UploadResponse
	json.NewDecoder(uploadW.Body).Decode(&uploadResponse)

	// Now view it (new URL pattern: /{id} instead of /v/{id})
	viewReq := httptest.NewRequest("GET", "/"+uploadResponse.ID, nil)
	viewW := httptest.NewRecorder()

	server.ServeHTTP(viewW, viewReq)

	if viewW.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", viewW.Code)
	}

	contentType := viewW.Header().Get("Content-Type")
	if contentType != "text/html; charset=utf-8" {
		t.Errorf("Expected HTML content type, got %s", contentType)
	}
}

func TestRawEndpoint(t *testing.T) {
	server, tmpDir := setupTestServer(t)
	defer os.RemoveAll(tmpDir)

	// First upload a file
	fileContent := "raw file content"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "test.txt")
	part.Write([]byte(fileContent))
	writer.Close()

	uploadReq := httptest.NewRequest("POST", "/upload", body)
	uploadReq.Header.Set("Content-Type", writer.FormDataContentType())
	uploadReq.Header.Set("X-API-Key", "testkey")
	uploadW := httptest.NewRecorder()

	server.ServeHTTP(uploadW, uploadReq)

	var uploadResponse UploadResponse
	json.NewDecoder(uploadW.Body).Decode(&uploadResponse)

	// Now get raw
	rawReq := httptest.NewRequest("GET", "/r/"+uploadResponse.ID, nil)
	rawW := httptest.NewRecorder()

	server.ServeHTTP(rawW, rawReq)

	if rawW.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rawW.Code)
	}

	responseBody, _ := io.ReadAll(rawW.Body)
	if string(responseBody) != fileContent {
		t.Errorf("Expected %q, got %q", fileContent, string(responseBody))
	}
}

func TestNotFound(t *testing.T) {
	server, tmpDir := setupTestServer(t)
	defer os.RemoveAll(tmpDir)
	defer server.Close()

	// Test with valid-looking ID that doesn't exist (5 alphanumeric chars)
	req := httptest.NewRequest("GET", "/AbCd1", nil)
	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

func TestListEndpoint(t *testing.T) {
	// Create test server with ShowList enabled
	tmpDir, err := os.MkdirTemp("", "mediashare_list_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	config := &Config{
		Port:           8080,
		APIKey:         "testkey",
		BaseURL:        "http://localhost:8080",
		StoragePath:    tmpDir,
		DBPath:         tmpDir + "/test.db",
		MaxFileSize:    10 * 1024 * 1024,
		RetentionHours: 72,
		ServiceName:    "MediaShareTest",
		Language:       "en",
		ShowList:       true,
	}

	server, err := NewServer(config)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	// First upload a file
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "test.mp4")
	part.Write([]byte("fake video content"))
	writer.WriteField("username", "testuser")
	writer.Close()

	uploadReq := httptest.NewRequest("POST", "/upload", body)
	uploadReq.Header.Set("Content-Type", writer.FormDataContentType())
	uploadReq.Header.Set("X-API-Key", "testkey")
	uploadW := httptest.NewRecorder()

	server.ServeHTTP(uploadW, uploadReq)

	// Now check the list page
	listReq := httptest.NewRequest("GET", "/", nil)
	listW := httptest.NewRecorder()

	server.ServeHTTP(listW, listReq)

	if listW.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", listW.Code)
	}

	contentType := listW.Header().Get("Content-Type")
	if contentType != "text/html; charset=utf-8" {
		t.Errorf("Expected HTML content type, got %s", contentType)
	}

	// Check that response contains the uploaded file info
	responseBody := listW.Body.String()
	if !strings.Contains(responseBody, "test.mp4") {
		t.Error("Expected list page to contain uploaded filename")
	}
	if !strings.Contains(responseBody, "testuser") {
		t.Error("Expected list page to contain username")
	}
}
