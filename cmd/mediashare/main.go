// Package main provides the entry point for the MediaShare service.
package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/ritlug/teleirc/internal/mediashare"
)

func main() {
	// Command line flags
	port := flag.Int("port", 0, "Port to listen on (overrides MEDIASHARE_PORT)")
	apiKey := flag.String("key", "", "API key for uploads (overrides MEDIASHARE_API_KEY)")
	baseURL := flag.String("url", "", "Base URL for generated links (overrides MEDIASHARE_BASE_URL)")
	storagePath := flag.String("storage", "", "Path to store files (overrides MEDIASHARE_STORAGE_PATH)")
	dbPath := flag.String("db", "", "Path to SQLite database (overrides MEDIASHARE_DB_PATH)")
	maxSize := flag.Int64("maxsize", 0, "Maximum file size in bytes (overrides MEDIASHARE_MAX_FILE_SIZE)")
	retention := flag.Int("retention", 0, "File retention in hours (overrides MEDIASHARE_RETENTION_HOURS)")
	serviceName := flag.String("name", "", "Service name for branding (overrides MEDIASHARE_SERVICE_NAME)")
	language := flag.String("lang", "", "Language code pl/en (overrides MEDIASHARE_LANGUAGE)")
	flag.Parse()

	// Load configuration from environment with flag overrides
	config := &mediashare.Config{
		Port:           getEnvInt("MEDIASHARE_PORT", 8080),
		APIKey:         getEnv("MEDIASHARE_API_KEY", ""),
		BaseURL:        getEnv("MEDIASHARE_BASE_URL", "http://localhost:8080"),
		StoragePath:    getEnv("MEDIASHARE_STORAGE_PATH", "./uploads"),
		DBPath:         getEnv("MEDIASHARE_DB_PATH", "./mediashare.db"),
		MaxFileSize:    getEnvInt64("MEDIASHARE_MAX_FILE_SIZE", 50*1024*1024), // 50MB default
		RetentionHours: getEnvInt("MEDIASHARE_RETENTION_HOURS", 72),
		ServiceName:    getEnv("MEDIASHARE_SERVICE_NAME", "MediaShare"),
		Language:       getEnv("MEDIASHARE_LANGUAGE", "pl"),
	}

	// Apply flag overrides
	if *port != 0 {
		config.Port = *port
	}
	if *apiKey != "" {
		config.APIKey = *apiKey
	}
	if *baseURL != "" {
		config.BaseURL = *baseURL
	}
	if *storagePath != "" {
		config.StoragePath = *storagePath
	}
	if *dbPath != "" {
		config.DBPath = *dbPath
	}
	if *maxSize != 0 {
		config.MaxFileSize = *maxSize
	}
	if *retention != 0 {
		config.RetentionHours = *retention
	}
	if *serviceName != "" {
		config.ServiceName = *serviceName
	}
	if *language != "" {
		config.Language = *language
	}

	// Ensure storage directory exists
	if err := os.MkdirAll(config.StoragePath, 0755); err != nil {
		log.Fatalf("Failed to create storage directory: %v", err)
	}

	// Start server
	server, err := mediashare.NewServer(config)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	if err := server.Start(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

func getEnvInt64(key string, defaultValue int64) int64 {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.ParseInt(value, 10, 64); err == nil {
			return intVal
		}
	}
	return defaultValue
}
