package mediashare

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

// Config holds the server configuration.
type Config struct {
	Port           int
	APIKey         string
	BaseURL        string
	StoragePath    string
	DBPath         string
	MaxFileSize    int64
	RetentionHours int
	ServiceName    string
	Language       string
}

// Server is the HTTP server for MediaShare.
type Server struct {
	config  *Config
	storage *Storage
	db      *Database
	i18n    *I18n
	mux     *http.ServeMux
}

// UploadResponse is returned after a successful upload.
type UploadResponse struct {
	Success  bool   `json:"success"`
	URL      string `json:"url"`
	RawURL   string `json:"raw_url"`
	ID       string `json:"id"`
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
}

// ErrorResponse is returned on errors.
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// idPattern matches 5-character alphanumeric IDs
var idPattern = regexp.MustCompile(`^[a-zA-Z0-9]{5}$`)

// NewServer creates a new MediaShare server.
func NewServer(config *Config) (*Server, error) {
	// Initialize database
	db, err := NewDatabase(config.DBPath)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	// Initialize i18n
	i18n := NewI18n(config.Language)

	s := &Server{
		config:  config,
		db:      db,
		storage: NewStorage(config.StoragePath, config.MaxFileSize, db),
		i18n:    i18n,
		mux:     http.NewServeMux(),
	}
	s.setupRoutes()

	// Start cleanup worker
	if config.RetentionHours > 0 {
		s.startCleanupWorker()
	}

	return s, nil
}

func (s *Server) setupRoutes() {
	s.mux.HandleFunc("/upload", s.handleUpload)
	s.mux.HandleFunc("/r/", s.handleRaw)
	s.mux.HandleFunc("/health", s.handleHealth)
	s.mux.HandleFunc("/", s.handleRoot)
}

// ServeHTTP implements http.Handler.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Add CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-API-Key")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	s.mux.ServeHTTP(w, r)
}

// Start begins listening on the configured port.
func (s *Server) Start() error {
	addr := fmt.Sprintf(":%d", s.config.Port)
	log.Printf("[%s] Server starting on %s", s.config.ServiceName, addr)
	log.Printf("[%s] Base URL: %s", s.config.ServiceName, s.config.BaseURL)
	log.Printf("[%s] Storage: %s", s.config.ServiceName, s.config.StoragePath)
	log.Printf("[%s] Language: %s", s.config.ServiceName, s.config.Language)
	log.Printf("[%s] Retention: %d hours", s.config.ServiceName, s.config.RetentionHours)
	return http.ListenAndServe(addr, s)
}

// Close cleans up resources.
func (s *Server) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"service": s.config.ServiceName,
	})
}

func (s *Server) handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		s.sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check API key
	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		apiKey = r.URL.Query().Get("key")
	}
	if s.config.APIKey != "" && apiKey != s.config.APIKey {
		s.sendError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Parse multipart form
	if err := r.ParseMultipartForm(s.config.MaxFileSize); err != nil {
		s.sendError(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		s.sendError(w, "No file provided", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Get username from form (optional)
	username := r.FormValue("username")

	// Store the file
	info, err := s.storage.Store(header.Filename, username, file)
	if err != nil {
		s.sendError(w, "Failed to store file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Build response URLs
	baseURL := strings.TrimSuffix(s.config.BaseURL, "/")
	response := UploadResponse{
		Success:  true,
		URL:      fmt.Sprintf("%s/%s", baseURL, info.ID),
		RawURL:   fmt.Sprintf("%s/r/%s", baseURL, info.ID),
		ID:       info.ID,
		Filename: info.Filename,
		Size:     info.Size,
	}

	log.Printf("[%s] Uploaded: %s (%s, %d bytes, user: %s)",
		s.config.ServiceName, info.ID, info.Filename, info.Size, username)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		s.sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Root path - show file list
	if r.URL.Path == "/" {
		s.handleList(w, r)
		return
	}

	// Extract ID from path: /{id}
	id := strings.TrimPrefix(r.URL.Path, "/")

	// Validate ID format (5 alphanumeric characters)
	if !idPattern.MatchString(id) {
		s.sendNotFound(w)
		return
	}

	// Get file info
	info, err := s.storage.Get(id)
	if err != nil || info == nil {
		s.sendNotFound(w)
		return
	}

	// Update last opened timestamp
	if err := s.storage.UpdateLastOpened(id); err != nil {
		log.Printf("[%s] Failed to update last opened: %v", s.config.ServiceName, err)
	}

	// Prepare template data
	baseURL := strings.TrimSuffix(s.config.BaseURL, "/")
	data := PageData{
		ID:          info.ID,
		Filename:    info.Filename,
		ContentType: info.ContentType,
		Size:        info.Size,
		IsVideo:     IsVideo(info.ContentType),
		IsAudio:     IsAudio(info.ContentType),
		IsImage:     IsImage(info.ContentType),
		RawURL:      fmt.Sprintf("%s/r/%s", baseURL, info.ID),
		Username:    info.Username,
		UploadedAt:  info.UploadedAt,
		ServiceName: s.config.ServiceName,
		Lang:        s.i18n.Lang(),
		T:           s.i18n.GetTranslations(),
		BaseURL:     baseURL,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := PlayerTemplate.Execute(w, data); err != nil {
		log.Printf("[%s] Template error: %v", s.config.ServiceName, err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (s *Server) handleList(w http.ResponseWriter, r *http.Request) {
	// Get recent files from database (last 50)
	records, err := s.db.GetRecentFiles(50)
	if err != nil {
		log.Printf("[%s] Failed to get recent files: %v", s.config.ServiceName, err)
		records = []*FileRecord{}
	}

	// Convert to template data
	baseURL := strings.TrimSuffix(s.config.BaseURL, "/")
	files := make([]FileListItem, 0, len(records))

	for _, rec := range records {
		// Determine content type category
		contentType := "file"
		if IsVideo(rec.ContentType) {
			contentType = "video"
		} else if IsAudio(rec.ContentType) {
			contentType = "audio"
		} else if IsImage(rec.ContentType) {
			contentType = "image"
		}

		// Format last opened
		lastOpened := s.i18n.T("never")
		if rec.LastOpenedAt != nil {
			lastOpened = rec.LastOpenedAt.Format("2006-01-02 15:04")
		}

		// Format username
		username := rec.Username
		if username == "" {
			username = s.i18n.T("anonymous")
		}

		files = append(files, FileListItem{
			ID:           rec.ID,
			Filename:     rec.Filename,
			ContentType:  contentType,
			Username:     username,
			UploadedAt:   rec.UploadedAt.Format("2006-01-02 15:04"),
			LastOpenedAt: lastOpened,
			URL:          fmt.Sprintf("%s/%s", baseURL, rec.ID),
			RawURL:       fmt.Sprintf("%s/r/%s", baseURL, rec.ID),
		})
	}

	data := ListPageData{
		ServiceName: s.config.ServiceName,
		Lang:        s.i18n.Lang(),
		T:           s.i18n.GetTranslations(),
		Files:       files,
		BaseURL:     baseURL,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := ListTemplate.Execute(w, data); err != nil {
		log.Printf("[%s] List template error: %v", s.config.ServiceName, err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (s *Server) handleRaw(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		s.sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from path: /r/{id}
	id := strings.TrimPrefix(r.URL.Path, "/r/")
	if id == "" || !idPattern.MatchString(id) {
		s.sendNotFound(w)
		return
	}

	// Get file info
	info, err := s.storage.Get(id)
	if err != nil || info == nil {
		s.sendNotFound(w)
		return
	}

	// Update last opened timestamp
	if err := s.storage.UpdateLastOpened(id); err != nil {
		log.Printf("[%s] Failed to update last opened: %v", s.config.ServiceName, err)
	}

	// Serve the file
	fullPath := s.storage.GetFullPath(info.Path)
	w.Header().Set("Content-Type", info.ContentType)
	w.Header().Set("Content-Disposition", fmt.Sprintf(`inline; filename="%s"`, info.Filename))
	http.ServeFile(w, r, fullPath)
}

func (s *Server) sendError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{
		Success: false,
		Error:   message,
	})
}

func (s *Server) sendNotFound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)

	data := NotFoundData{
		ServiceName: s.config.ServiceName,
		Lang:        s.i18n.Lang(),
		T:           s.i18n.GetTranslations(),
	}
	NotFoundTemplate.Execute(w, data)
}

// startCleanupWorker starts a background goroutine that periodically cleans up expired files.
func (s *Server) startCleanupWorker() {
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		// Run cleanup immediately on start
		s.cleanupExpiredFiles()

		for range ticker.C {
			s.cleanupExpiredFiles()
		}
	}()
	log.Printf("[%s] Cleanup worker started (every 1 hour, retention: %d hours)",
		s.config.ServiceName, s.config.RetentionHours)
}

// cleanupExpiredFiles removes files that haven't been opened within the retention period.
func (s *Server) cleanupExpiredFiles() {
	files, err := s.storage.GetExpiredFiles(s.config.RetentionHours)
	if err != nil {
		log.Printf("[%s] Cleanup error: %v", s.config.ServiceName, err)
		return
	}

	if len(files) == 0 {
		return
	}

	log.Printf("[%s] Cleaning up %d expired files...", s.config.ServiceName, len(files))

	for _, f := range files {
		// Remove file from disk
		fullPath := s.storage.GetFullPath(f.Path)
		if err := os.Remove(fullPath); err != nil && !os.IsNotExist(err) {
			log.Printf("[%s] Failed to remove file %s: %v", s.config.ServiceName, f.ID, err)
			continue
		}

		// Remove from database
		if err := s.db.Delete(f.ID); err != nil {
			log.Printf("[%s] Failed to delete record %s: %v", s.config.ServiceName, f.ID, err)
			continue
		}

		log.Printf("[%s] Cleaned up: %s (%s)", s.config.ServiceName, f.ID, f.Filename)
	}
}

// GetI18n returns the i18n instance for use by external code (like TeleIRC).
func (s *Server) GetI18n() *I18n {
	return s.i18n
}
